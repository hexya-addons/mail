package mail

import (
	"net/http"

	"github.com/hexya-erp/hexya/src/controllers"
)

func DefaultRequestUid(self interface{}) {
	//        """ For Anonymous people, they receive the access right of SUPERUSER_ID since they have NO access (auth=none)
	//            !!! Each time a method from this controller is call, there is a check if the user (who can be anonymous and Sudo access)
	//            can access to the resource.
	//        """
	//        return request.session.uid and request.session.uid or SUPERUSER_ID
}
func Poll(self interface{}, dbname interface{}, channels interface{}, last interface{}, options interface{}) {
	//        if request.session.uid:
	//            partner_id = request.env.user.partner_id.id
	//
	//            if partner_id:
	//                channels = list(channels)       # do not alter original list
	//                for mail_channel in request.env['mail.channel'].search([('channel_partner_ids', 'in', [partner_id])]):
	//                    channels.append(
	//                        (request.db, 'mail.channel', mail_channel.id))
	//                # personal and needaction channel
	//                channels.append((request.db, 'res.partner', partner_id))
	//                channels.append((request.db, 'ir.needaction', partner_id))
	//        return super(MailChatController, self)._poll(dbname, channels, last, options)
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/chat_post") {
		mail.ExtendController(http.MethodPost, "/chat_post", MailChatPost)
	} else {
		mail.AddController(http.MethodPost, "/chat_post", MailChatPost)
	}
}
func MailChatPost(self interface{}, uuid interface{}, message_content interface{}) {
	//        author_id = False  # message_post accept 'False' author_id, but not 'None'
	//        if request.session.uid:
	//            author_id = request.env['res.users'].sudo().browse(
	//                request.session.uid).partner_id.id
	//        mail_channel = request.env["mail.channel"].sudo().search(
	//            [('uuid', '=', uuid)], limit=1)
	//        message = mail_channel.sudo().with_context(mail_create_nosubscribe=True).message_post(author_id=author_id, email_from=False,
	//                                                                                              body=message_content, message_type='comment', subtype='mail.mt_comment', content_subtype='plaintext')
	//        return message and message.id or False
}
func init() {
	root := controllers.Registry
	var ok bool
	var mail *controllers.Group
	mail, ok = root.GetGroup("/mail")
	if !ok {
		mail = root.AddGroup("/mail")
	}
	if mail.HasController(http.MethodGet, "/chat_history") {
		mail.ExtendController(http.MethodPost, "/chat_history", MailChatHistory)
	} else {
		mail.AddController(http.MethodPost, "/chat_history", MailChatHistory)
	}
}
func MailChatHistory(self interface{}, uuid interface{}, last_id interface{}, limit interface{}) {
	//        channel = request.env["mail.channel"].sudo().search(
	//            [('uuid', '=', uuid)], limit=1)
	//        return channel.sudo().channel_fetch_message(last_id, limit)
}
