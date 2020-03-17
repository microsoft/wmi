// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DtcTask struct
type MSFT_DtcTask struct {
	cim.WmiInstance
}

//

// <param name="DtcName" type="string "></param>

// <param name="cmdletOutput" type="DtcInstance []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTask) Get( /* IN */ DtcName string,
	/* OUT */ cmdletOutput []DtcInstance) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", DtcName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LogPath" type="string "></param>
// <param name="StartType" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTask) Install( /* IN */ LogPath string,
	/* IN */ StartType string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Install", LogPath, StartType)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTask) Uninstall() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Uninstall")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DtcName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTask) Start( /* IN */ DtcName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Start", DtcName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DtcName" type="string "></param>
// <param name="Recursive" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTask) Stop( /* IN */ DtcName string,
	/* IN */ Recursive bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Stop", DtcName, Recursive)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
