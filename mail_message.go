package mail

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/hexya-addons/bus/bustypes"
	"github.com/hexya-addons/mail/mailtypes"
	"github.com/hexya-addons/web/webtypes"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/models/types/dates"
	"github.com/hexya-erp/hexya/src/templates"
	"github.com/hexya-erp/hexya/src/tools/emailutils"
	"github.com/hexya-erp/hexya/src/tools/strutils"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

var imageDataURL = regexp.MustCompile(`(data:image/[a-z]+?);base64,([a-z0-9+/\n]{3,}=*)\n*(['"])(?: data-filename="([^"]*)")?`)

func getDefaultFrom(env models.Environment) interface{} {
	currentUser := h.User().NewSet(env).CurrentUser()
	if currentUser.Email() != "" {
		addr := mail.Address{Name: currentUser.Name(), Address: currentUser.Email()}
		return addr.String()
	}
	panic(h.User().NewSet(env).T("Unable to post message, please configure the sender's email address."))
}

var fields_MailMessage = map[string]models.FieldDefinition{
	"Subject": fields.Char{},
	"Date": fields.DateTime{
		Default: func(env models.Environment) interface{} { return dates.Now() },
	},
	"Body": fields.HTML{
		String:  "Contents",
		Default: models.DefaultValue(""),
		// sanitize_style=True
	},
	"Attachments": fields.Many2Many{
		RelationModel:    h.Attachment(),
		M2MLinkModelName: "MessageAttachmentRel",
		M2MOurField:      "Message",
		M2MTheirField:    "Attachment",
		Help: "Attachments are linked to a document through model / res_id" +
			"and to the message through this field.",
	},
	"Parent": fields.Many2One{
		RelationModel: h.MailMessage(),
		String:        "Parent Message",
		Index:         true,
		OnDelete:      models.SetNull,
		Help:          "Initial thread message.",
	},
	"Children": fields.One2Many{
		RelationModel: h.MailMessage(),
		ReverseFK:     "Parent",
		String:        "Child Messages",
	},
	"ResModel": fields.Char{
		String: "Related Document Model",
		Index:  true,
	},
	"ResID": fields.Integer{
		String: "Related Document ID",
		Index:  true,
	},
	"RecordName": fields.Char{
		String: "Message Record Name",
		Help:   "Name get of the related document.",
	},
	"MessageType": fields.Selection{
		Selection: types.Selection{
			"email":        "Email",
			"comment":      "Comment",
			"notification": "System notification",
		},
		String:   "Type",
		Required: true,
		Default:  models.DefaultValue("email"),
		Help: "Message type: email for email message, notification for" +
			"system message, comment for other messages such as user replies",
	},
	"Subtype": fields.Many2One{
		RelationModel: h.MailMessageSubtype(),
		String:        "Subtype",
		OnDelete:      models.SetNull,
		Index:         true,
	},
	"MailActivityType": fields.Many2One{
		RelationModel: h.MailActivityType(),
		String:        "Mail Activity Type",
		Index:         true,
		OnDelete:      models.SetNull,
	},
	"EmailFrom": fields.Char{
		String:  "From",
		Default: getDefaultFrom,
		Help: "Email address of the sender. This field is set when no" +
			"matching partner is found and replaces the author_id field" +
			"in the chatter.",
	},
	"Author": fields.Many2One{
		RelationModel: h.Partner(),
		String:        "Author",
		Index:         true,
		OnDelete:      models.SetNull,
		Default: func(env models.Environment) interface{} {
			return h.User().NewSet(env).CurrentUser().Partner()
		},
		Help: "Author of the message. If not set, email_from may hold" +
			"an email address that did not match any partner.",
	},
	"AuthorAvatar": fields.Binary{
		String:   "Author's avatar",
		Related:  `AuthorId.ImageSmall`,
		ReadOnly: false,
	},

	// recipients: include inactive partners (they may have been archived after
	// the message was sent, but they should remain visible in the relation)
	"Partners": fields.Many2Many{
		RelationModel: h.Partner(),
		String:        "Recipients",
		Filter:        q.Partner().Active().Equals(true).Or().Active().Equals(false),
	},
	"NeedactionPartners": fields.Many2Many{
		RelationModel:    h.Partner(),
		M2MLinkModelName: "MailMessagePartnerNeedactionRel",
		String:           "Partners with Need Action",
		Filter:           q.Partner().Active().Equals(true).Or().Active().Equals(false),
	},
	"Needaction": fields.Boolean{
		String:  "Need Action",
		Compute: h.MailMessage().Methods().ComputeNeedaction(),
		// search='_search_needaction'
		Help: "Need Action",
	},
	"HasError": fields.Boolean{
		String:  "Has error",
		Compute: h.MailMessage().Methods().ComputeHasError(),
		// search='_search_has_error'
		Help: "Has error",
	},
	"Channels": fields.Many2Many{
		RelationModel:    h.MailChannel(),
		M2MLinkModelName: "MailMessageMailChannelRel",
	},

	// Notifications
	"Notifications": fields.One2Many{
		RelationModel: h.MailNotification(),
		ReverseFK:     "MailMessage",
	},

	// User interface
	"StarredPartners": fields.Many2Many{
		RelationModel:    h.Partner(),
		M2MLinkModelName: "MailMessagePartnerStarredRel",
		String:           "Favorited By",
	},
	"Starred": fields.Boolean{
		String:  "Starred",
		Compute: h.MailMessage().Methods().ComputeStarred(),
		Depends: []string{"StarredPartners"},
		// search='_search_starred'
		Help: "Current user has a starred notification linked to this message",
	},

	// Tracking
	"TrackingValues": fields.One2Many{
		RelationModel: h.MailTrackingValue(),
		ReverseFK:     "MailMessage",
		String:        "Tracking values",
		// groups="base.group_no_one"
		Help: "Tracked values are stored in a separate model. This field" +
			"allow to reconstruct the tracking and to generate statistics" +
			"on the model.",
	},

	// Mail Gateway
	"NoAutoThread": fields.Boolean{
		String: "No threading for answers",
		Help: "Answers do not go in the original document discussion thread." +
			"This has an impact on the generated message-id.",
	},
	"MessageID": fields.Char{
		String:   "Message-ID",
		Help:     "Message unique identifier",
		Index:    true,
		ReadOnly: true,
		NoCopy:   true,
	},
	"ReplyTo": fields.Char{
		String: "Reply-To",
		Help: "Reply email address. Setting the reply_to bypasses the" +
			"automatic thread creation.",
	},
	"MailServer": fields.Many2One{
		RelationModel: h.MailServer(),
		String:        "Outgoing mail server",
	},

	// Moderation
	"ModerationStatus": fields.Selection{
		Selection: types.Selection{
			"pending_moderation": "Pending Moderation",
			"accepted":           "Accepted",
			"rejected":           "Rejected",
		},
		String: "Moderation Status",
		Index:  true,
	},
	"Moderator": fields.Many2One{
		RelationModel: h.User(),
		String:        "Moderated By",
		Index:         true,
	},
	"NeedModeration": fields.Boolean{
		String:  "Need moderation",
		Compute: h.MailMessage().Methods().ComputeNeedModeration(),
		// search='_search_need_moderation'
	},

	// keep notification layout informations to be able to generate mail again
	"Layout": fields.Char{
		String: "Layout",
		NoCopy: true,
	},
	"AddSign": fields.Boolean{
		Default: models.DefaultValue(true),
	},
}

// ComputeNeedaction computes if there is a Need for action on a MailMessage,
// i.e. notified on my channel
func mailMessage_ComputeNeedaction(rs m.MailMessageSet) m.MailMessageData {
	res := h.MailMessage().NewData()
	myNotifications := h.MailNotification().NewSet(rs.Env()).Sudo().Search(q.MailNotification().
		MailMessage().In(rs).And().
		Partner().Equals(h.User().NewSet(rs.Env()).CurrentUser().Partner()).And().
		IsRead().Equals(false))
	for _, notif := range myNotifications.Records() {
		if notif.MailMessage().Equals(rs) {
			return res.SetNeedaction(true)
		}
	}
	return res.SetNeedaction(false)
}

// SearchNeedaction
func mailMessage_SearchNeedaction(rs m.MailMessageSet, operator interface{}, operand interface{}) {
	//        if operator == '=' and operand:
	//            return ['&', ('notification_ids.res_partner_id', '=', self.env.user.partner_id.id), ('notification_ids.is_read', '=', False)]
	//        return ['&', ('notification_ids.res_partner_id', '=', self.env.user.partner_id.id), ('notification_ids.is_read', '=', True)]
}

// ComputeHasError computes the HasError field from notifications
func mailMessage_ComputeHasError(rs m.MailMessageSet) m.MailMessageData {
	res := h.MailMessage().NewData()
	errorNotifications := h.MailNotification().NewSet(rs.Env()).Sudo().Search(q.MailNotification().
		MailMessage().In(rs).And().
		EmailStatus().In([]string{"bounce", "exception"}))
	for _, notif := range errorNotifications.Records() {
		if notif.MailMessage().Equals(rs) {
			return res.SetHasError(true)
		}
	}
	return res.SetHasError(false)
}

// SearchHasError
func mailMessage_SearchHasError(rs m.MailMessageSet, operator interface{}, operand interface{}) {
	//        if operator == '=' and operand:
	//            return [('notification_ids.email_status', 'in', ('bounce', 'exception'))]
	//        return ['!', ('notification_ids.email_status', 'in', ('bounce', 'exception'))]
}

