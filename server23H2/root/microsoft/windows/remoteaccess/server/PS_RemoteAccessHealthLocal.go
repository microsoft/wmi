// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_RemoteAccessHealthLocal struct
type PS_RemoteAccessHealthLocal struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessHealthLocalEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessHealthLocal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessHealthLocal{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessHealthLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessHealthLocal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessHealthLocal{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Refresh" type="bool "></param>

// <param name="cmdletOutput" type="RemoteAccessServerHealthLocal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessHealthLocal) Get( /* IN */ Refresh bool,
	/* OUT */ cmdletOutput RemoteAccessServerHealthLocal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Refresh)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
