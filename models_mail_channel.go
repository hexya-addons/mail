package mail

import (
	"fmt"

	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

//import re
//import uuid
func init() {
	h.MailChannelPartner().DeclareModel()

	h.MailChannelPartner().AddFields(map[string]models.FieldDefinition{
		"PartnerId": models.Many2OneField{
			RelationModel: h.Partner(),
			String:        "Recipient",
			OnDelete:      `cascade`,
		},
		"ChannelId": models.Many2OneField{
			RelationModel: h.MailChannel(),
			String:        "Channel",
			OnDelete:      `cascade`,
		},
		"SeenMessageId": models.Many2OneField{
			RelationModel: h.MailMessage(),
			String:        "Last Seen",
		},
		"FoldState": models.SelectionField{
			Selection: types.Selection{
				"open":   "Open",
				"folded": "Folded",
				"closed": "Closed",
			},
			String:  "Conversation Fold State",
			Default: models.DefaultValue("open"),
		},
		"IsMinimized": models.BooleanField{
			String: "Conversation is minimized",
		},
		"IsPinned": models.BooleanField{
			String:  "Is pinned on the interface",
			Default: models.DefaultValue(true),
		},
	})
	h.MailChannel().DeclareModel()

	//    _mail_flat_thread = False
	//    _mail_post_access = 'read'

	//    MAX_BOUNCE_LIMIT = 10
	h.MailChannel().Methods().GetDefaultImage().DeclareMethod(
		`GetDefaultImage`,
		func(rs m.MailChannelSet) {
			//        image_path = modules.get_module_resource(
			//            'mail', 'static/src/img', 'groupdefault.png')
			//        return tools.image_resize_image_big(open(image_path, 'rb').read().encode('base64'))
		})
	h.MailChannel().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:    "Name",
			Required:  true,
			Translate: true,
		},
		"ChannelType": models.SelectionField{
			Selection: types.Selection{
				"chat":    "Chat Discussion",
				"channel": "Channel",
			},
			String:  "Channel Type",
			Default: models.DefaultValue("channel"),
		},
		"Description": models.TextField{
			String: "Description",
		},
		"Uuid": models.CharField{
			String:  "UUID",
			Size:    50,
			Index:   true,
			Default: func(env models.Environment) interface{} { return fmt.Sprintf("%s", uuid.uuid4()) },
			NoCopy:  true,
		},
		"EmailSend": models.BooleanField{
			String:  "Send messages by email",
			Default: models.DefaultValue(false),
		},
		"ChannelLastSeenPartnerIds": models.One2ManyField{
			RelationModel: h.MailChannelPartner(),
			ReverseFK:     "",
			String:        "Last Seen",
		},
		"ChannelPartnerIds": models.Many2ManyField{
			RelationModel:    h.Partner(),
			M2MLinkModelName: "",
			M2MOurField:      "",
			M2MTheirField:    "",
			String:           "Listeners",
		},
		"ChannelMessageIds": models.Many2ManyField{
			RelationModel:    h.MailMessage(),
			M2MLinkModelName: "",
		},
		"IsMember": models.BooleanField{
			String:  "Is a member",
			Compute: h.MailChannel().Methods().ComputeIsMember(),
		},
		"Public": models.SelectionField{
			Selection: types.Selection{
				"public":  "Everyone",
				"private": "Invited people only",
				"groups":  "Selected group of users",
			},
			String:   "Privacy",
			Required: true,
			Default:  models.DefaultValue("groups"),
			Help: "This group is visible by non members. Invisible groups" +
				"can add members through the invite button.",
		},
		"GroupPublicId": models.Many2OneField{
			RelationModel: h.Group(),
			String:        "Authorized Group",
			Default:       func(env models.Environment) interface{} { return env.ref() },
		},
		"GroupIds": models.Many2ManyField{
			RelationModel: h.Group(),
			String:        "Auto Subscription",
			Help: "Members of those groups will automatically added as followers." +
				"Note that they will be able to manage their subscription" +
				"manually if necessary.",
		},
		"Image": models.BinaryField{
			String:  "Photo",
			Default: models.DefaultValue(_get_default_image),
			//attachment=True
			Help: "This field holds the image used as photo for the group," +
				"limited to 1024x1024px.",
		},
		"ImageMedium": models.BinaryField{
			String: "Medium-sized photo",
			//attachment=True
			Help: "Medium-sized photo of the group. It is automatically resized" +
				"as a 128x128px image, with aspect ratio preserved. Use" +
				"this field in form views or some kanban views.",
		},
		"ImageSmall": models.BinaryField{
			String: "Small-sized photo",
			//attachment=True
			Help: "Small-sized photo of the group. It is automatically resized" +
				"as a 64x64px image, with aspect ratio preserved. Use this" +
				"field anywhere a small image is required.",
		},
		"AliasId": models.Many2OneField{
			RelationModel: h.MailAlias(),
			String:        "Alias",
			OnDelete:      `restrict`,
			Required:      true,
			Help: "The email address associated with this group. New emails" +
				"received will automatically create new topics.",
		},
		"IsSubscribed": models.BooleanField{
			String:  "Is Subscribed",
			Compute: h.MailChannel().Methods().ComputeIsSubscribed(),
		},
	})
	h.MailChannel().Methods().ComputeIsSubscribed().DeclareMethod(
		`ComputeIsSubscribed`,
		func(rs h.MailChannelSet) h.MailChannelData {
			//        self.is_subscribed = self.env.user.partner_id in self.channel_partner_ids
		})
	h.MailChannel().Methods().ComputeIsMember().DeclareMethod(
		`ComputeIsMember`,
		func(rs h.MailChannelSet) h.MailChannelData {
			//        memberships = self.env['mail.channel.partner'].sudo().search([
			//            ('channel_id', 'in', self.ids),
			//            ('partner_id', '=', self.env.user.partner_id.id),
			//        ])
			//        membership_ids = memberships.mapped('channel_id')
			//        for record in self:
			//            record.is_member = record in membership_ids
		})
	h.MailChannel().Methods().Create().Extend(
		`Create`,
		func(rs m.MailChannelSet, vals models.RecordData) {
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
		})
	h.MailChannel().Methods().Unlink().Extend(
		`Unlink`,
		func(rs m.MailChannelSet) {
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
		})
	h.MailChannel().Methods().Write().Extend(
		`Write`,
		func(rs m.MailChannelSet, vals models.RecordData) {
			//        tools.image_resize_images(vals)
			//        result = super(Channel, self).write(vals)
			//        if vals.get('group_ids'):
			//            self._subscribe_users()
			//        return result
		})
	h.MailChannel().Methods().SubscribeUsers().DeclareMethod(
		`SubscribeUsers`,
		func(rs m.MailChannelSet) {
			//        for mail_channel in self:
			//            mail_channel.write({'channel_partner_ids': [(4, pid) for pid in mail_channel.mapped(
			//                'group_ids').mapped('users').mapped('partner_id').ids]})
		})
	h.MailChannel().Methods().ActionFollow().DeclareMethod(
		`ActionFollow`,
		func(rs m.MailChannelSet) {
			//        self.ensure_one()
			//        channel_partner = self.mapped('channel_last_seen_partner_ids').filtered(
			//            lambda cp: cp.partner_id == self.env.user.partner_id)
			//        if not channel_partner:
			//            return self.write({'channel_last_seen_partner_ids': [(0, 0, {'partner_id': self.env.user.partner_id.id})]})
		})
	h.MailChannel().Methods().ActionUnfollow().DeclareMethod(
		`ActionUnfollow`,
		func(rs m.MailChannelSet) {
			//        return self._action_unfollow(self.env.user.partner_id)
		})
	h.MailChannel().Methods().ActionUnfollow().DeclareMethod(
		`ActionUnfollow`,
		func(rs m.MailChannelSet, partner interface{}) {
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
		})
	h.MailChannel().Methods().NotificationRecipients().DeclareMethod(
		` All recipients of a message on a channel are considered as partners.
        This means they will receive a minimal email, without
a link to access
        in the backend. Mailing lists should indeed send
minimal emails to avoid
        the noise. `,
		func(rs m.MailChannelSet, message interface{}, groups interface{}) {
			//        groups = super(Channel, self)._notification_recipients(message, groups)
			//        for (index, (group_name, group_func, group_data)) in enumerate(groups):
			//            if group_name != 'customer':
			//                groups[index] = (group_name, lambda partner: False, group_data)
			//        return groups
		})
	h.MailChannel().Methods().MessageGetEmailValues().DeclareMethod(
		`MessageGetEmailValues`,
		func(rs m.MailChannelSet, notif_mail interface{}) {
			//        self.ensure_one()
			//        res = super(Channel, self).message_get_email_values(
			//            notif_mail=notif_mail)
			//        headers = {}
			//        if res.get('headers'):
			//            try:
			//                headers.update(safe_eval(res['headers']))
			//            except Exception:
			//                pass
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
		})
	h.MailChannel().Methods().MessageReceiveBounce().DeclareMethod(
		` Override bounce management to unsubscribe bouncing addresses `,
		func(rs m.MailChannelSet, email interface{}, partner interface{}, mail_id interface{}) {
			//        for p in partner:
			//            if p.message_bounce >= self.MAX_BOUNCE_LIMIT:
			//                self._action_unfollow(p)
			//        return super(Channel, self).message_receive_bounce(email, partner, mail_id=mail_id)
		})
	h.MailChannel().Methods().MessageGetRecipientValues().DeclareMethod(
		`MessageGetRecipientValues`,
		func(rs m.MailChannelSet, notif_message interface{}, recipient_ids interface{}) {
			//        if self.alias_domain and self.alias_name:
			//            return {
			//                'email_to': ','.join(formataddr((partner.name, partner.email)) for partner in self.env['res.partner'].sudo().browse(recipient_ids)),
			//                'recipient_ids': [],
			//            }
			//        return super(Channel, self).message_get_recipient_values(notif_message=notif_message, recipient_ids=recipient_ids)
		})
	h.MailChannel().Methods().MessagePost().DeclareMethod(
		`MessagePost`,
		func(rs m.MailChannelSet, body interface{}, subject interface{}, message_type interface{}, subtype interface{}, parent_id interface{}, attachments interface{}, content_subtype interface{}) {
			//        self.filtered(lambda channel: channel.channel_type == 'chat').mapped(
			//            'channel_last_seen_partner_ids').write({'is_pinned': True})
			//        message = super(Channel, self.with_context(mail_create_nosubscribe=True)).message_post(body=body, subject=subject,
			//                                                                                               message_type=message_type, subtype=subtype, parent_id=parent_id, attachments=attachments, content_subtype=content_subtype, **kwargs)
			//        return message
		})
	h.MailChannel().Methods().Init().DeclareMethod(
		`Init`,
		func(rs m.MailChannelSet) {
			//        self._cr.execute('SELECT indexname FROM pg_indexes WHERE indexname = %s',
			//                         ('mail_channel_partner_seen_message_id_idx'))
			//        if not self._cr.fetchone():
			//            self._cr.execute(
			//                'CREATE INDEX mail_channel_partner_seen_message_id_idx ON mail_channel_partner (channel_id,partner_id,seen_message_id)')
		})
	h.MailChannel().Methods().Broadcast().DeclareMethod(
		` Broadcast the current channel header to the given partner ids
            :param partner_ids : the partner to notify
        `,
		func(rs m.MailChannelSet, partner_ids interface{}) {
			//        notifications = self._channel_channel_notifications(partner_ids)
			//        self.env['bus.bus'].sendmany(notifications)
		})
	h.MailChannel().Methods().ChannelChannelNotifications().DeclareMethod(
		` Generate the bus notifications of current channel for
the given partner ids
            :param partner_ids : the partner to send the
current channel header
            :returns list of bus notifications (tuple (bus_channe,
message_content))
        `,
		func(rs m.MailChannelSet, partner_ids interface{}) {
			//        notifications = []
			//        for partner in self.env['res.partner'].browse(partner_ids):
			//            user_id = partner.user_ids and partner.user_ids[0] or False
			//            if user_id:
			//                for channel_info in self.sudo(user_id).channel_info():
			//                    notifications.append(
			//                        [(self._cr.dbname, 'res.partner', partner.id), channel_info])
			//        return notifications
		})
	h.MailChannel().Methods().Notify().DeclareMethod(
		` Broadcast the given message on the current channels.
            Send the message on the Bus Channel (uuid for
public mail.channel, and partner private bus channel (the tuple)).
            A partner will receive only on message on its
bus channel, even if this message belongs to multiple mail
channel. Then 'channel_ids' field
            of the received message indicates on wich mail
channel the message should be displayed.
            :param : mail.message to broadcast
        `,
		func(rs m.MailChannelSet, message interface{}) {
			//        message.ensure_one()
			//        notifications = self._channel_message_notifications(message)
			//        self.env['bus.bus'].sendmany(notifications)
		})
	h.MailChannel().Methods().ChannelMessageNotifications().DeclareMethod(
		` Generate the bus notifications for the given message
            :param message : the mail.message to sent
            :returns list of bus notifications (tuple (bus_channe,
message_content))
        `,
		func(rs m.MailChannelSet, message interface{}) {
			//        message_values = message.message_format()[0]
			//        notifications = []
			//        for channel in self:
			//            notifications.append(
			//                [(self._cr.dbname, 'mail.channel', channel.id), dict(message_values)])
			//            # add uuid to allow anonymous to listen
			//            if channel.public == 'public':
			//                notifications.append([channel.uuid, dict(message_values)])
			//        return notifications
		})
	h.MailChannel().Methods().ChannelInfo().DeclareMethod(
		` Get the informations header for the current channels
            :returns a list of channels values
            :rtype : list(dict)
        `,
		func(rs m.MailChannelSet, extra_info interface{}) {
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
		})
	h.MailChannel().Methods().ChannelFetchMessage().DeclareMethod(
		` Return message values of the current channel.
            :param last_id : last message id to start the research
            :param limit : maximum number of messages to fetch
            :returns list of messages values
            :rtype : list(dict)
        `,
		func(rs m.MailChannelSet, last_id interface{}, limit interface{}) {
			//        self.ensure_one()
			//        domain = [("channel_ids", "in", self.ids)]
			//        if last_id:
			//            domain.append(("id", "<", last_id))
			//        return self.env['mail.message'].message_fetch(domain=domain, limit=limit)
		})
	h.MailChannel().Methods().ChannelGet().DeclareMethod(
		` Get the canonical private channel between some partners,
create it if needed.
            To reuse an old channel (conversation), this
one must be private, and contains
            only the given partners.
            :param partners_to : list of res.partner ids
to add to the conversation
            :param pin : True if getting the channel should
pin it for the current user
            :returns a channel header, or False if the
users_to was False
            :rtype : dict
        `,
		func(rs m.MailChannelSet, partners_to interface{}, pin interface{}) {
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
		})
	h.MailChannel().Methods().ChannelGetAndMinimize().DeclareMethod(
		`ChannelGetAndMinimize`,
		func(rs m.MailChannelSet, partners_to interface{}) {
			//        channel = self.channel_get(partners_to)
			//        if channel:
			//            self.channel_minimize(channel['uuid'])
			//        return channel
		})
	h.MailChannel().Methods().ChannelFold().DeclareMethod(
		` Update the fold_state of the given session. In order to
syncronize web browser
            tabs, the change will be broadcast to himself
(the current user channel).
            Note: the user need to be logged
            :param state : the new status of the session
for the current user.
        `,
		func(rs m.MailChannelSet, uuid interface{}, state interface{}) {
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
		})
	h.MailChannel().Methods().ChannelMinimize().DeclareMethod(
		`ChannelMinimize`,
		func(rs m.MailChannelSet, uuid interface{}, minimized interface{}) {
			//        values = {
			//            'fold_state': minimized and 'open' or 'closed',
			//            'is_minimized': minimized
			//        }
			//        domain = [('partner_id', '=', self.env.user.partner_id.id),
			//                  ('channel_id.uuid', '=', uuid)]
			//        channel_partners = self.env['mail.channel.partner'].search(domain)
			//        channel_partners.write(values)
			//        self.env['bus.bus'].sendone(
			//            (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), channel_partners.channel_id.channel_info()[0])
		})
	h.MailChannel().Methods().ChannelPin().DeclareMethod(
		`ChannelPin`,
		func(rs m.MailChannelSet, uuid interface{}, pinned interface{}) {
			//        channel = self.search([('uuid', '=', uuid)])
			//        channel_partners = self.env['mail.channel.partner'].search(
			//            [('partner_id', '=', self.env.user.partner_id.id), ('channel_id', '=', channel.id)])
			//        if not pinned:
			//            self.env['bus.bus'].sendone(
			//                (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), channel.channel_info('unsubscribe')[0])
			//        if channel_partners:
			//            channel_partners.write({'is_pinned': pinned})
		})
	h.MailChannel().Methods().ChannelSeen().DeclareMethod(
		`ChannelSeen`,
		func(rs m.MailChannelSet) {
			//        self.ensure_one()
			//        if self.channel_message_ids.ids:
			//            # zero is the index of the last message
			//            last_message_id = self.channel_message_ids.ids[0]
			//            self.env['mail.channel.partner'].search([('channel_id', 'in', self.ids), (
			//                'partner_id', '=', self.env.user.partner_id.id)]).write({'seen_message_id': last_message_id})
			//            self.env['bus.bus'].sendone((self._cr.dbname, 'res.partner', self.env.user.partner_id.id), {
			//                                        'info': 'channel_seen', 'id': self.id, 'last_message_id': last_message_id})
			//            return last_message_id
		})
	h.MailChannel().Methods().ChannelInvite().DeclareMethod(
		` Add the given partner_ids to the current channels and
broadcast the channel header to them.
            :param partner_ids : list of partner id to add
        `,
		func(rs m.MailChannelSet, partner_ids interface{}) {
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
		})
	h.MailChannel().Methods().ChannelFetchSlot().DeclareMethod(
		` Return the channels of the user grouped by 'slot' (channel,
direct_message or private_group), and
            the mapping between partner_id/channel_id for
direct_message channels.
            :returns dict : the grouped channels and the mapping
        `,
		func(rs m.MailChannelSet) {
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
		})
	h.MailChannel().Methods().ChannelSearchToJoin().DeclareMethod(
		` Return the channel info of the channel the current partner can join
            :param name : the name of the researched channels
            :param domain : the base domain of the research
            :returns dict : channel dict
        `,
		func(rs m.MailChannelSet, name interface{}, domain interface{}) {
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
		})
	h.MailChannel().Methods().ChannelJoinAndGetInfo().DeclareMethod(
		`ChannelJoinAndGetInfo`,
		func(rs m.MailChannelSet) {
			//        self.ensure_one()
			//        if self.channel_type == 'channel' and not self.email_send:
			//            notification = _(
			//                '<div class="o_mail_notification">joined <a href="#" class="o_channel_redirect" data-oe-id="%s">#%s</a></div>') % (self.id, self.name)
			//            self.message_post(
			//                body=notification, message_type="notification", subtype="mail.mt_comment")
			//        self.action_follow()
			//        channel_info = self.channel_info()[0]
			//        self.env['bus.bus'].sendone(
			//            (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), channel_info)
			//        return channel_info
		})
	h.MailChannel().Methods().ChannelCreate().DeclareMethod(
		` Create a channel and add the current partner, broadcast
it (to make the user directly
            listen to it when polling)
            :param name : the name of the channel to create
            :param privacy : privacy of the channel. Should
be 'public' or 'private'.
            :return dict : channel header
        `,
		func(rs m.MailChannelSet, name interface{}, privacy interface{}) {
			//        new_channel = self.create({
			//            'name': name,
			//            'public': privacy,
			//            'email_send': False,
			//            'channel_partner_ids': [(4, self.env.user.partner_id.id)]
			//        })
			//        channel_info = new_channel.channel_info('creation')[0]
			//        notification = _('<div class="o_mail_notification">created <a href="#" class="o_channel_redirect" data-oe-id="%s">#%s</a></div>') % (
			//            new_channel.id, new_channel.name)
			//        new_channel.message_post(
			//            body=notification, message_type="notification", subtype="mail.mt_comment")
			//        self.env['bus.bus'].sendone(
			//            (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), channel_info)
			//        return channel_info
		})
	h.MailChannel().Methods().GetMentionSuggestions().DeclareMethod(
		` Return 'limit'-first channels' id, name and public fields
such that the name matches a
            'search' string. Exclude channels of type chat
(DM), and private channels the current
            user isn't registered to. `,
		func(rs m.MailChannelSet, search interface{}, limit interface{}) {
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
		})
	h.MailChannel().Methods().ChannelFetchListeners().DeclareMethod(
		` Return the id, name and email of partners listening to
the given channel `,
		func(rs m.MailChannelSet, uuid interface{}) {
			//        self._cr.execute("""
			//            SELECT P.id, P.name, P.email
			//            FROM mail_channel_partner CP
			//                INNER JOIN res_partner P ON CP.partner_id = P.id
			//                INNER JOIN mail_channel C ON CP.channel_id = C.id
			//            WHERE C.uuid = %s""", (uuid))
			//        return self._cr.dictfetchall()
		})
	h.MailChannel().Methods().ChannelFetchPreview().DeclareMethod(
		` Return the last message of the given channels `,
		func(rs m.MailChannelSet) {
			//        self._cr.execute("""
			//            SELECT mail_channel_id AS id, MAX(mail_message_id) AS message_id
			//            FROM mail_message_mail_channel_rel
			//            WHERE mail_channel_id IN %s
			//            GROUP BY mail_channel_id
			//            """, (tuple(self.ids)))
			//        channels_preview = dict((r['message_id'], r)
			//                                for r in self._cr.dictfetchall())
			//        last_messages = self.env['mail.message'].browse(
			//            channels_preview.keys()).message_format()
			//        for message in last_messages:
			//            channel = channels_preview[message['id']]
			//            del(channel['message_id'])
			//            channel['last_message'] = message
			//        return channels_preview.values()
		})
	h.MailChannel().Methods().GetMentionCommands().DeclareMethod(
		` Returns the allowed commands in channels `,
		func(rs m.MailChannelSet) {
			//        commands = []
			//        for n in dir(self):
			//            match = re.search('^_define_command_(.+?)$', n)
			//            if match:
			//                command = getattr(self, n)()
			//                command['name'] = match.group(1)
			//                commands.append(command)
			//        return commands
		})
	h.MailChannel().Methods().ExecuteCommand().DeclareMethod(
		` Executes a given command `,
		func(rs m.MailChannelSet, command interface{}) {
			//        self.ensure_one()
			//        command_callback = getattr(self, '_execute_command_' + command, False)
			//        if command_callback:
			//            command_callback(**kwargs)
		})
	h.MailChannel().Methods().SendTransientMessage().DeclareMethod(
		` Notifies partner_to that a message (not stored in DB) has been
            written in this channel `,
		func(rs m.MailChannelSet, partner_to interface{}, content interface{}) {
			//        self.env['bus.bus'].sendone((self._cr.dbname, 'res.partner', partner_to.id), {
			//            'body': "<span class='o_mail_notification'>" + content + "</span>",
			//            'channel_ids': [self.id],
			//            'info': 'transient_message',
			//        })
		})
	h.MailChannel().Methods().DefineCommandHelp().DeclareMethod(
		`DefineCommandHelp`,
		func(rs m.MailChannelSet) {
			//        return {'help': _("Show an helper message")}
		})
	h.MailChannel().Methods().ExecuteCommandHelp().DeclareMethod(
		`ExecuteCommandHelp`,
		func(rs m.MailChannelSet) {
			//        partner = self.env.user.partner_id
			//        if self.channel_type == 'channel':
			//            msg = _("You are in channel <b>#%s</b>.") % self.name
			//            if self.public == 'private':
			//                msg += _(" This channel is private. People must be invited to join it.")
			//        else:
			//            channel_partners = self.env['mail.channel.partner'].search(
			//                [('partner_id', '!=', partner.id), ('channel_id', '=', self.id)])
			//            msg = _("You are in a private conversation with <b>@%s</b>.") % (
			//                channel_partners[0].partner_id.name if channel_partners else _('Anonymous'))
			//        msg += _("""<br><br>
			//            You can mention someone by typing <b>@username</b>, this will grab its attention.<br>
			//            You can mention a channel by typing <b>#channel</b>.<br>
			//            You can execute a command by typing <b>/command</b>.<br>
			//            You can insert canned responses in your message by typing <b>:shortcut</b>.<br>""")
			//        self._send_transient_message(partner, msg)
		})
	h.MailChannel().Methods().DefineCommandLeave().DeclareMethod(
		`DefineCommandLeave`,
		func(rs m.MailChannelSet) {
			//        return {'help': _("Leave this channel")}
		})
	h.MailChannel().Methods().ExecuteCommandLeave().DeclareMethod(
		`ExecuteCommandLeave`,
		func(rs m.MailChannelSet) {
			//        if self.channel_type == 'channel':
			//            self.action_unfollow()
			//        else:
			//            self.channel_pin(self.uuid, False)
		})
	h.MailChannel().Methods().DefineCommandWho().DeclareMethod(
		`DefineCommandWho`,
		func(rs m.MailChannelSet) {
			//        return {
			//            'channel_types': ['channel', 'chat'],
			//            'help': _("List users in the current channel")
			//        }
		})
	h.MailChannel().Methods().ExecuteCommandWho().DeclareMethod(
		`ExecuteCommandWho`,
		func(rs m.MailChannelSet) {
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
		})
}
