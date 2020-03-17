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

// MSFT_DtcDefaultTask struct
type MSFT_DtcDefaultTask struct {
	cim.WmiInstance
}

//

// <param name="cmdletOutput" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcDefaultTask) Get( /* OUT */ cmdletOutput string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcDefaultTask) Set( /* IN */ ServerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", ServerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
