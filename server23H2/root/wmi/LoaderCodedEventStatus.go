// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// LoaderCodedEventStatus struct
type LoaderCodedEventStatus struct {
	*Image

	//
	BaseAddress uint64

	//
	Code int8

	//
	ErrorOpcode uint8

	//
	String string
}

func NewLoaderCodedEventStatusEx1(instance *cim.WmiInstance) (newInstance *LoaderCodedEventStatus, err error) {
	tmp, err := NewImageEx1(instance)

	if err != nil {
		return
	}
	newInstance = &LoaderCodedEventStatus{
		Image: tmp,
	}
	return
}

func NewLoaderCodedEventStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LoaderCodedEventStatus, err error) {
	tmp, err := NewImageEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LoaderCodedEventStatus{
		Image: tmp,
	}
	return
}

// SetBaseAddress sets the value of BaseAddress for the instance
func (instance *LoaderCodedEventStatus) SetPropertyBaseAddress(value uint64) (err error) {
	return instance.SetProperty("BaseAddress", (value))
}

// GetBaseAddress gets the value of BaseAddress for the instance
func (instance *LoaderCodedEventStatus) GetPropertyBaseAddress() (value uint64, err error) {
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
func (instance *LoaderCodedEventStatus) SetPropertyCode(value int8) (err error) {
	return instance.SetProperty("Code", (value))
}

// GetCode gets the value of Code for the instance
func (instance *LoaderCodedEventStatus) GetPropertyCode() (value int8, err error) {
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
func (instance *LoaderCodedEventStatus) SetPropertyErrorOpcode(value uint8) (err error) {
	return instance.SetProperty("ErrorOpcode", (value))
}

// GetErrorOpcode gets the value of ErrorOpcode for the instance
func (instance *LoaderCodedEventStatus) GetPropertyErrorOpcode() (value uint8, err error) {
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
func (instance *LoaderCodedEventStatus) SetPropertyString(value string) (err error) {
	return instance.SetProperty("String", (value))
}

// GetString gets the value of String for the instance
func (instance *LoaderCodedEventStatus) GetPropertyString() (value string, err error) {
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
