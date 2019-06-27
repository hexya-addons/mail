package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.MailWizardInvite().DeclareModel()

	h.MailWizardInvite().Methods().DefaultGet().Extend(
		`DefaultGet`,
		func(rs m.MailWizardInviteSet, fields interface{}) {
			//        result = super(Invite, self).default_get(fields)
			//        if self._context.get('mail_invite_follower_channel_only'):
			//            result['send_mail'] = False
			//        if 'message' not in fields:
			//            return result
			//        user_name = self.env.user.name_get()[0][1]
			//        model = result.get('res_model')
			//        res_id = result.get('res_id')
			//        if model and res_id:
			//            document = self.env['ir.model'].search(
			//                [('model', '=', model)]).name_get()[0][1]
			//            title = self.env[model].browse(res_id).display_name
			//            msg_fmt = _(
			//                '%(user_name)s invited you to follow %(document)s document: %(title)s')
			//        else:
			//            msg_fmt = _('%(user_name)s invited you to follow a new document')
			//        text = msg_fmt % locals()
			//        message = html.DIV(
			//            html.P(_('Hello,')),
			//            html.P(text)
			//        )
			//        result['message'] = etree.tostring(message)
			//        return result
		})
	h.MailWizardInvite().AddFields(map[string]models.FieldDefinition{
		"ResModel": models.CharField{
			String:   "Related Document Model",
			Required: true,
			Index:    true,
			Help:     "Model of the followed resource",
		},
		"ResId": models.IntegerField{
			String: "Related Document ID",
			Index:  true,
			Help:   "Id of the followed resource",
		},
		"PartnerIds": models.Many2ManyField{
			RelationModel: h.Partner(),
			String:        "Recipients",
			Help:          "List of partners that will be added as follower of the current document.",
		},
		"ChannelIds": models.Many2ManyField{
			RelationModel: h.MailChannel(),
			String:        "Channels",
			Help: "List of channels that will be added as listeners of the" +
				"current document.",
			Filter: q.ChannelType().Equals("channel"),
		},
		"Message": models.HTMLField{
			String: "Message",
		},
		"SendMail": models.BooleanField{
			String:  "Send Email",
			Default: models.DefaultValue(true),
			Help: "If checked, the partners will receive an email warning" +
				"they have been added in the document's followers.",
		},
	})
	h.MailWizardInvite().Methods().AddFollowers().DeclareMethod(
		`AddFollowers`,
		func(rs m.MailWizardInviteSet) {
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
			//            model_ids = self.env['ir.model'].search(
			//                [('model', '=', wizard.res_model)])
			//            model_name = model_ids.name_get()[0][1]
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
			//                })
			//                new_partners.with_context(auto_delete=True)._notify(
			//                    message, force_send=True, send_after_commit=False, user_signature=True)
			//                message.unlink()
			//        return {'type': 'ir.actions.act_window_close'}
		})
}
