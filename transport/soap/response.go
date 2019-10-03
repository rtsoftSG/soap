package soap

import (
	"encoding/xml"
	"time"
)

type Kso_Server_Output_Response struct {
	XMLName xml.Name `xml:"http://heartbeat/assets/soap/kso/kso.proto/KsoSrv_types/ Kso_Server_Output"`

	ClientID     string                     `xml:"-"`
	LastWindowTs time.Time                  `xml:"LastWindowTs,omitempty"`
	Actions      []Kso_Server_Action_Record `xml:"Actions,omitempty"`
}

type ActionType int

const (
	_ ActionType = iota
	ActionTypeUpdateConfig
	ActionTypeUploadSo
)

type Kso_Server_Action_Record struct {
	ActionType ActionType `xml:"ActionType"`
	ActionData []byte     `xml:"ActionData,omitempty"`
}
