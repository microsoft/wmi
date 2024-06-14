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

// ProvidersSettings struct
type ProvidersSettings struct {
	*EmbeddedObject

	//
	Providers []NameTypeElement
}

func NewProvidersSettingsEx1(instance *cim.WmiInstance) (newInstance *ProvidersSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProvidersSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewProvidersSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProvidersSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProvidersSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetProviders sets the value of Providers for the instance
func (instance *ProvidersSettings) SetPropertyProviders(value []NameTypeElement) (err error) {
	return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *ProvidersSettings) GetPropertyProviders() (value []NameTypeElement, err error) {
	retValue, err := instance.GetProperty("Providers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(NameTypeElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " NameTypeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, NameTypeElement(valuetmp))
	}

	return
}
