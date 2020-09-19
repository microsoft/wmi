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

// MSFT_DtcAdvancedHostSettingTask struct
type MSFT_DtcAdvancedHostSettingTask struct {
	*cim.WmiInstance
}

func NewMSFT_DtcAdvancedHostSettingTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_DtcAdvancedHostSettingTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcAdvancedHostSettingTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DtcAdvancedHostSettingTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DtcAdvancedHostSettingTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcAdvancedHostSettingTask{
		WmiInstance: tmp,
	}
	return
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
