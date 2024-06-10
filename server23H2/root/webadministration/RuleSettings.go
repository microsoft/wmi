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

// RuleSettings struct
type RuleSettings struct {
	*EmbeddedObject

	//
	Rules []RuleElement
}

func NewRuleSettingsEx1(instance *cim.WmiInstance) (newInstance *RuleSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RuleSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewRuleSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RuleSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RuleSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetRules sets the value of Rules for the instance
func (instance *RuleSettings) SetPropertyRules(value []RuleElement) (err error) {
	return instance.SetProperty("Rules", (value))
}

// GetRules gets the value of Rules for the instance
func (instance *RuleSettings) GetPropertyRules() (value []RuleElement, err error) {
	retValue, err := instance.GetProperty("Rules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(RuleElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " RuleElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, RuleElement(valuetmp))
	}

	return
}
