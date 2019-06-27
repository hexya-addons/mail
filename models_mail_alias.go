package mail

	import (
		"net/http"

		"github.com/hexya-erp/hexya/src/controllers"
		"github.com/hexya-erp/hexya/src/models"
		"github.com/hexya-erp/hexya/src/models/types"
		"github.com/hexya-erp/hexya/src/models/types/dates"
		"github.com/hexya-erp/pool/h"
		"github.com/hexya-erp/pool/q"
	)
	
//import logging
//import re
//import unicodedata
//_logger = logging.getLogger(__name__)
func RemoveAccents(input_str interface{})  {
//    """Suboptimal-but-better-than-nothing way to replace accented
//    latin letters by an ASCII equivalent. Will obviously change the
//    meaning of input_str and work only for some cases"""
//    input_str = ustr(input_str)
//    nkfd_form = unicodedata.normalize('NFKD', input_str)
//    return u''.join([c for c in nkfd_form if not unicodedata.combining(c)])
}
func init() {
h.MailAlias().DeclareModel()
h.MailAlias().AddSQLConstraint("alias_unique", "UNIQUE(alias_name)", "Unfortunately this email alias is already used, please choose a unique one")




h.MailAlias().AddFields(map[string]models.FieldDefinition{
"AliasName": models.CharField{
String: "Alias Name",
Help: "The name of the email alias, e.g. 'jobs' if you want to" + 
"catch emails for <jobs@example.odoo.com>",
},
"AliasModelId": models.Many2OneField{
RelationModel: h.IrModel(),
String: "Aliased Model",
Required: true,
OnDelete: `cascade`,
Help: "The model (Odoo Document Kind) to which this alias corresponds." + 
"Any incoming email that does not reply to an existing record" + 
"will cause the creation of a new record of this model (e.g." + 
"a Project Task)",
Filter: q.FieldId().Name().Equals("message_ids"),
},
"AliasUserId": models.Many2OneField{
RelationModel: h.User(),
String: "Owner",
Default: func (env models.Environment) interface{} { return env.Uid() },
Help: "The owner of records created upon receiving emails on this" + 
"alias. If this field is not set the system will attempt" + 
"to find the right owner based on the sender (From) address," + 
"or will use the Administrator account if no system user" + 
"is found for that address.",
},
"AliasDefaults": models.TextField{
String: "Default Values",
Required: true,
Default: models.DefaultValue("{}"),
Help: "A Python dictionary that will be evaluated to provide default" + 
"values when creating new records for this alias.",
},
"AliasForceThreadId": models.IntegerField{
String: "Record Thread ID",
Help: "Optional ID of a thread (record) to which all incoming" + 
"messages will be attached, even if they did not reply to" + 
"it. If set, this will disable the creation of new records completely.",
},
"AliasDomain": models.CharField{
String: "Alias domain",
Compute: h.MailAlias().Methods().GetAliasDomain(),
Default: func (env models.Environment) interface{} { return env["ir.config_parameter"].get_param() },
},
"AliasParentModelId": models.Many2OneField{
RelationModel: h.IrModel(),
String: "Parent Model",
Help: "Parent model holding the alias. The model holding the alias" + 
"reference is not necessarily the model given by alias_model_id" + 
"(example: project (parent_model) and task (model))",
},
"AliasParentThreadId": models.IntegerField{
String: "Parent Record Thread ID",
Help: "ID of the parent record holding the alias (example: project" + 
"holding the task creation alias)",
},
"AliasContact": models.SelectionField{
Selection: types.Selection{
"everyone": "Everyone",
"partners": "Authenticated Partners",
"followers": "Followers only",
},
Default: models.DefaultValue("everyone"),
String: "Alias Contact Security",
Required: true,
Help: "Policy to post a message on the document using the mailgateway." + 
"- everyone: everyone can post" + 
"- partners: only authenticated partners" + 
"- followers: only followers of the related document or" + 
"members of following channels" + 
"",
},

})
h.MailAlias().Methods().GetAliasDomain().DeclareMethod(
`GetAliasDomain`,
func(rs h.MailAliasSet) h.MailAliasData {
//        alias_domain = self.env["ir.config_parameter"].get_param(
//            "mail.catchall.domain")
//        for record in self:
//            record.alias_domain = alias_domain
})
h.MailAlias().Methods().CheckAliasDefaults().DeclareMethod(
`CheckAliasDefaults`,
func(rs m.MailAliasSet)  {
//        try:
//            dict(safe_eval(self.alias_defaults))
//        except Exception:
//            raise ValidationError(
//                _('Invalid expression, it must be a literal python dictionary definition e.g. "{\'field\': \'value\'}"'))
})
h.MailAlias().Methods().Create().Extend(
` Creates an email.alias record according to the values
provided in ``vals``,
            with 2 alterations: the ``alias_name`` value
may be suffixed in order to
            make it unique (and certain unsafe characters replaced), and
            he ``alias_model_id`` value will set to the
model ID of the ``model_name``
            context value, if provided.
        `,
func(rs m.MailAliasSet, vals models.RecordData)  {
//        model_name = self._context.get('alias_model_name')
//        parent_model_name = self._context.get('alias_parent_model_name')
//        if vals.get('alias_name'):
//            vals['alias_name'] = self._clean_and_make_unique(
//                vals.get('alias_name'))
//        if model_name:
//            model = self.env['ir.model'].search([('model', '=', model_name)])
//            vals['alias_model_id'] = model.id
//        if parent_model_name:
//            model = self.env['ir.model'].search(
//                [('model', '=', parent_model_name)])
//            vals['alias_parent_model_id'] = model.id
//        return super(Alias, self).create(vals)
})
h.MailAlias().Methods().Write().Extend(
`"give a unique alias name if given alias name is already assigned`,
func(rs m.MailAliasSet, vals models.RecordData)  {
//        if vals.get('alias_name') and self.ids:
//            vals['alias_name'] = self._clean_and_make_unique(
//                vals.get('alias_name'), alias_ids=self.ids)
//        return super(Alias, self).write(vals)
})
h.MailAlias().Methods().NameGet().Extend(
`Return the mail alias display alias_name, including the implicit
           mail catchall domain if exists from config otherwise
"New Alias".
           e.g. `jobs@mail.odoo.com` or `jobs` or 'New Alias'
        `,
func(rs m.MailAliasSet)  {
//        res = []
//        for record in self:
//            if record.alias_name and record.alias_domain:
//                res.append((record['id'], "%s@%s" %
//                            (record.alias_name, record.alias_domain)))
//            elif record.alias_name:
//                res.append((record['id'], "%s" % (record.alias_name)))
//            else:
//                res.append((record['id'], _("Inactive Alias")))
//        return res
})
h.MailAlias().Methods().FindUnique().DeclareMethod(
`Find a unique alias name similar to ``name``. If ``name`` is
           already taken, make a variant by adding an integer
suffix until
           an unused alias is found.
        `,
func(rs m.MailAliasSet, name interface{}, alias_ids interface{})  {
//        sequence = None
//        while True:
//            new_name = "%s%s" % (
//                name, sequence) if sequence is not None else name
//            domain = [('alias_name', '=', new_name)]
//            if alias_ids:
//                domain += [('id', 'not in', alias_ids)]
//            if not self.search(domain):
//                break
//            sequence = (sequence + 1) if sequence else 2
//        return new_name
})
h.MailAlias().Methods().CleanAndMakeUnique().DeclareMethod(
`CleanAndMakeUnique`,
func(rs m.MailAliasSet, name interface{}, alias_ids interface{})  {
//        name = remove_accents(name).lower().split('@')[0]
//        name = re.sub(r'[^\w+.]+', '-', name)
//        return self._find_unique(name, alias_ids=alias_ids)
})
h.MailAlias().Methods().OpenDocument().DeclareMethod(
`OpenDocument`,
func(rs m.MailAliasSet)  {
//        if not self.alias_model_id or not self.alias_force_thread_id:
//            return False
//        return {
//            'view_type': 'form',
//            'view_mode': 'form',
//            'res_model': self.alias_model_id.model,
//            'res_id': self.alias_force_thread_id,
//            'type': 'ir.actions.act_window',
//        }
})
h.MailAlias().Methods().OpenParentDocument().DeclareMethod(
`OpenParentDocument`,
func(rs m.MailAliasSet)  {
//        if not self.alias_parent_model_id or not self.alias_parent_thread_id:
//            return False
//        return {
//            'view_type': 'form',
//            'view_mode': 'form',
//            'res_model': self.alias_parent_model_id.model,
//            'res_id': self.alias_parent_thread_id,
//            'type': 'ir.actions.act_window',
//        }
})
h.MailAliasMixin().DeclareModel()


h.MailAliasMixin().AddFields(map[string]models.FieldDefinition{
"AliasId": models.Many2OneField{
RelationModel: h.MailAlias(),
String: "Alias",
OnDelete: `restrict`,
Required: true,
},
})
h.MailAliasMixin().Methods().GetAliasModelName().DeclareMethod(
` Return the model name for the alias. Incoming emails that are not
            replies to existing records will cause the
creation of a new record
            of this alias model. The value may depend on
``vals``, the dict of
            values passed to ``create`` when a record of
this model is created.
        `,
func(rs m.MailAliasMixinSet, vals interface{})  {
//        return None
})
h.MailAliasMixin().Methods().GetAliasValues().DeclareMethod(
` Return values to create an alias, or to write on the alias after its
            creation.
        `,
func(rs m.MailAliasMixinSet)  {
//        return {'alias_parent_thread_id': self.id}
})
h.MailAliasMixin().Methods().Create().Extend(
` Create a record with ``vals``, and create a corresponding alias. `,
func(rs m.MailAliasMixinSet, vals models.RecordData)  {
//        record = super(AliasMixin, self.with_context(
//            alias_model_name=self.get_alias_model_name(vals),
//            alias_parent_model_name=self._name)).create(vals)
//        record.alias_id.sudo().write(record.get_alias_values())
//        return record
})
h.MailAliasMixin().Methods().Unlink().Extend(
` Delete the given records, and cascade-delete their corresponding
alias. `,
func(rs m.MailAliasMixinSet)  {
//        aliases = self.mapped('alias_id')
//        res = super(AliasMixin, self).unlink()
//        aliases.unlink()
//        return res
})
h.MailAliasMixin().Methods().InitColumn().DeclareMethod(
` Create aliases for existing rows. `,
func(rs m.MailAliasMixinSet, name interface{})  {
//        super(AliasMixin, self)._init_column(name)
//        if name != 'alias_id':
//            return
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
//            record.with_context({'mail_notrack': True}).alias_id = alias
//            _logger.info('Mail alias created for %s %s (id %s)',
//                         record._name, record.display_name, record.id)
})
}