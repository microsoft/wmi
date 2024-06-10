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

// SiteMapSection struct
type SiteMapSection struct {
	*ConfigurationSectionWithCollection

	//
	DefaultProvider string

	//
	Enabled bool

	//
	Providers ProviderSettings
}

func NewSiteMapSectionEx1(instance *cim.WmiInstance) (newInstance *SiteMapSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SiteMapSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewSiteMapSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SiteMapSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SiteMapSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetDefaultProvider sets the value of DefaultProvider for the instance
func (instance *SiteMapSection) SetPropertyDefaultProvider(value string) (err error) {
	return instance.SetProperty("DefaultProvider", (value))
}

// GetDefaultProvider gets the value of DefaultProvider for the instance
func (instance *SiteMapSection) GetPropertyDefaultProvider() (value string, err error) {
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

// SetEnabled sets the value of Enabled for the instance
func (instance *SiteMapSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *SiteMapSection) GetPropertyEnabled() (value bool, err error) {
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

// SetProviders sets the value of Providers for the instance
func (instance *SiteMapSection) SetPropertyProviders(value ProviderSettings) (err error) {
	return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *SiteMapSection) GetPropertyProviders() (value ProviderSettings, err error) {
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
