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

// PS_BgpRoutingPolicyForPeer struct
type PS_BgpRoutingPolicyForPeer struct {
	*cim.WmiInstance
}

func NewPS_BgpRoutingPolicyForPeerEx1(instance *cim.WmiInstance) (newInstance *PS_BgpRoutingPolicyForPeer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_BgpRoutingPolicyForPeer{
		WmiInstance: tmp,
	}
	return
}

func NewPS_BgpRoutingPolicyForPeerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_BgpRoutingPolicyForPeer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_BgpRoutingPolicyForPeer{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Direction" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="PeerName" type="string []"></param>
// <param name="PolicyName" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRoutingPolicyForPeer) Add( /* IN */ PeerName []string,
	/* IN */ PolicyName []string,
	/* IN */ Direction uint32,
	/* IN */ RoutingDomain string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Add", PeerName, PolicyName, Direction, RoutingDomain, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Direction" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="PeerName" type="string []"></param>
// <param name="PolicyName" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRoutingPolicyForPeer) Remove( /* IN */ PeerName []string,
	/* IN */ PolicyName []string,
	/* IN */ Direction uint32,
	/* IN */ Force bool,
	/* IN */ RoutingDomain string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", PeerName, PolicyName, Direction, Force, RoutingDomain)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Direction" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="PeerName" type="string []"></param>
// <param name="PolicyName" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRoutingPolicyForPeer) Set( /* IN */ PolicyName []string,
	/* IN */ PeerName []string,
	/* IN */ Direction uint32,
	/* IN */ Force bool,
	/* IN */ RoutingDomain string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", PolicyName, PeerName, Direction, Force, RoutingDomain)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
