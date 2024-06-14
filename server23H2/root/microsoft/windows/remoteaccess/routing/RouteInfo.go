// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Routing
//////////////////////////////////////////////
package routing

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RouteInfo struct
type RouteInfo struct {
	*cim.WmiInstance

	//
	IPPrefix string

	//
	Metric uint32

	//
	NextHop string

	//
	PeerId string

	//
	Protocol string
}

func NewRouteInfoEx1(instance *cim.WmiInstance) (newInstance *RouteInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RouteInfo{
		WmiInstance: tmp,
	}
	return
}

func NewRouteInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RouteInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RouteInfo{
		WmiInstance: tmp,
	}
	return
}

// SetIPPrefix sets the value of IPPrefix for the instance
func (instance *RouteInfo) SetPropertyIPPrefix(value string) (err error) {
	return instance.SetProperty("IPPrefix", (value))
}

// GetIPPrefix gets the value of IPPrefix for the instance
func (instance *RouteInfo) GetPropertyIPPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("IPPrefix")
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

// SetMetric sets the value of Metric for the instance
func (instance *RouteInfo) SetPropertyMetric(value uint32) (err error) {
	return instance.SetProperty("Metric", (value))
}

// GetMetric gets the value of Metric for the instance
func (instance *RouteInfo) GetPropertyMetric() (value uint32, err error) {
	retValue, err := instance.GetProperty("Metric")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetNextHop sets the value of NextHop for the instance
func (instance *RouteInfo) SetPropertyNextHop(value string) (err error) {
	return instance.SetProperty("NextHop", (value))
}

// GetNextHop gets the value of NextHop for the instance
func (instance *RouteInfo) GetPropertyNextHop() (value string, err error) {
	retValue, err := instance.GetProperty("NextHop")
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

// SetPeerId sets the value of PeerId for the instance
func (instance *RouteInfo) SetPropertyPeerId(value string) (err error) {
	return instance.SetProperty("PeerId", (value))
}

// GetPeerId gets the value of PeerId for the instance
func (instance *RouteInfo) GetPropertyPeerId() (value string, err error) {
	retValue, err := instance.GetProperty("PeerId")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *RouteInfo) SetPropertyProtocol(value string) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *RouteInfo) GetPropertyProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("Protocol")
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
