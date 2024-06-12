// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_PciDeviceProperty struct
type MSNdis_PciDeviceProperty struct {
	*MSNdis

	//
	CurrentLinkSpeed uint32

	//
	CurrentLinkWidth uint32

	//
	CurrentPayloadSize uint32

	//
	CurrentSpeedAndMode uint32

	//
	DeviceType uint32

	//
	Header MSNdis_ObjectHeader

	//
	MaxLinkSpeed uint32

	//
	MaxLinkWidth uint32

	//
	MaxPayloadSize uint32

	//
	MaxReadRequestSize uint32
}

func NewMSNdis_PciDevicePropertyEx1(instance *cim.WmiInstance) (newInstance *MSNdis_PciDeviceProperty, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PciDeviceProperty{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_PciDevicePropertyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_PciDeviceProperty, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PciDeviceProperty{
		MSNdis: tmp,
	}
	return
}

// SetCurrentLinkSpeed sets the value of CurrentLinkSpeed for the instance
func (instance *MSNdis_PciDeviceProperty) SetPropertyCurrentLinkSpeed(value uint32) (err error) {
	return instance.SetProperty("CurrentLinkSpeed", (value))
}

// GetCurrentLinkSpeed gets the value of CurrentLinkSpeed for the instance
func (instance *MSNdis_PciDeviceProperty) GetPropertyCurrentLinkSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentLinkSpeed")
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

// SetCurrentLinkWidth sets the value of CurrentLinkWidth for the instance
func (instance *MSNdis_PciDeviceProperty) SetPropertyCurrentLinkWidth(value uint32) (err error) {
	return instance.SetProperty("CurrentLinkWidth", (value))
}

// GetCurrentLinkWidth gets the value of CurrentLinkWidth for the instance
func (instance *MSNdis_PciDeviceProperty) GetPropertyCurrentLinkWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentLinkWidth")
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

// SetCurrentPayloadSize sets the value of CurrentPayloadSize for the instance
func (instance *MSNdis_PciDeviceProperty) SetPropertyCurrentPayloadSize(value uint32) (err error) {
	return instance.SetProperty("CurrentPayloadSize", (value))
}

// GetCurrentPayloadSize gets the value of CurrentPayloadSize for the instance
func (instance *MSNdis_PciDeviceProperty) GetPropertyCurrentPayloadSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentPayloadSize")
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

// SetCurrentSpeedAndMode sets the value of CurrentSpeedAndMode for the instance
func (instance *MSNdis_PciDeviceProperty) SetPropertyCurrentSpeedAndMode(value uint32) (err error) {
	return instance.SetProperty("CurrentSpeedAndMode", (value))
}

// GetCurrentSpeedAndMode gets the value of CurrentSpeedAndMode for the instance
func (instance *MSNdis_PciDeviceProperty) GetPropertyCurrentSpeedAndMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentSpeedAndMode")
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

// SetDeviceType sets the value of DeviceType for the instance
func (instance *MSNdis_PciDeviceProperty) SetPropertyDeviceType(value uint32) (err error) {
	return instance.SetProperty("DeviceType", (value))
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *MSNdis_PciDeviceProperty) GetPropertyDeviceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceType")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_PciDeviceProperty) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_PciDeviceProperty) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetMaxLinkSpeed sets the value of MaxLinkSpeed for the instance
func (instance *MSNdis_PciDeviceProperty) SetPropertyMaxLinkSpeed(value uint32) (err error) {
	return instance.SetProperty("MaxLinkSpeed", (value))
}

// GetMaxLinkSpeed gets the value of MaxLinkSpeed for the instance
func (instance *MSNdis_PciDeviceProperty) GetPropertyMaxLinkSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLinkSpeed")
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

// SetMaxLinkWidth sets the value of MaxLinkWidth for the instance
func (instance *MSNdis_PciDeviceProperty) SetPropertyMaxLinkWidth(value uint32) (err error) {
	return instance.SetProperty("MaxLinkWidth", (value))
}

// GetMaxLinkWidth gets the value of MaxLinkWidth for the instance
func (instance *MSNdis_PciDeviceProperty) GetPropertyMaxLinkWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLinkWidth")
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

// SetMaxPayloadSize sets the value of MaxPayloadSize for the instance
func (instance *MSNdis_PciDeviceProperty) SetPropertyMaxPayloadSize(value uint32) (err error) {
	return instance.SetProperty("MaxPayloadSize", (value))
}

// GetMaxPayloadSize gets the value of MaxPayloadSize for the instance
func (instance *MSNdis_PciDeviceProperty) GetPropertyMaxPayloadSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxPayloadSize")
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

// SetMaxReadRequestSize sets the value of MaxReadRequestSize for the instance
func (instance *MSNdis_PciDeviceProperty) SetPropertyMaxReadRequestSize(value uint32) (err error) {
	return instance.SetProperty("MaxReadRequestSize", (value))
}

// GetMaxReadRequestSize gets the value of MaxReadRequestSize for the instance
func (instance *MSNdis_PciDeviceProperty) GetPropertyMaxReadRequestSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReadRequestSize")
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
