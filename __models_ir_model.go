package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

var fields_IrModel = map[string]models.FieldDefinition{
	"IsMailThread": fields.Boolean{
		String: "Mail Thread",
		//oldname='mail_thread'
		Default: models.DefaultValue(false),
		Help:    "Whether this model supports messages and notifications."},
}

// Unlink
func irModel_Unlink(rs m.IrModelSet) {
	//        models = tuple(self.mapped('model'))
	//        query = "DELETE FROM mail_followers WHERE res_model IN %s"
	//        self.env.cr.execute(query, [models])
	//        query = "DELETE FROM mail_message WHERE model in %s"
	//        self.env.cr.execute(query, [models])
	//        query = """
	//            SELECT DISTINCT store_fname
	//            FROM ir_attachment
	//            WHERE res_model IN %s
	//            EXCEPT
	//            SELECT store_fname
	//            FROM ir_attachment
	//            WHERE res_model not IN %s;
	//        """
	//        self.env.cr.execute(query, [models, models])
	//        fnames = self.env.cr.fetchall()
	//        query = """DELETE FROM ir_attachment WHERE res_model in %s"""
	//        self.env.cr.execute(query, [models])
	//        for (fname) in fnames:
	//            self.env['ir.attachment']._file_delete(fname)
	//        return super(IrModel, self).unlink()
}

// Write
func irModel_Write(rs m.IrModelSet, vals models.RecordData) {
	//        if self and 'is_mail_thread' in vals:
	//            if not all(rec.state == 'manual' for rec in self):
	//                raise UserError(_('Only custom models can be modified.'))
	//            if not all(rec.is_mail_thread <= vals['is_mail_thread'] for rec in self):
	//                raise UserError(
	//                    _('Field "Mail Thread" cannot be changed to "False".'))
	//            res = super(IrModel, self).write(vals)
	//            # setup models; this reloads custom models in registry
	//            self.pool.setup_models(self._cr)
	//            # update database schema of models
	//            models = self.pool.descendants(self.mapped('model'), '_inherits')
	//            self.pool.init_models(self._cr, models, dict(
	//                self._context, update_custom_fields=True))
	//        else:
	//            res = super(IrModel, self).write(vals)
	//        return res
}

// ReflectModelParams
func irModel_ReflectModelParams(rs m.IrModelSet, model interface{}) {
	//        vals = super(IrModel, self)._reflect_model_params(model)
	//        vals['is_mail_thread'] = issubclass(
	//            type(model), self.pool['mail.thread'])
	//        return vals
}

// Instanciate
func irModel_Instanciate(rs m.IrModelSet, model_data interface{}) {
	//        model_class = super(IrModel, self)._instanciate(model_data)
	//        if model_data.get('is_mail_thread') and model_class._name != 'mail.thread':
	//            parents = model_class._inherit or []
	//            parents = [parents] if isinstance(
	//                parents, pycompat.string_types) else parents
	//            model_class._inherit = parents + ['mail.thread']
	//        return model_class
}

var fields_IrModelFields = map[string]models.FieldDefinition{
	"TrackVisibility": fields.Selection{
		Selection: types.Selection{
			"onchange": "On Change",
			"always":   "Always",
		},
		String: "Tracking",
		Help: "When set, every modification to this field will be tracked" +
			"in the chatter."},
}

// ReflectFieldParams
func irModelFields_ReflectFieldParams(rs m.IrModelFieldsSet, field interface{}) {
	//        vals = super(IrModelField, self)._reflect_field_params(field)
	//        vals['track_visibility'] = getattr(field, 'track_visibility', None)
	//        return vals
}

// InstanciateAttrs
func irModelFields_InstanciateAttrs(rs m.IrModelFieldsSet, field_data interface{}) {
	//        attrs = super(IrModelField, self)._instanciate_attrs(field_data)
	//        if attrs and field_data.get('track_visibility'):
	//            attrs['track_visibility'] = field_data['track_visibility']
	//        return attrs
}
func init() {
	models.NewModel("IrModel")
	h.IrModel().AddFields(fields_IrModel)
	h.IrModel().Methods().Unlink().Extend(irModel_Unlink)
	h.IrModel().Methods().Write().Extend(irModel_Write)
	h.IrModel().NewMethod("ReflectModelParams", irModel_ReflectModelParams)
	h.IrModel().NewMethod("Instanciate", irModel_Instanciate)

	models.NewModel("IrModelFields")
	h.IrModelFields().AddFields(fields_IrModelFields)
	h.IrModelFields().NewMethod("ReflectFieldParams", irModelFields_ReflectFieldParams)
	h.IrModelFields().NewMethod("InstanciateAttrs", irModelFields_InstanciateAttrs)

}
