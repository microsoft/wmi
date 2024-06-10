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

// AppSettingsSection struct
type AppSettingsSection struct {
	*ConfigurationSectionWithCollection

	//
	AppSettings []KeyValueElement

	//
	File string
}

func NewAppSettingsSectionEx1(instance *cim.WmiInstance) (newInstance *AppSettingsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AppSettingsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewAppSettingsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AppSettingsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AppSettingsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAppSettings sets the value of AppSettings for the instance
func (instance *AppSettingsSection) SetPropertyAppSettings(value []KeyValueElement) (err error) {
	return instance.SetProperty("AppSettings", (value))
}

// GetAppSettings gets the value of AppSettings for the instance
func (instance *AppSettingsSection) GetPropertyAppSettings() (value []KeyValueElement, err error) {
	retValue, err := instance.GetProperty("AppSettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(KeyValueElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " KeyValueElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, KeyValueElement(valuetmp))
	}

	return
}

// SetFile sets the value of File for the instance
func (instance *AppSettingsSection) SetPropertyFile(value string) (err error) {
	return instance.SetProperty("File", (value))
}

// GetFile gets the value of File for the instance
func (instance *AppSettingsSection) GetPropertyFile() (value string, err error) {
	retValue, err := instance.GetProperty("File")
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
