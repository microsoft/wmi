// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetIPInterfaceRoute struct
type MSFT_NetIPInterfaceRoute struct { 
	*CIM_RouteUsesEndpoint
}

	func NewMSFT_NetIPInterfaceRouteEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPInterfaceRoute, err error) {tmp, err := NewCIM_RouteUsesEndpointEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetIPInterfaceRoute {
	CIM_RouteUsesEndpoint: tmp,
	}
	return
	}
	

	func NewMSFT_NetIPInterfaceRouteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetIPInterfaceRoute, err error) {tmp, err := NewCIM_RouteUsesEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetIPInterfaceRoute {
	CIM_RouteUsesEndpoint: tmp,
	}
	return
	}
	

