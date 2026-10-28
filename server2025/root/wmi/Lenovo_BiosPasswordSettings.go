// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Lenovo_BiosPasswordSettings struct
type Lenovo_BiosPasswordSettings struct {
	*Lenovo_BIOSElement

	//
	Active bool

	//
	HardDiskMasterPasswordState uint32

	//
	HardDiskUserPasswordState uint32

	//
	InstanceName string

	//
	MaxLength uint32

	//
	MinLength uint32

	//
	NvmeMasterPasswordState uint32

	//
	NvmeUserPasswordState uint32

	//
	PasswordMode uint32

	//
	PasswordState uint32

	//
	SupportedEncodings uint32

	//
	SupportedKeyboard uint32
}

func NewLenovo_BiosPasswordSettingsEx1(instance *cim.WmiInstance) (newInstance *Lenovo_BiosPasswordSettings, err error) {
	tmp, err := NewLenovo_BIOSElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Lenovo_BiosPasswordSettings{
		Lenovo_BIOSElement: tmp,
	}
	return
}

func NewLenovo_BiosPasswordSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Lenovo_BiosPasswordSettings, err error) {
	tmp, err := NewLenovo_BIOSElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Lenovo_BiosPasswordSettings{
		Lenovo_BIOSElement: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertyActive() (value bool, err error) {
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

// SetHardDiskMasterPasswordState sets the value of HardDiskMasterPasswordState for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertyHardDiskMasterPasswordState(value uint32) (err error) {
	return instance.SetProperty("HardDiskMasterPasswordState", (value))
}

// GetHardDiskMasterPasswordState gets the value of HardDiskMasterPasswordState for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertyHardDiskMasterPasswordState() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardDiskMasterPasswordState")
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

// SetHardDiskUserPasswordState sets the value of HardDiskUserPasswordState for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertyHardDiskUserPasswordState(value uint32) (err error) {
	return instance.SetProperty("HardDiskUserPasswordState", (value))
}

// GetHardDiskUserPasswordState gets the value of HardDiskUserPasswordState for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertyHardDiskUserPasswordState() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardDiskUserPasswordState")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertyInstanceName() (value string, err error) {
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

// SetMaxLength sets the value of MaxLength for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertyMaxLength(value uint32) (err error) {
	return instance.SetProperty("MaxLength", (value))
}

// GetMaxLength gets the value of MaxLength for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertyMaxLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLength")
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

// SetMinLength sets the value of MinLength for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertyMinLength(value uint32) (err error) {
	return instance.SetProperty("MinLength", (value))
}

// GetMinLength gets the value of MinLength for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertyMinLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinLength")
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

// SetNvmeMasterPasswordState sets the value of NvmeMasterPasswordState for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertyNvmeMasterPasswordState(value uint32) (err error) {
	return instance.SetProperty("NvmeMasterPasswordState", (value))
}

// GetNvmeMasterPasswordState gets the value of NvmeMasterPasswordState for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertyNvmeMasterPasswordState() (value uint32, err error) {
	retValue, err := instance.GetProperty("NvmeMasterPasswordState")
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

// SetNvmeUserPasswordState sets the value of NvmeUserPasswordState for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertyNvmeUserPasswordState(value uint32) (err error) {
	return instance.SetProperty("NvmeUserPasswordState", (value))
}

// GetNvmeUserPasswordState gets the value of NvmeUserPasswordState for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertyNvmeUserPasswordState() (value uint32, err error) {
	retValue, err := instance.GetProperty("NvmeUserPasswordState")
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

// SetPasswordMode sets the value of PasswordMode for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertyPasswordMode(value uint32) (err error) {
	return instance.SetProperty("PasswordMode", (value))
}

// GetPasswordMode gets the value of PasswordMode for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertyPasswordMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("PasswordMode")
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

// SetPasswordState sets the value of PasswordState for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertyPasswordState(value uint32) (err error) {
	return instance.SetProperty("PasswordState", (value))
}

// GetPasswordState gets the value of PasswordState for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertyPasswordState() (value uint32, err error) {
	retValue, err := instance.GetProperty("PasswordState")
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

// SetSupportedEncodings sets the value of SupportedEncodings for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertySupportedEncodings(value uint32) (err error) {
	return instance.SetProperty("SupportedEncodings", (value))
}

// GetSupportedEncodings gets the value of SupportedEncodings for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertySupportedEncodings() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedEncodings")
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

// SetSupportedKeyboard sets the value of SupportedKeyboard for the instance
func (instance *Lenovo_BiosPasswordSettings) SetPropertySupportedKeyboard(value uint32) (err error) {
	return instance.SetProperty("SupportedKeyboard", (value))
}

// GetSupportedKeyboard gets the value of SupportedKeyboard for the instance
func (instance *Lenovo_BiosPasswordSettings) GetPropertySupportedKeyboard() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedKeyboard")
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
