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

// MSFT_DtcLogTask struct
type MSFT_DtcLogTask struct {
	cim.WmiInstance
}

//

// <param name="DtcName" type="string "></param>

// <param name="cmdletOutput" type="DtcLogFileSettings "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcLogTask) Get( /* IN */ DtcName string,
	/* OUT */ cmdletOutput DtcLogFileSettings) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", DtcName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DtcName" type="string "></param>
// <param name="MaxSizeInMB" type="uint32 "></param>
// <param name="Path" type="string "></param>
// <param name="SizeInMB" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcLogTask) Set( /* IN */ DtcName string,
	/* IN */ Path string,
	/* IN */ SizeInMB uint32,
	/* IN */ MaxSizeInMB uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", DtcName, Path, SizeInMB, MaxSizeInMB)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DtcName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcLogTask) Reset( /* IN */ DtcName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Reset", DtcName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
