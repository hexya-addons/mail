package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

var fields_Email_templatePreview = map[string]models.FieldDefinition{
	"ResId": fields.Selection{
		Selection: _get_records,
		String:    "Sample Document"},

	"PartnerIds": fields.Many2Many{
		RelationModel: h.Partner(),
		String:        "Recipients"},

	"AttachmentIds": fields.Many2Many{
		String: "Attachments",
		Stored: false},
}

//  Return Records of particular Email Template's Model
func email_templatePreview_GetRecords(rs m.Email_templatePreviewSet) {
	//        template_id = self._context.get('template_id')
	//        default_res_id = self._context.get('default_res_id')
	//        if not template_id:
	//            return []
	//        template = self.env['mail.template'].browse(int(template_id))
	//        records = self.env[template.model_id.model].search(
	//            [], order="id desc", limit=10)
	//        records |= records.browse(default_res_id)
	//        return records.name_get()
}

// DefaultGet
func email_templatePreview_DefaultGet(rs m.Email_templatePreviewSet, fields interface{}) {
	//        result = super(TemplatePreview, self).default_get(fields)
	//        if 'res_id' in fields and not result.get('res_id'):
	//            records = self._get_records()
	//            # select first record as a Default
	//            result['res_id'] = records and records[0][0] or False
	//        if self._context.get('template_id') and 'model_id' in fields and not result.get('model_id'):
	//            result['model_id'] = self.env['mail.template'].browse(
	//                self._context['template_id']).model_id.id
	//        return result
}

// OnChangeResId
func email_templatePreview_OnChangeResId(rs m.Email_templatePreviewSet) {
	//        if not self.res_id:
	//            return {}
	//        mail_values = {}
	//        if self._context.get('template_id'):
	//            template = self.env['mail.template'].browse(
	//                self._context['template_id'])
	//            self.name = template.name
	//            mail_values = template.generate_email(self.res_id)
	//        for field in ['email_from', 'email_to', 'email_cc', 'reply_to', 'subject', 'body_html', 'partner_to', 'partner_ids', 'attachment_ids']:
	//            setattr(self, field, mail_values.get(field, False))
}
func init() {
	models.NewModel("Email_templatePreview")
	h.Email_templatePreview().AddFields(fields_Email_templatePreview)
	h.Email_templatePreview().NewMethod("GetRecords", email_templatePreview_GetRecords)
	h.Email_templatePreview().Methods().DefaultGet().Extend(email_templatePreview_DefaultGet)
	h.Email_templatePreview().NewMethod("OnChangeResId", email_templatePreview_OnChangeResId)

}
