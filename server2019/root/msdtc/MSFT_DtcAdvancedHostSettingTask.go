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

// MSFT_DtcAdvancedHostSettingTask struct
type MSFT_DtcAdvancedHostSettingTask struct {
	cim.WmiInstance
}

//

// <param name="Name" type="string "></param>
// <param name="Subkey" type="string "></param>

// <param name="cmdletOutput" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcAdvancedHostSettingTask) Get( /* IN */ Name string,
	/* IN */ Subkey string,
	/* OUT */ cmdletOutput string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Name, Subkey)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Name" type="string "></param>
// <param name="Subkey" type="string "></param>
// <param name="Type" type="string "></param>
// <param name="Value" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcAdvancedHostSettingTask) Set( /* IN */ Name string,
	/* IN */ Subkey string,
	/* IN */ Value string,
	/* IN */ Type string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", Name, Subkey, Value, Type)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
