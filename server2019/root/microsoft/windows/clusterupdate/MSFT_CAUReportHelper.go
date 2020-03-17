// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_CAUReportHelper struct
type MSFT_CAUReportHelper struct {
	cim.WmiInstance

	//
	OrchestratorGuid string
}

// SetOrchestratorGuid sets the value of OrchestratorGuid for the instance
func (instance *MSFT_CAUReportHelper) SetPropertyOrchestratorGuid(value string) (err error) {
	return instance.SetProperty("OrchestratorGuid", value)
}

// GetOrchestratorGuid gets the value of OrchestratorGuid for the instance
func (instance *MSFT_CAUReportHelper) GetPropertyOrchestratorGuid() (value string, err error) {
	retValue, err := instance.GetProperty("OrchestratorGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReportID" type="MSFT_CAURun_Report_ID []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUReportHelper) GetReportIDs( /* OUT */ ReportID []MSFT_CAURun_Report_ID) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetReportIDs")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ChunkSize" type="uint32 "></param>
// <param name="ReportId" type="MSFT_CAURun_Report_ID "></param>

// <param name="ReportChunks" type="MSFT_CAURun_Report_Chunk []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUReportHelper) GetReport( /* IN */ ReportId MSFT_CAURun_Report_ID,
	/* IN */ ChunkSize uint32,
	/* OUT */ ReportChunks []MSFT_CAURun_Report_Chunk) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetReport", ReportId, ChunkSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LastChunk" type="bool "></param>
// <param name="ReportChunk" type="MSFT_CAURun_Report_Chunk "></param>
// <param name="ReportSize" type="uint64 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUReportHelper) PutReport( /* IN */ ReportSize uint64,
	/* IN */ ReportChunk MSFT_CAURun_Report_Chunk,
	/* IN */ LastChunk bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("PutReport", ReportSize, ReportChunk, LastChunk)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
