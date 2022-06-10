package mail

import (
	"fmt"
	"strings"

	"github.com/hexya-addons/mail/mailtypes"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/tools/strutils"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// MailFollowers holds the data related to the follow mechanism inside
// Hexya. Partners can choose to follow documents (records) of any kind
// that inherits from MailThread. Following documents allow to receive
// notifications for new messages.
var fields_MailFollowers = map[string]models.FieldDefinition{
	"ResModel": fields.Char{
		String:   "Related Document Model Name",
		Required: true,
		Index:    true},

	"ResID": fields.Integer{
		String: "Related Document ID",
		Index:  true,
		Help:   "ID of the followed resource"},

	"Partner": fields.Many2One{
		RelationModel: h.Partner(),
		String:        "Related Partner",
		OnDelete:      models.Cascade,
		Index:         true},

	"Channel": fields.Many2One{
		RelationModel: h.MailChannel(),
		String:        "Listener",
		OnDelete:      models.Cascade,
		Index:         true},

	"Subtypes": fields.Many2Many{
		RelationModel: h.MailMessageSubtype(),
		String:        "Subtype",
		Help: "Message subtypes followed, meaning subtypes that will be" +
			"pushed onto the user's Wall."},
}

// Modifying followers change access rights to individual documents. As the
// cache may contain accessible/inaccessible data, one has to refresh it.

// InvalidateDocuments invalidates the cache of the documents followed by rs.
func mailFollowers_InvalidateDocuments(rs m.MailFollowersSet) {
	//        for record in self:
	//            if record.res_id:
	//                self.env[record.res_model].invalidate_cache(
	//                    ids=[record.res_id])
	for _, record := range rs.Records() {
		if record.ResID() != 0 {
			models.Registry.MustGet(rs.ModelName()).BrowseOne(rs.Env(), record.ResID()).InvalidateCache()
		}
	}
}

func mailFollowers_Create(rs m.MailFollowersSet, vals m.MailFollowersData) m.MailFollowersSet {
	res := rs.Super().Create(vals)
	res.InvalidateDocuments()
	return res
}

func mailFollowers_Write(rs m.MailFollowersSet, vals m.MailFollowersData) bool {
	if vals.HasResModel() || vals.HasResID() {
		rs.InvalidateDocuments()
	}
	res := rs.Super().Write(vals)
	if vals.HasResModel() || vals.HasResID() || vals.HasPartner() {
		rs.InvalidateDocuments()
	}
	return res
}

func mailFollowers_Unlink(rs m.MailFollowersSet) int64 {
	rs.InvalidateDocuments()
	return rs.Super().Unlink()
}

// GetRecipientData is a private method allowing to fetch recipients data based on a subtype.
// Purpose of this method is to fetch all data necessary to notify recipients
// in a single query. It fetches data from
//
// - followers (partners and channels) of records that follow the given subtype if records and subtype are set;
// - partners if pids is given;
// - channels if cids is given;
//
// Parameters:
//
// - records: fetch data from followers of records that follow subtype_id;
// - subtype: MailMessageSubtype to check against followers;
// - partners: additional set of partners from which to fetch recipient data;
// - cids: additional set of channels from which to fetch recipient data;
func mailFollowers_GetRecipientData(rs m.MailFollowersSet, records m.MailThreadSet, subtype m.MailMessageSubtypeSet,
	partners m.PartnerSet, channels m.MailChannelSet) []mailtypes.RecipientData {

	var res []mailtypes.RecipientData
	switch {
	case records.IsNotEmpty() && subtype.IsNotEmpty():
		var pidString string
		if partners.IsNotEmpty() {
			pidString = "OR partner.id IN (?)"
		}
		var cidString string
		if channels.IsNotEmpty() {
			cidString = "OR channel.id IN (?)"
		}
		query := fmt.Sprintf(`
WITH sub_followers AS (
    SELECT fol.id, fol.partner_id, fol.channel_id, subtype.internal
    FROM mail_followers fol
        RIGHT JOIN mail_followers_mail_message_subtype_rel subrel
        ON subrel.mail_followers_id = fol.id
        RIGHT JOIN mail_message_subtype subtype
        ON subtype.id = subrel.mail_message_subtype_id
    WHERE subrel.mail_message_subtype_id = %%s AND fol.res_model = %%s AND fol.res_id IN %%s
)
SELECT partner.id as pid, NULL AS cid,
        partner.active as active, partner.partner_share as pshare, NULL as ctype,
        users.notification_type AS notif, array_agg(groups.id) AS groups
    FROM res_partner partner
    LEFT JOIN res_users users ON users.partner_id = partner.id AND users.active
    LEFT JOIN res_groups_users_rel groups_rel ON groups_rel.uid = users.id
    LEFT JOIN res_groups groups ON groups.id = groups_rel.gid
    WHERE EXISTS (
        SELECT partner_id FROM sub_followers
        WHERE sub_followers.channel_id IS NULL
            AND sub_followers.partner_id = partner.id
            AND (coalesce(sub_followers.internal, false) <> TRUE OR coalesce(partner.partner_share, false) <> TRUE)
    ) %s
    GROUP BY partner.id, users.notification_type
UNION
SELECT NULL AS pid, channel.id AS cid,
       TRUE as active, NULL AS pshare, channel.channel_type AS ctype,
       CASE WHEN channel.email_send = TRUE THEN 'email' ELSE 'inbox' END AS notif, NULL AS groups
    FROM mail_channel channel
    WHERE EXISTS (
        SELECT channel_id FROM sub_followers WHERE partner_id IS NULL AND sub_followers.channel_id = channel.id
    ) %s`, pidString, cidString)
		params := []interface{}{subtype.ID(), records.ModelName(), records.Ids()}
		if partners.IsNotEmpty() {
			params = append(params, partners.Ids())
		}
		if channels.IsNotEmpty() {
			params = append(params, channels.Ids())
		}
		rs.Env().Cr().Select(res, query, params)
	case partners.IsNotEmpty() || channels.IsNotEmpty():
		var (
			params  []interface{}
			queries []string
		)
		if partners.IsNotEmpty() {
			queries = append(queries, `
SELECT partner.id as pid, NULL AS cid,
    partner.active as active, partner.partner_share as pshare, NULL as ctype,
    users.notification_type AS notif, NULL AS groups
FROM res_partner partner
LEFT JOIN res_users users ON users.partner_id = partner.id AND users.active
WHERE partner.id IN (?)`)
			params = append(params, partners.Ids())
		}
		if channels.IsNotEmpty() {
			queries = append(queries, `
SELECT NULL AS pid, channel.id AS cid,
    TRUE as active, NULL AS pshare, channel.channel_type AS ctype,
    CASE when channel.email_send = TRUE then 'email' else 'inbox' end AS notif, NULL AS groups
FROM mail_channel channel WHERE channel.id IN (?)`)
			params = append(params, channels.Ids())
		}
		query := strings.Join(queries, " UNION ")
		rs.Env().Cr().Select(res, query, params)
	}
	return res
}

// GetSubscriptionData is a private method allowing to fetch follower data from several
// documents of a given model.
// Followers can be filtered given partner IDs and channel IDs.
//
// Parameters:
//
// - resModel & resIds: the documents from which we want to have subscription data;
// - partners: optional partner to filter; if empty take all, otherwise limit to partners
// - channels: optional channel to filter; if empty take all, otherwise limit to channels
// - includePShare: optional join in partner to fetch their share status
func mailFollowers_GetSubscriptionData(rs m.MailFollowersSet, records m.MailThreadSet, partners m.PartnerSet,
	channels m.MailChannelSet, includePShare bool) []mailtypes.SubscriptionData {

	var (
		res          []mailtypes.SubscriptionData
		whereClauses []string
		whereParams  []interface{}
	)
	// base query: fetch followers of given documents
	for _, rec := range records.Records() {
		whereClauses = append(whereClauses, "fol.res_model = ? AND fol.res_id IN (?)")
		whereParams = append(whereParams, rec.ModelName(), rec.ID())
	}
	whereClause := strings.Join(whereClauses, " OR ")

	// additional: filter on optional partners / channels
	var subWhere []string
	if partners.IsNotEmpty() {
		subWhere = append(subWhere, "fol.partner_id IN (?)")
		whereParams = append(whereParams, partners.Ids())
	}
	if channels.IsNotEmpty() {
		subWhere = append(subWhere, "fol.channel_id IN (?)")
		whereParams = append(whereParams, channels.Ids())
	}
	if len(subWhere) > 0 {
		whereClause += fmt.Sprintf("AND (%s)", strings.Join(subWhere, " OR "))
	}
	var icFields, icJoin, icGroup string
	if includePShare {
		icFields = ", partner.partner_share"
		icJoin = "LEFT JOIN res_partner partner ON partner.id = fol.partner_id"
		icGroup = ", partner.partner_share"
	}
	query := fmt.Sprintf(`
SELECT fol.id, fol.res_id, fol.partner_id, fol.channel_id, array_agg(subtype.id)%s
FROM mail_followers fol
%s
LEFT JOIN mail_followers_mail_message_subtype_rel fol_rel ON fol_rel.mail_followers_id = fol.id
LEFT JOIN mail_message_subtype subtype ON subtype.id = fol_rel.mail_message_subtype_id
WHERE %s
GROUP BY fol.id%s`,
		icFields,
		icJoin,
		whereClause,
		icGroup,
	)
	rs.Env().Cr().Select(res, query, whereParams)
	return res
}

// InsertFollowers is the main internal method allowing to create or update followers
// for documents, given a resModel and the document resIds. This method
// does not handle access rights. This is the role of the caller to ensure there is
// no security breach.
//
// Parameters:
//
// - partnerSubtypes: optional subtypes for new partner followers.
//                    If not given, default ones are computed;
// - channelSubtypes: optional subtypes for new channel followers.
//                    If not given, default ones are computed;
// - customers: see AddDefaultFollowers;
// - checkExisting: see AddFollowers;
// - existingPolicy: see AddFollowers;
func mailFollowers_InsertFollowers(rs m.MailFollowersSet, resModel string, resIds []int64, partners m.PartnerSet,
	partnerSubtypes map[int64]m.MailMessageSubtypeSet, channels m.MailChannelSet, channelSubtypes map[int64]m.MailMessageSubtypeSet,
	customers m.PartnerSet, checkExisting bool, existingPolicy string) {
	sudoSelf := rs.Sudo()
	var (
		nData map[int64][]m.MailFollowersData
		uData map[int64]m.MailFollowersData
	)
	switch {
	case len(partnerSubtypes) == 0 && len(channelSubtypes) == 0:
		nData, uData = rs.AddDefaultFollowers(resModel, resIds, partners, channels, customers)
	default:
		nData, uData = rs.AddFollowers(resModel, resIds, partners, partnerSubtypes, channels, channelSubtypes, checkExisting, existingPolicy)
	}
	ctx := rs.Env().Context().Copy()
	if channels.IsNotEmpty() && rs.Env().Context().HasKey("default_partner_id") {
		ctx.Delete("default_partner_id")
	}
	for resID, values := range nData {
		for _, value := range values {
			sudoSelf.WithNewContext(ctx).Create(value.SetResID(resID))
		}
	}
	for folID, values := range uData {
		sudoSelf.BrowseOne(folID).Write(values)
	}
}

// AddDefaultFollowers id a shortcut to AddFollowers() that computes default subtypes.
// Existing followers are skipped as their subscription is considered as more important
// compared to new default subscription.
//
// Parameters:
//
// - customers: optional list of partners that are customers.
//              It is used if computing default subtype is necessary and allow to avoid
//              the check of partners being customers (no user or share user).
//              It is just a matter of saving queries if the info is already known;
//
// See also AddFollowers()
func mailFollowers_AddDefaultFollowers(rs m.MailFollowersSet, resModel string, resIds []int64, partners m.PartnerSet,
	channels m.MailChannelSet, customers m.PartnerSet) (map[int64][]m.MailFollowersData, map[int64]m.MailFollowersData) {

	if partners.IsEmpty() && channels.IsEmpty() {
		return make(map[int64][]m.MailFollowersData), make(map[int64]m.MailFollowersData)
	}
	subTypes, _, external := h.MailMessageSubtype().NewSet(rs.Env()).DefaultSubtypes(resModel)
	if partners.IsNotEmpty() && customers.IsEmpty() {
		customers = h.Partner().NewSet(rs.Env()).Sudo().Search(q.Partner().
			ID().In(partners.Ids()).And().
			PartnerShare().Equals(true))
	}
	cSTypes := make(map[int64]m.MailMessageSubtypeSet)
	for _, c := range channels.Records() {
		cSTypes[c.ID()] = subTypes
	}
	pSTypes := make(map[int64]m.MailMessageSubtypeSet)
	for _, p := range partners.Records() {
		st := subTypes
		if p.Intersect(customers).IsNotEmpty() {
			st = external
		}
		pSTypes[p.ID()] = st
	}
	return rs.AddFollowers(resModel, resIds, partners, pSTypes, channels, cSTypes, true, "skip")
}

// AddFollowers is an internal method that generates values to insert or update
// followers. Callers have to handle the result, for example by making a valid
// ORM command, inserting or updating directly follower records, ...
// This method returns two main data
//
// - first one is a dict which keys are resIds.
//   Value is a list of dict of values valid for creating new followers for the related res_id;
// - second one is a dict which keys are follower ids.
//   Value is a dict of values valid for updating the related follower record;
//
// Parameters:
//
// - checkExisting: if True, check for existing followers for given documents and handle
//                  them according to existing_policy parameter.
//                  Setting to False allows to save some computation if caller is sure
//                  there are no conflict for followers;
// - existingPolicy: if check_existing, tells what to do with already-existing followers:
//           * skip: simply skip existing followers, do not touch them;
//           * force: update existing with given subtypes only;
//           * replace: replace existing with nex subtypes (like force without old / new follower);
//           * update: gives an update dict allowing to add missing subtypes (no subtype removal);
func mailFollowers_AddFollowers(rs m.MailFollowersSet, records m.MailThreadSet, partners m.PartnerSet,
	partnerSubtypes map[int64]m.MailMessageSubtypeSet, channels m.MailChannelSet, channelSubtypes map[int64]m.MailMessageSubtypeSet,
	checkExisting bool, existingPolicy string) (map[int64][]m.MailFollowersData, map[int64]m.MailFollowersData) {

	rIds := records.Ids()
	if len(rIds) == 0 {
		rIds = []int64{0}
	}
	dataFols := make(map[int64]mailtypes.SubscriptionData)
	docPIds := make(map[int64]map[int64]bool)
	docCIds := make(map[int64]map[int64]bool)
	for _, i := range rIds {
		docPIds[i] = make(map[int64]bool)
		docCIds[i] = make(map[int64]bool)
	}
	var followerIds []int64
	if checkExisting && records.IsNotEmpty() {
		for _, sd := range rs.GetSubscriptionData(records, partners, channels, false) {
			if existingPolicy != "force" {
				switch {
				case sd.PartnerID != 0:
					docPIds[sd.ResID][sd.PartnerID] = true
				case sd.ChannelID != 0:
					docCIds[sd.ResID][sd.ChannelID] = true
				}
				dataFols[sd.FollowerID] = sd
				followerIds = append(followerIds, sd.FollowerID)
			}
		}
		if existingPolicy == "force" {
			h.MailFollowers().NewSet(rs.Env()).Sudo().Browse(followerIds).Unlink()
		}
	}

	nData := make(map[int64][]m.MailFollowersData)
	uData := make(map[int64]m.MailFollowersData)
	for _, resID := range rIds {
		for _, partner := range partners.Records() {
			switch {
			case !docPIds[resID][partner.ID()]:
				nData[resID] = append(nData[resID], h.MailFollowers().NewData().
					SetResModel(resModel).
					SetPartner(partner).
					SetSubtypes(partnerSubtypes[partner.ID()]))
			case strutils.IsIn(existingPolicy, "replace", "update"):
				var folID int64
				sIds := h.MailMessageSubtype().NewSet(rs.Env())
				for fID, val := range dataFols {
					if val.PartnerID == partner.ID() && val.ResID == resID {
						folID = fID
						sIds = h.MailMessageSubtype().Browse(rs.Env(), val.SubtypeIds)
						break
					}
				}
				newSids := partnerSubtypes[partner.ID()].Subtract(sIds)
				oldSids := sIds.Subtract(partnerSubtypes[partner.ID()])
				if folID != 0 && newSids.IsNotEmpty() {
					uData[folID] = h.MailFollowers().NewData().SetSubtypes(sIds.Union(newSids))
				}
				if folID != 0 && oldSids.IsNotEmpty() && existingPolicy == "replace" {
					uData[folID] = h.MailFollowers().NewData().SetSubtypes(sIds.Subtract(oldSids))
				}
			}
		}
		for _, channel := range channels.Records() {
			switch {
			case !docCIds[resID][channel.ID()]:
				nData[resID] = append(nData[resID], h.MailFollowers().NewData().
					SetResModel(resModel).
					SetChannel(channel).
					SetSubtypes(channelSubtypes[channel.ID()]))
			case strutils.IsIn(existingPolicy, "replace", "update"):
				var folID int64
				sIds := h.MailMessageSubtype().NewSet(rs.Env())
				for fID, val := range dataFols {
					if val.ChannelID == channel.ID() && val.ResID == resID {
						folID = fID
						sIds = h.MailMessageSubtype().Browse(rs.Env(), val.SubtypeIds)
						break
					}
				}
				newSids := channelSubtypes[channel.ID()].Subtract(sIds)
				oldSids := sIds.Subtract(channelSubtypes[channel.ID()])
				if folID != 0 && newSids.IsNotEmpty() {
					uData[folID] = h.MailFollowers().NewData().SetSubtypes(sIds.Union(newSids))
				}
				if folID != 0 && oldSids.IsNotEmpty() && existingPolicy == "replace" {
					uData[folID] = h.MailFollowers().NewData().SetSubtypes(sIds.Subtract(oldSids))
				}
			}
		}
	}
	return nData, uData
}

func mailFollowers_NameGet(rs m.MailFollowersSet) string {
	return rs.Partner().DisplayName()
}

func init() {
	models.NewModel("MailFollowers")
	h.MailFollowers().AddFields(fields_MailFollowers)
	h.MailFollowers().AddSQLConstraint("mail_followers_res_partner_res_model_id_uniq",
		"unique(res_model,res_id,partner_id)",
		"Error, a partner cannot follow twice the same object.")
	h.MailFollowers().AddSQLConstraint("mail_followers_res_channel_res_model_id_uniq",
		"unique(res_model,res_id,channel_id)",
		"Error, a channel cannot follow twice the same object.")
	h.MailFollowers().AddSQLConstraint("partner_xor_channel",
		"CHECK((partner_id IS NULL) != (channel_id IS NULL))",
		"Error: A follower must be either a partner or a channel (but not both).")
	h.MailFollowers().SetDescription("Document Followers")
	h.MailFollowers().NewMethod("InvalidateDocuments", mailFollowers_InvalidateDocuments)
	h.MailFollowers().Methods().Create().Extend(mailFollowers_Create)
	h.MailFollowers().Methods().Write().Extend(mailFollowers_Write)
	h.MailFollowers().Methods().Unlink().Extend(mailFollowers_Unlink)
	h.MailFollowers().NewMethod("GetRecipientData", mailFollowers_GetRecipientData)
	h.MailFollowers().NewMethod("GetSubscriptionData", mailFollowers_GetSubscriptionData)
	h.MailFollowers().NewMethod("InsertFollowers", mailFollowers_InsertFollowers)
	h.MailFollowers().NewMethod("AddDefaultFollowers", mailFollowers_AddDefaultFollowers)
	h.MailFollowers().NewMethod("AddFollowers", mailFollowers_AddFollowers)
	h.MailFollowers().Methods().NameGet().Extend(mailFollowers_NameGet)
}
