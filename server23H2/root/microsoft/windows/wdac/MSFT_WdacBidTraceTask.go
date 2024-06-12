// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.Wdac
//////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_WdacBidTraceTask struct
type MSFT_WdacBidTraceTask struct {
	*cim.WmiInstance
}

func NewMSFT_WdacBidTraceTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_WdacBidTraceTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_WdacBidTraceTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_WdacBidTraceTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_WdacBidTraceTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_WdacBidTraceTask{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="IncludeAllApplications" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) EnableByAllApp( /* IN */ PassThru bool,
	/* IN */ IncludeAllApplications bool,
	/* IN */ Platform string,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableByAllApp", PassThru, IncludeAllApplications, Platform)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Folder" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) EnableByFolder( /* IN */ PassThru bool,
	/* IN */ Platform string,
	/* IN */ Folder string,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableByFolder", PassThru, Platform, Folder)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InputObject" type="MSFT_WdacBidTrace []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) EnableByInputObject( /* IN */ PassThru bool,
	/* IN */ InputObject []MSFT_WdacBidTrace,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableByInputObject", PassThru, InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PassThru" type="bool "></param>
// <param name="Path" type="string "></param>
// <param name="Platform" type="string "></param>
// <param name="ProcessId" type="uint32 "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) EnableByPath( /* IN */ PassThru bool,
	/* IN */ Path string,
	/* IN */ Platform string,
	/* IN */ ProcessId uint32,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableByPath", PassThru, Path, Platform, ProcessId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IncludeAllApplications" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) DisableByAllApp( /* IN */ PassThru bool,
	/* IN */ IncludeAllApplications bool,
	/* IN */ Platform string,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisableByAllApp", PassThru, IncludeAllApplications, Platform)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Folder" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) DisableByFolder( /* IN */ PassThru bool,
	/* IN */ Folder string,
	/* IN */ Platform string,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisableByFolder", PassThru, Folder, Platform)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InputObject" type="MSFT_WdacBidTrace []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) DisableByInputObject( /* IN */ PassThru bool,
	/* IN */ InputObject []MSFT_WdacBidTrace,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisableByInputObject", PassThru, InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PassThru" type="bool "></param>
// <param name="Path" type="string "></param>
// <param name="Platform" type="string "></param>
// <param name="ProcessId" type="uint32 "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) DisableByPath( /* IN */ PassThru bool,
	/* IN */ Path string,
	/* IN */ ProcessId uint32,
	/* IN */ Platform string,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisableByPath", PassThru, Path, ProcessId, Platform)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IncludeAllApplications" type="bool "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) GetByAllApp( /* IN */ Platform string,
	/* IN */ IncludeAllApplications bool,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByAllApp", Platform, IncludeAllApplications)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Folder" type="string "></param>
// <param name="Platform" type="string "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) GetByFolder( /* IN */ Platform string,
	/* IN */ Folder string,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByFolder", Platform, Folder)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Path" type="string "></param>
// <param name="Platform" type="string "></param>
// <param name="ProcessId" type="uint32 "></param>

// <param name="cmdletOutput" type="MSFT_WdacBidTrace []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_WdacBidTraceTask) GetByPath( /* IN */ Platform string,
	/* IN */ Path string,
	/* IN */ ProcessId uint32,
	/* OUT */ cmdletOutput []MSFT_WdacBidTrace) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByPath", Platform, Path, ProcessId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
