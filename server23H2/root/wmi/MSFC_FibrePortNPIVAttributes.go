// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_FibrePortNPIVAttributes struct
type MSFC_FibrePortNPIVAttributes struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string

	//
	NumberVirtualPorts uint32

	//
	VirtualPorts []MSFC_VirtualFibrePortAttributes

	//
	WWNN []uint8

	//
	WWPN []uint8
}

func NewMSFC_FibrePortNPIVAttributesEx1(instance *cim.WmiInstance) (newInstance *MSFC_FibrePortNPIVAttributes, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_FibrePortNPIVAttributes{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_FibrePortNPIVAttributesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_FibrePortNPIVAttributes, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_FibrePortNPIVAttributes{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_FibrePortNPIVAttributes) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_FibrePortNPIVAttributes) GetPropertyActive() (value bool, err error) {
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
func (instance *MSFC_FibrePortNPIVAttributes) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_FibrePortNPIVAttributes) GetPropertyInstanceName() (value string, err error) {
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

// SetNumberVirtualPorts sets the value of NumberVirtualPorts for the instance
func (instance *MSFC_FibrePortNPIVAttributes) SetPropertyNumberVirtualPorts(value uint32) (err error) {
	return instance.SetProperty("NumberVirtualPorts", (value))
}

// GetNumberVirtualPorts gets the value of NumberVirtualPorts for the instance
func (instance *MSFC_FibrePortNPIVAttributes) GetPropertyNumberVirtualPorts() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberVirtualPorts")
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

// SetVirtualPorts sets the value of VirtualPorts for the instance
func (instance *MSFC_FibrePortNPIVAttributes) SetPropertyVirtualPorts(value []MSFC_VirtualFibrePortAttributes) (err error) {
	return instance.SetProperty("VirtualPorts", (value))
}

// GetVirtualPorts gets the value of VirtualPorts for the instance
func (instance *MSFC_FibrePortNPIVAttributes) GetPropertyVirtualPorts() (value []MSFC_VirtualFibrePortAttributes, err error) {
	retValue, err := instance.GetProperty("VirtualPorts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFC_VirtualFibrePortAttributes)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFC_VirtualFibrePortAttributes is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFC_VirtualFibrePortAttributes(valuetmp))
	}

	return
}

// SetWWNN sets the value of WWNN for the instance
func (instance *MSFC_FibrePortNPIVAttributes) SetPropertyWWNN(value []uint8) (err error) {
	return instance.SetProperty("WWNN", (value))
}

// GetWWNN gets the value of WWNN for the instance
func (instance *MSFC_FibrePortNPIVAttributes) GetPropertyWWNN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("WWNN")
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

// SetWWPN sets the value of WWPN for the instance
func (instance *MSFC_FibrePortNPIVAttributes) SetPropertyWWPN(value []uint8) (err error) {
	return instance.SetProperty("WWPN", (value))
}

// GetWWPN gets the value of WWPN for the instance
func (instance *MSFC_FibrePortNPIVAttributes) GetPropertyWWPN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("WWPN")
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
