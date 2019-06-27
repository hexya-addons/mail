package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.MailMessageSubtype().DeclareModel()

	h.MailMessageSubtype().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:    "Message Type",
			Required:  true,
			Translate: true,
			Help: "Message subtype gives a more precise type on the message," +
				"especially for system notifications. For example, it can" +
				"be a notification related to a new record (New), or to" +
				"a stage change in a process (Stage change). Message subtypes" +
				"allow to precisely tune the notifications the user want" +
				"to receive on its wall.",
		},
		"Description": models.TextField{
			String:    "Description",
			Translate: true,
			Help: "Description that will be added in the message posted for" +
				"this subtype. If void, the name will be added instead.",
		},
		"Internal": models.BooleanField{
			String: "Internal Only",
			Help: "Messages with internal subtypes will be visible only by" +
				"employees, aka members of base_user group",
		},
		"ParentId": models.Many2OneField{
			RelationModel: h.MailMessageSubtype(),
			String:        "Parent",
			OnDelete:      `set null`,
			Help: "Parent subtype, used for automatic subscription. This field" +
				"is not correctly named. For example on a project, the parent_id" +
				"of project subtypes refers to task-related subtypes.",
		},
		"RelationField": models.CharField{
			String: "Relation field",
			Help: "Field used to link the related model to the subtype model" +
				"when using automatic subscription on a related document." +
				"The field is used to compute getattr(related_document.relation_field).",
		},
		"ResModel": models.CharField{
			String: "Model",
			Help: "Model the subtype applies to. If False, this subtype applies" +
				"to all models.",
		},
		"Default": models.BooleanField{
			String:  "Default",
			Default: models.DefaultValue(true),
			Help:    "Activated by default when subscribing.",
		},
		"Sequence": models.IntegerField{
			String:  "Sequence",
			Default: models.DefaultValue(1),
			Help:    "Used to order subtypes.",
		},
		"Hidden": models.BooleanField{
			String: "Hidden",
			Help:   "Hide the subtype in the follower options",
		},
	})
}
