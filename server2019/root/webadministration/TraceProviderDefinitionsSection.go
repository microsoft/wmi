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

// TraceProviderDefinitionsSection struct
type TraceProviderDefinitionsSection struct {
	*ConfigurationSectionWithCollection

	//
	TraceProviderDefinitions []TraceProviderDefinition
}

func NewTraceProviderDefinitionsSectionEx1(instance *cim.WmiInstance) (newInstance *TraceProviderDefinitionsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TraceProviderDefinitionsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewTraceProviderDefinitionsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TraceProviderDefinitionsSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TraceProviderDefinitionsSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetTraceProviderDefinitions sets the value of TraceProviderDefinitions for the instance
func (instance *TraceProviderDefinitionsSection) SetPropertyTraceProviderDefinitions(value []TraceProviderDefinition) (err error) {
	return instance.SetProperty("TraceProviderDefinitions", (value))
}

// GetTraceProviderDefinitions gets the value of TraceProviderDefinitions for the instance
func (instance *TraceProviderDefinitionsSection) GetPropertyTraceProviderDefinitions() (value []TraceProviderDefinition, err error) {
	retValue, err := instance.GetProperty("TraceProviderDefinitions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(TraceProviderDefinition)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " TraceProviderDefinition is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TraceProviderDefinition(valuetmp))
	}

	return
}
