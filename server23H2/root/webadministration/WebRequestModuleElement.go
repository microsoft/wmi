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

// WebRequestModuleElement struct
type WebRequestModuleElement struct {
	*CollectionElement

	//
	Prefix string

	//
	Type string
}

func NewWebRequestModuleElementEx1(instance *cim.WmiInstance) (newInstance *WebRequestModuleElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebRequestModuleElement{
		CollectionElement: tmp,
	}
	return
}

func NewWebRequestModuleElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebRequestModuleElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebRequestModuleElement{
		CollectionElement: tmp,
	}
	return
}

// SetPrefix sets the value of Prefix for the instance
func (instance *WebRequestModuleElement) SetPropertyPrefix(value string) (err error) {
	return instance.SetProperty("Prefix", (value))
}

// GetPrefix gets the value of Prefix for the instance
func (instance *WebRequestModuleElement) GetPropertyPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("Prefix")
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

// SetType sets the value of Type for the instance
func (instance *WebRequestModuleElement) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *WebRequestModuleElement) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
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
