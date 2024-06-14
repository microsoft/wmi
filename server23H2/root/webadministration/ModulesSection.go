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

// ModulesSection struct
type ModulesSection struct {
	*ConfigurationSectionWithCollection

	//
	Modules []ModuleAction

	//
	RunAllManagedModulesForAllRequests bool

	//
	RunManagedModulesForWebDavRequests bool
}

func NewModulesSectionEx1(instance *cim.WmiInstance) (newInstance *ModulesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ModulesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewModulesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ModulesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ModulesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetModules sets the value of Modules for the instance
func (instance *ModulesSection) SetPropertyModules(value []ModuleAction) (err error) {
	return instance.SetProperty("Modules", (value))
}

// GetModules gets the value of Modules for the instance
func (instance *ModulesSection) GetPropertyModules() (value []ModuleAction, err error) {
	retValue, err := instance.GetProperty("Modules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ModuleAction)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ModuleAction is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ModuleAction(valuetmp))
	}

	return
}

// SetRunAllManagedModulesForAllRequests sets the value of RunAllManagedModulesForAllRequests for the instance
func (instance *ModulesSection) SetPropertyRunAllManagedModulesForAllRequests(value bool) (err error) {
	return instance.SetProperty("RunAllManagedModulesForAllRequests", (value))
}

// GetRunAllManagedModulesForAllRequests gets the value of RunAllManagedModulesForAllRequests for the instance
func (instance *ModulesSection) GetPropertyRunAllManagedModulesForAllRequests() (value bool, err error) {
	retValue, err := instance.GetProperty("RunAllManagedModulesForAllRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetRunManagedModulesForWebDavRequests sets the value of RunManagedModulesForWebDavRequests for the instance
func (instance *ModulesSection) SetPropertyRunManagedModulesForWebDavRequests(value bool) (err error) {
	return instance.SetProperty("RunManagedModulesForWebDavRequests", (value))
}

// GetRunManagedModulesForWebDavRequests gets the value of RunManagedModulesForWebDavRequests for the instance
func (instance *ModulesSection) GetPropertyRunManagedModulesForWebDavRequests() (value bool, err error) {
	retValue, err := instance.GetProperty("RunManagedModulesForWebDavRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}