package mail

import (
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

func autoVacuum_PowerOn(rs m.AutoVacuumSet) {
	h.MailThread().NewSet(rs.Env()).GarbageCollectAttachments()
	rs.Super().PowerOn()
}

func init() {
	h.AutoVacuum().Methods().PowerOn().Extend(autoVacuum_PowerOn)
}
