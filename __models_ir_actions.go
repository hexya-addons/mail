package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

var fields_IrActionsServer = map[string]models.FieldDefinition{
	"State": fields.Selection{
		// selection_add=[('email', 'Send Email'),('followers', 'Add Followers'),('next_activity', 'Create Next Activity'),]
	},

	"PartnerIds": fields.Many2Many{
		RelationModel: h.Partner(),
		String:        "Add Followers"},

	"ChannelIds": fields.Many2Many{
		RelationModel: h.MailChannel(),
		String:        "Add Channels"},

	"TemplateId": fields.Many2One{
		RelationModel: h.MailTemplate(),
		String:        "Email Template",
		OnDelete:      `set null`,
		Filter:        q.ModelId().Equals(model_id)},

	"ActivityTypeId": fields.Many2One{
		RelationModel: h.MailActivityType(),
		String:        "Activity",
		Filter:        q.ResModelId().Equals(False).Or().ResModelId().Equals(model_id)},

	"ActivitySummary": fields.Char{
		String: "Summary"},

	"ActivityNote": fields.HTML{
		String: "Note"},

	"ActivityDateDeadlineRange": fields.Integer{
		String: "Due Date In"},

	"ActivityDateDeadlineRangeType": fields.Selection{
		Selection: types.Selection{
			"days":   "Days",
			"weeks":  "Weeks",
			"months": "Months",
		},
		String:  "Due type",
		Default: models.DefaultValue("days")},

	"ActivityUserType": fields.Selection{
		Selection: types.Selection{
			"specific": "Specific User",
			"generic":  "Generic User From Record",
		},
		Default:  models.DefaultValue("specific"),
		Required: true,
		Help: "Use 'Specific User' to always assign the same user on the" +
			"next activity. Use 'Generic User From Record' to specify" +
			"the field name of the user to choose on the record."},

	"ActivityUserId": fields.Many2One{
		RelationModel: h.User(),
		String:        "Responsible"},

	"ActivityUserFieldName": fields.Char{
		String:  "User field name",
		Help:    "Technical name of the user on the record",
		Default: models.DefaultValue("user_id")},
}

// OnchangeActivityDateDeadlineRange
func irActionsServer_OnchangeActivityDateDeadlineRange(rs m.IrActionsServerSet) {
	//        if self.activity_date_deadline_range < 0:
	//            raise UserError(_("The 'Due Date In' value can't be negative."))
}

// OnChangeTemplateId
func irActionsServer_OnChangeTemplateId(rs m.IrActionsServerSet) {
	//        pass
}

// CheckMailThread
func irActionsServer_CheckMailThread(rs m.IrActionsServerSet) {
	//        for action in self:
	//            if action.state == 'followers' and not action.model_id.is_mail_thread:
	//                raise ValidationError(
	//                    _("Add Followers can only be done on a mail thread model"))
}

// CheckActivityMixin
func irActionsServer_CheckActivityMixin(rs m.IrActionsServerSet) {
	//        for action in self:
	//            if action.state == 'next_activity' and not issubclass(self.pool[action.model_id.model], self.pool['mail.thread']):
	//                raise ValidationError(
	//                    _("A next activity can only be planned on models that use the chatter"))
}

// RunActionFollowersMulti
func irActionsServer_RunActionFollowersMulti(rs m.IrActionsServerSet, action interface{}, eval_context interface{}) {
	//        Model = self.env[action.model_name]
	//        if self.partner_ids or self.channel_ids and hasattr(Model, 'message_subscribe'):
	//            records = Model.browse(self._context.get(
	//                'active_ids', self._context.get('active_id')))
	//            records.message_subscribe(
	//                self.partner_ids.ids, self.channel_ids.ids)
	//        return False
}

