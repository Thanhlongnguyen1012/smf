package models

type SMContextCreatedData struct {
	PduSessionID int32  `json:"pduSessionId"`
	SNssai       string `json:"sNssai,omitempty"`
}
