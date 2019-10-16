package soap

import (
	"encoding/xml"
)

type KsoServerResponse struct {
	XMLName xml.Name `xml:"http://casing/api/pb/kso.proto/KsoSrv_types/ Kso_Server_Response"`

	LastSaveEventLogsTs      int64                   `xml:"LastSaveEventLogsTs,omitempty"`      // unix
	LastSaveCalcResultsTs    int64                   `xml:"LastSaveCalcResultsTs,omitempty"`    // unix
	LastSaveMetaInfoTs       int64                   `xml:"LastSaveMetaInfoTs,omitempty"`       // unix
	LastSaveSysMonitorTs     int64                   `xml:"LastSaveSysMonitorTs,omitempty"`     // unix
	LastSavePackageCntTs     int64                   `xml:"LastSavePackageCntTs,omitempty"`     // unix
	LastSaveLostPackageCntTs int64                   `xml:"LastSaveLostPackageCntTs,omitempty"` // unix
	Actions                  []KsoServerActionRecord `xml:"Actions,omitempty"`
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
