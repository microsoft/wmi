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

// PS_BgpRoute struct
type PS_BgpRoute struct {
	*cim.WmiInstance
}

func NewPS_BgpRouteEx1(instance *cim.WmiInstance) (newInstance *PS_BgpRoute, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_BgpRoute{
		WmiInstance: tmp,
	}
	return
}

func NewPS_BgpRouteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_BgpRoute, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_BgpRoute{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Network" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="cmdletOutput" type="BgpRouteInfo []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRoute) Get( /* IN */ Network []string,
	/* IN */ RoutingDomain string,
	/* IN */ Type uint32,
	/* OUT */ cmdletOutput []BgpRouteInfo) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Network, RoutingDomain, Type)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
