package mail

//import base64
//import datetime
func SetUp3(self interface{}) {
	//        super(TestMailTemplate, self).setUp()
	//        self._attachments = [{
	//            'name': '_Test_First',
	//            'datas_fname':
	//            'first.txt',
	//            'datas': base64.b64encode('My first attachment'),
	//            'res_model': 'res.partner',
	//            'res_id': self.user_admin.partner_id.id
	//        }, {
	//            'name': '_Test_Second',
	//            'datas_fname': 'second.txt',
	//            'datas': base64.b64encode('My second attachment'),
	//            'res_model': 'res.partner',
	//            'res_id': self.user_admin.partner_id.id
	//        }]
	//        self.email_1 = 'test1@example.com'
	//        self.email_2 = 'test2@example.com'
	//        self.email_3 = self.partner_1.email
	//        self.email_template = self.env['mail.template'].create({
	//            'model_id': self.env['ir.model'].search([('model', '=', 'mail.channel')], limit=1).id,
	//            'name': 'Pigs Template',
	//            'subject': '${object.name}',
	//            'body_html': '${object.description}',
	//            'user_signature': False,
	//            'attachment_ids': [(0, 0, self._attachments[0]), (0, 0, self._attachments[1])],
	//            'partner_to': '%s,%s' % (self.partner_2.id, self.user_employee.partner_id.id),
	//            'email_to': '%s, %s' % (self.email_1, self.email_2),
	//            'email_cc': '%s' % self.email_3})
}
func TestComposerTemplateOnchange(self interface{}) {
	//        composer = self.env['mail.compose.message'].with_context({
	//            'default_composition_mode': 'comment',
	//            'default_model': 'mail.channel',
	//            'default_res_id': self.group_pigs.id,
	//            'default_use_template': False,
	//            'default_template_id': False
	//        }).create({'subject': 'Forget me subject', 'body': 'Dummy body'})
	//        values = composer.onchange_template_id(
	//            self.email_template.id, 'comment', 'mail.channel', self.group_pigs.id)['value']
	//        values = composer._convert_to_record(
	//            composer._convert_to_cache(values))
	//        recipients = values['partner_ids']
	//        attachments = values['attachment_ids']
	//        test_recipients = self.env['res.partner'].search([('email', 'in', [
	//                                                         'test1@example.com', 'test2@example.com'])]) | self.partner_1 | self.partner_2 | self.user_employee.partner_id
	//        test_attachments = self.env['ir.attachment'].search(
	//            [('name', 'in', ['_Test_First', '_Test_Second'])])
	//        self.assertEqual(values['subject'], self.group_pigs.name)
	//        self.assertEqual(values['body'], '<p>%s</p>' %
	//                         self.group_pigs.description)
	//        self.assertEqual(recipients, test_recipients)
	//        self.assertEqual(set(recipients.mapped('email')), set(
	//            [self.email_1, self.email_2, self.partner_1.email, self.partner_2.email, self.user_employee.email]))
	//        self.assertEqual(attachments, test_attachments)
	//        self.assertEqual(set(attachments.mapped('res_model')),
	//                         set(['res.partner']))
	//        self.assertEqual(set(attachments.mapped('res_id')),
	//                         set([self.user_admin.partner_id.id]))
}
func TestComposerTemplateSend(self interface{}) {
	//        self.group_pigs.with_context(use_template=False).message_post_with_template(
	//            self.email_template.id, composition_mode='comment')
	//        message = self.group_pigs.message_ids[0]
	//        test_recipients = self.env['res.partner'].search([('email', 'in', [
	//                                                         'test1@example.com', 'test2@example.com'])]) | self.partner_1 | self.partner_2 | self.user_employee.partner_id
	//        self.assertEqual(message.subject, self.group_pigs.name)
	//        self.assertEqual(message.body, '<p>%s</p>' %
	//                         self.group_pigs.description)
	//        self.assertEqual(message.partner_ids, test_recipients)
	//        self.assertEqual(set(message.attachment_ids.mapped(
	//            'res_model')), set(['mail.channel']))
	//        self.assertEqual(set(message.attachment_ids.mapped(
	//            'res_id')), set([self.group_pigs.id]))
}
func TestComposerTemplateMassMailing(self interface{}) {
	//        composer = self.env['mail.compose.message'].with_context({
	//            'default_composition_mode': 'mass_mail',
	//            'default_notify': True,
	//            'default_model': 'mail.channel',
	//            'default_res_id': self.group_pigs.id,
	//            'default_template_id': self.email_template.id,
	//            'active_ids': [self.group_pigs.id, self.group_public.id]
	//        }).create({})
	//        values = composer.onchange_template_id(
	//            self.email_template.id, 'mass_mail', 'mail.channel', self.group_pigs.id)['value']
	//        composer.write(values)
	//        composer.send_mail()
	//        message_1 = self.group_pigs.message_ids[0]
	//        message_2 = self.group_public.message_ids[0]
	//        self.assertEqual(message_1.subject, self.group_pigs.name,
	//                         'mail.message subject on Pigs incorrect')
	//        self.assertEqual(message_2.subject, self.group_public.name,
	//                         'mail.message subject on Bird incorrect')
	//        self.assertIn(self.group_pigs.description, message_1.body,
	//                      'mail.message body on Pigs incorrect')
	//        self.assertIn(self.group_public.description, message_2.body,
	//                      'mail.message body on Bird incorrect')
}
func TestMailTemplate(self interface{}) {
	//        mail_id = self.email_template.send_mail(self.group_pigs.id)
	//        mail = self.env['mail.mail'].browse(mail_id)
	//        self.assertEqual(mail.subject, self.group_pigs.name)
	//        self.assertEqual(mail.email_to, self.email_template.email_to)
	//        self.assertEqual(mail.email_cc, self.email_template.email_cc)
	//        self.assertEqual(mail.recipient_ids, self.partner_2 |
	//                         self.user_employee.partner_id)
}
func TestMessageComposeTemplateSave(self interface{}) {
	//        self.env['mail.compose.message'].with_context(
	//            {'default_composition_mode': 'comment',
	//             'default_model': 'mail.channel',
	//             'default_res_id': self.group_pigs.id,
	//             'active_ids': [self.group_pigs.id, self.group_public.id]
	//             }).create({
	//                 'subject': 'Forget me subject',
	//                 'body': '<p>Dummy body</p>'
	//             }).with_context({'default_model': 'mail.channel'}).save_as_template()
	//        last_template = self.env['mail.template'].search(
	//            [('model', '=', 'mail.channel'), ('subject', '=', 'Forget me subject')], limit=1)
	//        self.assertEqual(last_template.body_html, '<p>Dummy body</p>',
	//                         'email_template incorrect body_html')
}
func TestAddContextAction(self interface{}) {
	//        self.email_template.create_action()
	//        self.assertTrue(bool(self.email_template.ref_ir_act_window))
	//        self.assertTrue(bool(self.email_template.ref_ir_value))
	//        action = self.email_template.ref_ir_act_window
	//        self.assertEqual(action.name, 'Send Mail (%s)' %
	//                         self.email_template.name)
	//        value = self.email_template.ref_ir_value
	//        self.assertEqual(value.name, 'Send Mail (%s)' %
	//                         self.email_template.name)
}
func TestSetScheduledDateOnATemplate(self interface{}) {
	//        self.email_template_in_2_days = self.email_template.copy()
	//        self.email_template_in_2_days.write(
	//            {'scheduled_date': "${(datetime.datetime.now() + relativedelta(days=2)).strftime('%Y-%m-%d %H:%M')}"})
	//        mail_now_id = self.email_template.send_mail(self.group_pigs.id)
	//        mail_in_2_days_id = self.email_template_in_2_days.send_mail(
	//            self.group_pigs.id)
	//        mail_now = self.env['mail.mail'].browse(mail_now_id)
	//        mail_in_2_days = self.env['mail.mail'].browse(mail_in_2_days_id)
	//        self.assertEqual(bool(mail_now.scheduled_date), False)
	//        scheduled_date = datetime.datetime.strptime(
	//            mail_in_2_days.scheduled_date, '%Y-%m-%d %H:%M')
	//        date_in_2_days = datetime.datetime.today() + datetime.timedelta(days=2)
	//        self.assertEqual(scheduled_date.day, date_in_2_days.day)
	//        self.assertEqual(scheduled_date.month, date_in_2_days.month)
	//        self.assertEqual(scheduled_date.year, date_in_2_days.year)
	//        self.env['mail.mail'].process_email_queue(ids=[mail_now.id])
	//        self.assertTrue(len(self._mails) > 0)
	//        self.env['mail.mail'].process_email_queue(ids=[mail_in_2_days.id])
	//        self.assertEqual(mail_in_2_days.state, 'outgoing')
}
