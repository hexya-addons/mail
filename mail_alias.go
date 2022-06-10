package mail

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hexya-addons/mail/mailtypes"
	"github.com/hexya-erp/hexya/src/actions"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/tools/strutils"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// A Mail Alias is a mapping of an email address with a given Hexya Document
// model. It is used by Hexya's mail gateway when processing incoming emails
// sent to the system. If the recipient address (To) of the message matches
// a Mail Alias, the message will be either processed following the rules
// of that alias. If the message is a reply it will be attached to the
// existing discussion on the corresponding record, otherwise a new
// record of the corresponding model will be created.

// This is meant to be used in combination with a catch-all email configuration
// on the company's mail server, so that as soon as a new mail.alias is
// created, it becomes immediately usable and Hexya will accept email for it.

var fields_MailAlias = map[string]models.FieldDefinition{
	"AliasName": fields.Char{
		String: "Alias Name",
		Unique: true,
		Help: "The name of the email alias, e.g. 'jobs' if you want to" +
			"catch emails for <jobs@example.odoo.com>"},

	"AliasModel": fields.Char{
		String:   "Aliased Model",
		Required: true,
		Help: "The model (Odoo Document Kind) to which this alias corresponds." +
			"Any incoming email that does not reply to an existing record" +
			"will cause the creation of a new record of this model (e.g." +
			"a Project Task)"},

	"AliasUser": fields.Many2One{
		RelationModel: h.User(),
		String:        "Owner",
		Default: func(env models.Environment) interface{} {
			return h.User().NewSet(env).CurrentUser()
		},
		Help: "The owner of records created upon receiving emails on this" +
			"alias. If this field is not set the system will attempt" +
			"to find the right owner based on the sender (From) address," +
			"or will use the Administrator account if no system user" +
			"is found for that address."},

	"AliasDefaults": fields.Text{
		String:     "Default Values",
		Required:   true,
		Default:    models.DefaultValue("{}"),
		Constraint: h.MailAlias().Methods().CheckAliasDefaults(),
		Help: "A JSON object that will be evaluated to provide default" +
			"values when creating new records for this alias."},

	"AliasForceThreadID": fields.Integer{
		String: "Record Thread ID",
		Help: "Optional ID of a thread (record) to which all incoming" +
			"messages will be attached, even if they did not reply to" +
			"it. If set, this will disable the creation of new records completely."},

	"AliasDomain": fields.Char{
		String:  "Alias domain",
		Compute: h.MailAlias().Methods().ComputeAliasDomain(),
		Default: func(env models.Environment) interface{} {
			return h.ConfigParameter().NewSet(env).Sudo().GetParam("mail.catchall.domain", "")
		}},

	"AliasParentModel": fields.Char{
		String: "Parent Model",
		Help: "Parent model holding the alias. The model holding the alias" +
			"reference is not necessarily the model given by AliasModel" +
			"(example: project (ParentModel) and task (Model))"},

	"AliasParentThreadID": fields.Integer{
		String: "Parent Record Thread ID",
		Help: "ID of the parent record holding the alias (example: project" +
			"holding the task creation alias)"},

	"AliasContact": fields.Selection{
		Selection: types.Selection{
			"everyone":  "Everyone",
			"partners":  "Authenticated Partners",
			"followers": "Followers only",
		},
		Default:  models.DefaultValue("everyone"),
		String:   "Alias Contact Security",
		Required: true,
		Help: "Policy to post a message on the document using the mailgateway." +
			"- everyone: everyone can post" +
			"- partners: only authenticated partners" +
			"- followers: only followers of the related document or" +
			"members of following channels" +
			""},
}

// ComputeAliasDomain returns the AliasDomain defined in configuration
func mailAlias_ComputeAliasDomain(rs m.MailAliasSet) m.MailAliasData {
	res := h.MailAlias().NewData()
	res.SetAliasDomain(h.ConfigParameter().NewSet(rs.Env()).Sudo().GetParam("mail.catchall.domain", ""))
	return res
}

// CheckAliasDefaults
func mailAlias_CheckAliasDefaults(rs m.MailAliasSet) {
	// TODO find a way to do this

	//        try:
	//            dict(safe_eval(self.alias_defaults))
	//        except Exception:
	//            raise ValidationError(
	//                _('Invalid expression, it must be a literal python dictionary definition e.g. "{\'field\': \'value\'}"'))
}

func mailAlias_Create(rs m.MailAliasSet, vals m.MailAliasData) m.MailAliasSet {
	modelName := rs.Env().Context().GetString("alias_model_name")
	if modelName != "" {
		vals.SetAliasModel(modelName)
	}
	parentModelName := rs.Env().Context().GetString("alias_parent_model_name")
	if parentModelName != "" {
		vals.SetAliasParentModel(parentModelName)
	}
	return rs.Super().Create(vals)
}

