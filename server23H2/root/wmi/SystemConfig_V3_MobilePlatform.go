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

// SystemConfig_V3_MobilePlatform struct
type SystemConfig_V3_MobilePlatform struct {
	*SystemConfig_V3

	//
	BspVersion string

	//
	DeviceManufacturer string

	//
	DeviceManufacturerDisplayName string

	//
	DeviceModel string

	//
	DeviceModelDisplayName string

	//
	HardwareVersion string

	//
	MobileOperator string

	//
	MobileOperatorDisplayName string

	//
	OemSoftwareVersion string

	//
	RadioHardwareVersion string

	//
	RadioSoftwareVersion string

	//
	SocVersion string
}

func NewSystemConfig_V3_MobilePlatformEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V3_MobilePlatform, err error) {
	tmp, err := NewSystemConfig_V3Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V3_MobilePlatform{
		SystemConfig_V3: tmp,
	}
	return
}

func NewSystemConfig_V3_MobilePlatformEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V3_MobilePlatform, err error) {
	tmp, err := NewSystemConfig_V3Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V3_MobilePlatform{
		SystemConfig_V3: tmp,
	}
	return
}

// SetBspVersion sets the value of BspVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyBspVersion(value string) (err error) {
	return instance.SetProperty("BspVersion", (value))
}

// GetBspVersion gets the value of BspVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyBspVersion() (value string, err error) {
	retValue, err := instance.GetProperty("BspVersion")
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

// SetDeviceManufacturer sets the value of DeviceManufacturer for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyDeviceManufacturer(value string) (err error) {
	return instance.SetProperty("DeviceManufacturer", (value))
}

// GetDeviceManufacturer gets the value of DeviceManufacturer for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyDeviceManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceManufacturer")
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

// SetDeviceManufacturerDisplayName sets the value of DeviceManufacturerDisplayName for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyDeviceManufacturerDisplayName(value string) (err error) {
	return instance.SetProperty("DeviceManufacturerDisplayName", (value))
}

// GetDeviceManufacturerDisplayName gets the value of DeviceManufacturerDisplayName for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyDeviceManufacturerDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceManufacturerDisplayName")
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

// SetDeviceModel sets the value of DeviceModel for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyDeviceModel(value string) (err error) {
	return instance.SetProperty("DeviceModel", (value))
}

// GetDeviceModel gets the value of DeviceModel for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyDeviceModel() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceModel")
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

// SetDeviceModelDisplayName sets the value of DeviceModelDisplayName for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyDeviceModelDisplayName(value string) (err error) {
	return instance.SetProperty("DeviceModelDisplayName", (value))
}

// GetDeviceModelDisplayName gets the value of DeviceModelDisplayName for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyDeviceModelDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceModelDisplayName")
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

// SetHardwareVersion sets the value of HardwareVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyHardwareVersion(value string) (err error) {
	return instance.SetProperty("HardwareVersion", (value))
}

// GetHardwareVersion gets the value of HardwareVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyHardwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("HardwareVersion")
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

// SetMobileOperator sets the value of MobileOperator for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyMobileOperator(value string) (err error) {
	return instance.SetProperty("MobileOperator", (value))
}

// GetMobileOperator gets the value of MobileOperator for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyMobileOperator() (value string, err error) {
	retValue, err := instance.GetProperty("MobileOperator")
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

// SetMobileOperatorDisplayName sets the value of MobileOperatorDisplayName for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyMobileOperatorDisplayName(value string) (err error) {
	return instance.SetProperty("MobileOperatorDisplayName", (value))
}

// GetMobileOperatorDisplayName gets the value of MobileOperatorDisplayName for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyMobileOperatorDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("MobileOperatorDisplayName")
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

// SetOemSoftwareVersion sets the value of OemSoftwareVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyOemSoftwareVersion(value string) (err error) {
	return instance.SetProperty("OemSoftwareVersion", (value))
}

// GetOemSoftwareVersion gets the value of OemSoftwareVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyOemSoftwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("OemSoftwareVersion")
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

// SetRadioHardwareVersion sets the value of RadioHardwareVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyRadioHardwareVersion(value string) (err error) {
	return instance.SetProperty("RadioHardwareVersion", (value))
}

// GetRadioHardwareVersion gets the value of RadioHardwareVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyRadioHardwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("RadioHardwareVersion")
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

// SetRadioSoftwareVersion sets the value of RadioSoftwareVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertyRadioSoftwareVersion(value string) (err error) {
	return instance.SetProperty("RadioSoftwareVersion", (value))
}

// GetRadioSoftwareVersion gets the value of RadioSoftwareVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertyRadioSoftwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("RadioSoftwareVersion")
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

// SetSocVersion sets the value of SocVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) SetPropertySocVersion(value string) (err error) {
	return instance.SetProperty("SocVersion", (value))
}

// GetSocVersion gets the value of SocVersion for the instance
func (instance *SystemConfig_V3_MobilePlatform) GetPropertySocVersion() (value string, err error) {
	retValue, err := instance.GetProperty("SocVersion")
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
