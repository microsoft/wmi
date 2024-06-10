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

// PartialTrustVisibleAssembliesSection struct
type PartialTrustVisibleAssembliesSection struct {
	*ConfigurationSectionWithCollection

	//
	PartialTrustVisibleAssemblies []AssemblyCollectionElement
}

func NewPartialTrustVisibleAssembliesSectionEx1(instance *cim.WmiInstance) (newInstance *PartialTrustVisibleAssembliesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PartialTrustVisibleAssembliesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewPartialTrustVisibleAssembliesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PartialTrustVisibleAssembliesSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PartialTrustVisibleAssembliesSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetPartialTrustVisibleAssemblies sets the value of PartialTrustVisibleAssemblies for the instance
func (instance *PartialTrustVisibleAssembliesSection) SetPropertyPartialTrustVisibleAssemblies(value []AssemblyCollectionElement) (err error) {
	return instance.SetProperty("PartialTrustVisibleAssemblies", (value))
}

// GetPartialTrustVisibleAssemblies gets the value of PartialTrustVisibleAssemblies for the instance
func (instance *PartialTrustVisibleAssembliesSection) GetPropertyPartialTrustVisibleAssemblies() (value []AssemblyCollectionElement, err error) {
	retValue, err := instance.GetProperty("PartialTrustVisibleAssemblies")
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
