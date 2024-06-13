// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// LoaderCodedEvent struct
type LoaderCodedEvent struct {
	*Image_V2

	//
	BaseAddress uint64

	//
	Code int8

	//
	ErrorOpcode uint8

	//
	String string
}

func NewLoaderCodedEventEx1(instance *cim.WmiInstance) (newInstance *LoaderCodedEvent, err error) {
	tmp, err := NewImage_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &LoaderCodedEvent{
		Image_V2: tmp,
	}
	return
}

func NewLoaderCodedEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LoaderCodedEvent, err error) {
	tmp, err := NewImage_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LoaderCodedEvent{
		Image_V2: tmp,
	}
	return
}

// SetBaseAddress sets the value of BaseAddress for the instance
func (instance *LoaderCodedEvent) SetPropertyBaseAddress(value uint64) (err error) {
	return instance.SetProperty("BaseAddress", (value))
}

// GetBaseAddress gets the value of BaseAddress for the instance
func (instance *LoaderCodedEvent) GetPropertyBaseAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("BaseAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCode sets the value of Code for the instance
func (instance *LoaderCodedEvent) SetPropertyCode(value int8) (err error) {
	return instance.SetProperty("Code", (value))
}

// GetCode gets the value of Code for the instance
func (instance *LoaderCodedEvent) GetPropertyCode() (value int8, err error) {
	retValue, err := instance.GetProperty("Code")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int8(valuetmp)

	return
}

// SetErrorOpcode sets the value of ErrorOpcode for the instance
func (instance *LoaderCodedEvent) SetPropertyErrorOpcode(value uint8) (err error) {
	return instance.SetProperty("ErrorOpcode", (value))
}

// GetErrorOpcode gets the value of ErrorOpcode for the instance
func (instance *LoaderCodedEvent) GetPropertyErrorOpcode() (value uint8, err error) {
	retValue, err := instance.GetProperty("ErrorOpcode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetString sets the value of String for the instance
func (instance *LoaderCodedEvent) SetPropertyString(value string) (err error) {
	return instance.SetProperty("String", (value))
}

// GetString gets the value of String for the instance
func (instance *LoaderCodedEvent) GetPropertyString() (value string, err error) {
	retValue, err := instance.GetProperty("String")
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
