package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

//MODERATION_FIELDS = ['moderation', 'moderator_ids', 'moderation_ids', 'moderation_notify',
//'moderation_notify_msg', 'moderation_guidelines', 'moderation_guidelines_msg']

var fields_MailChannelPartner = map[string]models.FieldDefinition{
	"Partner": fields.Many2One{
		RelationModel: h.Partner(),
		String:        "Recipient",
		OnDelete:      models.Cascade},

	"PartnerEmail": fields.Char{
		String:   "Email",
		Related:  `PartnerId.Email`,
		ReadOnly: false},

	"Channel": fields.Many2One{
		RelationModel: h.MailChannel(),
		String:        "Channel",
		OnDelete:      models.Cascade},

	"SeenMessage": fields.Many2One{
		RelationModel: h.MailMessage(),
		String:        "Last Seen"},

	"FoldState": fields.Selection{
		Selection: types.Selection{
			"open":   "Open",
			"folded": "Folded",
			"closed": "Closed",
		},
		String:  "Conversation Fold State",
		Default: models.DefaultValue("open")},

	"IsMinimized": fields.Boolean{
		String: "Conversation is minimized"},

	"IsPinned": fields.Boolean{
		String:  "Is pinned on the interface",
		Default: models.DefaultValue(true)},
}

var fields_MailModeration = map[string]models.FieldDefinition{
	"Email": fields.Char{
		String:   "Email",
		Index:    true,
		Required: true},

	"Status": fields.Selection{
		Selection: types.Selection{
			"allow": "Always Allow",
			"ban":   "Permanent Ban",
		},
		String:   "Status",
		Required: true},

	"Channel": fields.Many2One{
		RelationModel: h.MailChannel(),
		String:        "Channel",
		Index:         true,
		Required:      true},
}

var fields_MailChannel = map[string]models.FieldDefinition{
	"Name": fields.Char{
		String:    "Name",
		Required:  true,
		Translate: true},

	"ChannelType": fields.Selection{
		Selection: types.Selection{
			"chat":    "Chat Discussion",
			"channel": "Channel",
		},
		String:  "Channel Type",
		Default: models.DefaultValue("channel")},

	"IsChat": fields.Boolean{
		String:  "Is a chat",
		Compute: h.MailChannel().Methods().ComputeIsChat(),
		Default: models.DefaultValue(false)},

	"Description": fields.Text{
		String: "Description"},

	"Uuid": fields.Char{
		String:  "UUID",
		Size:    50,
		Index:   true,
		Default: func(env models.Environment) interface{} { return str() },
		NoCopy:  true},

	"EmailSend": fields.Boolean{
		String:  "Send messages by email",
		Default: models.DefaultValue(false)},

	"ChannelLastSeenPartners": fields.One2Many{
		RelationModel: h.MailChannelPartner(),
		ReverseFK:     "",
		String:        "Last Seen"},

	"ChannelPartners": fields.Many2Many{
		RelationModel:    h.Partner(),
		M2MLinkModelName: "",
		M2MOurField:      "",
		M2MTheirField:    "",
		String:           "Listeners"},

	"ChannelMessages": fields.Many2Many{
		RelationModel:    h.MailMessage(),
		M2MLinkModelName: ""},

	"IsMember": fields.Boolean{
		String:  "Is a member",
		Compute: h.MailChannel().Methods().ComputeIsMember()},

	"Public": fields.Selection{
		Selection: types.Selection{
			"public":  "Everyone",
			"private": "Invited people only",
			"groups":  "Selected group of users",
		},
		String:   "Privacy",
		Required: true,
		Default:  models.DefaultValue("groups"),
		Help: "This group is visible by non members. Invisible groups" +
			"can add members through the invite button."},

	"GroupPublic": fields.Many2One{
		RelationModel: h.Group(),
		String:        "Authorized Group",
		Default:       func(env models.Environment) interface{} { return env.ref() }},

	"Groups": fields.Many2Many{
		RelationModel: h.Group(),
		String:        "Auto Subscription",
		Help: "Members of those groups will automatically added as followers." +
			"Note that they will be able to manage their subscription" +
			"manually if necessary."},

	"Image": fields.Binary{
		String:  "Photo",
		Default: models.DefaultValue(_get_default_image),
		// attachment=True
		Help: "This field holds the image used as photo for the group," +
			"limited to 1024x1024px."},

	"ImageMedium": fields.Binary{
		String: "Medium-sized photo",
		// attachment=True
		Help: "Medium-sized photo of the group. It is automatically resized" +
			"as a 128x128px image, with aspect ratio preserved. Use" +
			"this field in form views or some kanban views."},

	"ImageSmall": fields.Binary{
		String: "Small-sized photo",
		// attachment=True
		Help: "Small-sized photo of the group. It is automatically resized" +
			"as a 64x64px image, with aspect ratio preserved. Use this" +
			"field anywhere a small image is required."},

	"IsSubscribed": fields.Boolean{
		String:  "Is Subscribed",
		Compute: h.MailChannel().Methods().ComputeIsSubscribed()},

	"Moderation": fields.Boolean{
		String: "Moderate this channel"},

	"Moderators": fields.Many2Many{
		RelationModel:    h.User(),
		M2MLinkModelName: "",
		String:           "Moderators"},

	"IsModerator": fields.Boolean{
		Help:    "Current user is a moderator of the channel",
		String:  "Moderator",
		Compute: h.MailChannel().Methods().ComputeIsModerator()},

	"Moderations": fields.One2Many{
		RelationModel: h.MailModeration(),
		ReverseFK:     "",
		String:        "Moderated Emails",
		// groups="base.group_user"
	},

	"ModerationCount": fields.Integer{
		String:  "Moderated emails count",
		Compute: h.MailChannel().Methods().ComputeModerationCount(),
		// groups="base.group_user"
	},

	"ModerationNotify": fields.Boolean{
		String: "Automatic notification",
		Help: "People receive an automatic notification about their message" +
			"being waiting for moderation."},

	"ModerationNotifyMsg": fields.Text{
		String: "Notification message"},

	"ModerationGuidelines": fields.Boolean{
		String: "Send guidelines to new subscribers",
		Help: "Newcomers on this moderated channel will automatically" +
			"receive the guidelines."},

	"ModerationGuidelinesMsg": fields.Text{
		String: "Guidelines"},
}

// GetDefaultImage
func mailChannel_GetDefaultImage(rs m.MailChannelSet) {
	//        image_path = modules.get_module_resource(
	//            'mail', 'static/src/img', 'groupdefault.png')
	//        return tools.image_resize_image_big(base64.b64encode(open(image_path, 'rb').read()))
}

