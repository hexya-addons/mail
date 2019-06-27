package mail

import (
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.IrModel().DeclareModel()

	h.IrModel().Methods().Unlink().Extend(
		`Unlink`,
		func(rs m.IrModelSet) {
			//        query = "DELETE FROM mail_followers WHERE res_model IN %s"
			//        self.env.cr.execute(query, [tuple(self.mapped('model'))])
			//        return super(IrModel, self).unlink()
		})
}
