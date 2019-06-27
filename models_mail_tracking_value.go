package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.MailTrackingValue().DeclareModel()

	h.MailTrackingValue().AddFields(map[string]models.FieldDefinition{
		"Field": models.CharField{
			String:   "Changed Field",
			Required: true,
			ReadOnly: true,
		},
		"FieldDesc": models.CharField{
			String:   "Field Description",
			Required: true,
			ReadOnly: true,
		},
		"FieldType": models.CharField{
			String: "Field Type",
		},
		"OldValueInteger": models.IntegerField{
			String:   "Old Value Integer",
			ReadOnly: true,
		},
		"OldValueFloat": models.FloatField{
			String:   "Old Value Float",
			ReadOnly: true,
		},
		"OldValueMonetary": models.FloatField{
			String:   "Old Value Monetary",
			ReadOnly: true,
		},
		"OldValueChar": models.CharField{
			String:   "Old Value Char",
			ReadOnly: true,
		},
		"OldValueText": models.TextField{
			String:   "Old Value Text",
			ReadOnly: true,
		},
		"OldValueDatetime": models.DateTimeField{
			String:   "Old Value DateTime",
			ReadOnly: true,
		},
		"NewValueInteger": models.IntegerField{
			String:   "New Value Integer",
			ReadOnly: true,
		},
		"NewValueFloat": models.FloatField{
			String:   "New Value Float",
			ReadOnly: true,
		},
		"NewValueMonetary": models.FloatField{
			String:   "New Value Monetary",
			ReadOnly: true,
		},
		"NewValueChar": models.CharField{
			String:   "New Value Char",
			ReadOnly: true,
		},
		"NewValueText": models.TextField{
			String:   "New Value Text",
			ReadOnly: true,
		},
		"NewValueDatetime": models.DateTimeField{
			String:   "New Value Datetime",
			ReadOnly: true,
		},
		"MailMessageId": models.Many2OneField{
			RelationModel: h.MailMessage(),
			String:        "Message ID",
			Required:      true,
			Index:         true,
			OnDelete:      `cascade`,
		},
	})
	h.MailTrackingValue().Methods().CreateTrackingValues().DeclareMethod(
		`CreateTrackingValues`,
		func(rs m.MailTrackingValueSet, initial_value interface{}, new_value interface{}, col_name interface{}, col_info interface{}) {
			//        tracked = True
			//        values = {'field': col_name,
			//                  'field_desc': col_info['string'], 'field_type': col_info['type']}
			//        if col_info['type'] in ['integer', 'float', 'char', 'text', 'datetime', 'monetary']:
			//            values.update({
			//                'old_value_%s' % col_info['type']: initial_value,
			//                'new_value_%s' % col_info['type']: new_value
			//            })
			//        elif col_info['type'] == 'date':
			//            values.update({
			//                'old_value_datetime': initial_value and datetime.strftime(datetime.combine(datetime.strptime(initial_value, tools.DEFAULT_SERVER_DATE_FORMAT), datetime.min.time()), tools.DEFAULT_SERVER_DATETIME_FORMAT) or False,
			//                'new_value_datetime': new_value and datetime.strftime(datetime.combine(datetime.strptime(new_value, tools.DEFAULT_SERVER_DATE_FORMAT), datetime.min.time()), tools.DEFAULT_SERVER_DATETIME_FORMAT) or False,
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
		})
	h.MailTrackingValue().Methods().GetDisplayValue().DeclareMethod(
		`GetDisplayValue`,
		func(rs m.MailTrackingValueSet, typeName interface{}) {
			//        assert type in ('new', 'old')
			//        result = []
			//        for record in self:
			//            if record.field_type in ['integer', 'float', 'char', 'text', 'datetime', 'monetary']:
			//                result.append(getattr(record, '%s_value_%s' %
			//                                      (type, record.field_type)))
			//            elif record.field_type == 'date':
			//                if record['%s_value_datetime' % type]:
			//                    new_date = datetime.strptime(
			//                        record['%s_value_datetime' % type], tools.DEFAULT_SERVER_DATETIME_FORMAT).date()
			//                    result.append(new_date.strftime(
			//                        tools.DEFAULT_SERVER_DATE_FORMAT))
			//                else:
			//                    result.append(record['%s_value_datetime' % type])
			//            elif record.field_type == 'boolean':
			//                result.append(bool(record['%s_value_integer' % type]))
			//            else:
			//                result.append(record['%s_value_char' % type])
			//        return result
		})
	h.MailTrackingValue().Methods().GetOldDisplayValue().DeclareMethod(
		`GetOldDisplayValue`,
		func(rs m.MailTrackingValueSet) {
			//        return self.get_display_value('old')
		})
	h.MailTrackingValue().Methods().GetNewDisplayValue().DeclareMethod(
		`GetNewDisplayValue`,
		func(rs m.MailTrackingValueSet) {
			//        return self.get_display_value('new')
		})
}
