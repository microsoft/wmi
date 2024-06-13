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

// MSiSCSIInitiator_TargetAddress struct
type MSiSCSIInitiator_TargetAddress struct {
	*cim.WmiInstance

	//
	OSBusNumber uint32

	//
	OSDeviceName string

	//
	OSLunNumber uint32

	//
	OSTargetNumber uint32
}

func NewMSiSCSIInitiator_TargetAddressEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_TargetAddress, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_TargetAddress{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_TargetAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_TargetAddress, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_TargetAddress{
		WmiInstance: tmp,
	}
	return
}

// SetOSBusNumber sets the value of OSBusNumber for the instance
func (instance *MSiSCSIInitiator_TargetAddress) SetPropertyOSBusNumber(value uint32) (err error) {
	return instance.SetProperty("OSBusNumber", (value))
}

// GetOSBusNumber gets the value of OSBusNumber for the instance
func (instance *MSiSCSIInitiator_TargetAddress) GetPropertyOSBusNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSBusNumber")
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

// SetOSDeviceName sets the value of OSDeviceName for the instance
func (instance *MSiSCSIInitiator_TargetAddress) SetPropertyOSDeviceName(value string) (err error) {
	return instance.SetProperty("OSDeviceName", (value))
}

// GetOSDeviceName gets the value of OSDeviceName for the instance
func (instance *MSiSCSIInitiator_TargetAddress) GetPropertyOSDeviceName() (value string, err error) {
	retValue, err := instance.GetProperty("OSDeviceName")
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

// SetOSLunNumber sets the value of OSLunNumber for the instance
func (instance *MSiSCSIInitiator_TargetAddress) SetPropertyOSLunNumber(value uint32) (err error) {
	return instance.SetProperty("OSLunNumber", (value))
}

// GetOSLunNumber gets the value of OSLunNumber for the instance
func (instance *MSiSCSIInitiator_TargetAddress) GetPropertyOSLunNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSLunNumber")
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

// SetOSTargetNumber sets the value of OSTargetNumber for the instance
func (instance *MSiSCSIInitiator_TargetAddress) SetPropertyOSTargetNumber(value uint32) (err error) {
	return instance.SetProperty("OSTargetNumber", (value))
}

// GetOSTargetNumber gets the value of OSTargetNumber for the instance
func (instance *MSiSCSIInitiator_TargetAddress) GetPropertyOSTargetNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSTargetNumber")
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
