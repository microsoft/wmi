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

// WebRequestModulesSection struct
type WebRequestModulesSection struct {
	*ConfigurationSectionWithCollection

	//
	WebRequestModules []WebRequestModuleElement
}

func NewWebRequestModulesSectionEx1(instance *cim.WmiInstance) (newInstance *WebRequestModulesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebRequestModulesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewWebRequestModulesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebRequestModulesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebRequestModulesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetWebRequestModules sets the value of WebRequestModules for the instance
func (instance *WebRequestModulesSection) SetPropertyWebRequestModules(value []WebRequestModuleElement) (err error) {
	return instance.SetProperty("WebRequestModules", (value))
}

// GetWebRequestModules gets the value of WebRequestModules for the instance
func (instance *WebRequestModulesSection) GetPropertyWebRequestModules() (value []WebRequestModuleElement, err error) {
	retValue, err := instance.GetProperty("WebRequestModules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(WebRequestModuleElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " WebRequestModuleElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, WebRequestModuleElement(valuetmp))
	}

	return
}
