// Copyright 2020 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package mail

import (
	"github.com/hexya-addons/mail/mailtypes"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

var fields_Partner = map[string]models.FieldDefinition{
	"MessageBounce": fields.Integer{
		String: "Bounce",
		Help:   "Counter of the number of bounced emails for this contact",
	},
	"Channels": fields.Many2Many{
		RelationModel:    h.MailChannel(),
		M2MLinkModelName: "MailChannelPartner",
		M2MOurField:      "Partner",
		M2MTheirField:    "Channel",
		String:           "Channels",
		NoCopy:           true,
	},
}

// Notify
func partner_Notify(rs m.PartnerSet, message m.MailMessageSet, rData []mailtypes.PartnerData, record m.MailThreadSet,
	forceSend bool, sendAfterCommit bool, modelDescription string, mailAutoDelete bool) bool {

	return true
}

// NotifyByChat
func partner_NotifyByChat(rs m.PartnerSet, message m.MailMessageSet) {

}

func init() {
	h.Partner().InheritModel(h.MailThread())
	h.Partner().InheritModel(h.MailActivityMixin())
	h.Partner().InheritModel(h.MailBlacklistMixin())
	h.Partner().AddFields(fields_Partner)

	h.Partner().NewMethod("Notify", partner_Notify)
	h.Partner().NewMethod("NotifyByChat", partner_NotifyByChat)
	// h.Partner().Methods().MessageGetSuggestedRecipients().Extend()
	// h.Partner().Methods().MessageGetDefaultRecipients().Extend()
}
