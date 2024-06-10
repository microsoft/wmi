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

// DeviceFilter struct
type DeviceFilter struct {
	*CollectionElement

	//
	Argument string

	//
	Compare string

	//
	Method string

	//
	Name string

	//
	Type string
}

func NewDeviceFilterEx1(instance *cim.WmiInstance) (newInstance *DeviceFilter, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DeviceFilter{
		CollectionElement: tmp,
	}
	return
}

func NewDeviceFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DeviceFilter, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DeviceFilter{
		CollectionElement: tmp,
	}
	return
}

// SetArgument sets the value of Argument for the instance
func (instance *DeviceFilter) SetPropertyArgument(value string) (err error) {
	return instance.SetProperty("Argument", (value))
}

// GetArgument gets the value of Argument for the instance
func (instance *DeviceFilter) GetPropertyArgument() (value string, err error) {
	retValue, err := instance.GetProperty("Argument")
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

// SetCompare sets the value of Compare for the instance
func (instance *DeviceFilter) SetPropertyCompare(value string) (err error) {
	return instance.SetProperty("Compare", (value))
}

// GetCompare gets the value of Compare for the instance
func (instance *DeviceFilter) GetPropertyCompare() (value string, err error) {
	retValue, err := instance.GetProperty("Compare")
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

// SetMethod sets the value of Method for the instance
func (instance *DeviceFilter) SetPropertyMethod(value string) (err error) {
	return instance.SetProperty("Method", (value))
}

// GetMethod gets the value of Method for the instance
func (instance *DeviceFilter) GetPropertyMethod() (value string, err error) {
	retValue, err := instance.GetProperty("Method")
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
func (instance *DeviceFilter) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *DeviceFilter) GetPropertyName() (value string, err error) {
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

// SetType sets the value of Type for the instance
func (instance *DeviceFilter) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *DeviceFilter) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
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
