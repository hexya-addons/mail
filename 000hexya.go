package mail

import (
	_ "github.com/hexya-addons/base"
	_ "github.com/hexya-addons/baseSetup"
	_ "github.com/hexya-addons/bus"
	_ "github.com/hexya-addons/web"
	"github.com/hexya-addons/web/controllers"
	"github.com/hexya-erp/hexya/src/server"
	"github.com/hexya-erp/hexya/src/tools/logging"
)

const MODULE_NAME string = "mail"

var log logging.Logger

func init() {
	log = logging.GetLogger(MODULE_NAME)
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})
	controllers.BackendScss = append(controllers.BackendScss,
		"/static/mail/src/scss/announcement.scss",
		"/static/mail/src/scss/discuss.scss",
		"/static/mail/src/scss/abstract_thread_window.scss",
		"/static/mail/src/scss/thread_window.scss",
		"/static/mail/src/scss/composer.scss",
		"/static/mail/src/scss/chatter.scss",
		"/static/mail/src/scss/followers.scss",
		"/static/mail/src/scss/thread.scss",
		"/static/mail/src/scss/systray.scss",
		"/static/mail/src/scss/mail_activity.scss",
		"/static/mail/src/scss/activity_view.scss",
		"/static/mail/src/scss/kanban_view.scss",
		"/static/mail/src/scss/attachment_box.scss",
	)
}
