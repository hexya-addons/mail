package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.IrActionsServer().DeclareModel()

	h.IrActionsServer().Methods().GetStates().DeclareMethod(
		`GetStates`,
		func(rs m.IrActionsServerSet) {
			//        res = super(ServerActions, self)._get_states()
			//        res.insert(0, ('email', 'Send Email'))
			//        return res
		})
	h.IrActionsServer().AddFields(map[string]models.FieldDefinition{
		"EmailFrom": models.CharField{
			String:   "From",
			Related:  `TemplateId.EmailFrom`,
			ReadOnly: true,
		},
		"EmailTo": models.CharField{
			String:   "To (Emails)",
			Related:  `TemplateId.EmailTo`,
			ReadOnly: true,
		},
		"PartnerTo": models.CharField{
			String:   "To (Partners)",
			Related:  `TemplateId.PartnerTo`,
			ReadOnly: true,
		},
		"Subject": models.CharField{
			String:   "Subject",
			Related:  `TemplateId.Subject`,
			ReadOnly: true,
		},
		"BodyHtml": models.HTMLField{
			String:   "Body",
			Related:  `TemplateId.BodyHtml`,
			ReadOnly: true,
		},
		"TemplateId": models.Many2OneField{
			RelationModel: h.MailTemplate(),
			String:        "Email Template",
			OnDelete:      `set null`,
			Filter:        q.ModelId().Equals(model_id),
		},
	})
	h.IrActionsServer().Methods().OnChangeTemplateId().DeclareMethod(
		` Render the raw template in the server action fields. `,
		func(rs m.IrActionsServerSet) {
			//        if self.template_id and not self.template_id.email_from:
			//            raise UserError(_('Your template should define email_from'))
		})
	h.IrActionsServer().Methods().RunActionEmail().DeclareMethod(
		`RunActionEmail`,
		func(rs m.IrActionsServerSet, action interface{}, eval_context interface{}) {
			//        if not action.template_id or not self._context.get('active_id'):
			//            return False
			//        action.template_id.send_mail(self._context.get(
			//            'active_id'), force_send=False, raise_exception=False)
			//        return False
		})
	h.IrActionsServer().Methods().GetEvalContext().DeclareMethod(
		` Override the method giving the evaluation context but also the
        context used in all subsequent calls. Add the mail_notify_force_send
        key set to False in the context. This way all notification
emails linked
        to the currently executed action will be set in
the queue instead of
        sent directly. This will avoid possible break in transactions. `,
		func(rs m.IrActionsServerSet, action interface{}) {
			//        eval_context = super(
			//            ServerActions, self)._get_eval_context(action=action)
			//        ctx = dict(eval_context.get('context', {}))
			//        ctx['mail_notify_force_send'] = False
			//        eval_context['context'] = ctx
			//        return eval_context
		})
}
