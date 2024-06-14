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

// FullTrustAssembliesSection struct
type FullTrustAssembliesSection struct {
	*ConfigurationSectionWithCollection

	//
	FullTrustAssemblies []AssemblyCollectionElement
}

func NewFullTrustAssembliesSectionEx1(instance *cim.WmiInstance) (newInstance *FullTrustAssembliesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FullTrustAssembliesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewFullTrustAssembliesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FullTrustAssembliesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FullTrustAssembliesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetFullTrustAssemblies sets the value of FullTrustAssemblies for the instance
func (instance *FullTrustAssembliesSection) SetPropertyFullTrustAssemblies(value []AssemblyCollectionElement) (err error) {
	return instance.SetProperty("FullTrustAssemblies", (value))
}

// GetFullTrustAssemblies gets the value of FullTrustAssemblies for the instance
func (instance *FullTrustAssembliesSection) GetPropertyFullTrustAssemblies() (value []AssemblyCollectionElement, err error) {
	retValue, err := instance.GetProperty("FullTrustAssemblies")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(AssemblyCollectionElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " AssemblyCollectionElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, AssemblyCollectionElement(valuetmp))
	}

	return
}
