package mail

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

func attachment_PostAddCreate(rs m.AttachmentSet) {
	// Overrides behaviour when the attachment is created through the controller
	rs.Super().PostAddCreate()
	for _, rec := range rs.Records() {
		rec.RegisterAsMainAttachment(false)
	}
}

// RegisterAsMainAttachment registers this attachment as the
// main one of the model it is attached to.
func attachment_RegisterAsMainAttachment(rs m.AttachmentSet, force bool) {
	rs.EnsureOne()
	if rs.ResModel() == "" {
		return
	}
	relatedRecord := models.Registry.MustGet(rs.ResModel()).BrowseOne(rs.Env(), rs.ResID())
	messageMainAttachment, ok := relatedRecord.Model().Fields().Get("MessageMainAttachment")
	if relatedRecord.IsNotEmpty() && ok {
		if force || relatedRecord.Get(messageMainAttachment).(models.RecordSet).IsEmpty() {
			relatedRecord.Set(messageMainAttachment, rs)
		}
	}
}

func init() {
	h.Attachment().Methods().PostAddCreate().Extend(attachment_PostAddCreate)
	h.Attachment().NewMethod("RegisterAsMainAttachment", attachment_RegisterAsMainAttachment)
}
