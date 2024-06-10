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

// SystemDataSection struct
type SystemDataSection struct {
	*ConfigurationSectionWithCollection

	//
	DbProviderFactories DbProviderFactorySettings
}

func NewSystemDataSectionEx1(instance *cim.WmiInstance) (newInstance *SystemDataSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemDataSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewSystemDataSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemDataSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemDataSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetDbProviderFactories sets the value of DbProviderFactories for the instance
func (instance *SystemDataSection) SetPropertyDbProviderFactories(value DbProviderFactorySettings) (err error) {
	return instance.SetProperty("DbProviderFactories", (value))
}

// GetDbProviderFactories gets the value of DbProviderFactories for the instance
func (instance *SystemDataSection) GetPropertyDbProviderFactories() (value DbProviderFactorySettings, err error) {
	retValue, err := instance.GetProperty("DbProviderFactories")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DbProviderFactorySettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DbProviderFactorySettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DbProviderFactorySettings(valuetmp)

	return
}
