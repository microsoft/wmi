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

// PS_BgpRouter struct
type PS_BgpRouter struct {
	*cim.WmiInstance
}

func NewPS_BgpRouterEx1(instance *cim.WmiInstance) (newInstance *PS_BgpRouter, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_BgpRouter{
		WmiInstance: tmp,
	}
	return
}

func NewPS_BgpRouterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_BgpRouter, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_BgpRouter{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="BgpIdentifier" type="string "></param>
// <param name="ClientToClientReflection" type="uint32 "></param>
// <param name="ClusterId" type="uint32 "></param>
// <param name="CompareMEDAcrossASN" type="bool "></param>
// <param name="DefaultGatewayRouting" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="IPv6Routing" type="uint32 "></param>
// <param name="LocalASN" type="uint32 "></param>
// <param name="LocalIPv6Address" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RouteReflector" type="uint32 "></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="TransitRouting" type="uint32 "></param>

// <param name="cmdletOutput" type="BgpRouterConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouter) Add( /* IN */ BgpIdentifier string,
	/* IN */ LocalASN uint32,
	/* IN */ IPv6Routing uint32,
	/* IN */ CompareMEDAcrossASN bool,
	/* IN */ DefaultGatewayRouting bool,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ RoutingDomain string,
	/* IN */ LocalIPv6Address string,
	/* IN */ TransitRouting uint32,
	/* IN */ RouteReflector uint32,
	/* IN */ ClusterId uint32,
	/* IN */ ClientToClientReflection uint32,
	/* OUT */ cmdletOutput BgpRouterConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", BgpIdentifier, LocalASN, IPv6Routing, CompareMEDAcrossASN, DefaultGatewayRouting, PassThru, Force, RoutingDomain, LocalIPv6Address, TransitRouting, RouteReflector, ClusterId, ClientToClientReflection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="RoutingDomain" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouter) Remove( /* IN */ RoutingDomain []string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", RoutingDomain, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="RoutingDomain" type="string []"></param>

// <param name="cmdletOutput" type="BgpRouterConfig []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouter) Get( /* IN */ RoutingDomain []string,
	/* OUT */ cmdletOutput []BgpRouterConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", RoutingDomain)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BgpIdentifier" type="string "></param>
// <param name="ClientToClientReflection" type="uint32 "></param>
// <param name="ClusterId" type="uint32 "></param>
// <param name="CompareMEDAcrossASN" type="bool "></param>
// <param name="DefaultGatewayRouting" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="IPv6Routing" type="uint32 "></param>
// <param name="LocalASN" type="uint32 "></param>
// <param name="LocalIPv6Address" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RouteReflector" type="uint32 "></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="TransitRouting" type="uint32 "></param>

// <param name="cmdletOutput" type="BgpRouterConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouter) Set( /* IN */ BgpIdentifier string,
	/* IN */ LocalASN uint32,
	/* IN */ CompareMEDAcrossASN bool,
	/* IN */ DefaultGatewayRouting bool,
	/* IN */ IPv6Routing uint32,
	/* IN */ RoutingDomain string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ LocalIPv6Address string,
	/* IN */ TransitRouting uint32,
	/* IN */ RouteReflector uint32,
	/* IN */ ClusterId uint32,
	/* IN */ ClientToClientReflection uint32,
	/* OUT */ cmdletOutput BgpRouterConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", BgpIdentifier, LocalASN, CompareMEDAcrossASN, DefaultGatewayRouting, IPv6Routing, RoutingDomain, PassThru, Force, LocalIPv6Address, TransitRouting, RouteReflector, ClusterId, ClientToClientReflection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
