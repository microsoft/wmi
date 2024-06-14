// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_RemoteAccessLoadBalancerNode struct
type PS_RemoteAccessLoadBalancerNode struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessLoadBalancerNodeEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessLoadBalancerNode, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessLoadBalancerNode{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessLoadBalancerNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessLoadBalancerNode, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessLoadBalancerNode{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="InternalInterface" type="string "></param>
// <param name="InternetInterface" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RemoteAccessServer" type="string "></param>
// <param name="VpnIPAddressRange" type="string []"></param>

// <param name="cmdletOutput" type="RemoteAccessLoadBalancerNode "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLoadBalancerNode) Add( /* IN */ RemoteAccessServer string,
	/* IN */ InternalInterface string,
	/* IN */ InternetInterface string,
	/* IN */ VpnIPAddressRange []string,
	/* IN */ ComputerName string,
	/* IN */ Force bool,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput RemoteAccessLoadBalancerNode) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", RemoteAccessServer, InternalInterface, InternetInterface, VpnIPAddressRange, ComputerName, Force, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RemoteAccessServer" type="string "></param>

// <param name="cmdletOutput" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLoadBalancerNode) Remove( /* IN */ RemoteAccessServer string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", RemoteAccessServer, ComputerName, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
