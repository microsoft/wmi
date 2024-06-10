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

// RoleManagerSection struct
type RoleManagerSection struct {
	*ConfigurationSectionWithCollection

	//
	CacheRolesInCookie bool

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
	CreatePersistentCookie bool

	//
	DefaultProvider string

	//
	Domain string

	//
	Enabled bool

	//
	MaxCachedResults int32

	//
	Providers ProvidersSettings
}

func NewRoleManagerSectionEx1(instance *cim.WmiInstance) (newInstance *RoleManagerSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RoleManagerSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewRoleManagerSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RoleManagerSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RoleManagerSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetCacheRolesInCookie sets the value of CacheRolesInCookie for the instance
func (instance *RoleManagerSection) SetPropertyCacheRolesInCookie(value bool) (err error) {
	return instance.SetProperty("CacheRolesInCookie", (value))
}

// GetCacheRolesInCookie gets the value of CacheRolesInCookie for the instance
func (instance *RoleManagerSection) GetPropertyCacheRolesInCookie() (value bool, err error) {
	retValue, err := instance.GetProperty("CacheRolesInCookie")
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

// SetCookieName sets the value of CookieName for the instance
func (instance *RoleManagerSection) SetPropertyCookieName(value string) (err error) {
	return instance.SetProperty("CookieName", (value))
}

// GetCookieName gets the value of CookieName for the instance
func (instance *RoleManagerSection) GetPropertyCookieName() (value string, err error) {
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
func (instance *RoleManagerSection) SetPropertyCookiePath(value string) (err error) {
	return instance.SetProperty("CookiePath", (value))
}

// GetCookiePath gets the value of CookiePath for the instance
func (instance *RoleManagerSection) GetPropertyCookiePath() (value string, err error) {
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
func (instance *RoleManagerSection) SetPropertyCookieProtection(value int32) (err error) {
	return instance.SetProperty("CookieProtection", (value))
}

// GetCookieProtection gets the value of CookieProtection for the instance
func (instance *RoleManagerSection) GetPropertyCookieProtection() (value int32, err error) {
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
func (instance *RoleManagerSection) SetPropertyCookieRequireSSL(value bool) (err error) {
	return instance.SetProperty("CookieRequireSSL", (value))
}

// GetCookieRequireSSL gets the value of CookieRequireSSL for the instance
func (instance *RoleManagerSection) GetPropertyCookieRequireSSL() (value bool, err error) {
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
func (instance *RoleManagerSection) SetPropertyCookieSlidingExpiration(value bool) (err error) {
	return instance.SetProperty("CookieSlidingExpiration", (value))
}

// GetCookieSlidingExpiration gets the value of CookieSlidingExpiration for the instance
func (instance *RoleManagerSection) GetPropertyCookieSlidingExpiration() (value bool, err error) {
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
func (instance *RoleManagerSection) SetPropertyCookieTimeout(value string) (err error) {
	return instance.SetProperty("CookieTimeout", (value))
}

// GetCookieTimeout gets the value of CookieTimeout for the instance
func (instance *RoleManagerSection) GetPropertyCookieTimeout() (value string, err error) {
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

// SetCreatePersistentCookie sets the value of CreatePersistentCookie for the instance
func (instance *RoleManagerSection) SetPropertyCreatePersistentCookie(value bool) (err error) {
	return instance.SetProperty("CreatePersistentCookie", (value))
}

// GetCreatePersistentCookie gets the value of CreatePersistentCookie for the instance
func (instance *RoleManagerSection) GetPropertyCreatePersistentCookie() (value bool, err error) {
	retValue, err := instance.GetProperty("CreatePersistentCookie")
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

// SetDefaultProvider sets the value of DefaultProvider for the instance
func (instance *RoleManagerSection) SetPropertyDefaultProvider(value string) (err error) {
	return instance.SetProperty("DefaultProvider", (value))
}

// GetDefaultProvider gets the value of DefaultProvider for the instance
func (instance *RoleManagerSection) GetPropertyDefaultProvider() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultProvider")
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
func (instance *RoleManagerSection) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", (value))
}

// GetDomain gets the value of Domain for the instance
func (instance *RoleManagerSection) GetPropertyDomain() (value string, err error) {
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
func (instance *RoleManagerSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *RoleManagerSection) GetPropertyEnabled() (value bool, err error) {
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

// SetMaxCachedResults sets the value of MaxCachedResults for the instance
func (instance *RoleManagerSection) SetPropertyMaxCachedResults(value int32) (err error) {
	return instance.SetProperty("MaxCachedResults", (value))
}

// GetMaxCachedResults gets the value of MaxCachedResults for the instance
func (instance *RoleManagerSection) GetPropertyMaxCachedResults() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxCachedResults")
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

// SetProviders sets the value of Providers for the instance
func (instance *RoleManagerSection) SetPropertyProviders(value ProvidersSettings) (err error) {
	return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *RoleManagerSection) GetPropertyProviders() (value ProvidersSettings, err error) {
	retValue, err := instance.GetProperty("Providers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ProvidersSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ProvidersSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ProvidersSettings(valuetmp)

	return
}
