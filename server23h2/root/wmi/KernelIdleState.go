// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// KernelIdleState struct
type KernelIdleState struct {
	*cim.WmiInstance

	//
	Context uint32

	//
	DemotePercent uint8

	//
	IdleHandler uint32

	//
	Latency uint32

	//
	Power uint32

	//
	PromotePercent uint8

	//
	Reserved uint8

	//
	Reserved1 uint32

	//
	StateFlags uint32

	//
	StateType uint8

	//
	TimeCheck uint32
}

func NewKernelIdleStateEx1(instance *cim.WmiInstance) (newInstance *KernelIdleState, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &KernelIdleState{
		WmiInstance: tmp,
	}
	return
}

func NewKernelIdleStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KernelIdleState, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KernelIdleState{
		WmiInstance: tmp,
	}
	return
}

// SetContext sets the value of Context for the instance
func (instance *KernelIdleState) SetPropertyContext(value uint32) (err error) {
	return instance.SetProperty("Context", (value))
}

// GetContext gets the value of Context for the instance
func (instance *KernelIdleState) GetPropertyContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("Context")
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

// SetDemotePercent sets the value of DemotePercent for the instance
func (instance *KernelIdleState) SetPropertyDemotePercent(value uint8) (err error) {
	return instance.SetProperty("DemotePercent", (value))
}

// GetDemotePercent gets the value of DemotePercent for the instance
func (instance *KernelIdleState) GetPropertyDemotePercent() (value uint8, err error) {
	retValue, err := instance.GetProperty("DemotePercent")
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

// SetIdleHandler sets the value of IdleHandler for the instance
func (instance *KernelIdleState) SetPropertyIdleHandler(value uint32) (err error) {
	return instance.SetProperty("IdleHandler", (value))
}

// GetIdleHandler gets the value of IdleHandler for the instance
func (instance *KernelIdleState) GetPropertyIdleHandler() (value uint32, err error) {
	retValue, err := instance.GetProperty("IdleHandler")
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

// SetLatency sets the value of Latency for the instance
func (instance *KernelIdleState) SetPropertyLatency(value uint32) (err error) {
	return instance.SetProperty("Latency", (value))
}

// GetLatency gets the value of Latency for the instance
func (instance *KernelIdleState) GetPropertyLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("Latency")
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

// SetPower sets the value of Power for the instance
func (instance *KernelIdleState) SetPropertyPower(value uint32) (err error) {
	return instance.SetProperty("Power", (value))
}

// GetPower gets the value of Power for the instance
func (instance *KernelIdleState) GetPropertyPower() (value uint32, err error) {
	retValue, err := instance.GetProperty("Power")
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

// SetPromotePercent sets the value of PromotePercent for the instance
func (instance *KernelIdleState) SetPropertyPromotePercent(value uint8) (err error) {
	return instance.SetProperty("PromotePercent", (value))
}

// GetPromotePercent gets the value of PromotePercent for the instance
func (instance *KernelIdleState) GetPropertyPromotePercent() (value uint8, err error) {
	retValue, err := instance.GetProperty("PromotePercent")
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

// SetReserved sets the value of Reserved for the instance
func (instance *KernelIdleState) SetPropertyReserved(value uint8) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *KernelIdleState) GetPropertyReserved() (value uint8, err error) {
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

// SetReserved1 sets the value of Reserved1 for the instance
func (instance *KernelIdleState) SetPropertyReserved1(value uint32) (err error) {
	return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *KernelIdleState) GetPropertyReserved1() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved1")
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

// SetStateFlags sets the value of StateFlags for the instance
func (instance *KernelIdleState) SetPropertyStateFlags(value uint32) (err error) {
	return instance.SetProperty("StateFlags", (value))
}

// GetStateFlags gets the value of StateFlags for the instance
func (instance *KernelIdleState) GetPropertyStateFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("StateFlags")
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

// SetStateType sets the value of StateType for the instance
func (instance *KernelIdleState) SetPropertyStateType(value uint8) (err error) {
	return instance.SetProperty("StateType", (value))
}

// GetStateType gets the value of StateType for the instance
func (instance *KernelIdleState) GetPropertyStateType() (value uint8, err error) {
	retValue, err := instance.GetProperty("StateType")
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

// SetTimeCheck sets the value of TimeCheck for the instance
func (instance *KernelIdleState) SetPropertyTimeCheck(value uint32) (err error) {
	return instance.SetProperty("TimeCheck", (value))
}

// GetTimeCheck gets the value of TimeCheck for the instance
func (instance *KernelIdleState) GetPropertyTimeCheck() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimeCheck")
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
