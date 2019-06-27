package mail

import (
	"github.com/hexya-addons/web/controllers"
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "mail"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})
	controllers.BackendJS = append(controllers.BackendJS,
		"/static/mail/src/js/many2many_tags_email.js",
		"/static/mail/src/js/client_action.js",
		"/static/mail/src/js/chat_window.js",
		"/static/mail/src/js/extended_chat_window.js",
		"/static/mail/src/js/composer.js",
		"/static/mail/src/js/chat_manager.js",
		"/static/mail/src/js/chatter.js",
		"/static/mail/src/js/followers.js",
		"/static/mail/src/js/thread.js",
		"/static/mail/src/js/systray.js",
		"/static/mail/src/js/tour.js",
		"/static/mail/src/js/utils.js",
		"/static/mail/src/js/window_manager.js",
	)
	controllers.Backend = append(controllers.Backend,
		"/static/mail/src/less/announcement.less",
	)
	controllers.BackendLess = append(controllers.BackendLess,
		"/static/mail/src/less/client_action.less",
		"/static/mail/src/less/chat_window.less",
		"/static/mail/src/less/extended_chat_window.less",
		"/static/mail/src/less/composer.less",
		"/static/mail/src/less/chatter.less",
		"/static/mail/src/less/followers.less",
		"/static/mail/src/less/thread.less",
		"/static/mail/src/less/systray.less",
	)

}
