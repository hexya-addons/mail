package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

var fields_MailResendMessage = map[string]models.FieldDefinition{
	"MailMessageId": fields.Many2One{
		RelationModel: h.MailMessage(),
		String:        "Message",
		ReadOnly:      true},

	"PartnerIds": fields.One2Many{
		RelationModel: h.MailResendPartner(),
		ReverseFK:     "",
		String:        "Recipients"},

	"NotificationIds": fields.Many2Many{
		RelationModel: h.MailNotification(),
		String:        "Notifications",
		ReadOnly:      true},

	"HasCancel": fields.Boolean{
		Compute: h.MailResendMessage().Methods().ComputeHasCancel()},

	"PartnerReadonly": fields.Boolean{
		Compute: h.MailResendMessage().Methods().ComputePartnerReadonly()},
}

// ComputeHasCancel
func mailResendMessage_ComputeHasCancel(rs m.MailResendMessageSet) m.MailResendMessageData {
	//        self.has_cancel = self.partner_ids.filtered(lambda p: not p.resend)
}

// ComputePartnerReadonly
func mailResendMessage_ComputePartnerReadonly(rs m.MailResendMessageSet) m.MailResendMessageData {
	//        self.partner_readonly = not self.env['res.partner'].check_access_rights(
	//            'write', raise_exception=False)
}

// DefaultGet
func mailResendMessage_DefaultGet(rs m.MailResendMessageSet, fields interface{}) {
	//        rec = super(MailResendMessage, self).default_get(fields)
	//        message_id = self._context.get('mail_message_to_resend')
	//        if message_id:
	//            mail_message_id = self.env['mail.message'].browse(message_id)
	//            notification_ids = mail_message_id.notification_ids.filtered(
	//                lambda notif: notif.email_status in ('exception', 'bounce'))
	//            partner_ids = [(0, 0,
	//                            {
	//                                "partner_id": notif.res_partner_id.id,
	//                                "name": notif.res_partner_id.name,
	//                                "email": notif.res_partner_id.email,
	//                                "resend": True,
	//                                "message": notif.format_failure_reason(),
	//                            }
	//                            ) for notif in notification_ids]
	//            has_user = any(
	//                [notif.res_partner_id.user_ids for notif in notification_ids])
	//            if has_user:
	//                partner_readonly = not self.env['res.users'].check_access_rights(
	//                    'write', raise_exception=False)
	//            else:
	//                partner_readonly = not self.env['res.partner'].check_access_rights(
	//                    'write', raise_exception=False)
	//            rec['partner_readonly'] = partner_readonly
	//            rec['notification_ids'] = [(6, 0, notification_ids.ids)]
	//            rec['mail_message_id'] = mail_message_id.id
	//            rec['partner_ids'] = partner_ids
	//        else:
	//            raise UserError(_('No message_id found in context'))
	//        return rec
}

//  Process the wizard content and proceed with sending the related
//             email(s), rendering any template patterns on
// the fly if needed.
func mailResendMessage_ResendMailAction(rs m.MailResendMessageSet) {
	//        for wizard in self:
	//            "If a partner disappeared from partner list, we cancel the notification"
	//            to_cancel = wizard.partner_ids.filtered(
	//                lambda p: not p.resend).mapped("partner_id")
	//            to_send = wizard.partner_ids.filtered(
	//                lambda p: p.resend).mapped("partner_id")
	//            notif_to_cancel = wizard.notification_ids.filtered(
	//                lambda notif: notif.res_partner_id in to_cancel and notif.email_status in ('exception', 'bounce'))
	//            notif_to_cancel.sudo().write({'email_status': 'canceled'})
	//            if to_send:
	//                message = wizard.mail_message_id
	//                record = self.env[message.model].browse(
	//                    message.res_id) if message.model and message.res_id else None
	//
	//                rdata = []
	//                for pid, cid, active, pshare, ctype, notif, groups in self.env['mail.followers']._get_recipient_data(None, False, pids=to_send.ids):
	//                    if pid and notif == 'email' or not notif:
	//                        pdata = {'id': pid, 'share': pshare, 'active': active,
	//                                 'notif': 'email', 'groups': groups or []}
	//                        if not pshare and notif:  # has an user and is not shared, is therefore user
	//                            rdata.append(dict(pdata, type='user'))
	//                        elif pshare and notif:  # has an user and is shared, is therefore portal
	//                            rdata.append(dict(pdata, type='portal'))
	//                        else:  # has no user, is therefore customer
	//                            rdata.append(dict(pdata, type='customer'))
	//
	//                self.env['res.partner']._notify(
	//                    message,
	//                    rdata,
	//                    record,
	//                    force_send=True,
	//                    send_after_commit=False
	//                )
	//
	//            self.mail_message_id._notify_failure_update()
	//        return {'type': 'ir.actions.act_window_close'}
}

