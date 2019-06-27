package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.MailNotification().DeclareModel()

	h.MailNotification().AddFields(map[string]models.FieldDefinition{
		"MailMessageId": models.Many2OneField{
			RelationModel: h.MailMessage(),
			String:        "Message",
			Index:         true,
			OnDelete:      `cascade`,
			Required:      true,
		},
		"ResPartnerId": models.Many2OneField{
			RelationModel: h.Partner(),
			String:        "Needaction Recipient",
			Index:         true,
			OnDelete:      `cascade`,
			Required:      true,
		},
		"IsRead": models.BooleanField{
			String: "Is Read",
			Index:  true,
		},
		"IsEmail": models.BooleanField{
			String: "Sent by Email",
			Index:  true,
		},
		"EmailStatus": models.SelectionField{
			Selection: types.Selection{
				"ready":     "Ready to Send",
				"sent":      "Sent",
				"bounce":    "Bounced",
				"exception": "Exception",
			},
			String:  "Email Status",
			Default: models.DefaultValue("ready"),
			Index:   true,
		},
	})
	h.MailNotification().Methods().Init().DeclareMethod(
		`Init`,
		func(rs m.MailNotificationSet) {
			//        self._cr.execute('SELECT indexname FROM pg_indexes WHERE indexname = %s',
			//                         ('mail_notification_res_partner_id_is_read_email_status_mail_message_id'))
			//        if not self._cr.fetchone():
			//            self._cr.execute(
			//                'CREATE INDEX mail_notification_res_partner_id_is_read_email_status_mail_message_id ON mail_message_res_partner_needaction_rel (res_partner_id, is_read, email_status, mail_message_id)')
		})
}
