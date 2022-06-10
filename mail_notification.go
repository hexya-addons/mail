package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

var fields_MailNotification = map[string]models.FieldDefinition{
	"MailMessage": fields.Many2One{
		RelationModel: h.MailMessage(),
		String:        "Message",
		Index:         true,
		OnDelete:      models.Cascade,
		Required:      true,
	},
	"Partner": fields.Many2One{
		RelationModel: h.Partner(),
		String:        "Needaction Recipient",
		Index:         true,
		OnDelete:      models.Cascade,
		Required:      true,
	},
	"IsRead": fields.Boolean{
		String: "Is Read",
		Index:  true,
	},
	"IsEmail": fields.Boolean{
		String: "Sent by Email",
		Index:  true,
	},
	"EmailStatus": fields.Selection{
		Selection: types.Selection{
			"ready":     "Ready to Send",
			"sent":      "Sent",
			"bounce":    "Bounced",
			"exception": "Exception",
			"canceled":  "Canceled",
		},
		String:  "Email Status",
		Default: models.DefaultValue("ready"),
		Index:   true,
	},
	"Mail": fields.Many2One{
		RelationModel: h.MailMail(),
		String:        "Mail",
		Index:         true,
	},
	"FailureType": fields.Selection{
		Selection: types.Selection{
			"SMTP":      "Connection failed (outgoing mail server problem)",
			"RECIPIENT": "Invalid email address",
			"BOUNCE":    "Email address rejected by destination",
			"UNKNOWN":   "Unknown error",
		},
		String: "Failure type",
	},
	"FailureReason": fields.Text{
		String: "Failure reason",
		NoCopy: true,
	},
}

// Init
func mailNotification_Init(rs m.MailNotificationSet) {
	//        self._cr.execute('SELECT indexname FROM pg_indexes WHERE indexname = %s',
	//                         ('mail_notification_res_partner_id_is_read_email_status_mail_message_id'))
	//        if not self._cr.fetchone():
	//            self._cr.execute(
	//                'CREATE INDEX mail_notification_res_partner_id_is_read_email_status_mail_message_id ON mail_message_res_partner_needaction_rel (res_partner_id, is_read, email_status, mail_message_id)')
}

// FormatFailureReason
func mailNotification_FormatFailureReason(rs m.MailNotificationSet) {
	//        self.ensure_one()
	//        if self.failure_type != 'UNKNOWN':
	//            return dict(type(self).failure_type.selection).get(self.failure_type, _('No Error'))
	//        else:
	//            return _("Unknown error") + ": %s" % (self.failure_reason or '')
}
func init() {
	models.NewModel("MailNotification")
	h.MailNotification().AddFields(fields_MailNotification)
	h.MailNotification().NewMethod("Init", mailNotification_Init)
	h.MailNotification().NewMethod("FormatFailureReason", mailNotification_FormatFailureReason)

}
