package mail

func SetUpClass(cls interface{}) {
	//        super(TestMailGroup, cls).setUpClass()
	//        cls.registry('mail.channel')._revert_method(
	//            'message_get_recipient_values')
	//        cls.group_private = cls.env['mail.channel'].with_context({
	//            'mail_create_nolog': True,
	//            'mail_create_nosubscribe': True
	//        }).create({
	//            'name': 'Private',
	//            'public': 'private'}
	//        ).with_context({'mail_create_nosubscribe': False})
}
func TearDownClass(cls interface{}) {
	//        @api.multi
	//        def mail_group_message_get_recipient_values(self, notif_message=None, recipient_ids=None):
	//            return self.env['mail.thread'].message_get_recipient_values(notif_message=notif_message, recipient_ids=recipient_ids)
	//        cls.env['mail.channel']._patch_method(
	//            'message_get_recipient_values', mail_group_message_get_recipient_values)
	//        super(TestMailGroup, cls).tearDownClass()
}
func TestAccessRightsPublic(self interface{}) {
	//        self.group_public.sudo(self.user_public).read()
	//        with self.assertRaises(except_orm):
	//            self.group_pigs.sudo(self.user_public).read()
	//        self.group_private.write(
	//            {'channel_partner_ids': [(4, self.user_public.partner_id.id)]})
	//        self.group_private.sudo(self.user_public).read()
	//        with self.assertRaises(AccessError):
	//            self.env['mail.channel'].sudo(
	//                self.user_public).create({'name': 'Test'})
	//        with self.assertRaises(AccessError):
	//            self.group_public.sudo(self.user_public).write(
	//                {'name': 'Broutouschnouk'})
	//        with self.assertRaises(AccessError):
	//            self.group_public.sudo(self.user_public).unlink()
}
func TestAccessRightsGroups(self interface{}) {
	//        self.group_pigs.sudo(self.user_employee).read()
	//        self.env['mail.channel'].sudo(
	//            self.user_employee).create({'name': 'Test'})
	//        self.group_pigs.sudo(self.user_employee).write({'name': 'modified'})
	//        self.group_pigs.sudo(self.user_employee).unlink()
	//        with self.assertRaises(except_orm):
	//            self.group_private.sudo(self.user_employee).read()
	//        with self.assertRaises(AccessError):
	//            self.group_private.sudo(self.user_employee).write(
	//                {'name': 're-modified'})
}
func TestAccessRightsFollowersKo(self interface{}) {
	//        with self.assertRaises(AccessError):
	//            self.group_private.sudo(self.user_portal).name
}
func TestAccessRightsFollowersPortal(self interface{}) {
	//        self.group_private.write(
	//            {'channel_partner_ids': [(4, self.user_portal.partner_id.id)]})
	//        chell_pigs = self.group_private.sudo(self.user_portal)
	//        trigger_read = chell_pigs.name
	//        for message in chell_pigs.message_ids:
	//            trigger_read = message.subject
	//        for partner in chell_pigs.message_partner_ids:
	//            if partner.id == self.user_portal.partner_id.id:
	//                # Chell can read her own partner record
	//                continue
	//            # TODO Change the except_orm to Warning
	//            with self.assertRaises(except_orm):
	//                trigger_read = partner.name
}
func TestMailGroupNotificationRecipientsGrouped(self interface{}) {
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.domain', 'schlouby.fr')
	//        self.group_private.write({'alias_name': 'Test'})
	//        self.group_private.message_subscribe_users(
	//            [self.user_employee.id, self.user_portal.id])
	//        self.group_private.message_post(
	//            body="Test", message_type='comment', subtype='mt_comment')
	//        sent_emails = self._mails
	//        self.assertEqual(len(sent_emails), 1)
	//        for email in sent_emails:
	//            self.assertEqual(
	//                set(email['email_to']),
	//                set([formataddr((self.user_employee.name, self.user_employee.email)), formataddr((self.user_portal.name, self.user_portal.email))]))
}
func TestMailGroupNotificationRecipientsSeparated(self interface{}) {
	//        self.group_private.write({'alias_name': False})
	//        self.group_private.message_subscribe_users(
	//            [self.user_employee.id, self.user_portal.id])
	//        self.group_private.message_post(
	//            body="Test", message_type='comment', subtype='mt_comment')
	//        sent_emails = self._mails
	//        self.assertEqual(len(sent_emails), 2)
	//        for email in sent_emails:
	//            self.assertIn(
	//                email['email_to'][0],
	//                [formataddr((self.user_employee.name, self.user_employee.email)), formataddr((self.user_portal.name, self.user_portal.email))])
}
