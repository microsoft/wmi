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
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RouteEvent2 struct
type RouteEvent2 struct {
	*cim.WmiInstance

	//
	Routes []RouteInfo

	//
	RoutingDomainID string
}

func NewRouteEvent2Ex1(instance *cim.WmiInstance) (newInstance *RouteEvent2, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RouteEvent2{
		WmiInstance: tmp,
	}
	return
}

func NewRouteEvent2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RouteEvent2, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RouteEvent2{
		WmiInstance: tmp,
	}
	return
}

// SetRoutes sets the value of Routes for the instance
func (instance *RouteEvent2) SetPropertyRoutes(value []RouteInfo) (err error) {
	return instance.SetProperty("Routes", (value))
}

// GetRoutes gets the value of Routes for the instance
func (instance *RouteEvent2) GetPropertyRoutes() (value []RouteInfo, err error) {
	retValue, err := instance.GetProperty("Routes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(RouteInfo)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " RouteInfo is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, RouteInfo(valuetmp))
	}

	return
}

// SetRoutingDomainID sets the value of RoutingDomainID for the instance
func (instance *RouteEvent2) SetPropertyRoutingDomainID(value string) (err error) {
	return instance.SetProperty("RoutingDomainID", (value))
}

// GetRoutingDomainID gets the value of RoutingDomainID for the instance
func (instance *RouteEvent2) GetPropertyRoutingDomainID() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomainID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
