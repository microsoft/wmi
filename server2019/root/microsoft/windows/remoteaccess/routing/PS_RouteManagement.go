// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Routing
//
// ////////////////////////////////////////////
package routing

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_RouteManagement struct
type PS_RouteManagement struct {
	*cim.WmiInstance
}

func NewPS_RouteManagementEx1(instance *cim.WmiInstance) (newInstance *PS_RouteManagement, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RouteManagement{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RouteManagementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RouteManagement, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RouteManagement{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="RoutingDomainID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RouteManagement) Register( /* IN */ RoutingDomainID string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Register", RoutingDomainID)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="RoutingDomainID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RouteManagement) Unregister( /* IN */ RoutingDomainID string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Unregister", RoutingDomainID)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="RoutingDomainID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Routes" type="RouteInfo []"></param>
func (instance *PS_RouteManagement) GetRoutes( /* IN */ RoutingDomainID string,
	/* OUT */ Routes []RouteInfo) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetRoutes", RoutingDomainID)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
