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

// ServiceDescriptionFormatSettings struct
type ServiceDescriptionFormatSettings struct {
	*EmbeddedObject

	//
	ServiceDescriptionFormatExtensionTypes []TypeElement
}

func NewServiceDescriptionFormatSettingsEx1(instance *cim.WmiInstance) (newInstance *ServiceDescriptionFormatSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServiceDescriptionFormatSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewServiceDescriptionFormatSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServiceDescriptionFormatSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServiceDescriptionFormatSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetServiceDescriptionFormatExtensionTypes sets the value of ServiceDescriptionFormatExtensionTypes for the instance
func (instance *ServiceDescriptionFormatSettings) SetPropertyServiceDescriptionFormatExtensionTypes(value []TypeElement) (err error) {
	return instance.SetProperty("ServiceDescriptionFormatExtensionTypes", (value))
}

// GetServiceDescriptionFormatExtensionTypes gets the value of ServiceDescriptionFormatExtensionTypes for the instance
func (instance *ServiceDescriptionFormatSettings) GetPropertyServiceDescriptionFormatExtensionTypes() (value []TypeElement, err error) {
	retValue, err := instance.GetProperty("ServiceDescriptionFormatExtensionTypes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TypeElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TypeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TypeElement(valuetmp))
	}

	return
}