// ComputeStarred computes if the message is starred by the current user.
func mailMessage_ComputeStarred(rs m.MailMessageSet) m.MailMessageData {
	res := h.MailMessage().NewData()
	isStarred := h.User().NewSet(rs.Env()).CurrentUser().Partner().Intersect(rs.StarredPartners()).IsNotEmpty()
	return res.SetStarred(isStarred)
}

// SearchStarred
func mailMessage_SearchStarred(rs m.MailMessageSet, operator interface{}, operand interface{}) {
	//        if operator == '=' and operand:
	//            return [('starred_partner_ids', 'in', [self.env.user.partner_id.id])]
	//        return [('starred_partner_ids', 'not in', [self.env.user.partner_id.id])]
}

// ComputeNeedModeration computes if this message needs moderation
func mailMessage_ComputeNeedModeration(_ m.MailMessageSet) m.MailMessageData {
	res := h.MailMessage().NewData().SetNeedModeration(false)
	return res
}

// SearchNeedModeration
func mailMessage_SearchNeedModeration(rs m.MailMessageSet, operator interface{}, operand interface{}) {
	//        if operator == '=' and operand is True:
	//            return ['&', '&',
	//                    ('moderation_status', '=', 'pending_moderation'),
	//                    ('model', '=', 'mail.channel'),
	//                    ('res_id', 'in', self.env.user.moderation_channel_ids.ids)]
	//        return ValueError(_('Unsupported search filter on moderation status'))
}

// ------------------------------------------------------
// Notification API
// ------------------------------------------------------

// MarkAllAsRead removes all needactions of the current partner.
// If channels is given, restrict to messages written in one of those channels.
func mailMessage_MarkAllAsRead(rs m.MailMessageSet, channels m.MailChannelSet, domain q.MailMessageCondition) []int64 {
	var ids []int64
	currentUser := h.User().NewSet(rs.Env()).CurrentUser()
	partner := currentUser.Partner()
	deleteMode := !currentUser.Share()
	switch {
	case domain.IsEmpty() && deleteMode:
		query := `DELETE FROM mail_message_partner_needaction_rel WHERE partner_id IN (?)`
		args := []interface{}{partner.ID()}
		if channels.IsNotEmpty() {
			query += `
AND mail_message_id in
	(SELECT mail_message_id
	FROM mail_message_mail_channel_rel
	WHERE mail_channel_id in (?))`
			args = append(args, channels.Ids())
		}
		query += ` RETURNING mail_message_id as id`

		rs.Env().Cr().Select(&ids, query, args)
	default:
		// not really efficient method: it does one db request for the
		// search, and one for each message in the result set to remove the
		// current user from the relation.
		msgCond := q.MailMessage().NeedactionPartners().In(partner)
		if channels.IsNotEmpty() {
			msgCond.And().Channels().In(channels)
		}
		unreadMessages := h.MailMessage().Search(rs.Env(), msgCond.AndCond(domain))
		notifications := h.MailNotification().NewSet(rs.Env()).Sudo().Search(q.MailNotification().
			MailMessage().In(unreadMessages).And().
			Partner().Equals(partner).And().
			IsRead().Equals(false))
		switch deleteMode {
		case true:
			notifications.Unlink()
		case false:
			notifications.SetIsRead(true)
		}
		ids = unreadMessages.Ids()
	}

	notification := map[string]interface{}{
		"type":        "mark_as_read",
		"message_ids": ids,
		"channel_ids": channels.Ids(),
	}
	h.BusBus().NewSet(rs.Env()).Sendone(fmt.Sprintf("res.partner.%d", partner.ID()), notification)
	return ids
}

// SetMessageDone removes the needaction from messages for the current partner.
func mailMessage_SetMessageDone(rs m.MailMessageSet) {
	currentUser := h.User().NewSet(rs.Env()).CurrentUser()
	partner := currentUser.Partner()
	deleteMode := !currentUser.Share()
	notifications := h.MailNotification().NewSet(rs.Env()).Sudo().Search(q.MailNotification().
		MailMessage().In(rs).And().
		Partner().Equals(partner).And().
		IsRead().Equals(false))
	if notifications.IsEmpty() {
		return
	}
	messages := h.MailMessage().NewSet(rs.Env())
	for _, notif := range notifications.Records() {
		messages = messages.Union(notif.MailMessage())
	}
	type group struct {
		messages m.MailMessageSet
		channels m.MailChannelSet
	}
	var groups []group
	currentChannels := messages.Records()[0].Channels()

	currentGroup := h.MailMessage().NewSet(rs.Env())
	for _, record := range messages.Records() {
		if record.Channels().Equals(currentChannels) {
			currentGroup = currentGroup.Union(record)
			continue
		}
		groups = append(groups, group{
			messages: currentGroup,
			channels: currentChannels,
		})
		currentGroup = record
		currentChannels = record.Channels()
	}
	groups = append(groups, group{
		messages: currentGroup,
		channels: currentChannels,
	})

	switch deleteMode {
	case true:
		notifications.Unlink()
	case false:
		notifications.SetIsRead(true)
	}

	for _, grp := range groups {
		notification := map[string]interface{}{
			"type":        "mark_as_read",
			"message_ids": grp.messages.Ids(),
			"channel_ids": grp.channels.Ids(),
		}
		h.BusBus().NewSet(rs.Env()).Sendone(fmt.Sprintf("res.partner.%d", partner.ID()), notification)
	}
}

//  Unstar messages for the current partner.
func mailMessage_UnstarAll(rs m.MailMessageSet) {
	currentUser := h.User().NewSet(rs.Env()).CurrentUser()
	partner := currentUser.Partner()
	starredMessages := h.MailMessage().Search(rs.Env(), q.MailMessage().StarredPartners().In(partner))
	for _, starredMsg := range starredMessages.Records() {
		starredMsg.SetStarredPartners(starredMsg.Partners().Subtract(partner))
	}
	notification := map[string]interface{}{
		"type":        "toggle_star",
		"message_ids": starredMessages.Ids(),
		"starred":     false,
	}
	h.BusBus().NewSet(rs.Env()).Sendone(fmt.Sprintf("res.partner.%d", partner.ID()), notification)
}

// ToggleMessageStarred toggles messages as (un)starred.
// Technically, the notifications related to uid are set to (un)starred.
func mailMessage_ToggleMessageStarred(rs m.MailMessageSet) {
	rs.EnsureOne()
	rs.CheckAccessRule("read")
	currentUser := h.User().NewSet(rs.Env()).CurrentUser()
	partner := currentUser.Partner()
	starred := !rs.Starred()
	switch starred {
	case true:
		rs.Sudo().SetStarredPartners(rs.StarredPartners().Union(partner))
	case false:
		rs.Sudo().SetStarredPartners(rs.StarredPartners().Subtract(partner))
	}
	notification := map[string]interface{}{
		"type":        "toggle_star",
		"message_ids": rs.Ids(),
		"starred":     starred,
	}
	h.BusBus().NewSet(rs.Env()).Sendone(fmt.Sprintf("res.partner.%d", partner.ID()), notification)
}

// ------------------------------------------------------
// Message loading for web interface
// ------------------------------------------------------