func mailAlias_Write(rs m.MailAliasSet, vals m.MailAliasData) bool {
	// give a unique alias name if given alias name is already assigned
	if vals.AliasName() != "" && rs.IsNotEmpty() {
		vals.SetAliasName(rs.CleanAndMakeUnique(vals.AliasName(), rs.Ids()))
	}
	return rs.Super().Write(vals)
}

func mailAlias_NameGet(rs m.MailAliasSet) string {
	switch {
	case rs.AliasName() != "" && rs.AliasDomain() != "":
		return fmt.Sprintf("%s@%s", rs.AliasName(), rs.AliasDomain())
	case rs.AliasName() != "":
		return rs.AliasName()
	default:
		return rs.T("Inactive Alias")
	}
}

// FindUnique find a unique alias name similar to ``name``.
// If ``name`` is already taken, make a variant by adding an integer
// suffix until an unused alias is found.
func mailAlias_FindUnique(rs m.MailAliasSet, name string, aliasIds []int64) string {
	var (
		sequence int
		newName  string
	)
	for {
		switch sequence {
		case 0:
			newName = name
		default:
			newName = fmt.Sprintf("%s%d", name, sequence)
		}
		cond := q.MailAlias().AliasName().Equals(newName)
		if len(aliasIds) > 0 {
			cond = cond.And().ID().NotIn(aliasIds)
		}
		if h.MailAlias().Search(rs.Env(), cond).IsEmpty() {
			break
		}
		if sequence == 0 {
			sequence = 1
		}
		sequence++
	}
	return newName
}

// CleanAndMakeUnique returns a clean and unique mail alias
func mailAlias_CleanAndMakeUnique(rs m.MailAliasSet, name string, aliasIds []int64) string {
	// when an alias name appears to already be an email, we keep the local part only
	name = strings.Split(strings.ToLower(strutils.RemoveAccent(name)), "@")[0]
	name = regexp.MustCompile(`[^\w+.]+`).ReplaceAllString(name, "-")
	return rs.FindUnique(name, aliasIds)
}

// OpenDocument returns an action that opens the record this activity is attached to
func mailAlias_OpenDocument(rs m.MailAliasSet) *actions.Action {
	if rs.AliasModel() == "" || rs.AliasForceThreadID() == 0 {
		return &actions.Action{}
	}
	return &actions.Action{
		Type:     actions.ActionActWindow,
		ViewMode: "form",
		Model:    rs.AliasModel(),
		ResID:    rs.AliasForceThreadID(),
	}
}

// OpenParentDocument returns an action that opens the parent record this activity is attached to
func mailAlias_OpenParentDocument(rs m.MailAliasSet) *actions.Action {
	if rs.AliasParentModel() == "" || rs.AliasParentThreadID() == 0 {
		return &actions.Action{}
	}
	return &actions.Action{
		Type:     actions.ActionActWindow,
		ViewMode: "form",
		Model:    rs.AliasParentModel(),
		ResID:    rs.AliasParentThreadID(),
	}
}

var fields_MailAliasMixin = map[string]models.FieldDefinition{
	"Alias": fields.Many2One{
		RelationModel: h.MailAlias(),
		String:        "Alias",
		OnDelete:      models.Restrict,
		Required:      true,
		Embed:         true,
	},
}

// GetAliasModelName returns the model name for the alias.
// Incoming emails that are not replies to existing records will cause the
// creation of a new record of this alias model.
// The value may depend on vals, RecordData passed to Create() when a record of
// this model is created.
func mailAliasMixin_GetAliasModelName(rs m.MailAliasMixinSet, vals models.RecordData) string {
	return ""
}

// GetAliasValues return values to create an alias,
// or to write on the alias after its creation.
func mailAliasMixin_GetAliasValues(rs m.MailAliasMixinSet) m.MailAliasData {
	return h.MailAlias().NewData().SetAliasParentThreadID(rs.ID())
}

func mailAliasMixin_Create(rs m.MailAliasMixinSet, vals m.MailAliasMixinData) m.MailAliasMixinSet {
	record := rs.Super().
		WithContext("alias_model_name", rs.GetAliasModelName(vals.Underlying())).
		WithContext("alias_parent_model_name", rs.ModelName()).
		Create(vals)
	record.Alias().Sudo().Write(rs.GetAliasValues())
	return record
}

func mailAliasMixin_Unlink(rs m.MailAliasMixinSet) int64 {
	aliases := h.MailAlias().NewSet(rs.Env())
	for _, rec := range rs.Records() {
		aliases = aliases.Union(rec.Alias())
	}
	res := rs.Super().Unlink()
	aliases.Unlink()
	return res
}

