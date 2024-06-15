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

// ISR_V1 struct
type ISR_V1 struct {
	*PerfInfo_V1

	//
	InitialTime interface{}

	//
	ReturnValue uint32

	//
	Routine uint32
}

func NewISR_V1Ex1(instance *cim.WmiInstance) (newInstance *ISR_V1, err error) {
	tmp, err := NewPerfInfo_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ISR_V1{
		PerfInfo_V1: tmp,
	}
	return
}

func NewISR_V1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISR_V1, err error) {
	tmp, err := NewPerfInfo_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISR_V1{
		PerfInfo_V1: tmp,
	}
	return
}

// SetInitialTime sets the value of InitialTime for the instance
func (instance *ISR_V1) SetPropertyInitialTime(value interface{}) (err error) {
	return instance.SetProperty("InitialTime", (value))
}

// GetInitialTime gets the value of InitialTime for the instance
func (instance *ISR_V1) GetPropertyInitialTime() (value interface{}, err error) {
	retValue, err := instance.GetProperty("InitialTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetReturnValue sets the value of ReturnValue for the instance
func (instance *ISR_V1) SetPropertyReturnValue(value uint32) (err error) {
	return instance.SetProperty("ReturnValue", (value))
}

// GetReturnValue gets the value of ReturnValue for the instance
func (instance *ISR_V1) GetPropertyReturnValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReturnValue")
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

// SetRoutine sets the value of Routine for the instance
func (instance *ISR_V1) SetPropertyRoutine(value uint32) (err error) {
	return instance.SetProperty("Routine", (value))
}

// GetRoutine gets the value of Routine for the instance
func (instance *ISR_V1) GetPropertyRoutine() (value uint32, err error) {
	retValue, err := instance.GetProperty("Routine")
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
