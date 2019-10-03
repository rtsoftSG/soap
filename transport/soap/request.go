package soap

import (
	"encoding/xml"
	"time"
)

type Kso_Window_Input_Request struct {
	XMLName xml.Name `xml:"http://heartbeat/assets/soap/kso/kso.proto/KsoSrv_types/ Kso_Window_Input"`

	ClientID string `xml:"ClientId"`
}

type Kso_Save_Event_Logs_Input_Request struct {
	XMLName xml.Name `xml:"http://heartbeat/assets/soap/kso/kso.proto/KsoSrv_types/ Kso_Save_Event_Logs_Input"`

	ClientID string                  `xml:"ClientId"`
	Events   []Kso_Event_Logs_Record `xml:"Events,omitempty"`
}

type Kso_Event_Logs_Record struct {
	Ts        time.Time `xml:"Ts,omitempty"`
	EventName string    `xml:"EventName"`
	EventData []byte    `xml:"EventData,omitempty"`
}
