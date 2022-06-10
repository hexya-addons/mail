package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

var fields_MailShortcode = map[string]models.FieldDefinition{
	"Source": fields.Char{
		String:   "Shortcut",
		Required: true,
		Index:    true,
		Help:     "The shortcut which must be replaced in the Chat Messages"},

	"Substitution": fields.Text{
		String:   "Substitution",
		Required: true,
		Index:    true,
		Help:     "The escaped html code replacing the shortcut"},

	"Description": fields.Char{
		String: "Description"},
}

func init() {
	models.NewModel("MailShortcode")
	h.MailShortcode().AddFields(fields_MailShortcode)

}
