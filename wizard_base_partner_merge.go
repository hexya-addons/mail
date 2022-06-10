package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

var fields_BasePartnerMergeAutomaticWizard = map[string]models.FieldDefinition{}

// LogMergeOperation
func basePartnerMergeAutomaticWizard_LogMergeOperation(rs m.BasePartnerMergeAutomaticWizardSet, src_partners interface{}, dst_partner interface{}) {
	//        super(MergePartnerAutomatic, self)._log_merge_operation(
	//            src_partners, dst_partner)
	//        dst_partner.message_post(body='%s %s' % (_("Merged with the following partners:"), ", ".join(
	//            '%s <%s> (ID %s)' % (p.name, p.email or 'n/a', p.id) for p in src_partners)))
}
func init() {
	h.BasePartnerMergeAutomaticWizard().AddFields(fields_BasePartnerMergeAutomaticWizard)
	h.BasePartnerMergeAutomaticWizard().NewMethod("LogMergeOperation", basePartnerMergeAutomaticWizard_LogMergeOperation)

}
