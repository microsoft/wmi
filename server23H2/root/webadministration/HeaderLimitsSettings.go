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

// HeaderLimitsSettings struct
type HeaderLimitsSettings struct {
	*EmbeddedObject

	//
	HeaderLimits []HeaderLimitsElement
}

func NewHeaderLimitsSettingsEx1(instance *cim.WmiInstance) (newInstance *HeaderLimitsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HeaderLimitsSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewHeaderLimitsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeaderLimitsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeaderLimitsSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetHeaderLimits sets the value of HeaderLimits for the instance
func (instance *HeaderLimitsSettings) SetPropertyHeaderLimits(value []HeaderLimitsElement) (err error) {
	return instance.SetProperty("HeaderLimits", (value))
}

// GetHeaderLimits gets the value of HeaderLimits for the instance
func (instance *HeaderLimitsSettings) GetPropertyHeaderLimits() (value []HeaderLimitsElement, err error) {
	retValue, err := instance.GetProperty("HeaderLimits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(HeaderLimitsElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " HeaderLimitsElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, HeaderLimitsElement(valuetmp))
	}

	return
}
