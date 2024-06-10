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

// PS_RemoteAccessLoadBalancer struct
type PS_RemoteAccessLoadBalancer struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessLoadBalancerEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessLoadBalancer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessLoadBalancer{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessLoadBalancerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessLoadBalancer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessLoadBalancer{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="Disable" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="RemoteAccessLoadBalancer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLoadBalancer) SetByDisableLoadBalancing( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ Disable bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput RemoteAccessLoadBalancer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByDisableLoadBalancing", ComputerName, PassThru, Disable, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="InternalDedicatedIPAddress" type="string []"></param>
// <param name="InternalVirtualIPAddress" type="string []"></param>
// <param name="InternetDedicatedIPAddress" type="string []"></param>
// <param name="InternetVirtualIPAddress" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="UseThirdPartyLoadBalancer" type="bool "></param>

// <param name="cmdletOutput" type="RemoteAccessLoadBalancer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLoadBalancer) SetByEnableLoadBalancing( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ UseThirdPartyLoadBalancer bool,
	/* IN */ InternetDedicatedIPAddress []string,
	/* IN */ InternalDedicatedIPAddress []string,
	/* IN */ InternetVirtualIPAddress []string,
	/* IN */ InternalVirtualIPAddress []string,
	/* OUT */ cmdletOutput RemoteAccessLoadBalancer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByEnableLoadBalancing", ComputerName, PassThru, UseThirdPartyLoadBalancer, InternetDedicatedIPAddress, InternalDedicatedIPAddress, InternetVirtualIPAddress, InternalVirtualIPAddress)
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
// <param name="ThirdPartyLoadBalancer" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessLoadBalancer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLoadBalancer) SetByThirdPartyLoadBalancer( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ ThirdPartyLoadBalancer string,
	/* OUT */ cmdletOutput RemoteAccessLoadBalancer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByThirdPartyLoadBalancer", ComputerName, PassThru, ThirdPartyLoadBalancer)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Brief" type="bool "></param>
// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessLoadBalancer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLoadBalancer) Get( /* IN */ ComputerName string,
	/* IN */ EntrypointName string,
	/* IN */ Brief bool,
	/* OUT */ cmdletOutput RemoteAccessLoadBalancer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", ComputerName, EntrypointName, Brief)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
