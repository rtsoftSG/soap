package soap

import (
	"encoding/xml"
)

type KsoServerResponse struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Server_Output"`

	EndTs   int64                   `xml:"EndTs,omitempty"` // unix
	Actions []KsoServerActionRecord `xml:"Actions,omitempty"`
}

type ActionType int

const (
	_ ActionType = iota
	ActionTypeUpdateConfig
	ActionTypeUploadSo
)

type KsoServerActionRecord struct {
	ActionType ActionType `xml:"ActionType"`
	ActionData string     `xml:"ActionData,omitempty"` // base64
}
