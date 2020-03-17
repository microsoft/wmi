// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.AppBackgroundTask
//////////////////////////////////////////////
package appbackgroundtask

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_BackgroundTask struct
type PS_BackgroundTask struct {
	cim.WmiInstance
}

//

// <param name="TaskID" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BackgroundTask) Start( /* IN */ TaskID []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Start", TaskID)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="TaskID" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BackgroundTask) Unregister( /* IN */ TaskID []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Unregister", TaskID)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="IncludeResourceUsage" type="bool "></param>
// <param name="PackageFamilyName" type="string "></param>

// <param name="cmdletOutput" type="MSFT_BackgroundTask []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BackgroundTask) Get( /* IN */ PackageFamilyName string,
	/* IN */ IncludeResourceUsage bool,
	/* OUT */ cmdletOutput []MSFT_BackgroundTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", PackageFamilyName, IncludeResourceUsage)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BackgroundTask) Enable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BackgroundTask) Disable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="mode" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BackgroundTask) Set( /* IN */ mode string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", mode)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
