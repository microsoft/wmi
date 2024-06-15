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

// ISR struct
type ISR struct {
	*PerfInfo_V2

	//
	InitialTime interface{}

	//
	Reserved uint8

	//
	ReturnValue uint8

	//
	Routine uint32

	//
	Vector uint16
}

func NewISREx1(instance *cim.WmiInstance) (newInstance *ISR, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ISR{
		PerfInfo_V2: tmp,
	}
	return
}

func NewISREx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISR, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISR{
		PerfInfo_V2: tmp,
	}
	return
}

// SetInitialTime sets the value of InitialTime for the instance
func (instance *ISR) SetPropertyInitialTime(value interface{}) (err error) {
	return instance.SetProperty("InitialTime", (value))
}

// GetInitialTime gets the value of InitialTime for the instance
func (instance *ISR) GetPropertyInitialTime() (value interface{}, err error) {
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

// SetReserved sets the value of Reserved for the instance
func (instance *ISR) SetPropertyReserved(value uint8) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *ISR) GetPropertyReserved() (value uint8, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetReturnValue sets the value of ReturnValue for the instance
func (instance *ISR) SetPropertyReturnValue(value uint8) (err error) {
	return instance.SetProperty("ReturnValue", (value))
}

// GetReturnValue gets the value of ReturnValue for the instance
func (instance *ISR) GetPropertyReturnValue() (value uint8, err error) {
	retValue, err := instance.GetProperty("ReturnValue")
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

// SetRoutine sets the value of Routine for the instance
func (instance *ISR) SetPropertyRoutine(value uint32) (err error) {
	return instance.SetProperty("Routine", (value))
}

// GetRoutine gets the value of Routine for the instance
func (instance *ISR) GetPropertyRoutine() (value uint32, err error) {
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

// SetVector sets the value of Vector for the instance
func (instance *ISR) SetPropertyVector(value uint16) (err error) {
	return instance.SetProperty("Vector", (value))
}

// GetVector gets the value of Vector for the instance
func (instance *ISR) GetPropertyVector() (value uint16, err error) {
	retValue, err := instance.GetProperty("Vector")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
