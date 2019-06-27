package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.MailShortcode().DeclareModel()

	h.MailShortcode().AddFields(map[string]models.FieldDefinition{
		"Source": models.CharField{
			String:   "Shortcut",
			Required: true,
			Index:    true,
			Help:     "The shortcut which must be replaced in the Chat Messages",
		},
		"Substitution": models.TextField{
			String:   "Substitution",
			Required: true,
			Index:    true,
			Help:     "The escaped html code replacing the shortcut",
		},
		"Description": models.CharField{
			String: "Description",
		},
		"ShortcodeType": models.SelectionField{
			Selection: types.Selection{
				"image": "Smiley",
				"text":  "Canned Response",
			},
			Required: true,
			Default:  models.DefaultValue("text"),
			Help: "* Smiley are only used for HTML code to display an image" +
				"* Text (default value) is used to substitute text with another text",
		},
	})
}
