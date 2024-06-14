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

// SoapExtensionTypesInfo struct
type SoapExtensionTypesInfo struct {
	*EmbeddedObject

	//
	SoapExtensionTypes []SoapExtensionTypeElement
}

func NewSoapExtensionTypesInfoEx1(instance *cim.WmiInstance) (newInstance *SoapExtensionTypesInfo, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SoapExtensionTypesInfo{
		EmbeddedObject: tmp,
	}
	return
}

func NewSoapExtensionTypesInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SoapExtensionTypesInfo, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SoapExtensionTypesInfo{
		EmbeddedObject: tmp,
	}
	return
}

// SetSoapExtensionTypes sets the value of SoapExtensionTypes for the instance
func (instance *SoapExtensionTypesInfo) SetPropertySoapExtensionTypes(value []SoapExtensionTypeElement) (err error) {
	return instance.SetProperty("SoapExtensionTypes", (value))
}

// GetSoapExtensionTypes gets the value of SoapExtensionTypes for the instance
func (instance *SoapExtensionTypesInfo) GetPropertySoapExtensionTypes() (value []SoapExtensionTypeElement, err error) {
	retValue, err := instance.GetProperty("SoapExtensionTypes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SoapExtensionTypeElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SoapExtensionTypeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SoapExtensionTypeElement(valuetmp))
	}

	return
}
