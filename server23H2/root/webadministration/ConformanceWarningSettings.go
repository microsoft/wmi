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

// ConformanceWarningSettings struct
type ConformanceWarningSettings struct {
	*EmbeddedObject

	//
	ConformanceWarnings []ConformanceWarning
}

func NewConformanceWarningSettingsEx1(instance *cim.WmiInstance) (newInstance *ConformanceWarningSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ConformanceWarningSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewConformanceWarningSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ConformanceWarningSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ConformanceWarningSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetConformanceWarnings sets the value of ConformanceWarnings for the instance
func (instance *ConformanceWarningSettings) SetPropertyConformanceWarnings(value []ConformanceWarning) (err error) {
	return instance.SetProperty("ConformanceWarnings", (value))
}

// GetConformanceWarnings gets the value of ConformanceWarnings for the instance
func (instance *ConformanceWarningSettings) GetPropertyConformanceWarnings() (value []ConformanceWarning, err error) {
	retValue, err := instance.GetProperty("ConformanceWarnings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ConformanceWarning)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ConformanceWarning is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ConformanceWarning(valuetmp))
	}

	return
}
