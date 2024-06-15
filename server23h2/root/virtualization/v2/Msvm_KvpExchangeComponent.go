// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_KvpExchangeComponent struct
type Msvm_KvpExchangeComponent struct {
	*CIM_LogicalDevice

	//
	GuestExchangeItems []string

	//
	GuestIntrinsicExchangeItems []string
}

func NewMsvm_KvpExchangeComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_KvpExchangeComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_KvpExchangeComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewMsvm_KvpExchangeComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_KvpExchangeComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_KvpExchangeComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetGuestExchangeItems sets the value of GuestExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponent) SetPropertyGuestExchangeItems(value []string) (err error) {
	return instance.SetProperty("GuestExchangeItems", (value))
}

// GetGuestExchangeItems gets the value of GuestExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponent) GetPropertyGuestExchangeItems() (value []string, err error) {
	retValue, err := instance.GetProperty("GuestExchangeItems")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetGuestIntrinsicExchangeItems sets the value of GuestIntrinsicExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponent) SetPropertyGuestIntrinsicExchangeItems(value []string) (err error) {
	return instance.SetProperty("GuestIntrinsicExchangeItems", (value))
}

// GetGuestIntrinsicExchangeItems gets the value of GuestIntrinsicExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponent) GetPropertyGuestIntrinsicExchangeItems() (value []string, err error) {
	retValue, err := instance.GetProperty("GuestIntrinsicExchangeItems")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
