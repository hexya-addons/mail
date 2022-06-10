package mail

import (
	"fmt"
	"net/mail"
	"strconv"
	"strings"

	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/models/types/dates"
	"github.com/hexya-erp/hexya/src/tools/emailutils"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
	"github.com/jaytaylor/html2text"
	"github.com/jordan-wright/email"
)

// MailMail is a model holding RFC2822 email messages to send.
// This model also provides facilities to queue and send new email messages.
var fields_MailMail = map[string]models.FieldDefinition{
	"MailMessage": fields.Many2One{
		RelationModel: h.MailMessage(),
		String:        "Message",
		Required:      true,
		OnDelete:      models.Cascade,
		Index:         true,
		Embed:         true,
	},
	"BodyHtml": fields.Text{
		String: "Rich-text Contents",
		Help:   "Rich-text/HTML message",
	},
	"References": fields.Text{
		String:   "References",
		Help:     "Message references, such as identifiers of previous messages",
		ReadOnly: true,
	},
	"Headers": fields.Text{
		String: "Headers",
		NoCopy: true,
	},
	// Auto-detected based on Create() - if 'MailMessage' was passed then this mail is a notification
	// and during Unlink() we will not cascade delete the parent and its attachments
	"Notification": fields.Boolean{
		String: "Is Notification",
		Help:   "Mail has been created to notify people of an existing mail.message",
	},
	// recipients: include inactive partners (they may have been archived after
	// the message was sent, but they should remain visible in the relation)
	"EmailTo": fields.Text{
		String: "To",
		Help:   "Message recipients (emails)",
	},
	"EmailCc": fields.Char{
		String: "Cc",
		Help:   "Carbon copy message recipients",
	},
	"Recipients": fields.Many2Many{
		RelationModel: h.Partner(),
		String:        "To (Partners)",
		Filter:        q.Partner().Active().Equals(true).Or().Active().Equals(false),
	},
	"State": fields.Selection{
		Selection: types.Selection{
			"outgoing":  "Outgoing",
			"sent":      "Sent",
			"received":  "Received",
			"exception": "Delivery Failed",
			"cancel":    "Cancelled",
		},
		String:   "Status",
		ReadOnly: true,
		NoCopy:   true,
		Default:  models.DefaultValue("outgoing"),
	},
	"AutoDelete": fields.Boolean{
		String: "Auto Delete",
		Help:   "Permanently delete this email after sending it, to save space",
	},
	"FailureReason": fields.Text{
		String:   "Failure Reason",
		ReadOnly: true,
		Help: "Failure reason. This is usually the exception thrown by" +
			"the email server, stored to ease the debugging of mailing issues.",
	},
	"ScheduledDate": fields.DateTime{
		String: "Scheduled Send Date",
		Help: "If set, the queue manager will send the email after the" +
			"date. If not set, the email will be send as soon as possible.",
	},
}

func mailMail_Create(rs m.MailMailSet, values m.MailMailData) m.MailMailSet {
	if !values.HasNotification() && values.MailMessage().IsNotEmpty() {
		values.SetNotification(true)
	}
	newMail := rs.Super().Create(values)
	if values.Attachments().IsNotEmpty() {
		newMail.Attachments().Check("read", nil)
	}
	return newMail
}

func mailMail_Write(rs m.MailMailSet, vals m.MailMailData) bool {
	res := rs.Super().Write(vals)
	if vals.Attachments().IsNotEmpty() {
		for _, mail := range rs.Records() {
			mail.Attachments().Check("read", nil)
		}
	}
	return res
}

func mailMail_Unlink(rs m.MailMailSet) int64 {
	var mailMsgCascadeIds []int64
	for _, mail := range rs.Records() {
		if !mail.Notification() {
			mailMsgCascadeIds = append(mailMsgCascadeIds, mail.MailMessage().ID())
		}
	}
	res := rs.Super().Unlink()
	if len(mailMsgCascadeIds) > 0 {
		h.MailMessage().Browse(rs.Env(), mailMsgCascadeIds).Unlink()
	}
	return res
}

