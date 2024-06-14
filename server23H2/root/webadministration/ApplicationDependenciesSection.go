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

// ApplicationDependenciesSection struct
type ApplicationDependenciesSection struct {
	*ConfigurationSectionWithCollection

	//
	ApplicationDependencies []ApplicationDependency
}

func NewApplicationDependenciesSectionEx1(instance *cim.WmiInstance) (newInstance *ApplicationDependenciesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ApplicationDependenciesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewApplicationDependenciesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ApplicationDependenciesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ApplicationDependenciesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetApplicationDependencies sets the value of ApplicationDependencies for the instance
func (instance *ApplicationDependenciesSection) SetPropertyApplicationDependencies(value []ApplicationDependency) (err error) {
	return instance.SetProperty("ApplicationDependencies", (value))
}

// GetApplicationDependencies gets the value of ApplicationDependencies for the instance
func (instance *ApplicationDependenciesSection) GetPropertyApplicationDependencies() (value []ApplicationDependency, err error) {
	retValue, err := instance.GetProperty("ApplicationDependencies")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ApplicationDependency)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ApplicationDependency is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ApplicationDependency(valuetmp))
	}

	return
}
