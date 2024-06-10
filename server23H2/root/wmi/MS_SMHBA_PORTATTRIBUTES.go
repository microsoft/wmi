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

// MS_SMHBA_PORTATTRIBUTES struct
type MS_SMHBA_PORTATTRIBUTES struct {
	*cim.WmiInstance

	//
	OSDeviceName string

	//
	PortSpecificAttributes []uint8

	//
	PortSpecificAttributesSize uint32

	//
	PortState uint32

	//
	PortType uint32

	//
	Reserved uint64
}

func NewMS_SMHBA_PORTATTRIBUTESEx1(instance *cim.WmiInstance) (newInstance *MS_SMHBA_PORTATTRIBUTES, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_PORTATTRIBUTES{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SMHBA_PORTATTRIBUTESEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SMHBA_PORTATTRIBUTES, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_PORTATTRIBUTES{
		WmiInstance: tmp,
	}
	return
}

// SetOSDeviceName sets the value of OSDeviceName for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) SetPropertyOSDeviceName(value string) (err error) {
	return instance.SetProperty("OSDeviceName", (value))
}

// GetOSDeviceName gets the value of OSDeviceName for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) GetPropertyOSDeviceName() (value string, err error) {
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

// SetPortSpecificAttributes sets the value of PortSpecificAttributes for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) SetPropertyPortSpecificAttributes(value []uint8) (err error) {
	return instance.SetProperty("PortSpecificAttributes", (value))
}

// GetPortSpecificAttributes gets the value of PortSpecificAttributes for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) GetPropertyPortSpecificAttributes() (value []uint8, err error) {
	retValue, err := instance.GetProperty("PortSpecificAttributes")
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

// SetPortSpecificAttributesSize sets the value of PortSpecificAttributesSize for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) SetPropertyPortSpecificAttributesSize(value uint32) (err error) {
	return instance.SetProperty("PortSpecificAttributesSize", (value))
}

// GetPortSpecificAttributesSize gets the value of PortSpecificAttributesSize for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) GetPropertyPortSpecificAttributesSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortSpecificAttributesSize")
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

// SetPortState sets the value of PortState for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) SetPropertyPortState(value uint32) (err error) {
	return instance.SetProperty("PortState", (value))
}

// GetPortState gets the value of PortState for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) GetPropertyPortState() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortState")
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

// SetPortType sets the value of PortType for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) SetPropertyPortType(value uint32) (err error) {
	return instance.SetProperty("PortType", (value))
}

// GetPortType gets the value of PortType for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) GetPropertyPortType() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortType")
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

// SetReserved sets the value of Reserved for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) SetPropertyReserved(value uint64) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *MS_SMHBA_PORTATTRIBUTES) GetPropertyReserved() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reserved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