// DefaultGet
func mailChannel_DefaultGet(rs m.MailChannelSet, fields interface{}) {
	//        res = super(Channel, self).default_get(fields)
	//        if not res.get('alias_contact') and (not fields or 'alias_contact' in fields):
	//            res['alias_contact'] = 'everyone' if res.get(
	//                'public', 'private') == 'public' else 'followers'
	//        return res
}

// ComputeIsSubscribed
func mailChannel_ComputeIsSubscribed(rs m.MailChannelSet) m.MailChannelData {
	//        self.is_subscribed = self.env.user.partner_id in self.channel_partner_ids
}

// ComputeIsModerator
func mailChannel_ComputeIsModerator(rs m.MailChannelSet) m.MailChannelData {
	//        for channel in self:
	//            channel.is_moderator = self.env.user in channel.moderator_ids
}

// ComputeModerationCount
func mailChannel_ComputeModerationCount(rs m.MailChannelSet) m.MailChannelData {
	//        read_group_res = self.env['mail.moderation'].read_group(
	//            [('channel_id', 'in', self.ids)], ['channel_id'], 'channel_id')
	//        data = dict((res['channel_id'][0], res['channel_id_count'])
	//                    for res in read_group_res)
	//        for channel in self:
	//            channel.moderation_count = data.get(channel.id, 0)
}

// CheckModeratorEmail
func mailChannel_CheckModeratorEmail(rs m.MailChannelSet) {
	//        if any(not moderator.email for channel in self for moderator in channel.moderator_ids):
	//            raise ValidationError(_("Moderators must have an email address."))
}

// CheckModeratorIsMember
func mailChannel_CheckModeratorIsMember(rs m.MailChannelSet) {
	//        for channel in self:
	//            if not (channel.mapped('moderator_ids.partner_id') <= channel.sudo().channel_partner_ids):
	//                raise ValidationError(
	//                    _("Moderators should be members of the channel they moderate."))
}

// CheckModerationParameters
func mailChannel_CheckModerationParameters(rs m.MailChannelSet) {
	//        if any(not channel.email_send and channel.moderation for channel in self):
	//            raise ValidationError(_('Only mailing lists can be moderated.'))
}

// CheckModeratorExistence
func mailChannel_CheckModeratorExistence(rs m.MailChannelSet) {
	//        if any(not channel.moderator_ids for channel in self if channel.moderation):
	//            raise ValidationError(
	//                _('Moderated channels must have moderators.'))
}

// ComputeIsMember
func mailChannel_ComputeIsMember(rs m.MailChannelSet) m.MailChannelData {
	//        memberships = self.env['mail.channel.partner'].sudo().search([
	//            ('channel_id', 'in', self.ids),
	//            ('partner_id', '=', self.env.user.partner_id.id),
	//        ])
	//        membership_ids = memberships.mapped('channel_id')
	//        for record in self:
	//            record.is_member = record in membership_ids
}

// ComputeIsChat
func mailChannel_ComputeIsChat(rs m.MailChannelSet) m.MailChannelData {
	//        for record in self:
	//            if record.channel_type == 'chat':
	//                record.is_chat = True
}

// OnchangePublic
func mailChannel_OnchangePublic(rs m.MailChannelSet) {
	//        if self.public != 'public' and self.alias_contact == 'everyone':
	//            self.alias_contact = 'followers'
}

// OnchangeModeratorIds
func mailChannel_OnchangeModeratorIds(rs m.MailChannelSet) {
	//        missing_partners = self.mapped(
	//            'moderator_ids.partner_id') - self.mapped('channel_last_seen_partner_ids.partner_id')
	//        for partner in missing_partners:
	//            self.channel_last_seen_partner_ids += self.env['mail.channel.partner'].new(
	//                {'partner_id': partner.id})
}

// OnchangeEmailSend
func mailChannel_OnchangeEmailSend(rs m.MailChannelSet) {
	//        if not self.email_send:
	//            self.moderation = False
}

// OnchangeModeration
func mailChannel_OnchangeModeration(rs m.MailChannelSet) {
	//        if not self.moderation:
	//            self.moderation_notify = False
	//            self.moderation_guidelines = False
	//            self.moderator_ids = False
	//        else:
	//            self.moderator_ids |= self.env.user
}

// Create
func mailChannel_Create(rs m.MailChannelSet, vals models.RecordData) {
	//        if not vals.get('image'):
	//            defaults = self.default_get(['image'])
	//            vals['image'] = defaults['image']
	//        tools.image_resize_images(vals)
	//        channel = super(Channel, self.with_context(
	//            alias_model_name=self._name, alias_parent_model_name=self._name, mail_create_nolog=True, mail_create_nosubscribe=True)
	//        ).create(vals)
	//        channel.alias_id.write(
	//            {"alias_force_thread_id": channel.id, 'alias_parent_thread_id': channel.id})
	//        if vals.get('group_ids'):
	//            channel._subscribe_users()
	//        if not self._context.get('mail_channel_noautofollow'):
	//            channel.message_subscribe(channel_ids=[channel.id])
	//        return channel
}

// Unlink
func mailChannel_Unlink(rs m.MailChannelSet) {
	//        aliases = self.mapped('alias_id')
	//        try:
	//            all_emp_group = self.env.ref('mail.channel_all_employees')
	//        except ValueError:
	//            all_emp_group = None
	//        if all_emp_group and all_emp_group in self:
	//            raise UserError(
	//                _('You cannot delete those groups, as the Whole Company group is required by other modules.'))
	//        res = super(Channel, self).unlink()
	//        aliases.sudo().unlink()
	//        return res
}

// Write
func mailChannel_Write(rs m.MailChannelSet, vals models.RecordData) {
	//        if any(key for key in MODERATION_FIELDS if vals.get(key)) and any(self.env.user not in channel.moderator_ids for channel in self if channel.moderation):
	//            if not self.env.user.has_group('base.group_system'):
	//                raise UserError(
	//                    _("You do not have the rights to modify fields related to moderation on one of the channels you are modifying."))
	//        tools.image_resize_images(vals)
	//        result = super(Channel, self).write(vals)
	//        if vals.get('group_ids'):
	//            self._subscribe_users()
	//        if vals.get('moderation') is False:
	//            self.env['mail.message'].search([
	//                ('moderation_status', '=', 'pending_moderation'),
	//                ('model', '=', 'mail.channel'),
	//                ('res_id', 'in', self.ids)
	//            ])._moderate_accept()
	//        return result
}

// GetAliasModelName
func mailChannel_GetAliasModelName(rs m.MailChannelSet, vals interface{}) {
	//        return vals.get('alias_model', 'mail.channel')
}

