package mail

func Test00StateEmail(self interface{}) {
	//        """ Test ir.actions.server email type """
	//        email_template = self.env['mail.template'].create({
	//            'name': 'TestTemplate',
	//            'email_from': 'myself@example.com',
	//            'email_to': 'brigitte@example.com',
	//            'partner_to': '%s' % self.test_partner.id,
	//            'model_id': self.res_partner_model.id,
	//            'subject': 'About ${object.name}',
	//            'body_html': '<p>Dear ${object.name}, your parent is ${object.parent_id and object.parent_id.name or "False"}</p>',
	//        })
	//        self.action.write({'state': 'email', 'template_id': email_template.id})
	//        run_res = self.action.with_context(self.context).run()
	//        self.assertFalse(
	//            run_res, 'ir_actions_server: email server action correctly finished should return False')
	//        mail = self.env['mail.mail'].search(
	//            [('subject', '=', 'About TestingPartner')])
	//        self.assertEqual(len(mail), 1, 'ir_actions_server: TODO')
	//        self.assertEqual(mail.body, '<p>Dear TestingPartner, your parent is False</p>',
	//                         'ir_actions_server: TODO')
}
