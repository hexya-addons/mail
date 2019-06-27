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
	
func init() {
h.MailFollowers().DeclareModel()
h.MailFollowers().AddSQLConstraint("mail_followers_res_partner_res_model_id_uniq", "unique(res_model,res_id,partner_id)", "Error, a partner cannot follow twice the same object.")
h.MailFollowers().AddSQLConstraint("mail_followers_res_channel_res_model_id_uniq", "unique(res_model,res_id,channel_id)", "Error, a channel cannot follow twice the same object.")
h.MailFollowers().AddSQLConstraint("partner_xor_channel", "CHECK((partner_id IS NULL) != (channel_id IS NULL))", "Error: A follower must be either a partner or a channel (but not both).")




h.MailFollowers().AddFields(map[string]models.FieldDefinition{
"ResModel": models.CharField{
String: "Related Document Model",
Required: true,
Index: true,
Help: "Model of the followed resource",
},
"ResId": models.IntegerField{
String: "Related Document ID",
Index: true,
Help: "Id of the followed resource",
},
"PartnerId": models.Many2OneField{
RelationModel: h.Partner(),
String: "Related Partner",
OnDelete: `cascade`,
Index: true,
},
"ChannelId": models.Many2OneField{
RelationModel: h.MailChannel(),
String: "Listener",
OnDelete: `cascade`,
Index: true,
},
"SubtypeIds": models.Many2ManyField{
RelationModel: h.MailMessageSubtype(),
String: "Subtype",
Help: "Message subtypes followed, meaning subtypes that will be" + 
"pushed onto the user's Wall.",
},
})
h.MailFollowers().Methods().AddFollowerCommand().DeclareMethod(
` Please upate me
        :param force: if True, delete existing followers
before creating new one
                      using the subtypes given in the parameters
        `,
func(rs m.MailFollowersSet, res_model interface{}, res_ids interface{}, partner_data interface{}, channel_data interface{}, force interface{})  {
//        force_mode = force or (all(data for data in partner_data.values()) and all(
//            data for data in channel_data.values()))
//        generic = []
//        specific = {}
//        existing = {}  # {res_id: follower_ids}
//        p_exist = {}  # {partner_id: res_ids}
//        c_exist = {}  # {channel_id: res_ids}
//        followers = self.sudo().search([
//            '&',
//            '&', ('res_model', '=', res_model), ('res_id', 'in', res_ids),
//            '|', ('partner_id', 'in', partner_data.keys()), ('channel_id', 'in', channel_data.keys())])
//        if force_mode:
//            followers.unlink()
//        else:
//            for follower in followers:
//                existing.setdefault(follower.res_id, list()).append(follower)
//                if follower.partner_id:
//                    p_exist.setdefault(follower.partner_id.id,
//                                       list()).append(follower.res_id)
//                if follower.channel_id:
//                    c_exist.setdefault(follower.channel_id.id,
//                                       list()).append(follower.res_id)
//        default_subtypes = self.env['mail.message.subtype'].search([
//            ('default', '=', True),
//            '|', ('res_model', '=', res_model), ('res_model', '=', False)])
//        external_default_subtypes = default_subtypes.filtered(
//            lambda subtype: not subtype.internal)
//        if force_mode:
//            employee_pids = self.env['res.users'].sudo().search(
//                [('partner_id', 'in', partner_data.keys()), ('share', '=', False)]).mapped('partner_id').ids
//            for pid, data in partner_data.iteritems():
//                if not data:
//                    if pid not in employee_pids:
//                        partner_data[pid] = external_default_subtypes.ids
//                    else:
//                        partner_data[pid] = default_subtypes.ids
//            for cid, data in channel_data.iteritems():
//                if not data:
//                    channel_data[cid] = default_subtypes.ids
//        gen_new_pids = [pid for pid in partner_data.keys()
//                        if pid not in p_exist]
//        gen_new_cids = [cid for cid in channel_data.keys()
//                        if cid not in c_exist]
//        for pid in gen_new_pids:
//            generic.append([0, 0, {'res_model': res_model, 'partner_id': pid, 'subtype_ids': [
//                           (6, 0, partner_data.get(pid) or default_subtypes.ids)]}])
//        for cid in gen_new_cids:
//            generic.append([0, 0, {'res_model': res_model, 'channel_id': cid, 'subtype_ids': [
//                           (6, 0, channel_data.get(cid) or default_subtypes.ids)]}])
//        if not force_mode:
//            for res_id in res_ids:
//                command = []
//                doc_followers = existing.get(res_id, list())
//
//                new_pids = set(partner_data.keys(
//                )) - set([sub.partner_id.id for sub in doc_followers if sub.partner_id]) - set(gen_new_pids)
//                new_cids = set(channel_data.keys(
//                )) - set([sub.channel_id.id for sub in doc_followers if sub.channel_id]) - set(gen_new_cids)
//
//                # subscribe new followers
//                for new_pid in new_pids:
//                    command.append((0, 0, {
//                        'res_model': res_model,
//                        'partner_id': new_pid,
//                        'subtype_ids': [(6, 0, partner_data.get(new_pid) or default_subtypes.ids)],
//                    }))
//                for new_cid in new_cids:
//                    command.append((0, 0, {
//                        'res_model': res_model,
//                        'channel_id': new_cid,
//                        'subtype_ids': [(6, 0, channel_data.get(new_cid) or default_subtypes.ids)],
//                    }))
//                if command:
//                    specific[res_id] = command
//        return generic, specific
})
h.MailFollowers().Methods().InvalidateDocuments().DeclareMethod(
` Invalidate the cache of the documents followed by ``self``. `,
func(rs m.MailFollowersSet)  {
//        for record in self:
//            if record.res_id:
//                self.env[record.res_model].invalidate_cache(
//                    ids=[record.res_id])
})
h.MailFollowers().Methods().Create().Extend(
`Create`,
func(rs m.MailFollowersSet, vals models.RecordData)  {
//        res = super(Followers, self).create(vals)
//        res._invalidate_documents()
//        return res
})
h.MailFollowers().Methods().Write().Extend(
`Write`,
func(rs m.MailFollowersSet, vals models.RecordData)  {
//        if 'res_model' in vals or 'res_id' in vals:
//            self._invalidate_documents()
//        res = super(Followers, self).write(vals)
//        self._invalidate_documents()
//        return res
})
h.MailFollowers().Methods().Unlink().Extend(
`Unlink`,
func(rs m.MailFollowersSet)  {
//        self._invalidate_documents()
//        return super(Followers, self).unlink()
})

}