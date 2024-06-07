// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// PS_BgpPeer struct
type PS_BgpPeer struct { 
	*cim.WmiInstance
}

	func NewPS_BgpPeerEx1(instance *cim.WmiInstance) (newInstance *PS_BgpPeer, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_BgpPeer {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_BgpPeerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_BgpPeer, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_BgpPeer {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="GracefulRestart" type="bool "></param>
// <param name="GracefulRestartTime" type="uint16 "></param>
// <param name="HoldTimeSec" type="uint16 "></param>
// <param name="IdleHoldTimeSec" type="uint16 "></param>
// <param name="LocalASN" type="uint32 "></param>
// <param name="LocalIPAddress" type="string "></param>
// <param name="MaxAllowedPrefix" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="OperationMode" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PeerASN" type="uint32 "></param>
// <param name="PeeringMode" type="uint32 "></param>
// <param name="PeerIPAddress" type="string "></param>
// <param name="RouteReflectorClient" type="bool "></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="Weight" type="uint16 "></param>

// <param name="cmdletOutput" type="BgpPeerConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpPeer) Add( /* IN */ Name string,
 /* IN */ LocalIPAddress string,
 /* IN */ PeerIPAddress string,
 /* IN */ LocalASN uint32,
 /* IN */ PeerASN uint32,
 /* IN */ OperationMode uint32,
 /* IN */ PeeringMode uint32,
 /* IN */ HoldTimeSec uint16,
 /* IN */ IdleHoldTimeSec uint16,
 /* IN */ Weight uint16,
 /* IN */ RouteReflectorClient bool,
 /* IN */ MaxAllowedPrefix uint32,
 /* IN */ PassThru bool,
 /* IN */ RoutingDomain string,
 /* IN */ GracefulRestart bool,
 /* IN */ GracefulRestartTime uint16,
 /* OUT */ cmdletOutput BgpPeerConfig) (result uint32, err error) {retVal, err := instance.InvokeMethod("Add" , Name, LocalIPAddress, PeerIPAddress, LocalASN, PeerASN, OperationMode, PeeringMode, HoldTimeSec, IdleHoldTimeSec, Weight, RouteReflectorClient, MaxAllowedPrefix, PassThru, RoutingDomain, GracefulRestart, GracefulRestartTime)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ClearPrefixLimit" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="GracefulRestart" type="bool "></param>
// <param name="GracefulRestartTime" type="uint16 "></param>
// <param name="HoldTimeSec" type="uint16 "></param>
// <param name="IdleHoldTimeSec" type="uint16 "></param>
// <param name="LocalASN" type="uint32 "></param>
// <param name="LocalIPAddress" type="string "></param>
// <param name="MaxAllowedPrefix" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="OperationMode" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PeerASN" type="uint32 "></param>
// <param name="PeeringMode" type="uint32 "></param>
// <param name="PeerIPAddress" type="string "></param>
// <param name="RouteReflectorClient" type="bool "></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="Weight" type="uint16 "></param>

// <param name="cmdletOutput" type="BgpPeerConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpPeer) Set( /* IN */ Name string,
 /* IN */ LocalIPAddress string,
 /* IN */ PeerIPAddress string,
 /* IN */ LocalASN uint32,
 /* IN */ PeerASN uint32,
 /* IN */ OperationMode uint32,
 /* IN */ PeeringMode uint32,
 /* IN */ HoldTimeSec uint16,
 /* IN */ IdleHoldTimeSec uint16,
 /* IN */ Weight uint16,
 /* IN */ RouteReflectorClient bool,
 /* IN */ MaxAllowedPrefix uint32,
 /* IN */ RoutingDomain string,
 /* IN */ PassThru bool,
 /* IN */ Force bool,
 /* IN */ ClearPrefixLimit bool,
 /* IN */ GracefulRestart bool,
 /* IN */ GracefulRestartTime uint16,
 /* OUT */ cmdletOutput BgpPeerConfig) (result uint32, err error) {retVal, err := instance.InvokeMethod("Set" , Name, LocalIPAddress, PeerIPAddress, LocalASN, PeerASN, OperationMode, PeeringMode, HoldTimeSec, IdleHoldTimeSec, Weight, RouteReflectorClient, MaxAllowedPrefix, RoutingDomain, PassThru, Force, ClearPrefixLimit, GracefulRestart, GracefulRestartTime)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Name" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="BgpPeerConfig []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpPeer) Get( /* IN */ Name []string,
 /* IN */ RoutingDomain string,
 /* OUT */ cmdletOutput []BgpPeerConfig) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" , Name, RoutingDomain)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="AllRoutingDomains" type="bool "></param>

// <param name="cmdletOutput" type="BgpPeerConfig []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpPeer) GetForAllRoutingDomains( /* IN */ AllRoutingDomains bool,
 /* OUT */ cmdletOutput []BgpPeerConfig) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetForAllRoutingDomains" , AllRoutingDomains)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Force" type="bool "></param>
// <param name="Name" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpPeer) Remove( /* IN */ Name []string,
 /* IN */ RoutingDomain string,
 /* IN */ Force bool) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Remove" , Name, RoutingDomain, Force);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="Name" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpPeer) Start( /* IN */ RoutingDomain string,
 /* IN */ Name []string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Start" , RoutingDomain, Name);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="Force" type="bool "></param>
// <param name="Name" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpPeer) Stop( /* IN */ Name []string,
 /* IN */ RoutingDomain string,
 /* IN */ Force bool) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Stop" , Name, RoutingDomain, Force);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