func mailMail_DefaultGet(rs m.MailMailSet) m.MailMailData {
	if _, ok := h.MailMail().NewSet(rs.Env()).FieldGet(h.MailMail().Fields().MessageType()).Selection[rs.Env().Context().GetString("default_type")]; !ok {
		return rs.WithContext("default_type", "").Super().DefaultGet()
	}
	return rs.Super().DefaultGet()
}

// MarkOutgoing change theses mails' status to outgoing
func mailMail_MarkOutgoing(rs m.MailMailSet) bool {
	rs.SetState("outgoing")
	return true
}

// Cancel change theses mails' status to cancel
func mailMail_Cancel(rs m.MailMailSet) bool {
	rs.SetState("cancel")
	return true

}

// ProcessEmailQueue sends immediately queued messages, committing after each
// message is sent - this is not transactional and should not be called during
// another transaction!
//
// emails: is an optional recordset of MailMail to send.
//         If passed no search is performed, and these MailMail are used instead.
func mailMail_ProcessEmailQueue(rs m.MailMailSet, emails m.MailMailSet) bool {
	if emails.IsEmpty() {
		emails = h.MailMail().Search(rs.Env(), q.MailMail().
			State().Equals("outgoing").AndCond(
			q.MailMail().ScheduledDate().Lower(dates.Now()).Or().
				ScheduledDate().IsNull())).Limit(10000)
	}
	res := emails.Send(true, false)
	return res
}

// PostprocessSentMessage performs any post-processing necessary after sending a MailMail
// successfully, including deleting it completely along with its attachment if the
// AutoDelete flag of the mail was set.
// Overridden by subclasses for extra post-processing behaviors.
func mailMail_PostprocessSentMessage(rs m.MailMailSet, successPartners m.PartnerSet, failureReason string, failureType string) bool {
	notifMails := h.MailMail().NewSet(rs.Env())
	mailToDelete := h.MailMail().NewSet(rs.Env())
	for _, mail := range rs.Records() {
		if mail.Notification() {
			notifMails = notifMails.Union(mail)
		}
		if mail.AutoDelete() {
			mailToDelete = mailToDelete.Union(mail)
		}
	}
	defer func() {
		if failureType == "" || failureType == "RECIPIENT" {
			// if we have another error, we want to keep the mail.
			mailToDelete.Sudo().Unlink()
		}
	}()
	if notifMails.IsEmpty() {
		return true
	}
	notifications := h.MailNotification().Search(rs.Env(), q.MailNotification().
		IsEmail().Equals(true).And().
		Mail().In(notifMails).And().
		EmailStatus().NotIn([]string{"sent", "canceled"}))
	if notifications.IsEmpty() {
		return true
	}

	// find all notification linked to a failure
	failed := h.MailNotification().NewSet(rs.Env())
	if failureType != "" {
		failed = notifications.Filtered(func(r m.MailNotificationSet) bool {
			return r.Partner().Intersect(successPartners).IsEmpty()
		})
		failed.Sudo().Write(h.MailNotification().NewData().
			SetEmailStatus("exception").
			SetFailureType(failureType).
			SetFailureReason(failureReason))
		messages := h.MailMessage().NewSet(rs.Env())
		for _, notif := range notifications.Records() {
			if notif.MailMessage().ResID() != 0 && notif.MailMessage().ResModel() != "" {
				messages = messages.Union(notif.MailMessage())
			}
		}
		messages.NotifyFailureUpdate()
		notifications.Subtract(failed).Sudo().Write(h.MailNotification().NewData().
			SetEmailStatus("sent").
			SetFailureType("").
			SetFailureReason(""))
	}
	return true
}

