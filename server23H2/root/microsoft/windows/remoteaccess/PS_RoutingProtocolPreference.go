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

// PS_RoutingProtocolPreference struct
type PS_RoutingProtocolPreference struct {
	*cim.WmiInstance
}

func NewPS_RoutingProtocolPreferenceEx1(instance *cim.WmiInstance) (newInstance *PS_RoutingProtocolPreference, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RoutingProtocolPreference{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RoutingProtocolPreferenceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RoutingProtocolPreference, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RoutingProtocolPreference{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Level" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="ProtocolPreference "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RoutingProtocolPreference) Set( /* IN */ Name string,
	/* IN */ Level uint32,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput ProtocolPreference) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", Name, Level, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="cmdletOutput" type="ProtocolPreference "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RoutingProtocolPreference) Get( /* OUT */ cmdletOutput ProtocolPreference) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
