// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Server
//
// ////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_RemoteAccessUserActivityLocal struct
type PS_RemoteAccessUserActivityLocal struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessUserActivityLocalEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessUserActivityLocal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessUserActivityLocal{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessUserActivityLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessUserActivityLocal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessUserActivityLocal{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="EndDateTime" type="string "></param>
// <param name="HostIPAddress" type="string "></param>
// <param name="StartDateTime" type="string "></param>
// <param name="UserName" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessUserActivityLocal []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessUserActivityLocal) GetByAccountingUserActivity( /* IN */ StartDateTime string,
	/* IN */ EndDateTime string,
	/* IN */ HostIPAddress string,
	/* IN */ UserName string,
	/* OUT */ cmdletOutput []RemoteAccessUserActivityLocal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByAccountingUserActivity", StartDateTime, EndDateTime, HostIPAddress, UserName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="HostIPAddress" type="string "></param>
// <param name="UserName" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessUserActivityLocal []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessUserActivityLocal) GetByActiveUserActivity( /* IN */ UserName string,
	/* IN */ HostIPAddress string,
	/* OUT */ cmdletOutput []RemoteAccessUserActivityLocal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByActiveUserActivity", UserName, HostIPAddress)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SessionId" type="uint64 "></param>

// <param name="cmdletOutput" type="RemoteAccessUserActivityLocal []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessUserActivityLocal) GetBySessionId( /* IN */ SessionId uint64,
	/* OUT */ cmdletOutput []RemoteAccessUserActivityLocal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetBySessionId", SessionId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
