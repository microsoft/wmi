// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Uev
//////////////////////////////////////////////
package uev

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ConfigurationItem struct
type ConfigurationItem struct {
	*cim.WmiInstance

	// Is setting valid
	IsValid bool

	// Setting name
	SettingName string

	// Setting source
	SettingSource string

	// Setting value
	SettingValue string
}

func NewConfigurationItemEx1(instance *cim.WmiInstance) (newInstance *ConfigurationItem, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ConfigurationItem{
		WmiInstance: tmp,
	}
	return
}

func NewConfigurationItemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ConfigurationItem, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ConfigurationItem{
		WmiInstance: tmp,
	}
	return
}

// SetIsValid sets the value of IsValid for the instance
func (instance *ConfigurationItem) SetPropertyIsValid(value bool) (err error) {
	return instance.SetProperty("IsValid", (value))
}

// GetIsValid gets the value of IsValid for the instance
func (instance *ConfigurationItem) GetPropertyIsValid() (value bool, err error) {
	retValue, err := instance.GetProperty("IsValid")
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

// SetSettingName sets the value of SettingName for the instance
func (instance *ConfigurationItem) SetPropertySettingName(value string) (err error) {
	return instance.SetProperty("SettingName", (value))
}

// GetSettingName gets the value of SettingName for the instance
func (instance *ConfigurationItem) GetPropertySettingName() (value string, err error) {
	retValue, err := instance.GetProperty("SettingName")
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

// SetSettingSource sets the value of SettingSource for the instance
func (instance *ConfigurationItem) SetPropertySettingSource(value string) (err error) {
	return instance.SetProperty("SettingSource", (value))
}

// GetSettingSource gets the value of SettingSource for the instance
func (instance *ConfigurationItem) GetPropertySettingSource() (value string, err error) {
	retValue, err := instance.GetProperty("SettingSource")
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

// SetSettingValue sets the value of SettingValue for the instance
func (instance *ConfigurationItem) SetPropertySettingValue(value string) (err error) {
	return instance.SetProperty("SettingValue", (value))
}

// GetSettingValue gets the value of SettingValue for the instance
func (instance *ConfigurationItem) GetPropertySettingValue() (value string, err error) {
	retValue, err := instance.GetProperty("SettingValue")
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
