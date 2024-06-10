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

// IsapiFiltersSection struct
type IsapiFiltersSection struct {
	*ConfigurationSectionWithCollection

	//
	IsapiFilters []IsapiFilterElement
}

func NewIsapiFiltersSectionEx1(instance *cim.WmiInstance) (newInstance *IsapiFiltersSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &IsapiFiltersSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewIsapiFiltersSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *IsapiFiltersSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &IsapiFiltersSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetIsapiFilters sets the value of IsapiFilters for the instance
func (instance *IsapiFiltersSection) SetPropertyIsapiFilters(value []IsapiFilterElement) (err error) {
	return instance.SetProperty("IsapiFilters", (value))
}

// GetIsapiFilters gets the value of IsapiFilters for the instance
func (instance *IsapiFiltersSection) GetPropertyIsapiFilters() (value []IsapiFilterElement, err error) {
	retValue, err := instance.GetProperty("IsapiFilters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(IsapiFilterElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " IsapiFilterElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, IsapiFilterElement(valuetmp))
	}

	return
}
