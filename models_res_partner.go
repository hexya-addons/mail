package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

//import logging
//import threading
//_logger = logging.getLogger(__name__)
func init() {

	//    _mail_flat_thread = False
	//    _mail_mass_mailing = _('Customers')
	h.Partner().AddFields(map[string]models.FieldDefinition{
		"MessageBounce": models.IntegerField{
			String: "Bounce",
			Help:   "Counter of the number of bounced emails for this contact",
		},
		"NotifyEmail": models.SelectionField{
			Selection: types.Selection{
				"none":   "Never",
				"always": "All Messages",
			},
			String:   "Email Messages and Notifications",
			Required: true,
			//oldname='notification_email_send'
			Default: models.DefaultValue("always"),
			Help: "Policy to receive emails for new messages pushed to your personal Inbox:" +
				"- Never: no emails are sent" +
				"- All Messages: for every notification you receive in your Inbox",
		},
		"OptOut": models.BooleanField{
			String: "Opt-Out",
			Help: "If opt-out is checked, this contact has refused to receive" +
				"emails for mass mailing and marketing campaign. Filter" +
				"'Available for Mass Mailing' allows users to filter the" +
				"partners when performing mass mailing.",
		},
		"ChannelIds": models.Many2ManyField{
			RelationModel:    h.MailChannel(),
			M2MLinkModelName: "",
			M2MOurField:      "",
			M2MTheirField:    "",
			String:           "Channels",
			NoCopy:           true,
		},
	})
	h.Partner().Methods().MessageGetSuggestedRecipients().DeclareMethod(
		`MessageGetSuggestedRecipients`,
		func(rs m.PartnerSet) {
			//        recipients = super(Partner, self).message_get_suggested_recipients()
			//        for partner in self:
			//            partner._message_add_suggested_recipient(
			//                recipients, partner=partner, reason=_('Partner Profile'))
			//        return recipients
		})
	h.Partner().Methods().MessageGetDefaultRecipients().DeclareMethod(
		`MessageGetDefaultRecipients`,
		func(rs m.PartnerSet) {
			//        return dict((res_id, {'partner_ids': [res_id], 'email_to': False, 'email_cc': False}) for res_id in self.ids)
		})
	h.Partner().Methods().NotifyPrepareTemplateContext().DeclareMethod(
		`NotifyPrepareTemplateContext`,
		func(rs m.PartnerSet, message interface{}) {
			//        signature = ""
			//        if message.author_id and message.author_id.user_ids and message.author_id.user_ids[0].signature:
			//            signature = message.author_id.user_ids[0].signature
			//        elif message.author_id:
			//            signature = "<p>-- <br/>%s</p>" % message.author_id.name
			//        if message.author_id and message.author_id.user_ids:
			//            user = message.author_id.user_ids[0]
			//        else:
			//            user = self.env.user
			//        if user.company_id.website:
			//            website_url = 'http://%s' % user.company_id.website if not user.company_id.website.lower().startswith(('http:',
			//                                                                                                                   'https:')) else user.company_id.website
			//        else:
			//            website_url = False
			//        model_name = False
			//        if message.model:
			//            model_name = self.env['ir.model'].sudo().search(
			//                [('model', '=', self.env[message.model]._name)]).name_get()[0][1]
			//        record_name = message.record_name
			//        tracking = []
			//        for tracking_value in self.env['mail.tracking.value'].sudo().search([('mail_message_id', '=', message.id)]):
			//            tracking.append((tracking_value.field_desc,
			//                             tracking_value.get_old_display_value()[0],
			//                             tracking_value.get_new_display_value()[0]))
			//        is_discussion = message.subtype_id.id == self.env['ir.model.data'].xmlid_to_res_id(
			//            'mail.mt_comment')
			//        record = False
			//        if message.res_id and message.model in self.env:
			//            record = self.env[message.model].browse(message.res_id)
			//        company = user.company_id
			//        if record and hasattr(record, 'company_id'):
			//            company = record.company_id
			//        company_name = company.name
			//        return {
			//            'signature': signature,
			//            'website_url': website_url,
			//            'company': company,
			//            'company_name': company_name,
			//            'model_name': model_name,
			//            'record': record,
			//            'record_name': record_name,
			//            'tracking': tracking,
			//            'is_discussion': is_discussion,
			//            'subtype': message.subtype_id,
			//        }
		})
	h.Partner().Methods().NotifyPrepareEmailValues().DeclareMethod(
		`NotifyPrepareEmailValues`,
		func(rs m.PartnerSet, message interface{}) {
			//        references = message.parent_id.message_id if message.parent_id else False
			//        custom_values = dict()
			//        if message.res_id and message.model in self.env and hasattr(self.env[message.model], 'message_get_email_values'):
			//            custom_values = self.env[message.model].browse(
			//                message.res_id).message_get_email_values(message)
			//        mail_values = {
			//            'mail_message_id': message.id,
			//            'mail_server_id': message.mail_server_id.id,
			//            'auto_delete': self._context.get('mail_auto_delete', True),
			//            'references': references,
			//        }
			//        mail_values.update(custom_values)
			//        return mail_values
		})
	h.Partner().Methods().NotifySend().DeclareMethod(
		`NotifySend`,
		func(rs m.PartnerSet, body interface{}, subject interface{}, recipients interface{}) {
			//        emails = self.env['mail.mail']
			//        recipients_nbr, recipients_max = len(recipients), 50
			//        email_chunks = [recipients[x:x + recipients_max]
			//                        for x in xrange(0, len(recipients), recipients_max)]
			//        for email_chunk in email_chunks:
			//            # TDE FIXME: missing message parameter. So we will find mail_message_id
			//            # in the mail_values and browse it. It should already be in the
			//            # cache so should not impact performances.
			//            mail_message_id = mail_values.get('mail_message_id')
			//            message = self.env['mail.message'].browse(
			//                mail_message_id) if mail_message_id else None
			//            if message and message.model and message.res_id and message.model in self.env and hasattr(self.env[message.model], 'message_get_recipient_values'):
			//                tig = self.env[message.model].browse(message.res_id)
			//                recipient_values = tig.message_get_recipient_values(
			//                    notif_message=message, recipient_ids=email_chunk.ids)
			//            else:
			//                recipient_values = self.env['mail.thread'].message_get_recipient_values(
			//                    notif_message=None, recipient_ids=email_chunk.ids)
			//            create_values = {
			//                'body_html': body,
			//                'subject': subject,
			//            }
			//            create_values.update(mail_values)
			//            create_values.update(recipient_values)
			//            emails |= self.env['mail.mail'].create(create_values)
			//        return emails, recipients_nbr
		})
	h.Partner().Methods().NotifyUdpateNotifications().DeclareMethod(
		`NotifyUdpateNotifications`,
		func(rs m.PartnerSet, emails interface{}) {
			//        for email in emails:
			//            notifications = self.env['mail.notification'].sudo().search([
			//                ('mail_message_id', '=', email.mail_message_id.id),
			//                ('res_partner_id', 'in', email.recipient_ids.ids)])
			//            notifications.write({
			//                'is_email': True,
			//                'email_status': 'ready',
			//            })
		})
	h.Partner().Methods().Notify().DeclareMethod(
		`Notify`,
		func(rs m.PartnerSet, message interface{}, force_send interface{}, send_after_commit interface{}, user_signature interface{}) {
			//        message_sudo = message.sudo()
			//        email_channels = message.channel_ids.filtered(
			//            lambda channel: channel.email_send)
			//        self.sudo().search([
			//            '|',
			//            ('id', 'in', self.ids),
			//            ('channel_ids', 'in', email_channels.ids),
			//            ('email', '!=', message_sudo.author_id and message_sudo.author_id.email or message.email_from),
			//            ('notify_email', '!=', 'none')])._notify_by_email(message, force_send=force_send, send_after_commit=send_after_commit, user_signature=user_signature)
			//        self._notify_by_chat(message)
			//        return True
		})
	h.Partner().Methods().NotifyByEmail().DeclareMethod(
		` Method to send email linked to notified messages. The recipients are
        the recordset on which this method is called.

        :param boolean force_send: send notification emails
now instead of letting the scheduler handle the email queue
        :param boolean send_after_commit: send notification
emails after the transaction end instead of durign the
                                          transaction;
this option is used only if force_send is True
        :param user_signature: add current user signature
to notification emails `,
		func(rs m.PartnerSet, message interface{}, force_send interface{}, send_after_commit interface{}, user_signature interface{}) {
			//        if not self.ids:
			//            return True
			//        base_template = None
			//        if message.model and self._context.get('custom_layout', False):
			//            base_template = self.env.ref(
			//                self._context['custom_layout'], raise_if_not_found=False)
			//        if not base_template:
			//            base_template = self.env.ref(
			//                'mail.mail_template_data_notification_email_default')
			//        base_template_ctx = self._notify_prepare_template_context(message)
			//        if not user_signature:
			//            base_template_ctx['signature'] = False
			//        base_mail_values = self._notify_prepare_email_values(message)
			//        if message.model and message.res_id and hasattr(self.env[message.model], '_message_notification_recipients'):
			//            recipients = self.env[message.model].browse(
			//                message.res_id)._message_notification_recipients(message, self)
			//        else:
			//            recipients = self.env['mail.thread']._message_notification_recipients(
			//                message, self)
			//        emails = self.env['mail.mail']
			//        recipients_nbr, recipients_max = 0, 50
			//        for email_type, recipient_template_values in recipients.iteritems():
			//            if recipient_template_values['followers']:
			//                # generate notification email content
			//                # fixme: set button_unfollow to none
			//                template_fol_values = dict(
			//                    base_template_ctx, **recipient_template_values)
			//                template_fol_values['has_button_follow'] = False
			//                template_fol = base_template.with_context(
			//                    **template_fol_values)
			//                # generate templates for followers and not followers
			//                fol_values = template_fol.generate_email(
			//                    message.id, fields=['body_html', 'subject'])
			//                # send email
			//                new_emails, new_recipients_nbr = self._notify_send(
			//                    fol_values['body'], fol_values['subject'], recipient_template_values['followers'], **base_mail_values)
			//                # update notifications
			//                self._notify_udpate_notifications(new_emails)
			//
			//                emails |= new_emails
			//                recipients_nbr += new_recipients_nbr
			//            if recipient_template_values['not_followers']:
			//                # generate notification email content
			//                # fixme: set button_follow to none
			//                template_not_values = dict(
			//                    base_template_ctx, **recipient_template_values)
			//                template_not_values['has_button_unfollow'] = False
			//                template_not = base_template.with_context(
			//                    **template_not_values)
			//                # generate templates for followers and not followers
			//                not_values = template_not.generate_email(
			//                    message.id, fields=['body_html', 'subject'])
			//                # send email
			//                new_emails, new_recipients_nbr = self._notify_send(
			//                    not_values['body'], not_values['subject'], recipient_template_values['not_followers'], **base_mail_values)
			//                # update notifications
			//                self._notify_udpate_notifications(new_emails)
			//
			//                emails |= new_emails
			//                recipients_nbr += new_recipients_nbr
			//        test_mode = getattr(threading.currentThread(), 'testing', False)
			//        if force_send and recipients_nbr < recipients_max and \
			//                (not self.pool._init or test_mode):
			//            email_ids = emails.ids
			//            dbname = self.env.cr.dbname
			//            _context = self._context
			//
			//            def send_notifications():
			//                db_registry = registry(dbname)
			//                with api.Environment.manage(), db_registry.cursor() as cr:
			//                    env = api.Environment(cr, SUPERUSER_ID, _context)
			//                    env['mail.mail'].browse(email_ids).send()
			//
			//            # unless asked specifically, send emails after the transaction to
			//            # avoid side effects due to emails being sent while the transaction fails
			//            if not test_mode and send_after_commit:
			//                self._cr.after('commit', send_notifications)
			//            else:
			//                emails.send()
			//        return True
		})
	h.Partner().Methods().NotifyByChat().DeclareMethod(
		` Broadcast the message to all the partner since `,
		func(rs m.PartnerSet, message interface{}) {
			//        message_values = message.message_format()[0]
			//        notifications = []
			//        for partner in self:
			//            notifications.append(
			//                [(self._cr.dbname, 'ir.needaction', partner.id), dict(message_values)])
			//        self.env['bus.bus'].sendmany(notifications)
		})
	h.Partner().Methods().GetNeedactionCount().DeclareMethod(
		` compute the number of needaction of the current user `,
		func(rs m.PartnerSet) {
			//        if self.env.user.partner_id:
			//            self.env.cr.execute("""
			//                SELECT count(*) as needaction_count
			//                FROM mail_message_res_partner_needaction_rel R
			//                WHERE R.res_partner_id = %s AND (R.is_read = false OR R.is_read IS NULL)""", (self.env.user.partner_id.id))
			//            return self.env.cr.dictfetchall()[0].get('needaction_count')
			//        _logger.error('Call to needaction_count without partner_id')
			//        return 0
		})
	h.Partner().Methods().GetStarredCount().DeclareMethod(
		` compute the number of starred of the current user `,
		func(rs m.PartnerSet) {
			//        if self.env.user.partner_id:
			//            self.env.cr.execute("""
			//                SELECT count(*) as starred_count
			//                FROM mail_message_res_partner_starred_rel R
			//                WHERE R.res_partner_id = %s """, (self.env.user.partner_id.id))
			//            return self.env.cr.dictfetchall()[0].get('starred_count')
			//        _logger.error('Call to starred_count without partner_id')
			//        return 0
		})
	h.Partner().Methods().GetStaticMentionSuggestions().DeclareMethod(
		` To be overwritten to return the id, name and email of
partners used as static mention
            suggestions loaded once at webclient initialization
and stored client side. `,
		func(rs m.PartnerSet) {
			//        return []
		})
	h.Partner().Methods().GetMentionSuggestions().DeclareMethod(
		` Return 'limit'-first partners' id, name and email such
that the name or email matches a
            'search' string. Prioritize users, and then
extend the research to all partners. `,
		func(rs m.PartnerSet, search interface{}, limit interface{}) {
			//        search_dom = expression.OR(
			//            [[('name', 'ilike', search)], [('email', 'ilike', search)]])
			//        fields = ['id', 'name', 'email']
			//        domain = expression.AND([[('user_ids.id', '!=', False)], search_dom])
			//        users = self.search_read(domain, fields, limit=limit)
			//        partners = []
			//        if len(users) < limit:
			//            partners = self.search_read(search_dom, fields, limit=limit)
			//            # Remove duplicates
			//            partners = [p for p in partners if not len(
			//                [u for u in users if u['id'] == p['id']])]
			//        return [users, partners]
		})
}
