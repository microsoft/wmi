// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSE8125520_1F5F_48D5_9C53_BB40C321ED3D.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_RegistryKey struct
type RSOP_RegistryKey struct {
	*RSOP_SecuritySettings

	//
	Mode RegistryKey_Mode

	//
	Path string

	//
	SDDLString string
}

func NewRSOP_RegistryKeyEx1(instance *cim.WmiInstance) (newInstance *RSOP_RegistryKey, err error) {
	tmp, err := NewRSOP_SecuritySettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_RegistryKey{
		RSOP_SecuritySettings: tmp,
	}
	return
}

func NewRSOP_RegistryKeyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_RegistryKey, err error) {
	tmp, err := NewRSOP_SecuritySettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_RegistryKey{
		RSOP_SecuritySettings: tmp,
	}
	return
}

// SetMode sets the value of Mode for the instance
func (instance *RSOP_RegistryKey) SetPropertyMode(value RegistryKey_Mode) (err error) {
	return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *RSOP_RegistryKey) GetPropertyMode() (value RegistryKey_Mode, err error) {
	retValue, err := instance.GetProperty("Mode")
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

	value = RegistryKey_Mode(valuetmp)

	return
}

// SetPath sets the value of Path for the instance
func (instance *RSOP_RegistryKey) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *RSOP_RegistryKey) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetSDDLString sets the value of SDDLString for the instance
func (instance *RSOP_RegistryKey) SetPropertySDDLString(value string) (err error) {
	return instance.SetProperty("SDDLString", (value))
}

// GetSDDLString gets the value of SDDLString for the instance
func (instance *RSOP_RegistryKey) GetPropertySDDLString() (value string, err error) {
	retValue, err := instance.GetProperty("SDDLString")
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