// InitColumn create aliases for existing rows
func mailAliasMixin_InitColumn(rs m.MailAliasMixinSet, name string) {
	//        super(AliasMixin, self)._init_column(name)
	//        if name != 'alias_id':
	//            return
	//        IM = self.env['ir.model']
	//        IM._reflect_model(self)
	//        IM._reflect_model(self.env[self.get_alias_model_name({})])
	//        alias_ctx = {
	//            'alias_model_name': self.get_alias_model_name({}),
	//            'alias_parent_model_name': self._name,
	//        }
	//        alias_model = self.env['mail.alias'].sudo(
	//        ).with_context(alias_ctx).browse([])
	//        child_ctx = {
	//            'active_test': False,       # retrieve all records
	//            'prefetch_fields': False,   # do not prefetch fields on records
	//        }
	//        child_model = self.sudo().with_context(child_ctx).browse([])
	//        for record in child_model.search([('alias_id', '=', False)]):
	//            # create the alias, and link it to the current record
	//            alias = alias_model.create(record.get_alias_values())
	//            record.with_context(mail_notrack=True).alias_id = alias
	//            _logger.info('Mail alias created for %s %s (id %s)',
	//                         record._name, record.display_name, record.id)
}

// AliasCheckContact is the main mixin method that inheriting models may inherit in order
// to implement a specifc behavior.
func mailAliasMixin_AliasCheckContact(rs m.MailAliasMixinSet, message string,
	messageDict map[string]interface{}, alias m.MailAliasSet) mailtypes.CheckContactResult {

	return rs.AliasCheckContactOnRecord(rs, message, messageDict, alias)
}

// AliasCheckContactOnRecord is a generic method that takes a record not necessarily inheriting from
// MailAliasMixin.
func mailAliasMixin_AliasCheckContactOnRecord(rs m.MailAliasMixinSet, record models.RecordSet, message string,
	messageDict map[string]interface{}, alias m.MailAliasSet) mailtypes.CheckContactResult {

	//        author = self.env['res.partner'].browse(
	//            message_dict.get('author_id', False))
	author := h.Partner().BrowseOne(rs.Env(), messageDict["author_id"].(int64))
	switch alias.AliasContact() {
	case "followers":
		if record.IsEmpty() {
			return mailtypes.CheckContactResult{
				ErrorMessage: rs.T("incorrectly configured alias (unknown reference record)"),
			}
		}
		recordThread, ok := record.(m.MailThreadSet)
		if !ok {
			return mailtypes.CheckContactResult{
				ErrorMessage: rs.T("incorrectly configured alias (not a MailThread)"),
			}
		}
		channelPartners := h.Partner().NewSet(rs.Env())
		for _, channel := range recordThread.MessageChannels().Records() {
			channelPartners = channelPartners.Union(channel.ChannelPartners())
		}
		acceptedPartners := recordThread.MessagePartners().Union(channelPartners)
		if author.IsEmpty() || author.Intersect(acceptedPartners).IsEmpty() {
			return mailtypes.CheckContactResult{
				ErrorMessage: rs.T("restricted to followers"),
			}
		}
	case "partners":
		if author.IsEmpty() {
			return mailtypes.CheckContactResult{
				ErrorMessage: rs.T("restricted to known authors"),
			}
		}
	}
	return mailtypes.CheckContactResult{
		OK: true,
	}
}

func init() {
	models.NewModel("MailAlias")
	h.MailAlias().SetDefaultOrder("AliasModel", "AliasName")
	h.MailAlias().SetDescription("Email Alias")
	h.MailAlias().AddFields(fields_MailAlias)
	h.MailAlias().NewMethod("ComputeAliasDomain", mailAlias_ComputeAliasDomain)
	h.MailAlias().NewMethod("CheckAliasDefaults", mailAlias_CheckAliasDefaults)
	h.MailAlias().Methods().Create().Extend(mailAlias_Create)
	h.MailAlias().Methods().Write().Extend(mailAlias_Write)
	h.MailAlias().Methods().NameGet().Extend(mailAlias_NameGet)
	h.MailAlias().NewMethod("FindUnique", mailAlias_FindUnique)
	h.MailAlias().NewMethod("CleanAndMakeUnique", mailAlias_CleanAndMakeUnique)
	h.MailAlias().NewMethod("OpenDocument", mailAlias_OpenDocument)
	h.MailAlias().NewMethod("OpenParentDocument", mailAlias_OpenParentDocument)

	models.NewModel("MailAliasMixin")
	h.MailAliasMixin().AddFields(fields_MailAliasMixin)
	h.MailAliasMixin().SetDescription("Email Aliases Mixin")
	h.MailAliasMixin().NewMethod("GetAliasModelName", mailAliasMixin_GetAliasModelName)
	h.MailAliasMixin().NewMethod("GetAliasValues", mailAliasMixin_GetAliasValues)
	h.MailAliasMixin().Methods().Create().Extend(mailAliasMixin_Create)
	h.MailAliasMixin().Methods().Unlink().Extend(mailAliasMixin_Unlink)
	h.MailAliasMixin().NewMethod("InitColumn", mailAliasMixin_InitColumn)
	h.MailAliasMixin().NewMethod("AliasCheckContact", mailAliasMixin_AliasCheckContact)
	h.MailAliasMixin().NewMethod("AliasCheckContactOnRecord", mailAliasMixin_AliasCheckContactOnRecord)

}
