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

// SessionStateSection struct
type SessionStateSection struct {
	*ConfigurationSectionWithCollection

	//
	AllowCustomSqlDatabase bool

	//
	CompressionEnabled bool

	//
	Cookieless int32

	//
	CookieName string

	//
	CustomProvider string

	//
	Mode int32

	//
	PartitionResolverType string

	//
	Providers ProviderSettings

	//
	RegenerateExpiredSessionId bool

	//
	SessionIDManagerType string

	//
	SqlCommandTimeout string

	//
	SqlConnectionRetryInterval string

	//
	SqlConnectionString string

	//
	StateConnectionString string

	//
	StateNetworkTimeout string

	//
	Timeout string

	//
	UseHostingIdentity bool
}

func NewSessionStateSectionEx1(instance *cim.WmiInstance) (newInstance *SessionStateSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SessionStateSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewSessionStateSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SessionStateSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SessionStateSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAllowCustomSqlDatabase sets the value of AllowCustomSqlDatabase for the instance
func (instance *SessionStateSection) SetPropertyAllowCustomSqlDatabase(value bool) (err error) {
	return instance.SetProperty("AllowCustomSqlDatabase", (value))
}

// GetAllowCustomSqlDatabase gets the value of AllowCustomSqlDatabase for the instance
func (instance *SessionStateSection) GetPropertyAllowCustomSqlDatabase() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowCustomSqlDatabase")
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

// SetCompressionEnabled sets the value of CompressionEnabled for the instance
func (instance *SessionStateSection) SetPropertyCompressionEnabled(value bool) (err error) {
	return instance.SetProperty("CompressionEnabled", (value))
}

// GetCompressionEnabled gets the value of CompressionEnabled for the instance
func (instance *SessionStateSection) GetPropertyCompressionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("CompressionEnabled")
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

// SetCookieless sets the value of Cookieless for the instance
func (instance *SessionStateSection) SetPropertyCookieless(value int32) (err error) {
	return instance.SetProperty("Cookieless", (value))
}

// GetCookieless gets the value of Cookieless for the instance
func (instance *SessionStateSection) GetPropertyCookieless() (value int32, err error) {
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
func (instance *SessionStateSection) SetPropertyCookieName(value string) (err error) {
	return instance.SetProperty("CookieName", (value))
}

// GetCookieName gets the value of CookieName for the instance
func (instance *SessionStateSection) GetPropertyCookieName() (value string, err error) {
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

// SetCustomProvider sets the value of CustomProvider for the instance
func (instance *SessionStateSection) SetPropertyCustomProvider(value string) (err error) {
	return instance.SetProperty("CustomProvider", (value))
}

// GetCustomProvider gets the value of CustomProvider for the instance
func (instance *SessionStateSection) GetPropertyCustomProvider() (value string, err error) {
	retValue, err := instance.GetProperty("CustomProvider")
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

// SetMode sets the value of Mode for the instance
func (instance *SessionStateSection) SetPropertyMode(value int32) (err error) {
	return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *SessionStateSection) GetPropertyMode() (value int32, err error) {
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

// SetPartitionResolverType sets the value of PartitionResolverType for the instance
func (instance *SessionStateSection) SetPropertyPartitionResolverType(value string) (err error) {
	return instance.SetProperty("PartitionResolverType", (value))
}

// GetPartitionResolverType gets the value of PartitionResolverType for the instance
func (instance *SessionStateSection) GetPropertyPartitionResolverType() (value string, err error) {
	retValue, err := instance.GetProperty("PartitionResolverType")
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

// SetProviders sets the value of Providers for the instance
func (instance *SessionStateSection) SetPropertyProviders(value ProviderSettings) (err error) {
	return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *SessionStateSection) GetPropertyProviders() (value ProviderSettings, err error) {
	retValue, err := instance.GetProperty("Providers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ProviderSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ProviderSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ProviderSettings(valuetmp)

	return
}

// SetRegenerateExpiredSessionId sets the value of RegenerateExpiredSessionId for the instance
func (instance *SessionStateSection) SetPropertyRegenerateExpiredSessionId(value bool) (err error) {
	return instance.SetProperty("RegenerateExpiredSessionId", (value))
}

// GetRegenerateExpiredSessionId gets the value of RegenerateExpiredSessionId for the instance
func (instance *SessionStateSection) GetPropertyRegenerateExpiredSessionId() (value bool, err error) {
	retValue, err := instance.GetProperty("RegenerateExpiredSessionId")
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

// SetSessionIDManagerType sets the value of SessionIDManagerType for the instance
func (instance *SessionStateSection) SetPropertySessionIDManagerType(value string) (err error) {
	return instance.SetProperty("SessionIDManagerType", (value))
}

// GetSessionIDManagerType gets the value of SessionIDManagerType for the instance
func (instance *SessionStateSection) GetPropertySessionIDManagerType() (value string, err error) {
	retValue, err := instance.GetProperty("SessionIDManagerType")
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

// SetSqlCommandTimeout sets the value of SqlCommandTimeout for the instance
func (instance *SessionStateSection) SetPropertySqlCommandTimeout(value string) (err error) {
	return instance.SetProperty("SqlCommandTimeout", (value))
}

// GetSqlCommandTimeout gets the value of SqlCommandTimeout for the instance
func (instance *SessionStateSection) GetPropertySqlCommandTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("SqlCommandTimeout")
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

// SetSqlConnectionRetryInterval sets the value of SqlConnectionRetryInterval for the instance
func (instance *SessionStateSection) SetPropertySqlConnectionRetryInterval(value string) (err error) {
	return instance.SetProperty("SqlConnectionRetryInterval", (value))
}

// GetSqlConnectionRetryInterval gets the value of SqlConnectionRetryInterval for the instance
func (instance *SessionStateSection) GetPropertySqlConnectionRetryInterval() (value string, err error) {
	retValue, err := instance.GetProperty("SqlConnectionRetryInterval")
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

// SetSqlConnectionString sets the value of SqlConnectionString for the instance
func (instance *SessionStateSection) SetPropertySqlConnectionString(value string) (err error) {
	return instance.SetProperty("SqlConnectionString", (value))
}

// GetSqlConnectionString gets the value of SqlConnectionString for the instance
func (instance *SessionStateSection) GetPropertySqlConnectionString() (value string, err error) {
	retValue, err := instance.GetProperty("SqlConnectionString")
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

// SetStateConnectionString sets the value of StateConnectionString for the instance
func (instance *SessionStateSection) SetPropertyStateConnectionString(value string) (err error) {
	return instance.SetProperty("StateConnectionString", (value))
}

// GetStateConnectionString gets the value of StateConnectionString for the instance
func (instance *SessionStateSection) GetPropertyStateConnectionString() (value string, err error) {
	retValue, err := instance.GetProperty("StateConnectionString")
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

// SetStateNetworkTimeout sets the value of StateNetworkTimeout for the instance
func (instance *SessionStateSection) SetPropertyStateNetworkTimeout(value string) (err error) {
	return instance.SetProperty("StateNetworkTimeout", (value))
}

// GetStateNetworkTimeout gets the value of StateNetworkTimeout for the instance
func (instance *SessionStateSection) GetPropertyStateNetworkTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("StateNetworkTimeout")
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

// SetTimeout sets the value of Timeout for the instance
func (instance *SessionStateSection) SetPropertyTimeout(value string) (err error) {
	return instance.SetProperty("Timeout", (value))
}

// GetTimeout gets the value of Timeout for the instance
func (instance *SessionStateSection) GetPropertyTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("Timeout")
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

// SetUseHostingIdentity sets the value of UseHostingIdentity for the instance
func (instance *SessionStateSection) SetPropertyUseHostingIdentity(value bool) (err error) {
	return instance.SetProperty("UseHostingIdentity", (value))
}

// GetUseHostingIdentity gets the value of UseHostingIdentity for the instance
func (instance *SessionStateSection) GetPropertyUseHostingIdentity() (value bool, err error) {
	retValue, err := instance.GetProperty("UseHostingIdentity")
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
