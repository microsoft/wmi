// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS814BDA2A_0BB7_4B36_B7B7_0A0141C4C544.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_RegistryPolicySetting struct
type RSOP_RegistryPolicySetting struct {
	*RSOP_PolicySetting

	//
	command string

	//
	deleted bool

	//
	registryKey string

	//
	value []uint8

	//
	valueName string

	//
	valueType uint32
}

func NewRSOP_RegistryPolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_RegistryPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_RegistryPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_RegistryPolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_RegistryPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_RegistryPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// Setcommand sets the value of command for the instance
func (instance *RSOP_RegistryPolicySetting) SetPropertycommand(value string) (err error) {
	return instance.SetProperty("command", (value))
}

// Getcommand gets the value of command for the instance
func (instance *RSOP_RegistryPolicySetting) GetPropertycommand() (value string, err error) {
	retValue, err := instance.GetProperty("command")
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

// Setdeleted sets the value of deleted for the instance
func (instance *RSOP_RegistryPolicySetting) SetPropertydeleted(value bool) (err error) {
	return instance.SetProperty("deleted", (value))
}

// Getdeleted gets the value of deleted for the instance
func (instance *RSOP_RegistryPolicySetting) GetPropertydeleted() (value bool, err error) {
	retValue, err := instance.GetProperty("deleted")
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

// SetregistryKey sets the value of registryKey for the instance
func (instance *RSOP_RegistryPolicySetting) SetPropertyregistryKey(value string) (err error) {
	return instance.SetProperty("registryKey", (value))
}

// GetregistryKey gets the value of registryKey for the instance
func (instance *RSOP_RegistryPolicySetting) GetPropertyregistryKey() (value string, err error) {
	retValue, err := instance.GetProperty("registryKey")
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

// Setvalue sets the value of value for the instance
func (instance *RSOP_RegistryPolicySetting) SetPropertyvalue(value []uint8) (err error) {
	return instance.SetProperty("value", (value))
}

// Getvalue gets the value of value for the instance
func (instance *RSOP_RegistryPolicySetting) GetPropertyvalue() (value []uint8, err error) {
	retValue, err := instance.GetProperty("value")
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

// SetvalueName sets the value of valueName for the instance
func (instance *RSOP_RegistryPolicySetting) SetPropertyvalueName(value string) (err error) {
	return instance.SetProperty("valueName", (value))
}

// GetvalueName gets the value of valueName for the instance
func (instance *RSOP_RegistryPolicySetting) GetPropertyvalueName() (value string, err error) {
	retValue, err := instance.GetProperty("valueName")
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

// SetvalueType sets the value of valueType for the instance
func (instance *RSOP_RegistryPolicySetting) SetPropertyvalueType(value uint32) (err error) {
	return instance.SetProperty("valueType", (value))
}

// GetvalueType gets the value of valueType for the instance
func (instance *RSOP_RegistryPolicySetting) GetPropertyvalueType() (value uint32, err error) {
	retValue, err := instance.GetProperty("valueType")
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
