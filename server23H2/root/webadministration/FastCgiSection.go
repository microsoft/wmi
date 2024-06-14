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

// FastCgiSection struct
type FastCgiSection struct {
	*ConfigurationSectionWithCollection

	//
	FastCgi []FastCgiApplicationElement
}

func NewFastCgiSectionEx1(instance *cim.WmiInstance) (newInstance *FastCgiSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FastCgiSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewFastCgiSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FastCgiSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FastCgiSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetFastCgi sets the value of FastCgi for the instance
func (instance *FastCgiSection) SetPropertyFastCgi(value []FastCgiApplicationElement) (err error) {
	return instance.SetProperty("FastCgi", (value))
}

// GetFastCgi gets the value of FastCgi for the instance
func (instance *FastCgiSection) GetPropertyFastCgi() (value []FastCgiApplicationElement, err error) {
	retValue, err := instance.GetProperty("FastCgi")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FastCgiApplicationElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FastCgiApplicationElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FastCgiApplicationElement(valuetmp))
	}

	return
}
