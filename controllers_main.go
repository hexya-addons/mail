package mail

import (
	"net/http"

	"github.com/hexya-erp/hexya/src/controllers"
)

//import base64
//import json
//import logging
//import psycopg2
//import werkzeug
//_logger = logging.getLogger(__name__)
//    _cp_path = '/mail'
func RedirectToMessaging(cls interface{}) {
	//        messaging_action = request.env['mail.thread']._get_inbox_action_xml_id(
	//        )
	//        url = '/web#%s' % url_encode({'action': messaging_action})
	//        return werkzeug.utils.redirect(url)
}
func CheckToken(cls interface{}, token interface{}) {
	//        base_link = request.httprequest.path
	//        params = dict(request.params)
	//        params.pop('token', '')
	//        valid_token = request.env['mail.thread']._generate_notification_token(
	//            base_link, params)
	//        return consteq(valid_token, str(token))
}
func CheckTokenAndRecordOrRedirect(cls interface{}, model interface{}, res_id interface{}, token interface{}) {
	//        comparison = cls._check_token(token)
	//        if not comparison:
	//            _logger.warning(_('Invalid token in route %s') %
	//                            request.httprequest.url)
	//            return comparison, None, cls._redirect_to_messaging()
	//        try:
	//            record = request.env[model].browse(res_id).exists()
	//        except Exception:
	//            record = None
	//            redirect = cls._redirect_to_messaging()
	//        else:
	//            redirect = cls._redirect_to_record(model, res_id)
	//        return comparison, record, redirect
}
func RedirectToRecord(cls interface{}, model interface{}, res_id interface{}) {
	//        uid = request.session.uid
	//        if not model or not res_id or model not in request.env:
	//            return cls._redirect_to_messaging()
	//        RecordModel = request.env[model]
	//        record_sudo = RecordModel.sudo().browse(res_id).exists()
	//        if not record_sudo:
	//            # record does not seem to exist -> redirect to login
	//            return cls._redirect_to_messaging()
	//        record_action = record_sudo.get_access_action()
	//        record_target_type = record_action.pop('target_type', 'dummy')
	//        if record_action['type'] == 'ir.actions.act_url':
	//            if record_target_type == 'public' and not uid:
	//                return werkzeug.utils.redirect(record_action['url'])
	//            else:
	//                # user connected or non-public URL, handled below
	//                pass
	//        # other choice: act_window (no support of anything else currently)
	//        elif not record_action['type'] == 'ir.actions.act_window':
	//            return cls._redirect_to_messaging()
	//        if uid is not None:
	//            if not RecordModel.sudo(uid).check_access_rights('read', raise_exception=False):
	//                return cls._redirect_to_messaging()
	//            try:
	//                record_sudo.sudo(uid).check_access_rule('read')
	//            except AccessError:
	//                return cls._redirect_to_messaging()
	//            if record_action['type'] == 'ir.actions.act_url':
	//                return werkzeug.utils.redirect(record_action['url'])
	//        else:
	//            # Specific case in 10.0 only: not logged users could receive an act_url that is
	//            # not public. As we don't handle fully access tokens in 10.0 we have to redirect
	//            # to the login to avoid access issues and/or crash when computing url_params.
	//            # CHS-note: do not forward-port me as in saas-16 it is already managed
	//            if record_action['type'] == 'ir.actions.act_url':
	//                return cls._redirect_to_messaging()
	//        url_params = {
	//            'view_type': record_action['view_type'],
	//            'model': model,
	//            'id': res_id,
	//            'active_id': res_id,
	//            'view_id': record_sudo.get_formview_id(),
	//            'action': record_action.get('id'),
	//        }
	//        url = '/web?#%s' % url_encode(url_params)
	//        return werkzeug.utils.redirect(url)
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/receive") {
		mail.ExtendController(http.MethodPost, "/receive", Receive)
	} else {
		mail.AddController(http.MethodPost, "/receive", Receive)
	}
}
func Receive(self interface{}, req interface{}) {
	//        """ End-point to receive mail from an external SMTP server. """
	//        dbs = req.jsonrequest.get('databases')
	//        for db in dbs:
	//            message = dbs[db].decode('base64')
	//            try:
	//                db_registry = registry(db)
	//                with db_registry.cursor() as cr:
	//                    env = api.Environment(cr, SUPERUSER_ID, {})
	//                    env['mail.thread'].message_process(None, message)
	//            except psycopg2.Error:
	//                pass
	//        return True
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/read_followers") {
		mail.ExtendController(http.MethodPost, "/read_followers", ReadFollowers)
	} else {
		mail.AddController(http.MethodPost, "/read_followers", ReadFollowers)
	}
}
func ReadFollowers(self interface{}, follower_ids interface{}, res_model interface{}) {
	//        followers = []
	//        is_editable = request.env.user.has_group('base.group_no_one')
	//        partner_id = request.env.user.partner_id
	//        follower_id = None
	//        follower_recs = request.env['mail.followers'].sudo().browse(
	//            follower_ids)
	//        res_ids = follower_recs.mapped('res_id')
	//        request.env[res_model].browse(res_ids).check_access_rule("read")
	//        for follower in follower_recs:
	//            is_uid = partner_id == follower.partner_id
	//            follower_id = follower.id if is_uid else follower_id
	//            followers.append({
	//                'id': follower.id,
	//                'name': follower.partner_id.name or follower.channel_id.name,
	//                'email': follower.partner_id.email if follower.partner_id else None,
	//                'res_model': 'res.partner' if follower.partner_id else 'mail.channel',
	//                'res_id': follower.partner_id.id or follower.channel_id.id,
	//                'is_editable': is_editable,
	//                'is_uid': is_uid,
	//            })
	//        return {
	//            'followers': followers,
	//            'subtypes': self.read_subscription_data(res_model, follower_id) if follower_id else None
	//        }
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/read_subscription_data") {
		mail.ExtendController(http.MethodPost, "/read_subscription_data", ReadSubscriptionData)
	} else {
		mail.AddController(http.MethodPost, "/read_subscription_data", ReadSubscriptionData)
	}
}
func ReadSubscriptionData(self interface{}, res_model interface{}, follower_id interface{}) {
	//        """ Computes:
	//            - message_subtype_data: data about document subtypes: which are
	//                available, which are followed if any """
	//        followers = request.env['mail.followers'].browse(follower_id)
	//        subtypes = request.env['mail.message.subtype'].search(
	//            ['&', ('hidden', '=', False), '|', ('res_model', '=', res_model), ('res_model', '=', False)])
	//        subtypes_list = [{
	//            'name': subtype.name,
	//            'res_model': subtype.res_model,
	//            'sequence': subtype.sequence,
	//            'default': subtype.default,
	//            'internal': subtype.internal,
	//            'followed': subtype.id in followers.mapped('subtype_ids').ids,
	//            'parent_model': subtype.parent_id and subtype.parent_id.res_model or False,
	//            'id': subtype.id
	//        } for subtype in subtypes]
	//        subtypes_list = sorted(subtypes_list, key=itemgetter(
	//            'parent_model', 'res_model', 'internal', 'sequence'))
	//        return subtypes_list
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/view") {
		mail.ExtendController(http.MethodPost, "/view", MailActionView)
	} else {
		mail.AddController(http.MethodPost, "/view", MailActionView)
	}
}
func MailActionView(self interface{}, model interface{}, res_id interface{}, message_id interface{}) {
	//        """ Generic access point from notification emails. The heuristic to
	//        choose where to redirect the user is the following :
	//
	//         - find a public URL
	//         - if none found
	//          - users with a read access are redirected to the document
	//          - users without read access are redirected to the Messaging
	//          - not logged users are redirected to the login page
	//        """
	//        if message_id:
	//            try:
	//                message = request.env['mail.message'].sudo().browse(
	//                    int(message_id)).exists()
	//            except:
	//                message = request.env['mail.message']
	//            if message:
	//                model, res_id = message.model, message.res_id
	//            else:
	//                # either a wrong message_id, either someone trying ids -> just go to messaging
	//                return self._redirect_to_messaging()
	//        elif res_id and isinstance(res_id, basestring):
	//            res_id = int(res_id)
	//        return self._redirect_to_record(model, res_id)
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/follow") {
		mail.ExtendController(http.MethodPost, "/follow", MailActionFollow)
	} else {
		mail.AddController(http.MethodPost, "/follow", MailActionFollow)
	}
}
func MailActionFollow(self interface{}, model interface{}, res_id interface{}, token interface{}) {
	//        comparison, record, redirect = self._check_token_and_record_or_redirect(
	//            model, int(res_id), token)
	//        if comparison and record:
	//            try:
	//                record.sudo().message_subscribe_users()
	//            except Exception:
	//                return self._redirect_to_messaging()
	//        return redirect
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/unfollow") {
		mail.ExtendController(http.MethodPost, "/unfollow", MailActionUnfollow)
	} else {
		mail.AddController(http.MethodPost, "/unfollow", MailActionUnfollow)
	}
}
func MailActionUnfollow(self interface{}, model interface{}, res_id interface{}, token interface{}) {
	//        comparison, record, redirect = self._check_token_and_record_or_redirect(
	//            model, int(res_id), token)
	//        if comparison and record:
	//            try:
	//                # TDE CHECKME: is sudo really necessary ?
	//                record.sudo().message_unsubscribe_users([request.uid])
	//            except Exception:
	//                return self._redirect_to_messaging()
	//        return redirect
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/new") {
		mail.ExtendController(http.MethodPost, "/new", MailActionNew)
	} else {
		mail.AddController(http.MethodPost, "/new", MailActionNew)
	}
}
func MailActionNew(self interface{}, model interface{}, res_id interface{}, action_id interface{}) {
	//        if model not in request.env:
	//            return self._redirect_to_messaging()
	//        params = {'view_type': 'form', 'model': model}
	//        if action_id:
	//            # Probably something to do
	//            params['action'] = action_id
	//        return werkzeug.utils.redirect('/web?#%s' % url_encode(params))
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/assign") {
		mail.ExtendController(http.MethodPost, "/assign", MailActionAssign)
	} else {
		mail.AddController(http.MethodPost, "/assign", MailActionAssign)
	}
}
func MailActionAssign(self interface{}, model interface{}, res_id interface{}, token interface{}) {
	//        comparison, record, redirect = self._check_token_and_record_or_redirect(
	//            model, int(res_id), token)
	//        if comparison and record:
	//            try:
	//                record.write({'user_id': request.uid})
	//            except Exception:
	//                return self._redirect_to_messaging()
	//        return redirect
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	var res_model *controllers.Group
	res_model, ok = mail.GetGroup("/:res_model")
	if !ok {
		res_model = mail.AddGroup("/:res_model")
	}
	var res_id *controllers.Group
	res_id, ok = res_model.GetGroup("/:res_id")
	if !ok {
		res_id = res_model.AddGroup("/:res_id")
	}
	var avatar *controllers.Group
	avatar, ok = res_id.GetGroup("/avatar")
	if !ok {
		avatar = res_id.AddGroup("/avatar")
	}
	if avatar.HasController(http.MethodGet, "/:partner_id") {
		avatar.ExtendController(http.MethodPost, "/:partner_id", Avatar)
	} else {
		avatar.AddController(http.MethodPost, "/:partner_id", Avatar)
	}
}
func Avatar(self interface{}, res_model interface{}, res_id interface{}, partner_id interface{}) {
	//        headers = [('Content-Type', 'image/png')]
	//        status = 200
	//        content = 'R0lGODlhAQABAIABAP///wAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=='
	//        if res_model in request.env:
	//            try:
	//                # if the current user has access to the document, get the partner avatar as sudo()
	//                request.env[res_model].browse(res_id).check_access_rule('read')
	//                if partner_id in request.env[res_model].browse(res_id).sudo().exists().message_ids.mapped('author_id').ids:
	//                    status, headers, _content = binary_content(
	//                        model='res.partner', id=partner_id, field='image_medium', default_mimetype='image/png', env=request.env(user=SUPERUSER_ID))
	//                    # binary content return an empty string and not a placeholder if obj[field] is False
	//                    if _content != '':
	//                        content = _content
	//                    if status == 304:
	//                        return werkzeug.wrappers.Response(status=304)
	//            except AccessError:
	//                pass
	//        image_base64 = base64.b64decode(content)
	//        headers.append(('Content-Length', len(image_base64)))
	//        response = request.make_response(image_base64, headers)
	//        response.status = str(status)
	//        return response
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/needaction") {
		mail.ExtendController(http.MethodPost, "/needaction", Needaction)
	} else {
		mail.AddController(http.MethodPost, "/needaction", Needaction)
	}
}
func Needaction(self interface{}) {
	//        return request.env['res.partner'].get_needaction_count()
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/client_action") {
		mail.ExtendController(http.MethodPost, "/client_action", MailClientAction)
	} else {
		mail.AddController(http.MethodPost, "/client_action", MailClientAction)
	}
}
func MailClientAction(self interface{}) {
	//        values = {
	//            'needaction_inbox_counter': request.env['res.partner'].get_needaction_count(),
	//            'starred_counter': request.env['res.partner'].get_starred_count(),
	//            'channel_slots': request.env['mail.channel'].channel_fetch_slot(),
	//            'commands': request.env['mail.channel'].get_mention_commands(),
	//            'mention_partner_suggestions': request.env['res.partner'].get_static_mention_suggestions(),
	//            'shortcodes': request.env['mail.shortcode'].sudo().search_read([], ['shortcode_type', 'source', 'substitution', 'description']),
	//            'menu_id': request.env['ir.model.data'].xmlid_to_res_id('mail.mail_channel_menu_root_chat'),
	//        }
	//        return values
}