// SubscribeUsers
func mailChannel_SubscribeUsers(rs m.MailChannelSet) {
	//        for mail_channel in self:
	//            mail_channel.write({'channel_partner_ids': [(4, pid) for pid in mail_channel.mapped(
	//                'group_ids').mapped('users').mapped('partner_id').ids]})
}

// ActionFollow
func mailChannel_ActionFollow(rs m.MailChannelSet) {
	//        self.ensure_one()
	//        channel_partner = self.mapped('channel_last_seen_partner_ids').filtered(
	//            lambda cp: cp.partner_id == self.env.user.partner_id)
	//        if not channel_partner:
	//            return self.write({'channel_last_seen_partner_ids': [(0, 0, {'partner_id': self.env.user.partner_id.id})]})
	//        return False
}

// ActionUnfollow
func mailChannel_ActionUnfollow(rs m.MailChannelSet) {
	//        return self._action_unfollow(self.env.user.partner_id)
}

// ActionUnfollow
func mailChannel_ActionUnfollow(rs m.MailChannelSet, partner interface{}) {
	//        channel_info = self.channel_info('unsubscribe')[0]
	//        result = self.write({'channel_partner_ids': [(3, partner.id)]})
	//        self.env['bus.bus'].sendone(
	//            (self._cr.dbname, 'res.partner', partner.id), channel_info)
	//        if not self.email_send:
	//            notification = _(
	//                '<div class="o_mail_notification">left <a href="#" class="o_channel_redirect" data-oe-id="%s">#%s</a></div>') % (self.id, self.name)
	//            # post 'channel left' message as root since the partner just unsubscribed from the channel
	//            self.sudo().message_post(body=notification, message_type="notification",
	//                                     subtype="mail.mt_comment", author_id=partner.id)
	//        return result
}

//  All recipients of a message on a channel are considered as partners.
//         This means they will receive a minimal email, without
// a link to access
//         in the backend. Mailing lists should indeed send
// minimal emails to avoid
//         the noise.
func mailChannel_NotifyGetGroups(rs m.MailChannelSet, message interface{}, groups interface{}) {
	//        groups = super(Channel, self)._notify_get_groups(message, groups)
	//        for (index, (group_name, group_func, group_data)) in enumerate(groups):
	//            if group_name != 'customer':
	//                groups[index] = (group_name, lambda partner: False, group_data)
	//        return groups
}

// NotifySpecificEmailValues
func mailChannel_NotifySpecificEmailValues(rs m.MailChannelSet, message interface{}) {
	//        res = super(Channel, self)._notify_specific_email_values(message)
	//        try:
	//            headers = safe_eval(res.get('headers', dict()))
	//        except Exception:
	//            headers = {}
	//        headers['Precedence'] = 'list'
	//        headers['X-Auto-Response-Suppress'] = 'OOF'
	//        if self.alias_domain and self.alias_name:
	//            headers['List-Id'] = '<%s.%s>' % (self.alias_name,
	//                                              self.alias_domain)
	//            headers['List-Post'] = '<mailto:%s@%s>' % (
	//                self.alias_name, self.alias_domain)
	//            # Avoid users thinking it was a personal message
	//            # X-Forge-To: will replace To: after SMTP envelope is determined by ir.mail.server
	//            list_to = '"%s" <%s@%s>' % (
	//                self.name, self.alias_name, self.alias_domain)
	//            headers['X-Forge-To'] = list_to
	//        res['headers'] = repr(headers)
	//        return res
}

//  Override bounce management to unsubscribe bouncing addresses
func mailChannel_MessageReceiveBounce(rs m.MailChannelSet, email interface{}, partner interface{}, mail_id interface{}) {
	//        for p in partner:
	//            if p.message_bounce >= self.MAX_BOUNCE_LIMIT:
	//                self._action_unfollow(p)
	//        return super(Channel, self).message_receive_bounce(email, partner, mail_id=mail_id)
}

// NotifyEmailRecipients
func mailChannel_NotifyEmailRecipients(rs m.MailChannelSet, message interface{}, recipient_ids interface{}) {
	//        whitelist = self.env['res.partner'].sudo().search(
	//            [('id', 'in', recipient_ids)]).filtered(lambda p: not p.is_blacklisted)
	//        if self.alias_domain and self.alias_name:
	//            return {
	//                'email_to': ','.join(formataddr((partner.name, partner.email)) for partner in whitelist if partner.email),
	//                'recipient_ids': [],
	//            }
	//        return super(Channel, self)._notify_email_recipients(message, whitelist.ids)
}

//  This method is used to compute moderation status before the creation
//         of a message.  For this operation the message's
// author email address is required.
//         This address is returned with status for other computations.
func mailChannel_ExtractModerationValues(rs m.MailChannelSet, message_type interface{}) {
	//        moderation_status = 'accepted'
	//        email = ''
	//        if self.moderation and message_type in ['email', 'comment']:
	//            author_id = kwargs.get('author_id')
	//            if author_id and isinstance(author_id, pycompat.integer_types):
	//                email = self.env['res.partner'].browse([author_id]).email
	//            elif author_id:
	//                email = author_id.email
	//            elif kwargs.get('email_from'):
	//                email = tools.email_split(kwargs['email_from'])[0]
	//            else:
	//                email = self.env.user.email
	//            if email in self.mapped('moderator_ids.email'):
	//                return moderation_status, email
	//            status = self.env['mail.moderation'].sudo().search(
	//                [('email', '=', email), ('channel_id', 'in', self.ids)]).mapped('status')
	//            if status and status[0] == 'allow':
	//                moderation_status = 'accepted'
	//            elif status and status[0] == 'ban':
	//                moderation_status = 'rejected'
	//            else:
	//                moderation_status = 'pending_moderation'
	//        return moderation_status, email
}

// MessagePost
func mailChannel_MessagePost(rs m.MailChannelSet, message_type interface{}) {
	//        moderation_status, email = self._extract_moderation_values(
	//            message_type, **kwargs)
	//        if moderation_status == 'rejected':
	//            return self.env['mail.message']
	//        self.filtered(lambda channel: channel.is_chat).mapped(
	//            'channel_last_seen_partner_ids').write({'is_pinned': True})
	//        message = super(Channel, self.with_context(mail_create_nosubscribe=True)).message_post(
	//            message_type=message_type, moderation_status=moderation_status, **kwargs)
	//        if self.moderation_notify and self.moderation_notify_msg and message_type == 'email' and moderation_status == 'pending_moderation':
	//            self.env['mail.mail'].create({
	//                'body_html': self.moderation_notify_msg,
	//                'subject': 'Re: %s' % (kwargs.get('subject', '')),
	//                'email_to': email,
	//                'auto_delete': True,
	//                'state': 'outgoing'
	//            })
	//        return message
}

