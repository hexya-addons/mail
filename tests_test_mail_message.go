package mail

//import itertools
func SetUp2(self interface{}) {
	//        super(TestMailMessage, self).setUp()
	//        self.group_private = self.env['mail.channel'].with_context({
	//            'mail_create_nolog': True,
	//            'mail_create_nosubscribe': True
	//        }).create({
	//            'name': 'Private',
	//            'public': 'private'}
	//        ).with_context({'mail_create_nosubscribe': False})
	//        self.message = self.env['mail.message'].create({
	//            'body': 'My Body',
	//            'model': 'mail.channel',
	//            'res_id': self.group_private.id,
	//        })
}
func TestMailMessageValuesBasic(self interface{}) {
	//        self.env['ir.config_parameter'].search(
	//            [('key', '=', 'mail.catchall.domain')]).unlink()
	//        msg = self.env['mail.message'].sudo(self.user_employee).create({
	//            'reply_to': 'test.reply@example.com',
	//            'email_from': 'test.from@example.com',
	//        })
	//        self.assertIn('-private', msg.message_id.split('@')
	//                      [0], 'mail_message: message_id for a void message should be a "private" one')
	//        self.assertEqual(msg.reply_to, 'test.reply@example.com')
	//        self.assertEqual(msg.email_from, 'test.from@example.com')
}
func TestMailMessageValuesDefault(self interface{}) {
	//        self.env['ir.config_parameter'].search(
	//            [('key', '=', 'mail.catchall.domain')]).unlink()
	//        msg = self.env['mail.message'].sudo(self.user_employee).create({})
	//        self.assertIn('-private', msg.message_id.split('@')
	//                      [0], 'mail_message: message_id for a void message should be a "private" one')
	//        self.assertEqual(msg.reply_to, '%s <%s>' %
	//                         (self.user_employee.name, self.user_employee.email))
	//        self.assertEqual(msg.email_from, '%s <%s>' %
	//                         (self.user_employee.name, self.user_employee.email))
}
func TestMailMessageValuesAlias(self interface{}) {
	//        alias_domain = 'example.com'
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.domain', alias_domain)
	//        self.env['ir.config_parameter'].search(
	//            [('key', '=', 'mail.catchall.alias')]).unlink()
	//        msg = self.env['mail.message'].sudo(self.user_employee).create({})
	//        self.assertIn('-private', msg.message_id.split('@')
	//                      [0], 'mail_message: message_id for a void message should be a "private" one')
	//        self.assertEqual(msg.reply_to, '%s <%s>' %
	//                         (self.user_employee.name, self.user_employee.email))
	//        self.assertEqual(msg.email_from, '%s <%s>' %
	//                         (self.user_employee.name, self.user_employee.email))
}
func TestMailMessageValuesAliasCatchall(self interface{}) {
	//        alias_domain = 'example.com'
	//        alias_catchall = 'pokemon'
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.domain', alias_domain)
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.alias', alias_catchall)
	//        msg = self.env['mail.message'].sudo(self.user_employee).create({})
	//        self.assertIn('-private', msg.message_id.split('@')
	//                      [0], 'mail_message: message_id for a void message should be a "private" one')
	//        self.assertEqual(msg.reply_to, '%s <%s@%s>' % (
	//            self.env.user.company_id.name, alias_catchall, alias_domain))
	//        self.assertEqual(msg.email_from, '%s <%s>' %
	//                         (self.user_employee.name, self.user_employee.email))
}
func TestMailMessageValuesDocumentNoAlias(self interface{}) {
	//        self.env['ir.config_parameter'].search(
	//            [('key', '=', 'mail.catchall.domain')]).unlink()
	//        msg = self.env['mail.message'].sudo(self.user_employee).create({
	//            'model': 'mail.channel',
	//            'res_id': self.group_pigs.id
	//        })
	//        self.assertIn('-openerp-%d-mail.channel' % self.group_pigs.id, msg.message_id.split(
	//            '@')[0], 'mail_message: message_id for a void message should be a "private" one')
	//        self.assertEqual(msg.reply_to, '%s <%s>' %
	//                         (self.user_employee.name, self.user_employee.email))
	//        self.assertEqual(msg.email_from, '%s <%s>' %
	//                         (self.user_employee.name, self.user_employee.email))
}
func TestMailMessageValuesDocumentAlias(self interface{}) {
	//        alias_domain = 'example.com'
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.domain', alias_domain)
	//        self.env['ir.config_parameter'].search(
	//            [('key', '=', 'mail.catchall.alias')]).unlink()
	//        msg = self.env['mail.message'].sudo(self.user_employee).create({
	//            'model': 'mail.channel',
	//            'res_id': self.group_pigs.id
	//        })
	//        self.assertIn('-openerp-%d-mail.channel' % self.group_pigs.id, msg.message_id.split(
	//            '@')[0], 'mail_message: message_id for a void message should be a "private" one')
	//        self.assertEqual(msg.reply_to, '%s %s <%s@%s>' % (
	//            self.env.user.company_id.name, self.group_pigs.name, self.group_pigs.alias_name, alias_domain))
	//        self.assertEqual(msg.email_from, '%s <%s>' %
	//                         (self.user_employee.name, self.user_employee.email))
}
func TestMailMessageValuesDocumentAliasCatchall(self interface{}) {
	//        alias_domain = 'example.com'
	//        alias_catchall = 'pokemon'
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.domain', alias_domain)
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.alias', alias_catchall)
	//        msg = self.env['mail.message'].sudo(self.user_employee).create({
	//            'model': 'mail.channel',
	//            'res_id': self.group_pigs.id
	//        })
	//        self.assertIn('-openerp-%d-mail.channel' % self.group_pigs.id, msg.message_id.split(
	//            '@')[0], 'mail_message: message_id for a void message should be a "private" one')
	//        self.assertEqual(msg.reply_to, '%s %s <%s@%s>' % (
	//            self.env.user.company_id.name, self.group_pigs.name, self.group_pigs.alias_name, alias_domain))
	//        self.assertEqual(msg.email_from, '%s <%s>' %
	//                         (self.user_employee.name, self.user_employee.email))
}
func TestMailMessageValuesNoAutoThread(self interface{}) {
	//        msg = self.env['mail.message'].sudo(self.user_employee).create({
	//            'model': 'mail.channel',
	//            'res_id': self.group_pigs.id,
	//            'no_auto_thread': True,
	//        })
	//        self.assertIn('reply_to', msg.message_id.split('@')[0])
	//        self.assertNotIn('mail.channel', msg.message_id.split('@')[0])
	//        self.assertNotIn('-%d-' % self.group_pigs.id,
	//                         msg.message_id.split('@')[0])
}
func TestMailMessageNotifyFromMailMail(self interface{}) {
	//        self.email_to_list = []
	//        mail = self.env['mail.mail'].create({
	//            'body_html': '<p>Test</p>',
	//            'email_to': 'test@example.com',
	//            'partner_ids': [(4, self.user_employee.partner_id.id)]
	//        })
	//        self.email_to_list.extend(itertools.chain.from_iterable(
	//            sent_email['email_to'] for sent_email in self._mails if sent_email.get('email_to')))
	//        self.assertNotIn(u'Ernest Employee <e.e@example.com>',
	//                         self.email_to_list)
	//        mail.send()
	//        self.email_to_list.extend(itertools.chain.from_iterable(
	//            sent_email['email_to'] for sent_email in self._mails if sent_email.get('email_to')))
	//        self.assertNotIn(u'Ernest Employee <e.e@example.com>',
	//                         self.email_to_list)
	//        self.assertIn(u'test@example.com', self.email_to_list)
}
func TestMailMessageAccessSearch(self interface{}) {
	//        msg1 = self.env['mail.message'].create({
	//            'subject': '_Test', 'body': 'A', 'subtype_id': self.ref('mail.mt_comment')})
	//        msg2 = self.env['mail.message'].create({
	//            'subject': '_Test', 'body': 'A+B', 'subtype_id': self.ref('mail.mt_comment'),
	//            'partner_ids': [(6, 0, [self.user_public.partner_id.id])]})
	//        msg3 = self.env['mail.message'].create({
	//            'subject': '_Test', 'body': 'A Pigs', 'subtype_id': False,
	//            'model': 'mail.channel', 'res_id': self.group_pigs.id})
	//        msg4 = self.env['mail.message'].create({
	//            'subject': '_Test', 'body': 'A+P Pigs', 'subtype_id': self.ref('mail.mt_comment'),
	//            'model': 'mail.channel', 'res_id': self.group_pigs.id,
	//            'partner_ids': [(6, 0, [self.user_public.partner_id.id])]})
	//        msg5 = self.env['mail.message'].create({
	//            'subject': '_Test', 'body': 'A+E Pigs', 'subtype_id': self.ref('mail.mt_comment'),
	//            'model': 'mail.channel', 'res_id': self.group_pigs.id,
	//            'partner_ids': [(6, 0, [self.user_employee.partner_id.id])]})
	//        msg6 = self.env['mail.message'].create({
	//            'subject': '_Test', 'body': 'A Birds', 'subtype_id': self.ref('mail.mt_comment'),
	//            'model': 'mail.channel', 'res_id': self.group_private.id})
	//        msg7 = self.env['mail.message'].sudo(self.user_employee).create({
	//            'subject': '_Test', 'body': 'B', 'subtype_id': self.ref('mail.mt_comment')})
	//        msg8 = self.env['mail.message'].sudo(self.user_employee).create({
	//            'subject': '_Test', 'body': 'B+E', 'subtype_id': self.ref('mail.mt_comment'),
	//            'partner_ids': [(6, 0, [self.user_employee.partner_id.id])]})
	//        messages = self.env['mail.message'].sudo(
	//            self.user_public).search([('subject', 'like', '_Test')])
	//        self.assertEqual(messages, msg2 | msg4)
	//        messages = self.env['mail.message'].sudo(self.user_employee).search(
	//            [('subject', 'like', '_Test'), ('body', 'ilike', 'A')])
	//        self.assertEqual(messages, msg3 | msg4 | msg5)
	//        messages = self.env['mail.message'].sudo(
	//            self.user_employee).search([('subject', 'like', '_Test')])
	//        self.assertEqual(messages, msg3 | msg4 | msg5 | msg7 | msg8)
	//        messages = self.env['mail.message'].search(
	//            [('subject', 'like', '_Test')])
	//        self.assertEqual(messages, msg1 | msg2 | msg3 |
	//                         msg4 | msg5 | msg6 | msg7 | msg8)
	//        messages = self.env['mail.message'].sudo(
	//            self.user_portal).search([('subject', 'like', '_Test')])
	//        self.assertFalse(messages)
	//        self.group_pigs.write({'public': 'public'})
	//        messages = self.env['mail.message'].sudo(
	//            self.user_portal).search([('subject', 'like', '_Test')])
	//        self.assertEqual(messages, msg4 | msg5)
}
func TestMailMessageAccessReadCrash(self interface{}) {
	//        with self.assertRaises(except_orm):
	//            self.message.sudo(self.user_employee).read()
}
func TestMailMessageAccessReadCrashPortal(self interface{}) {
	//        with self.assertRaises(except_orm):
	//            self.message.sudo(self.user_portal).read(
	//                ['body', 'message_type', 'subtype_id'])
}
func TestMailMessageAccessReadOkPortal(self interface{}) {
	//        self.message.write({'subtype_id': self.ref(
	//            'mail.mt_comment'), 'res_id': self.group_public.id})
	//        self.message.sudo(self.user_portal).read(
	//            ['body', 'message_type', 'subtype_id'])
}
func TestMailMessageAccessReadNotification(self interface{}) {
	//        attachment = self.env['ir.attachment'].create({
	//            'datas': 'My attachment'.encode('base64'),
	//            'name': 'doc.txt',
	//            'datas_fname': 'doc.txt'})
	//        self.message.write({'attachment_ids': [(4, attachment.id)]})
	//        self.message.write(
	//            {'partner_ids': [(4, self.user_employee.partner_id.id)]})
	//        self.message.sudo(self.user_employee).read()
	//        attachment.sudo(self.user_employee).read(['name', 'datas'])
}
func TestMailMessageAccessReadAuthor(self interface{}) {
	//        self.message.write({'author_id': self.user_employee.partner_id.id})
	//        self.message.sudo(self.user_employee).read()
}
func TestMailMessageAccessReadDoc(self interface{}) {
	//        self.message.write(
	//            {'model': 'mail.channel', 'res_id': self.group_public.id})
	//        self.message.sudo(self.user_employee).read()
}
func TestMailMessageAccessCreateCrashPublic(self interface{}) {
	//        with self.assertRaises(AccessError):
	//            self.env['mail.message'].sudo(self.user_public).create(
	//                {'model': 'mail.channel', 'res_id': self.group_pigs.id, 'body': 'Test'})
	//        with self.assertRaises(AccessError):
	//            self.env['mail.message'].sudo(self.user_public).create(
	//                {'model': 'mail.channel', 'res_id': self.group_public.id, 'body': 'Test'})
}
func TestMailMessageAccessCreateCrash(self interface{}) {
	//        with self.assertRaises(except_orm):
	//            self.env['mail.message'].sudo(self.user_employee).create(
	//                {'model': 'mail.channel', 'res_id': self.group_private.id, 'body': 'Test'})
}
func TestMailMessageAccessCreateDoc(self interface{}) {
	//        Message = self.env['mail.message'].sudo(self.user_employee)
	//        Message.create(
	//            {'model': 'mail.channel', 'res_id': self.group_public.id, 'body': 'Test'})
	//        with self.assertRaises(except_orm):
	//            Message.create(
	//                {'model': 'mail.channel', 'res_id': self.group_private.id, 'body': 'Test'})
}
func TestMailMessageAccessCreatePrivate(self interface{}) {
	//        self.env['mail.message'].sudo(
	//            self.user_employee).create({'body': 'Test'})
}
func TestMailMessageAccessCreateReply(self interface{}) {
	//        self.message.write(
	//            {'partner_ids': [(4, self.user_employee.partner_id.id)]})
	//        self.env['mail.message'].sudo(self.user_employee).create(
	//            {'model': 'mail.channel', 'res_id': self.group_private.id, 'body': 'Test', 'parent_id': self.message.id})
}
func TestMessageSetStar(self interface{}) {
	//        msg = self.group_pigs.message_post(body='My Body', subject='1')
	//        msg_emp = self.env['mail.message'].sudo(
	//            self.user_employee).browse(msg.id)
	//        msg.toggle_message_starred()
	//        self.assertTrue(msg.starred)
	//        msg_emp.toggle_message_starred()
	//        self.assertTrue(msg_emp.starred)
	//        msg.toggle_message_starred()
	//        self.assertFalse(msg.starred)
	//        self.assertTrue(msg_emp.starred)
}
func Test60CacheInvalidation(self interface{}) {
	//        msg_cnt = len(self.group_pigs.message_ids)
	//        self.group_pigs.message_post(body='Hi!', subject='test')
	//        self.assertEqual(len(self.group_pigs.message_ids), msg_cnt + 1)
}