// MessagePostprocess applies post-processing on the given messages.
// messages must be values taken from rs.
// This method will handle partners in batch to avoid doing numerous queries.
func mailMessage_MessageReadDictPostprocess(rs m.MailMessageSet, messages *[]*mailtypes.MailMessageInfo) bool {
	// 1. Aggregate partners (author_id and partner_ids), attachments and tracking values
	partners := h.Partner().NewSet(rs.Env()).Sudo()
	attachements := h.Attachment().NewSet(rs.Env())
	for _, message := range rs.Sudo().Records() {
		if message.Author().IsNotEmpty() {
			partners = partners.Union(message.Author())
		}
		if message.Partners().IsNotEmpty() {
			partners = partners.Union(message.Partners())
		}
		if message.NeedactionPartners().IsNotEmpty() {
			partners = partners.Union(message.NeedactionPartners())
		}
		if message.Attachments().IsNotEmpty() {
			attachements = attachements.Union(message.Attachments())
		}
	}

	partnerTree := make(map[int64]webtypes.RecordIDWithName)
	for _, partner := range partners.Records() {
		partnerTree[partner.ID()] = webtypes.RecordIDWithName{
			ID:   partner.ID(),
			Name: partner.NameGet(),
		}
	}

	// 2. Attachments as SUPERUSER, because could receive msg and attachments for doc uid cannot see
	attachmentsData := attachements.Sudo().Load(
		q.Attachment().ID(), q.Attachment().DatasFname(), q.Attachment().Name(), q.Attachment().MimeType())
	attachmentsTree := make(map[int64]*mailtypes.MailAttachmentInfo)
	for _, att := range attachmentsData.Records() {
		attachmentsTree[att.ID()] = &mailtypes.MailAttachmentInfo{
			ID:       att.ID(),
			Filename: att.DatasFname(),
			Name:     att.Name(),
			Mimetype: att.MimeType(),
		}
	}

	// 3. Tracking values
	trackingValues := h.MailTrackingValue().NewSet(rs.Env()).Sudo().Search(q.MailTrackingValue().MailMessage().In(rs))
	messageToTracking := make(map[int64][]int64)
	trackingTree := make(map[int64]*mailtypes.TrackingData)
	for _, tracking := range trackingValues.Records() {
		trackingTree[tracking.ID()] = &mailtypes.TrackingData{
			ID:           tracking.ID(),
			ChangedField: tracking.FieldDesc(),
			OldValue:     tracking.GetOldDisplayValue(),
			NewValue:     tracking.GetNewDisplayValue(),
			FieldType:    tracking.FieldType(),
		}
	}

	// 4. Update message dictionaries
	for i, messageDict := range *messages {
		messageID := messageDict.ID
		message := rs.Filtered(func(r m.MailMessageSet) bool {
			return r.ID() == messageID
		})
		author := webtypes.RecordIDWithName{
			Name: message.EmailFrom(),
		}
		if message.Author().IsNotEmpty() {
			author = partnerTree[message.Author().ID()]
		}
		var partnerIds []webtypes.RecordIDWithName
		for _, partner := range message.Partners().Records() {
			if p, ok := partnerTree[partner.ID()]; ok {
				partnerIds = append(partnerIds, p)
			}
		}
		// we read customer_email_status before filtering inactive user because we don't want to miss a red enveloppe
		var customerEmailStatus string
		allSent := true
		for _, notif := range message.Notifications().Records() {
			switch notif.EmailStatus() {
			case "sent":
			case "exception":
				customerEmailStatus = "exception"
				break
			case "bounce":
				customerEmailStatus = "bounce"
				break
			default:
				allSent = false
			}
		}
		if customerEmailStatus == "" {
			customerEmailStatus = "ready"
			if allSent {
				customerEmailStatus = "sent"
			}
		}
		var customerEmailData []mailtypes.CustomerEmailInfo
		for _, notification := range message.Notifications().Filtered(func(r m.MailNotificationSet) bool {
			return (strutils.IsIn(r.EmailStatus(), "bounce", "exception", "canceled") ||
				r.Partner().PartnerShare()) &&
				r.Partner().Active()
		}).Records() {
			customerEmailData = append(customerEmailData, mailtypes.CustomerEmailInfo{
				ID:          partnerTree[notification.Partner().ID()].ID,
				Name:        partnerTree[notification.Partner().ID()].Name,
				EmailStatus: notification.EmailStatus(),
			})
		}

		var attachments []*mailtypes.MailAttachmentInfo
		if message.Attachments().IsNotEmpty() {
			var hasAccessToModel bool
			resModel := models.Registry.MustGet(message.ResModel())
			if message.ResModel() != "" {
				hasAccessToModel = rs.Env().Pool(message.ResModel()).CheckExecutionPermission(resModel.Methods().MustGet("Load"), true)
			}
			mainAttachment := h.Attachment().NewSet(rs.Env())
			if hasAccessToModel && message.ResID() != 0 &&
				resModel.Search(rs.Env(), resModel.Field(models.ID).Equals(message.ID())).SearchCount() > 0 {
				mainAttachment = resModel.BrowseOne(rs.Env(), message.ResID()).Wrap().(m.MailThreadSet).MessageMainAttachment()
			}
			for _, attachment := range message.Attachments().Records() {
				if _, ok := attachmentsTree[attachment.ID()]; ok {
					attachmentsTree[attachment.ID()].IsMain = attachment.Equals(mainAttachment)
				}
				attachments = append(attachments, attachmentsTree[attachment.ID()])
			}
		}
		var trackingVals []*mailtypes.TrackingData
		for _, trackingValueID := range messageToTracking[messageID] {
			if _, ok := trackingTree[trackingValueID]; ok {
				trackingVals = append(trackingVals, trackingTree[trackingValueID])
			}
		}
		(*messages)[i].Author = author
		(*messages)[i].Partners = partnerIds
		(*messages)[i].CustomerEmailStatus = customerEmailStatus
		(*messages)[i].CustomerEmailData = customerEmailData
		(*messages)[i].Attachments = attachments
		(*messages)[i].TrackingValues = trackingVals
	}
	return true
}

// MessageFetchFailed returns the formatted mail failues
func mailMessage_MessageFetchFailed(rs m.MailMessageSet) []*mailtypes.MailFailureInfo {
	messages := h.MailMessage().Search(rs.Env(), q.MailMessage().
		HasError().Equals(true).And().
		Author().Equals(h.User().NewSet(rs.Env()).CurrentUser().Partner()))
	return messages.FormatMailFailures()
}

// MessageFetch gets a limited amount of formatted messages with provided domain.
//
// Parameters:
// - cond: the condition to filter messages;
// - limit: the maximum amount of messages to get;
// - moderatedChannels: if not empty, it contains a moderated channel. Fetched messages
// 						should include pending moderation messages for moderators. If the
// 						current user is not moderator, it should still get self-authored
// 						messages that are pending moderation;
func mailMessage_MessageFetch(rs m.MailMessageSet, cond q.MailMessageCondition, limit int, moderatedChannel m.MailChannelSet) []*mailtypes.MailMessageInfo {
	messages := h.MailMessage().Search(rs.Env(), cond).Limit(limit)
	if moderatedChannel.IsNotEmpty() {
		// Split load moderated and regular messages, as the ORed domain can
		// cause performance issues on large databases.
		moderatedMessageCond := q.MailMessage().
			ResModel().Equals("MailChannel").And().
			ResID().In(moderatedChannel.Ids()).AndCond(
			q.MailMessage().Author().Equals(h.User().NewSet(rs.Env()).CurrentUser().Partner()).Or().
				NeedModeration().Equals(true))
		messages = messages.Union(h.MailMessage().Search(rs.Env(), moderatedMessageCond).Limit(limit))
		// Truncate the results to limit
		messages = messages.SortedByField(q.MailMessage().ID(), true).Limit(limit)
	}
	return messages.MessageFormat()
}

// MessageFormat get the message values in the format for web client.
// Since message values can be broadcasted, computed fields MUST NOT BE READ and broadcasted.
func mailMessage_MessageFormat(rs m.MailMessageSet) []*mailtypes.MailMessageInfo {
	var messageValues []*mailtypes.MailMessageInfo
	for _, msg := range rs.Load(q.MailMessage().ID(), q.MailMessage().Body(), q.MailMessage().Date(),
		q.MailMessage().Author(), q.MailMessage().EmailFrom(),
		q.MailMessage().MessageType(), q.MailMessage().Subtype(), q.MailMessage().Subject(),
		q.MailMessage().ResModel(), q.MailMessage().ResID(), q.MailMessage().RecordName(),
		q.MailMessage().Channels(), q.MailMessage().Partners(),
		q.MailMessage().StarredPartners(),
		q.MailMessage().ModerationStatus(),
	).Records() {
		var partnersWithNames []webtypes.RecordIDWithName
		for _, p := range msg.Partners().Records() {
			partnersWithNames = append(partnersWithNames, webtypes.RecordIDWithName{
				ID:   p.ID(),
				Name: p.Name(),
			})
		}
		messageValues = append(messageValues, &mailtypes.MailMessageInfo{
			ID: msg.ID(), Body: msg.Body(), Date: msg.Date(), Author: webtypes.RecordIDWithName{ID: msg.Author().ID(), Name: msg.Author().Name()},
			EmailFrom:   msg.EmailFrom(),
			MessageType: msg.MessageType(), Subtype: webtypes.RecordIDWithName{ID: msg.Subtype().ID(), Name: msg.Subtype().Name()},
			Subject: msg.Subject(),
			Model:   msg.ResModel(), ResID: msg.ResID(), RecordName: msg.RecordName(),
			ChannelIds: msg.Channels().Ids(), Partners: partnersWithNames,
			StarredPartners:  msg.StarredPartners().Ids(),
			ModerationStatue: msg.ModerationStatus(),
		})
	}

	rs.MessageReadDictPostprocess(&messageValues)

	// Add subtype data (is_note flag, is_discussion flag , subtype_description). Do it as sudo
	// because portal / public may have to look for internal subtypes
	subTypes := h.MailMessageSubtype().NewSet(rs.Env())
	for _, r := range rs.Records() {
		subTypes = subTypes.Union(r.Sudo().Subtype())
	}

	comment := h.MailMessageSubtype().NewSet(rs.Env()).GetRecord("mail_mt_comment")
	note := h.MailMessageSubtype().NewSet(rs.Env()).GetRecord("mail_mt_note")

	// fetch notification status
	notifDict := make(map[int64][]int64)
	notifs := h.MailNotification().NewSet(rs.Env()).Sudo().Search(q.MailNotification().
		MailMessage().In(rs).And().
		IsRead().Equals(false))
	for _, notif := range notifs.Records() {
		notifDict[notif.MailMessage().ID()] = append(notifDict[notif.MailMessage().ID()], notif.Partner().ID())
	}

	for i, message := range messageValues {
		messageValues[i].NeedactionPartnerIds = notifDict[message.ID]
		messageValues[i].IsNote = message.Subtype.ID == note.ID()
		messageValues[i].IsDiscussion = message.Subtype.ID == comment.ID()
		messageValues[i].IsNotification = messageValues[i].IsNote && message.Model == "" && message.ResID == 0
		messageValues[i].SubtypeDescription = subTypes.Filtered(func(r m.MailMessageSubtypeSet) bool {
			return r.ID() == message.Subtype.ID
		}).Description()
	}
	return messageValues
}