// AliasCheckContact
func mailChannel_AliasCheckContact(rs m.MailChannelSet, message interface{}, message_dict interface{}, alias interface{}) {
	//        if alias.alias_contact == 'followers' and self.ids:
	//            author = self.env['res.partner'].browse(
	//                message_dict.get('author_id', False))
	//            if not author or author not in self.channel_partner_ids:
	//                return {
	//                    'error_message': _('restricted to channel members'),
	//                }
	//            return True
	//        return super(Channel, self)._alias_check_contact(message, message_dict, alias)
}

// Init
func mailChannel_Init(rs m.MailChannelSet) {
	//        self._cr.execute('SELECT indexname FROM pg_indexes WHERE indexname = %s',
	//                         ('mail_channel_partner_seen_message_id_idx'))
	//        if not self._cr.fetchone():
	//            self._cr.execute(
	//                'CREATE INDEX mail_channel_partner_seen_message_id_idx ON mail_channel_partner (channel_id,partner_id,seen_message_id)')
}

//  Send guidelines to all channel members.
func mailChannel_SendGuidelines(rs m.MailChannelSet) {
	//        if self.env.user in self.moderator_ids or self.env.user.has_group('base.group_system'):
	//            success = self._send_guidelines(self.channel_partner_ids)
	//            if not success:
	//                raise UserError(
	//                    _('View "mail.mail_channel_send_guidelines" was not found. No email has been sent. Please contact an administrator to fix this issue.'))
	//        else:
	//            raise UserError(
	//                _("Only an administrator or a moderator can send guidelines to channel members!"))
}

//  Send guidelines of a given channel. Returns False if template
// used for guidelines
//         not found. Caller may have to handle this return value.
func mailChannel_SendGuidelines(rs m.MailChannelSet, partners interface{}) {
	//        self.ensure_one()
	//        view = self.env.ref(
	//            'mail.mail_channel_send_guidelines', raise_if_not_found=False)
	//        if not view:
	//            _logger.warning(
	//                'View "mail.mail_channel_send_guidelines" was not found.')
	//            return False
	//        banned_emails = self.env['mail.moderation'].sudo().search([
	//            ('status', '=', 'ban'),
	//            ('channel_id', 'in', self.ids)
	//        ]).mapped('email')
	//        for partner in partners.filtered(lambda p: p.email and not (p.email in banned_emails)):
	//            create_values = {
	//                'body_html': view.render({'channel': self, 'partner': partner}, engine='ir.qweb', minimal_qcontext=True),
	//                'subject': _("Guidelines of channel %s") % self.name,
	//                'email_from': partner.company_id.catchall or partner.company_id.email,
	//                'recipient_ids': [(4, partner.id)]
	//            }
	//            mail = self.env['mail.mail'].create(create_values)
	//            mail.send()
	//        return True
}

//  This method adds emails into either white or black of
// the channel list of emails
//             according to status. If an email in emails
// is already moderated, the method updates the email status.
//             :param emails: list of email addresses to put
// in white or black list of channel.
//             :param status: value is 'allow' or 'ban'. Emails
// are put in white list if 'allow', in black list if 'ban'.
//
func mailChannel_UpdateModerationEmail(rs m.MailChannelSet, emails []string, status string) bool {
	//        self.ensure_one()
	//        splitted_emails = [tools.email_split(
	//            email)[0] for email in emails if tools.email_split(email)]
	//        moderated = self.env['mail.moderation'].sudo().search([
	//            ('email', 'in', splitted_emails),
	//            ('channel_id', 'in', self.ids)
	//        ])
	//        cmds = [(1, record.id, {'status': status}) for record in moderated]
	//        not_moderated = [
	//            email for email in splitted_emails if email not in moderated.mapped('email')]
	//        cmds += [(0, 0, {'email': email, 'status': status})
	//                 for email in not_moderated]
	//        return self.write({'moderation_ids': cmds})
}

//  Broadcast the current channel header to the given partner ids
//             :param partner_ids : the partner to notify
//
func mailChannel_Broadcast(rs m.MailChannelSet, partner_ids interface{}) {
	//        notifications = self._channel_channel_notifications(partner_ids)
	//        self.env['bus.bus'].sendmany(notifications)
}

//  Generate the bus notifications of current channel for
// the given partner ids
//             :param partner_ids : the partner to send the
// current channel header
//             :returns list of bus notifications (tuple (bus_channe,
// message_content))
//
func mailChannel_ChannelChannelNotifications(rs m.MailChannelSet, partner_ids interface{}) {
	//        notifications = []
	//        for partner in self.env['res.partner'].browse(partner_ids):
	//            user_id = partner.user_ids and partner.user_ids[0] or False
	//            if user_id:
	//                for channel_info in self.sudo(user_id).channel_info():
	//                    notifications.append(
	//                        [(self._cr.dbname, 'res.partner', partner.id), channel_info])
	//        return notifications
}

//  Broadcast the given message on the current channels.
//             Send the message on the Bus Channel (uuid for
// public mail.channel, and partner private bus channel (the tuple)).
//             A partner will receive only on message on its
// bus channel, even if this message belongs to multiple mail
// channel. Then 'channel_ids' field
//             of the received message indicates on wich mail
// channel the message should be displayed.
//             :param : mail.message to broadcast
//
func mailChannel_Notify(rs m.MailChannelSet, message interface{}) {
	//        if not self:
	//            return
	//        message.ensure_one()
	//        notifications = self._channel_message_notifications(message)
	//        self.env['bus.bus'].sendmany(notifications)
}

//  Generate the bus notifications for the given message
//             :param message : the mail.message to sent
//             :returns list of bus notifications (tuple (bus_channe,
// message_content))
//
func mailChannel_ChannelMessageNotifications(rs m.MailChannelSet, message interface{}) {
	//        message_values = message.message_format()[0]
	//        notifications = []
	//        for channel in self:
	//            notifications.append(
	//                [(self._cr.dbname, 'mail.channel', channel.id), dict(message_values)])
	//            # add uuid to allow anonymous to listen
	//            if channel.public == 'public':
	//                notifications.append([channel.uuid, dict(message_values)])
	//        return notifications
}

