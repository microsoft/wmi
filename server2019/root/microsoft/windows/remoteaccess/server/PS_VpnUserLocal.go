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

// PS_VpnUserLocal struct
type PS_VpnUserLocal struct {
	*cim.WmiInstance
}

func NewPS_VpnUserLocalEx1(instance *cim.WmiInstance) (newInstance *PS_VpnUserLocal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnUserLocal{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnUserLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnUserLocal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnUserLocal{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="UserName" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="StatusCodes" type="uint32 []"></param>
func (instance *PS_VpnUserLocal) DisconnectByUserName( /* IN */ UserName []string,
	/* OUT */ StatusCodes []uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisconnectByUserName", UserName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="HostIPAddress" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="StatusCodes" type="uint32 []"></param>
func (instance *PS_VpnUserLocal) DisconnectByHostIP( /* IN */ HostIPAddress []string,
	/* OUT */ StatusCodes []uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisconnectByHostIP", HostIPAddress)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
