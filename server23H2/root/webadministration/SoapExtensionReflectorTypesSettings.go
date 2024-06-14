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

// SoapExtensionReflectorTypesSettings struct
type SoapExtensionReflectorTypesSettings struct {
	*EmbeddedObject

	//
	SoapExtensionReflectorTypes []TypeElement
}

func NewSoapExtensionReflectorTypesSettingsEx1(instance *cim.WmiInstance) (newInstance *SoapExtensionReflectorTypesSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SoapExtensionReflectorTypesSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewSoapExtensionReflectorTypesSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SoapExtensionReflectorTypesSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SoapExtensionReflectorTypesSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetSoapExtensionReflectorTypes sets the value of SoapExtensionReflectorTypes for the instance
func (instance *SoapExtensionReflectorTypesSettings) SetPropertySoapExtensionReflectorTypes(value []TypeElement) (err error) {
	return instance.SetProperty("SoapExtensionReflectorTypes", (value))
}

// GetSoapExtensionReflectorTypes gets the value of SoapExtensionReflectorTypes for the instance
func (instance *SoapExtensionReflectorTypesSettings) GetPropertySoapExtensionReflectorTypes() (value []TypeElement, err error) {
	retValue, err := instance.GetProperty("SoapExtensionReflectorTypes")
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
