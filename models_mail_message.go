package mail

	import (
		"net/http"

		"github.com/hexya-erp/hexya/src/controllers"
		"github.com/hexya-erp/hexya/src/models"
		"github.com/hexya-erp/hexya/src/models/types"
		"github.com/hexya-erp/hexya/src/models/types/dates"
		"github.com/hexya-erp/pool/h"
		"github.com/hexya-erp/pool/q"
	)
	
//import logging
//_logger = logging.getLogger(__name__)
func init() {
h.MailMessage().DeclareModel()





//    _message_read_limit = 30
h.MailMessage().Methods().GetDefaultFrom().DeclareMethod(
`GetDefaultFrom`,
func(rs m.MailMessageSet)  {
//        if self.env.user.email:
//            return formataddr((self.env.user.name, self.env.user.email))
//        raise UserError(
//            _("Unable to send email, please configure the sender's email address."))
})
h.MailMessage().Methods().GetDefaultAuthor().DeclareMethod(
`GetDefaultAuthor`,
func(rs m.MailMessageSet)  {
//        return self.env.user.partner_id
})
h.MailMessage().AddFields(map[string]models.FieldDefinition{
"Subject": models.CharField{
String: "Subject",
},
"Date": models.DateTimeField{
String: "Date",
Default: func (env models.Environment) interface{} { return dates.Now() },
},
"Body": models.HTMLField{
String: "Contents",
Default: models.DefaultValue(""),
//sanitize_style=True
//strip_classes=True
},
"AttachmentIds": models.Many2ManyField{
RelationModel: h.Attachment(),
M2MLinkModelName: "",
M2MOurField: "",
M2MTheirField: "",
String: "Attachments",
Help: "Attachments are linked to a document through model / res_id" + 
"and to the message through this field.",
},
"ParentId": models.Many2OneField{
RelationModel: h.MailMessage(),
String: "Parent Message",
Index: true,
OnDelete: `set null`,
Help: "Initial thread message.",
},
"ChildIds": models.One2ManyField{
RelationModel: h.MailMessage(),
ReverseFK: "",
String: "Child Messages",
},
"Model": models.CharField{
String: "Related Document Model",
Index: true,
},
"ResId": models.IntegerField{
String: "Related Document ID",
Index: true,
},
"RecordName": models.CharField{
String: "Message Record Name",
Help: "Name get of the related document.",
},
"MessageType": models.SelectionField{
Selection: types.Selection{
"email": "Email",
"comment": "Comment",
"notification": "System notification",
},
String: "Type",
Required: true,
Default: models.DefaultValue("email"),
Help: "Message type: email for email message, notification for" + 
"system message, comment for other messages such as user replies",
//oldname='type'
},
"SubtypeId": models.Many2OneField{
RelationModel: h.MailMessageSubtype(),
String: "Subtype",
OnDelete: `set null`,
Index: true,
},
"EmailFrom": models.CharField{
String: "From",
Default: models.DefaultValue(_get_default_from),
Help: "Email address of the sender. This field is set when no" + 
"matching partner is found and replaces the author_id field" + 
"in the chatter.",
},
"AuthorId": models.Many2OneField{
RelationModel: h.Partner(),
String: "Author",
Index: true,
OnDelete: `set null`,
Default: models.DefaultValue(_get_default_author),
Help: "Author of the message. If not set, email_from may hold" + 
"an email address that did not match any partner.",
},
"AuthorAvatar": models.BinaryField{
String: "Author's avatar",
Related: `AuthorId.ImageSmall`,
},
"PartnerIds": models.Many2ManyField{
RelationModel: h.Partner(),
String: "Recipients",
},
"NeedactionPartnerIds": models.Many2ManyField{
RelationModel: h.Partner(),
M2MLinkModelName: "",
String: "Partners with Need Action",
},
"Needaction": models.BooleanField{
String: "Need Action",
Compute: h.MailMessage().Methods().GetNeedaction(),
//search='_search_needaction'
Help: "Need Action",
},
"ChannelIds": models.Many2ManyField{
RelationModel: h.MailChannel(),
M2MLinkModelName: "",
String: "Channels",
},
"NotificationIds": models.One2ManyField{
RelationModel: h.MailNotification(),
ReverseFK: "",
String: "Notifications",

NoCopy: true,
},
"StarredPartnerIds": models.Many2ManyField{
RelationModel: h.Partner(),
M2MLinkModelName: "",
String: "Favorited By",
},
"Starred": models.BooleanField{
String: "Starred",
Compute: h.MailMessage().Methods().GetStarred(),
//search='_search_starred'
Help: "Current user has a starred notification linked to this message",
},
"TrackingValueIds": models.One2ManyField{
RelationModel: h.MailTrackingValue(),
ReverseFK: "",
String: "Tracking values",
//groups="base.group_no_one"
Help: "Tracked values are stored in a separate model. This field" + 
"allow to reconstruct the tracking and to generate statistics" + 
"on the model.",
},
"NoAutoThread": models.BooleanField{
String: "No threading for answers",
Help: "Answers do not go in the original document discussion thread." + 
"This has an impact on the generated message-id.",
},
"MessageId": models.CharField{
String: "Message-Id",
Help: "Message unique identifier",
Index: true,
ReadOnly: true,
NoCopy: true,
},
"ReplyTo": models.CharField{
String: "Reply-To",
Help: "Reply email address. Setting the reply_to bypasses the" + 
"automatic thread creation.",
},
"MailServerId": models.Many2OneField{
RelationModel: h.IrMail_server(),
String: "Outgoing mail server",
},
})
h.MailMessage().Methods().GetNeedaction().DeclareMethod(
` Need action on a mail.message = notified on my channel `,
func(rs h.MailMessageSet) h.MailMessageData {
//        my_messages = self.env['mail.notification'].sudo().search([
//            ('mail_message_id', 'in', self.ids),
//            ('res_partner_id', '=', self.env.user.partner_id.id),
//            ('is_read', '=', False)]).mapped('mail_message_id')
//        for message in self:
//            message.needaction = message in my_messages
})
h.MailMessage().Methods().SearchNeedaction().DeclareMethod(
`SearchNeedaction`,
func(rs m.MailMessageSet, operator interface{}, operand interface{})  {
//        if operator == '=' and operand:
//            return ['&', ('notification_ids.res_partner_id', '=', self.env.user.partner_id.id), ('notification_ids.is_read', '=', False)]
//        return ['&', ('notification_ids.res_partner_id', '=', self.env.user.partner_id.id), ('notification_ids.is_read', '=', True)]
})
h.MailMessage().Methods().GetStarred().DeclareMethod(
` Compute if the message is starred by the current user. `,
func(rs h.MailMessageSet) h.MailMessageData {
//        starred = self.sudo().filtered(
//            lambda msg: self.env.user.partner_id in msg.starred_partner_ids)
//        for message in self:
//            message.starred = message in starred
})
h.MailMessage().Methods().SearchStarred().DeclareMethod(
`SearchStarred`,
func(rs m.MailMessageSet, operator interface{}, operand interface{})  {
//        if operator == '=' and operand:
//            return [('starred_partner_ids', 'in', [self.env.user.partner_id.id])]
//        return [('starred_partner_ids', 'not in', [self.env.user.partner_id.id])]
})
h.MailMessage().Methods().NeedactionDomainGet().DeclareMethod(
`NeedactionDomainGet`,
func(rs m.MailMessageSet)  {
//        return [('needaction', '=', True)]
})
h.MailMessage().Methods().MarkAllAsRead().DeclareMethod(
` Remove all needactions of the current partner. If channel_ids is
            given, restrict to messages written in one
of those channels. `,
func(rs m.MailMessageSet, channel_ids interface{}, domain interface{})  {
//        partner_id = self.env.user.partner_id.id
//        delete_mode = not self.env.user.share
//        if not domain and delete_mode:
//            query = "DELETE FROM mail_message_res_partner_needaction_rel WHERE res_partner_id IN %s"
//            args = [(partner_id)]
//            if channel_ids:
//                query += """
//                    AND mail_message_id in
//                        (SELECT mail_message_id
//                        FROM mail_message_mail_channel_rel
//                        WHERE mail_channel_id in %s)"""
//                args += [tuple(channel_ids)]
//            query += " RETURNING mail_message_id as id"
//            self._cr.execute(query, args)
//            self.invalidate_cache()
//
//            ids = [m['id'] for m in self._cr.dictfetchall()]
//        else:
//            # not really efficient method: it does one db request for the
//            # search, and one for each message in the result set to remove the
//            # current user from the relation.
//            msg_domain = [('needaction_partner_ids', 'in', partner_id)]
//            if channel_ids:
//                msg_domain += [('channel_ids', 'in', channel_ids)]
//            unread_messages = self.search(expression.AND([msg_domain, domain]))
//            notifications = self.env['mail.notification'].sudo().search([
//                ('mail_message_id', 'in', unread_messages.ids),
//                ('res_partner_id', '=', self.env.user.partner_id.id),
//                ('is_read', '=', False)])
//            if delete_mode:
//                notifications.unlink()
//            else:
//                notifications.write({'is_read': True})
//            ids = unread_messages.mapped('id')
//        notification = {'type': 'mark_as_read',
//                        'message_ids': ids, 'channel_ids': channel_ids}
//        self.env['bus.bus'].sendone(
//            (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), notification)
//        return ids
})
h.MailMessage().Methods().MarkAsUnread().DeclareMethod(
` Add needactions to messages for the current partner. `,
func(rs m.MailMessageSet, channel_ids interface{})  {
//        partner_id = self.env.user.partner_id.id
//        for message in self:
//            message.write({'needaction_partner_ids': [(4, partner_id)]})
//        ids = [m.id for m in self]
//        notification = {'type': 'mark_as_unread',
//                        'message_ids': ids, 'channel_ids': channel_ids}
//        self.env['bus.bus'].sendone(
//            (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), notification)
})
h.MailMessage().Methods().SetMessageDone().DeclareMethod(
` Remove the needaction from messages for the current partner. `,
func(rs m.MailMessageSet)  {
//        partner_id = self.env.user.partner_id
//        delete_mode = not self.env.user.share
//        notifications = self.env['mail.notification'].sudo().search([
//            ('mail_message_id', 'in', self.ids),
//            ('res_partner_id', '=', partner_id.id),
//            ('is_read', '=', False)])
//        if not notifications:
//            return
//        groups = []
//        messages = notifications.mapped('mail_message_id')
//        current_channel_ids = messages[0].channel_ids
//        current_group = []
//        for record in messages:
//            if record.channel_ids == current_channel_ids:
//                current_group.append(record.id)
//            else:
//                groups.append((current_group, current_channel_ids))
//                current_group = [record.id]
//                current_channel_ids = record.channel_ids
//        groups.append((current_group, current_channel_ids))
//        current_group = [record.id]
//        current_channel_ids = record.channel_ids
//        if delete_mode:
//            notifications.unlink()
//        else:
//            notifications.write({'is_read': True})
//        for (msg_ids, channel_ids) in groups:
//            notification = {'type': 'mark_as_read', 'message_ids': msg_ids, 'channel_ids': [
//                c.id for c in channel_ids]}
//            self.env['bus.bus'].sendone(
//                (self._cr.dbname, 'res.partner', partner_id.id), notification)
})
h.MailMessage().Methods().UnstarAll().DeclareMethod(
` Unstar messages for the current partner. `,
func(rs m.MailMessageSet)  {
//        partner_id = self.env.user.partner_id.id
//        starred_messages = self.search(
//            [('starred_partner_ids', 'in', partner_id)])
//        starred_messages.write({'starred_partner_ids': [(3, partner_id)]})
//        ids = [m.id for m in starred_messages]
//        notification = {'type': 'toggle_star',
//                        'message_ids': ids, 'starred': False}
//        self.env['bus.bus'].sendone(
//            (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), notification)
})
h.MailMessage().Methods().ToggleMessageStarred().DeclareMethod(
` Toggle messages as (un)starred. Technically, the notifications related
            to uid are set to (un)starred.
        `,
func(rs m.MailMessageSet)  {
//        self.check_access_rule('read')
//        starred = not self.starred
//        if starred:
//            self.sudo().write({'starred_partner_ids': [
//                (4, self.env.user.partner_id.id)]})
//        else:
//            self.sudo().write({'starred_partner_ids': [
//                (3, self.env.user.partner_id.id)]})
//        notification = {'type': 'toggle_star',
//                        'message_ids': [self.id], 'starred': starred}
//        self.env['bus.bus'].sendone(
//            (self._cr.dbname, 'res.partner', self.env.user.partner_id.id), notification)
})
h.MailMessage().Methods().MessageReadDictPostprocess().DeclareMethod(
` Post-processing on values given by message_read. This method will
            handle partners in batch to avoid doing numerous queries.

            :param list messages: list of message, as get_dict result
            :param dict message_tree: {[msg.id]: msg browse
record as super user}
        `,
func(rs m.MailMessageSet, messages interface{}, message_tree interface{})  {
//        partners = self.env['res.partner'].sudo()
//        attachments = self.env['ir.attachment']
//        message_ids = message_tree.keys()
//        for key, message in message_tree.iteritems():
//            if message.author_id:
//                partners |= message.author_id
//            if message.subtype_id and message.partner_ids:  # take notified people of message with a subtype
//                partners |= message.partner_ids
//            # take specified people of message without a subtype (log)
//            elif not message.subtype_id and message.partner_ids:
//                partners |= message.partner_ids
//            if message.needaction_partner_ids:  # notified
//                partners |= message.needaction_partner_ids
//            if message.attachment_ids:
//                attachments |= message.attachment_ids
//        partners_names = partners.name_get()
//        partner_tree = dict((partner[0], partner)
//                            for partner in partners_names)
//        attachments_data = attachments.sudo().read(
//            ['id', 'datas_fname', 'name', 'mimetype'])
//        attachments_tree = dict((attachment['id'], {
//            'id': attachment['id'],
//            'filename': attachment['datas_fname'],
//            'name': attachment['name'],
//            'mimetype': attachment['mimetype'],
//        }) for attachment in attachments_data)
//        tracking_values = self.env['mail.tracking.value'].sudo().search(
//            [('mail_message_id', 'in', message_ids)])
//        message_to_tracking = dict()
//        tracking_tree = dict.fromkeys(tracking_values.ids, False)
//        for tracking in tracking_values:
//            message_to_tracking.setdefault(
//                tracking.mail_message_id.id, list()).append(tracking.id)
//            tracking_tree[tracking.id] = {
//                'id': tracking.id,
//                'changed_field': tracking.field_desc,
//                'old_value': tracking.get_old_display_value()[0],
//                'new_value': tracking.get_new_display_value()[0],
//                'field_type': tracking.field_type,
//            }
//        for message_dict in messages:
//            message_id = message_dict.get('id')
//            message = message_tree[message_id]
//            if message.author_id:
//                author = partner_tree[message.author_id.id]
//            else:
//                author = (0, message.email_from)
//            partner_ids = []
//            if message.subtype_id:
//                partner_ids = [partner_tree[partner.id] for partner in message.partner_ids
//                               if partner.id in partner_tree]
//            else:
//                partner_ids = [partner_tree[partner.id] for partner in message.partner_ids
//                               if partner.id in partner_tree]
//
//            customer_email_data = []
//            for notification in message.notification_ids.filtered(lambda notif: notif.res_partner_id.partner_share and notif.res_partner_id.active):
//                customer_email_data.append((partner_tree[notification.res_partner_id.id][0],
//                                            partner_tree[notification.res_partner_id.id][1], notification.email_status))
//
//            attachment_ids = []
//            for attachment in message.attachment_ids:
//                if attachment.id in attachments_tree:
//                    attachment_ids.append(attachments_tree[attachment.id])
//            tracking_value_ids = []
//            for tracking_value_id in message_to_tracking.get(message_id, list()):
//                if tracking_value_id in tracking_tree:
//                    tracking_value_ids.append(tracking_tree[tracking_value_id])
//
//            message_dict.update({
//                'author_id': author,
//                'partner_ids': partner_ids,
//                'customer_email_status': (all(d[2] == 'sent' for d in customer_email_data) and 'sent') or
//                (any(d[2] == 'exception' for d in customer_email_data) and 'exception') or
//                (any(
//                    d[2] == 'bounce' for d in customer_email_data) and 'bounce') or 'ready',
//                'customer_email_data': customer_email_data,
//                'attachment_ids': attachment_ids,
//                'tracking_value_ids': tracking_value_ids,
//            })
//        return True
})
h.MailMessage().Methods().MessageFetch().DeclareMethod(
`MessageFetch`,
func(rs m.MailMessageSet, domain interface{}, limit interface{})  {
//        return self.search(domain, limit=limit).message_format()
})
h.MailMessage().Methods().MessageFormat().DeclareMethod(
` Get the message values in the format for web client. Since
message values can be broadcasted,
            computed fields MUST NOT BE READ and broadcasted.
            :returns list(dict).
             Example :
                {
                    'body': HTML content of the message
                    'model': u'res.partner',
                    'record_name': u'Agrolait',
                    'attachment_ids': [
                        {
                            'file_type_icon': u'webimage',
                            'id': 45,
                            'name': u'sample.png',
                            'filename': u'sample.png'
                        }
                    ],
                    'needaction_partner_ids': [], # list of partner ids
                    'res_id': 7,
                    'tracking_value_ids': [
                        {
                            'old_value': "",
                            'changed_field': "Customer",
                            'id': 2965,
                            'new_value': "Axelor"
                        }
                    ],
                    'author_id': (3, u'Administrator'),
                    'email_from': 'sacha@pokemon.com' #
email address or False
                    'subtype_id': (1, u'Discussions'),
                    'channel_ids': [], # list of channel ids
                    'date': '2015-06-30 08:22:33',
                    'partner_ids': [[7, "Sacha Du Bourg-Palette"]],
# list of partner name_get
                    'message_type': u'comment',
                    'id': 59,
                    'subject': False
                    'is_note': True # only if the subtype is internal
                }
        `,
func(rs m.MailMessageSet)  {
//        message_values = self.read([
//            'id', 'body', 'date', 'author_id', 'email_from',  # base message fields
//            'message_type', 'subtype_id', 'subject',  # message specific
//            'model', 'res_id', 'record_name',  # document related
//            'channel_ids', 'partner_ids',  # recipients
//            'needaction_partner_ids',  # list of partner ids for whom the message is a needaction
//            'starred_partner_ids',  # list of partner ids for whom the message is starred
//        ])
//        message_tree = dict((m.id, m) for m in self.sudo())
//        self._message_read_dict_postprocess(message_values, message_tree)
//        subtypes = self.env['mail.message.subtype'].sudo().search(
//            [('id', 'in', [msg['subtype_id'][0] for msg in message_values if msg['subtype_id']])]).read(['internal', 'description'])
//        subtypes_dict = dict((subtype['id'], subtype) for subtype in subtypes)
//        for message in message_values:
//            message['is_note'] = message['subtype_id'] and subtypes_dict[message['subtype_id'][0]]['internal']
//            message['subtype_description'] = message['subtype_id'] and subtypes_dict[message['subtype_id'][0]]['description']
//        return message_values
})
h.MailMessage().Methods().Init().DeclareMethod(
`Init`,
func(rs m.MailMessageSet)  {
//        self._cr.execute(
//            """SELECT indexname FROM pg_indexes WHERE indexname = 'mail_message_model_res_id_idx'""")
//        if not self._cr.fetchone():
//            self._cr.execute(
//                """CREATE INDEX mail_message_model_res_id_idx ON mail_message (model, res_id)""")
})
h.MailMessage().Methods().FindAllowedModelWise().DeclareMethod(
`FindAllowedModelWise`,
func(rs m.MailMessageSet, doc_model interface{}, doc_dict interface{})  {
//        doc_ids = doc_dict.keys()
//        allowed_doc_ids = self.env[doc_model].with_context(
//            active_test=False).search([('id', 'in', doc_ids)]).ids
//        return set([message_id for allowed_doc_id in allowed_doc_ids for message_id in doc_dict[allowed_doc_id]])
})
h.MailMessage().Methods().FindAllowedDocIds().DeclareMethod(
`FindAllowedDocIds`,
func(rs m.MailMessageSet, model_ids interface{})  {
//        IrModelAccess = self.env['ir.model.access']
//        allowed_ids = set()
//        for doc_model, doc_dict in model_ids.iteritems():
//            if not IrModelAccess.check(doc_model, 'read', False):
//                continue
//            allowed_ids |= self._find_allowed_model_wise(doc_model, doc_dict)
//        return allowed_ids
})
h.MailMessage().Methods().Search().Extend(
` Override that adds specific access rights of mail.message, to remove
        ids uid could not see according to our custom rules.
Please refer to
        check_access_rule for more details about those rules.

        Non employees users see only message with subtype
(aka do not see
        internal logs).

        After having received ids of a classic search, keep only:
        - if author_id == pid, uid is the author, OR
        - uid belongs to a notified channel, OR
        - uid is in the specified recipients, OR
        - uid has a notification on the message, OR
        - uid have read access to the related document is model, res_id
        - otherwise: remove the id
        `,
func(rs m.MailMessageSet, args models.Conditioner, offset interface{}, limit interface{}, order interface{}, count interface{}, access_rights_uid interface{})  {
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
//        self._cr.execute("""
//            SELECT DISTINCT m.id, m.model, m.res_id, m.author_id,
//                            COALESCE(partner_rel.res_partner_id, needaction_rel.res_partner_id),
//                            channel_partner.channel_id as channel_id
//            FROM "%s" m
//            LEFT JOIN "mail_message_res_partner_rel" partner_rel
//            ON partner_rel.mail_message_id = m.id AND partner_rel.res_partner_id = %%(pid)s
//            LEFT JOIN "mail_message_res_partner_needaction_rel" needaction_rel
//            ON needaction_rel.mail_message_id = m.id AND needaction_rel.res_partner_id = %%(pid)s
//            LEFT JOIN "mail_message_mail_channel_rel" channel_rel
//            ON channel_rel.mail_message_id = m.id
//            LEFT JOIN "mail_channel" channel
//            ON channel.id = channel_rel.mail_channel_id
//            LEFT JOIN "mail_channel_partner" channel_partner
//            ON channel_partner.channel_id = channel.id AND channel_partner.partner_id = %%(pid)s
//            WHERE m.id = ANY (%%(ids)s)""" % self._table, dict(pid=pid, ids=ids))
//        for id, rmod, rid, author_id, partner_id, channel_id in self._cr.fetchall():
//            if author_id == pid:
//                author_ids.add(id)
//            elif partner_id == pid:
//                partner_ids.add(id)
//            elif channel_id:
//                channel_ids.add(id)
//            elif rmod and rid:
//                model_ids.setdefault(rmod, {}).setdefault(rid, set()).add(id)
//        allowed_ids = self._find_allowed_doc_ids(model_ids)
//        final_ids = author_ids | partner_ids | channel_ids | allowed_ids
//        if count:
//            return len(final_ids)
//        else:
//            # re-construct a list based on ids, because set did not keep the original order
//            id_list = [id for id in ids if id in final_ids]
//            return id_list
})
h.MailMessage().Methods().CheckAccessRule().DeclareMethod(
` Access rules of mail.message:
            - read: if
                - author_id == pid, uid is the author OR
                - uid is in the recipients (partner_ids) OR
                - uid has been notified (needaction) OR
                - uid is member of a listern channel (channel_ids.partner_ids)
OR
                - uid have read access to the related document
if model, res_id
                - otherwise: raise
            - create: if
                - no model, no res_id (private message) OR
                - pid in message_follower_ids if model, res_id OR
                - uid can read the parent OR
                - uid have write or create access on the
related document if model, res_id, OR
                - otherwise: raise
            - write: if
                - author_id == pid, uid is the author, OR
                - uid is in the recipients (partner_ids) OR
                - uid has write or create access on the
related document if model, res_id
                - otherwise: raise
            - unlink: if
                - uid has write or create access on the
related document if model, res_id
                - otherwise: raise

        Specific case: non employee users see only messages
with subtype (aka do
        not see internal logs).
        `,
func(rs m.MailMessageSet, operation interface{})  {
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
//        if operation in ['read', 'write']:
//            self._cr.execute("""
//                SELECT DISTINCT m.id, m.model, m.res_id, m.author_id, m.parent_id,
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
//                WHERE m.id = ANY (%%(ids)s)""" % self._table, dict(pid=self.env.user.partner_id.id, ids=self.ids))
//            for mid, rmod, rid, author_id, parent_id, partner_id, channel_id in self._cr.fetchall():
//                message_values[mid] = {
//                    'model': rmod,
//                    'res_id': rid,
//                    'author_id': author_id,
//                    'parent_id': parent_id,
//                    'notified': any((message_values[mid].get('notified'), partner_id, channel_id))
//                }
//        else:
//            self._cr.execute(
//                """SELECT DISTINCT id, model, res_id, author_id, parent_id FROM "%s" WHERE id = ANY (%%s)""" % self._table, (self.ids))
//            for mid, rmod, rid, author_id, parent_id in self._cr.fetchall():
//                message_values[mid] = {
//                    'model': rmod, 'res_id': rid, 'author_id': author_id, 'parent_id': parent_id}
//        author_ids = []
//        if operation == 'read' or operation == 'write':
//            author_ids = [mid for mid, message in message_values.iteritems()
//                          if message.get('author_id') and message.get('author_id') == self.env.user.partner_id.id]
//        elif operation == 'create':
//            author_ids = [mid for mid, message in message_values.iteritems()
//                          if not message.get('model') and not message.get('res_id')]
//        notified_ids = []
//        if operation == 'create':
//            # TDE: probably clean me
//            parent_ids = [message.get('parent_id') for mid, message in message_values.iteritems()
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
//            notified_ids += [mid for mid, message in message_values.iteritems()
//                             if message.get('parent_id') in not_parent_ids]
//        other_ids = set(self.ids).difference(
//            set(author_ids), set(notified_ids))
//        model_record_ids = _generate_model_record_ids(
//            message_values, other_ids)
//        if operation in ['read', 'write']:
//            notified_ids = [
//                mid for mid, message in message_values.iteritems() if message.get('notified')]
//        elif operation == 'create':
//            for doc_model, doc_ids in model_record_ids.items():
//                followers = self.env['mail.followers'].sudo().search([
//                    ('res_model', '=', doc_model),
//                    ('res_id', 'in', list(doc_ids)),
//                    ('partner_id', '=', self.env.user.partner_id.id),
//                ])
//                fol_mids = [follower.res_id for follower in followers]
//                notified_ids += [mid for mid, message in message_values.iteritems()
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
//            document_related_ids += [mid for mid, message in message_values.iteritems()
//                                     if message.get('model') == model and message.get('res_id') in mids.ids]
//        other_ids = other_ids.difference(set(document_related_ids))
//        if not other_ids:
//            return
//        raise AccessError(
//            _('The requested operation cannot be completed due to security restrictions. Please contact your system administrator.\n\n(Document type: %s, Operation: %s)') % (
//                self._description, operation)
//            + ' - ({} {}, {} {})'.format(_('Records:'),
//                                         list(other_ids)[:6], _('User:'), self._uid)
//        )
})
h.MailMessage().Methods().GetRecordName().DeclareMethod(
` Return the related document name, using name_get. It is done using
            SUPERUSER_ID, to be sure to have the record
name correctly stored. `,
func(rs m.MailMessageSet, values interface{})  {
//        model = values.get('model', self.env.context.get('default_model'))
//        res_id = values.get('res_id', self.env.context.get('default_res_id'))
//        if not model or not res_id or model not in self.env:
//            return False
//        return self.env[model].sudo().browse(res_id).name_get()[0][1]
})
h.MailMessage().Methods().GetReplyTo().DeclareMethod(
` Return a specific reply_to: alias of the document through
        message_get_reply_to or take the email_from `,
func(rs m.MailMessageSet, values interface{})  {
//        model, res_id, email_from = values.get('model', self._context.get('default_model')), values.get(
//            'res_id', self._context.get('default_res_id')), values.get('email_from')  # ctx values / defualt_get res ?
//        if model:
//            # return self.env[model].browse(res_id).message_get_reply_to([res_id], default=email_from)[res_id]
//            return self.env[model].message_get_reply_to([res_id], default=email_from)[res_id]
//        else:
//            # return self.env['mail.thread'].message_get_reply_to(default=email_from)[None]
//            return self.env['mail.thread'].message_get_reply_to([None], default=email_from)[None]
})
h.MailMessage().Methods().GetMessageId().DeclareMethod(
`GetMessageId`,
func(rs m.MailMessageSet, values interface{})  {
//        if values.get('no_auto_thread', False) is True:
//            message_id = tools.generate_tracking_message_id('reply_to')
//        elif values.get('res_id') and values.get('model'):
//            message_id = tools.generate_tracking_message_id(
//                '%(res_id)s-%(model)s' % values)
//        else:
//            message_id = tools.generate_tracking_message_id('private')
//        return message_id
})
h.MailMessage().Methods().InvalidateDocuments().DeclareMethod(
` Invalidate the cache of the documents followed by ``self``. `,
func(rs m.MailMessageSet)  {
//        for record in self:
//            if record.model and record.res_id:
//                self.env[record.model].invalidate_cache(ids=[record.res_id])
})
h.MailMessage().Methods().Create().Extend(
`Create`,
func(rs m.MailMessageSet, values models.RecordData)  {
//        if self.env.context.get('default_starred'):
//            self = self.with_context({'default_starred_partner_ids': [
//                                     (4, self.env.user.partner_id.id)]})
//        if 'email_from' not in values:  # needed to compute reply_to
//            values['email_from'] = self._get_default_from()
//        if not values.get('message_id'):
//            values['message_id'] = self._get_message_id(values)
//        if 'reply_to' not in values:
//            values['reply_to'] = self._get_reply_to(values)
//        if 'record_name' not in values and 'default_record_name' not in self.env.context:
//            values['record_name'] = self._get_record_name(values)
//        tracking_values_cmd = values.pop('tracking_value_ids', False)
//        message = super(Message, self).create(values)
//        if values.get('attachment_ids'):
//            message.attachment_ids.check(mode='read')
//        if tracking_values_cmd:
//            message.sudo().write({'tracking_value_ids': tracking_values_cmd})
//        message._invalidate_documents()
//        if not self.env.context.get('message_create_from_mail_mail'):
//            message._notify(force_send=self.env.context.get('mail_notify_force_send', True),
//                            user_signature=self.env.context.get('mail_notify_user_signature', True))
//        return message
})
h.MailMessage().Methods().Read().Extend(
` Override to explicitely call check_access_rule, that is not called
            by the ORM. It instead directly fetches ir.rules
and apply them. `,
func(rs m.MailMessageSet, fields []string, load interface{})  {
//        self.check_access_rule('read')
//        return super(Message, self).read(fields=fields, load=load)
})
h.MailMessage().Methods().Write().Extend(
`Write`,
func(rs m.MailMessageSet, vals models.RecordData)  {
//        if 'model' in vals or 'res_id' in vals:
//            self._invalidate_documents()
//        res = super(Message, self).write(vals)
//        if vals.get('attachment_ids'):
//            for mail in self:
//                mail.attachment_ids.check(mode='read')
//        self._invalidate_documents()
//        return res
})
h.MailMessage().Methods().Unlink().Extend(
`Unlink`,
func(rs m.MailMessageSet)  {
//        self.check_access_rule('unlink')
//        self.mapped('attachment_ids').filtered(
//            lambda attach: attach.res_model == self._name and (
//                attach.res_id in self.ids or attach.res_id == 0)
//        ).unlink()
//        self._invalidate_documents()
//        return super(Message, self).unlink()
})
h.MailMessage().Methods().Notify().DeclareMethod(
` Add the related record followers to the destination partner_ids
if is not a private message.
            Call mail_notification.notify to manage the email sending
        `,
func(rs m.MailMessageSet, force_send interface{}, send_after_commit interface{}, user_signature interface{})  {
//        group_user = self.env.ref('base.group_user')
//        self_sudo = self.sudo()
//        self.ensure_one()  # tde: not sure, just for testinh, will see
//        partners = self.env['res.partner'] | self.partner_ids
//        channels = self.env['mail.channel'] | self.channel_ids
//        if self_sudo.subtype_id and self.model and self.res_id:
//            followers = self.env['mail.followers'].sudo().search([
//                ('res_model', '=', self.model),
//                ('res_id', '=', self.res_id)
//            ]).filtered(lambda fol: self.subtype_id in fol.subtype_ids)
//            if self_sudo.subtype_id.internal:
//                followers = followers.filtered(lambda fol: fol.channel_id or (
//                    fol.partner_id.user_ids and group_user in fol.partner_id.user_ids[0].mapped('groups_id')))
//            channels = self_sudo.channel_ids | followers.mapped('channel_id')
//            partners = self_sudo.partner_ids | followers.mapped('partner_id')
//        else:
//            channels = self_sudo.channel_ids
//            partners = self_sudo.partner_ids
//        if not self._context.get('mail_notify_author', False) and self_sudo.author_id:
//            partners = partners - self_sudo.author_id
//        message_values = {
//            'channel_ids': [(6, 0, channels.ids)],
//            'needaction_partner_ids': [(6, 0, partners.ids)]
//        }
//        if self.model and self.res_id and hasattr(self.env[self.model], 'message_get_message_notify_values'):
//            message_values.update(self.env[self.model].browse(
//                self.res_id).message_get_message_notify_values(self, message_values))
//        self.write(message_values)
//        partners._notify(self, force_send=force_send,
//                         send_after_commit=send_after_commit, user_signature=user_signature)
//        channels._notify(self)
//        if self.parent_id:
//            self.parent_id.invalidate_cache()
//        return True
})
}