// FormatMailFailures returns a shorter message to notify a failure update
func mailMessage_FormatMailFailures(rs m.MailMessageSet) []*mailtypes.MailFailureInfo {
	var failureInfos []*mailtypes.MailFailureInfo
	// For each channel, build the information header and include the logged partner information
	for _, message := range rs.Records() {
		if message.ResModel() != "" && message.ResID() != 0 {
			record := models.Registry.MustGet(message.ResModel()).BrowseOne(rs.Env(), message.ResID())
			if !record.CheckExecutionPermission(record.Model().Methods().MustGet("Load")) {
				continue
			}
		}
		notifications := make(map[int64]mailtypes.PartnerEmailStatus)
		for _, notif := range message.Notifications().Sudo().Records() {
			notifications[notif.Partner().ID()] = mailtypes.PartnerEmailStatus{
				EmailStatus: notif.EmailStatus(),
				Name:        notif.Partner().Name(),
			}
		}
		info := mailtypes.MailFailureInfo{
			MessageID:       message.ID(),
			RecordName:      message.RecordName(),
			ModelName:       models.Registry.MustGet(message.ResModel()).Description(),
			UUID:            message.ID(),
			ResID:           message.ResID(),
			Model:           message.ResModel(),
			LastMessageDate: message.Date(),
			ModuleIcon:      "/static/mail/src/img/smiley/mailfailure.jpg",
			Notifications:   notifications,
		}
		failureInfos = append(failureInfos, &info)
	}
	return failureInfos
}

// NotifyFailureUpdate notifies mail failures on the event bus
func mailMessage_NotifyFailureUpdate(rs m.MailMessageSet) {
	msgsByAuthors := make(map[int64]m.MailMessageSet)
	for _, rec := range rs.Records() {
		if _, ok := msgsByAuthors[rec.Author().ID()]; !ok {
			msgsByAuthors[rec.Author().ID()] = rec
			continue
		}
		msgsByAuthors[rec.Author().ID()] = msgsByAuthors[rec.Author().ID()].Union(rec)
	}
	for authorID, recs := range msgsByAuthors {
		h.BusBus().NewSet(rs.Env()).Sendone(
			fmt.Sprintf("res.partner.%s", authorID),
			map[string]interface{}{
				"type":     "mail_failure",
				"elements": recs.FormatMailFailures(),
			})
	}
}

// ------------------------------------------------------
//  MailMessage internals
// ------------------------------------------------------

// Init creates the index on ResModel and ResID
func mailMessage_Init(rs m.MailMessageSet) {
	var indexNames []string
	rs.Env().Cr().Select(&indexNames, `SELECT indexname FROM pg_indexes WHERE indexname = 'mail_message_model_res_id_idx'`)
	if len(indexNames) == 0 {
		rs.Env().Cr().Execute(`CREATE INDEX mail_message_model_res_id_idx ON mail_message (model, res_id)`)
	}
}

// FindAllowedModelWise
func mailMessage_FindAllowedModelWise(rs m.MailMessageSet, doc_model interface{}, doc_dict interface{}) {
	//        doc_ids = list(doc_dict)
	//        allowed_doc_ids = self.env[doc_model].with_context(
	//            active_test=False).search([('id', 'in', doc_ids)]).ids
	//        return set([message_id for allowed_doc_id in allowed_doc_ids for message_id in doc_dict[allowed_doc_id]])
}

// FindAllowedDocIds
func mailMessage_FindAllowedDocIds(rs m.MailMessageSet, model_ids interface{}) {
	//        IrModelAccess = self.env['ir.model.access']
	//        allowed_ids = set()
	//        for doc_model, doc_dict in model_ids.items():
	//            if not IrModelAccess.check(doc_model, 'read', False):
	//                continue
	//            allowed_ids |= self._find_allowed_model_wise(doc_model, doc_dict)
	//        return allowed_ids
}

//  Override that adds specific access rights of mail.message, to remove
//         ids uid could not see according to our custom rules.
// Please refer to
//         check_access_rule for more details about those rules.
//
//         Non employees users see only message with subtype
// (aka do not see
//         internal logs).
//
//         After having received ids of a classic search, keep only:
//         - if author_id == pid, uid is the author, OR
//         - uid belongs to a notified channel, OR
//         - uid is in the specified recipients, OR
//         - uid has a notification on the message
//         - otherwise: remove the id
//
func mailMessage_Search(rs m.MailMessageSet, args q.MailMessageCondition, offset int, limit int, order interface{}, count interface{}, access_rights_uid interface{}) {
	//        if self._uid == SUPERUSER_ID:
	//            return super(Message, self)._search(
	//                args, offset=offset, limit=limit, order=order,
	//                count=count, access_rights_uid=access_rights_uid)
	//        if not self.env['res.users'].has_group('base.group_user'):
	//            args = ['&', '&', ('subtype_id', '!=', False),
	//                    ('subtype_id.internal', '=', False)] + list(args)
	//        ids = super(Message, self)._search(
	//            args, offset=offset, limit=limit, order=order,
	//            count=False, access_rights_uid=access_rights_uid)
	//        if not ids and count:
	//            return 0
	//        elif not ids:
	//            return ids
	//        pid = self.env.user.partner_id.id
	//        author_ids, partner_ids, channel_ids, allowed_ids = set(
	//            []), set([]), set([]), set([])
	//        model_ids = {}
	//        super(Message, self.sudo(access_rights_uid or self._uid)
	//              ).check_access_rights('read')
	//        for sub_ids in self._cr.split_for_in_conditions(ids):
	//            self._cr.execute("""
	//                SELECT DISTINCT m.id, m.model, m.res_id, m.author_id,
	//                                COALESCE(partner_rel.res_partner_id, needaction_rel.res_partner_id),
	//                                channel_partner.channel_id as channel_id
	//                FROM "%s" m
	//                LEFT JOIN "mail_message_res_partner_rel" partner_rel
	//                ON partner_rel.mail_message_id = m.id AND partner_rel.res_partner_id = %%(pid)s
	//                LEFT JOIN "mail_message_res_partner_needaction_rel" needaction_rel
	//                ON needaction_rel.mail_message_id = m.id AND needaction_rel.res_partner_id = %%(pid)s
	//                LEFT JOIN "mail_message_mail_channel_rel" channel_rel
	//                ON channel_rel.mail_message_id = m.id
	//                LEFT JOIN "mail_channel" channel
	//                ON channel.id = channel_rel.mail_channel_id
	//                LEFT JOIN "mail_channel_partner" channel_partner
	//                ON channel_partner.channel_id = channel.id AND channel_partner.partner_id = %%(pid)s
	//                WHERE m.id = ANY (%%(ids)s)""" % self._table, dict(pid=pid, ids=list(sub_ids)))
	//            for id, rmod, rid, author_id, partner_id, channel_id in self._cr.fetchall():
	//                if author_id == pid:
	//                    author_ids.add(id)
	//                elif partner_id == pid:
	//                    partner_ids.add(id)
	//                elif channel_id:
	//                    channel_ids.add(id)
	//                elif rmod and rid:
	//                    model_ids.setdefault(rmod, {}).setdefault(
	//                        rid, set()).add(id)
	//        allowed_ids = self._find_allowed_doc_ids(model_ids)
	//        final_ids = author_ids | partner_ids | channel_ids | allowed_ids
	//        if count:
	//            return len(final_ids)
	//        else:
	//            # re-construct a list based on ids, because set did not keep the original order
	//            id_list = [id for id in ids if id in final_ids]
	//            return id_list
}

