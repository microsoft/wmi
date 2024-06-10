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

// PS_BgpRouteFlapDampening struct
type PS_BgpRouteFlapDampening struct {
	*cim.WmiInstance
}

func NewPS_BgpRouteFlapDampeningEx1(instance *cim.WmiInstance) (newInstance *PS_BgpRouteFlapDampening, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_BgpRouteFlapDampening{
		WmiInstance: tmp,
	}
	return
}

func NewPS_BgpRouteFlapDampeningEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_BgpRouteFlapDampening, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_BgpRouteFlapDampening{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Force" type="bool "></param>
// <param name="HalfLife" type="uint32 "></param>
// <param name="HalfLifeUnreachable" type="uint32 "></param>
// <param name="MaxSuppressTime" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="ReuseThreshold" type="uint32 "></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="SuppressThreshold" type="uint32 "></param>

// <param name="cmdletOutput" type="BgpRouteFlapDampeningConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouteFlapDampening) Set( /* IN */ RoutingDomain string,
	/* IN */ ReuseThreshold uint32,
	/* IN */ SuppressThreshold uint32,
	/* IN */ HalfLife uint32,
	/* IN */ HalfLifeUnreachable uint32,
	/* IN */ MaxSuppressTime uint32,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput BgpRouteFlapDampeningConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", RoutingDomain, ReuseThreshold, SuppressThreshold, HalfLife, HalfLifeUnreachable, MaxSuppressTime, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="BgpRouteFlapDampeningConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouteFlapDampening) Get( /* IN */ RoutingDomain string,
	/* OUT */ cmdletOutput BgpRouteFlapDampeningConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", RoutingDomain)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="BgpRouteFlapDampeningConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouteFlapDampening) Enable( /* IN */ RoutingDomain string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput BgpRouteFlapDampeningConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", RoutingDomain, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouteFlapDampening) Disable( /* IN */ RoutingDomain string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable", RoutingDomain, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Prefix" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRouteFlapDampening) Clear( /* IN */ RoutingDomain string,
	/* IN */ Force bool,
	/* IN */ Prefix []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Clear", RoutingDomain, Force, Prefix)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
