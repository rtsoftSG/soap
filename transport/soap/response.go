package soap

import (
	"encoding/xml"
)

type Kso_Server_Output_Response struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Server_Output"`

	AppID   string                     `xml:"-"`
	EndTs   int64                      `xml:"EndTs,omitempty"` // unix
	Actions []Kso_Server_Action_Record `xml:"Actions,omitempty"`
}

type ActionType int

const (
	_ ActionType = iota
	ActionTypeUpdateConfig
	ActionTypeUploadSo
)

type Kso_Server_Action_Record struct {
	ActionType ActionType `xml:"ActionType"`
	ActionData string     `xml:"ActionData,omitempty"` // base64
}