//  Access rules of mail.message:
//             - read: if
//                 - author_id == pid, uid is the author OR
//                 - uid is in the recipients (partner_ids) OR
//                 - uid has been notified (needaction) OR
//                 - uid is member of a listern channel (channel_ids.partner_ids)
// OR
//                 - uid have read access to the related document
// if model, res_id
//                 - otherwise: raise
//             - create: if
//                 - no model, no res_id (private message) OR
//                 - pid in message_follower_ids if model, res_id OR
//                 - uid can read the parent OR
//                 - uid have write or create access on the
// related document if model, res_id, OR
//                 - otherwise: raise
//             - write: if
//                 - author_id == pid, uid is the author, OR
//                 - uid is in the recipients (partner_ids) OR
//                 - uid is moderator of the channel and moderation_status
// is pending_moderation OR
//                 - uid has write or create access on the
// related document if model, res_id and moderation_status
// is not pending_moderation
//                 - otherwise: raise
//             - unlink: if
//                 - uid is moderator of the channel and moderation_status
// is pending_moderation OR
//                 - uid has write or create access on the
// related document if model, res_id and moderation_status
// is not pending_moderation
//                 - otherwise: raise
//
//         Specific case: non employee users see only messages
// with subtype (aka do
//         not see internal logs).
//
func mailMessage_CheckAccessRule(rs m.MailMessageSet, operation string) {
	//        def _generate_model_record_ids(msg_val, msg_ids):
	//            """ :param model_record_ids: {'model': {'res_id': (msg_id, msg_id)}, ... }
	//                :param message_values: {'msg_id': {'model': .., 'res_id': .., 'author_id': ..}}
	//            """
	//            model_record_ids = {}
	//            for id in msg_ids:
	//                vals = msg_val.get(id, {})
	//                if vals.get('model') and vals.get('res_id'):
	//                    model_record_ids.setdefault(
	//                        vals['model'], set()).add(vals['res_id'])
	//            return model_record_ids
	//        if self._uid == SUPERUSER_ID:
	//            return
	//        if not self.env['res.users'].has_group('base.group_user'):
	//            self._cr.execute('''SELECT DISTINCT message.id, message.subtype_id, subtype.internal
	//                                FROM "%s" AS message
	//                                LEFT JOIN "mail_message_subtype" as subtype
	//                                ON message.subtype_id = subtype.id
	//                                WHERE message.message_type = %%s AND (message.subtype_id IS NULL OR subtype.internal IS TRUE) AND message.id = ANY (%%s)''' % (self._table), ('comment', self.ids))
	//            if self._cr.fetchall():
	//                raise AccessError(
	//                    _('The requested operation cannot be completed due to security restrictions. Please contact your system administrator.\n\n(Document type: %s, Operation: %s)') % (
	//                        self._description, operation)
	//                    + ' - ({} {}, {} {})'.format(_('Records:'),
	//                                                 self.ids[:6], _('User:'), self._uid)
	//                )
	//        message_values = dict((res_id, {}) for res_id in self.ids)
	//        if operation == 'read':
	//            self._cr.execute("""
	//                SELECT DISTINCT m.id, m.model, m.res_id, m.author_id, m.parent_id,
	//                                COALESCE(partner_rel.res_partner_id, needaction_rel.res_partner_id),
	//                                channel_partner.channel_id as channel_id, m.moderation_status
	//                FROM "%s" m
	//                LEFT JOIN "mail_message_res_partner_rel" partner_rel
	//                ON partner_rel.mail_message_id = m.id AND partner_rel.res_partner_id = %%(pid)s
	//                LEFT JOIN "mail_message_res_partner_needaction_rel" needaction_rel
	//                ON needaction_rel.mail_message_id = m.id AND needaction_rel.res_partner_id = %%(pid)s
	//                LEFT JOIN "mail_message_mail_channel_rel" channel_rel
	//                ON channel_rel.mail_message_id = m.id
	//                LEFT JOIN "mail_channel" channel
	//                ON channel.id = channel_rel.mail_channel_id
	//                LEFT JOIN "mail_channel_partner" channel_partner
	//                ON channel_partner.channel_id = channel.id AND channel_partner.partner_id = %%(pid)s
	//                WHERE m.id = ANY (%%(ids)s)""" % self._table, dict(pid=self.env.user.partner_id.id, ids=self.ids))
	//            for mid, rmod, rid, author_id, parent_id, partner_id, channel_id, moderation_status in self._cr.fetchall():
	//                message_values[mid] = {
	//                    'model': rmod,
	//                    'res_id': rid,
	//                    'author_id': author_id,
	//                    'parent_id': parent_id,
	//                    'moderation_status': moderation_status,
	//                    'moderator_id': False,
	//                    'notified': any((message_values[mid].get('notified'), partner_id, channel_id))
	//                }
	//        elif operation == 'write':
	//            self._cr.execute("""
	//                SELECT DISTINCT m.id, m.model, m.res_id, m.author_id, m.parent_id, m.moderation_status,
	//                                COALESCE(partner_rel.res_partner_id, needaction_rel.res_partner_id),
	//                                channel_partner.channel_id as channel_id, channel_moderator_rel.res_users_id as moderator_id
	//                FROM "%s" m
	//                LEFT JOIN "mail_message_res_partner_rel" partner_rel
	//                ON partner_rel.mail_message_id = m.id AND partner_rel.res_partner_id = %%(pid)s
	//                LEFT JOIN "mail_message_res_partner_needaction_rel" needaction_rel
	//                ON needaction_rel.mail_message_id = m.id AND needaction_rel.res_partner_id = %%(pid)s
	//                LEFT JOIN "mail_message_mail_channel_rel" channel_rel
	//                ON channel_rel.mail_message_id = m.id
	//                LEFT JOIN "mail_channel" channel
	//                ON channel.id = channel_rel.mail_channel_id
	//                LEFT JOIN "mail_channel_partner" channel_partner
	//                ON channel_partner.channel_id = channel.id AND channel_partner.partner_id = %%(pid)s
	//                LEFT JOIN "mail_channel" moderated_channel
	//                ON m.moderation_status = 'pending_moderation' AND m.res_id = moderated_channel.id
	//                LEFT JOIN "mail_channel_moderator_rel" channel_moderator_rel
	//                ON channel_moderator_rel.mail_channel_id = moderated_channel.id AND channel_moderator_rel.res_users_id = %%(uid)s
	//                WHERE m.id = ANY (%%(ids)s)""" % self._table, dict(pid=self.env.user.partner_id.id, uid=self.env.user.id, ids=self.ids))
	//            for mid, rmod, rid, author_id, parent_id, moderation_status, partner_id, channel_id, moderator_id in self._cr.fetchall():
	//                message_values[mid] = {
	//                    'model': rmod,
	//                    'res_id': rid,
	//                    'author_id': author_id,
	//                    'parent_id': parent_id,
	//                    'moderation_status': moderation_status,
	//                    'moderator_id': moderator_id,
	//                    'notified': any((message_values[mid].get('notified'), partner_id, channel_id))
	//                }
	//        elif operation == 'create':
	//            self._cr.execute(
	//                """SELECT DISTINCT id, model, res_id, author_id, parent_id, moderation_status FROM "%s" WHERE id = ANY (%%s)""" % self._table, (self.ids))
	//            for mid, rmod, rid, author_id, parent_id, moderation_status in self._cr.fetchall():
	//                message_values[mid] = {
	//                    'model': rmod,
	//                    'res_id': rid,
	//                    'author_id': author_id,
	//                    'parent_id': parent_id,
	//                    'moderation_status': moderation_status,
	//                    'moderator_id': False
	//                }
	//        else:  # unlink
	//            self._cr.execute("""SELECT DISTINCT m.id, m.model, m.res_id, m.author_id, m.parent_id, m.moderation_status, channel_moderator_rel.res_users_id as moderator_id
	//                FROM "%s" m
	//                LEFT JOIN "mail_channel" moderated_channel
	//                ON m.moderation_status = 'pending_moderation' AND m.res_id = moderated_channel.id
	//                LEFT JOIN "mail_channel_moderator_rel" channel_moderator_rel
	//                ON channel_moderator_rel.mail_channel_id = moderated_channel.id AND channel_moderator_rel.res_users_id = (%%s)
	//                WHERE m.id = ANY (%%s)""" % self._table, (self.env.user.id, self.ids))
	//            for mid, rmod, rid, author_id, parent_id, moderation_status, moderator_id in self._cr.fetchall():
	//                message_values[mid] = {
	//                    'model': rmod,
	//                    'res_id': rid,
	//                    'author_id': author_id,
	//                    'parent_id': parent_id,
	//                    'moderation_status': moderation_status,
	//                    'moderator_id': moderator_id
	//                }
	//        author_ids = []
	//        if operation == 'read':
	//            author_ids = [mid for mid, message in message_values.items()
	//                          if message.get('author_id') and message.get('author_id') == self.env.user.partner_id.id]
	//        elif operation == 'write':
	//            author_ids = [mid for mid, message in message_values.items()
	//                          if message.get('moderation_status') != 'pending_moderation' and message.get('author_id') == self.env.user.partner_id.id]
	//        elif operation == 'create':
	//            author_ids = [mid for mid, message in message_values.items()
	//                          if not message.get('model') and not message.get('res_id')]
	//        notified_ids = []
	//        if operation == 'create':
	//            # TDE: probably clean me
	//            parent_ids = [message.get('parent_id') for message in message_values.values()
	//                          if message.get('parent_id')]
	//            self._cr.execute("""SELECT DISTINCT m.id, partner_rel.res_partner_id, channel_partner.partner_id FROM "%s" m
	//                LEFT JOIN "mail_message_res_partner_rel" partner_rel
	//                ON partner_rel.mail_message_id = m.id AND partner_rel.res_partner_id = (%%s)
	//                LEFT JOIN "mail_message_mail_channel_rel" channel_rel
	//                ON channel_rel.mail_message_id = m.id
	//                LEFT JOIN "mail_channel" channel
	//                ON channel.id = channel_rel.mail_channel_id
	//                LEFT JOIN "mail_channel_partner" channel_partner
	//                ON channel_partner.channel_id = channel.id AND channel_partner.partner_id = (%%s)
	//                WHERE m.id = ANY (%%s)""" % self._table, (self.env.user.partner_id.id, self.env.user.partner_id.id, parent_ids))
	//            not_parent_ids = [mid[0]
	//                              for mid in self._cr.fetchall() if any([mid[1], mid[2]])]
	//            notified_ids += [mid for mid, message in message_values.items()
	//                             if message.get('parent_id') in not_parent_ids]
	//        moderator_ids = []
	//        if operation in ['write', 'unlink']:
	//            moderator_ids = [
	//                mid for mid, message in message_values.items() if message.get('moderator_id')]
	//        other_ids = set(self.ids).difference(set(author_ids),
	//                                             set(notified_ids), set(moderator_ids))
	//        model_record_ids = _generate_model_record_ids(
	//            message_values, other_ids)
	//        if operation in ['read', 'write']:
	//            notified_ids = [
	//                mid for mid, message in message_values.items() if message.get('notified')]
	//        elif operation == 'create':
	//            for doc_model, doc_ids in model_record_ids.items():
	//                followers = self.env['mail.followers'].sudo().search([
	//                    ('res_model', '=', doc_model),
	//                    ('res_id', 'in', list(doc_ids)),
	//                    ('partner_id', '=', self.env.user.partner_id.id),
	//                ])
	//                fol_mids = [follower.res_id for follower in followers]
	//                notified_ids += [mid for mid, message in message_values.items()
	//                                 if message.get('model') == doc_model and message.get('res_id') in fol_mids]
	//        other_ids = other_ids.difference(set(notified_ids))
	//        model_record_ids = _generate_model_record_ids(
	//            message_values, other_ids)
	//        document_related_ids = []
	//        for model, doc_ids in model_record_ids.items():
	//            DocumentModel = self.env[model]
	//            mids = DocumentModel.browse(doc_ids).exists()
	//            if hasattr(DocumentModel, 'check_mail_message_access'):
	//                DocumentModel.check_mail_message_access(
	//                    mids.ids, operation)  # ?? mids ?
	//            else:
	//                self.env['mail.thread'].check_mail_message_access(
	//                    mids.ids, operation, model_name=model)
	//            if operation in ['write', 'unlink']:
	//                document_related_ids += [mid for mid, message in message_values.items()
	//                                         if message.get('model') == model and message.get('res_id') in mids.ids and
	//                                         message.get('moderation_status') != 'pending_moderation']
	//            else:
	//                document_related_ids += [mid for mid, message in message_values.items()
	//                                         if message.get('model') == model and message.get('res_id') in mids.ids]
	//        other_ids = other_ids.difference(set(document_related_ids))
	//        if not (other_ids and self.browse(other_ids).exists()):
	//            return
	//        raise AccessError(
	//            _('The requested operation cannot be completed due to security restrictions. Please contact your system administrator.\n\n(Document type: %s, Operation: %s)') % (
	//                self._description, operation)
	//            + ' - ({} {}, {} {})'.format(_('Records:'),
	//                                         list(other_ids)[:6], _('User:'), self._uid)
	//        )
}

