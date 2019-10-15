package soap

import (
	"encoding/xml"
)

type KsoOpenWindowRequest struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Open_Window_Request"`

	AppID string `xml:"AppId"`
}

type KsoSaveEventLogsRequest struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Save_Event_Logs_Request"`

	AppID  string               `xml:"AppId"`
	Events []KsoEventLogsRecord `xml:"Events,omitempty"`
}

type KsoEventLogsRecord struct {
	Ts        int64  `xml:"Ts"` // unix
	EventName string `xml:"EventName"`
	EventData string `xml:"EventData,omitempty"` // base64
}

type SavePackageCntRequest struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Save_Package_Cnt_Request"`

	AppID string             `xml:"AppId"`
	Pkg   []PackageCntRecord `xml:"Pkg,omitempty"`
}

type PackageCntRecord struct {
	Ts      int64  `xml:"Ts"` // unix
	ActorId string `xml:"ActorId"`
	Cnt     int    `xml:"Cnt,omitempty"`
}
