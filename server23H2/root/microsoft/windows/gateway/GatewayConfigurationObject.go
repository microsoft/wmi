// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Gateway
//////////////////////////////////////////////
package gateway

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// GatewayConfigurationObject struct
type GatewayConfigurationObject struct {
	*cim.WmiInstance

	//
	RoutingDomainName string

	//
	Standby bool

	//
	Version string
}

func NewGatewayConfigurationObjectEx1(instance *cim.WmiInstance) (newInstance *GatewayConfigurationObject, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &GatewayConfigurationObject{
		WmiInstance: tmp,
	}
	return
}

func NewGatewayConfigurationObjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *GatewayConfigurationObject, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &GatewayConfigurationObject{
		WmiInstance: tmp,
	}
	return
}

// SetRoutingDomainName sets the value of RoutingDomainName for the instance
func (instance *GatewayConfigurationObject) SetPropertyRoutingDomainName(value string) (err error) {
	return instance.SetProperty("RoutingDomainName", (value))
}

// GetRoutingDomainName gets the value of RoutingDomainName for the instance
func (instance *GatewayConfigurationObject) GetPropertyRoutingDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomainName")
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

// SetStandby sets the value of Standby for the instance
func (instance *GatewayConfigurationObject) SetPropertyStandby(value bool) (err error) {
	return instance.SetProperty("Standby", (value))
}

// GetStandby gets the value of Standby for the instance
func (instance *GatewayConfigurationObject) GetPropertyStandby() (value bool, err error) {
	retValue, err := instance.GetProperty("Standby")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetVersion sets the value of Version for the instance
func (instance *GatewayConfigurationObject) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *GatewayConfigurationObject) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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
