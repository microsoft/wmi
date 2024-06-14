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

// SystemConfig_V2_MobilePlatform struct
type SystemConfig_V2_MobilePlatform struct {
	*SystemConfig_V2

	//
	BootLoaderVersion string

	//
	FirmwareRevision string

	//
	FriendlyName string

	//
	HardwareRevision string

	//
	HardwareVariant string

	//
	Manufacturer string

	//
	ManufacturerDisplayName string

	//
	ManufacturerModelName string

	//
	MobileOperatorDisplayName string

	//
	MobileOperatorName string

	//
	ModelName string

	//
	RadioHardwareRevision string

	//
	RadioSoftwareRevision string

	//
	ROMVersion string

	//
	SOCVersion string
}

func NewSystemConfig_V2_MobilePlatformEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_MobilePlatform, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_MobilePlatform{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_MobilePlatformEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_MobilePlatform, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_MobilePlatform{
		SystemConfig_V2: tmp,
	}
	return
}

// SetBootLoaderVersion sets the value of BootLoaderVersion for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyBootLoaderVersion(value string) (err error) {
	return instance.SetProperty("BootLoaderVersion", (value))
}

// GetBootLoaderVersion gets the value of BootLoaderVersion for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyBootLoaderVersion() (value string, err error) {
	retValue, err := instance.GetProperty("BootLoaderVersion")
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

// SetFirmwareRevision sets the value of FirmwareRevision for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyFirmwareRevision(value string) (err error) {
	return instance.SetProperty("FirmwareRevision", (value))
}

// GetFirmwareRevision gets the value of FirmwareRevision for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyFirmwareRevision() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareRevision")
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

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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

// SetHardwareRevision sets the value of HardwareRevision for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyHardwareRevision(value string) (err error) {
	return instance.SetProperty("HardwareRevision", (value))
}

// GetHardwareRevision gets the value of HardwareRevision for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyHardwareRevision() (value string, err error) {
	retValue, err := instance.GetProperty("HardwareRevision")
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

// SetHardwareVariant sets the value of HardwareVariant for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyHardwareVariant(value string) (err error) {
	return instance.SetProperty("HardwareVariant", (value))
}

// GetHardwareVariant gets the value of HardwareVariant for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyHardwareVariant() (value string, err error) {
	retValue, err := instance.GetProperty("HardwareVariant")
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

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
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

// SetManufacturerDisplayName sets the value of ManufacturerDisplayName for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyManufacturerDisplayName(value string) (err error) {
	return instance.SetProperty("ManufacturerDisplayName", (value))
}

// GetManufacturerDisplayName gets the value of ManufacturerDisplayName for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyManufacturerDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("ManufacturerDisplayName")
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

// SetManufacturerModelName sets the value of ManufacturerModelName for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyManufacturerModelName(value string) (err error) {
	return instance.SetProperty("ManufacturerModelName", (value))
}

// GetManufacturerModelName gets the value of ManufacturerModelName for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyManufacturerModelName() (value string, err error) {
	retValue, err := instance.GetProperty("ManufacturerModelName")
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
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyMobileOperatorDisplayName(value string) (err error) {
	return instance.SetProperty("MobileOperatorDisplayName", (value))
}

// GetMobileOperatorDisplayName gets the value of MobileOperatorDisplayName for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyMobileOperatorDisplayName() (value string, err error) {
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

// SetMobileOperatorName sets the value of MobileOperatorName for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyMobileOperatorName(value string) (err error) {
	return instance.SetProperty("MobileOperatorName", (value))
}

// GetMobileOperatorName gets the value of MobileOperatorName for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyMobileOperatorName() (value string, err error) {
	retValue, err := instance.GetProperty("MobileOperatorName")
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

// SetModelName sets the value of ModelName for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyModelName(value string) (err error) {
	return instance.SetProperty("ModelName", (value))
}

// GetModelName gets the value of ModelName for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyModelName() (value string, err error) {
	retValue, err := instance.GetProperty("ModelName")
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

// SetRadioHardwareRevision sets the value of RadioHardwareRevision for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyRadioHardwareRevision(value string) (err error) {
	return instance.SetProperty("RadioHardwareRevision", (value))
}

// GetRadioHardwareRevision gets the value of RadioHardwareRevision for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyRadioHardwareRevision() (value string, err error) {
	retValue, err := instance.GetProperty("RadioHardwareRevision")
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

// SetRadioSoftwareRevision sets the value of RadioSoftwareRevision for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyRadioSoftwareRevision(value string) (err error) {
	return instance.SetProperty("RadioSoftwareRevision", (value))
}

// GetRadioSoftwareRevision gets the value of RadioSoftwareRevision for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyRadioSoftwareRevision() (value string, err error) {
	retValue, err := instance.GetProperty("RadioSoftwareRevision")
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

// SetROMVersion sets the value of ROMVersion for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertyROMVersion(value string) (err error) {
	return instance.SetProperty("ROMVersion", (value))
}

// GetROMVersion gets the value of ROMVersion for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertyROMVersion() (value string, err error) {
	retValue, err := instance.GetProperty("ROMVersion")
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

// SetSOCVersion sets the value of SOCVersion for the instance
func (instance *SystemConfig_V2_MobilePlatform) SetPropertySOCVersion(value string) (err error) {
	return instance.SetProperty("SOCVersion", (value))
}

// GetSOCVersion gets the value of SOCVersion for the instance
func (instance *SystemConfig_V2_MobilePlatform) GetPropertySOCVersion() (value string, err error) {
	retValue, err := instance.GetProperty("SOCVersion")
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
