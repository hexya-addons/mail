package mail

import (
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.IrAutovacuum().DeclareModel()

	h.IrAutovacuum().Methods().PowerOn().DeclareMethod(
		`PowerOn`,
		func(rs m.IrAutovacuumSet) {
			//        self.env['mail.thread']._garbage_collect_attachments()
			//        return super(AutoVacuum, self).power_on(*args, **kwargs)
		})
}
