package mail

func TestInviteEmail(self interface{}) {
	//        mail_invite = self.env['mail.wizard.invite'].with_context({
	//            'default_res_model': 'mail.channel',
	//            'default_res_id': self.group_pigs.id
	//        }).sudo(self.user_employee.id).create({
	//            'partner_ids': [(4, self.user_portal.partner_id.id), (4, self.partner_1.id)],
	//            'send_mail': True})
	//        mail_invite.add_followers()
	//        self.assertEqual(self.group_pigs.message_partner_ids,
	//                         self.user_portal.partner_id | self.partner_1,
	//                         'invite wizard: Pigs followers after invite is incorrect, should be Admin + added follower')
	//        self.assertEqual(self.group_pigs.message_follower_ids.mapped('channel_id'),
	//                         self.env['mail.channel'],
	//                         'invite wizard: Pigs followers after invite is incorrect, should not have channels')
	//        self.assertEqual(len(
	//            self._mails), 2, 'invite wizard: sent email number incorrect, should be only for Bert')
	//        self.assertEqual(self._mails[0].get('subject'), 'Invitation to follow Discussion channel: Pigs',
	//                         'invite wizard: subject of invitation email is incorrect')
	//        self.assertEqual(self._mails[1].get('subject'), 'Invitation to follow Discussion channel: Pigs',
	//                         'invite wizard: subject of invitation email is incorrect')
	//        self.assertIn('%s invited you to follow Discussion channel document: Pigs' % self.user_employee.name,
	//                      self._mails[0].get('body'),
	//                      'invite wizard: body of invitation email is incorrect')
	//        self.assertIn('%s invited you to follow Discussion channel document: Pigs' % self.user_employee.name,
	//                      self._mails[1].get('body'),
	//                      'invite wizard: body of invitation email is incorrect')
}
