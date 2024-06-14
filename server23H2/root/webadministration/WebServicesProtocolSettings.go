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

// WebServicesProtocolSettings struct
type WebServicesProtocolSettings struct {
	*EmbeddedObject

	//
	Protocols []WebServicesProtocolElement
}

func NewWebServicesProtocolSettingsEx1(instance *cim.WmiInstance) (newInstance *WebServicesProtocolSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebServicesProtocolSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewWebServicesProtocolSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebServicesProtocolSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebServicesProtocolSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetProtocols sets the value of Protocols for the instance
func (instance *WebServicesProtocolSettings) SetPropertyProtocols(value []WebServicesProtocolElement) (err error) {
	return instance.SetProperty("Protocols", (value))
}

// GetProtocols gets the value of Protocols for the instance
func (instance *WebServicesProtocolSettings) GetPropertyProtocols() (value []WebServicesProtocolElement, err error) {
	retValue, err := instance.GetProperty("Protocols")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(WebServicesProtocolElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " WebServicesProtocolElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, WebServicesProtocolElement(valuetmp))
	}

	return
}
