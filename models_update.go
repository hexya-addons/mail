package mail

import (
	"github.com/hexya-erp/pool/h"
)

//import datetime
//import logging
//import werkzeug.urls
//import urllib2
//_logger = logging.getLogger(__name__)
func init() {
	h.Publisher_warrantyContract().DeclareModel()

	h.Publisher_warrantyContract().Methods().GetMessage().DeclareMethod(
		`GetMessage`,
		func(rs m.Publisher_warrantyContractSet) {
			//        Users = self.env['res.users']
			//        IrParamSudo = self.env['ir.config_parameter'].sudo()
			//        dbuuid = IrParamSudo.get_param('database.uuid')
			//        db_create_date = IrParamSudo.get_param('database.create_date')
			//        limit_date = datetime.datetime.now()
			//        limit_date = limit_date - datetime.timedelta(15)
			//        limit_date_str = limit_date.strftime(
			//            misc.DEFAULT_SERVER_DATETIME_FORMAT)
			//        nbr_users = Users.search_count([('active', '=', True)])
			//        nbr_active_users = Users.search_count(
			//            [("login_date", ">=", limit_date_str), ('active', '=', True)])
			//        nbr_share_users = 0
			//        nbr_active_share_users = 0
			//        if "share" in Users._fields:
			//            nbr_share_users = Users.search_count(
			//                [("share", "=", True), ('active', '=', True)])
			//            nbr_active_share_users = Users.search_count(
			//                [("share", "=", True), ("login_date", ">=", limit_date_str), ('active', '=', True)])
			//        user = self.env.user
			//        domain = [('application', '=', True), ('state', 'in',
			//                                               ['installed', 'to upgrade', 'to remove'])]
			//        apps = self.env['ir.module.module'].sudo(
			//        ).search_read(domain, ['name'])
			//        enterprise_code = IrParamSudo.get_param('database.enterprise_code')
			//        web_base_url = IrParamSudo.get_param('web.base.url')
			//        msg = {
			//            "dbuuid": dbuuid,
			//            "nbr_users": nbr_users,
			//            "nbr_active_users": nbr_active_users,
			//            "nbr_share_users": nbr_share_users,
			//            "nbr_active_share_users": nbr_active_share_users,
			//            "dbname": self._cr.dbname,
			//            "db_create_date": db_create_date,
			//            "version": release.version,
			//            "language": user.lang,
			//            "web_base_url": web_base_url,
			//            "apps": [app['name'] for app in apps],
			//            "enterprise_code": enterprise_code,
			//        }
			//        if user.partner_id.company_id:
			//            company_id = user.partner_id.company_id
			//            msg.update(company_id.read(["name", "email", "phone"])[0])
			//        return msg
		})
	h.Publisher_warrantyContract().Methods().GetSysLogs().DeclareMethod(
		`
        Utility method to send a publisher warranty get logs messages.
        `,
		func(rs m.Publisher_warrantyContractSet) {
			//        msg = self._get_message()
			//        arguments = {'arg0': msg, "action": "update"}
			//        arguments_raw = werkzeug.urls.url_encode(arguments)
			//        url = config.get("publisher_warranty_url")
			//        uo = urllib2.urlopen(url, arguments_raw, timeout=30)
			//        try:
			//            submit_result = uo.read()
			//            return literal_eval(submit_result)
			//        finally:
			//            uo.close()
		})
	h.Publisher_warrantyContract().Methods().UpdateNotification().DeclareMethod(
		`
        Send a message to OpenERP's publisher warranty
server to check the
        validity of the contracts, get notifications, etc...

        @param cron_mode: If true, catch all exceptions
(appropriate for usage in a cron).
        @type cron_mode: boolean
        `,
		func(rs m.Publisher_warrantyContractSet, cron_mode interface{}) {
			//        try:
			//            try:
			//                result = self._get_sys_logs()
			//            except Exception:
			//                if cron_mode:   # we don't want to see any stack trace in cron
			//                    return False
			//                _logger.debug(
			//                    "Exception while sending a get logs messages", exc_info=1)
			//                raise UserError(
			//                    _("Error during communication with the publisher warranty server."))
			//            # old behavior based on res.log; now on mail.message, that is not necessarily installed
			//            user = self.env['res.users'].sudo().browse(SUPERUSER_ID)
			//            poster = self.sudo().env.ref('mail.channel_all_employees')
			//            if not (poster and poster.exists()):
			//                if not user.exists():
			//                    return True
			//                poster = user
			//            for message in result["messages"]:
			//                try:
			//                    poster.message_post(body=message, subtype='mt_comment', partner_ids=[
			//                                        user.partner_id.id])
			//                except Exception:
			//                    pass
			//            if result.get('enterprise_info'):
			//                # Update expiration date
			//                self.env['ir.config_parameter'].sudo().set_param(
			//                    'database.expiration_date', result['enterprise_info'].get('expiration_date'), ['base.group_user'])
			//                self.env['ir.config_parameter'].sudo().set_param('database.expiration_reason',
			//                                                                 result['enterprise_info'].get('expiration_reason', 'trial'), ['base.group_system'])
			//                self.env['ir.config_parameter'].sudo().set_param(
			//                    'database.enterprise_code', result['enterprise_info'].get('enterprise_code'), ['base.group_user'])
			//
			//        except Exception:
			//            if cron_mode:
			//                return False    # we don't want to see any stack trace in cron
			//            else:
			//                raise
			//        return True
		})
}
