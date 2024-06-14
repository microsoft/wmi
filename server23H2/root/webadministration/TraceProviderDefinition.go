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

// TraceProviderDefinition struct
type TraceProviderDefinition struct {
	*CollectionElement

	//
	Areas TraceAreaSettings

	//
	Guid string

	//
	Name string
}

func NewTraceProviderDefinitionEx1(instance *cim.WmiInstance) (newInstance *TraceProviderDefinition, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TraceProviderDefinition{
		CollectionElement: tmp,
	}
	return
}

func NewTraceProviderDefinitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TraceProviderDefinition, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TraceProviderDefinition{
		CollectionElement: tmp,
	}
	return
}

// SetAreas sets the value of Areas for the instance
func (instance *TraceProviderDefinition) SetPropertyAreas(value TraceAreaSettings) (err error) {
	return instance.SetProperty("Areas", (value))
}

// GetAreas gets the value of Areas for the instance
func (instance *TraceProviderDefinition) GetPropertyAreas() (value TraceAreaSettings, err error) {
	retValue, err := instance.GetProperty("Areas")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(TraceAreaSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " TraceAreaSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TraceAreaSettings(valuetmp)

	return
}

// SetGuid sets the value of Guid for the instance
func (instance *TraceProviderDefinition) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *TraceProviderDefinition) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *TraceProviderDefinition) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *TraceProviderDefinition) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
