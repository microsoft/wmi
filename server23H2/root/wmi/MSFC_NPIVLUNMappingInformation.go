// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_NPIVLUNMappingInformation struct
type MSFC_NPIVLUNMappingInformation struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string

	//
	OSBus uint8

	//
	OSLUN uint8

	//
	OSTarget uint8

	//
	WWPNPhysicalPort []uint8

	//
	WWPNVirtualPort []uint8
}

func NewMSFC_NPIVLUNMappingInformationEx1(instance *cim.WmiInstance) (newInstance *MSFC_NPIVLUNMappingInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_NPIVLUNMappingInformation{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_NPIVLUNMappingInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_NPIVLUNMappingInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_NPIVLUNMappingInformation{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_NPIVLUNMappingInformation) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_NPIVLUNMappingInformation) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSFC_NPIVLUNMappingInformation) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_NPIVLUNMappingInformation) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetOSBus sets the value of OSBus for the instance
func (instance *MSFC_NPIVLUNMappingInformation) SetPropertyOSBus(value uint8) (err error) {
	return instance.SetProperty("OSBus", (value))
}

// GetOSBus gets the value of OSBus for the instance
func (instance *MSFC_NPIVLUNMappingInformation) GetPropertyOSBus() (value uint8, err error) {
	retValue, err := instance.GetProperty("OSBus")
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

// SetOSLUN sets the value of OSLUN for the instance
func (instance *MSFC_NPIVLUNMappingInformation) SetPropertyOSLUN(value uint8) (err error) {
	return instance.SetProperty("OSLUN", (value))
}

// GetOSLUN gets the value of OSLUN for the instance
func (instance *MSFC_NPIVLUNMappingInformation) GetPropertyOSLUN() (value uint8, err error) {
	retValue, err := instance.GetProperty("OSLUN")
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

// SetOSTarget sets the value of OSTarget for the instance
func (instance *MSFC_NPIVLUNMappingInformation) SetPropertyOSTarget(value uint8) (err error) {
	return instance.SetProperty("OSTarget", (value))
}

// GetOSTarget gets the value of OSTarget for the instance
func (instance *MSFC_NPIVLUNMappingInformation) GetPropertyOSTarget() (value uint8, err error) {
	retValue, err := instance.GetProperty("OSTarget")
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

// SetWWPNPhysicalPort sets the value of WWPNPhysicalPort for the instance
func (instance *MSFC_NPIVLUNMappingInformation) SetPropertyWWPNPhysicalPort(value []uint8) (err error) {
	return instance.SetProperty("WWPNPhysicalPort", (value))
}

// GetWWPNPhysicalPort gets the value of WWPNPhysicalPort for the instance
func (instance *MSFC_NPIVLUNMappingInformation) GetPropertyWWPNPhysicalPort() (value []uint8, err error) {
	retValue, err := instance.GetProperty("WWPNPhysicalPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetWWPNVirtualPort sets the value of WWPNVirtualPort for the instance
func (instance *MSFC_NPIVLUNMappingInformation) SetPropertyWWPNVirtualPort(value []uint8) (err error) {
	return instance.SetProperty("WWPNVirtualPort", (value))
}

// GetWWPNVirtualPort gets the value of WWPNVirtualPort for the instance
func (instance *MSFC_NPIVLUNMappingInformation) GetPropertyWWPNVirtualPort() (value []uint8, err error) {
	retValue, err := instance.GetProperty("WWPNVirtualPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}
