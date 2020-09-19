// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DtcAdvancedSettingTask struct
type MSFT_DtcAdvancedSettingTask struct {
	*cim.WmiInstance
}

func NewMSFT_DtcAdvancedSettingTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_DtcAdvancedSettingTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcAdvancedSettingTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DtcAdvancedSettingTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DtcAdvancedSettingTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcAdvancedSettingTask{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="DtcName" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="Subkey" type="string "></param>

// <param name="cmdletOutput" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcAdvancedSettingTask) Get( /* IN */ DtcName string,
	/* IN */ Subkey string,
	/* IN */ Name string,
	/* OUT */ cmdletOutput string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", DtcName, Subkey, Name)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DtcName" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="Subkey" type="string "></param>
// <param name="Type" type="string "></param>
// <param name="Value" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcAdvancedSettingTask) Set( /* IN */ DtcName string,
	/* IN */ Subkey string,
	/* IN */ Name string,
	/* IN */ Value string,
	/* IN */ Type string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", DtcName, Subkey, Name, Value, Type)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
