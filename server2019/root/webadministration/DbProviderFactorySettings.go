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

// DbProviderFactorySettings struct
type DbProviderFactorySettings struct {
	*EmbeddedObject

	//
	DbProviderFactories []ProviderFactoryElement
}

func NewDbProviderFactorySettingsEx1(instance *cim.WmiInstance) (newInstance *DbProviderFactorySettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DbProviderFactorySettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewDbProviderFactorySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DbProviderFactorySettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DbProviderFactorySettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetDbProviderFactories sets the value of DbProviderFactories for the instance
func (instance *DbProviderFactorySettings) SetPropertyDbProviderFactories(value []ProviderFactoryElement) (err error) {
	return instance.SetProperty("DbProviderFactories", (value))
}

// GetDbProviderFactories gets the value of DbProviderFactories for the instance
func (instance *DbProviderFactorySettings) GetPropertyDbProviderFactories() (value []ProviderFactoryElement, err error) {
	retValue, err := instance.GetProperty("DbProviderFactories")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ProviderFactoryElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ProviderFactoryElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ProviderFactoryElement(valuetmp))
	}

	return
}