// CancelMailAction
func mailResendMessage_CancelMailAction(rs m.MailResendMessageSet) {
	//        for wizard in self:
	//            for notif in wizard.notification_ids:
	//                notif.filtered(lambda notif: notif.email_status in (
	//                    'exception', 'bounce')).sudo().write({'email_status': 'canceled'})
	//            wizard.mail_message_id._notify_failure_update()
	//        return {'type': 'ir.actions.act_window_close'}
}

var fields_MailResendPartner = map[string]models.FieldDefinition{
	"PartnerId": fields.Many2One{
		RelationModel: h.Partner(),
		String:        "Partner",
		Required:      true,
		OnDelete:      `cascade`},

	"Name": fields.Char{
		Related: `PartnerId.Name`,
		//related_sudo=False
		ReadOnly: false},

	"Email": fields.Char{
		Related: `PartnerId.Email`,
		//related_sudo=False
		ReadOnly: false},

	"Resend": fields.Boolean{
		String:  "Send Again",
		Default: models.DefaultValue(true)},

	"ResendWizardId": fields.Many2One{
		RelationModel: h.MailResendMessage(),
		String:        "Resend wizard"},

	"Message": fields.Char{
		String: "Help message"},
}

var fields_MailResendCancel = map[string]models.FieldDefinition{
	"Model": fields.Char{
		String: "Model"},

	"HelpMessage": fields.Char{
		String:  "Help message",
		Compute: h.MailResendCancel().Methods().ComputeHelpMessage()},
}

// ComputeHelpMessage
func mailResendCancel_ComputeHelpMessage(rs m.MailResendCancelSet) m.MailResendCancelData {
	//        for wizard in self:
	//            wizard.help_message = _(
	//                "Are you sure you want to discard %s mail delivery failures. You won't be able to re-send these mails later!") % (wizard._context.get('unread_counter'))
}

// CancelResendAction
func mailResendCancel_CancelResendAction(rs m.MailResendCancelSet) {
	//        author_id = self.env.user.partner_id.id
	//        for wizard in self:
	//            self._cr.execute("""
	//                                SELECT notif.id, mes.id
	//                                FROM mail_message_res_partner_needaction_rel notif
	//                                JOIN mail_message mes
	//                                    ON notif.mail_message_id = mes.id
	//                                WHERE notif.email_status IN ('bounce', 'exception')
	//                                    AND mes.model = %s
	//                                    AND mes.author_id = %s
	//                            """, (wizard.model, author_id))
	//            res = self._cr.fetchall()
	//            notif_ids = [row[0] for row in res]
	//            messages_ids = list(set([row[1] for row in res]))
	//            if notif_ids:
	//                self.env["mail.notification"].browse(
	//                    notif_ids).sudo().write({'email_status': 'canceled'})
	//                self.env["mail.message"].browse(
	//                    messages_ids)._notify_failure_update()
	//        return {'type': 'ir.actions.act_window_close'}
}
func init() {
	models.NewModel("MailResendMessage")
	h.MailResendMessage().AddFields(fields_MailResendMessage)
	h.MailResendMessage().NewMethod("ComputeHasCancel", mailResendMessage_ComputeHasCancel)
	h.MailResendMessage().NewMethod("ComputePartnerReadonly", mailResendMessage_ComputePartnerReadonly)
	h.MailResendMessage().Methods().DefaultGet().Extend(mailResendMessage_DefaultGet)
	h.MailResendMessage().NewMethod("ResendMailAction", mailResendMessage_ResendMailAction)
	h.MailResendMessage().NewMethod("CancelMailAction", mailResendMessage_CancelMailAction)

	models.NewModel("MailResendPartner")
	h.MailResendPartner().AddFields(fields_MailResendPartner)

	models.NewModel("MailResendCancel")
	h.MailResendCancel().AddFields(fields_MailResendCancel)
	h.MailResendCancel().NewMethod("ComputeHelpMessage", mailResendCancel_ComputeHelpMessage)
	h.MailResendCancel().NewMethod("CancelResendAction", mailResendCancel_CancelResendAction)

}
