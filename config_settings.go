package mail

import (
	"github.com/hexya-addons/base/basetypes"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/types/dates"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

var fields_ConfigSettings = map[string]models.FieldDefinition{
	"FailCounter": fields.Integer{
		String:   "Fail Mail",
		ReadOnly: true,
		GoType:   new(int),
	},
	"AliasDomain": fields.Char{
		String: "Alias Domain",
		Help: "If you have setup a catch-all email domain redirected to" +
			"the Odoo server, enter the domain name here.",
	},
}

func configSettings_ConfigFields(rs m.ConfigSettingsSet) basetypes.ConfigFieldsMap {
	res := rs.Super().ConfigFields()
	res[h.ConfigSettings().Fields().AliasDomain()] = "mail.catchall.domain"
	return res
}

func configSettings_GetValues(rs m.ConfigSettingsSet) m.ConfigSettingsData {
	res := rs.Super().GetValues()
	previousDate := dates.Now().AddDate(0, 0, -30)
	res.SetFailCounter(h.MailMail().NewSet(rs.Env()).Sudo().Search(
		q.MailMail().
			Date().GreaterOrEqual(previousDate).
			And().State().Equals("exception")).SearchCount())
	return res
}

func configSettings_SetValues(rs m.ConfigSettingsSet) {
	rs.Super().SetValues()
	h.ConfigParameter().NewSet(rs.Env()).SetParam("mail.catchall.domain", rs.AliasDomain())
}

func init() {
	h.ConfigSettings().AddFields(fields_ConfigSettings)
	h.ConfigSettings().Methods().ConfigFields().Extend(configSettings_ConfigFields)
	h.ConfigSettings().Methods().GetValues().Extend(configSettings_GetValues)
	h.ConfigSettings().Methods().SetValues().Extend(configSettings_SetValues)
}
