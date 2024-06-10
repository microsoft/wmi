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

// HealthMonitoringSection struct
type HealthMonitoringSection struct {
	*ConfigurationSectionWithCollection

	//
	BufferModes BufferModeSettings

	//
	Enabled bool

	//
	EventMappings EventMappingSettings

	//
	HeartbeatInterval string

	//
	Profiles ProfileSettings

	//
	Providers ProviderSettings

	//
	Rules RuleSettings
}

func NewHealthMonitoringSectionEx1(instance *cim.WmiInstance) (newInstance *HealthMonitoringSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HealthMonitoringSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewHealthMonitoringSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HealthMonitoringSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HealthMonitoringSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetBufferModes sets the value of BufferModes for the instance
func (instance *HealthMonitoringSection) SetPropertyBufferModes(value BufferModeSettings) (err error) {
	return instance.SetProperty("BufferModes", (value))
}

// GetBufferModes gets the value of BufferModes for the instance
func (instance *HealthMonitoringSection) GetPropertyBufferModes() (value BufferModeSettings, err error) {
	retValue, err := instance.GetProperty("BufferModes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BufferModeSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BufferModeSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BufferModeSettings(valuetmp)

	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *HealthMonitoringSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *HealthMonitoringSection) GetPropertyEnabled() (value bool, err error) {
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

// SetEventMappings sets the value of EventMappings for the instance
func (instance *HealthMonitoringSection) SetPropertyEventMappings(value EventMappingSettings) (err error) {
	return instance.SetProperty("EventMappings", (value))
}

// GetEventMappings gets the value of EventMappings for the instance
func (instance *HealthMonitoringSection) GetPropertyEventMappings() (value EventMappingSettings, err error) {
	retValue, err := instance.GetProperty("EventMappings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(EventMappingSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " EventMappingSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = EventMappingSettings(valuetmp)

	return
}

// SetHeartbeatInterval sets the value of HeartbeatInterval for the instance
func (instance *HealthMonitoringSection) SetPropertyHeartbeatInterval(value string) (err error) {
	return instance.SetProperty("HeartbeatInterval", (value))
}

// GetHeartbeatInterval gets the value of HeartbeatInterval for the instance
func (instance *HealthMonitoringSection) GetPropertyHeartbeatInterval() (value string, err error) {
	retValue, err := instance.GetProperty("HeartbeatInterval")
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

// SetProfiles sets the value of Profiles for the instance
func (instance *HealthMonitoringSection) SetPropertyProfiles(value ProfileSettings) (err error) {
	return instance.SetProperty("Profiles", (value))
}

// GetProfiles gets the value of Profiles for the instance
func (instance *HealthMonitoringSection) GetPropertyProfiles() (value ProfileSettings, err error) {
	retValue, err := instance.GetProperty("Profiles")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ProfileSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ProfileSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ProfileSettings(valuetmp)

	return
}

// SetProviders sets the value of Providers for the instance
func (instance *HealthMonitoringSection) SetPropertyProviders(value ProviderSettings) (err error) {
	return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *HealthMonitoringSection) GetPropertyProviders() (value ProviderSettings, err error) {
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

// SetRules sets the value of Rules for the instance
func (instance *HealthMonitoringSection) SetPropertyRules(value RuleSettings) (err error) {
	return instance.SetProperty("Rules", (value))
}

// GetRules gets the value of Rules for the instance
func (instance *HealthMonitoringSection) GetPropertyRules() (value RuleSettings, err error) {
	retValue, err := instance.GetProperty("Rules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RuleSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RuleSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RuleSettings(valuetmp)

	return
}