//  Get the informations header for the current channels
//             :returns a list of channels values
//             :rtype : list(dict)
//
func mailChannel_ChannelInfo(rs m.MailChannelSet, extra_info interface{}) {
	//        channel_infos = []
	//        partner_channels = self.env['mail.channel.partner']
	//        if self.env.user and self.env.user.partner_id:
	//            partner_channels = self.env['mail.channel.partner'].search(
	//                [('partner_id', '=', self.env.user.partner_id.id), ('channel_id', 'in', self.ids)])
	//        for channel in self:
	//            info = {
	//                'id': channel.id,
	//                'name': channel.name,
	//                'uuid': channel.uuid,
	//                'state': 'open',
	//                'is_minimized': False,
	//                'channel_type': channel.channel_type,
	//                'public': channel.public,
	//                'mass_mailing': channel.email_send,
	//                'moderation': channel.moderation,
	//                'is_moderator': self.env.uid in channel.moderator_ids.ids,
	//                'group_based_subscription': bool(channel.group_ids),
	//                'create_uid': channel.create_uid.id,
	//            }
	//            if extra_info:
	//                info['info'] = extra_info
	//            # add the partner for 'direct mesage' channel
	//            if channel.channel_type == 'chat':
	//                info['direct_partner'] = (channel.sudo()
	//                                          .with_context(active_test=False)
	//                                          .channel_partner_ids
	//                                          .filtered(lambda p: p.id != self.env.user.partner_id.id)
	//                                          .read(['id', 'name', 'im_status']))
	//
	//            # add last message preview (only used in mobile)
	//            if self._context.get('isMobile', False):
	//                last_message = channel.channel_fetch_preview()
	//                if last_message:
	//                    info['last_message'] = last_message[0].get('last_message')
	//
	//            # add user session state, if available and if user is logged
	//            if partner_channels.ids:
	//                partner_channel = partner_channels.filtered(
	//                    lambda c: channel.id == c.channel_id.id)
	//                if len(partner_channel) >= 1:
	//                    partner_channel = partner_channel[0]
	//                    info['state'] = partner_channel.fold_state or 'open'
	//                    info['is_minimized'] = partner_channel.is_minimized
	//                    info['seen_message_id'] = partner_channel.seen_message_id.id
	//                # add needaction and unread counter, since the user is logged
	//                info['message_needaction_counter'] = channel.message_needaction_counter
	//                info['message_unread_counter'] = channel.message_unread_counter
	//            channel_infos.append(info)
	//        return channel_infos
}

//  Return message values of the current channel.
//             :param last_id : last message id to start the research
//             :param limit : maximum number of messages to fetch
//             :returns list of messages values
//             :rtype : list(dict)
//
func mailChannel_ChannelFetchMessage(rs m.MailChannelSet, last_id interface{}, limit interface{}) {
	//        self.ensure_one()
	//        domain = [("channel_ids", "in", self.ids)]
	//        if last_id:
	//            domain.append(("id", "<", last_id))
	//        return self.env['mail.message'].message_fetch(domain=domain, limit=limit)
}

//  Get the canonical private channel between some partners,
// create it if needed.
//             To reuse an old channel (conversation), this
// one must be private, and contains
//             only the given partners.
//             :param partners_to : list of res.partner ids
// to add to the conversation
//             :param pin : True if getting the channel should
// pin it for the current user
//             :returns a channel header, or False if the
// users_to was False
//             :rtype : dict
//
func mailChannel_ChannelGet(rs m.MailChannelSet, partners_to interface{}, pin interface{}) {
	//        if partners_to:
	//            partners_to.append(self.env.user.partner_id.id)
	//            # determine type according to the number of partner in the channel
	//            self.env.cr.execute("""
	//                SELECT P.channel_id as channel_id
	//                FROM mail_channel C, mail_channel_partner P
	//                WHERE P.channel_id = C.id
	//                    AND C.public LIKE 'private'
	//                    AND P.partner_id IN %s
	//                    AND channel_type LIKE 'chat'
	//                GROUP BY P.channel_id
	//                HAVING array_agg(P.partner_id ORDER BY P.partner_id) = %s
	//            """, (tuple(partners_to), sorted(list(partners_to))))
	//            result = self.env.cr.dictfetchall()
	//            if result:
	//                # get the existing channel between the given partners
	//                channel = self.browse(result[0].get('channel_id'))
	//                # pin up the channel for the current partner
	//                if pin:
	//                    self.env['mail.channel.partner'].search(
	//                        [('partner_id', '=', self.env.user.partner_id.id), ('channel_id', '=', channel.id)]).write({'is_pinned': True})
	//            else:
	//                # create a new one
	//                channel = self.create({
	//                    'channel_partner_ids': [(4, partner_id) for partner_id in partners_to],
	//                    'public': 'private',
	//                    'channel_type': 'chat',
	//                    'email_send': False,
	//                    'name': ', '.join(self.env['res.partner'].sudo().browse(partners_to).mapped('name')),
	//                })
	//                # broadcast the channel header to the other partner (not me)
	//                channel._broadcast(partners_to)
	//            return channel.channel_info()[0]
	//        return False
}

// ChannelGetAndMinimize
func mailChannel_ChannelGetAndMinimize(rs m.MailChannelSet, partners_to interface{}) {
	//        channel = self.channel_get(partners_to)
	//        if channel:
	//            self.channel_minimize(channel['uuid'])
	//        return channel
}

//  Update the fold_state of the given session. In order to
// syncronize web browser
//             tabs, the change will be broadcast to himself
// (the current user channel).
//             Note: the user need to be logged
//             :param state : the new status of the session
// for the current user.
//
func mailChannel_ChannelFold(rs m.MailChannelSet, uuid interface{}, state interface{}) {
	//        domain = [('partner_id', '=', self.env.user.partner_id.id),
	//                  ('channel_id.uuid', '=', uuid)]
	//        for session_state in self.env['mail.channel.partner'].search(domain):
	//            if not state:
	//                state = session_state.fold_state
	//                if session_state.fold_state == 'open':
	//                    state = 'folded'
	//                else:
	//                    state = 'open'
	//            session_state.write({
	//                'fold_state': state,
	//                'is_minimized': bool(state != 'closed'),
	//            })
	//            self.env['bus.bus'].sendone(
	//                (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), session_state.channel_id.channel_info()[0])
}

// ChannelMinimize
func mailChannel_ChannelMinimize(rs m.MailChannelSet, uuid interface{}, minimized interface{}) {
	//        values = {
	//            'fold_state': minimized and 'open' or 'closed',
	//            'is_minimized': minimized
	//        }
	//        domain = [('partner_id', '=', self.env.user.partner_id.id),
	//                  ('channel_id.uuid', '=', uuid)]
	//        channel_partners = self.env['mail.channel.partner'].search(
	//            domain, limit=1)
	//        channel_partners.write(values)
	//        self.env['bus.bus'].sendone(
	//            (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), channel_partners.channel_id.channel_info()[0])
}

// ChannelPin
func mailChannel_ChannelPin(rs m.MailChannelSet, uuid interface{}, pinned interface{}) {
	//        channel = self.search([('uuid', '=', uuid)])
	//        channel_partners = self.env['mail.channel.partner'].search(
	//            [('partner_id', '=', self.env.user.partner_id.id), ('channel_id', '=', channel.id)])
	//        if not pinned:
	//            self.env['bus.bus'].sendone(
	//                (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), channel.channel_info('unsubscribe')[0])
	//        if channel_partners:
	//            channel_partners.write({'is_pinned': pinned})
}

