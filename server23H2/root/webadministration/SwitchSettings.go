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

// SwitchSettings struct
type SwitchSettings struct {
	*EmbeddedObject

	//
	Name string

	//
	Value string
}

func NewSwitchSettingsEx1(instance *cim.WmiInstance) (newInstance *SwitchSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SwitchSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewSwitchSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SwitchSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SwitchSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *SwitchSettings) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *SwitchSettings) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetValue sets the value of Value for the instance
func (instance *SwitchSettings) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", (value))
}

// GetValue gets the value of Value for the instance
func (instance *SwitchSettings) GetPropertyValue() (value string, err error) {
	retValue, err := instance.GetProperty("Value")
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