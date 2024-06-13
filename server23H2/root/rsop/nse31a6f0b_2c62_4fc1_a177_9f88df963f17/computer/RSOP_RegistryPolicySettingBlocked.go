// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NSE31A6F0B_2C62_4FC1_A177_9F88DF963F17.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_RegistryPolicySettingBlocked struct
type RSOP_RegistryPolicySettingBlocked struct {
	*RSOP_PolicySettingBlocked

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

func NewRSOP_RegistryPolicySettingBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_RegistryPolicySettingBlocked, err error) {
	tmp, err := NewRSOP_PolicySettingBlockedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_RegistryPolicySettingBlocked{
		RSOP_PolicySettingBlocked: tmp,
	}
	return
}

func NewRSOP_RegistryPolicySettingBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_RegistryPolicySettingBlocked, err error) {
	tmp, err := NewRSOP_PolicySettingBlockedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_RegistryPolicySettingBlocked{
		RSOP_PolicySettingBlocked: tmp,
	}
	return
}

// Setcommand sets the value of command for the instance
func (instance *RSOP_RegistryPolicySettingBlocked) SetPropertycommand(value string) (err error) {
	return instance.SetProperty("command", (value))
}

// Getcommand gets the value of command for the instance
func (instance *RSOP_RegistryPolicySettingBlocked) GetPropertycommand() (value string, err error) {
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
func (instance *RSOP_RegistryPolicySettingBlocked) SetPropertydeleted(value bool) (err error) {
	return instance.SetProperty("deleted", (value))
}

// Getdeleted gets the value of deleted for the instance
func (instance *RSOP_RegistryPolicySettingBlocked) GetPropertydeleted() (value bool, err error) {
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
func (instance *RSOP_RegistryPolicySettingBlocked) SetPropertyregistryKey(value string) (err error) {
	return instance.SetProperty("registryKey", (value))
}

// GetregistryKey gets the value of registryKey for the instance
func (instance *RSOP_RegistryPolicySettingBlocked) GetPropertyregistryKey() (value string, err error) {
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
func (instance *RSOP_RegistryPolicySettingBlocked) SetPropertyvalue(value []uint8) (err error) {
	return instance.SetProperty("value", (value))
}

// Getvalue gets the value of value for the instance
func (instance *RSOP_RegistryPolicySettingBlocked) GetPropertyvalue() (value []uint8, err error) {
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
func (instance *RSOP_RegistryPolicySettingBlocked) SetPropertyvalueName(value string) (err error) {
	return instance.SetProperty("valueName", (value))
}

// GetvalueName gets the value of valueName for the instance
func (instance *RSOP_RegistryPolicySettingBlocked) GetPropertyvalueName() (value string, err error) {
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
func (instance *RSOP_RegistryPolicySettingBlocked) SetPropertyvalueType(value uint32) (err error) {
	return instance.SetProperty("valueType", (value))
}

// GetvalueType gets the value of valueType for the instance
func (instance *RSOP_RegistryPolicySettingBlocked) GetPropertyvalueType() (value uint32, err error) {
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
