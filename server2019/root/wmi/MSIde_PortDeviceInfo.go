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
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSIde_PortDeviceInfo struct
type MSIde_PortDeviceInfo struct {
	*MSIde

	//
	Active bool

	//
	Bus uint8

	//
	InstanceName string

	//
	Lun uint8

	//
	Target uint8
}

func NewMSIde_PortDeviceInfoEx1(instance *cim.WmiInstance) (newInstance *MSIde_PortDeviceInfo, err error) {
	tmp, err := NewMSIdeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSIde_PortDeviceInfo{
		MSIde: tmp,
	}
	return
}

func NewMSIde_PortDeviceInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSIde_PortDeviceInfo, err error) {
	tmp, err := NewMSIdeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSIde_PortDeviceInfo{
		MSIde: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSIde_PortDeviceInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSIde_PortDeviceInfo) GetPropertyActive() (value bool, err error) {
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

// SetBus sets the value of Bus for the instance
func (instance *MSIde_PortDeviceInfo) SetPropertyBus(value uint8) (err error) {
	return instance.SetProperty("Bus", (value))
}

// GetBus gets the value of Bus for the instance
func (instance *MSIde_PortDeviceInfo) GetPropertyBus() (value uint8, err error) {
	retValue, err := instance.GetProperty("Bus")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSIde_PortDeviceInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSIde_PortDeviceInfo) GetPropertyInstanceName() (value string, err error) {
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

// SetLun sets the value of Lun for the instance
func (instance *MSIde_PortDeviceInfo) SetPropertyLun(value uint8) (err error) {
	return instance.SetProperty("Lun", (value))
}

// GetLun gets the value of Lun for the instance
func (instance *MSIde_PortDeviceInfo) GetPropertyLun() (value uint8, err error) {
	retValue, err := instance.GetProperty("Lun")
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

// SetTarget sets the value of Target for the instance
func (instance *MSIde_PortDeviceInfo) SetPropertyTarget(value uint8) (err error) {
	return instance.SetProperty("Target", (value))
}

// GetTarget gets the value of Target for the instance
func (instance *MSIde_PortDeviceInfo) GetPropertyTarget() (value uint8, err error) {
	retValue, err := instance.GetProperty("Target")
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
