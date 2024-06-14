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

// AlwaysAllowedQueryStringSettings struct
type AlwaysAllowedQueryStringSettings struct {
	*EmbeddedObject

	//
	AlwaysAllowedQueryStrings []QueryStringElement
}

func NewAlwaysAllowedQueryStringSettingsEx1(instance *cim.WmiInstance) (newInstance *AlwaysAllowedQueryStringSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AlwaysAllowedQueryStringSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewAlwaysAllowedQueryStringSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AlwaysAllowedQueryStringSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AlwaysAllowedQueryStringSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAlwaysAllowedQueryStrings sets the value of AlwaysAllowedQueryStrings for the instance
func (instance *AlwaysAllowedQueryStringSettings) SetPropertyAlwaysAllowedQueryStrings(value []QueryStringElement) (err error) {
	return instance.SetProperty("AlwaysAllowedQueryStrings", (value))
}

// GetAlwaysAllowedQueryStrings gets the value of AlwaysAllowedQueryStrings for the instance
func (instance *AlwaysAllowedQueryStringSettings) GetPropertyAlwaysAllowedQueryStrings() (value []QueryStringElement, err error) {
	retValue, err := instance.GetProperty("AlwaysAllowedQueryStrings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(QueryStringElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " QueryStringElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, QueryStringElement(valuetmp))
	}

	return
}
