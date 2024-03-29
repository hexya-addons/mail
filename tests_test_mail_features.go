package mail

func TestAliasSetup(self interface{}) {
	//        alias = self.env['mail.alias'].with_context(
	//            alias_model_name='mail.channel').create({'alias_name': 'b4r+_#_R3wl$$'})
	//        self.assertEqual(alias.alias_name, 'b4r+_-_r3wl-',
	//                         'Disallowed chars should be replaced by hyphens')
}
func Test10CacheInvalidation(self interface{}) {
	//        """ Test that creating a mail-thread record does not invalidate the whole cache. """
	//        record = self.env['res.partner'].new({'name': 'Brave New Partner'})
	//        self.assertTrue(record.name)
	//        self.env['res.partner'].create({'name': 'Actual Partner'})
	//        self.assertTrue(record.name)
}
func TestNeedaction(self interface{}) {
	//        na_emp1_base = self.env['mail.message'].sudo(
	//            self.user_employee)._needaction_count(domain=[])
	//        na_emp2_base = self.env['mail.message'].sudo(
	//        )._needaction_count(domain=[])
	//        self.group_pigs.message_post(body='Test', message_type='comment', subtype='mail.mt_comment', partner_ids=[
	//                                     self.user_employee.partner_id.id])
	//        na_emp1_new = self.env['mail.message'].sudo(
	//            self.user_employee)._needaction_count(domain=[])
	//        na_emp2_new = self.env['mail.message'].sudo(
	//        )._needaction_count(domain=[])
	//        self.assertEqual(na_emp1_new, na_emp1_base + 1)
	//        self.assertEqual(na_emp2_new, na_emp2_base)
}
func TestMarkAllAsRead(self interface{}) {
	//        emp_partner = self.user_employee.partner_id.sudo(self.user_employee.id)
	//        group_private = self.env['mail.channel'].with_context({
	//            'mail_create_nolog': True,
	//            'mail_create_nosubscribe': True,
	//            'mail_channel_noautofollow': True,
	//        }).create({
	//            'name': 'Private',
	//            'description': 'Private James R.',
	//            'public': 'private',
	//            'alias_name': 'private',
	//            'alias_contact': 'followers'}
	//        ).with_context({'mail_create_nosubscribe': False})
	//        group_private.message_post(body='Test', message_type='comment',
	//                                   subtype='mail.mt_comment', partner_ids=[emp_partner.id])
	//        emp_partner.env['mail.message'].mark_all_as_read(
	//            channel_ids=[], domain=[])
	//        na_count = emp_partner.get_needaction_count()
	//        self.assertEqual(
	//            na_count, 0, "mark all as read should conclude all needactions")
	//        new_msg = group_private.message_post(
	//            body='Zest', message_type='comment', subtype='mail.mt_comment', partner_ids=[emp_partner.id])
	//        needaction_accessible = len(
	//            emp_partner.env['mail.message'].search([['needaction', '=', True]]))
	//        self.assertEqual(needaction_accessible, 1,
	//                         "a new message to a partner is readable to that partner")
	//        new_msg.sudo().partner_ids = self.env['res.partner']
	//        emp_partner.env['mail.message'].search([['needaction', '=', True]])
	//        needaction_length = len(
	//            emp_partner.env['mail.message'].search([['needaction', '=', True]]))
	//        self.assertEqual(needaction_length, 1,
	//                         "message should still be readable when notified")
	//        na_count = emp_partner.get_needaction_count()
	//        self.assertEqual(
	//            na_count, 1, "message not accessible is currently still counted")
	//        emp_partner.env['mail.message'].mark_all_as_read(
	//            channel_ids=[], domain=[])
	//        na_count = emp_partner.get_needaction_count()
	//        self.assertEqual(
	//            na_count, 0, "mark all read should conclude all needactions even inacessible ones")
}
func TestMarkAllAsReadShare(self interface{}) {
	//        portal_partner = self.user_portal.partner_id.sudo(self.user_portal.id)
	//        self.group_pigs.message_post(body='Test', message_type='comment',
	//                                     subtype='mail.mt_comment', partner_ids=[portal_partner.id])
	//        portal_partner.env['mail.message'].mark_all_as_read(
	//            channel_ids=[], domain=[])
	//        na_count = portal_partner.get_needaction_count()
	//        self.assertEqual(
	//            na_count, 0, "mark all as read should conclude all needactions")
	//        new_msg = self.group_pigs.message_post(
	//            body='Zest', message_type='comment', subtype='mail.mt_comment', partner_ids=[portal_partner.id])
	//        needaction_accessible = len(
	//            portal_partner.env['mail.message'].search([['needaction', '=', True]]))
	//        self.assertEqual(needaction_accessible, 1,
	//                         "a new message to a partner is readable to that partner")
	//        new_msg.sudo().partner_ids = self.env['res.partner']
	//        needaction_length = len(
	//            portal_partner.env['mail.message'].search([['needaction', '=', True]]))
	//        self.assertEqual(needaction_length, 1,
	//                         "message should still be readable when notified")
	//        na_count = portal_partner.get_needaction_count()
	//        self.assertEqual(
	//            na_count, 1, "message not accessible is currently still counted")
	//        portal_partner.env['mail.message'].mark_all_as_read(
	//            channel_ids=[], domain=[])
	//        na_count = portal_partner.get_needaction_count()
	//        self.assertEqual(
	//            na_count, 0, "mark all read should conclude all needactions even inacessible ones")
}
func TestPostNoSubscribeAuthor(self interface{}) {
	//        original = self.group_pigs.message_follower_ids
	//        self.group_pigs.sudo(self.user_employee).with_context({'mail_create_nosubscribe': True}).message_post(
	//            body='Test Body', message_type='comment', subtype='mt_comment')
	//        self.assertEqual(self.group_pigs.message_follower_ids.mapped(
	//            'partner_id'), original.mapped('partner_id'))
	//        self.assertEqual(self.group_pigs.message_follower_ids.mapped(
	//            'channel_id'), original.mapped('channel_id'))
}
func TestPostNoSubscribeRecipients(self interface{}) {
	//        original = self.group_pigs.message_follower_ids
	//        self.group_pigs.sudo(self.user_employee).with_context({'mail_create_nosubscribe': True}).message_post(
	//            body='Test Body', message_type='comment', subtype='mt_comment', partner_ids=[(4, self.partner_1.id), (4, self.partner_2.id)])
	//        self.assertEqual(self.group_pigs.message_follower_ids.mapped(
	//            'partner_id'), original.mapped('partner_id'))
	//        self.assertEqual(self.group_pigs.message_follower_ids.mapped(
	//            'channel_id'), original.mapped('channel_id'))
}
func TestPostSubscribeRecipients(self interface{}) {
	//        original = self.group_pigs.message_follower_ids
	//        self.group_pigs.sudo(self.user_employee).with_context({'mail_create_nosubscribe': True, 'mail_post_autofollow': True}).message_post(
	//            body='Test Body', message_type='comment', subtype='mt_comment', partner_ids=[(4, self.partner_1.id), (4, self.partner_2.id)])
	//        self.assertEqual(self.group_pigs.message_follower_ids.mapped(
	//            'partner_id'), original.mapped('partner_id') | self.partner_1 | self.partner_2)
	//        self.assertEqual(self.group_pigs.message_follower_ids.mapped(
	//            'channel_id'), original.mapped('channel_id'))
}
func TestPostSubscribeRecipientsPartial(self interface{}) {
	//        original = self.group_pigs.message_follower_ids
	//        self.group_pigs.sudo(self.user_employee).with_context({'mail_create_nosubscribe': True, 'mail_post_autofollow': True, 'mail_post_autofollow_partner_ids': [self.partner_2.id]}).message_post(
	//            body='Test Body', message_type='comment', subtype='mt_comment', partner_ids=[(4, self.partner_1.id), (4, self.partner_2.id)])
	//        self.assertEqual(self.group_pigs.message_follower_ids.mapped(
	//            'partner_id'), original.mapped('partner_id') | self.partner_2)
	//        self.assertEqual(self.group_pigs.message_follower_ids.mapped(
	//            'channel_id'), original.mapped('channel_id'))
}
func TestPostNotifications(self interface{}) {
	//        _body, _body_alt = '<p>Test Body</p>', 'Test Body'
	//        _subject = 'Test Subject'
	//        _attachments = [
	//            ('List1', 'My first attachment'),
	//            ('List2', 'My second attachment')
	//        ]
	//        _attach_1 = self.env['ir.attachment'].sudo(self.user_employee).create({
	//            'name': 'Attach1', 'datas_fname': 'Attach1',
	//            'datas': 'bWlncmF0aW9uIHRlc3Q=',
	//            'res_model': 'mail.compose.message', 'res_id': 0})
	//        _attach_2 = self.env['ir.attachment'].sudo(self.user_employee).create({
	//            'name': 'Attach2', 'datas_fname': 'Attach2',
	//            'datas': 'bWlncmF0aW9uIHRlc3Q=',
	//            'res_model': 'mail.compose.message', 'res_id': 0})
	//        self.partner_2.write({'notify_email': 'none'})
	//        self.user_admin.write({'notify_email': 'always'})
	//        self.group_pigs.message_subscribe_users(user_ids=[self.env.user.id])
	//        _domain = 'schlouby.fr'
	//        _catchall = 'test_catchall'
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.domain', _domain)
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.alias', _catchall)
	//        msg = self.group_pigs.sudo(self.user_employee).message_post(
	//            body=_body, subject=_subject, partner_ids=[
	//                self.partner_1.id, self.partner_2.id],
	//            attachment_ids=[_attach_1.id,
	//                            _attach_2.id], attachments=_attachments,
	//            message_type='comment', subtype='mt_comment')
	//        self.assertEqual(msg.subject, _subject)
	//        self.assertEqual(msg.body, _body)
	//        self.assertEqual(msg.partner_ids, self.partner_1 | self.partner_2)
	//        self.assertEqual(msg.needaction_partner_ids,
	//                         self.env.user.partner_id | self.partner_1 | self.partner_2)
	//        self.assertEqual(msg.channel_ids, self.env['mail.channel'])
	//        self.assertEqual(set(msg.attachment_ids.mapped('res_model')), set(['mail.channel']),
	//                         'message_post: all atttachments should be linked to the mail.channel model')
	//        self.assertEqual(set(msg.attachment_ids.mapped('res_id')), set([self.group_pigs.id]),
	//                         'message_post: all atttachments should be linked to the pigs group')
	//        self.assertEqual(set([x.decode('base64') for x in msg.attachment_ids.mapped('datas')]),
	//                         set(['migration test', _attachments[0][1], _attachments[1][1]]))
	//        self.assertTrue(set([_attach_1.id, _attach_2.id]).issubset(msg.attachment_ids.ids),
	//                        'message_post: mail.message attachments duplicated')
	//        self.assertFalse(self.env['mail.mail'].search([('mail_message_id', '=', msg.message_id)]),
	//                         'message_post: mail.mail notifications should have been auto-deleted')
	//        self.assertEqual(set(m['email_from'] for m in self._mails),
	//                         set(['%s <%s>' % (self.user_employee.name,
	//                                           self.user_employee.email)]),
	//                         'message_post: notification email wrong email_from: should use sender email')
	//        self.assertEqual(set(m['email_to'][0] for m in self._mails),
	//                         set(['%s <%s>' % (self.partner_1.name, self.partner_1.email),
	//                              '%s <%s>' % (self.env.user.name, self.env.user.email)]))
	//        self.assertFalse(any(len(m['email_to']) != 1 for m in self._mails),
	//                         'message_post: notification email should be sent to one partner at a time')
	//        self.assertEqual(set(m['reply_to'] for m in self._mails),
	//                         set(['%s %s <%s@%s>' % (self.env.user.company_id.name,
	//                                                 self.group_pigs.name, self.group_pigs.alias_name, _domain)]),
	//                         'message_post: notification email should use group aliases and data for reply to')
	//        self.assertTrue(all(_subject in m['subject'] for m in self._mails))
	//        self.assertTrue(all(_body in m['body'] for m in self._mails))
	//        self.assertTrue(all(_body_alt in m['body'] for m in self._mails))
	//        self.assertFalse(any(m['references'] for m in self._mails))
}
func TestPostAnswer(self interface{}) {
	//        _body = '<p>Test Body</p>'
	//        _subject = 'Test Subject'
	//        _domain = 'schlouby.fr'
	//        _catchall = 'test_catchall'
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.domain', _domain)
	//        self.env['ir.config_parameter'].set_param(
	//            'mail.catchall.alias', _catchall)
	//        parent_msg = self.group_pigs.sudo(self.user_employee).message_post(
	//            body=_body, subject=_subject,
	//            message_type='comment', subtype='mt_comment')
	//        self.assertEqual(parent_msg.partner_ids, self.env['res.partner'])
	//        msg = self.group_pigs.sudo(self.user_employee).message_post(
	//            body=_body, subject=_subject, partner_ids=[self.partner_1.id],
	//            message_type='comment', subtype='mt_comment', parent_id=parent_msg.id)
	//        self.assertEqual(msg.parent_id.id, parent_msg.id)
	//        self.assertEqual(msg.partner_ids, self.partner_1)
	//        self.assertTrue(all('openerp-%d-mail.channel' %
	//                            self.group_pigs.id in m['references'] for m in self._mails))
	//        new_msg = self.group_pigs.sudo(self.user_employee).message_post(
	//            body=_body, subject=_subject,
	//            message_type='comment', subtype='mt_comment', parent_id=msg.id)
	//        self.assertEqual(new_msg.parent_id.id, parent_msg.id,
	//                         'message_post: flatten error')
	//        self.assertFalse(new_msg.partner_ids)
}
func TestMessageCompose(self interface{}) {
	//        composer = self.env['mail.compose.message'].with_context({
	//            'default_composition_mode': 'comment',
	//            'default_model': 'mail.channel',
	//            'default_res_id': self.group_pigs.id,
	//        }).sudo(self.user_employee).create({
	//            'body': '<p>Test Body</p>',
	//            'partner_ids': [(4, self.partner_1.id), (4, self.partner_2.id)]
	//        })
	//        self.assertEqual(composer.composition_mode,  'comment')
	//        self.assertEqual(composer.model, 'mail.channel')
	//        self.assertEqual(composer.subject, 'Re: %s' % self.group_pigs.name)
	//        self.assertEqual(composer.record_name, self.group_pigs.name)
	//        composer.send_mail()
	//        message = self.group_pigs.message_ids[0]
	//        composer = self.env['mail.compose.message'].with_context({
	//            'default_composition_mode': 'comment',
	//            'default_res_id': self.group_pigs.id,
	//            'default_parent_id': message.id
	//        }).sudo(self.user_employee).create({})
	//        self.assertEqual(composer.model, 'mail.channel')
	//        self.assertEqual(composer.res_id, self.group_pigs.id)
	//        self.assertEqual(composer.parent_id, message)
	//        self.assertEqual(composer.subject, 'Re: %s' % self.group_pigs.name)
}
func TestMessageComposeMassMail(self interface{}) {
	//        composer = self.env['mail.compose.message'].with_context({
	//            'default_composition_mode': 'mass_mail',
	//            'default_model': 'mail.channel',
	//            'default_res_id': False,
	//            'active_ids': [self.group_pigs.id, self.group_public.id]
	//        }).sudo(self.user_employee).create({
	//            'subject': 'Testing ${object.name}',
	//            'body': '<p>${object.description}</p>',
	//            'partner_ids': [(4, self.partner_1.id), (4, self.partner_2.id)]
	//        })
	//        composer.with_context({
	//            'default_res_id': -1,
	//            'active_ids': [self.group_pigs.id, self.group_public.id]
	//        }).send_mail()
	//        mails = self.env['mail.mail'].search([('subject', 'ilike', 'Testing')])
	//        for mail in mails:
	//            self.assertEqual(mail.recipient_ids, self.partner_1 | self.partner_2,
	//                             'compose wizard: mail_mail mass mailing: mail.mail in mass mail incorrect recipients')
	//        message1 = self.group_pigs.message_ids[0]
	//        self.assertEqual(message1.subject, 'Testing %s' % self.group_pigs.name)
	//        self.assertEqual(message1.body, '<p>%s</p>' %
	//                         self.group_pigs.description)
	//        message1 = self.group_public.message_ids[0]
	//        self.assertEqual(message1.subject, 'Testing %s' %
	//                         self.group_public.name)
	//        self.assertEqual(message1.body, '<p>%s</p>' %
	//                         self.group_public.description)
}
func TestMessageComposeMassMailActiveDomain(self interface{}) {
	//        self.env['mail.compose.message'].with_context({
	//            'default_composition_mode': 'mass_mail',
	//            'default_model': 'mail.channel',
	//            'active_ids': [self.group_pigs.id],
	//            'active_domain': [('name', 'in', ['%s' % self.group_pigs.name, '%s' % self.group_public.name])],
	//        }).sudo(self.user_employee).create({
	//            'subject': 'From Composer Test',
	//            'body': '${object.description}',
	//        }).send_mail()
	//        self.assertEqual(
	//            self.group_pigs.message_ids[0].subject, 'From Composer Test')
	//        self.assertEqual(
	//            self.group_public.message_ids[0].subject, 'From Composer Test')
}
