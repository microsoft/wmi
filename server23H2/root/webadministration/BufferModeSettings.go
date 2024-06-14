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

// BufferModeSettings struct
type BufferModeSettings struct {
	*EmbeddedObject

	//
	BufferModes []BufferModeElement
}

func NewBufferModeSettingsEx1(instance *cim.WmiInstance) (newInstance *BufferModeSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BufferModeSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewBufferModeSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BufferModeSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BufferModeSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetBufferModes sets the value of BufferModes for the instance
func (instance *BufferModeSettings) SetPropertyBufferModes(value []BufferModeElement) (err error) {
	return instance.SetProperty("BufferModes", (value))
}

// GetBufferModes gets the value of BufferModes for the instance
func (instance *BufferModeSettings) GetPropertyBufferModes() (value []BufferModeElement, err error) {
	retValue, err := instance.GetProperty("BufferModes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(BufferModeElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " BufferModeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, BufferModeElement(valuetmp))
	}

	return
}