// When an activity is set on update of a record,
//         update might be triggered many times by recomputes.
//         When need to know it to skip these steps.
//         Except if the computed field is supposed to trigger the action
//         
func irActionsServer_IsRecompute(rs m.IrActionsServerSet, action interface{}) {
	//        records = self.env[action.model_name].browse(
	//            self._context.get('active_ids', self._context.get('active_id')))
	//        old_values = action._context.get('old_values')
	//        if old_values:
	//            domain_post = action._context.get('domain_post')
	//            tracked_fields = []
	//            if domain_post:
	//                for leaf in domain_post:
	//                    if isinstance(leaf, (tuple, list)):
	//                        tracked_fields.append(leaf[0])
	//            fields_to_check = [field for record, field_names in old_values.items(
	//            ) for field in field_names if field not in tracked_fields]
	//            if fields_to_check:
	//                field = records._fields[fields_to_check[0]]
	//                # Pick an arbitrary field; if it is marked to be recomputed,
	//                # it means we are in an extraneous write triggered by the recompute.
	//                # In this case, we should not create a new activity.
	//                if records._recompute_check(field):
	//                    return True
	//        return False
}

// RunActionEmail
func irActionsServer_RunActionEmail(rs m.IrActionsServerSet, action interface{}, eval_context interface{}) {
	//        if not action.template_id or not self._context.get('active_id') or self._is_recompute(action):
	//            return False
	//        cleaned_ctx = dict(self.env.context)
	//        cleaned_ctx.pop('default_type', None)
	//        cleaned_ctx.pop('default_parent_id', None)
	//        action.template_id.with_context(cleaned_ctx).send_mail(
	//            self._context.get('active_id'), force_send=False, raise_exception=False)
	//        return False
}

// RunActionNextActivity
func irActionsServer_RunActionNextActivity(rs m.IrActionsServerSet, action interface{}, eval_context interface{}) {
	//        if not action.activity_type_id or not self._context.get('active_id') or self._is_recompute(action):
	//            return False
	//        records = self.env[action.model_name].browse(
	//            self._context.get('active_ids', self._context.get('active_id')))
	//        vals = {
	//            'summary': action.activity_summary or '',
	//            'note': action.activity_note or '',
	//            'activity_type_id': action.activity_type_id.id,
	//        }
	//        if action.activity_date_deadline_range > 0:
	//            vals['date_deadline'] = fields.Date.context_today(action) + relativedelta(
	//                **{action.activity_date_deadline_range_type: action.activity_date_deadline_range})
	//        for record in records:
	//            user = False
	//            if action.activity_user_type == 'specific':
	//                user = action.activity_user_id
	//            elif action.activity_user_type == 'generic' and action.activity_user_field_name in record:
	//                user = record[action.activity_user_field_name]
	//            if user:
	//                vals['user_id'] = user.id
	//            record.activity_schedule(**vals)
	//        return False
}

//  Override the method giving the evaluation context but also the
//         context used in all subsequent calls. Add the mail_notify_force_send
//         key set to False in the context. This way all notification
// emails linked
//         to the currently executed action will be set in
// the queue instead of
//         sent directly. This will avoid possible break in transactions. 
func irActionsServer_GetEvalContext(rs m.IrActionsServerSet, action interface{}) {
	//        eval_context = super(
	//            ServerActions, self)._get_eval_context(action=action)
	//        ctx = dict(eval_context['env'].context)
	//        ctx['mail_notify_force_send'] = False
	//        eval_context['env'].context = ctx
	//        return eval_context
}
func init() {
	models.NewModel("IrActionsServer")
	h.IrActionsServer().AddFields(fields_IrActionsServer)
	h.IrActionsServer().NewMethod("OnchangeActivityDateDeadlineRange", irActionsServer_OnchangeActivityDateDeadlineRange)
	h.IrActionsServer().NewMethod("OnChangeTemplateId", irActionsServer_OnChangeTemplateId)
	h.IrActionsServer().NewMethod("CheckMailThread", irActionsServer_CheckMailThread)
	h.IrActionsServer().NewMethod("CheckActivityMixin", irActionsServer_CheckActivityMixin)
	h.IrActionsServer().NewMethod("RunActionFollowersMulti", irActionsServer_RunActionFollowersMulti)
	h.IrActionsServer().NewMethod("IsRecompute", irActionsServer_IsRecompute)
	h.IrActionsServer().NewMethod("RunActionEmail", irActionsServer_RunActionEmail)
	h.IrActionsServer().NewMethod("RunActionNextActivity", irActionsServer_RunActionNextActivity)
	h.IrActionsServer().NewMethod("GetEvalContext", irActionsServer_GetEvalContext)

}
