// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DtcDefaultTask struct
type MSFT_DtcDefaultTask struct {
	*cim.WmiInstance
}

func NewMSFT_DtcDefaultTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_DtcDefaultTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcDefaultTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DtcDefaultTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DtcDefaultTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcDefaultTask{
		WmiInstance: tmp,
	}
	return
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