// ChannelSeen
func mailChannel_ChannelSeen(rs m.MailChannelSet) {
	//        self.ensure_one()
	//        if self.channel_message_ids.ids:
	//            # zero is the index of the last message
	//            last_message_id = self.channel_message_ids.ids[0]
	//            self.env['mail.channel.partner'].search([('channel_id', 'in', self.ids), (
	//                'partner_id', '=', self.env.user.partner_id.id)]).write({'seen_message_id': last_message_id})
	//            self.env['bus.bus'].sendone((self._cr.dbname, 'res.partner', self.env.user.partner_id.id), {
	//                                        'info': 'channel_seen', 'id': self.id, 'last_message_id': last_message_id})
	//            return last_message_id
}

//  Add the given partner_ids to the current channels and
// broadcast the channel header to them.
//             :param partner_ids : list of partner id to add
//
func mailChannel_ChannelInvite(rs m.MailChannelSet, partner_ids interface{}) {
	//        partners = self.env['res.partner'].browse(partner_ids)
	//        for channel in self:
	//            partners_to_add = partners - channel.channel_partner_ids
	//            channel.write({'channel_last_seen_partner_ids': [
	//                          (0, 0, {'partner_id': partner_id}) for partner_id in partners_to_add.ids]})
	//            for partner in partners_to_add:
	//                if partner.id != self.env.user.partner_id.id:
	//                    notification = _('<div class="o_mail_notification">%(author)s invited %(new_partner)s to <a href="#" class="o_channel_redirect" data-oe-id="%(channel_id)s">#%(channel_name)s</a></div>') % {
	//                        'author': self.env.user.display_name,
	//                        'new_partner': partner.display_name,
	//                        'channel_id': channel.id,
	//                        'channel_name': channel.name,
	//                    }
	//                else:
	//                    notification = _(
	//                        '<div class="o_mail_notification">joined <a href="#" class="o_channel_redirect" data-oe-id="%s">#%s</a></div>') % (channel.id, channel.name)
	//                self.message_post(body=notification, message_type="notification",
	//                                  subtype="mail.mt_comment", author_id=partner.id)
	//        self._broadcast(partner_ids)
}

//  Broadcast the typing notification to channel members
//             :param is_typing: (boolean) tells whether the
// current user is typing or not
//             :param is_website_user: (boolean) tells whether
// the user that notifies comes
//               from the website-side. This is useful in
// order to distinguish operator and
//               unlogged users for livechat, because unlogged
// users have the same
//               partner_id as the admin (default: False).
//
func mailChannel_NotifyTyping(rs m.MailChannelSet, is_typing interface{}, is_website_user interface{}) {
	//        notifications = []
	//        for channel in self:
	//            data = {
	//                'info': 'typing_status',
	//                'is_typing': is_typing,
	//                'is_website_user': is_website_user,
	//                'partner_id': self.env.user.partner_id.id,
	//            }
	//            # notify backend users
	//            notifications.append(
	//                [(self._cr.dbname, 'mail.channel', channel.id), data])
	//            notifications.append([channel.uuid, data])  # notify frontend users
	//        self.env['bus.bus'].sendmany(notifications)
}

//  Return the channels of the user grouped by 'slot' (channel,
// direct_message or private_group), and
//             the mapping between partner_id/channel_id for
// direct_message channels.
//             :returns dict : the grouped channels and the mapping
//
func mailChannel_ChannelFetchSlot(rs m.MailChannelSet) {
	//        values = {}
	//        my_partner_id = self.env.user.partner_id.id
	//        pinned_channels = self.env['mail.channel.partner'].search(
	//            [('partner_id', '=', my_partner_id), ('is_pinned', '=', True)]).mapped('channel_id')
	//        values['channel_channel'] = self.search([('channel_type', '=', 'channel'), ('public', 'in', [
	//                                                'public', 'groups']), ('channel_partner_ids', 'in', [my_partner_id])]).channel_info()
	//        direct_message_channels = self.search(
	//            [('channel_type', '=', 'chat'), ('id', 'in', pinned_channels.ids)])
	//        values['channel_direct_message'] = direct_message_channels.channel_info()
	//        values['channel_private_group'] = self.search([('channel_type', '=', 'channel'), (
	//            'public', '=', 'private'), ('channel_partner_ids', 'in', [my_partner_id])]).channel_info()
	//        return values
}

//  Return the channel info of the channel the current partner can join
//             :param name : the name of the researched channels
//             :param domain : the base domain of the research
//             :returns dict : channel dict
//
func mailChannel_ChannelSearchToJoin(rs m.MailChannelSet, name interface{}, domain interface{}) {
	//        if not domain:
	//            domain = []
	//        domain = expression.AND([
	//            [('channel_type', '=', 'channel')],
	//            [('channel_partner_ids', 'not in', [self.env.user.partner_id.id])],
	//            [('public', '!=', 'private')],
	//            domain
	//        ])
	//        if name:
	//            domain = expression.AND(
	//                [domain, [('name', 'ilike', '%'+name+'%')]])
	//        return self.search(domain).read(['name', 'public', 'uuid', 'channel_type'])
}

// ChannelJoinAndGetInfo
func mailChannel_ChannelJoinAndGetInfo(rs m.MailChannelSet) {
	//        self.ensure_one()
	//        added = self.action_follow()
	//        if added and self.channel_type == 'channel' and not self.email_send:
	//            notification = _(
	//                '<div class="o_mail_notification">joined <a href="#" class="o_channel_redirect" data-oe-id="%s">#%s</a></div>') % (self.id, self.name)
	//            self.message_post(
	//                body=notification, message_type="notification", subtype="mail.mt_comment")
	//        if added and self.moderation_guidelines:
	//            self._send_guidelines(self.env.user.partner_id)
	//        channel_info = self.channel_info('join')[0]
	//        self.env['bus.bus'].sendone(
	//            (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), channel_info)
	//        return channel_info
}

//  Create a channel and add the current partner, broadcast
// it (to make the user directly
//             listen to it when polling)
//             :param name : the name of the channel to create
//             :param privacy : privacy of the channel. Should
// be 'public' or 'private'.
//             :return dict : channel header
//
func mailChannel_ChannelCreate(rs m.MailChannelSet, name interface{}, privacy interface{}) {
	//        new_channel = self.create({
	//            'name': name,
	//            'public': privacy,
	//            'email_send': False,
	//            'channel_partner_ids': [(4, self.env.user.partner_id.id)]
	//        })
	//        notification = _('<div class="o_mail_notification">created <a href="#" class="o_channel_redirect" data-oe-id="%s">#%s</a></div>') % (
	//            new_channel.id, new_channel.name)
	//        new_channel.message_post(
	//            body=notification, message_type="notification", subtype="mail.mt_comment")
	//        channel_info = new_channel.channel_info('creation')[0]
	//        self.env['bus.bus'].sendone(
	//            (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), channel_info)
	//        return channel_info
}

