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

// ProfileElement struct
type ProfileElement struct {
	*CollectionElement

	//
	Custom string

	//
	MaxLimit int32

	//
	MinInstances int32

	//
	MinInterval string

	//
	Name string
}

func NewProfileElementEx1(instance *cim.WmiInstance) (newInstance *ProfileElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProfileElement{
		CollectionElement: tmp,
	}
	return
}

func NewProfileElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProfileElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProfileElement{
		CollectionElement: tmp,
	}
	return
}

// SetCustom sets the value of Custom for the instance
func (instance *ProfileElement) SetPropertyCustom(value string) (err error) {
	return instance.SetProperty("Custom", (value))
}

// GetCustom gets the value of Custom for the instance
func (instance *ProfileElement) GetPropertyCustom() (value string, err error) {
	retValue, err := instance.GetProperty("Custom")
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

// SetMaxLimit sets the value of MaxLimit for the instance
func (instance *ProfileElement) SetPropertyMaxLimit(value int32) (err error) {
	return instance.SetProperty("MaxLimit", (value))
}

// GetMaxLimit gets the value of MaxLimit for the instance
func (instance *ProfileElement) GetPropertyMaxLimit() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxLimit")
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

// SetMinInstances sets the value of MinInstances for the instance
func (instance *ProfileElement) SetPropertyMinInstances(value int32) (err error) {
	return instance.SetProperty("MinInstances", (value))
}

// GetMinInstances gets the value of MinInstances for the instance
func (instance *ProfileElement) GetPropertyMinInstances() (value int32, err error) {
	retValue, err := instance.GetProperty("MinInstances")
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

// SetMinInterval sets the value of MinInterval for the instance
func (instance *ProfileElement) SetPropertyMinInterval(value string) (err error) {
	return instance.SetProperty("MinInterval", (value))
}

// GetMinInterval gets the value of MinInterval for the instance
func (instance *ProfileElement) GetPropertyMinInterval() (value string, err error) {
	retValue, err := instance.GetProperty("MinInterval")
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
func (instance *ProfileElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ProfileElement) GetPropertyName() (value string, err error) {
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
