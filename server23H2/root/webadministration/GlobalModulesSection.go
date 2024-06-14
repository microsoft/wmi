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

// GlobalModulesSection struct
type GlobalModulesSection struct {
	*ConfigurationSectionWithCollection

	//
	GlobalModules []GlobalModuleElement
}

func NewGlobalModulesSectionEx1(instance *cim.WmiInstance) (newInstance *GlobalModulesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &GlobalModulesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewGlobalModulesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *GlobalModulesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &GlobalModulesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetGlobalModules sets the value of GlobalModules for the instance
func (instance *GlobalModulesSection) SetPropertyGlobalModules(value []GlobalModuleElement) (err error) {
	return instance.SetProperty("GlobalModules", (value))
}

// GetGlobalModules gets the value of GlobalModules for the instance
func (instance *GlobalModulesSection) GetPropertyGlobalModules() (value []GlobalModuleElement, err error) {
	retValue, err := instance.GetProperty("GlobalModules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(GlobalModuleElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " GlobalModuleElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, GlobalModuleElement(valuetmp))
	}

	return
}
