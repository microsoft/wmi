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

// DenyStringsSettings struct
type DenyStringsSettings struct {
	*EmbeddedObject

	//
	DenyStrings []DenyStringElement
}

func NewDenyStringsSettingsEx1(instance *cim.WmiInstance) (newInstance *DenyStringsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DenyStringsSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewDenyStringsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DenyStringsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DenyStringsSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetDenyStrings sets the value of DenyStrings for the instance
func (instance *DenyStringsSettings) SetPropertyDenyStrings(value []DenyStringElement) (err error) {
	return instance.SetProperty("DenyStrings", (value))
}

// GetDenyStrings gets the value of DenyStrings for the instance
func (instance *DenyStringsSettings) GetPropertyDenyStrings() (value []DenyStringElement, err error) {
	retValue, err := instance.GetProperty("DenyStrings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DenyStringElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DenyStringElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DenyStringElement(valuetmp))
	}

	return
}
