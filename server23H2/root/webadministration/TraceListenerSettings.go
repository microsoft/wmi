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

// TraceListenerSettings struct
type TraceListenerSettings struct {
	*EmbeddedObject

	//
	Listeners []ListenerElement
}

func NewTraceListenerSettingsEx1(instance *cim.WmiInstance) (newInstance *TraceListenerSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TraceListenerSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewTraceListenerSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TraceListenerSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TraceListenerSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetListeners sets the value of Listeners for the instance
func (instance *TraceListenerSettings) SetPropertyListeners(value []ListenerElement) (err error) {
	return instance.SetProperty("Listeners", (value))
}

// GetListeners gets the value of Listeners for the instance
func (instance *TraceListenerSettings) GetPropertyListeners() (value []ListenerElement, err error) {
	retValue, err := instance.GetProperty("Listeners")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ListenerElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ListenerElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ListenerElement(valuetmp))
	}

	return
}