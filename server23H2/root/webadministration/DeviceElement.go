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

// DeviceElement struct
type DeviceElement struct {
	*CollectionElement

	//
	Device []AdapterElement

	//
	InheritsFrom string

	//
	Name string

	//
	PageAdapter string

	//
	PredicateClass string

	//
	PredicateMethod string
}

func NewDeviceElementEx1(instance *cim.WmiInstance) (newInstance *DeviceElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DeviceElement{
		CollectionElement: tmp,
	}
	return
}

func NewDeviceElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DeviceElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DeviceElement{
		CollectionElement: tmp,
	}
	return
}

// SetDevice sets the value of Device for the instance
func (instance *DeviceElement) SetPropertyDevice(value []AdapterElement) (err error) {
	return instance.SetProperty("Device", (value))
}

// GetDevice gets the value of Device for the instance
func (instance *DeviceElement) GetPropertyDevice() (value []AdapterElement, err error) {
	retValue, err := instance.GetProperty("Device")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(AdapterElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " AdapterElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, AdapterElement(valuetmp))
	}

	return
}

// SetInheritsFrom sets the value of InheritsFrom for the instance
func (instance *DeviceElement) SetPropertyInheritsFrom(value string) (err error) {
	return instance.SetProperty("InheritsFrom", (value))
}

// GetInheritsFrom gets the value of InheritsFrom for the instance
func (instance *DeviceElement) GetPropertyInheritsFrom() (value string, err error) {
	retValue, err := instance.GetProperty("InheritsFrom")
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
func (instance *DeviceElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *DeviceElement) GetPropertyName() (value string, err error) {
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

// SetPageAdapter sets the value of PageAdapter for the instance
func (instance *DeviceElement) SetPropertyPageAdapter(value string) (err error) {
	return instance.SetProperty("PageAdapter", (value))
}

// GetPageAdapter gets the value of PageAdapter for the instance
func (instance *DeviceElement) GetPropertyPageAdapter() (value string, err error) {
	retValue, err := instance.GetProperty("PageAdapter")
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

// SetPredicateClass sets the value of PredicateClass for the instance
func (instance *DeviceElement) SetPropertyPredicateClass(value string) (err error) {
	return instance.SetProperty("PredicateClass", (value))
}

// GetPredicateClass gets the value of PredicateClass for the instance
func (instance *DeviceElement) GetPropertyPredicateClass() (value string, err error) {
	retValue, err := instance.GetProperty("PredicateClass")
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

// SetPredicateMethod sets the value of PredicateMethod for the instance
func (instance *DeviceElement) SetPropertyPredicateMethod(value string) (err error) {
	return instance.SetProperty("PredicateMethod", (value))
}

// GetPredicateMethod gets the value of PredicateMethod for the instance
func (instance *DeviceElement) GetPropertyPredicateMethod() (value string, err error) {
	retValue, err := instance.GetProperty("PredicateMethod")
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
