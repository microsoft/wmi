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

// PS_BgpRoutingPolicy struct
type PS_BgpRoutingPolicy struct {
	*cim.WmiInstance
}

func NewPS_BgpRoutingPolicyEx1(instance *cim.WmiInstance) (newInstance *PS_BgpRoutingPolicy, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_BgpRoutingPolicy{
		WmiInstance: tmp,
	}
	return
}

func NewPS_BgpRoutingPolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_BgpRoutingPolicy, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_BgpRoutingPolicy{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="AddCommunity" type="string []"></param>
// <param name="ClearMED" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="IgnorePrefix" type="string []"></param>
// <param name="MatchASNRange" type="uint32 []"></param>
// <param name="MatchCommunity" type="string []"></param>
// <param name="MatchNextHop" type="string []"></param>
// <param name="MatchPrefix" type="string []"></param>
// <param name="Name" type="string "></param>
// <param name="NewLocalPref" type="uint32 "></param>
// <param name="NewMED" type="uint32 "></param>
// <param name="NewNextHop" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PolicyType" type="uint32 "></param>
// <param name="RemoveAllCommunities" type="bool "></param>
// <param name="RemoveCommunity" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="BgpRoutingPolicyConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRoutingPolicy) Add( /* IN */ PassThru bool,
	/* IN */ RoutingDomain string,
	/* IN */ IgnorePrefix []string,
	/* IN */ Name string,
	/* IN */ MatchCommunity []string,
	/* IN */ NewLocalPref uint32,
	/* IN */ AddCommunity []string,
	/* IN */ Force bool,
	/* IN */ NewMED uint32,
	/* IN */ PolicyType uint32,
	/* IN */ NewNextHop string,
	/* IN */ RemoveCommunity []string,
	/* IN */ MatchPrefix []string,
	/* IN */ MatchASNRange []uint32,
	/* IN */ MatchNextHop []string,
	/* IN */ ClearMED bool,
	/* IN */ RemoveAllCommunities bool,
	/* OUT */ cmdletOutput BgpRoutingPolicyConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", PassThru, RoutingDomain, IgnorePrefix, Name, MatchCommunity, NewLocalPref, AddCommunity, Force, NewMED, PolicyType, NewNextHop, RemoveCommunity, MatchPrefix, MatchASNRange, MatchNextHop, ClearMED, RemoveAllCommunities)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Name" type="string []"></param>
// <param name="PolicyType" type="uint32 "></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="BgpRoutingPolicyConfig []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRoutingPolicy) Get( /* IN */ Name []string,
	/* IN */ PolicyType uint32,
	/* IN */ RoutingDomain string,
	/* OUT */ cmdletOutput []BgpRoutingPolicyConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Name, PolicyType, RoutingDomain)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllRoutingDomains" type="bool "></param>

// <param name="cmdletOutput" type="BgpRoutingPolicyConfig []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRoutingPolicy) GetForAllRoutingDomains( /* IN */ AllRoutingDomains bool,
	/* OUT */ cmdletOutput []BgpRoutingPolicyConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetForAllRoutingDomains", AllRoutingDomains)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Name" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRoutingPolicy) Remove( /* IN */ Name []string,
	/* IN */ RoutingDomain string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", Name, RoutingDomain, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AddCommunity" type="string []"></param>
// <param name="ClearMED" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="IgnorePrefix" type="string []"></param>
// <param name="MatchASNRange" type="uint32 []"></param>
// <param name="MatchCommunity" type="string []"></param>
// <param name="MatchNextHop" type="string []"></param>
// <param name="MatchPrefix" type="string []"></param>
// <param name="Name" type="string "></param>
// <param name="NewLocalPref" type="uint32 "></param>
// <param name="NewMED" type="uint32 "></param>
// <param name="NewNextHop" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PolicyType" type="uint32 "></param>
// <param name="RemoveAllCommunities" type="bool "></param>
// <param name="RemoveCommunity" type="string []"></param>
// <param name="RemovePolicyClause" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="BgpRoutingPolicyConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpRoutingPolicy) Set( /* IN */ Name string,
	/* IN */ PassThru bool,
	/* IN */ PolicyType uint32,
	/* IN */ RoutingDomain string,
	/* IN */ Force bool,
	/* IN */ AddCommunity []string,
	/* IN */ RemoveCommunity []string,
	/* IN */ RemovePolicyClause []string,
	/* IN */ MatchASNRange []uint32,
	/* IN */ IgnorePrefix []string,
	/* IN */ MatchCommunity []string,
	/* IN */ MatchPrefix []string,
	/* IN */ MatchNextHop []string,
	/* IN */ ClearMED bool,
	/* IN */ NewLocalPref uint32,
	/* IN */ NewMED uint32,
	/* IN */ NewNextHop string,
	/* IN */ RemoveAllCommunities bool,
	/* OUT */ cmdletOutput BgpRoutingPolicyConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", Name, PassThru, PolicyType, RoutingDomain, Force, AddCommunity, RemoveCommunity, RemovePolicyClause, MatchASNRange, IgnorePrefix, MatchCommunity, MatchPrefix, MatchNextHop, ClearMED, NewLocalPref, NewMED, NewNextHop, RemoveAllCommunities)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
