package soap

import (
	"encoding/xml"
)

type Kso_Window_Input_Request struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Window_Input"`

	AppID string `xml:"AppId"`
}

type Kso_Save_Event_Logs_Input_Request struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Save_Event_Logs_Input"`

	AppID  string                  `xml:"AppId"`
	Events []Kso_Event_Logs_Record `xml:"Events,omitempty"`
}

type Kso_Event_Logs_Record struct {
	Ts        int64  `xml:"Ts"` // unix
	EventName string `xml:"EventName"`
	EventData string `xml:"EventData,omitempty"` // base64
}
