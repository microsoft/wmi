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

// CachingSection struct
type CachingSection struct {
	*ConfigurationSectionWithCollection

	//
	Enabled bool

	//
	EnableKernelCache bool

	//
	MaxCacheSize uint32

	//
	MaxResponseSize uint32

	//
	Profiles CachingProfileSettings
}

func NewCachingSectionEx1(instance *cim.WmiInstance) (newInstance *CachingSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CachingSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewCachingSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CachingSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CachingSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *CachingSection) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *CachingSection) GetPropertyEnabled() (value bool, err error) {
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

// SetEnableKernelCache sets the value of EnableKernelCache for the instance
func (instance *CachingSection) SetPropertyEnableKernelCache(value bool) (err error) {
	return instance.SetProperty("EnableKernelCache", (value))
}

// GetEnableKernelCache gets the value of EnableKernelCache for the instance
func (instance *CachingSection) GetPropertyEnableKernelCache() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableKernelCache")
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

// SetMaxCacheSize sets the value of MaxCacheSize for the instance
func (instance *CachingSection) SetPropertyMaxCacheSize(value uint32) (err error) {
	return instance.SetProperty("MaxCacheSize", (value))
}

// GetMaxCacheSize gets the value of MaxCacheSize for the instance
func (instance *CachingSection) GetPropertyMaxCacheSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCacheSize")
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

// SetMaxResponseSize sets the value of MaxResponseSize for the instance
func (instance *CachingSection) SetPropertyMaxResponseSize(value uint32) (err error) {
	return instance.SetProperty("MaxResponseSize", (value))
}

// GetMaxResponseSize gets the value of MaxResponseSize for the instance
func (instance *CachingSection) GetPropertyMaxResponseSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxResponseSize")
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

// SetProfiles sets the value of Profiles for the instance
func (instance *CachingSection) SetPropertyProfiles(value CachingProfileSettings) (err error) {
	return instance.SetProperty("Profiles", (value))
}

// GetProfiles gets the value of Profiles for the instance
func (instance *CachingSection) GetPropertyProfiles() (value CachingProfileSettings, err error) {
	retValue, err := instance.GetProperty("Profiles")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CachingProfileSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CachingProfileSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CachingProfileSettings(valuetmp)

	return
}
