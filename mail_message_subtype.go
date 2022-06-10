package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

var fields_MailMessageSubtype = map[string]models.FieldDefinition{
	"Name": fields.Char{
		String:    "Message Type",
		Required:  true,
		Translate: true,
		Help: "Message subtype gives a more precise type on the message," +
			"especially for system notifications. For example, it can" +
			"be a notification related to a new record (New), or to" +
			"a stage change in a process (Stage change). Message subtypes" +
			"allow to precisely tune the notifications the user want" +
			"to receive on its wall."},

	"Description": fields.Text{
		String:    "Description",
		Translate: true,
		Help: "Description that will be added in the message posted for" +
			"this subtype. If void, the name will be added instead."},

	"Internal": fields.Boolean{
		String: "Internal Only",
		Help: "Messages with internal subtypes will be visible only by" +
			"employees, aka members of base_user group"},

	"Parent": fields.Many2One{
		RelationModel: h.MailMessageSubtype(),
		String:        "Parent",
		OnDelete:      models.SetNull,
		Help: "Parent subtype, used for automatic subscription. This field" +
			"is not correctly named. For example on a project, the parent_id" +
			"of project subtypes refers to task-related subtypes."},

	"RelationField": fields.Char{
		String: "Relation field",
		Help: "Field used to link the related model to the subtype model" +
			"when using automatic subscription on a related document." +
			"The field is used to compute getattr(related_document.relation_field)."},

	"ResModel": fields.Char{
		String: "Model",
		Help: "Model the subtype applies to. If False, this subtype applies" +
			"to all models."},

	"Default": fields.Boolean{
		String:  "Default",
		Default: models.DefaultValue(true),
		Help:    "Activated by default when subscribing."},

	"Sequence": fields.Integer{
		String:  "Sequence",
		Default: models.DefaultValue(1),
		Help:    "Used to order subtypes."},

	"Hidden": fields.Boolean{
		String: "Hidden",
		Help:   "Hide the subtype in the follower options"},
}

// GetAutoSubscriptionSubtypes returns data related to auto subscription based on subtype matching.
//
// Example with tasks and project :
//
//          * generic: discussion, res_model = False
//          * task: new, res_model = project.task
//          * project: task_new, parent_id = new, res_model = project.project, field = project_id
//
// Returned data (in order):
//
//          * all: all subtypes that are generic or related to task and project
//          * def: for task, default subtypes ids
//          * internal: for task, internal-only default subtypes ids
//          * parent: dict(parent subtype id, child subtype id), i.e. {task_new.id: new.id}
//          * relation: dict(parent_model, relation_fields), i.e. {'project.project': ['project_id']}
func mailMessageSubtype_GetAutoSubscriptionSubtypes(rs m.MailMessageSubtypeSet, modelName string) (
	m.MailMessageSubtypeSet, m.MailMessageSubtypeSet, m.MailMessageSubtypeSet, map[int64]m.MailMessageSubtypeSet, map[string]models.FieldNames) {

	subtypes := h.MailMessageSubtype().NewSet(rs.Env()).Sudo().Search(q.MailMessageSubtype().
		ResModel().IsNull().Or().
		ResModel().Equals(modelName).Or().
		ParentFilteredOn(q.MailMessageSubtype().ResModel().Equals(modelName)))
	all := h.MailMessageSubtype().NewSet(rs.Env())
	def := h.MailMessageSubtype().NewSet(rs.Env())
	internal := h.MailMessageSubtype().NewSet(rs.Env())
	parent := make(map[int64]m.MailMessageSubtypeSet)
	relation := make(map[string]models.FieldNames)
	for _, subtype := range subtypes.Records() {
		switch {
		case subtype.ResModel() == "" || subtype.ResModel() == modelName:
			all = all.Union(subtype)
			if subtype.Default() {
				def = def.Union(subtype)
			}
		case subtype.RelationField() != "":
			parent[subtype.ID()] = subtype.Parent()
			relation[subtype.ResModel()] = append(relation[subtype.ResModel()], rs.Collection().Model().FieldName(subtype.RelationField()))
		}
		if subtype.Internal() {
			internal = internal.Union(subtype)
		}
	}
	return all, def, internal, parent, relation
}

// DefaultSubtypes retrieves the default subtypes (all, internal, external)
// for the given model.
func mailMessageSubtype_DefaultSubtypes(rs m.MailMessageSubtypeSet, modelName string) (m.MailMessageSubtypeSet, m.MailMessageSubtypeSet, m.MailMessageSubtypeSet) {
	cond := q.MailMessageSubtype().Default().Equals(true).
		AndCond(q.MailMessageSubtype().ResModel().Equals(modelName).Or().ResModel().IsNull())
	subTypes := h.MailMessageSubtype().Search(rs.Env(), cond)
	internal := subTypes.Filtered(func(r m.MailMessageSubtypeSet) bool {
		return r.Internal()
	})
	return subTypes, internal, subTypes.Subtract(internal)
}

func init() {
	models.NewModel("MailMessageSubtype")
	h.MailMessageSubtype().AddFields(fields_MailMessageSubtype)
	h.MailMessageSubtype().NewMethod("GetAutoSubscriptionSubtypes", mailMessageSubtype_GetAutoSubscriptionSubtypes)
	h.MailMessageSubtype().NewMethod("DefaultSubtypes", mailMessageSubtype_DefaultSubtypes)

}
