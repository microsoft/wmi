// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_VpnUser struct
type PS_VpnUser struct {
	*cim.WmiInstance
}

func NewPS_VpnUserEx1(instance *cim.WmiInstance) (newInstance *PS_VpnUser, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnUser{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnUserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnUser, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnUser{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="HostIPAddress" type="string []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnUser) DisconnectByHostIP( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ HostIPAddress []string,
	/* OUT */ cmdletOutput []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisconnectByHostIP", ComputerName, PassThru, HostIPAddress)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="UserName" type="string []"></param>

// <param name="cmdletOutput" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnUser) DisconnectByUserName( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ UserName []string,
	/* OUT */ cmdletOutput []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisconnectByUserName", ComputerName, PassThru, UserName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
