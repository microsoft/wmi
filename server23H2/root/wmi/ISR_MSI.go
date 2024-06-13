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

// ISR_MSI struct
type ISR_MSI struct {
	*PerfInfo_V2

	//
	InitialTime interface{}

	//
	MessageNumber uint32

	//
	Reserved uint8

	//
	ReturnValue uint8

	//
	Routine uint32

	//
	Vector uint16
}

func NewISR_MSIEx1(instance *cim.WmiInstance) (newInstance *ISR_MSI, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ISR_MSI{
		PerfInfo_V2: tmp,
	}
	return
}

func NewISR_MSIEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISR_MSI, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISR_MSI{
		PerfInfo_V2: tmp,
	}
	return
}

// SetInitialTime sets the value of InitialTime for the instance
func (instance *ISR_MSI) SetPropertyInitialTime(value interface{}) (err error) {
	return instance.SetProperty("InitialTime", (value))
}

// GetInitialTime gets the value of InitialTime for the instance
func (instance *ISR_MSI) GetPropertyInitialTime() (value interface{}, err error) {
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

// SetMessageNumber sets the value of MessageNumber for the instance
func (instance *ISR_MSI) SetPropertyMessageNumber(value uint32) (err error) {
	return instance.SetProperty("MessageNumber", (value))
}

// GetMessageNumber gets the value of MessageNumber for the instance
func (instance *ISR_MSI) GetPropertyMessageNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("MessageNumber")
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

// SetReserved sets the value of Reserved for the instance
func (instance *ISR_MSI) SetPropertyReserved(value uint8) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *ISR_MSI) GetPropertyReserved() (value uint8, err error) {
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
func (instance *ISR_MSI) SetPropertyReturnValue(value uint8) (err error) {
	return instance.SetProperty("ReturnValue", (value))
}

// GetReturnValue gets the value of ReturnValue for the instance
func (instance *ISR_MSI) GetPropertyReturnValue() (value uint8, err error) {
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
func (instance *ISR_MSI) SetPropertyRoutine(value uint32) (err error) {
	return instance.SetProperty("Routine", (value))
}

// GetRoutine gets the value of Routine for the instance
func (instance *ISR_MSI) GetPropertyRoutine() (value uint32, err error) {
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
func (instance *ISR_MSI) SetPropertyVector(value uint16) (err error) {
	return instance.SetProperty("Vector", (value))
}

// GetVector gets the value of Vector for the instance
func (instance *ISR_MSI) GetPropertyVector() (value uint16, err error) {
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
