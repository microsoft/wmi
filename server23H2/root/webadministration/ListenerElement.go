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

// ListenerElement struct
type ListenerElement struct {
	*TypedElement

	//
	Filter TypedElement

	//
	Name string

	//
	TraceOutputOptions int32
}

func NewListenerElementEx1(instance *cim.WmiInstance) (newInstance *ListenerElement, err error) {
	tmp, err := NewTypedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ListenerElement{
		TypedElement: tmp,
	}
	return
}

func NewListenerElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ListenerElement, err error) {
	tmp, err := NewTypedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ListenerElement{
		TypedElement: tmp,
	}
	return
}

// SetFilter sets the value of Filter for the instance
func (instance *ListenerElement) SetPropertyFilter(value TypedElement) (err error) {
	return instance.SetProperty("Filter", (value))
}

// GetFilter gets the value of Filter for the instance
func (instance *ListenerElement) GetPropertyFilter() (value TypedElement, err error) {
	retValue, err := instance.GetProperty("Filter")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(TypedElement)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " TypedElement is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TypedElement(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *ListenerElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ListenerElement) GetPropertyName() (value string, err error) {
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

// SetTraceOutputOptions sets the value of TraceOutputOptions for the instance
func (instance *ListenerElement) SetPropertyTraceOutputOptions(value int32) (err error) {
	return instance.SetProperty("TraceOutputOptions", (value))
}

// GetTraceOutputOptions gets the value of TraceOutputOptions for the instance
func (instance *ListenerElement) GetPropertyTraceOutputOptions() (value int32, err error) {
	retValue, err := instance.GetProperty("TraceOutputOptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
