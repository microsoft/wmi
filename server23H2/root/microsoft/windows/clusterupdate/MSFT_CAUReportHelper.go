// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_CAUReportHelper struct
type MSFT_CAUReportHelper struct {
	*cim.WmiInstance

	//
	OrchestratorGuid string
}

func NewMSFT_CAUReportHelperEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAUReportHelper, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CAUReportHelper{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CAUReportHelperEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CAUReportHelper, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAUReportHelper{
		WmiInstance: tmp,
	}
	return
}

// SetOrchestratorGuid sets the value of OrchestratorGuid for the instance
func (instance *MSFT_CAUReportHelper) SetPropertyOrchestratorGuid(value string) (err error) {
	return instance.SetProperty("OrchestratorGuid", (value))
}

// GetOrchestratorGuid gets the value of OrchestratorGuid for the instance
func (instance *MSFT_CAUReportHelper) GetPropertyOrchestratorGuid() (value string, err error) {
	retValue, err := instance.GetProperty("OrchestratorGuid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

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
