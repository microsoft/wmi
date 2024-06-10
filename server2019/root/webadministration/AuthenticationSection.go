// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AuthenticationSection struct
type AuthenticationSection struct {
	*ConfigurationSectionWithCollection

	//
	Forms FormsAuthenticationConfiguration

	//
	Mode int32

	//
	Passport PassportAuthentication
}

func NewAuthenticationSectionEx1(instance *cim.WmiInstance) (newInstance *AuthenticationSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AuthenticationSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewAuthenticationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AuthenticationSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AuthenticationSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetForms sets the value of Forms for the instance
func (instance *AuthenticationSection) SetPropertyForms(value FormsAuthenticationConfiguration) (err error) {
	return instance.SetProperty("Forms", (value))
}

// GetForms gets the value of Forms for the instance
func (instance *AuthenticationSection) GetPropertyForms() (value FormsAuthenticationConfiguration, err error) {
	retValue, err := instance.GetProperty("Forms")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FormsAuthenticationConfiguration)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FormsAuthenticationConfiguration is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FormsAuthenticationConfiguration(valuetmp)

	return
}

// SetMode sets the value of Mode for the instance
func (instance *AuthenticationSection) SetPropertyMode(value int32) (err error) {
	return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *AuthenticationSection) GetPropertyMode() (value int32, err error) {
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

	value = int32(valuetmp)

	return
}

// SetPassport sets the value of Passport for the instance
func (instance *AuthenticationSection) SetPropertyPassport(value PassportAuthentication) (err error) {
	return instance.SetProperty("Passport", (value))
}

// GetPassport gets the value of Passport for the instance
func (instance *AuthenticationSection) GetPropertyPassport() (value PassportAuthentication, err error) {
	retValue, err := instance.GetProperty("Passport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(PassportAuthentication)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " PassportAuthentication is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PassportAuthentication(valuetmp)

	return
}
