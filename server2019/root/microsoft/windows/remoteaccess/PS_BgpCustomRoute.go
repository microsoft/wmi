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

// PS_BgpCustomRoute struct
type PS_BgpCustomRoute struct { 
	*cim.WmiInstance
}

	func NewPS_BgpCustomRouteEx1(instance *cim.WmiInstance) (newInstance *PS_BgpCustomRoute, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_BgpCustomRoute {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_BgpCustomRouteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_BgpCustomRoute, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_BgpCustomRoute {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="Interface" type="string []"></param>
// <param name="Network" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="BgpCustomNetworkInfo "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpCustomRoute) Add( /* IN */ PassThru bool,
 /* IN */ RoutingDomain string,
 /* IN */ Interface []string,
 /* IN */ Network []string,
 /* OUT */ cmdletOutput BgpCustomNetworkInfo) (result uint32, err error) {retVal, err := instance.InvokeMethod("Add" , PassThru, RoutingDomain, Interface, Network)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Force" type="bool "></param>
// <param name="Interface" type="string []"></param>
// <param name="Network" type="string []"></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpCustomRoute) Remove( /* IN */ Network []string,
 /* IN */ Interface []string,
 /* IN */ Force bool,
 /* IN */ RoutingDomain string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Remove" , Network, Interface, Force, RoutingDomain);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="BgpCustomNetworkInfo "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpCustomRoute) Get( /* IN */ RoutingDomain string,
 /* OUT */ cmdletOutput BgpCustomNetworkInfo) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" , RoutingDomain)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="AllRoutingDomains" type="bool "></param>

// <param name="cmdletOutput" type="BgpCustomNetworkInfo []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_BgpCustomRoute) GetForAllRoutingDomains( /* IN */ AllRoutingDomains bool,
 /* OUT */ cmdletOutput []BgpCustomNetworkInfo) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetForAllRoutingDomains" , AllRoutingDomains)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


