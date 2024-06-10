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

// TraceAreaElement struct
type TraceAreaElement struct {
	*CollectionElement

	//
	Areas string

	//
	provider string

	//
	Verbosity int32
}

func NewTraceAreaElementEx1(instance *cim.WmiInstance) (newInstance *TraceAreaElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TraceAreaElement{
		CollectionElement: tmp,
	}
	return
}

func NewTraceAreaElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TraceAreaElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TraceAreaElement{
		CollectionElement: tmp,
	}
	return
}

// SetAreas sets the value of Areas for the instance
func (instance *TraceAreaElement) SetPropertyAreas(value string) (err error) {
	return instance.SetProperty("Areas", (value))
}

// GetAreas gets the value of Areas for the instance
func (instance *TraceAreaElement) GetPropertyAreas() (value string, err error) {
	retValue, err := instance.GetProperty("Areas")
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

// Setprovider sets the value of provider for the instance
func (instance *TraceAreaElement) SetPropertyprovider(value string) (err error) {
	return instance.SetProperty("provider", (value))
}

// Getprovider gets the value of provider for the instance
func (instance *TraceAreaElement) GetPropertyprovider() (value string, err error) {
	retValue, err := instance.GetProperty("provider")
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

// SetVerbosity sets the value of Verbosity for the instance
func (instance *TraceAreaElement) SetPropertyVerbosity(value int32) (err error) {
	return instance.SetProperty("Verbosity", (value))
}

// GetVerbosity gets the value of Verbosity for the instance
func (instance *TraceAreaElement) GetPropertyVerbosity() (value int32, err error) {
	retValue, err := instance.GetProperty("Verbosity")
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
