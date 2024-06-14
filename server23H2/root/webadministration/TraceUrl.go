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

// TraceUrl struct
type TraceUrl struct {
	*CollectionElement

	//
	CustomActionExe string

	//
	CustomActionParams string

	//
	CustomActionTriggerLimit uint32

	//
	FailureDefinitions FailureDefinition

	//
	Path string

	//
	TraceAreas TraceUrlAreaSettings
}

func NewTraceUrlEx1(instance *cim.WmiInstance) (newInstance *TraceUrl, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TraceUrl{
		CollectionElement: tmp,
	}
	return
}

func NewTraceUrlEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TraceUrl, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TraceUrl{
		CollectionElement: tmp,
	}
	return
}

// SetCustomActionExe sets the value of CustomActionExe for the instance
func (instance *TraceUrl) SetPropertyCustomActionExe(value string) (err error) {
	return instance.SetProperty("CustomActionExe", (value))
}

// GetCustomActionExe gets the value of CustomActionExe for the instance
func (instance *TraceUrl) GetPropertyCustomActionExe() (value string, err error) {
	retValue, err := instance.GetProperty("CustomActionExe")
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

// SetCustomActionParams sets the value of CustomActionParams for the instance
func (instance *TraceUrl) SetPropertyCustomActionParams(value string) (err error) {
	return instance.SetProperty("CustomActionParams", (value))
}

// GetCustomActionParams gets the value of CustomActionParams for the instance
func (instance *TraceUrl) GetPropertyCustomActionParams() (value string, err error) {
	retValue, err := instance.GetProperty("CustomActionParams")
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

// SetCustomActionTriggerLimit sets the value of CustomActionTriggerLimit for the instance
func (instance *TraceUrl) SetPropertyCustomActionTriggerLimit(value uint32) (err error) {
	return instance.SetProperty("CustomActionTriggerLimit", (value))
}

// GetCustomActionTriggerLimit gets the value of CustomActionTriggerLimit for the instance
func (instance *TraceUrl) GetPropertyCustomActionTriggerLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("CustomActionTriggerLimit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetFailureDefinitions sets the value of FailureDefinitions for the instance
func (instance *TraceUrl) SetPropertyFailureDefinitions(value FailureDefinition) (err error) {
	return instance.SetProperty("FailureDefinitions", (value))
}

// GetFailureDefinitions gets the value of FailureDefinitions for the instance
func (instance *TraceUrl) GetPropertyFailureDefinitions() (value FailureDefinition, err error) {
	retValue, err := instance.GetProperty("FailureDefinitions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FailureDefinition)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FailureDefinition is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FailureDefinition(valuetmp)

	return
}

// SetPath sets the value of Path for the instance
func (instance *TraceUrl) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *TraceUrl) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetTraceAreas sets the value of TraceAreas for the instance
func (instance *TraceUrl) SetPropertyTraceAreas(value TraceUrlAreaSettings) (err error) {
	return instance.SetProperty("TraceAreas", (value))
}

// GetTraceAreas gets the value of TraceAreas for the instance
func (instance *TraceUrl) GetPropertyTraceAreas() (value TraceUrlAreaSettings, err error) {
	retValue, err := instance.GetProperty("TraceAreas")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(TraceUrlAreaSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " TraceUrlAreaSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TraceUrlAreaSettings(valuetmp)

	return
}
