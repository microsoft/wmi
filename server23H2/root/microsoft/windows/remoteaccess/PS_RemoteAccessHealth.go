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

// PS_RemoteAccessHealth struct
type PS_RemoteAccessHealth struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessHealthEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessHealth, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessHealth{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessHealthEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessHealth, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessHealth{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Cluster" type="bool "></param>
// <param name="ComputerName" type="string "></param>
// <param name="Refresh" type="bool "></param>

// <param name="cmdletOutput" type="RemoteAccessHealthMonitor []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessHealth) GetByCluster( /* IN */ ComputerName string,
	/* IN */ Refresh bool,
	/* IN */ Cluster bool,
	/* OUT */ cmdletOutput []RemoteAccessHealthMonitor) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByCluster", ComputerName, Refresh, Cluster)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>
// <param name="Refresh" type="bool "></param>

// <param name="cmdletOutput" type="RemoteAccessHealthMonitor []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessHealth) GetByEntrypointName( /* IN */ ComputerName string,
	/* IN */ Refresh bool,
	/* IN */ EntrypointName string,
	/* OUT */ cmdletOutput []RemoteAccessHealthMonitor) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByEntrypointName", ComputerName, Refresh, EntrypointName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
