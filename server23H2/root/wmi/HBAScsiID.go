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

// HBAScsiID struct
type HBAScsiID struct {
	*cim.WmiInstance

	//
	OSDeviceName []uint16

	//
	ScsiBusNumber uint32

	//
	ScsiOSLun uint32

	//
	ScsiTargetNumber uint32
}

func NewHBAScsiIDEx1(instance *cim.WmiInstance) (newInstance *HBAScsiID, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HBAScsiID{
		WmiInstance: tmp,
	}
	return
}

func NewHBAScsiIDEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HBAScsiID, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HBAScsiID{
		WmiInstance: tmp,
	}
	return
}

// SetOSDeviceName sets the value of OSDeviceName for the instance
func (instance *HBAScsiID) SetPropertyOSDeviceName(value []uint16) (err error) {
	return instance.SetProperty("OSDeviceName", (value))
}

// GetOSDeviceName gets the value of OSDeviceName for the instance
func (instance *HBAScsiID) GetPropertyOSDeviceName() (value []uint16, err error) {
	retValue, err := instance.GetProperty("OSDeviceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetScsiBusNumber sets the value of ScsiBusNumber for the instance
func (instance *HBAScsiID) SetPropertyScsiBusNumber(value uint32) (err error) {
	return instance.SetProperty("ScsiBusNumber", (value))
}

// GetScsiBusNumber gets the value of ScsiBusNumber for the instance
func (instance *HBAScsiID) GetPropertyScsiBusNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScsiBusNumber")
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

// SetScsiOSLun sets the value of ScsiOSLun for the instance
func (instance *HBAScsiID) SetPropertyScsiOSLun(value uint32) (err error) {
	return instance.SetProperty("ScsiOSLun", (value))
}

// GetScsiOSLun gets the value of ScsiOSLun for the instance
func (instance *HBAScsiID) GetPropertyScsiOSLun() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScsiOSLun")
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

// SetScsiTargetNumber sets the value of ScsiTargetNumber for the instance
func (instance *HBAScsiID) SetPropertyScsiTargetNumber(value uint32) (err error) {
	return instance.SetProperty("ScsiTargetNumber", (value))
}

// GetScsiTargetNumber gets the value of ScsiTargetNumber for the instance
func (instance *HBAScsiID) GetPropertyScsiTargetNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScsiTargetNumber")
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