// ------------------------------------------------------
// mail_mail formatting, tools and send mechanism
// ------------------------------------------------------

// SendPrepareBody return a specific email body.
// The main purpose of this method is to be inherited to add
// custom content depending on some module.
func mailMail_SendPrepareBody(rs m.MailMailSet) string {
	rs.EnsureOne()
	return rs.BodyHtml()
}

// SendPrepareValues return values for specific email values, depending on a
// partner, or generic to the whole recipients given by mail.email_to.
func mailMail_SendPrepareValues(rs m.MailMailSet, partner m.PartnerSet) m.MailMailData {
	rs.Env()
	body := rs.SendPrepareBody()
	bodyAlternative, err := html2text.FromString(body, html2text.Options{PrettyTables: true})
	if err != nil {
		bodyAlternative = ""
	}
	emailsTo := emailutils.SplitAddresses(rs.EmailTo())
	emailTo := strings.Join(emailsTo, ",")
	if partner.IsNotEmpty() {
		partnerAddr := mail.Address{
			Name:    partner.Name(),
			Address: partner.Email(),
		}
		emailTo = partnerAddr.String()
	}
	res := h.MailMail().NewData().
		SetBodyHtml(body).
		SetEmailTo(emailTo).
		SetBody(bodyAlternative)
	return res
}

// SplitByServer returns a list of MailMail recordsets with the same server.
//
// The same server may repeat in order to limit batch size according to
// the `mail.session.batch.size` system parameter.
func mailMail_SplitByServer(rs m.MailMailSet) []m.MailMailSet {
	groups := make(map[int64]m.MailMailSet)
	var res []m.MailMailSet
	batchSize, err := strconv.Atoi(h.ConfigParameter().NewSet(rs.Env()).Sudo().GetParam("mail.session.batch.size", "1000"))
	if err != nil {
		batchSize = 1000
	}
	rs.Load(h.MailMail().Fields().MailServer())
	for _, msg := range rs.Records() {
		if groups[msg.MailServer().ID()].Len() > batchSize {
			res = append(res, groups[msg.MailServer().ID()])
			delete(groups, msg.MailServer().ID())
		}
		groups[msg.MailServer().ID()] = groups[msg.MailServer().ID()].Union(msg)
	}
	for _, records := range groups {
		if records.IsNotEmpty() {
			res = append(res, records)
		}
	}
	return res
}

// Send sends the selected emails immediately, ignoring their current
// state (mails that have already been sent should not be passed
// unless they should actually be re-sent).
//
// Emails successfully delivered are marked as 'sent', and those
// that fail to be deliver are marked as 'exception', and the
// corresponding error mail is output in the server logs.
func mailMail_Send(rs m.MailMailSet, autoCommit bool, panicOnError bool) bool {
	//        for server_id, batch_ids in self._split_by_server():
	//            smtp_session = None
	//            try:
	//                smtp_session = self.env['ir.mail_server'].connect(
	//                    mail_server_id=server_id)
	//            except Exception as exc:
	//                if raise_exception:
	//                    # To be consistent and backward compatible with mail_mail.send() raised
	//                    # exceptions, it is encapsulated into an Odoo MailDeliveryException
	//                    raise MailDeliveryException(
	//                        _('Unable to connect to SMTP Server'), exc)
	//                else:
	//                    batch = self.browse(batch_ids)
	//                    batch.write({'state': 'exception', 'failure_reason': exc})
	//                    batch._postprocess_sent_message(
	//                        success_pids=[], failure_type="SMTP")
	//            else:
	//                self.browse(batch_ids)._send(
	//                    auto_commit=auto_commit,
	//                    raise_exception=raise_exception,
	//                    smtp_session=smtp_session)
	//                _logger.info(
	//                    'Sent batch %s emails via mail server ID #%s',
	//                    len(batch_ids), server_id)
	//            finally:
	//                if smtp_session:
	//                    smtp_session.quit()
	for _, batch := range rs.SplitByServer() {
		batch.DoSend(autoCommit, panicOnError)
	}
	return true
}

