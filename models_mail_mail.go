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
	
//import base64
//import datetime
//import logging
//import psycopg2
//import threading
//_logger = logging.getLogger(__name__)
func init() {
h.MailMail().DeclareModel()





h.MailMail().AddFields(map[string]models.FieldDefinition{
"MailMessageId": models.Many2OneField{
RelationModel: h.MailMessage(),
String: "Message",
Required: true,
OnDelete: `cascade`,
Index: true,

},
"BodyHtml": models.TextField{
String: "Rich-text Contents",
Help: "Rich-text/HTML message",
},
"References": models.TextField{
String: "References",
Help: "Message references, such as identifiers of previous messages",
ReadOnly: true,
},
"Headers": models.TextField{
String: "Headers",
NoCopy: true,
},
"Notification": models.BooleanField{
String: "Is Notification",
Help: "Mail has been created to notify people of an existing mail.message",
},
"EmailTo": models.TextField{
String: "To",
Help: "Message recipients (emails)",
},
"EmailCc": models.CharField{
String: "Cc",
Help: "Carbon copy message recipients",
},
"RecipientIds": models.Many2ManyField{
RelationModel: h.Partner(),
String: "To (Partners)",
},
"State": models.SelectionField{
Selection: types.Selection{
"outgoing": "Outgoing",
"sent": "Sent",
"received": "Received",
"exception": "Delivery Failed",
"cancel": "Cancelled",
},
String: "Status",
ReadOnly: true,
NoCopy: true,
Default: models.DefaultValue("outgoing"),
},
"AutoDelete": models.BooleanField{
String: "Auto Delete",
Help: "Permanently delete this email after sending it, to save space",
},
"FailureReason": models.TextField{
String: "Failure Reason",
ReadOnly: true,
Help: "Failure reason. This is usually the exception thrown by" + 
"the email server, stored to ease the debugging of mailing issues.",
},
"ScheduledDate": models.CharField{
String: "Scheduled Send Date",
Help: "If set, the queue manager will send the email after the" + 
"date. If not set, the email will be send as soon as possible.",
},
})
h.MailMail().Methods().Create().Extend(
`Create`,
func(rs m.MailMailSet, values models.RecordData)  {
//        if 'notification' not in values and values.get('mail_message_id'):
//            values['notification'] = True
//        if not values.get('mail_message_id'):
//            self = self.with_context(message_create_from_mail_mail=True)
//        new_mail = super(MailMail, self).create(values)
//        if values.get('attachment_ids'):
//            new_mail.attachment_ids.check(mode='read')
//        return new_mail
})
h.MailMail().Methods().Write().Extend(
`Write`,
func(rs m.MailMailSet, vals models.RecordData)  {
//        res = super(MailMail, self).write(vals)
//        if vals.get('attachment_ids'):
//            for mail in self:
//                mail.attachment_ids.check(mode='read')
//        return res
})
h.MailMail().Methods().Unlink().Extend(
`Unlink`,
func(rs m.MailMailSet)  {
//        to_cascade = self.search(
//            [('notification', '=', False), ('id', 'in', self.ids)]).mapped('mail_message_id')
//        res = super(MailMail, self).unlink()
//        to_cascade.unlink()
//        return res
})
h.MailMail().Methods().DefaultGet().Extend(
`DefaultGet`,
func(rs m.MailMailSet, fields interface{})  {
//        if self._context.get('default_type') not in type(self).message_type.base_field.selection:
//            self = self.with_context(dict(self._context, default_type=None))
//        return super(MailMail, self).default_get(fields)
})
h.MailMail().Methods().MarkOutgoing().DeclareMethod(
`MarkOutgoing`,
func(rs m.MailMailSet)  {
//        return self.write({'state': 'outgoing'})
})
h.MailMail().Methods().Cancel().DeclareMethod(
`Cancel`,
func(rs m.MailMailSet)  {
//        return self.write({'state': 'cancel'})
})
h.MailMail().Methods().ProcessEmailQueue().DeclareMethod(
`Send immediately queued messages, committing after each
           message is sent - this is not transactional and should
           not be called during another transaction!

           :param list ids: optional list of emails ids
to send. If passed
                            no search is performed, and
these ids are used
                            instead.
           :param dict context: if a 'filters' key is present
in context,
                                this value will be used as an additional
                                filter to further restrict the outgoing
                                messages to send (by default
all 'outgoing'
                                messages are sent).
        `,
func(rs m.MailMailSet, ids interface{})  {
//        if not self.ids:
//            filters = ['&',
//                       ('state', '=', 'outgoing'),
//                       '|',
//                       ('scheduled_date', '<', datetime.datetime.now()),
//                       ('scheduled_date', '=', False)]
//            if 'filters' in self._context:
//                filters.extend(self._context['filters'])
//            ids = self.search(filters).ids
//        res = None
//        try:
//            # auto-commit except in testing mode
//            auto_commit = not getattr(
//                threading.currentThread(), 'testing', False)
//            res = self.browse(ids).send(auto_commit=auto_commit)
//        except Exception:
//            _logger.exception("Failed processing mail queue")
//        return res
})
h.MailMail().Methods().PostprocessSentMessage().DeclareMethod(
`Perform any post-processing necessary after sending ``mail``
        successfully, including deleting it completely along with its
        attachment if the ``auto_delete`` flag of the mail was set.
        Overridden by subclasses for extra post-processing behaviors.

        :param browse_record mail: the mail that was just sent
        :return: True
        `,
func(rs m.MailMailSet, mail_sent interface{})  {
//        notif_emails = self.filtered(lambda email: email.notification)
//        if notif_emails:
//            notifications = self.env['mail.notification'].search([
//                ('mail_message_id', 'in', notif_emails.mapped('mail_message_id').ids),
//                ('res_partner_id', 'in', notif_emails.mapped('recipient_ids').ids),
//                ('is_email', '=', True)])
//            if mail_sent:
//                notifications.write({
//                    'email_status': 'sent',
//                })
//            else:
//                notifications.write({
//                    'email_status': 'exception',
//                })
//        if mail_sent:
//            self.sudo().filtered(lambda self: self.auto_delete).unlink()
//        return True
})
h.MailMail().Methods().SendGetMailBody().DeclareMethod(
`Return a specific ir_email body. The main purpose of this method
        is to be inherited to add custom content depending
on some module.`,
func(rs m.MailMailSet, partner interface{})  {
//        self.ensure_one()
//        body = self.body_html or ''
//        return body
})
h.MailMail().Methods().SendGetMailTo().DeclareMethod(
`Forge the email_to with the following heuristic:
          - if 'partner', recipient specific (Partner Name <email>)
          - else fallback on mail.email_to splitting `,
func(rs m.MailMailSet, partner interface{})  {
//        self.ensure_one()
//        if partner:
//            email_to = [formataddr((partner.name, partner.email))]
//        else:
//            email_to = tools.email_split_and_format(self.email_to)
//        return email_to
})
h.MailMail().Methods().SendGetEmailDict().DeclareMethod(
`Return a dictionary for specific email values, depending on a
        partner, or generic to the whole recipients given
by mail.email_to.

            :param browse_record mail: mail.mail browse_record
            :param browse_record partner: specific recipient partner
        `,
func(rs m.MailMailSet, partner interface{})  {
//        self.ensure_one()
//        body = self.send_get_mail_body(partner=partner)
//        body_alternative = tools.html2plaintext(body)
//        res = {
//            'body': body,
//            'body_alternative': body_alternative,
//            'email_to': self.send_get_mail_to(partner=partner),
//        }
//        return res
})
h.MailMail().Methods().Send().DeclareMethod(
` Sends the selected emails immediately, ignoring their current
            state (mails that have already been sent should
not be passed
            unless they should actually be re-sent).
            Emails successfully delivered are marked as
'sent', and those
            that fail to be deliver are marked as 'exception', and the
            corresponding error mail is output in the server logs.

            :param bool auto_commit: whether to force a
commit of the mail status
                after sending each mail (meant only for
scheduler processing);
                should never be True during normal transactions
(default: False)
            :param bool raise_exception: whether to raise
an exception if the
                email sending process has failed
            :return: True
        `,
func(rs m.MailMailSet, auto_commit interface{}, raise_exception interface{})  {
//        IrMailServer = self.env['ir.mail_server']
//        for mail_id in self.ids:
//            try:
//                mail = self.browse(mail_id)
//                # TDE note: remove me when model_id field is present on mail.message - done here to avoid doing it multiple times in the sub method
//                if mail.model:
//                    model = self.env['ir.model'].sudo().search(
//                        [('model', '=', mail.model)])[0]
//                else:
//                    model = None
//                if model:
//                    mail = mail.with_context(model_name=model.name)
//
//                # load attachment binary data with a separate read(), as prefetching all
//                # `datas` (binary field) could bloat the browse cache, triggerring
//                # soft/hard mem limits with temporary data.
//                attachments = [(a['datas_fname'], base64.b64decode(a['datas']))
//                               for a in mail.attachment_ids.sudo().read(['datas_fname', 'datas'])]
//
//                # specific behavior to customize the send email for notified partners
//                email_list = []
//                if mail.email_to:
//                    email_list.append(mail.send_get_email_dict())
//                for partner in mail.recipient_ids:
//                    email_list.append(
//                        mail.send_get_email_dict(partner=partner))
//
//                # headers
//                headers = {}
//                bounce_alias = self.env['ir.config_parameter'].get_param(
//                    "mail.bounce.alias")
//                catchall_domain = self.env['ir.config_parameter'].get_param(
//                    "mail.catchall.domain")
//                if bounce_alias and catchall_domain:
//                    if mail.model and mail.res_id:
//                        headers['Return-Path'] = '%s+%d-%s-%d@%s' % (
//                            bounce_alias, mail.id, mail.model, mail.res_id, catchall_domain)
//                    else:
//                        headers['Return-Path'] = '%s+%d@%s' % (
//                            bounce_alias, mail.id, catchall_domain)
//                if mail.headers:
//                    try:
//                        headers.update(safe_eval(mail.headers))
//                    except Exception:
//                        pass
//
//                # Writing on the mail object may fail (e.g. lock on user) which
//                # would trigger a rollback *after* actually sending the email.
//                # To avoid sending twice the same email, provoke the failure earlier
//                mail.write({
//                    'state': 'exception',
//                    'failure_reason': _('Error without exception. Probably due do sending an email without computed recipients.'),
//                })
//                mail_sent = False
//
//                # Update notification in a transient exception state to avoid concurrent
//                # update in case an email bounces while sending all emails related to current
//                # mail record.
//                notifs = self.env['mail.notification'].search([
//                    ('is_email', '=', True),
//                    ('mail_message_id', 'in', mail.mapped('mail_message_id').ids),
//                    ('res_partner_id', 'in', mail.mapped('recipient_ids').ids),
//                    ('email_status', 'not in', ('sent', 'canceled'))
//                ])
//                if notifs:
//                    notifs.sudo().write({
//                        'email_status': 'exception',
//                    })
//
//                # build an RFC2822 email.message.Message object and send it without queuing
//                res = None
//                for email in email_list:
//                    msg = IrMailServer.build_email(
//                        email_from=mail.email_from,
//                        email_to=email.get('email_to'),
//                        subject=mail.subject,
//                        body=email.get('body'),
//                        body_alternative=email.get('body_alternative'),
//                        email_cc=tools.email_split(mail.email_cc),
//                        reply_to=mail.reply_to,
//                        attachments=attachments,
//                        message_id=mail.message_id,
//                        references=mail.references,
//                        object_id=mail.res_id and (
//                            '%s-%s' % (mail.res_id, mail.model)),
//                        subtype='html',
//                        subtype_alternative='plain',
//                        headers=headers)
//                    try:
//                        res = IrMailServer.send_email(
//                            msg, mail_server_id=mail.mail_server_id.id)
//                    except AssertionError as error:
//                        if error.message == IrMailServer.NO_VALID_RECIPIENT:
//                            # No valid recipient found for this particular
//                            # mail item -> ignore error to avoid blocking
//                            # delivery to next recipients, if any. If this is
//                            # the only recipient, the mail will show as failed.
//                            _logger.info("Ignoring invalid recipients for mail.mail %s: %s",
//                                         mail.message_id, email.get('email_to'))
//                        else:
//                            raise
//                if res:
//                    mail.write({'state': 'sent', 'message_id': res,
//                                'failure_reason': False})
//                    mail_sent = True
//
//                # /!\ can't use mail.state here, as mail.refresh() will cause an error
//                # see revid:odo@openerp.com-20120622152536-42b2s28lvdv3odyr in 6.1
//                if mail_sent:
//                    _logger.info(
//                        'Mail with ID %r and Message-Id %r successfully sent', mail.id, mail.message_id)
//                mail._postprocess_sent_message(mail_sent=mail_sent)
//            except MemoryError:
//                # prevent catching transient MemoryErrors, bubble up to notify user or abort cron job
//                # instead of marking the mail as failed
//                _logger.exception(
//                    'MemoryError while processing mail with ID %r and Msg-Id %r. Consider raising the --limit-memory-hard startup option',
//                    mail.id, mail.message_id)
//                raise
//            except psycopg2.Error:
//                # If an error with the database occurs, chances are that the cursor is unusable.
//                # This will lead to an `psycopg2.InternalError` being raised when trying to write
//                # `state`, shadowing the original exception and forbid a retry on concurrent
//                # update. Let's bubble it.
//                raise
//            except Exception as e:
//                failure_reason = tools.ustr(e)
//                _logger.exception(
//                    'failed sending mail (id: %s) due to %s', mail.id, failure_reason)
//                mail.write(
//                    {'state': 'exception', 'failure_reason': failure_reason})
//                mail._postprocess_sent_message(mail_sent=False)
//                if raise_exception:
//                    if isinstance(e, AssertionError):
//                        # get the args of the original error, wrap into a value and throw a MailDeliveryException
//                        # that is an except_orm, with name and value as arguments
//                        value = '. '.join(e.args)
//                        raise MailDeliveryException(
//                            _("Mail Delivery Failed"), value)
//                    raise
//
//            if auto_commit is True:
//                self._cr.commit()
//        return True
})
}