package soap

import (
	"encoding/xml"
)

type OpenWindowRequest struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Open_Window_Request"`

	AppID string `xml:"AppId"`
}

type SaveEventLogsRequest struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Save_Event_Logs_Request"`

	AppID  string            `xml:"AppId"`
	Events []EventLogsRecord `xml:"Events"`
}

type EventLogsRecord struct {
	Ts        int64  `xml:"Ts"` // unix
	EventName string `xml:"EventName"`
	EventData string `xml:"EventData,omitempty"` // base64
}

type SavePackageCntRequest struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Save_Package_Cnt_Request"`

	AppID string             `xml:"AppId"`
	Pkg   []PackageCntRecord `xml:"Pkg"`
}

type PackageCntRecord struct {
	Ts      int64  `xml:"Ts"` // unix
	ActorId string `xml:"ActorId"`
	Cnt     int    `xml:"Cnt"`
}

type SaveMetaInfoRequest struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Save_Meta_Info_Request"`

	AppID string       `xml:"AppId"`
	Info  []DataRecord `xml:"Info"`
}

type SaveSysMonitorRequest struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Save_Sys_Monitor_Request"`

	AppID   string       `xml:"AppId"`
	Metrics []DataRecord `xml:"Metrics"`
}

type DataRecord struct {
	Ts   int64  `xml:"Ts"`   // unix
	Data string `xml:"Data"` // base64
}

type SaveCalcResultsRequest struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Save_Calc_Results_Request"`

	AppID    string              `xml:"AppId"`
	CalcName string              `xml:"CalcName"`
	CalcRes  []CalcResultsRecord `xml:"CalcRes"`
}

type CalcResultsRecord struct {
	Ts  int64   `xml:"Ts"` // unix
	Ua  float32 `xml:"Ua"`
	Ub  float32 `xml:"Ub"`
	Uc  float32 `xml:"Uc"`
	Uab float32 `xml:"Uab"`
	Ubc float32 `xml:"Ubc"`
	Uca float32 `xml:"Uca"`
	Ia  float32 `xml:"Ia"`
	Ib  float32 `xml:"Ib"`
	Ic  float32 `xml:"Ic"`
	Pa  float32 `xml:"Pa"`
	Qd  int     `xml:"Qd"`
}