// GetRecordName returns the related document name, using NameGet.
// It is done using SUPERUSER_ID, to be sure to have the record
// name correctly stored.
func mailMessage_GetRecordName(rs m.MailMessageSet, values m.MailMessageData) string {
	model := rs.Env().Context().GetString("default_res_model")
	if values.ResModel() != "" {
		model = values.ResModel()
	}
	resID := rs.Env().Context().GetInteger("default_res_id")
	if values.ResID() != 0 {
		resID = values.ResID()
	}
	_, modelExists := models.Registry.Get(model)
	if model == "" || resID == 0 || !modelExists {
		return ""
	}
	return models.Registry.MustGet(model).BrowseOne(rs.Sudo().Env(), resID).Wrap().(m.CommonMixinSet).NameGet()
}

// GetReplyTo returns a specific reply_to for the document
func mailMessage_GetReplyTo(rs m.MailMessageSet, values m.MailMessageData) string {
	model := rs.Env().Context().GetString("default_res_model")
	if values.ResModel() != "" {
		model = values.ResModel()
	}
	resID := rs.Env().Context().GetInteger("default_res_id")
	if values.ResID() != 0 {
		resID = values.ResID()
	}
	records := models.Registry.MustGet(model).BrowseOne(rs.Env(), resID)
	return h.MailThread().NewSet(rs.Env()).NotifyGetReplyToOnRecords(values.EmailFrom(), records, h.Company().NewSet(rs.Env()), map[int64]string{})[resID]
}

// GetMessageID returns a Message-ID RFC822 header field string
func mailMessage_GetMessageID(rs m.MailMessageSet, values m.MailMessageData) string {
	switch {
	case values.NoAutoThread():
		return emailutils.GenerateTrackingMessageID("reply-to")
	case values.ResModel() != "" && values.ResID() != 0:
		return emailutils.GenerateTrackingMessageID(fmt.Sprintf("%d-%s", values.ResID(), values.ResModel()))
	default:
		return emailutils.GenerateTrackingMessageID("private")
	}
}

// InvalidateDocuments invalidate the cache of the documents followed by ``self``.
func mailMessage_InvalidateDocuments(rs m.MailMessageSet) {
	for _, record := range rs.Records() {
		if record.ResModel() != "" && record.ResID() != 0 {
			relRS := models.Registry.MustGet(record.ResModel()).BrowseOne(rs.Env(), record.ResID())
			if _, ok := relRS.Wrap().(m.MailThreadSet); !ok {
				continue
			}
			relRS.InvalidateCache()
		}
	}
}

func mailMessage_Create(rs m.MailMessageSet, values m.MailMessageData) m.MailMessageSet {
	// coming from mail.js that does not have pid in its values
	if rs.Env().Context().HasKey("default_starred") {
		values.SetStarredPartners(values.StarredPartners().Union(h.User().NewSet(rs.Env()).CurrentUser().Partner()))
	}
	if !values.HasEmailFrom() {
		values.SetEmailFrom(getDefaultFrom(rs.Env()).(string))
	}
	if values.MessageID() == "" {
		values.SetMessageID(rs.GetMessageID(values))
	}
	if !values.HasReplyTo() {
		values.SetReplyTo(rs.GetReplyTo(values))
	}
	if !values.HasRecordName() && !rs.Env().Context().HasKey("default_record_name") {
		values.SetRecordName(rs.GetRecordName(values))
	}
	if !values.HasAttachments() {
		values.SetAttachments(h.Attachment().NewSet(rs.Env()))
	}
	if values.HasBody() {
		dataToURL := make(map[string][2]string)
		base64ToBoundary := func(match []string) string {
			if len(match) < 5 {
				log.Warn("Badly formatted base64 embedded image")
				return ""
			}
			if match[2] == "" {
				log.Warn("Badly formatted base64 embedded image")
				return ""
			}
			if _, ok := dataToURL[match[2]]; !ok {
				name := fmt.Sprintf("image%d", len(dataToURL))
				if match[4] != "" {
					name = match[4]
				}
				attachment := h.Attachment().Create(rs.Env(), h.Attachment().NewData().
					SetName(name).
					SetDatas(match[2]).
					SetDatasFname(name).
					SetResModel(values.ResModel()).
					SetResID(values.ResID()))
				attachment.GenerateAccessToken()
				values.SetAttachments(values.Attachments().Union(attachment))
				dataToURL[match[2]] = [2]string{
					fmt.Sprintf("/web/image/%d?access_token=%s", attachment.ID(), attachment.AccessToken()),
					name,
				}
			}
			return fmt.Sprintf("%s%s alt=\"%s\"", dataToURL[match[2]][0], match[3], dataToURL[match[2]][1])
		}
		// TODO finish implementation
		matches := imageDataURL.FindAllStringSubmatch(values.Body(), -1)
		var body string
		for _, match := range matches {
			base64ToBoundary(match)
		}
		values.SetBody(body)
	}
	// delegate creation of tracking after the create as sudo to avoid access rights issues
	trackingValues := values.TrackingValues()
	values.UnsetTrackingValues()

	message := rs.Super().Create(values)
	if values.HasAttachments() {
		message.Attachments().Check("read", nil)
	}
	if trackingValues.IsNotEmpty() {
		message.Sudo().Write(h.MailMessage().NewData().SetTrackingValues(trackingValues))
	}
	if values.HasResModel() && values.HasResID() {
		message.InvalidateDocuments()
	}
	return message
}

func mailMessage_Write(rs m.MailMessageSet, vals m.MailMessageData) bool {
	if vals.HasResModel() || vals.HasResID() {
		rs.InvalidateDocuments()
	}
	res := rs.Super().Write(vals)
	if vals.Attachments().IsNotEmpty() {
		for _, msg := range rs.Records() {
			msg.Attachments().Check("read", nil)
		}
	}
	if vals.HasNotifications() || vals.HasResModel() || vals.HasResID() {
		rs.InvalidateDocuments()
	}
	return res
}

func mailMessage_Unlink(rs m.MailMessageSet) int64 {
	if rs.IsEmpty() {
		return 0
	}
	rs.CheckAccessRule("unlink")
	attachments := h.Attachment().NewSet(rs.Env())
	for _, rec := range rs.Records() {
		for _, attach := range rec.Attachments().Records() {
			if attach.ResModel() == rs.ModelName() && (attach.ResID() == 0 || attach.ResID() == rec.ID()) {
				attachments = attachments.Union(attach)
			}
		}
	}
	attachments.Unlink()
	rs.InvalidateDocuments()
	return rs.Super().Unlink()
}

// ------------------------------------------------------
// Messaging API
// ------------------------------------------------------

// Notify is the main notification method. This method basically does two things
//
// * call NotifyComputeRecipients that computes recipients to
//   notify based on message record or message creation values if given
//   (to optimize performance if we already have data computed);
// * call NotifyRecipients that performs the notification process;
//
// Parameters:
//
// record: record on which the message is posted, if any;
// msgVals: dictionary of values used to create the message.
//           If given it is used instead of accessing rs to lesen query count in some
//           simple cases where no notification is actually required;
// forceSend: tells whether to send notification emails within the
//             current transaction or to use the email queue;
// sendAfterCommit: if forceSend, tells whether to send emails after
//                  the transaction has been committed using a post-commit hook;
// modelDescription: optional data used in notification process (see notification templates);
// mailAutoDelete: delete notification emails once sent;
func mailMessage_Notify(rs m.MailMessageSet, record m.MailThreadSet, msgVals m.MailMessageData, forceSend bool,
	sendAfterCommit bool, modelDescription string, mailAutoDelete bool) bool {
	if msgVals == nil {
		msgVals = h.MailMessage().NewData(rs.Env())
	}
	rData := rs.NotifyComputeRecipients(record, msgVals)
	return rs.NotifyRecipients(rData, record, msgVals, forceSend, sendAfterCommit, modelDescription, mailAutoDelete)
}

