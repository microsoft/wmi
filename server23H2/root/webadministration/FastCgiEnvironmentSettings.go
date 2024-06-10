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

// FastCgiEnvironmentSettings struct
type FastCgiEnvironmentSettings struct {
	*EmbeddedObject

	//
	EnvironmentVariables []FastCgiEnvironmentElement
}

func NewFastCgiEnvironmentSettingsEx1(instance *cim.WmiInstance) (newInstance *FastCgiEnvironmentSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FastCgiEnvironmentSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFastCgiEnvironmentSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FastCgiEnvironmentSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FastCgiEnvironmentSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetEnvironmentVariables sets the value of EnvironmentVariables for the instance
func (instance *FastCgiEnvironmentSettings) SetPropertyEnvironmentVariables(value []FastCgiEnvironmentElement) (err error) {
	return instance.SetProperty("EnvironmentVariables", (value))
}

// GetEnvironmentVariables gets the value of EnvironmentVariables for the instance
func (instance *FastCgiEnvironmentSettings) GetPropertyEnvironmentVariables() (value []FastCgiEnvironmentElement, err error) {
	retValue, err := instance.GetProperty("EnvironmentVariables")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FastCgiEnvironmentElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FastCgiEnvironmentElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FastCgiEnvironmentElement(valuetmp))
	}

	return
}
