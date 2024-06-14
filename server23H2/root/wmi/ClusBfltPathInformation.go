// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// ClusBfltPathInformation struct
type ClusBfltPathInformation struct {
	*cim.WmiInstance

	// Attributes.
	Attributes uint32

	// Bus Type.
	BusType uint32

	// Device Guid.
	DeviceGuid string

	// Device Number.
	DeviceNumber uint32

	// Type.
	DeviceType uint32

	// Id.
	Id uint32

	// Type.
	PathType uint32

	// Registration Key.
	RegistrationKey uint64

	// Status.
	Status uint32
}

func NewClusBfltPathInformationEx1(instance *cim.WmiInstance) (newInstance *ClusBfltPathInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ClusBfltPathInformation{
		WmiInstance: tmp,
	}
	return
}

func NewClusBfltPathInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClusBfltPathInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClusBfltPathInformation{
		WmiInstance: tmp,
	}
	return
}

// SetAttributes sets the value of Attributes for the instance
func (instance *ClusBfltPathInformation) SetPropertyAttributes(value uint32) (err error) {
	return instance.SetProperty("Attributes", (value))
}

// GetAttributes gets the value of Attributes for the instance
func (instance *ClusBfltPathInformation) GetPropertyAttributes() (value uint32, err error) {
	retValue, err := instance.GetProperty("Attributes")
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

// SetBusType sets the value of BusType for the instance
func (instance *ClusBfltPathInformation) SetPropertyBusType(value uint32) (err error) {
	return instance.SetProperty("BusType", (value))
}

// GetBusType gets the value of BusType for the instance
func (instance *ClusBfltPathInformation) GetPropertyBusType() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusType")
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

// SetDeviceGuid sets the value of DeviceGuid for the instance
func (instance *ClusBfltPathInformation) SetPropertyDeviceGuid(value string) (err error) {
	return instance.SetProperty("DeviceGuid", (value))
}

// GetDeviceGuid gets the value of DeviceGuid for the instance
func (instance *ClusBfltPathInformation) GetPropertyDeviceGuid() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceGuid")
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

// SetDeviceNumber sets the value of DeviceNumber for the instance
func (instance *ClusBfltPathInformation) SetPropertyDeviceNumber(value uint32) (err error) {
	return instance.SetProperty("DeviceNumber", (value))
}

// GetDeviceNumber gets the value of DeviceNumber for the instance
func (instance *ClusBfltPathInformation) GetPropertyDeviceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceNumber")
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
func (instance *ClusBfltPathInformation) SetPropertyDeviceType(value uint32) (err error) {
	return instance.SetProperty("DeviceType", (value))
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *ClusBfltPathInformation) GetPropertyDeviceType() (value uint32, err error) {
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

// SetId sets the value of Id for the instance
func (instance *ClusBfltPathInformation) SetPropertyId(value uint32) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *ClusBfltPathInformation) GetPropertyId() (value uint32, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetPathType sets the value of PathType for the instance
func (instance *ClusBfltPathInformation) SetPropertyPathType(value uint32) (err error) {
	return instance.SetProperty("PathType", (value))
}

// GetPathType gets the value of PathType for the instance
func (instance *ClusBfltPathInformation) GetPropertyPathType() (value uint32, err error) {
	retValue, err := instance.GetProperty("PathType")
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

// SetRegistrationKey sets the value of RegistrationKey for the instance
func (instance *ClusBfltPathInformation) SetPropertyRegistrationKey(value uint64) (err error) {
	return instance.SetProperty("RegistrationKey", (value))
}

// GetRegistrationKey gets the value of RegistrationKey for the instance
func (instance *ClusBfltPathInformation) GetPropertyRegistrationKey() (value uint64, err error) {
	retValue, err := instance.GetProperty("RegistrationKey")
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

// SetStatus sets the value of Status for the instance
func (instance *ClusBfltPathInformation) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *ClusBfltPathInformation) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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
