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

// WebDavPropertiesElement struct
type WebDavPropertiesElement struct {
	*CollectionElement

	//
	PropertyStore string

	//
	XmlNamespace string
}

func NewWebDavPropertiesElementEx1(instance *cim.WmiInstance) (newInstance *WebDavPropertiesElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebDavPropertiesElement{
		CollectionElement: tmp,
	}
	return
}

func NewWebDavPropertiesElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebDavPropertiesElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebDavPropertiesElement{
		CollectionElement: tmp,
	}
	return
}

// SetPropertyStore sets the value of PropertyStore for the instance
func (instance *WebDavPropertiesElement) SetPropertyPropertyStore(value string) (err error) {
	return instance.SetProperty("PropertyStore", (value))
}

// GetPropertyStore gets the value of PropertyStore for the instance
func (instance *WebDavPropertiesElement) GetPropertyPropertyStore() (value string, err error) {
	retValue, err := instance.GetProperty("PropertyStore")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetXmlNamespace sets the value of XmlNamespace for the instance
func (instance *WebDavPropertiesElement) SetPropertyXmlNamespace(value string) (err error) {
	return instance.SetProperty("XmlNamespace", (value))
}

// GetXmlNamespace gets the value of XmlNamespace for the instance
func (instance *WebDavPropertiesElement) GetPropertyXmlNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("XmlNamespace")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
