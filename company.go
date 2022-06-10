package mail

import (
	"fmt"

	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

var fields_Company = map[string]models.FieldDefinition{
	"Catchall": fields.Char{
		String:  "Catchall Email",
		Compute: h.Company().Methods().ComputeCatchall()},
}

// ComputeCatchall computes the catchall email of this company
func company_ComputeCatchall(rs m.CompanySet) m.CompanyData {
	res := h.Company().NewData().SetCatchall("")
	alias := h.ConfigParameter().NewSet(rs.Env()).Sudo().GetParam("mail.catchall.alias", "")
	domain := h.ConfigParameter().NewSet(rs.Env()).Sudo().GetParam("mail.catchall.domain", "")
	if alias != "" && domain != "" {
		res.SetCatchall(fmt.Sprintf("%s@%s", alias, domain))
	}
	return res
}

func init() {
	h.Company().AddFields(fields_Company)
	h.Company().NewMethod("ComputeCatchall", company_ComputeCatchall)
}
