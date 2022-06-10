package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

var fields_MailBlacklist = map[string]models.FieldDefinition{
	"Email": fields.Char{
		String:   "Email Address",
		Required: true,
		Index:    true,
		Help:     "This field is case insensitive.",
		// track_visibility=True
	},

	"Active": fields.Boolean{
		Default: models.DefaultValue(true),
		// track_visibility=True
	},
}

// Create
func mailBlacklist_Create(rs m.MailBlacklistSet, values models.RecordData) {
	//        new_values = []
	//        all_emails = []
	//        for value in values:
	//            email = self._sanitize_email(value.get('email'))
	//            if not email:
	//                raise UserError(_('Invalid email address %r') % value['email'])
	//            if email in all_emails:
	//                continue
	//            all_emails.append(email)
	//            new_value = dict(value, email=email)
	//            new_values.append(new_value)
	//        """ To avoid crash during import due to unique email, return the existing records if any """
	//        sql = '''SELECT email, id FROM mail_blacklist WHERE email = ANY(%s)'''
	//        emails = [v['email'] for v in new_values]
	//        self._cr.execute(sql, (emails))
	//        bl_entries = dict(self._cr.fetchall())
	//        to_create = [v for v in new_values
	//                     if v['email'] not in bl_entries]
	//        results = super(MailBlackList, self).create(to_create)
	//        return self.env['mail.blacklist'].browse(bl_entries.values()) | results
}

// Write
func mailBlacklist_Write(rs m.MailBlacklistSet, values models.RecordData) {
	//        if 'email' in values:
	//            values['email'] = self._sanitize_email(values['email'])
	//        return super(MailBlackList, self).write(values)
}

//  Override _search in order to grep search on email field and make it
//         lower-case and sanitized
func mailBlacklist_Search(rs m.MailBlacklistSet, args model_mixin.Condition, offset interface{}, limit interface{}, order interface{}, count interface{}, access_rights_uid interface{}) {
	//        if args:
	//            new_args = []
	//            for arg in args:
	//                if isinstance(arg, (list, tuple)) and arg[0] == 'email' and isinstance(arg[2], tools.pycompat.text_type):
	//                    sanitized = self.env['mail.blacklist']._sanitize_email(
	//                        arg[2])
	//                    if sanitized:
	//                        new_args.append([arg[0], arg[1], sanitized])
	//                    else:
	//                        new_args.append(arg)
	//                else:
	//                    new_args.append(arg)
	//        else:
	//            new_args = args
	//        return super(MailBlackList, self)._search(new_args, offset=offset, limit=limit, order=order, count=count, access_rights_uid=access_rights_uid)
}

// Add
func mailBlacklist_Add(rs m.MailBlacklistSet, email interface{}) {
	//        sanitized = self._sanitize_email(email)
	//        record = self.env["mail.blacklist"].with_context(
	//            active_test=False).search([('email', '=', sanitized)])
	//        if len(record) > 0:
	//            record.write({'active': True})
	//        else:
	//            record = self.create({'email': email})
	//        return record
}

// Remove
func mailBlacklist_Remove(rs m.MailBlacklistSet, email interface{}) {
	//        sanitized = self._sanitize_email(email)
	//        record = self.env["mail.blacklist"].with_context(
	//            active_test=False).search([('email', '=', sanitized)])
	//        if len(record) > 0:
	//            record.write({'active': False})
	//        else:
	//            record = record.create({'email': email, 'active': False})
	//        return record
}

//  Sanitize and standardize blacklist entries: all emails should be
//         only real email extracted from strings (A <a@a>
// -> a@a)  and should be
//         lower case.
func mailBlacklist_SanitizeEmail(rs m.MailBlacklistSet, email interface{}) {
	//        emails = tools.email_split(email)
	//        if not emails or len(emails) != 1:
	//            return False
	//        return emails[0].lower()
}