// DoSend sends these MailMail records.
func mailMail_DoSend(rs m.MailMailSet, autoCommit bool, panicOnError bool) bool {
	for _, msg := range rs.Records() {
		if msg.State() != "outgoing" {
			if msg.State() != "exception" && msg.AutoDelete() {
				msg.Sudo().Unlink()
			}
			continue
		}
		// TODO
		// # remove attachments if user send the link with the access_token
		// body = mail.body_html or ''
		// attachments = mail.attachment_ids
		// for link in re.findall(r'/web/(?:content|image)/([0-9]+)', body):
		//     attachments = attachments - IrAttachment.browse(int(link))
		//
		// # load attachment binary data with a separate read(), as prefetching all
		// # `datas` (binary field) could bloat the browse cache, triggerring
		// # soft/hard mem limits with temporary data.
		// attachments = [(a['datas_fname'], base64.b64decode(a['datas']), a['mimetype'])
		//                for a in attachments.sudo().read(['datas_fname', 'datas', 'mimetype']) if a['datas'] is not False]

		// specific behavior to customize the send email for notified partners
		var emailList []m.MailMailData
		if msg.EmailTo() != "" {
			emailList = append(emailList, msg.SendPrepareValues(nil))
		}
		for _, partner := range msg.Recipients().Records() {
			values := msg.SendPrepareValues(partner)
			values.SetPartners(partner)
			emailList = append(emailList, values)
		}

		// headers
		var emailMsg email.Email
		bounceAlias := h.ConfigParameter().NewSet(rs.Env()).Sudo().GetParam("mail.bounce.alias", "")
		catchallDomain := h.ConfigParameter().NewSet(rs.Env()).Sudo().GetParam("mail.catchall.domain", "")
		if bounceAlias != "" && catchallDomain != "" {
			emailMsg.Headers.Set("Return-Path", fmt.Sprintf("%s+%d@%s", bounceAlias, msg.ID(), catchallDomain))
			if msg.ResModel() != "" && msg.ResID() != 0 {
				emailMsg.Headers.Set("Return-Path", fmt.Sprintf("%s+%d-%s-%d@%s", bounceAlias, msg.ID(), msg.ResModel(), msg.ResID(), catchallDomain))
			}
		}
		if msg.Headers() != "" {
			headers := strings.Split(msg.Headers(), "\n")
			for _, header := range headers {
				toks := strings.Split(header, ":")
				if len(toks) != 2 {
					continue
				}
				emailMsg.Headers.Set(strings.TrimSpace(toks[0]), strings.TrimSpace(toks[1]))
			}
		}

		// Writing on the mail object may fail (e.g. lock on user) which
		// would trigger a rollback *after* actually sending the email.
		// To avoid sending twice the same email, provoke the failure earlier
		msg.Write(h.MailMail().NewData().
			SetState("exception").
			SetFailureReason(rs.T("Error without exception. Probably due do sending an email without computed recipients.")))

		// Update notification in a transient exception state to avoid concurrent
		// update in case an email bounces while sending all emails related to current
		// mail record.
		notifs := h.MailNotification().Search(rs.Env(), q.MailNotification().
			IsEmail().Equals(true).And().
			Mail().Equals(msg).And().
			EmailStatus().NotIn([]string{"sent", "canceled"}))
		if notifs.IsNotEmpty() {
			notifMsg := rs.T("Error without exception. Probably due do concurrent access update of notification records. Please see with an administrator.")
			notifs.Sudo().Write(h.MailNotification().NewData().
				SetEmailStatus("exception").
				SetFailureType("UNKNOWN").
				SetFailureReason(notifMsg))
		}

		// build an RFC2822 email.Email object and send it without queuing
		emailMsg.From = msg.EmailFrom()
		emailMsg.To = emailutils.SplitAddresses(msg.EmailTo())
		emailMsg.Cc = emailutils.SplitAddresses(msg.EmailCc())
		emailMsg.ReplyTo = emailutils.SplitAddresses(msg.ReplyTo())
		emailMsg.Subject = msg.Subject()

		// TODO
		// h.MailServer().NewSet(rs.Env()).BuildEmail(emailMsg, )

	}
	//        IrMailServer = self.env['ir.mail_server']
	//        IrAttachment = self.env['ir.attachment']
	//        for mail_id in self.ids:
	//            success_pids = []
	//            failure_type = None
	//            processing_pid = None
	//            mail = None
	//            try:
	//                mail = self.browse(mail_id)
	//                if mail.state != 'outgoing':
	//                    if mail.state != 'exception' and mail.auto_delete:
	//                        mail.sudo().unlink()
	//                    continue
	//
	//                # remove attachments if user send the link with the access_token
	//                body = mail.body_html or ''
	//                attachments = mail.attachment_ids
	//                for link in re.findall(r'/web/(?:content|image)/([0-9]+)', body):
	//                    attachments = attachments - IrAttachment.browse(int(link))
	//
	//                # load attachment binary data with a separate read(), as prefetching all
	//                # `datas` (binary field) could bloat the browse cache, triggerring
	//                # soft/hard mem limits with temporary data.
	//                attachments = [(a['datas_fname'], base64.b64decode(a['datas']), a['mimetype'])
	//                               for a in attachments.sudo().read(['datas_fname', 'datas', 'mimetype']) if a['datas'] is not False]
	//
	//                # specific behavior to customize the send email for notified partners
	//                email_list = []
	//                if mail.email_to:
	//                    email_list.append(mail._send_prepare_values())
	//                for partner in mail.recipient_ids:
	//                    values = mail._send_prepare_values(partner=partner)
	//                    values['partner_id'] = partner
	//                    email_list.append(values)
	//
	//                # headers
	//                headers = {}
	//                ICP = self.env['ir.config_parameter'].sudo()
	//                bounce_alias = ICP.get_param("mail.bounce.alias")
	//                catchall_domain = ICP.get_param("mail.catchall.domain")
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
	//                # Update notification in a transient exception state to avoid concurrent
	//                # update in case an email bounces while sending all emails related to current
	//                # mail record.
	//                notifs = self.env['mail.notification'].search([
	//                    ('is_email', '=', True),
	//                    ('mail_id', 'in', mail.ids),
	//                    ('email_status', 'not in', ('sent', 'canceled'))
	//                ])
	//                if notifs:
	//                    notif_msg = _(
	//                        'Error without exception. Probably due do concurrent access update of notification records. Please see with an administrator.')
	//                    notifs.sudo().write({
	//                        'email_status': 'exception',
	//                        'failure_type': 'UNKNOWN',
	//                        'failure_reason': notif_msg,
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
	//                    processing_pid = email.pop("partner_id", None)
	//                    try:
	//                        res = IrMailServer.send_email(
	//                            msg, mail_server_id=mail.mail_server_id.id, smtp_session=smtp_session)
	//                        if processing_pid:
	//                            success_pids.append(processing_pid)
	//                        processing_pid = None
	//                    except AssertionError as error:
	//                        if str(error) == IrMailServer.NO_VALID_RECIPIENT:
	//                            failure_type = "RECIPIENT"
	//                            # No valid recipient found for this particular
	//                            # mail item -> ignore error to avoid blocking
	//                            # delivery to next recipients, if any. If this is
	//                            # the only recipient, the mail will show as failed.
	//                            _logger.info("Ignoring invalid recipients for mail.mail %s: %s",
	//                                         mail.message_id, email.get('email_to'))
	//                        else:
	//                            raise
	//                if res:  # mail has been sent at least once, no major exception occured
	//                    mail.write({'state': 'sent', 'message_id': res,
	//                                'failure_reason': False})
	//                    _logger.info(
	//                        'Mail with ID %r and Message-Id %r successfully sent', mail.id, mail.message_id)
	//                    # /!\ can't use mail.state here, as mail.refresh() will cause an error
	//                    # see revid:odo@openerp.com-20120622152536-42b2s28lvdv3odyr in 6.1
	//                mail._postprocess_sent_message(
	//                    success_pids=success_pids, failure_type=failure_type)
	//            except MemoryError:
	//                # prevent catching transient MemoryErrors, bubble up to notify user or abort cron job
	//                # instead of marking the mail as failed
	//                _logger.exception(
	//                    'MemoryError while processing mail with ID %r and Msg-Id %r. Consider raising the --limit-memory-hard startup option',
	//                    mail.id, mail.message_id)
	//                # mail status will stay on ongoing since transaction will be rollback
	//                raise
	//            except (psycopg2.Error, smtplib.SMTPServerDisconnected):
	//                # If an error with the database or SMTP session occurs, chances are that the cursor
	//                # or SMTP session are unusable, causing further errors when trying to save the state.
	//                _logger.exception(
	//                    'Exception while processing mail with ID %r and Msg-Id %r.',
	//                    mail.id, mail.message_id)
	//                raise
	//            except Exception as e:
	//                failure_reason = tools.ustr(e)
	//                _logger.exception(
	//                    'failed sending mail (id: %s) due to %s', mail.id, failure_reason)
	//                mail.write(
	//                    {'state': 'exception', 'failure_reason': failure_reason})
	//                mail._postprocess_sent_message(
	//                    success_pids=success_pids, failure_reason=failure_reason, failure_type='UNKNOWN')
	//                if raise_exception:
	//                    if isinstance(e, (AssertionError, UnicodeEncodeError)):
	//                        if isinstance(e, UnicodeEncodeError):
	//                            value = "Invalid text: %s" % e.object
	//                        else:
	//                            # get the args of the original error, wrap into a value and throw a MailDeliveryException
	//                            # that is an except_orm, with name and value as arguments
	//                            value = '. '.join(e.args)
	//                        raise MailDeliveryException(
	//                            _("Mail Delivery Failed"), value)
	//                    raise
	//
	//            if auto_commit is True:
	//                self._cr.commit()
	//        return True
	return true
}

