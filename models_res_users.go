package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {

	h.User().AddFields(map[string]models.FieldDefinition{
		"AliasId": models.Many2OneField{
			RelationModel: h.MailAlias(),
			String:        "Alias",
			OnDelete:      `set null`,
			Required:      false,
			Help: "Email address internally associated with this user. Incoming" +
				"emails will appear in the user's notifications.",
			NoCopy: true,
		},
		"AliasContact": models.SelectionField{
			Selection: types.Selection{
				"everyone":  "Everyone",
				"partners":  "Authenticated Partners",
				"followers": "Followers only",
			},
			String:  "Alias Contact Security",
			Related: `AliasId.AliasContact`,
		},
	})
	h.User().Methods().Init().DeclareMethod(
		` Override of __init__ to add access rights on notification_email_send
            and alias fields. Access rights are disabled
by default, but allowed
            on some specific fields defined in self.SELF_{READ/WRITE}ABLE_FIELDS.
        `,
		func(rs m.UserSet, pool interface{}, cr interface{}) {
			//        init_res = super(Users, self).__init__(pool, cr)
			//        type(self).SELF_WRITEABLE_FIELDS = list(self.SELF_WRITEABLE_FIELDS)
			//        type(self).SELF_WRITEABLE_FIELDS.extend(['notify_email'])
			//        type(self).SELF_READABLE_FIELDS = list(self.SELF_READABLE_FIELDS)
			//        type(self).SELF_READABLE_FIELDS.extend(['notify_email'])
			//        return init_res
		})
	h.User().Methods().Create().Extend(
		`Create`,
		func(rs m.UserSet, values models.RecordData) {
			//        if not values.get('login', False):
			//            action = self.env.ref('base.action_res_users')
			//            msg = _(
			//                "You cannot create a new user from here.\n To create new user please go to configuration panel.")
			//            raise exceptions.RedirectWarning(
			//                msg, action.id, _('Go to the configuration panel'))
			//        user = super(Users, self).create(values)
			//        user._create_welcome_message()
			//        return user
		})
	h.User().Methods().Write().Extend(
		`Write`,
		func(rs m.UserSet, vals models.RecordData) {
			//        write_res = super(Users, self).write(vals)
			//        sel_groups = [vals[k]
			//                      for k in vals if is_selection_groups(k) and vals[k]]
			//        if vals.get('groups_id'):
			//            # form: {'group_ids': [(3, 10), (3, 3), (4, 10), (4, 3)]} or {'group_ids': [(6, 0, [ids]}
			//            user_group_ids = [command[1]
			//                              for command in vals['groups_id'] if command[0] == 4]
			//            user_group_ids += [id for command in vals['groups_id']
			//                               if command[0] == 6 for id in command[2]]
			//            self.env['mail.channel'].search(
			//                [('group_ids', 'in', user_group_ids)])._subscribe_users()
			//        elif sel_groups:
			//            self.env['mail.channel'].search(
			//                [('group_ids', 'in', sel_groups)])._subscribe_users()
			//        return write_res
		})
	h.User().Methods().CreateWelcomeMessage().DeclareMethod(
		`CreateWelcomeMessage`,
		func(rs m.UserSet) {
			//        self.ensure_one()
			//        if not self.has_group('base.group_user'):
			//            return False
			//        company_name = self.company_id.name if self.company_id else ''
			//        body = _('%s has joined the %s network.') % (self.name, company_name)
			//        return self.partner_id.sudo().message_post(body=body)
		})
	h.User().Methods().MessagePostGetPid().DeclareMethod(
		`MessagePostGetPid`,
		func(rs m.UserSet) {
			//        self.ensure_one()
			//        if 'thread_model' in self.env.context:
			//            self = self.with_context(thread_model='res.users')
			//        return self.partner_id.id
		})
	h.User().Methods().MessagePost().DeclareMethod(
		` Redirect the posting of message on res.users as a private discussion.
            This is done because when giving the context
of Chatter on the
            various mailboxes, we do not have access to
the current partner_id. `,
		func(rs m.UserSet) {
			//        current_pids = []
			//        partner_ids = kwargs.get('partner_ids', [])
			//        user_pid = self._message_post_get_pid()
			//        for partner_id in partner_ids:
			//            if isinstance(partner_id, (list, tuple)) and partner_id[0] == 4 and len(partner_id) == 2:
			//                current_pids.append(partner_id[1])
			//            elif isinstance(partner_id, (list, tuple)) and partner_id[0] == 6 and len(partner_id) == 3:
			//                current_pids.append(partner_id[2])
			//            elif isinstance(partner_id, (int, long)):
			//                current_pids.append(partner_id)
			//        if user_pid not in current_pids:
			//            partner_ids.append(user_pid)
			//        kwargs['partner_ids'] = partner_ids
			//        return self.env['mail.thread'].message_post(**kwargs)
		})
	h.User().Methods().MessageUpdate().DeclareMethod(
		`MessageUpdate`,
		func(rs m.UserSet, msg_dict interface{}, update_vals interface{}) {
			//        return True
		})
	h.User().Methods().MessageSubscribe().DeclareMethod(
		`MessageSubscribe`,
		func(rs m.UserSet, partner_ids interface{}, channel_ids interface{}, subtype_ids interface{}, force interface{}) {
			//        return True
		})
	h.User().Methods().MessagePartnerInfoFromEmails().DeclareMethod(
		`MessagePartnerInfoFromEmails`,
		func(rs m.UserSet, emails interface{}, link_mail interface{}) {
			//        return self.env['mail.thread'].message_partner_info_from_emails(emails, link_mail=link_mail)
		})
	h.User().Methods().MessageGetSuggestedRecipients().DeclareMethod(
		`MessageGetSuggestedRecipients`,
		func(rs m.UserSet) {
			//        return dict((res_id, list()) for res_id in self._ids)
		})

	h.Group().Methods().Write().Extend(
		`Write`,
		func(rs m.GroupSet, vals models.RecordData, context interface{}) {
			//        write_res = super(res_groups_mail_channel, self).write(vals)
			//        if vals.get('users'):
			//            # form: {'group_ids': [(3, 10), (3, 3), (4, 10), (4, 3)]} or {'group_ids': [(6, 0, [ids]}
			//            user_ids = [command[1]
			//                        for command in vals['users'] if command[0] == 4]
			//            user_ids += [id for command in vals['users']
			//                         if command[0] == 6 for id in command[2]]
			//            self.env['mail.channel'].search(
			//                [('group_ids', 'in', self._ids)])._subscribe_users()
			//        return write_res
		})
}