var fields_MailBlacklistMixin = map[string]models.FieldDefinition{
	"IsBlacklisted": fields.Boolean{
		String:  "Blacklist",
		Compute: h.MailBlacklistMixin().Methods().ComputeIsBlacklisted(),
		// compute_sudo=True
		Stored: false,
		// search="_search_is_blacklisted"
		// groups="base.group_user"
		Help: "If the email address is on the blacklist, the contact won't" +
			"receive mass mailing anymore, from any list"},
}

// AssertPrimaryEmail
func mailBlacklistMixin_AssertPrimaryEmail(rs m.MailBlacklistMixinSet) {
	//        if not hasattr(self, "_primary_email") or \
	//                not isinstance(self._primary_email, (list, tuple)) or \
	//                not len(self._primary_email) == 1:
	//            raise UserError(
	//                _('Invalid primary email field on model %s') % self._name)
	//        field_name = self._primary_email[0]
	//        if field_name not in self._fields or self._fields[field_name].type != 'char':
	//            raise UserError(
	//                _('Invalid primary email field on model %s') % self._name)
}

// SearchIsBlacklisted
func mailBlacklistMixin_SearchIsBlacklisted(rs m.MailBlacklistMixinSet, operator interface{}, value interface{}) {
	//        self._assert_primary_email()
	//        if operator != '=':
	//            if operator == '!=' and isinstance(value, bool):
	//                value = not value
	//            else:
	//                raise NotImplementedError()
	//        [email_field] = self._primary_email
	//        if value:
	//            query = """
	//                SELECT m.id
	//                    FROM mail_blacklist bl
	//                    JOIN %s m
	//                    ON (LOWER(substring(m.%s, '([^ ,;<@]+@[^> ,;]+)')) = bl.email AND bl.active)
	//            """
	//        else:
	//            query = """
	//                SELECT m.id
	//                    FROM %s m
	//                    LEFT JOIN mail_blacklist bl
	//                    ON (LOWER(substring(m.%s, '([^ ,;<@]+@[^> ,;]+)')) = bl.email AND bl.active)
	//                    WHERE bl.id IS NULL
	//            """
	//        self._cr.execute(query % (self._table, email_field))
	//        res = self._cr.fetchall()
	//        if not res:
	//            return [(0, '=', 1)]
	//        return [('id', 'in', [r[0] for r in res])]
}

// ComputeIsBlacklisted
func mailBlacklistMixin_ComputeIsBlacklisted(rs m.MailBlacklistMixinSet) m.MailBlacklistMixinData {
	//        self._assert_primary_email()
	//        [email_field] = self._primary_email
	//        sanitized = [self.env['mail.blacklist']._sanitize_email(
	//            email) for email in self.mapped(email_field)]
	//        blacklist = set(self.env['mail.blacklist'].sudo().search(
	//            [('email', 'in', sanitized)]).mapped('email'))
	//        for record in self:
	//            record.is_blacklisted = self.env['mail.blacklist']._sanitize_email(
	//                record[email_field]) in blacklist
}
func init() {
	models.NewModel("MailBlacklist")
	h.MailBlacklist().AddFields(fields_MailBlacklist)
	h.MailBlacklist().Methods().Create().Extend(mailBlacklist_Create)
	h.MailBlacklist().Methods().Write().Extend(mailBlacklist_Write)
	h.MailBlacklist().Methods().Search().Extend(mailBlacklist_Search)
	h.MailBlacklist().NewMethod("Add", mailBlacklist_Add)
	h.MailBlacklist().NewMethod("Remove", mailBlacklist_Remove)
	h.MailBlacklist().NewMethod("SanitizeEmail", mailBlacklist_SanitizeEmail)

	models.NewModel("MailBlacklistMixin")
	h.MailBlacklistMixin().AddFields(fields_MailBlacklistMixin)
	h.MailBlacklistMixin().NewMethod("AssertPrimaryEmail", mailBlacklistMixin_AssertPrimaryEmail)
	h.MailBlacklistMixin().NewMethod("SearchIsBlacklisted", mailBlacklistMixin_SearchIsBlacklisted)
	h.MailBlacklistMixin().NewMethod("ComputeIsBlacklisted", mailBlacklistMixin_ComputeIsBlacklisted)

}
