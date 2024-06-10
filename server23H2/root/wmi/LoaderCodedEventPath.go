// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// LoaderCodedEventPath struct
type LoaderCodedEventPath struct {
	*Image

	//
	BaseAddress uint64

	//
	Code int8

	//
	ErrorOpcode uint8

	//
	String1 string

	//
	String2 string
}

func NewLoaderCodedEventPathEx1(instance *cim.WmiInstance) (newInstance *LoaderCodedEventPath, err error) {
	tmp, err := NewImageEx1(instance)

	if err != nil {
		return
	}
	newInstance = &LoaderCodedEventPath{
		Image: tmp,
	}
	return
}

func NewLoaderCodedEventPathEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LoaderCodedEventPath, err error) {
	tmp, err := NewImageEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LoaderCodedEventPath{
		Image: tmp,
	}
	return
}

// SetBaseAddress sets the value of BaseAddress for the instance
func (instance *LoaderCodedEventPath) SetPropertyBaseAddress(value uint64) (err error) {
	return instance.SetProperty("BaseAddress", (value))
}

// GetBaseAddress gets the value of BaseAddress for the instance
func (instance *LoaderCodedEventPath) GetPropertyBaseAddress() (value uint64, err error) {
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
func (instance *LoaderCodedEventPath) SetPropertyCode(value int8) (err error) {
	return instance.SetProperty("Code", (value))
}

// GetCode gets the value of Code for the instance
func (instance *LoaderCodedEventPath) GetPropertyCode() (value int8, err error) {
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
func (instance *LoaderCodedEventPath) SetPropertyErrorOpcode(value uint8) (err error) {
	return instance.SetProperty("ErrorOpcode", (value))
}

// GetErrorOpcode gets the value of ErrorOpcode for the instance
func (instance *LoaderCodedEventPath) GetPropertyErrorOpcode() (value uint8, err error) {
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

// SetString1 sets the value of String1 for the instance
func (instance *LoaderCodedEventPath) SetPropertyString1(value string) (err error) {
	return instance.SetProperty("String1", (value))
}

// GetString1 gets the value of String1 for the instance
func (instance *LoaderCodedEventPath) GetPropertyString1() (value string, err error) {
	retValue, err := instance.GetProperty("String1")
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

// SetString2 sets the value of String2 for the instance
func (instance *LoaderCodedEventPath) SetPropertyString2(value string) (err error) {
	return instance.SetProperty("String2", (value))
}

// GetString2 gets the value of String2 for the instance
func (instance *LoaderCodedEventPath) GetPropertyString2() (value string, err error) {
	retValue, err := instance.GetProperty("String2")
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