// NotifyComputeRecipients compute recipients to notify based on subtype and followers.
// This method returns data structured as expected for NotifyRecipients.
func mailMessage_NotifyComputeRecipients(rs m.MailMessageSet, record m.MailThreadSet, msgVals m.MailMessageData) mailtypes.RecipientsData {
	msgSudo := rs.Sudo()
	partners := msgSudo.Partners()
	if msgVals.HasPartners() {
		partners = msgVals.Partners()
	}
	channels := msgSudo.Channels()
	if msgVals.HasChannels() {
		channels = msgVals.Channels()
	}
	subtype := msgSudo.Subtype()
	if msgVals.HasSubtype() {
		subtype = msgVals.Subtype()
	}
	var recipientData mailtypes.RecipientsData
	res := h.MailFollowers().NewSet(rs.Env()).GetRecipientData(record, subtype, partners, channels)

	author := h.Partner().NewSet(rs.Env())
	if len(res) != 0 {
		author = rs.Author()
		if msgVals.Author().IsNotEmpty() {
			author = msgVals.Author()
		}
	}

	for _, data := range res {
		if data.PartnerID != 0 && data.PartnerID == author.ID() && !rs.Env().Context().GetBool("mail_notify_author") {
			// do not notify the author of its own messages
			continue
		}
		switch {
		case data.PartnerID != 0:
			if !data.Active {
				// avoid to notify inactive partner by email
				continue
			}
			pData := mailtypes.PartnerData{
				ID:               data.PartnerID,
				Active:           data.Active,
				Share:            data.Share,
				Groups:           data.Groups,
				NotificationType: "email",
			}
			switch data.NotificationType {
			case "inbox":
				pData.NotificationType = data.NotificationType
				pData.Type = "user"
				recipientData.Partners = append(recipientData.Partners, pData)
			default:
				switch {
				case !data.Share && data.NotificationType != "":
					// has a user and is not shared, is therefore user
					pData.Type = "user"
					recipientData.Partners = append(recipientData.Partners, pData)
				case data.Share && data.NotificationType != "":
					// has a user but is shared, is therefore portal
					pData.Type = "portal"
					recipientData.Partners = append(recipientData.Partners, pData)
				default:
					// has no user, is therefore customer
					pData.Type = "customer"
					recipientData.Partners = append(recipientData.Partners, pData)
				}
			}
		case data.ChannelID != 0:
			recipientData.Channels = append(recipientData.Channels, mailtypes.ChannelData{
				ID:               data.ChannelID,
				NotificationType: data.NotificationType,
				Type:             data.NotificationType,
			})
		}
	}
	return recipientData
}

// NotifyRecipients is the main method implementing the notification process.
func mailMessage_NotifyRecipients(rs m.MailMessageSet, rdata mailtypes.RecipientsData, record m.MailThreadSet, msgVals m.MailMessageData,
	forceSend bool, SendAfterCommit bool, modelDescription string, mailAutoDelete bool) bool {

	rs.EnsureOne()
	var channelIDs, partnerIDs, emailCIDs, inboxPIDs []int64
	var partnerEmailRData []mailtypes.PartnerData
	for _, r := range rdata.Channels {
		channelIDs = append(channelIDs, r.ID)
		if r.NotificationType == "email" {
			emailCIDs = append(emailCIDs, r.ID)
		}
	}
	for _, r := range rdata.Partners {
		partnerIDs = append(partnerIDs, r.ID)
		if r.NotificationType == "inbox" {
			inboxPIDs = append(inboxPIDs, r.ID)
		}
		if r.NotificationType == "email" {
			partnerEmailRData = append(partnerEmailRData, r)
		}
	}
	messageValues := h.MailMessage().NewData()
	if len(rdata.Channels) > 0 {
		messageValues.SetChannels(h.MailChannel().Browse(rs.Env(), channelIDs))
	}
	if len(rdata.Partners) > 0 {
		messageValues.SetNeedactionPartners(h.Partner().Browse(rs.Env(), partnerIDs))
	}
	if record.IsNotEmpty() {
		messageValues = record.NotifyCustomizeRecipients(rs, messageValues, rdata)
	}
	rs.Write(messageValues)

	// notify partners and channels
	if len(emailCIDs) > 0 {
		newPIDs := h.Partner().NewSet(rs.Env()).Sudo().Search(q.Partner().
			ID().NotIn(partnerIDs).And().
			ChannelsFilteredOn(q.MailChannel().ID().NotIn(emailCIDs)).And().
			Email().NotIn([]string{rs.Author().Email(), rs.EmailFrom()}))
		for _, partner := range newPIDs.Records() {
			rdata.Partners = append(rdata.Partners, mailtypes.PartnerData{
				ID:               partner.ID(),
				Share:            true,
				NotificationType: "email",
				Type:             "customer",
				Groups:           []string{},
			})
		}
	}

	if len(partnerEmailRData) > 0 {
		h.Partner().NewSet(rs.Env()).Notify(rs, partnerEmailRData, record, forceSend, SendAfterCommit, modelDescription, mailAutoDelete)
	}
	if len(inboxPIDs) > 0 {
		h.Partner().Browse(rs.Env(), inboxPIDs).NotifyByChat(rs)
	}
	if len(rdata.Channels) > 0 {
		h.MailChannel().NewSet(rs.Env()).Sudo().Browse(channelIDs).Notify(rs)
	}
	return true
}

// --------------------------------------------------
// Moderation
// --------------------------------------------------

// Moderate messages. A check is done on moderation status of the
// current user to ensure we only moderate valid messages.
// Title and comment are used for the reject email if any.
func mailMessage_Moderate(rs m.MailMessageSet, decision string, title string, comment string) {
	moderatedChannels := h.User().NewSet(rs.Env()).CurrentUser().ModerationChannels()
	toModerate := h.MailMessage().NewSet(rs.Env())
	for _, message := range rs.Records() {
		if message.ResModel() == "MailChannel" &&
			moderatedChannels.Intersect(h.MailChannel().BrowseOne(rs.Env(), message.ResID())).IsNotEmpty() &&
			message.ModerationStatus() == "pending_moderation" {
			toModerate = toModerate.Union(message)
		}
	}
	if toModerate.IsNotEmpty() {
		toModerate.DoModerate(decision)
	}
}

// DoModerate moderates these messages with the given decision, which can be :
//
// * accept       - moderate message and broadcast that message to followers of relevant channels.
// * reject       - message will be deleted from the database without broadcast an email sent to the
//                  author with an explanation that the moderators can edit.
// * discard      - message will be deleted from the database without broadcast.
// * allow        - add email address to white list people of specific channel,
//                  so that next time if a message come from same email address on same channel,
//                  it will be automatically broadcasted to relevant channels without any approval from moderator.
// * ban          - add email address to black list of emails for the specific channel.
//                  From next time, a person sending a message using that email address will not need moderation.
//                  MessagePost will not create messages with the corresponding expeditor.
//
// Title and comment are used for the reject email if any.
func mailMessage_DoModerate(rs m.MailMessageSet, decision string, title string, comment string) {
	updateEmails := func(status string) {
		channels := h.MailChannel().NewSet(rs.Env())
		for _, rec := range rs.Records() {
			channels = channels.Union(h.MailChannel().BrowseOne(rs.Env(), rec.ID()))
		}
		for _, channel := range channels.Records() {
			var emails []string
			for _, msg := range rs.Records() {
				if msg.ResID() == channel.ID() {
					emails = append(emails, msg.EmailFrom())
				}
			}
			channel.UpdateModerationEmail(emails, status)
		}
	}

	switch decision {
	case "accept":
		rs.ModerateAccept()
	case "reject":
		rs.ModerateSendRejectEmail(title, comment)
		rs.ModerateDiscard()
	case "discard":
		rs.ModerateDiscard()
	case "allow":
		updateEmails("allow")
		rs.SearchFromSameAuthors().ModerateAccept()
	case "ban":
		updateEmails("ban")
		rs.SearchFromSameAuthors().ModerateDiscard()
	}
}

// ModerateAccept accepts the given messages
func mailMessage_ModerateAccept(rs m.MailMessageSet) {
	rs.Write(h.MailMessage().NewData().
		SetModerationStatus("accepted").
		SetModerator(h.User().NewSet(rs.Env()).CurrentUser()))
	for _, message := range rs.Records() {
		record := h.MailThread().NewSet(rs.Env())
		if message.ResModel() != "" && message.ResID() != 0 {
			record = models.Registry.MustGet(message.ResModel()).BrowseOne(rs.Env(), message.ResID()).Wrap().(m.MailThreadSet)
		}
		message.Notify(record, h.MailMessage().NewData(), false, true, "", true)
	}
}

// ModerateSendRejectEmail sends an email to the sender of this message explaining the reason to reject it.
func mailMessage_ModerateSendRejectEmail(rs m.MailMessageSet, subject string, comment string) {
	for _, msg := range rs.Records() {
		if msg.EmailFrom() == "" {
			continue
		}
		currentUser := h.User().NewSet(rs.Env()).CurrentUser()
		emailFrom := currentUser.Company().Catchall()
		if h.User().NewSet(rs.Env()).CurrentUser().Partner().Email() != "" {
			emailAddr := mail.Address{
				Name:    currentUser.Partner().Name(),
				Address: currentUser.Partner().Email(),
			}
			emailFrom = emailAddr.String()
		}
		//            body_html = tools.append_content_to_html(
		//                '<div>%s</div>' % tools.ustr(comment), msg.body)
		// FIXME
		bodyHTML := fmt.Sprintf("<div>%s</div>", comment)
		vals := h.MailMail().NewData().
			SetSubject(subject).
			SetBodyHtml(bodyHTML).
			SetEmailFrom(emailFrom).
			SetEmailTo(msg.EmailFrom()).
			SetAutoDelete(true).
			SetState("outgoing")
		h.MailMail().NewSet(rs.Env()).Sudo().Create(vals)
	}
}

// SearchFromSameAuthors returns all pending moderation messages that have same EmailFrom and
// same ResID as given recordset.
func mailMessage_SearchFromSameAuthors(rs m.MailMessageSet) m.MailMessageSet {
	messages := h.MailMessage().NewSet(rs.Env()).Sudo()
	for _, message := range rs.Records() {
		messages = messages.Union(messages.Search(q.MailMessage().
			ModerationStatus().Equals("pending_moderation").And().
			EmailFrom().Equals(message.EmailFrom()).And().
			ResModel().Equals("MailChannel").And().
			ResID().Equals(message.ResID()),
		))
	}
	return messages
}

