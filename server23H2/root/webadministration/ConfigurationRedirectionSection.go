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

// ConfigurationRedirectionSection struct
type ConfigurationRedirectionSection struct {
	*ConfigurationSection

	//
	Enabled bool

	//
	EnableUncPolling bool

	//
	Password string

	//
	PollingPeriod string

	//
	RedirectionPath string

	//
	UserName string
}

func NewConfigurationRedirectionSectionEx1(instance *cim.WmiInstance) (newInstance *ConfigurationRedirectionSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ConfigurationRedirectionSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewConfigurationRedirectionSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ConfigurationRedirectionSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ConfigurationRedirectionSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *ConfigurationRedirectionSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *ConfigurationRedirectionSection) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetEnableUncPolling sets the value of EnableUncPolling for the instance
func (instance *ConfigurationRedirectionSection) SetPropertyEnableUncPolling(value bool) (err error) {
	return instance.SetProperty("EnableUncPolling", (value))
}

// GetEnableUncPolling gets the value of EnableUncPolling for the instance
func (instance *ConfigurationRedirectionSection) GetPropertyEnableUncPolling() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableUncPolling")
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

// SetPassword sets the value of Password for the instance
func (instance *ConfigurationRedirectionSection) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *ConfigurationRedirectionSection) GetPropertyPassword() (value string, err error) {
	retValue, err := instance.GetProperty("Password")
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

// SetPollingPeriod sets the value of PollingPeriod for the instance
func (instance *ConfigurationRedirectionSection) SetPropertyPollingPeriod(value string) (err error) {
	return instance.SetProperty("PollingPeriod", (value))
}

// GetPollingPeriod gets the value of PollingPeriod for the instance
func (instance *ConfigurationRedirectionSection) GetPropertyPollingPeriod() (value string, err error) {
	retValue, err := instance.GetProperty("PollingPeriod")
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

// SetRedirectionPath sets the value of RedirectionPath for the instance
func (instance *ConfigurationRedirectionSection) SetPropertyRedirectionPath(value string) (err error) {
	return instance.SetProperty("RedirectionPath", (value))
}

// GetRedirectionPath gets the value of RedirectionPath for the instance
func (instance *ConfigurationRedirectionSection) GetPropertyRedirectionPath() (value string, err error) {
	retValue, err := instance.GetProperty("RedirectionPath")
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

// SetUserName sets the value of UserName for the instance
func (instance *ConfigurationRedirectionSection) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *ConfigurationRedirectionSection) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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
