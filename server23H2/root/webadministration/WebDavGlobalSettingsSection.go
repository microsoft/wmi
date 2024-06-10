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

// WebDavGlobalSettingsSection struct
type WebDavGlobalSettingsSection struct {
	*ConfigurationSectionWithCollection

	//
	LockStores WebDavLockStoresSettings

	//
	PropertyStores WebDavPropertyStoresSettings
}

func NewWebDavGlobalSettingsSectionEx1(instance *cim.WmiInstance) (newInstance *WebDavGlobalSettingsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebDavGlobalSettingsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewWebDavGlobalSettingsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebDavGlobalSettingsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebDavGlobalSettingsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetLockStores sets the value of LockStores for the instance
func (instance *WebDavGlobalSettingsSection) SetPropertyLockStores(value WebDavLockStoresSettings) (err error) {
	return instance.SetProperty("LockStores", (value))
}

// GetLockStores gets the value of LockStores for the instance
func (instance *WebDavGlobalSettingsSection) GetPropertyLockStores() (value WebDavLockStoresSettings, err error) {
	retValue, err := instance.GetProperty("LockStores")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WebDavLockStoresSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WebDavLockStoresSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WebDavLockStoresSettings(valuetmp)

	return
}

// SetPropertyStores sets the value of PropertyStores for the instance
func (instance *WebDavGlobalSettingsSection) SetPropertyPropertyStores(value WebDavPropertyStoresSettings) (err error) {
	return instance.SetProperty("PropertyStores", (value))
}

// GetPropertyStores gets the value of PropertyStores for the instance
func (instance *WebDavGlobalSettingsSection) GetPropertyPropertyStores() (value WebDavPropertyStoresSettings, err error) {
	retValue, err := instance.GetProperty("PropertyStores")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(WebDavPropertyStoresSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " WebDavPropertyStoresSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WebDavPropertyStoresSettings(valuetmp)

	return
}
