// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_NextHopRoute struct
type CIM_NextHopRoute struct {
	*CIM_ManagedElement

	//
	AdminDistance uint16

	//
	DestinationAddress string

	//
	IsStatic bool

	//
	RouteMetric uint16

	//
	TypeOfRoute uint16
}

func NewCIM_NextHopRouteEx1(instance *cim.WmiInstance) (newInstance *CIM_NextHopRoute, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_NextHopRoute{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_NextHopRouteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_NextHopRoute, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_NextHopRoute{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetAdminDistance sets the value of AdminDistance for the instance
func (instance *CIM_NextHopRoute) SetPropertyAdminDistance(value uint16) (err error) {
	return instance.SetProperty("AdminDistance", value)
}

// GetAdminDistance gets the value of AdminDistance for the instance
func (instance *CIM_NextHopRoute) GetPropertyAdminDistance() (value uint16, err error) {
	retValue, err := instance.GetProperty("AdminDistance")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestinationAddress sets the value of DestinationAddress for the instance
func (instance *CIM_NextHopRoute) SetPropertyDestinationAddress(value string) (err error) {
	return instance.SetProperty("DestinationAddress", value)
}

// GetDestinationAddress gets the value of DestinationAddress for the instance
func (instance *CIM_NextHopRoute) GetPropertyDestinationAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsStatic sets the value of IsStatic for the instance
func (instance *CIM_NextHopRoute) SetPropertyIsStatic(value bool) (err error) {
	return instance.SetProperty("IsStatic", value)
}

// GetIsStatic gets the value of IsStatic for the instance
func (instance *CIM_NextHopRoute) GetPropertyIsStatic() (value bool, err error) {
	retValue, err := instance.GetProperty("IsStatic")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRouteMetric sets the value of RouteMetric for the instance
func (instance *CIM_NextHopRoute) SetPropertyRouteMetric(value uint16) (err error) {
	return instance.SetProperty("RouteMetric", value)
}

// GetRouteMetric gets the value of RouteMetric for the instance
func (instance *CIM_NextHopRoute) GetPropertyRouteMetric() (value uint16, err error) {
	retValue, err := instance.GetProperty("RouteMetric")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTypeOfRoute sets the value of TypeOfRoute for the instance
func (instance *CIM_NextHopRoute) SetPropertyTypeOfRoute(value uint16) (err error) {
	return instance.SetProperty("TypeOfRoute", value)
}

// GetTypeOfRoute gets the value of TypeOfRoute for the instance
func (instance *CIM_NextHopRoute) GetPropertyTypeOfRoute() (value uint16, err error) {
	retValue, err := instance.GetProperty("TypeOfRoute")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
