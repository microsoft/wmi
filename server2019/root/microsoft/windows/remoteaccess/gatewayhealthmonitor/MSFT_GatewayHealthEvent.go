// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.GatewayHealthMonitor
//
// ////////////////////////////////////////////
package gatewayhealthmonitor

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_GatewayHealthEvent struct
type MSFT_GatewayHealthEvent struct {
	*cim.WmiInstance

	//
	HealthStatus bool

	//
	Id string

	//
	Values []MSFT_GatewayHealthContext
}

func NewMSFT_GatewayHealthEventEx1(instance *cim.WmiInstance) (newInstance *MSFT_GatewayHealthEvent, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_GatewayHealthEvent{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_GatewayHealthEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_GatewayHealthEvent, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_GatewayHealthEvent{
		WmiInstance: tmp,
	}
	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_GatewayHealthEvent) SetPropertyHealthStatus(value bool) (err error) {
	return instance.SetProperty("HealthStatus", (value))
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_GatewayHealthEvent) GetPropertyHealthStatus() (value bool, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
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

// SetId sets the value of Id for the instance
func (instance *MSFT_GatewayHealthEvent) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSFT_GatewayHealthEvent) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetValues sets the value of Values for the instance
func (instance *MSFT_GatewayHealthEvent) SetPropertyValues(value []MSFT_GatewayHealthContext) (err error) {
	return instance.SetProperty("Values", (value))
}

// GetValues gets the value of Values for the instance
func (instance *MSFT_GatewayHealthEvent) GetPropertyValues() (value []MSFT_GatewayHealthContext, err error) {
	retValue, err := instance.GetProperty("Values")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_GatewayHealthContext)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_GatewayHealthContext is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_GatewayHealthContext(valuetmp))
	}

	return
}