// ModerateDiscard notify deletion of messages to their moderators and authors
// and then delete them.
func mailMessage_ModerateDiscard(rs m.MailMessageSet) {
	var (
		channels   = h.MailChannel().NewSet(rs.Env())
		moderators = h.User().NewSet(rs.Env())
		authors    = h.Partner().NewSet(rs.Env())
	)
	for _, message := range rs.Records() {
		channel := h.MailChannel().BrowseOne(rs.Env(), message.ResID())
		channels = channels.Union(channel)
		moderators = moderators.Union(channel.Moderators())
		authors = authors.Union(message.Author())
	}
	partnerToPID := make(map[int64]map[int64]bool)
	for _, moderator := range moderators.Records() {
		if _, ok := partnerToPID[moderator.Partner().ID()]; !ok {
			partnerToPID[moderator.Partner().ID()] = make(map[int64]bool)
		}
		modMessages := rs.Filtered(func(r m.MailMessageSet) bool {
			for _, ch := range moderator.ModerationChannels().Records() {
				if r.ResID() == ch.ID() {
					return true
				}
			}
			return false
		})
		for _, msg := range modMessages.Records() {
			partnerToPID[moderator.Partner().ID()][msg.ID()] = true
		}
	}
	for _, author := range authors.Records() {
		if _, ok := partnerToPID[author.ID()]; !ok {
			partnerToPID[author.ID()] = make(map[int64]bool)
		}
		authorMessages := rs.Filtered(func(r m.MailMessageSet) bool {
			if r.Author().Equals(author) {
				return true
			}
			return false
		})
		for _, msg := range authorMessages.Records() {
			partnerToPID[author.ID()][msg.ID()] = true
		}
	}
	var notifications []*bustypes.Notification
	for partnerID, messages := range partnerToPID {
		var msgList []int64
		for msg := range messages {
			msgList = append(msgList, msg)
		}
		notifications = append(notifications, &bustypes.Notification{
			Channel: fmt.Sprintf("res.partner.%d", partnerID),
			Message: map[string]interface{}{
				"type":        "deletion",
				"message_ids": msgList,
			},
		})
	}
	h.BusBus().NewSet(rs.Env()).Sendmany(notifications)
	rs.Unlink()
}

// NotifyPendingByChat generates the bus notifications for the given message and send them
// to the appropriate moderators and the author (if the author has not been elected moderator
// meanwhile). The author notification can be considered as a feedback to the author.
func mailMessage_NotifyPendingByChat(rs m.MailMessageSet) {
	rs.EnsureOne()
	message := rs.MessageFormat()[0]
	partners := h.Partner().NewSet(rs.Env())
	for _, moderator := range h.MailChannel().BrowseOne(rs.Env(), rs.ResID()).Moderators().Records() {
		partners = partners.Union(moderator.Partner())
	}
	var notifications []*bustypes.Notification
	for _, partner := range partners.Records() {
		notifications = append(notifications, &bustypes.Notification{
			Channel: fmt.Sprintf("res.partner.%d", partner.ID()),
			Message: map[string]interface{}{
				"type":    "moderator",
				"message": message,
			},
		})
	}
	if rs.Author().Intersect(partners).IsEmpty() {
		notifications = append(notifications, &bustypes.Notification{
			Channel: fmt.Sprintf("res.partner.%d", rs.Author().ID()),
			Message: map[string]interface{}{
				"type":    "author",
				"message": message,
			},
		})
	}
	h.BusBus().NewSet(rs.Env()).Sendmany(notifications)
}

// NotifyModerators pushes a notification (Inbox/email) to moderators having messages
// waiting for moderation. This method is called once a day by a cron.
func mailMessage_NotifyModerators(rs m.MailMessageSet) {
	messages := h.MailMessage().Search(rs.Env(), q.MailMessage().
		ModerationStatus().Equals("pending_moderation"))
	channels := h.MailChannel().NewSet(rs.Env())
	moderatorsToNotify := h.User().NewSet(rs.Env())
	for _, message := range messages.Records() {
		channel := h.MailChannel().BrowseOne(rs.Env(), message.ResID())
		channels = channels.Union(channel)
		moderatorsToNotify = moderatorsToNotify.Union(channel.Moderators())
	}
	// FIXME
	// template := templates.Registry
	MailThread := h.MailThread().NewSet(rs.Env()).WithContext("mail_notify_author", true)
	//        for moderator in moderators_to_notify:
	//            MailThread.message_notify(
	//                moderator.partner_id.ids,
	//                # tocheck: target language
	//                subject=_('Message are pending moderation'),
	//                body=template.render(
	//                    {'record': moderator.partner_id}, engine='ir.qweb', minimal_qcontext=True),
	//                email_from=moderator.company_id.catchall or moderator.company_id.email)
	for _, moderator := range moderatorsToNotify.Records() {
		// FIXME
		MailThread.MessageNotify(moderator.Partner(), "", rs.T("Messages are pending moderation"))
	}
}

func mailMessage_NameGet(rs m.MailMessageSet) string {
	return rs.RecordName()
}

func init() {
	models.NewModel("MailMessage")
	h.MailMessage().AddFields(fields_MailMessage)
	h.MailMessage().SetDescription("Message")
	h.MailMessage().SetDefaultOrder("ID DESC")
	h.MailMessage().NewMethod("ComputeNeedaction", mailMessage_ComputeNeedaction)
	h.MailMessage().NewMethod("SearchNeedaction", mailMessage_SearchNeedaction)
	h.MailMessage().NewMethod("ComputeHasError", mailMessage_ComputeHasError)
	h.MailMessage().NewMethod("SearchHasError", mailMessage_SearchHasError)
	h.MailMessage().NewMethod("ComputeStarred", mailMessage_ComputeStarred)
	h.MailMessage().NewMethod("SearchStarred", mailMessage_SearchStarred)
	h.MailMessage().NewMethod("ComputeNeedModeration", mailMessage_ComputeNeedModeration)
	h.MailMessage().NewMethod("SearchNeedModeration", mailMessage_SearchNeedModeration)
	h.MailMessage().NewMethod("MarkAllAsRead", mailMessage_MarkAllAsRead)
	h.MailMessage().NewMethod("SetMessageDone", mailMessage_SetMessageDone)
	h.MailMessage().NewMethod("UnstarAll", mailMessage_UnstarAll)
	h.MailMessage().NewMethod("ToggleMessageStarred", mailMessage_ToggleMessageStarred)
	h.MailMessage().NewMethod("MessageReadDictPostprocess", mailMessage_MessageReadDictPostprocess)
	h.MailMessage().NewMethod("MessageFetchFailed", mailMessage_MessageFetchFailed)
	h.MailMessage().NewMethod("MessageFetch", mailMessage_MessageFetch)
	h.MailMessage().NewMethod("MessageFormat", mailMessage_MessageFormat)
	h.MailMessage().NewMethod("FormatMailFailures", mailMessage_FormatMailFailures)
	h.MailMessage().NewMethod("NotifyFailureUpdate", mailMessage_NotifyFailureUpdate)
	h.MailMessage().NewMethod("Init", mailMessage_Init)
	h.MailMessage().NewMethod("FindAllowedModelWise", mailMessage_FindAllowedModelWise)
	h.MailMessage().NewMethod("FindAllowedDocIds", mailMessage_FindAllowedDocIds)
	// h.MailMessage().Methods().Search().Extend(mailMessage_Search)
	h.MailMessage().NewMethod("CheckAccessRule", mailMessage_CheckAccessRule)
	h.MailMessage().NewMethod("GetRecordName", mailMessage_GetRecordName)
	h.MailMessage().NewMethod("GetReplyTo", mailMessage_GetReplyTo)
	h.MailMessage().NewMethod("GetMessageID", mailMessage_GetMessageID)
	h.MailMessage().NewMethod("InvalidateDocuments", mailMessage_InvalidateDocuments)
	h.MailMessage().Methods().Create().Extend(mailMessage_Create)
	h.MailMessage().Methods().Write().Extend(mailMessage_Write)
	h.MailMessage().Methods().Unlink().Extend(mailMessage_Unlink)
	h.MailMessage().NewMethod("Notify", mailMessage_Notify)
	h.MailMessage().NewMethod("NotifyComputeRecipients", mailMessage_NotifyComputeRecipients)
	h.MailMessage().NewMethod("NotifyRecipients", mailMessage_NotifyRecipients)
	h.MailMessage().NewMethod("Moderate", mailMessage_Moderate)
	h.MailMessage().NewMethod("DoModerate", mailMessage_DoModerate)
	h.MailMessage().NewMethod("ModerateAccept", mailMessage_ModerateAccept)
	h.MailMessage().NewMethod("ModerateSendRejectEmail", mailMessage_ModerateSendRejectEmail)
	h.MailMessage().NewMethod("SearchFromSameAuthors", mailMessage_SearchFromSameAuthors)
	h.MailMessage().NewMethod("ModerateDiscard", mailMessage_ModerateDiscard)
	h.MailMessage().NewMethod("NotifyPendingByChat", mailMessage_NotifyPendingByChat)
	h.MailMessage().NewMethod("NotifyModerators", mailMessage_NotifyModerators)
	h.MailMessage().Methods().NameGet().Extend(mailMessage_NameGet)
}
