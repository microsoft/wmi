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

// TraceUrlAreaSettings struct
type TraceUrlAreaSettings struct {
	*EmbeddedObject

	//
	TraceAreas []TraceAreaElement
}

func NewTraceUrlAreaSettingsEx1(instance *cim.WmiInstance) (newInstance *TraceUrlAreaSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TraceUrlAreaSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewTraceUrlAreaSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TraceUrlAreaSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TraceUrlAreaSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetTraceAreas sets the value of TraceAreas for the instance
func (instance *TraceUrlAreaSettings) SetPropertyTraceAreas(value []TraceAreaElement) (err error) {
	return instance.SetProperty("TraceAreas", (value))
}

// GetTraceAreas gets the value of TraceAreas for the instance
func (instance *TraceUrlAreaSettings) GetPropertyTraceAreas() (value []TraceAreaElement, err error) {
	retValue, err := instance.GetProperty("TraceAreas")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TraceAreaElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TraceAreaElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TraceAreaElement(valuetmp))
	}

	return
}
