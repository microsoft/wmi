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

// AppliesToSettings struct
type AppliesToSettings struct {
	*EmbeddedObject

	//
	AppliesTo []AppliesToElement
}

func NewAppliesToSettingsEx1(instance *cim.WmiInstance) (newInstance *AppliesToSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AppliesToSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewAppliesToSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AppliesToSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AppliesToSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAppliesTo sets the value of AppliesTo for the instance
func (instance *AppliesToSettings) SetPropertyAppliesTo(value []AppliesToElement) (err error) {
	return instance.SetProperty("AppliesTo", (value))
}

// GetAppliesTo gets the value of AppliesTo for the instance
func (instance *AppliesToSettings) GetPropertyAppliesTo() (value []AppliesToElement, err error) {
	retValue, err := instance.GetProperty("AppliesTo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(AppliesToElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " AppliesToElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, AppliesToElement(valuetmp))
	}

	return
}