//  Return 'limit'-first channels' id, name and public fields
// such that the name matches a
//             'search' string. Exclude channels of type chat
// (DM), and private channels the current
//             user isn't registered to.
func mailChannel_GetMentionSuggestions(rs m.MailChannelSet, search interface{}, limit interface{}) {
	//        domain = expression.AND([
	//            [('name', 'ilike', search)],
	//            [('channel_type', '=', 'channel')],
	//            expression.OR([
	//                [('public', '!=', 'private')],
	//                [('channel_partner_ids', 'in', [
	//                    self.env.user.partner_id.id])]
	//            ])
	//        ])
	//        return self.search_read(domain, ['id', 'name', 'public'], limit=limit)
}

//  Return the id, name and email of partners listening to
// the given channel
func mailChannel_ChannelFetchListeners(rs m.MailChannelSet, uuid interface{}) {
	//        self._cr.execute("""
	//            SELECT P.id, P.name, P.email
	//            FROM mail_channel_partner CP
	//                INNER JOIN res_partner P ON CP.partner_id = P.id
	//                INNER JOIN mail_channel C ON CP.channel_id = C.id
	//            WHERE C.uuid = %s""", (uuid))
	//        return self._cr.dictfetchall()
}

// ChannelFetchListenersWhereClause
func mailChannel_ChannelFetchListenersWhereClause(rs m.MailChannelSet, uuid interface{}) {
	//        return ("C.uuid = %s", (uuid))
}

//  Return the last message of the given channels
func mailChannel_ChannelFetchPreview(rs m.MailChannelSet) {
	//        self._cr.execute("""
	//            SELECT mail_channel_id AS id, MAX(mail_message_id) AS message_id
	//            FROM mail_message_mail_channel_rel
	//            WHERE mail_channel_id IN %s
	//            GROUP BY mail_channel_id
	//            """, (tuple(self.ids)))
	//        channels_preview = dict((r['message_id'], r)
	//                                for r in self._cr.dictfetchall())
	//        last_messages = self.env['mail.message'].browse(
	//            channels_preview).message_format()
	//        for message in last_messages:
	//            channel = channels_preview[message['id']]
	//            del(channel['message_id'])
	//            channel['last_message'] = message
	//        return list(channels_preview.values())
}

//  Returns the allowed commands in channels
func mailChannel_GetMentionCommands(rs m.MailChannelSet) {
	//        commands = []
	//        for n in dir(self):
	//            match = re.search('^_define_command_(.+?)$', n)
	//            if match:
	//                command = getattr(self, n)()
	//                command['name'] = match.group(1)
	//                commands.append(command)
	//        return commands
}

//  Executes a given command
func mailChannel_ExecuteCommand(rs m.MailChannelSet, command interface{}) {
	//        self.ensure_one()
	//        command_callback = getattr(self, '_execute_command_' + command, False)
	//        if command_callback:
	//            command_callback(**kwargs)
}

//  Notifies partner_to that a message (not stored in DB) has been
//             written in this channel
func mailChannel_SendTransientMessage(rs m.MailChannelSet, partner_to interface{}, content interface{}) {
	//        self.env['bus.bus'].sendone((self._cr.dbname, 'res.partner', partner_to.id), {
	//            'body': "<span class='o_mail_notification'>" + content + "</span>",
	//            'channel_ids': [self.id],
	//            'info': 'transient_message',
	//        })
}

// DefineCommandHelp
func mailChannel_DefineCommandHelp(rs m.MailChannelSet) {
	//        return {'help': _("Show an helper message")}
}

// ExecuteCommandHelp
func mailChannel_ExecuteCommandHelp(rs m.MailChannelSet) {
	//        partner = self.env.user.partner_id
	//        if self.channel_type == 'channel':
	//            msg = _("You are in channel <b>#%s</b>.") % self.name
	//            if self.public == 'private':
	//                msg += _(" This channel is private. People must be invited to join it.")
	//        else:
	//            all_channel_partners = self.env['mail.channel.partner'].with_context(
	//                active_test=False)
	//            channel_partners = all_channel_partners.search(
	//                [('partner_id', '!=', partner.id), ('channel_id', '=', self.id)])
	//            msg = _("You are in a private conversation with <b>@%s</b>.") % (
	//                channel_partners[0].partner_id.name if channel_partners else _('Anonymous'))
	//        msg += _("""<br><br>
	//            Type <b>@username</b> to mention someone, and grab his attention.<br>
	//            Type <b>#channel</b>.to mention a channel.<br>
	//            Type <b>/command</b> to execute a command.<br>
	//            Type <b>:shortcut</b> to insert canned responses in your message.<br>""")
	//        self._send_transient_message(partner, msg)
}

// DefineCommandLeave
func mailChannel_DefineCommandLeave(rs m.MailChannelSet) {
	//        return {'help': _("Leave this channel")}
}

// ExecuteCommandLeave
func mailChannel_ExecuteCommandLeave(rs m.MailChannelSet) {
	//        if self.channel_type == 'channel':
	//            self.action_unfollow()
	//        else:
	//            self.channel_pin(self.uuid, False)
}

// DefineCommandWho
func mailChannel_DefineCommandWho(rs m.MailChannelSet) {
	//        return {
	//            'channel_types': ['channel', 'chat'],
	//            'help': _("List users in the current channel")
	//        }
}

