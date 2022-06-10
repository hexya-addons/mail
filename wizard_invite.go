package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

var fields_MailWizardInvite = map[string]models.FieldDefinition{
	"ResModel": fields.Char{
		String:   "Related Document Model",
		Required: true,
		Index:    true,
		Help:     "Model of the followed resource"},

	"ResId": fields.Integer{
		String: "Related Document ID",
		Index:  true,
		Help:   "Id of the followed resource"},

	"PartnerIds": fields.Many2Many{
		RelationModel: h.Partner(),
		String:        "Recipients",
		Help:          "List of partners that will be added as follower of the current document."},

	"ChannelIds": fields.Many2Many{
		RelationModel: h.MailChannel(),
		String:        "Channels",
		Help: "List of channels that will be added as listeners of the" +
			"current document.",
		Filter: q.ChannelType().Equals("channel")},

	"Message": fields.HTML{
		String: "Message"},

	"SendMail": fields.Boolean{
		String:  "Send Email",
		Default: models.DefaultValue(true),
		Help: "If checked, the partners will receive an email warning" +
			"they have been added in the document's followers."},
}

// DefaultGet
func mailWizardInvite_DefaultGet(rs m.MailWizardInviteSet, fields interface{}) {
	//        result = super(Invite, self).default_get(fields)
	//        if self._context.get('mail_invite_follower_channel_only'):
	//            result['send_mail'] = False
	//        if 'message' not in fields:
	//            return result
	//        user_name = self.env.user.name_get()[0][1]
	//        model = result.get('res_model')
	//        res_id = result.get('res_id')
	//        if model and res_id:
	//            document = self.env['ir.model']._get(model).display_name
	//            title = self.env[model].browse(res_id).display_name
	//            msg_fmt = _(
	//                '%(user_name)s invited you to follow %(document)s document: %(title)s')
	//        else:
	//            msg_fmt = _('%(user_name)s invited you to follow a new document.')
	//        text = msg_fmt % locals()
	//        message = html.DIV(
	//            html.P(_('Hello,')),
	//            html.P(text)
	//        )
	//        result['message'] = etree.tostring(message)
	//        return result
}

// AddFollowers
func mailWizardInvite_AddFollowers(rs m.MailWizardInviteSet) {
	//        email_from = self.env['mail.message']._get_default_from()
	//        for wizard in self:
	//            Model = self.env[wizard.res_model]
	//            document = Model.browse(wizard.res_id)
	//
	//            # filter partner_ids to get the new followers, to avoid sending email to already following partners
	//            new_partners = wizard.partner_ids - document.message_partner_ids
	//            new_channels = wizard.channel_ids - document.message_channel_ids
	//            document.message_subscribe(new_partners.ids, new_channels.ids)
	//
	//            model_name = self.env['ir.model']._get(
	//                wizard.res_model).display_name
	//            # send an email if option checked and if a message exists (do not send void emails)
	//            # when deleting the message, cleditor keeps a <br>
	//            if wizard.send_mail and wizard.message and not wizard.message == '<br>':
	//                message = self.env['mail.message'].create({
	//                    'subject': _('Invitation to follow %s: %s') % (model_name, document.name_get()[0][1]),
	//                    'body': wizard.message,
	//                    'record_name': document.name_get()[0][1],
	//                    'email_from': email_from,
	//                    'reply_to': email_from,
	//                    'model': wizard.res_model,
	//                    'res_id': wizard.res_id,
	//                    'no_auto_thread': True,
	//                    'add_sign': True,
	//                })
	//                partners_data = []
	//                recipient_data = self.env['mail.followers']._get_recipient_data(
	//                    document, False, pids=new_partners.ids)
	//                for pid, cid, active, pshare, ctype, notif, groups in recipient_data:
	//                    pdata = {'id': pid, 'share': pshare, 'active': active,
	//                             'notif': 'email', 'groups': groups or []}
	//                    if not pshare and notif:  # has an user and is not shared, is therefore user
	//                        partners_data.append(dict(pdata, type='user'))
	//                    elif pshare and notif:  # has an user and is shared, is therefore portal
	//                        partners_data.append(dict(pdata, type='portal'))
	//                    else:  # has no user, is therefore customer
	//                        partners_data.append(dict(pdata, type='customer'))
	//                self.env['res.partner'].with_context(auto_delete=True)._notify(
	//                    message, partners_data, document,
	//                    force_send=True, send_after_commit=False)
	//                message.unlink()
	//        return {'type': 'ir.actions.act_window_close'}
}
func init() {
	models.NewModel("MailWizardInvite")
	h.MailWizardInvite().AddFields(fields_MailWizardInvite)
	h.MailWizardInvite().Methods().DefaultGet().Extend(mailWizardInvite_DefaultGet)
	h.MailWizardInvite().NewMethod("AddFollowers", mailWizardInvite_AddFollowers)

}
