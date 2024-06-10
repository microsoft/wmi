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

// RedirectHeaderSettings struct
type RedirectHeaderSettings struct {
	*EmbeddedObject

	//
	RedirectHeaders []NameValueConfigurationElement
}

func NewRedirectHeaderSettingsEx1(instance *cim.WmiInstance) (newInstance *RedirectHeaderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RedirectHeaderSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewRedirectHeaderSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RedirectHeaderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RedirectHeaderSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetRedirectHeaders sets the value of RedirectHeaders for the instance
func (instance *RedirectHeaderSettings) SetPropertyRedirectHeaders(value []NameValueConfigurationElement) (err error) {
	return instance.SetProperty("RedirectHeaders", (value))
}

// GetRedirectHeaders gets the value of RedirectHeaders for the instance
func (instance *RedirectHeaderSettings) GetPropertyRedirectHeaders() (value []NameValueConfigurationElement, err error) {
	retValue, err := instance.GetProperty("RedirectHeaders")
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
