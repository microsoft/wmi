// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SystemConfig_Virtualization struct
type SystemConfig_Virtualization struct {
	*SystemConfig_V2

	//
	HvciEnabled uint8

	//
	HyperVisorEnabled uint8

	//
	Reserved uint8

	//
	VbsEnabled uint8
}

func NewSystemConfig_VirtualizationEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_Virtualization, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_Virtualization{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_VirtualizationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_Virtualization, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_Virtualization{
		SystemConfig_V2: tmp,
	}
	return
}

// SetHvciEnabled sets the value of HvciEnabled for the instance
func (instance *SystemConfig_Virtualization) SetPropertyHvciEnabled(value uint8) (err error) {
	return instance.SetProperty("HvciEnabled", (value))
}

// GetHvciEnabled gets the value of HvciEnabled for the instance
func (instance *SystemConfig_Virtualization) GetPropertyHvciEnabled() (value uint8, err error) {
	retValue, err := instance.GetProperty("HvciEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetHyperVisorEnabled sets the value of HyperVisorEnabled for the instance
func (instance *SystemConfig_Virtualization) SetPropertyHyperVisorEnabled(value uint8) (err error) {
	return instance.SetProperty("HyperVisorEnabled", (value))
}

// GetHyperVisorEnabled gets the value of HyperVisorEnabled for the instance
func (instance *SystemConfig_Virtualization) GetPropertyHyperVisorEnabled() (value uint8, err error) {
	retValue, err := instance.GetProperty("HyperVisorEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *SystemConfig_Virtualization) SetPropertyReserved(value uint8) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *SystemConfig_Virtualization) GetPropertyReserved() (value uint8, err error) {
	retValue, err := instance.GetProperty("Reserved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetVbsEnabled sets the value of VbsEnabled for the instance
func (instance *SystemConfig_Virtualization) SetPropertyVbsEnabled(value uint8) (err error) {
	return instance.SetProperty("VbsEnabled", (value))
}

// GetVbsEnabled gets the value of VbsEnabled for the instance
func (instance *SystemConfig_Virtualization) GetPropertyVbsEnabled() (value uint8, err error) {
	retValue, err := instance.GetProperty("VbsEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
