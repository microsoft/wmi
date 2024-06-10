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

// MSFC_NPIVCapabilities struct
type MSFC_NPIVCapabilities struct {
	*cim.WmiInstance

	//
	Active bool

	//
	DhChapAvailableOnPhysicalPort bool

	//
	DhChapAvailableOnVirtualPorts bool

	//
	InstanceName string

	//
	MaxVirtualPortCount uint16
}

func NewMSFC_NPIVCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSFC_NPIVCapabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_NPIVCapabilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_NPIVCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_NPIVCapabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_NPIVCapabilities{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_NPIVCapabilities) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_NPIVCapabilities) GetPropertyActive() (value bool, err error) {
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

// SetDhChapAvailableOnPhysicalPort sets the value of DhChapAvailableOnPhysicalPort for the instance
func (instance *MSFC_NPIVCapabilities) SetPropertyDhChapAvailableOnPhysicalPort(value bool) (err error) {
	return instance.SetProperty("DhChapAvailableOnPhysicalPort", (value))
}

// GetDhChapAvailableOnPhysicalPort gets the value of DhChapAvailableOnPhysicalPort for the instance
func (instance *MSFC_NPIVCapabilities) GetPropertyDhChapAvailableOnPhysicalPort() (value bool, err error) {
	retValue, err := instance.GetProperty("DhChapAvailableOnPhysicalPort")
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

// SetDhChapAvailableOnVirtualPorts sets the value of DhChapAvailableOnVirtualPorts for the instance
func (instance *MSFC_NPIVCapabilities) SetPropertyDhChapAvailableOnVirtualPorts(value bool) (err error) {
	return instance.SetProperty("DhChapAvailableOnVirtualPorts", (value))
}

// GetDhChapAvailableOnVirtualPorts gets the value of DhChapAvailableOnVirtualPorts for the instance
func (instance *MSFC_NPIVCapabilities) GetPropertyDhChapAvailableOnVirtualPorts() (value bool, err error) {
	retValue, err := instance.GetProperty("DhChapAvailableOnVirtualPorts")
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
func (instance *MSFC_NPIVCapabilities) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_NPIVCapabilities) GetPropertyInstanceName() (value string, err error) {
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

// SetMaxVirtualPortCount sets the value of MaxVirtualPortCount for the instance
func (instance *MSFC_NPIVCapabilities) SetPropertyMaxVirtualPortCount(value uint16) (err error) {
	return instance.SetProperty("MaxVirtualPortCount", (value))
}

// GetMaxVirtualPortCount gets the value of MaxVirtualPortCount for the instance
func (instance *MSFC_NPIVCapabilities) GetPropertyMaxVirtualPortCount() (value uint16, err error) {
	retValue, err := instance.GetProperty("MaxVirtualPortCount")
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
