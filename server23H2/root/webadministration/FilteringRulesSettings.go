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

// FilteringRulesSettings struct
type FilteringRulesSettings struct {
	*EmbeddedObject

	//
	FilteringRules []FilteringRulesElement
}

func NewFilteringRulesSettingsEx1(instance *cim.WmiInstance) (newInstance *FilteringRulesSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FilteringRulesSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFilteringRulesSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FilteringRulesSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FilteringRulesSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetFilteringRules sets the value of FilteringRules for the instance
func (instance *FilteringRulesSettings) SetPropertyFilteringRules(value []FilteringRulesElement) (err error) {
	return instance.SetProperty("FilteringRules", (value))
}

// GetFilteringRules gets the value of FilteringRules for the instance
func (instance *FilteringRulesSettings) GetPropertyFilteringRules() (value []FilteringRulesElement, err error) {
	retValue, err := instance.GetProperty("FilteringRules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FilteringRulesElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FilteringRulesElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FilteringRulesElement(valuetmp))
	}

	return
}
