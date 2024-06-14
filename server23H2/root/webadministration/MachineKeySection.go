// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MachineKeySection struct
type MachineKeySection struct {
	*ConfigurationSection

	//
	ApplicationName string

	//
	CompatibilityMode int32

	//
	DataProtectorType string

	//
	Decryption string

	//
	DecryptionKey string

	//
	Validation int32

	//
	ValidationKey string
}

func NewMachineKeySectionEx1(instance *cim.WmiInstance) (newInstance *MachineKeySection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MachineKeySection{
		ConfigurationSection: tmp,
	}
	return
}

func NewMachineKeySectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MachineKeySection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MachineKeySection{
		ConfigurationSection: tmp,
	}
	return
}

// SetApplicationName sets the value of ApplicationName for the instance
func (instance *MachineKeySection) SetPropertyApplicationName(value string) (err error) {
	return instance.SetProperty("ApplicationName", (value))
}

// GetApplicationName gets the value of ApplicationName for the instance
func (instance *MachineKeySection) GetPropertyApplicationName() (value string, err error) {
	retValue, err := instance.GetProperty("ApplicationName")
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

// SetCompatibilityMode sets the value of CompatibilityMode for the instance
func (instance *MachineKeySection) SetPropertyCompatibilityMode(value int32) (err error) {
	return instance.SetProperty("CompatibilityMode", (value))
}

// GetCompatibilityMode gets the value of CompatibilityMode for the instance
func (instance *MachineKeySection) GetPropertyCompatibilityMode() (value int32, err error) {
	retValue, err := instance.GetProperty("CompatibilityMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetDataProtectorType sets the value of DataProtectorType for the instance
func (instance *MachineKeySection) SetPropertyDataProtectorType(value string) (err error) {
	return instance.SetProperty("DataProtectorType", (value))
}

// GetDataProtectorType gets the value of DataProtectorType for the instance
func (instance *MachineKeySection) GetPropertyDataProtectorType() (value string, err error) {
	retValue, err := instance.GetProperty("DataProtectorType")
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

// SetDecryption sets the value of Decryption for the instance
func (instance *MachineKeySection) SetPropertyDecryption(value string) (err error) {
	return instance.SetProperty("Decryption", (value))
}

// GetDecryption gets the value of Decryption for the instance
func (instance *MachineKeySection) GetPropertyDecryption() (value string, err error) {
	retValue, err := instance.GetProperty("Decryption")
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

// SetDecryptionKey sets the value of DecryptionKey for the instance
func (instance *MachineKeySection) SetPropertyDecryptionKey(value string) (err error) {
	return instance.SetProperty("DecryptionKey", (value))
}

// GetDecryptionKey gets the value of DecryptionKey for the instance
func (instance *MachineKeySection) GetPropertyDecryptionKey() (value string, err error) {
	retValue, err := instance.GetProperty("DecryptionKey")
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

// SetValidation sets the value of Validation for the instance
func (instance *MachineKeySection) SetPropertyValidation(value int32) (err error) {
	return instance.SetProperty("Validation", (value))
}

// GetValidation gets the value of Validation for the instance
func (instance *MachineKeySection) GetPropertyValidation() (value int32, err error) {
	retValue, err := instance.GetProperty("Validation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetValidationKey sets the value of ValidationKey for the instance
func (instance *MachineKeySection) SetPropertyValidationKey(value string) (err error) {
	return instance.SetProperty("ValidationKey", (value))
}

// GetValidationKey gets the value of ValidationKey for the instance
func (instance *MachineKeySection) GetPropertyValidationKey() (value string, err error) {
	retValue, err := instance.GetProperty("ValidationKey")
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
