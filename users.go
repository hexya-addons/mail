package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

var fields_User = map[string]models.FieldDefinition{
	"Alias": fields.Many2One{
		RelationModel: h.MailAlias(),
		String:        "Alias",
		OnDelete:      models.SetNull,
		Required:      false,
		Help: "Email address internally associated with this user. Incoming" +
			"emails will appear in the user's notifications.",
		NoCopy: true,
	},

	"AliasContact": fields.Selection{
		Selection: types.Selection{
			"everyone":  "Everyone",
			"partners":  "Authenticated Partners",
			"followers": "Followers only",
		},
		String:   "Alias Contact Security",
		Related:  `AliasId.AliasContact`,
		ReadOnly: false},

	"NotificationType": fields.Selection{
		Selection: types.Selection{
			"email": "Handle by Emails",
			"inbox": "Handle in Odoo",
		},
		String:   "Notification Management",
		Required: true,
		Default:  models.DefaultValue("email"),
		Help: "Policy on how to handle Chatter notifications:" +
			"- Handle by Emails: notifications are sent to your email address" +
			"- Handle in Odoo: notifications appear in your Odoo Inbox"},

	"IsModerator": fields.Boolean{
		String:  "Is moderator",
		Compute: h.User().Methods().ComputeIsModerator()},

	"ModerationCounter": fields.Integer{
		String:  "Moderation count",
		Compute: h.User().Methods().ComputeModerationCounter()},

	"ModerationChannels": fields.Many2Many{
		RelationModel:    h.MailChannel(),
		M2MLinkModelName: "",
		String:           "Moderated channels"},
}

// ComputeIsModerator
func user_ComputeIsModerator(rs m.UserSet) m.UserData {
	//        moderated = self.env['mail.channel'].search([
	//            ('id', 'in', self.mapped('moderation_channel_ids').ids),
	//            ('moderation', '=', True),
	//            ('moderator_ids', 'in', self.ids)
	//        ])
	//        user_ids = moderated.mapped('moderator_ids')
	//        for user in self:
	//            user.is_moderator = user in user_ids
}

// ComputeModerationCounter
func user_ComputeModerationCounter(rs m.UserSet) m.UserData {
	//        self._cr.execute("""
	//SELECT channel_moderator.res_users_id, COUNT(msg.id)
	//FROM "mail_channel_moderator_rel" AS channel_moderator
	//JOIN "mail_message" AS msg
	//ON channel_moderator.mail_channel_id = msg.res_id
	//    AND channel_moderator.res_users_id IN %s
	//    AND msg.model = 'mail.channel'
	//    AND msg.moderation_status = 'pending_moderation'
	//GROUP BY channel_moderator.res_users_id""", [tuple(self.ids)])
	//        result = dict(self._cr.fetchall())
	//        for user in self:
	//            user.moderation_counter = result.get(user.id, 0)
}

//  Override of __init__ to add access rights on notification_email_send
//             and alias fields. Access rights are disabled
// by default, but allowed
//             on some specific fields defined in self.SELF_{READ/WRITE}ABLE_FIELDS.
//
func user_Init(rs m.UserSet, pool interface{}, cr interface{}) {
	//        init_res = super(Users, self).__init__(pool, cr)
	//        type(self).SELF_WRITEABLE_FIELDS = list(self.SELF_WRITEABLE_FIELDS)
	//        type(self).SELF_WRITEABLE_FIELDS.extend(['notification_type'])
	//        type(self).SELF_READABLE_FIELDS = list(self.SELF_READABLE_FIELDS)
	//        type(self).SELF_READABLE_FIELDS.extend(['notification_type'])
	//        return init_res
}

// Create
func user_Create(rs m.UserSet, values models.RecordData) {
	//        if not values.get('login', False):
	//            action = self.env.ref('base.action_res_users')
	//            msg = _(
	//                "You cannot create a new user from here.\n To create new user please go to configuration panel.")
	//            raise exceptions.RedirectWarning(
	//                msg, action.id, _('Go to the configuration panel'))
	//        user = super(Users, self).create(values)
	//        self.env['mail.channel'].search(
	//            [('group_ids', 'in', user.groups_id.ids)])._subscribe_users()
	//        return user
}

// Write
func user_Write(rs m.UserSet, vals models.RecordData) {
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
}

// SystrayGetActivities
func user_SystrayGetActivities(rs m.UserSet) {
	//        query = """SELECT m.id, count(*), act.res_model as model,
	//                        CASE
	//                            WHEN %(today)s::date - act.date_deadline::date = 0 Then 'today'
	//                            WHEN %(today)s::date - act.date_deadline::date > 0 Then 'overdue'
	//                            WHEN %(today)s::date - act.date_deadline::date < 0 Then 'planned'
	//                        END AS states
	//                    FROM mail_activity AS act
	//                    JOIN ir_model AS m ON act.res_model_id = m.id
	//                    WHERE user_id = %(user_id)s
	//                    GROUP BY m.id, states, act.res_model;
	//                    """
	//        self.env.cr.execute(query, {
	//            'today': fields.Date.context_today(self),
	//            'user_id': self.env.uid,
	//        })
	//        activity_data = self.env.cr.dictfetchall()
	//        model_ids = [a['id'] for a in activity_data]
	//        model_names = {n[0]: n[1]
	//                       for n in self.env['ir.model'].browse(model_ids).name_get()}
	//        user_activities = {}
	//        for activity in activity_data:
	//            if not user_activities.get(activity['model']):
	//                user_activities[activity['model']] = {
	//                    'name': model_names[activity['id']],
	//                    'model': activity['model'],
	//                    'type': 'activity',
	//                    'icon': modules.module.get_module_icon(self.env[activity['model']]._original_module),
	//                    'total_count': 0, 'today_count': 0, 'overdue_count': 0, 'planned_count': 0,
	//                }
	//            user_activities[activity['model']]['%s_count' %
	//                                               activity['states']] += activity['count']
	//            if activity['states'] in ('today', 'overdue'):
	//                user_activities[activity['model']
	//                                ]['total_count'] += activity['count']
	//        return list(user_activities.values())
}

var fields_Group = map[string]models.FieldDefinition{}

// Write
func group_Write(rs m.GroupSet, vals models.RecordData, context interface{}) {
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
}
func init() {
	models.NewModel("User")
	h.User().AddFields(fields_User)
	h.User().NewMethod("ComputeIsModerator", user_ComputeIsModerator)
	h.User().NewMethod("ComputeModerationCounter", user_ComputeModerationCounter)
	h.User().NewMethod("Init", user_Init)
	h.User().Methods().Create().Extend(user_Create)
	h.User().Methods().Write().Extend(user_Write)
	h.User().NewMethod("SystrayGetActivities", user_SystrayGetActivities)

	models.NewModel("Group")
	h.Group().AddFields(fields_Group)
	h.Group().Methods().Write().Extend(group_Write)

}
