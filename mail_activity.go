package mail

import (
	"fmt"
	"time"

	"github.com/hexya-erp/hexya/src/actions"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/models/types/dates"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

var fields_MailActivityType = map[string]models.FieldDefinition{
	"Name": fields.Char{
		String:    "Name",
		Required:  true,
		Translate: true},
	"Summary": fields.Char{
		String:    "Summary",
		Translate: true},
	"Sequence": fields.Integer{
		String:  "Sequence",
		Default: models.DefaultValue(10)},
	"Active": fields.Boolean{
		Default: models.DefaultValue(true)},
	"DelayCount": fields.Integer{
		String:  "After",
		Default: models.DefaultValue(0),
		GoType:  new(int),
		Help: "Number of days/week/month before executing the action." +
			"It allows to plan the action deadline."},
	"DelayUnit": fields.Selection{
		Selection: types.Selection{
			"days":   "days",
			"weeks":  "weeks",
			"months": "months",
		},
		String:   "Delay units",
		Help:     "Unit of delay",
		Required: true,
		Default:  models.DefaultValue("days")},
	"DelayFrom": fields.Selection{
		Selection: types.Selection{
			"current_date":      "after validation date",
			"previous_activity": "after previous activity deadline",
		},
		String:   "Delay Type",
		Help:     "Type of delay",
		Required: true,
		Default:  models.DefaultValue("previous_activity")},
	"Icon": fields.Char{
		String: "Icon",
		Help:   "Font awesome icon e.g. fa-tasks"},
	"DecorationType": fields.Selection{
		Selection: types.Selection{
			"warning": "Alert",
			"danger":  "Error",
		},
		String: "Decoration Type",
		Help:   "Change the background color of the related activities of this type."},
	"ResModel": fields.Char{
		String:   "Model",
		Index:    true,
		OnChange: h.MailActivityType().Methods().OnchangeResModel(),
		Help: "Specify a model if the activity should be specific to a" +
			"model and not available when managing activities for other models."},
	"DefaultNextType": fields.Many2One{
		RelationModel: h.MailActivityType(),
		String:        "Default Next Activity",
		Filter: q.MailActivity().ResModel().Equals("").Or().ResModel().EqualsFunc(func(r models.RecordSet) string {
			return r.Collection().Wrap().(m.MailActivityTypeSet).ResModel()
		})},
	"ForceNext": fields.Boolean{
		String:  "Auto Schedule Next Activity",
		Default: models.DefaultValue(false)},
	"NextTypes": fields.Many2Many{
		RelationModel:    h.MailActivityType(),
		M2MLinkModelName: "MailActivityRel",
		M2MOurField:      "Activity",
		M2MTheirField:    "Recommended",
		Filter: q.MailActivity().ResModel().Equals("").Or().ResModel().EqualsFunc(func(r models.RecordSet) string {
			return r.Collection().Wrap().(m.MailActivityTypeSet).ResModel()
		}),
		String: "Recommended Next Activities"},
	"PreviousTypes": fields.Many2Many{
		RelationModel:    h.MailActivityType(),
		M2MLinkModelName: "MailActivityRel",
		M2MOurField:      "Recommended",
		M2MTheirField:    "Activity",
		Filter: q.MailActivity().ResModel().Equals("").Or().ResModel().EqualsFunc(func(r models.RecordSet) string {
			return r.Collection().Wrap().(m.MailActivityTypeSet).ResModel()
		}),
		String: "Preceding Activities"},
	"Category": fields.Selection{
		Selection: types.Selection{
			"default": "Other",
		},
		Default: models.DefaultValue("default"),
		String:  "Category",
		Help:    "Categories may trigger specific behavior like opening calendar view"},
	"MailTemplates": fields.Many2Many{
		RelationModel: h.MailTemplate(),
		String:        "Mails templates"},
	"InitialResModel": fields.Char{
		String:  "Initial model",
		Compute: h.MailActivityType().Methods().ComputeInitialResModel(),
		Stored:  false,
		Help: "Technical field to keep trace of the model at the beginning" +
			"of the edition for UX related behaviour"},
	"ResModelChange": fields.Boolean{
		String:  "Model has change",
		Help:    "Technical field for UX related behaviour",
		Default: models.DefaultValue(false),
		Stored:  false},
}

// OnchangeResModel updates the mail templates based on the model
func mailActivityType_OnchangeResModel(rs m.MailActivityTypeSet) m.MailActivityTypeData {
	res := h.MailActivityType().NewData()
	res.SetMailTemplates(rs.MailTemplates().Filtered(func(r m.MailTemplateSet) bool {
		return r.ResModel() == rs.ResModel()
	}))
	res.SetResModelChange(rs.InitialResModel() != "" && rs.InitialResModel() != rs.ResModel())
	return res
}

// ComputeInitialResModel returns the initial model (which is the model)
func mailActivityType_ComputeInitialResModel(rs m.MailActivityTypeSet) m.MailActivityTypeData {
	res := h.MailActivityType().NewData()
	res.SetInitialResModel(rs.ResModel())
	return res
}

var fields_MailActivity = map[string]models.FieldDefinition{
	"ResID": fields.Integer{
		String:   "Related Document ID",
		Index:    true,
		Required: true},
	"ResModel": fields.Char{
		String:   "Document Model",
		Index:    true,
		Required: true},
	"ResName": fields.Char{
		String:   "Document Name",
		Compute:  h.MailActivity().Methods().ComputeResName(),
		Stored:   true,
		Help:     "Display name of the related document.",
		ReadOnly: true},
	"ActivityType": fields.Many2One{
		RelationModel: h.MailActivityType(),
		String:        "Activity",
		Filter: q.MailActivity().ResModel().Equals("").Or().ResModel().EqualsFunc(func(r models.RecordSet) string {
			return r.Collection().Wrap().(m.MailActivitySet).ResModel()
		}),
		OnChange: h.MailActivity().Methods().OnchangeActivityType(),
		OnDelete: `restrict`},
	"ActivityCategory": fields.Selection{
		Related:  `ActivityTypeId.Category`,
		ReadOnly: true},
	"ActivityDecoration": fields.Selection{
		Related:  `ActivityTypeId.DecorationType`,
		ReadOnly: true},
	"Icon": fields.Char{
		String:   "Icon",
		Related:  `ActivityTypeId.Icon`,
		ReadOnly: true},
	"Summary": fields.Char{
		String: "Summary"},
	"Note": fields.HTML{
		String: "Note",
	},
	"Feedback": fields.HTML{
		String: "Feedback"},
	"DateDeadline": fields.Date{
		String:   "Due Date",
		Index:    true,
		Required: true,
		Default:  func(env models.Environment) interface{} { return dates.Today() }},
	"Automated": fields.Boolean{
		String:   "Automated activity",
		ReadOnly: true,
		Help: "Indicates this activity has been created automatically" +
			"and not by any user."},
	"User": fields.Many2One{
		RelationModel: h.User(),
		String:        "Assigned to",
		Default:       func(env models.Environment) interface{} { return env.Uid() },
		Index:         true,
		Required:      true},
	"CreateUser": fields.Many2One{
		RelationModel: h.User(),
		String:        "Creator",
		Default:       func(env models.Environment) interface{} { return env.Uid() },
		Index:         true},
	"State": fields.Selection{
		Selection: types.Selection{
			"overdue": "Overdue",
			"today":   "Today",
			"planned": "Planned",
		},
		String:  "State",
		Compute: h.MailActivity().Methods().ComputeState()},
	"RecommendedActivityType": fields.Many2One{
		RelationModel: h.MailActivityType(),
		String:        "Recommended Activity Type",
		OnChange:      h.MailActivity().Methods().OnchangeRecommendedActivityType()},
	"PreviousActivityType": fields.Many2One{
		RelationModel: h.MailActivityType(),
		String:        "Previous Activity Type",
		OnChange:      h.MailActivity().Methods().OnchangePreviousActivityType(),
		ReadOnly:      true},
	"HasRecommendedActivities": fields.Boolean{
		String:  "Next activities available",
		Compute: h.MailActivity().Methods().ComputeHasRecommendedActivities(),
		Help:    "Technical field for UX purpose"},
	"MailTemplates": fields.Many2Many{
		Related:  `ActivityTypeId.MailTemplateIds`,
		ReadOnly: false},
	"ForceNext": fields.Boolean{
		Related:  `ActivityTypeId.ForceNext`,
		ReadOnly: false},
}

// ComputeHasRecommendedActivities computes if the associated activity type has next activities
func mailActivity_ComputeHasRecommendedActivities(rs m.MailActivitySet) m.MailActivityData {
	res := h.MailActivity().NewData()
	res.SetHasRecommendedActivities(rs.PreviousActivityType().NextTypes().IsNotEmpty())
	return res
}

// OnchangePreviousActivityType updates this activity type to be the previous type's default next type
func mailActivity_OnchangePreviousActivityType(rs m.MailActivitySet) m.MailActivityData {
	res := h.MailActivity().NewData()
	if rs.PreviousActivityType().DefaultNextType().IsNotEmpty() {
		rs.SetActivityType(rs.PreviousActivityType().DefaultNextType())
	}
	return res
}

// ComputeResName computes the name of the related record
func mailActivity_ComputeResName(rs m.MailActivitySet) m.MailActivityData {
	res := h.MailActivity().NewData()
	resName := models.Registry.MustGet(rs.ResModel()).BrowseOne(rs.Env(), rs.ResID()).Call("NameGet").(string)
	res.SetResName(resName)
	return res
}

// ComputeState computes the state of the activity
func mailActivity_ComputeState(rs m.MailActivitySet) m.MailActivityData {
	res := h.MailActivity().NewData()
	if rs.DateDeadline().IsZero() {
		return res
	}
	tz := rs.User().Sudo().TZ()
	res.SetState(rs.ComputeStateFromDate(rs.DateDeadline(), tz))
	return res
}

// ComputeStateFromDate returns a state string from the given dateDeadline in the given timezone
func mailActivity_ComputeStateFromDate(_ m.MailActivitySet, dateDeadline dates.Date, tz string) string {
	today := dates.Today().Time
	if tz != "" {
		loc, err := time.LoadLocation(tz)
		if err == nil {
			y, mo, d := today.In(loc).Date()
			today = time.Date(y, mo, d, 0, 0, 0, 0, time.UTC)
		}
	}
	diff := dateDeadline.Time.Sub(today)
	switch {
	case diff.Hours() < 0:
		return "overdue"
	case diff.Hours() < 24:
		return "today"
	default:
		return "planned"
	}
}

// OnchangeActivityType is triggered when activity type is changed to
// update this activity with values of the selected activity type.
func mailActivity_OnchangeActivityType(rs m.MailActivitySet) m.MailActivityData {
	res := h.MailActivity().NewData()
	if rs.ActivityType().IsNotEmpty() {
		res.SetSummary(rs.ActivityType().Summary())
		contextToday := dates.Today().In(rs.Env().Context().GetString("tz"))
		base := dates.Date{
			Time: time.Date(contextToday.Year(), contextToday.Month(), contextToday.Day(), 0, 0, 0, 0, time.UTC),
		}
		if rs.ActivityType().DelayFrom() == "previous_activity" && rs.Env().Context().HasKey("activity_previous_deadline") {
			base = rs.Env().Context().GetDate("activity_previous_deadline")
		}
		switch rs.ActivityType().DelayUnit() {
		case "days":
			res.SetDateDeadline(base.AddDate(0, 0, rs.ActivityType().DelayCount()))
		case "weeks":
			res.SetDateDeadline(base.AddWeeks(rs.ActivityType().DelayCount()))
		case "months":
			res.SetDateDeadline(base.AddDate(0, rs.ActivityType().DelayCount(), 0))
		}
	}
	return res
}

// OnchangeRecommendedActivityType updates this activity when the recommended activity type is changed.
func mailActivity_OnchangeRecommendedActivityType(rs m.MailActivitySet) m.MailActivityData {
	res := h.MailActivity().NewData()
	if rs.RecommendedActivityType().IsNotEmpty() {
		res.SetActivityType(rs.RecommendedActivityType())
	}
	return res
}

// CheckAccess checks whether the current user is allowed to execute the operation
//
// * create: check write rights on related document;
// * write: rule OR write rights on document;
// * unlink: rule OR write rights on document;
//
func mailActivity_CheckAccess(rs m.MailActivitySet, operation string) {
	// TODO Implement CheckAccess
	//
	//        self.check_access_rights(
	//            operation, raise_exception=True)  # will raise an AccessError
	//        if operation in ('write', 'unlink'):
	//            try:
	//                self.check_access_rule(operation)
	//            except exceptions.AccessError:
	//                pass
	//            else:
	//                return
	//        doc_operation = 'read' if operation == 'read' else 'write'
	//        activity_to_documents = dict()
	//        for activity in self.sudo():
	//            activity_to_documents.setdefault(
	//                activity.res_model, list()).append(activity.res_id)
	//        for model, res_ids in activity_to_documents.items():
	//            self.env[model].check_access_rights(
	//                doc_operation, raise_exception=True)
	//            try:
	//                self.env[model].browse(
	//                    res_ids).check_access_rule(doc_operation)
	//            except exceptions.AccessError:
	//                raise exceptions.AccessError(
	//                    _('The requested operation cannot be completed due to security restrictions. Please contact your system administrator.\n\n(Document type: %s, Operation: %s)') % (
	//                        self._description, operation)
	//                    + ' - ({} {}, {} {})'.format(_('Records:'),
	//                                                 res_ids[:6], _('User:'), self._uid)
	//                )
}

//  Check assigned user (user_id field) has access to the document. Purpose
//         is to allow assigned user to handle their activities.
// For that purpose
//         assigned user should be able to at least read the
// document. We therefore
//         raise an UserError if the assigned user has no
// access to the document.
func mailActivity_CheckAccessAssignation(rs m.MailActivitySet) {
	// TODO implement ChecAccessAssignation
	//
	//        for activity in self:
	//            model = self.env[activity.res_model].sudo(activity.user_id.id)
	//            try:
	//                model.check_access_rights('read')
	//            except exceptions.AccessError:
	//                raise exceptions.UserError(
	//                    _('Assigned user %s has no access to the document and is not able to handle this activity.') %
	//                    activity.user_id.display_name)
	//            else:
	//                try:
	//                    target_user = activity.user_id
	//                    target_record = self.env[activity.res_model].browse(
	//                        activity.res_id)
	//                    if hasattr(target_record, 'company_id') and (
	//                        target_record.company_id != target_user.company_id and (
	//                            len(target_user.sudo().company_ids) > 1)):
	//                        return  # in that case we skip the check, assuming it would fail because of the company
	//                    model.browse(activity.res_id).check_access_rule('read')
	//                except exceptions.AccessError:
	//                    raise exceptions.UserError(
	//                        _('Assigned user %s has no access to the document and is not able to handle this activity.') %
	//                        activity.user_id.display_name)
}

func mailActivity_Create(rs m.MailActivitySet, values m.MailActivityData) m.MailActivitySet {
	// already compute default values to be sure those are computed using the current user
	valuesDefaults := rs.DefaultGet()
	valuesDefaults.MergeWith(values)
	// continue as sudo because activities are somewhat protected
	activity := rs.Super().Sudo().Create(valuesDefaults)
	activityUser := activity.Sudo()
	activityUser.CheckAccess("create")
	partnerID := activityUser.User().Partner().ID()

	// send a notification to assigned user; in case of manually done activity also check
	// target has rights on document otherwise we prevent its creation. Automated activities
	// are checked since they are integrated into business flows that should not crash.
	if !activityUser.User().Equals(h.User().NewSet(rs.Env()).CurrentUser()) {
		if !activityUser.Automated() {
			activityUser.CheckAccessAssignation()
		}
		if !rs.Env().Context().GetBool("mail_activity_quick_update") {
			activityUser.Sudo().ActionNotify()
		}
	}
	record := models.Registry.MustGet(activityUser.ResModel()).BrowseOne(rs.Env(), activityUser.ResID()).Wrap("MailThread").(m.MailThreadSet)
	record.MessageSubscribe(partnerID, nil, nil)
	if activity.DateDeadline().LowerEqual(dates.Today()) {
		h.BusBus().NewSet(rs.Env()).Sendone(fmt.Sprintf("partner_%d", activity.User().Partner().ID()),
			map[string]interface{}{
				"type":             "activity_updated",
				"activity_created": true,
			})
	}
	return activityUser
}

func mailActivity_Write(rs m.MailActivitySet, values m.MailActivityData) bool {
	rs.CheckAccess("write")
	preResponsibles := h.Partner().NewSet(rs.Env())
	if values.User().IsNotEmpty() {
		for _, rec := range rs.Records() {
			preResponsibles = preResponsibles.Union(rec.User().Partner())
		}
	}
	res := rs.Super().Sudo().Write(values)
	if values.User().IsEmpty() {
		return res
	}
	if !values.User().Equals(h.User().NewSet(rs.Env()).CurrentUser()) {
		toCheck := rs.Filtered(func(r m.MailActivitySet) bool {
			return !r.Automated()
		})
		toCheck.CheckAccessAssignation()
		if !rs.Env().Context().GetBool("mail_activity_quick_update") {
			rs.ActionNotify()
		}
	}
	for _, activity := range rs.Records() {
		record := models.Registry.MustGet(activity.ResModel()).BrowseOne(rs.Env(), activity.ResID()).Wrap("MailThread").(m.MailThreadSet)
		record.MessageSubscribe(activity.User().Partner().ID, nil, nil)
		h.BusBus().NewSet(rs.Env()).Sendone(fmt.Sprintf("partner_%d", activity.User().Partner().ID()),
			map[string]interface{}{
				"type":             "activity_updated",
				"activity_created": true,
			})
	}
	for _, activity := range rs.Records() {
		if activity.DateDeadline().Greater(dates.Today()) {
			continue
		}
		for _, partner := range preResponsibles.Records() {
			h.BusBus().NewSet(rs.Env()).Sendone(fmt.Sprintf("partner_%d", partner.ID()),
				map[string]interface{}{
					"type":             "activity_updated",
					"activity_deleted": true,
				})
		}
	}
	return res
}

func mailActivity_Unlink(rs m.MailActivitySet) int64 {
	rs.CheckAccess("unlink")
	for _, activity := range rs.Records() {
		if activity.DateDeadline().LowerEqual(dates.Today()) {
			h.BusBus().NewSet(rs.Env()).Sendone(fmt.Sprintf("partner_%d", activity.User().Partner().ID()),
				map[string]interface{}{
					"type":             "activity_updated",
					"activity_deleted": true,
				})
		}
	}
	return rs.Super().Sudo().Unlink()
}

// ActionNotify notifies the user that she has been assigned the activity
func mailActivity_ActionNotify(rs m.MailActivitySet) {
	//        body_template = self.env.ref('mail.message_activity_assigned')
	//        for activity in self:
	//            model_description = self.env['ir.model']._get(
	//                activity.res_model).display_name
	//            body = body_template.render(
	//                dict(activity=activity, model_description=model_description),
	//                engine='ir.qweb',
	//                minimal_qcontext=True
	//            )
	//            self.env['mail.thread'].message_notify(
	//                partner_ids=activity.user_id.partner_id.ids,
	//                body=body,
	//                subject=_('%s: %s assigned to you') % (
	//                    activity.res_name, activity.summary or activity.activity_type_id.name),
	//                record_name=activity.res_name,
	//                model_description=model_description,
	//                notif_layout='mail.mail_notification_light'
	//            )

	// for _, activity := range rs.Records() {
	// modelDescription := models.Registry.MustGet(activity.ResModel()).Description()
	// var body http.ResponseWriter
	// templates.Registry.Instance("mail_message_activity_assigned", map[string]interface{}{
	//
	// }).Render(body)
	// }
}

// ActionDone is a wrapper without feedback because web button add context as
// parameter, therefore setting context to feedback
func mailActivity_ActionDone(rs m.MailActivitySet) int64 {
	return rs.ActionFeedback("")
}

// ActionFeedback closes the activity with the given feedback
func mailActivity_ActionFeedback(rs m.MailActivitySet, feedback string) int64 {
	message := h.MailMessage().NewSet(rs.Env())
	if feedback != "" {
		rs.SetFeedback(feedback)
	}
	// Search for all attachments linked to the activities we are about to unlink. This way, we
	// can link them to the message posted and prevent their deletion.
	attachments := h.Attachment().NewSet(rs.Env()).Search(q.Attachment().
		ResModel().Equals(rs.ModelName()).
		And().ResID().In(rs.Ids()))
	attachments.Load(q.MailActivity().ID(), q.MailActivity().ResID())
	var activityID int64
	activityAttachments := make(map[int64][]int64)
	for _, attachment := range attachments.Records() {
		activityID = attachment.ResID()
		activityAttachments[activityID] = append(activityAttachments[activityID], attachment.ID())
	}

	for _, activity := range rs.Records() {
		record := models.Registry.MustGet(activity.ResModel()).BrowseOne(rs.Env(), activity.ResID()).Wrap("MailThread").(m.MailThreadSet)
		//            record.message_post_with_view(
		//                'mail.message_activity_done',
		//                values={'activity': activity},
		//                subtype_id=self.env['ir.model.data'].xmlid_to_res_id(
		//                    'mail.mt_activities'),
		//                mail_activity_type_id=activity.activity_type_id.id)
		// TODO finish implementation here
		record.MessagePostWithView("mail.message_activity_done")

		// Moving the attachments in the message
		// directly, see route /web_editor/attachment/add
		activityMessage := record.Messages().Records()[0]
		messageAttachments := h.Attachment().Browse(rs.Env(), activityAttachments[activity.ID()])
		if messageAttachments.IsNotEmpty() {
			messageAttachments.Write(h.Attachment().NewData().
				SetResID(activityMessage.ID()).
				SetResModel(activityMessage.ModelName()))
			activityMessage.SetAttachmentIds(messageAttachments)
		}
		message = message.Union(activityMessage)
	}
	rs.Unlink()
	return message.ID()
}

// ActionDoneScheduleNext is wrapper without feedback because web button add context as
// parameter, therefore setting context to feedback
func mailActivity_ActionDoneScheduleNext(rs m.MailActivitySet) *actions.Action {
	return rs.ActionFeedbackScheduleNext("")
}

// ActionFeedbackScheduleNext closes this action with the given feedback
// and returns an UI action to schedule the next activity.
func mailActivity_ActionFeedbackScheduleNext(rs m.MailActivitySet, feedback string) *actions.Action {
	ctx := rs.Env().Context().Cleaned("default_").
		WithKey("default_previous_activity_type_id", rs.ActivityType().ID()).
		WithKey("activity_previous_deadline", rs.DateDeadline()).
		WithKey("default_res_id", rs.ResID()).
		WithKey("default_res_model", rs.ResModel())
	rs.ActionFeedback(feedback)
	if rs.ForceNext() {
		Activity := h.MailActivity().NewSet(rs.Env()).WithNewContext(ctx)
		res := Activity.New(Activity.DefaultGet())
		res.OnchangePreviousActivityType()
		res.OnchangeActivityType()
		Activity.Create(res.First())
		return &actions.Action{}
	}
	return &actions.Action{
		Name:     rs.T("Schedule an Activity"),
		Context:  ctx,
		ViewMode: "form",
		Model:    "MailActivity",
		Type:     actions.ActionActWindow,
		Target:   "new",
	}
}

// ActionCloseDialog returns the close window action
func mailActivity_ActionCloseDialog(rs m.MailActivitySet) *actions.Action {
	return &actions.Action{
		Type: actions.ActionCloseWindow,
	}
}

// ActivityFormat
func mailActivity_ActivityFormat(rs m.MailActivitySet) []m.MailActivityData {
	// TODO Check that it works or make method adapter
	return rs.All()
}

// GetActivityData returns the necessary data for the web client
func mailActivity_GetActivityData(rs m.MailActivitySet, res_model interface{}, domain interface{}) {
	//        activity_domain = [('res_model', '=', res_model)]
	//        if domain:
	//            res = self.env[res_model].search(domain)
	//            activity_domain.append(('res_id', 'in', res.ids))
	//        grouped_activities = self.env['mail.activity'].read_group(
	//            activity_domain,
	//            ['res_id', 'activity_type_id',
	//                'ids:array_agg(id)', 'date_deadline:min(date_deadline)'],
	//            ['res_id', 'activity_type_id'],
	//            lazy=False)
	//        if not domain:
	//            res_ids = tuple(a['res_id'] for a in grouped_activities)
	//            res = self.env[res_model].search([('id', 'in', res_ids)])
	//            grouped_activities = [
	//                a for a in grouped_activities if a['res_id'] in res.ids]
	//        activity_type_ids = self.env['mail.activity.type']
	//        res_id_to_deadline = {}
	//        activity_data = defaultdict(dict)
	//        for group in grouped_activities:
	//            res_id = group['res_id']
	//            activity_type_id = group['activity_type_id'][0]
	//            activity_type_ids |= self.env['mail.activity.type'].browse(
	//                activity_type_id)  # we will get the name when reading mail_template_ids
	//            res_id_to_deadline[res_id] = group['date_deadline'] if (
	//                res_id not in res_id_to_deadline or group['date_deadline'] < res_id_to_deadline[res_id]) else res_id_to_deadline[res_id]
	//            state = self._compute_state_from_date(
	//                group['date_deadline'], self.user_id.sudo().tz)
	//            activity_data[res_id][activity_type_id] = {
	//                'count': group['__count'],
	//                'ids': group['ids'],
	//                'state': state,
	//                'o_closest_deadline': group['date_deadline'],
	//            }
	//        res_ids_sorted = sorted(
	//            res_id_to_deadline, key=lambda item: res_id_to_deadline[item])
	//        res_id_to_name = dict(
	//            self.env[res_model].browse(res_ids_sorted).name_get())
	//        activity_type_infos = []
	//        for elem in sorted(activity_type_ids, key=lambda item: item.sequence):
	//            mail_template_info = []
	//            for mail_template_id in elem.mail_template_ids:
	//                mail_template_info.append(
	//                    {"id": mail_template_id.id, "name": mail_template_id.name})
	//            activity_type_infos.append(
	//                [elem.id, elem.name, mail_template_info])
	//        return {
	//            'activity_types': activity_type_infos,
	//            'res_ids': [(rid, res_id_to_name[rid]) for rid in res_ids_sorted],
	//            'grouped_activities': activity_data,
	//            'model': res_model,
	//        }
}

var fields_MailActivityMixin = map[string]models.FieldDefinition{
	"Activities": fields.One2Many{
		JSON:          "activity_ids",
		RelationModel: h.MailActivity(),
		ReverseFK:     "ResID",
		String:        "Activities",
		// groups="base.group_user"
		Filter: q.MailActivity().ResModel().EqualsFunc(func(r models.RecordSet) string {
			return r.ModelName()
		})},

	"ActivityState": fields.Selection{
		Selection: types.Selection{
			"overdue": "Overdue",
			"today":   "Today",
			"planned": "Planned",
		},
		String:  "Activity State",
		Compute: h.MailActivityMixin().Methods().ComputeActivityState(),
		// groups="base.group_user"
		Help: "Status based on activities" +
			"Overdue: Due date is already passed" +
			"Today: Activity date is today" +
			"Planned: Future activities."},

	"ActivityUser": fields.Many2One{
		RelationModel: h.User(),
		String:        "Responsible User",
		Related:       `Activities.User`,
		ReadOnly:      false,
		// search='_search_activity_user_id'
		// groups="base.group_user"
	},

	"ActivityType": fields.Many2One{
		RelationModel: h.MailActivityType(),
		String:        "Next Activity Type",
		Related:       `Activities.ActivityType`,
		ReadOnly:      false,
		// search='_search_activity_type_id'
		// groups="base.group_user"
	},

	"ActivityDateDeadline": fields.Date{
		String:  "Next Activity Deadline",
		Compute: h.MailActivityMixin().Methods().ComputeActivityDateDeadline(),
		// search='_search_activity_date_deadline'
		ReadOnly: true,
		Stored:   false,
		// groups="base.group_user"
	},

	"ActivitySummary": fields.Char{
		String:   "Next Activity Summary",
		Related:  `Activities.Summary`,
		ReadOnly: false,
		// search='_search_activity_summary'
		// groups="base.group_user"
	},
}

// ComputeActivityState returns the activity state of this record from its activities
func mailActivityMixin_ComputeActivityState(rs m.MailActivityMixinSet) m.MailActivityMixinData {
	res := h.MailActivityMixin().NewData()
	states := make(map[string]bool)
	for _, activity := range rs.Activities().Records() {
		states[activity.State()] = true
	}
	switch {
	case states["overdue"]:
		res.SetActivityState("overdue")
	case states["today"]:
		res.SetActivityState("today")
	case states["planned"]:
		res.SetActivityState("planned")
	}
	return res
}

// ComputeActivityDateDeadline computes this record's deadline from its activities
func mailActivityMixin_ComputeActivityDateDeadline(rs m.MailActivityMixinSet) m.MailActivityMixinData {
	res := h.MailActivityMixin().NewData()
	if rs.Activities().IsNotEmpty() {
		res.SetActivityDateDeadline(rs.Activities().Records()[0].DateDeadline())
	}
	return res
}

// SearchActivityDateDeadline
func mailActivityMixin_SearchActivityDateDeadline(rs m.MailActivityMixinSet, operator interface{}, operand interface{}) {
	//        if operator == '=' and not operand:
	//            return [('activity_ids', '=', False)]
	//        return [('activity_ids.date_deadline', operator, operand)]
}

// SearchActivityUserId
func mailActivityMixin_SearchActivityUser(rs m.MailActivityMixinSet, operator interface{}, operand interface{}) {
	//        return [('activity_ids.user_id', operator, operand)]
}

// SearchActivityTypeId
func mailActivityMixin_SearchActivityType(rs m.MailActivityMixinSet, operator interface{}, operand interface{}) {
	//        return [('activity_ids.activity_type_id', operator, operand)]
}

// SearchActivitySummary
func mailActivityMixin_SearchActivitySummary(rs m.MailActivityMixinSet, operator interface{}, operand interface{}) {
	//        return [('activity_ids.summary', operator, operand)]
}

func mailActivityMixin_Write(rs m.MailActivityMixinSet, vals m.MailActivityMixinData) bool {
	// Delete activities of archived record.
	activeField, exists := rs.Collection().Model().Fields().Get("active")
	if !exists {
		return rs.Super().Write(vals)
	}
	if vals.Has(activeField) && !vals.Get(activeField).(bool) {
		h.MailActivity().NewSet(rs.Env()).Sudo().Search(q.MailActivity().
			ResModel().Equals(rs.ModelName()).And().
			ResID().In(rs.Ids())).Unlink()
	}
	return rs.Super().Write(vals)
}

func mailActivityMixin_Unlink(rs m.MailActivityMixinSet) int64 {
	// Override unlink to delete records activities through (res_model, res_id).
	res := rs.Super().Unlink()
	h.MailActivity().NewSet(rs.Env()).Sudo().Search(q.MailActivity().
		ResModel().Equals(rs.ModelName()).And().
		ResID().In(rs.Ids())).Unlink()
	return res
}

func mailActivityMixin_ToggleActive(rs m.MailActivityMixinSet) {
	// Before archiving the record we should also remove its ongoing
	// activities. Otherwise they stay in the systray
	// and concerning archived records it makes no sense.
	activeField, exists := rs.Collection().Model().Fields().Get("active")
	if !exists {
		return
	}
	recordsToDeactivate := rs.Filtered(func(r m.MailActivityMixinSet) bool {
		return r.Get(activeField).(bool)
	})
	if recordsToDeactivate.IsNotEmpty() {
		// use a sudo to bypass every access rights; all activities should be removed
		h.MailActivity().NewSet(rs.Env()).Sudo().Search(q.MailActivity().
			ResModel().Equals(rs.ModelName()).And().
			ResID().In(recordsToDeactivate.Ids())).Unlink()
	}
	rs.Super().ToggleActive()
}

// ActivitySendMail automatically sends an email based on the given MailTemplate, given its ID.
func mailActivityMixin_ActivitySendMail(rs m.MailActivityMixinSet, template_id int64) bool {
	//        template = self.env['mail.template'].browse(template_id).exists()
	//        if not template:
	//            return False
	//        for record in self.with_context(mail_post_autofollow=True):
	//        return True
	template := h.MailTemplate().BrowseOne(rs.Env(), template_id).Fetch()
	if template.IsEmpty() {
		return false
	}
	for _, record := range rs.WithContext("mailçpost_autofollow", true).Records() {
		//            record.message_post_with_template(
		//                template_id,
		//                composition_mode='comment'
		//            )
		// TODO add missing params
		record.MessagePostWithTemplate(template_id)
	}
	return true
}

// ActivitySchedule schedules an activity on each record of the current record set.
// This method allow to provide as parameter act_type_xmlid.
// This is an external ID of activity type instead of directly giving an activityTypID.
// It is useful to avoid having various "GetRecord()" in the code and allow
// to let the mixin handle access rights.
func mailActivityMixin_ActivitySchedule(rs m.MailActivityMixinSet, actTypeExternalID string, dateDeadline dates.Date,
	summary string, note string, actValues m.MailActivityData) m.MailActivitySet {

	activities := h.MailActivity().NewSet(rs.Env())
	if rs.Env().Context().GetBool("mail_activity_automation_skip") {
		return activities
	}
	if dateDeadline.IsZero() {
		dateDeadline = dates.TZToday(rs.Env().Context().GetString("tz"))
	}
	activityType := h.MailActivityType().NewSet(rs.Env())
	switch actTypeExternalID {
	case "":
		activityType = actValues.ActivityType()
	default:
		activityType = h.MailActivityType().NewSet(rs.Env()).Sudo().GetRecord(actTypeExternalID)
	}
	for _, record := range rs.Records() {
		sum := summary
		if sum == "" {
			sum = activityType.Summary()
		}
		createVals := h.MailActivity().NewData().
			SetActivityType(activityType).
			SetSummary(sum).
			SetAutomated(true).
			SetNote(note).
			SetDateDeadline(dateDeadline).
			SetResModel(rs.ModelName()).
			SetResID(record.ID())
		createVals.MergeWith(actValues)
		activities = activities.Union(h.MailActivity().Create(rs.Env(), createVals))
	}
	return activities
}

// ActivityScheduleWithView is a helper method to schedule an activity on each record of
// the current record set.
// This method allow to the same mecanism as `ActivitySchedule()`, but provides
// 2 additionnal parameters:
// 	- templateID: id of the of the hweb template to render
//  - renderContext: the values required to render the given hweb template
func mailActivityMixin_ActivityScheduleWithView(rs m.MailActivityMixinSet, actTypeExternalID string, dateDeadline dates.Date,
	summary string, templateID string, renderContext map[string]interface{}, actValues m.MailActivityData) m.MailActivitySet {

	activities := h.MailActivity().NewSet(rs.Env())
	if rs.Env().Context().GetBool("mail_activity_automation_skip") {
		return activities
	}
	if renderContext == nil {
		renderContext = make(map[string]interface{})
	}
	if templateID == "" {
		return activities
	}
	for _, record := range rs.Records() {
		renderContext["object"] = record
		//            note = views.render(
		//                render_context, engine='ir.qweb', minimal_qcontext=True)
		// TODO implement me
		note := ""
		activities = activities.Union(record.ActivitySchedule(actTypeExternalID, dateDeadline, summary, note, actValues))
	}
	return activities
}

// ActivityReschedule reschedules some automated activities.
// Activities to reschedule are selected based on type external ids and optionally by
// user. Purpose is to be able to
//  - update the deadline to dateDeadline;
//  - update the responsible to newUser;
func mailActivityMixin_ActivityReschedule(rs m.MailActivityMixinSet, actTypeExternalIds []string, user m.UserSet,
	dateDeadline dates.Date, newUser m.UserSet) m.MailActivitySet {

	activities := h.MailActivity().NewSet(rs.Env())
	if rs.Env().Context().GetBool("mail_activity_automation_skip") {
		return activities
	}
	activityTypes := h.MailActivityType().NewSet(rs.Env())
	for _, extID := range actTypeExternalIds {
		activityTypes = activityTypes.Union(h.MailActivityType().NewSet(rs.Env()).GetRecord(extID))
	}
	cond := q.MailActivity().
		ResModel().Equals(rs.ModelName()).And().
		ResID().In(rs.Ids()).And().
		Automated().Equals(true).And().
		ActivityType().In(activityTypes)
	if user.IsNotEmpty() {
		cond = cond.And().User().Equals(user)
	}
	activities = h.MailActivity().NewSet(rs.Env()).Search(cond)
	if activities.IsNotEmpty() {
		writeVals := h.MailActivity().NewData()
		if !dateDeadline.IsZero() {
			writeVals.SetDateDeadline(dateDeadline)
		}
		if newUser.IsNotEmpty() {
			writeVals.SetUser(newUser)
		}
	}
	return activities
}

// ActivityFeedback set activities as done, limiting to some activity types and
// optionally to a given user.
func mailActivityMixin_ActivityFeedback(rs m.MailActivityMixinSet, actTypeExternalIds []string, user m.UserSet,
	feedback string) bool {

	if rs.Env().Context().GetBool("mail_activity_automation_skip") {
		return false
	}
	activityTypes := h.MailActivityType().NewSet(rs.Env())
	for _, extID := range actTypeExternalIds {
		activityTypes = activityTypes.Union(h.MailActivityType().NewSet(rs.Env()).GetRecord(extID))
	}
	cond := q.MailActivity().
		ResModel().Equals(rs.ModelName()).And().
		ResID().In(rs.Ids()).And().
		Automated().Equals(true).And().
		ActivityType().In(activityTypes)
	if user.IsNotEmpty() {
		cond = cond.And().User().Equals(user)
	}
	activities := h.MailActivity().NewSet(rs.Env()).Search(cond)
	if activities.IsNotEmpty() {
		activities.ActionFeedback(feedback)
	}
	return true
}

// ActivityUnlink unlinks activities, limiting to some activity types and optionally to a given user.
func mailActivityMixin_ActivityUnlink(rs m.MailActivityMixinSet, actTypeExternalIds []string, user m.UserSet) bool {

	if rs.Env().Context().GetBool("mail_activity_automation_skip") {
		return false
	}
	activityTypes := h.MailActivityType().NewSet(rs.Env())
	for _, extID := range actTypeExternalIds {
		activityTypes = activityTypes.Union(h.MailActivityType().NewSet(rs.Env()).GetRecord(extID))
	}
	cond := q.MailActivity().
		ResModel().Equals(rs.ModelName()).And().
		ResID().In(rs.Ids()).And().
		Automated().Equals(true).And().
		ActivityType().In(activityTypes)
	if user.IsNotEmpty() {
		cond = cond.And().User().Equals(user)
	}
	h.MailActivity().NewSet(rs.Env()).Search(cond).Unlink()
	return true
}

func init() {
	models.NewModel("MailActivityType")
	h.MailActivityType().AddFields(fields_MailActivityType)
	h.MailActivityType().NewMethod("OnchangeResModel", mailActivityType_OnchangeResModel)
	h.MailActivityType().NewMethod("ComputeInitialResModel", mailActivityType_ComputeInitialResModel)

	models.NewModel("MailActivity")
	h.MailActivity().AddFields(fields_MailActivity)
	h.MailActivity().NewMethod("ComputeHasRecommendedActivities", mailActivity_ComputeHasRecommendedActivities)
	h.MailActivity().NewMethod("OnchangePreviousActivityType", mailActivity_OnchangePreviousActivityType)
	h.MailActivity().NewMethod("ComputeResName", mailActivity_ComputeResName)
	h.MailActivity().NewMethod("ComputeState", mailActivity_ComputeState)
	h.MailActivity().NewMethod("ComputeStateFromDate", mailActivity_ComputeStateFromDate)
	h.MailActivity().NewMethod("OnchangeActivityType", mailActivity_OnchangeActivityType)
	h.MailActivity().NewMethod("OnchangeRecommendedActivityType", mailActivity_OnchangeRecommendedActivityType)
	h.MailActivity().NewMethod("CheckAccess", mailActivity_CheckAccess)
	h.MailActivity().NewMethod("CheckAccessAssignation", mailActivity_CheckAccessAssignation)
	h.MailActivity().Methods().Create().Extend(mailActivity_Create)
	h.MailActivity().Methods().Write().Extend(mailActivity_Write)
	h.MailActivity().Methods().Unlink().Extend(mailActivity_Unlink)
	h.MailActivity().NewMethod("ActionNotify", mailActivity_ActionNotify)
	h.MailActivity().NewMethod("ActionDone", mailActivity_ActionDone)
	h.MailActivity().NewMethod("ActionFeedback", mailActivity_ActionFeedback)
	h.MailActivity().NewMethod("ActionDoneScheduleNext", mailActivity_ActionDoneScheduleNext)
	h.MailActivity().NewMethod("ActionFeedbackScheduleNext", mailActivity_ActionFeedbackScheduleNext)
	h.MailActivity().NewMethod("ActionCloseDialog", mailActivity_ActionCloseDialog)
	h.MailActivity().NewMethod("ActivityFormat", mailActivity_ActivityFormat)
	h.MailActivity().NewMethod("GetActivityData", mailActivity_GetActivityData)

	models.NewMixinModel("MailActivityMixin")
	h.MailActivityMixin().InheritModel(h.ModelMixin())
	h.MailActivityMixin().InheritModel(h.MailThread())
	h.MailActivityMixin().AddFields(fields_MailActivityMixin)
	h.MailActivityMixin().NewMethod("ComputeActivityState", mailActivityMixin_ComputeActivityState)
	h.MailActivityMixin().NewMethod("ComputeActivityDateDeadline", mailActivityMixin_ComputeActivityDateDeadline)
	h.MailActivityMixin().NewMethod("SearchActivityDateDeadline", mailActivityMixin_SearchActivityDateDeadline)
	h.MailActivityMixin().NewMethod("SearchActivityUserId", mailActivityMixin_SearchActivityUser)
	h.MailActivityMixin().NewMethod("SearchActivityTypeId", mailActivityMixin_SearchActivityType)
	h.MailActivityMixin().NewMethod("SearchActivitySummary", mailActivityMixin_SearchActivitySummary)
	h.MailActivityMixin().Methods().Write().Extend(mailActivityMixin_Write)
	h.MailActivityMixin().Methods().Unlink().Extend(mailActivityMixin_Unlink)
	h.MailActivityMixin().Methods().ToggleActive().Extend(mailActivityMixin_ToggleActive)
	h.MailActivityMixin().NewMethod("ActivitySendMail", mailActivityMixin_ActivitySendMail)
	h.MailActivityMixin().NewMethod("ActivitySchedule", mailActivityMixin_ActivitySchedule)
	h.MailActivityMixin().NewMethod("ActivityScheduleWithView", mailActivityMixin_ActivityScheduleWithView)
	h.MailActivityMixin().NewMethod("ActivityReschedule", mailActivityMixin_ActivityReschedule)
	h.MailActivityMixin().NewMethod("ActivityFeedback", mailActivityMixin_ActivityFeedback)
	h.MailActivityMixin().NewMethod("ActivityUnlink", mailActivityMixin_ActivityUnlink)
}
