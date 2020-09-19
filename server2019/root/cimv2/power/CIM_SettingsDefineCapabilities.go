// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_SettingsDefineCapabilities struct
type CIM_SettingsDefineCapabilities struct {
	*CIM_Component

	//
	PropertyPolicy uint16

	//
	ValueRange uint16

	//
	ValueRole uint16
}

func NewCIM_SettingsDefineCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *CIM_SettingsDefineCapabilities, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SettingsDefineCapabilities{
		CIM_Component: tmp,
	}
	return
}

func NewCIM_SettingsDefineCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SettingsDefineCapabilities, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SettingsDefineCapabilities{
		CIM_Component: tmp,
	}
	return
}

// SetPropertyPolicy sets the value of PropertyPolicy for the instance
func (instance *CIM_SettingsDefineCapabilities) SetPropertyPropertyPolicy(value uint16) (err error) {
	return instance.SetProperty("PropertyPolicy", (value))
}

// GetPropertyPolicy gets the value of PropertyPolicy for the instance
func (instance *CIM_SettingsDefineCapabilities) GetPropertyPropertyPolicy() (value uint16, err error) {
	retValue, err := instance.GetProperty("PropertyPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetValueRange sets the value of ValueRange for the instance
func (instance *CIM_SettingsDefineCapabilities) SetPropertyValueRange(value uint16) (err error) {
	return instance.SetProperty("ValueRange", (value))
}

// GetValueRange gets the value of ValueRange for the instance
func (instance *CIM_SettingsDefineCapabilities) GetPropertyValueRange() (value uint16, err error) {
	retValue, err := instance.GetProperty("ValueRange")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetValueRole sets the value of ValueRole for the instance
func (instance *CIM_SettingsDefineCapabilities) SetPropertyValueRole(value uint16) (err error) {
	return instance.SetProperty("ValueRole", (value))
}

// GetValueRole gets the value of ValueRole for the instance
func (instance *CIM_SettingsDefineCapabilities) GetPropertyValueRole() (value uint16, err error) {
	retValue, err := instance.GetProperty("ValueRole")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
