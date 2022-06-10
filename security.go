package mail

import (
	"github.com/hexya-addons/base"
	"github.com/hexya-erp/hexya/src/models/security"
	"github.com/hexya-erp/pool/h"
)

//vars

var ()

func init() {
	//group_init

	//rights
	h.MailMessage().Methods().Load().AllowGroup(security.GroupEveryone)
	h.MailMessage().Methods().AllowAllToGroup(base.GroupPortal)
	h.MailMessage().Methods().AllowAllToGroup(base.GroupUser)
	h.MailMail().Methods().Load().AllowGroup(security.GroupEveryone)
	h.MailMail().Methods().Load().AllowGroup(base.GroupPortal)
	h.MailMail().Methods().Write().AllowGroup(base.GroupPortal)
	h.MailMail().Methods().Create().AllowGroup(base.GroupPortal)
	h.MailMail().Methods().Load().AllowGroup(base.GroupUser)
	h.MailMail().Methods().Write().AllowGroup(base.GroupUser)
	h.MailMail().Methods().Create().AllowGroup(base.GroupUser)
	h.MailMail().Methods().AllowAllToGroup(base.GroupSystem)
	h.MailFollowers().Methods().Load().AllowGroup(security.GroupEveryone)
	h.MailFollowers().Methods().Load().AllowGroup(base.GroupPortal)
	h.MailFollowers().Methods().Write().AllowGroup(base.GroupPortal)
	h.MailFollowers().Methods().Create().AllowGroup(base.GroupPortal)
	h.MailFollowers().Methods().Load().AllowGroup(base.GroupUser)
	h.MailFollowers().Methods().Write().AllowGroup(base.GroupUser)
	h.MailFollowers().Methods().Create().AllowGroup(base.GroupUser)
	h.MailFollowers().Methods().AllowAllToGroup(base.GroupSystem)
	h.MailNotification().Methods().Load().AllowGroup(base.GroupPortal)
	h.MailNotification().Methods().Load().AllowGroup(base.GroupUser)
	h.MailNotification().Methods().Write().AllowGroup(base.GroupUser)
	h.MailNotification().Methods().Create().AllowGroup(base.GroupUser)
	h.MailNotification().Methods().AllowAllToGroup(base.GroupSystem)
	h.MailChannel().Methods().Load().AllowGroup(security.GroupEveryone)
	h.MailChannel().Methods().AllowAllToGroup(base.GroupUser)
	h.MailChannelPartner().Methods().Load().AllowGroup(base.GroupPublic)
	h.MailChannelPartner().Methods().AllowAllToGroup(base.GroupPortal)
	h.MailChannelPartner().Methods().AllowAllToGroup(base.GroupUser)
	h.MailModeration().Methods().AllowAllToGroup(base.GroupUser)
	h.MailAlias().Methods().Load().AllowGroup(security.GroupEveryone)
	h.MailAlias().Methods().AllowAllToGroup(base.GroupUser)
	h.MailAlias().Methods().AllowAllToGroup(base.GroupSystem)
	h.MailMessageSubtype().Methods().Load().AllowGroup(security.GroupEveryone)
	h.MailMessageSubtype().Methods().AllowAllToGroup(base.GroupUser)
	h.MailTrackingValue().Methods().AllowAllToGroup(base.GroupSystem)
	h.MailThread().Methods().AllowAllToGroup(security.GroupEveryone)
	h.PublisherWarrantyContract().Methods().AllowAllToGroup(security.GroupEveryone)
	h.MailTemplate().Methods().Load().AllowGroup(base.GroupUser)
	h.MailTemplate().Methods().Write().AllowGroup(base.GroupUser)
	h.MailTemplate().Methods().Create().AllowGroup(base.GroupUser)
	h.MailTemplate().Methods().AllowAllToGroup(base.GroupSystem)
	h.MailShortcode().Methods().AllowAllToGroup(base.GroupUser)
	h.MailShortcode().Methods().Load().AllowGroup(base.GroupPortal)
	h.MailActivity().Methods().AllowAllToGroup(base.GroupUser)
	h.MailActivityType().Methods().Load().AllowGroup(base.GroupUser)
	h.MailActivityType().Methods().AllowAllToGroup(base.GroupSystem)
	h.MailBlacklist().Methods().AllowAllToGroup(base.GroupSystem)
}
