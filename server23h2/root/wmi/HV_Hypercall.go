// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HV_Hypercall struct
type HV_Hypercall struct {
	*PerfInfo_V2

	//
	CallCode uint32

	//
	IsFast uint8

	//
	IsNested uint8
}

func NewHV_HypercallEx1(instance *cim.WmiInstance) (newInstance *HV_Hypercall, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &HV_Hypercall{
		PerfInfo_V2: tmp,
	}
	return
}

func NewHV_HypercallEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HV_Hypercall, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HV_Hypercall{
		PerfInfo_V2: tmp,
	}
	return
}

// SetCallCode sets the value of CallCode for the instance
func (instance *HV_Hypercall) SetPropertyCallCode(value uint32) (err error) {
	return instance.SetProperty("CallCode", (value))
}

// GetCallCode gets the value of CallCode for the instance
func (instance *HV_Hypercall) GetPropertyCallCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("CallCode")
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

// SetIsFast sets the value of IsFast for the instance
func (instance *HV_Hypercall) SetPropertyIsFast(value uint8) (err error) {
	return instance.SetProperty("IsFast", (value))
}

// GetIsFast gets the value of IsFast for the instance
func (instance *HV_Hypercall) GetPropertyIsFast() (value uint8, err error) {
	retValue, err := instance.GetProperty("IsFast")
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

// SetIsNested sets the value of IsNested for the instance
func (instance *HV_Hypercall) SetPropertyIsNested(value uint8) (err error) {
	return instance.SetProperty("IsNested", (value))
}

// GetIsNested gets the value of IsNested for the instance
func (instance *HV_Hypercall) GetPropertyIsNested() (value uint8, err error) {
	retValue, err := instance.GetProperty("IsNested")
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