func mailMail_NameGet(rs m.MailMailSet) string {
	return rs.Subject()
}

func init() {
	models.NewModel("MailMail")
	h.MailMail().AddFields(fields_MailMail)
	h.MailMail().SetDescription("Outgoing Mails")
	h.MailMail().SetDefaultOrder("ID DESC")
	h.MailMail().Methods().Create().Extend(mailMail_Create)
	h.MailMail().Methods().Write().Extend(mailMail_Write)
	h.MailMail().Methods().Unlink().Extend(mailMail_Unlink)
	h.MailMail().Methods().DefaultGet().Extend(mailMail_DefaultGet)
	h.MailMail().NewMethod("MarkOutgoing", mailMail_MarkOutgoing)
	h.MailMail().NewMethod("Cancel", mailMail_Cancel)
	h.MailMail().NewMethod("ProcessEmailQueue", mailMail_ProcessEmailQueue)
	h.MailMail().NewMethod("PostprocessSentMessage", mailMail_PostprocessSentMessage)
	h.MailMail().NewMethod("SendPrepareBody", mailMail_SendPrepareBody)
	h.MailMail().NewMethod("SendPrepareValues", mailMail_SendPrepareValues)
	h.MailMail().NewMethod("SplitByServer", mailMail_SplitByServer)
	h.MailMail().NewMethod("Send", mailMail_Send)
	h.MailMail().NewMethod("DoSend", mailMail_DoSend)
	h.MailMail().Methods().NameGet().Extend(mailMail_NameGet)
}
