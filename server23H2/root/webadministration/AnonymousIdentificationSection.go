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

// AnonymousIdentificationSection struct
type AnonymousIdentificationSection struct {
	*ConfigurationSection

	//
	Cookieless int32

	//
	CookieName string

	//
	CookiePath string

	//
	CookieProtection int32

	//
	CookieRequireSSL bool

	//
	CookieSlidingExpiration bool

	//
	CookieTimeout string

	//
	Domain string

	//
	Enabled bool
}

func NewAnonymousIdentificationSectionEx1(instance *cim.WmiInstance) (newInstance *AnonymousIdentificationSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AnonymousIdentificationSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewAnonymousIdentificationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AnonymousIdentificationSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AnonymousIdentificationSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetCookieless sets the value of Cookieless for the instance
func (instance *AnonymousIdentificationSection) SetPropertyCookieless(value int32) (err error) {
	return instance.SetProperty("Cookieless", (value))
}

// GetCookieless gets the value of Cookieless for the instance
func (instance *AnonymousIdentificationSection) GetPropertyCookieless() (value int32, err error) {
	retValue, err := instance.GetProperty("Cookieless")
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

// SetCookieName sets the value of CookieName for the instance
func (instance *AnonymousIdentificationSection) SetPropertyCookieName(value string) (err error) {
	return instance.SetProperty("CookieName", (value))
}

// GetCookieName gets the value of CookieName for the instance
func (instance *AnonymousIdentificationSection) GetPropertyCookieName() (value string, err error) {
	retValue, err := instance.GetProperty("CookieName")
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

// SetCookiePath sets the value of CookiePath for the instance
func (instance *AnonymousIdentificationSection) SetPropertyCookiePath(value string) (err error) {
	return instance.SetProperty("CookiePath", (value))
}

// GetCookiePath gets the value of CookiePath for the instance
func (instance *AnonymousIdentificationSection) GetPropertyCookiePath() (value string, err error) {
	retValue, err := instance.GetProperty("CookiePath")
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

// SetCookieProtection sets the value of CookieProtection for the instance
func (instance *AnonymousIdentificationSection) SetPropertyCookieProtection(value int32) (err error) {
	return instance.SetProperty("CookieProtection", (value))
}

// GetCookieProtection gets the value of CookieProtection for the instance
func (instance *AnonymousIdentificationSection) GetPropertyCookieProtection() (value int32, err error) {
	retValue, err := instance.GetProperty("CookieProtection")
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

// SetCookieRequireSSL sets the value of CookieRequireSSL for the instance
func (instance *AnonymousIdentificationSection) SetPropertyCookieRequireSSL(value bool) (err error) {
	return instance.SetProperty("CookieRequireSSL", (value))
}

// GetCookieRequireSSL gets the value of CookieRequireSSL for the instance
func (instance *AnonymousIdentificationSection) GetPropertyCookieRequireSSL() (value bool, err error) {
	retValue, err := instance.GetProperty("CookieRequireSSL")
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

// SetCookieSlidingExpiration sets the value of CookieSlidingExpiration for the instance
func (instance *AnonymousIdentificationSection) SetPropertyCookieSlidingExpiration(value bool) (err error) {
	return instance.SetProperty("CookieSlidingExpiration", (value))
}

// GetCookieSlidingExpiration gets the value of CookieSlidingExpiration for the instance
func (instance *AnonymousIdentificationSection) GetPropertyCookieSlidingExpiration() (value bool, err error) {
	retValue, err := instance.GetProperty("CookieSlidingExpiration")
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

// SetCookieTimeout sets the value of CookieTimeout for the instance
func (instance *AnonymousIdentificationSection) SetPropertyCookieTimeout(value string) (err error) {
	return instance.SetProperty("CookieTimeout", (value))
}

// GetCookieTimeout gets the value of CookieTimeout for the instance
func (instance *AnonymousIdentificationSection) GetPropertyCookieTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("CookieTimeout")
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

// SetDomain sets the value of Domain for the instance
func (instance *AnonymousIdentificationSection) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", (value))
}

// GetDomain gets the value of Domain for the instance
func (instance *AnonymousIdentificationSection) GetPropertyDomain() (value string, err error) {
	retValue, err := instance.GetProperty("Domain")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *AnonymousIdentificationSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *AnonymousIdentificationSection) GetPropertyEnabled() (value bool, err error) {
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