// ExecuteCommandWho
func mailChannel_ExecuteCommandWho(rs m.MailChannelSet) {
	//        partner = self.env.user.partner_id
	//        members = [
	//            '<a href="#" data-oe-id=' +
	//            str(p.id)+' data-oe-model="res.partner">@'+p.name+'</a>'
	//            for p in self.channel_partner_ids[:30] if p != partner
	//        ]
	//        if len(members) == 0:
	//            msg = _("You are alone in this channel.")
	//        else:
	//            dots = "..." if len(members) != len(
	//                self.channel_partner_ids) - 1 else ""
	//            msg = _("Users in this channel: %s %s and you.") % (
	//                ", ".join(members), dots)
	//        self._send_transient_message(partner, msg)
}
func init() {
	models.NewModel("MailChannelPartner")
	h.MailChannelPartner().AddFields(fields_MailChannelPartner)

	models.NewModel("MailModeration")
	h.MailModeration().AddFields(fields_MailModeration)

	models.NewModel("MailChannel")
	h.MailChannel().AddFields(fields_MailChannel)
	h.MailChannel().NewMethod("GetDefaultImage", mailChannel_GetDefaultImage)
	h.MailChannel().Methods().DefaultGet().Extend(mailChannel_DefaultGet)
	h.MailChannel().NewMethod("ComputeIsSubscribed", mailChannel_ComputeIsSubscribed)
	h.MailChannel().NewMethod("ComputeIsModerator", mailChannel_ComputeIsModerator)
	h.MailChannel().NewMethod("ComputeModerationCount", mailChannel_ComputeModerationCount)
	h.MailChannel().NewMethod("CheckModeratorEmail", mailChannel_CheckModeratorEmail)
	h.MailChannel().NewMethod("CheckModeratorIsMember", mailChannel_CheckModeratorIsMember)
	h.MailChannel().NewMethod("CheckModerationParameters", mailChannel_CheckModerationParameters)
	h.MailChannel().NewMethod("CheckModeratorExistence", mailChannel_CheckModeratorExistence)
	h.MailChannel().NewMethod("ComputeIsMember", mailChannel_ComputeIsMember)
	h.MailChannel().NewMethod("ComputeIsChat", mailChannel_ComputeIsChat)
	h.MailChannel().NewMethod("OnchangePublic", mailChannel_OnchangePublic)
	h.MailChannel().NewMethod("OnchangeModeratorIds", mailChannel_OnchangeModeratorIds)
	h.MailChannel().NewMethod("OnchangeEmailSend", mailChannel_OnchangeEmailSend)
	h.MailChannel().NewMethod("OnchangeModeration", mailChannel_OnchangeModeration)
	h.MailChannel().Methods().Create().Extend(mailChannel_Create)
	h.MailChannel().Methods().Unlink().Extend(mailChannel_Unlink)
	h.MailChannel().Methods().Write().Extend(mailChannel_Write)
	h.MailChannel().NewMethod("GetAliasModelName", mailChannel_GetAliasModelName)
	h.MailChannel().NewMethod("SubscribeUsers", mailChannel_SubscribeUsers)
	h.MailChannel().NewMethod("ActionFollow", mailChannel_ActionFollow)
	h.MailChannel().NewMethod("ActionUnfollow", mailChannel_ActionUnfollow)
	h.MailChannel().NewMethod("ActionUnfollow", mailChannel_ActionUnfollow)
	h.MailChannel().NewMethod("NotifyGetGroups", mailChannel_NotifyGetGroups)
	h.MailChannel().NewMethod("NotifySpecificEmailValues", mailChannel_NotifySpecificEmailValues)
	h.MailChannel().NewMethod("MessageReceiveBounce", mailChannel_MessageReceiveBounce)
	h.MailChannel().NewMethod("NotifyEmailRecipients", mailChannel_NotifyEmailRecipients)
	h.MailChannel().NewMethod("ExtractModerationValues", mailChannel_ExtractModerationValues)
	h.MailChannel().NewMethod("MessagePost", mailChannel_MessagePost)
	h.MailChannel().NewMethod("AliasCheckContact", mailChannel_AliasCheckContact)
	h.MailChannel().NewMethod("Init", mailChannel_Init)
	h.MailChannel().NewMethod("SendGuidelines", mailChannel_SendGuidelines)
	h.MailChannel().NewMethod("SendGuidelines", mailChannel_SendGuidelines)
	h.MailChannel().NewMethod("UpdateModerationEmail", mailChannel_UpdateModerationEmail)
	h.MailChannel().NewMethod("Broadcast", mailChannel_Broadcast)
	h.MailChannel().NewMethod("ChannelChannelNotifications", mailChannel_ChannelChannelNotifications)
	h.MailChannel().NewMethod("Notify", mailChannel_Notify)
	h.MailChannel().NewMethod("ChannelMessageNotifications", mailChannel_ChannelMessageNotifications)
	h.MailChannel().NewMethod("ChannelInfo", mailChannel_ChannelInfo)
	h.MailChannel().NewMethod("ChannelFetchMessage", mailChannel_ChannelFetchMessage)
	h.MailChannel().NewMethod("ChannelGet", mailChannel_ChannelGet)
	h.MailChannel().NewMethod("ChannelGetAndMinimize", mailChannel_ChannelGetAndMinimize)
	h.MailChannel().NewMethod("ChannelFold", mailChannel_ChannelFold)
	h.MailChannel().NewMethod("ChannelMinimize", mailChannel_ChannelMinimize)
	h.MailChannel().NewMethod("ChannelPin", mailChannel_ChannelPin)
	h.MailChannel().NewMethod("ChannelSeen", mailChannel_ChannelSeen)
	h.MailChannel().NewMethod("ChannelInvite", mailChannel_ChannelInvite)
	h.MailChannel().NewMethod("NotifyTyping", mailChannel_NotifyTyping)
	h.MailChannel().NewMethod("ChannelFetchSlot", mailChannel_ChannelFetchSlot)
	h.MailChannel().NewMethod("ChannelSearchToJoin", mailChannel_ChannelSearchToJoin)
	h.MailChannel().NewMethod("ChannelJoinAndGetInfo", mailChannel_ChannelJoinAndGetInfo)
	h.MailChannel().NewMethod("ChannelCreate", mailChannel_ChannelCreate)
	h.MailChannel().NewMethod("GetMentionSuggestions", mailChannel_GetMentionSuggestions)
	h.MailChannel().NewMethod("ChannelFetchListeners", mailChannel_ChannelFetchListeners)
	h.MailChannel().NewMethod("ChannelFetchListenersWhereClause", mailChannel_ChannelFetchListenersWhereClause)
	h.MailChannel().NewMethod("ChannelFetchPreview", mailChannel_ChannelFetchPreview)
	h.MailChannel().NewMethod("GetMentionCommands", mailChannel_GetMentionCommands)
	h.MailChannel().NewMethod("ExecuteCommand", mailChannel_ExecuteCommand)
	h.MailChannel().NewMethod("SendTransientMessage", mailChannel_SendTransientMessage)
	h.MailChannel().NewMethod("DefineCommandHelp", mailChannel_DefineCommandHelp)
	h.MailChannel().NewMethod("ExecuteCommandHelp", mailChannel_ExecuteCommandHelp)
	h.MailChannel().NewMethod("DefineCommandLeave", mailChannel_DefineCommandLeave)
	h.MailChannel().NewMethod("ExecuteCommandLeave", mailChannel_ExecuteCommandLeave)
	h.MailChannel().NewMethod("DefineCommandWho", mailChannel_DefineCommandWho)
	h.MailChannel().NewMethod("ExecuteCommandWho", mailChannel_ExecuteCommandWho)

}
