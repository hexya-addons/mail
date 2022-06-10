package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

var fields_MailTrackingValue = map[string]models.FieldDefinition{
	"Field": fields.Char{
		String:   "Changed Field",
		Required: true,
		ReadOnly: true,
	},
	"FieldDesc": fields.Char{
		String:   "Field Description",
		Required: true,
		ReadOnly: true,
	},
	"FieldType": fields.Char{
		String: "Field Type",
	},
	"OldValueInteger": fields.Integer{
		String:   "Old Value Integer",
		ReadOnly: true,
	},
	"OldValueFloat": fields.Float{
		String:   "Old Value Float",
		ReadOnly: true,
	},
	"OldValueMonetary": fields.Float{
		String:   "Old Value Monetary",
		ReadOnly: true,
	},
	"OldValueChar": fields.Char{
		String:   "Old Value Char",
		ReadOnly: true,
	},
	"OldValueText": fields.Text{
		String:   "Old Value Text",
		ReadOnly: true,
	},
	"OldValueDatetime": fields.DateTime{
		String:   "Old Value DateTime",
		ReadOnly: true,
	},
	"NewValueInteger": fields.Integer{
		String:   "New Value Integer",
		ReadOnly: true,
	},
	"NewValueFloat": fields.Float{
		String:   "New Value Float",
		ReadOnly: true,
	},
	"NewValueMonetary": fields.Float{
		String:   "New Value Monetary",
		ReadOnly: true,
	},
	"NewValueChar": fields.Char{
		String:   "New Value Char",
		ReadOnly: true,
	},
	"NewValueText": fields.Text{
		String:   "New Value Text",
		ReadOnly: true,
	},
	"NewValueDatetime": fields.DateTime{
		String:   "New Value Datetime",
		ReadOnly: true,
	},
	"MailMessage": fields.Many2One{
		RelationModel: h.MailMessage(),
		String:        "Message ID",
		Required:      true,
		Index:         true,
		OnDelete:      models.Cascade,
	},
	"TrackSequence": fields.Integer{
		String:   "Tracking field sequence",
		ReadOnly: true,
		Default:  models.DefaultValue(100),
	},
}

// CreateTrackingValues
func mailTrackingValue_CreateTrackingValues(rs m.MailTrackingValueSet, initial_value interface{}, new_value interface{}, col_name interface{}, col_info interface{}, track_sequence interface{}) {
	//        tracked = True
	//        values = {'field': col_name, 'field_desc': col_info['string'],
	//                  'field_type': col_info['type'], 'track_sequence': track_sequence}
	//        if col_info['type'] in ['integer', 'float', 'char', 'text', 'datetime', 'monetary']:
	//            values.update({
	//                'old_value_%s' % col_info['type']: initial_value,
	//                'new_value_%s' % col_info['type']: new_value
	//            })
	//        elif col_info['type'] == 'date':
	//            values.update({
	//                'old_value_datetime': initial_value and fields.Datetime.to_string(datetime.combine(fields.Date.from_string(initial_value), datetime.min.time())) or False,
	//                'new_value_datetime': new_value and fields.Datetime.to_string(datetime.combine(fields.Date.from_string(new_value), datetime.min.time())) or False,
	//            })
	//        elif col_info['type'] == 'boolean':
	//            values.update({
	//                'old_value_integer': initial_value,
	//                'new_value_integer': new_value
	//            })
	//        elif col_info['type'] == 'selection':
	//            values.update({
	//                'old_value_char': initial_value and dict(col_info['selection'])[initial_value] or '',
	//                'new_value_char': new_value and dict(col_info['selection'])[new_value] or ''
	//            })
	//        elif col_info['type'] == 'many2one':
	//            values.update({
	//                'old_value_integer': initial_value and initial_value.id or 0,
	//                'new_value_integer': new_value and new_value.id or 0,
	//                'old_value_char': initial_value and initial_value.sudo().name_get()[0][1] or '',
	//                'new_value_char': new_value and new_value.sudo().name_get()[0][1] or ''
	//            })
	//        else:
	//            tracked = False
	//        if tracked:
	//            return values
	//        return {}
}

// GetDisplayValue
func mailTrackingValue_GetDisplayValue(rs m.MailTrackingValueSet, newValue bool) string {
	rs.EnsureOne()
	return ""
	//        assert type in ('new', 'old')
	//        result = []
	//        for record in self:
	//            if record.field_type in ['integer', 'float', 'char', 'text', 'monetary']:
	//                result.append(getattr(record, '%s_value_%s' %
	//                                      (type, record.field_type)))
	//            elif record.field_type == 'datetime':
	//                if record['%s_value_datetime' % type]:
	//                    new_datetime = getattr(record, '%s_value_datetime' % type)
	//                    result.append('%sZ' % new_datetime)
	//                else:
	//                    result.append(record['%s_value_datetime' % type])
	//            elif record.field_type == 'date':
	//                if record['%s_value_datetime' % type]:
	//                    new_date = record['%s_value_datetime' % type]
	//                    result.append(fields.Date.to_string(new_date))
	//                else:
	//                    result.append(record['%s_value_datetime' % type])
	//            elif record.field_type == 'boolean':
	//                result.append(bool(record['%s_value_integer' % type]))
	//            else:
	//                result.append(record['%s_value_char' % type])
	//        return result
}

// GetOldDisplayValue
func mailTrackingValue_GetOldDisplayValue(rs m.MailTrackingValueSet) string {
	return rs.GetDisplayValue(false)
}

// GetNewDisplayValue
func mailTrackingValue_GetNewDisplayValue(rs m.MailTrackingValueSet) string {
	return rs.GetDisplayValue(true)
}

func init() {
	models.NewModel("MailTrackingValue")
	h.MailTrackingValue().AddFields(fields_MailTrackingValue)
	h.MailTrackingValue().NewMethod("CreateTrackingValues", mailTrackingValue_CreateTrackingValues)
	h.MailTrackingValue().NewMethod("GetDisplayValue", mailTrackingValue_GetDisplayValue)
	h.MailTrackingValue().NewMethod("GetOldDisplayValue", mailTrackingValue_GetOldDisplayValue)
	h.MailTrackingValue().NewMethod("GetNewDisplayValue", mailTrackingValue_GetNewDisplayValue)

}
