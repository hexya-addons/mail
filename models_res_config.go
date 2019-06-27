package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

//import urlparse
//import datetime
func init() {
	h.BaseConfigSettings().DeclareModel()

	h.BaseConfigSettings().AddFields(map[string]models.FieldDefinition{
		"FailCounter": models.IntegerField{
			String:   "Fail Mail",
			ReadOnly: true,
		},
		"AliasDomain": models.CharField{
			String: "Alias Domain",
			Help: "If you have setup a catch-all email domain redirected to" +
				"the Odoo server, enter the domain name here.",
		},
	})
	h.BaseConfigSettings().Methods().GetDefaultFailCounter().DeclareMethod(
		`GetDefaultFailCounter`,
		func(rs m.BaseConfigSettingsSet, fields interface{}) {
			//        previous_date = datetime.datetime.now() - datetime.timedelta(days=30)
			//        return {
			//            'fail_counter': self.env['mail.mail'].sudo().search_count([('date', '>=', previous_date.strftime(tools.DEFAULT_SERVER_DATETIME_FORMAT)), ('state', '=', 'exception')]),
			//        }
		})
	h.BaseConfigSettings().Methods().GetDefaultAliasDomain().DeclareMethod(
		`GetDefaultAliasDomain`,
		func(rs m.BaseConfigSettingsSet, fields interface{}) {
			//        alias_domain = self.env["ir.config_parameter"].get_param(
			//            "mail.catchall.domain", default=None)
			//        if alias_domain is None:
			//            domain = self.env["ir.config_parameter"].get_param("web.base.url")
			//            try:
			//                alias_domain = urlparse.urlsplit(domain).netloc.split(':')[0]
			//            except Exception:
			//                pass
			//        return {'alias_domain': alias_domain or False}
		})
	h.BaseConfigSettings().Methods().SetAliasDomain().DeclareMethod(
		`SetAliasDomain`,
		func(rs m.BaseConfigSettingsSet) {
			//        for record in self:
			//            self.env['ir.config_parameter'].set_param(
			//                "mail.catchall.domain", record.alias_domain or '')
		})
}
