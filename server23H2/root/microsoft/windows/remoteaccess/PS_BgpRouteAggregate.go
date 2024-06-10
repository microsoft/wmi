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

// PS_BgpRouteAggregate struct
type PS_BgpRouteAggregate struct {
	*cim.WmiInstance
}

func NewPS_BgpRouteAggregateEx1(instance *cim.WmiInstance) (newInstance *PS_BgpRouteAggregate, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_BgpRouteAggregate{
		WmiInstance: tmp,
	}
	return
}

func NewPS_BgpRouteAggregateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_BgpRouteAggregate, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_BgpRouteAggregate{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="AttributePolicy" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Prefix" type="string "></param>
// <param name="PreserveASPath" type="uint32 "></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="SummaryOnly" type="uint32 "></param>

// <param name="cmdletOutput" type="BgpRouteAggregateConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouteAggregate) Add( /* IN */ RoutingDomain string,
	/* IN */ Prefix string,
	/* IN */ SummaryOnly uint32,
	/* IN */ AttributePolicy []string,
	/* IN */ PreserveASPath uint32,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput BgpRouteAggregateConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", RoutingDomain, Prefix, SummaryOnly, AttributePolicy, PreserveASPath, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AttributePolicy" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Prefix" type="string "></param>
// <param name="PreserveASPath" type="uint32 "></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="SummaryOnly" type="uint32 "></param>

// <param name="cmdletOutput" type="BgpRouteAggregateConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouteAggregate) Set( /* IN */ RoutingDomain string,
	/* IN */ AttributePolicy []string,
	/* IN */ PreserveASPath uint32,
	/* IN */ Prefix string,
	/* IN */ SummaryOnly uint32,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput BgpRouteAggregateConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", RoutingDomain, AttributePolicy, PreserveASPath, Prefix, SummaryOnly, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Prefix" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="BgpRouteAggregateConfig []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouteAggregate) Get( /* IN */ RoutingDomain string,
	/* IN */ Prefix []string,
	/* OUT */ cmdletOutput []BgpRouteAggregateConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", RoutingDomain, Prefix)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Prefix" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouteAggregate) Remove( /* IN */ RoutingDomain string,
	/* IN */ Prefix []string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", RoutingDomain, Prefix, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
