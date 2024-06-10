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

// ListenerSettings struct
type ListenerSettings struct {
	*EmbeddedObject

	//
	SharedListeners []ListenerElement
}

func NewListenerSettingsEx1(instance *cim.WmiInstance) (newInstance *ListenerSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ListenerSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewListenerSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ListenerSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ListenerSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetSharedListeners sets the value of SharedListeners for the instance
func (instance *ListenerSettings) SetPropertySharedListeners(value []ListenerElement) (err error) {
	return instance.SetProperty("SharedListeners", (value))
}

// GetSharedListeners gets the value of SharedListeners for the instance
func (instance *ListenerSettings) GetPropertySharedListeners() (value []ListenerElement, err error) {
	retValue, err := instance.GetProperty("SharedListeners")
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
