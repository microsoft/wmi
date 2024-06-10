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

// CustomHeaderSettings struct
type CustomHeaderSettings struct {
	*EmbeddedObject

	//
	CustomHeaders []NameValueConfigurationElement
}

func NewCustomHeaderSettingsEx1(instance *cim.WmiInstance) (newInstance *CustomHeaderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CustomHeaderSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewCustomHeaderSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CustomHeaderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CustomHeaderSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetCustomHeaders sets the value of CustomHeaders for the instance
func (instance *CustomHeaderSettings) SetPropertyCustomHeaders(value []NameValueConfigurationElement) (err error) {
	return instance.SetProperty("CustomHeaders", (value))
}

// GetCustomHeaders gets the value of CustomHeaders for the instance
func (instance *CustomHeaderSettings) GetPropertyCustomHeaders() (value []NameValueConfigurationElement, err error) {
	retValue, err := instance.GetProperty("CustomHeaders")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(NameValueConfigurationElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " NameValueConfigurationElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, NameValueConfigurationElement(valuetmp))
	}

	return
}
