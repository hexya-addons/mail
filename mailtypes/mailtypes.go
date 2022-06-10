// Copyright 2020 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package mailtypes

import (
	"encoding/json"

	"github.com/hexya-addons/web/webtypes"
	"github.com/hexya-erp/hexya/src/models/types/dates"
)

// A CheckContactResult holds values of the result of a contact check
type CheckContactResult struct {
	OK            bool
	ErrorMessage  string
	ErrorTemplate string
}

// RecipientData holds all data necessary to notify recipients
type RecipientData struct {
	PartnerID        int64    `db:"pid"`
	ChannelID        int64    `db:"cid"`
	Active           bool     `db:"active"`
	Share            bool     `db:"share"`
	NotificationType string   `db:"notif"`
	Groups           []string `db:"groups"`
}

// SubscriptionData holds all follower data for a document
type SubscriptionData struct {
	FollowerID   int64   `db:"id"`
	ResID        int64   `db:"res_id"`
	PartnerID    int64   `db:"partner_id"`
	ChannelID    int64   `db:"channel_id"`
	SubtypeIds   []int64 `db:"subtype_ids"`
	PartnerShare bool    `db:"partner_share"`
}

// TrackingData is a lightweight container for holding a MailTrackingValue
// when interacting with the client.
type TrackingData struct {
	ID           int64  `json:"id"`
	OldValue     string `json:"old_value"`
	NewValue     string `json:"new_value"`
	FieldType    string `json:"field_type"`
	ChangedField string `json:"changed_field"`
}

// A PartnerEmailStatus links a partner name with an EmailStatus
type PartnerEmailStatus struct {
	Name        string
	EmailStatus string
}

// MarshalJSON a partner email status as a list
func (p PartnerEmailStatus) MarshalJSON() ([]byte, error) {
	aux := [2]interface{}{p.EmailStatus, p.Name}
	return json.Marshal(aux)
}

// A MailFailureInfo is a shorter message to notify a failure update
type MailFailureInfo struct {
	MessageID       int64
	RecordName      string
	ModelName       string
	UUID            int64
	ResID           int64
	Model           string
	LastMessageDate dates.DateTime
	ModuleIcon      string
	Notifications   map[int64]PartnerEmailStatus
}

// MailAttachmentInfo is the format of an attachment in a mail message for the client
type MailAttachmentInfo struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Filename     string `json:"filename"`
	FileTypeIcon string `json:"file_type_icon"`
	Mimetype     string `json:"mimetype"`
	IsMain       bool   `json:"is_main"`
}

// CustomerEmailInfo holds a customer ID, name and email status
type CustomerEmailInfo struct {
	ID          int64
	Name        string
	EmailStatus string
}

// MarshalJSON a customer info as a list
func (c CustomerEmailInfo) MarshalJSON() ([]byte, error) {
	aux := [3]interface{}{c.ID, c.Name, c.EmailStatus}
	return json.Marshal(aux)
}

// MailMessageInfo is the format of a mail message for the client
type MailMessageInfo struct {
	ID                   int64                       `json:"id"`
	Body                 string                      `json:"body"`
	Model                string                      `json:"model"`
	ResID                int64                       `json:"res_id"`
	RecordName           string                      `json:"record_name"`
	Attachments          []*MailAttachmentInfo       `json:"attachment_ids"`
	NeedactionPartnerIds []int64                     `json:"needaction_partner_ids"`
	TrackingValues       []*TrackingData             `json:"tracking_value_ids"`
	Author               webtypes.RecordIDWithName   `json:"author_id"`
	EmailFrom            string                      `json:"email_from"`
	Subtype              webtypes.RecordIDWithName   `json:"subtype_id"`
	SubtypeDescription   string                      `json:"subtype_description"`
	ChannelIds           []int64                     `json:"channel_ids"`
	Date                 dates.DateTime              `json:"date"`
	Partners             []webtypes.RecordIDWithName `json:"partner_ids"`
	MessageType          string                      `json:"message_type"`
	Subject              string                      `json:"subject"`
	IsNote               bool                        `json:"is_note"`
	IsDiscussion         bool                        `json:"is_discussion"`
	IsNotification       bool                        `json:"is_notification"`
	ModerationStatue     string                      `json:"moderation_statue"`
	CustomerEmailData    []CustomerEmailInfo         `json:"customer_email_data"`
	CustomerEmailStatus  string                      `json:"customer_email_status"`
	StarredPartners      []int64                     `json:"starred_partner_ids"`
}

// PartnerData as used in RecipientsData
type PartnerData struct {
	ID               int64    `json:"id"`
	Share            bool     `json:"share"`
	Active           bool     `json:"active"`
	NotificationType string   `json:"notif"`
	Type             string   `json:"type"`
	Groups           []string `json:"groups"`
}

// ChannelData as used in RecipientsData
type ChannelData struct {
	ID               int64  `json:"id"`
	NotificationType string `json:"notif"`
	Type             string `json:"type"`
}

// A RecipientsData is used to notify recipients
type RecipientsData struct {
	Partners []PartnerData
	Channels []ChannelData
}
