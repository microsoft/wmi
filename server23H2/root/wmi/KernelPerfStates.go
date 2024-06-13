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

// KernelPerfStates struct
type KernelPerfStates struct {
	*MSKernelPpmClass

	//
	Active bool

	//
	BusyAdjThreshold uint8

	//
	Count uint32

	//
	CurrentState uint32

	//
	FeedbackHandler uint32

	//
	InstanceName string

	//
	LowestPerfState uint32

	//
	MaxFrequency uint32

	//
	MaxPerfState uint32

	//
	MinPerfState uint32

	//
	PolicyType uint8

	//
	PStateContext uint32

	//
	PStateHandler uint32

	//
	Reserved uint8

	//
	Reserved1 uint32

	//
	Reserved2 uint64

	//
	State []KernelPerfState

	//
	TargetProcessors uint64

	//
	ThermalConstraint uint32

	//
	TimerInterval uint32

	//
	TStateContext uint32

	//
	TStateHandler uint32

	//
	Type uint8
}

func NewKernelPerfStatesEx1(instance *cim.WmiInstance) (newInstance *KernelPerfStates, err error) {
	tmp, err := NewMSKernelPpmClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KernelPerfStates{
		MSKernelPpmClass: tmp,
	}
	return
}

func NewKernelPerfStatesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KernelPerfStates, err error) {
	tmp, err := NewMSKernelPpmClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KernelPerfStates{
		MSKernelPpmClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *KernelPerfStates) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *KernelPerfStates) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetBusyAdjThreshold sets the value of BusyAdjThreshold for the instance
func (instance *KernelPerfStates) SetPropertyBusyAdjThreshold(value uint8) (err error) {
	return instance.SetProperty("BusyAdjThreshold", (value))
}

// GetBusyAdjThreshold gets the value of BusyAdjThreshold for the instance
func (instance *KernelPerfStates) GetPropertyBusyAdjThreshold() (value uint8, err error) {
	retValue, err := instance.GetProperty("BusyAdjThreshold")
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

// SetCount sets the value of Count for the instance
func (instance *KernelPerfStates) SetPropertyCount(value uint32) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *KernelPerfStates) GetPropertyCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("Count")
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

// SetCurrentState sets the value of CurrentState for the instance
func (instance *KernelPerfStates) SetPropertyCurrentState(value uint32) (err error) {
	return instance.SetProperty("CurrentState", (value))
}

// GetCurrentState gets the value of CurrentState for the instance
func (instance *KernelPerfStates) GetPropertyCurrentState() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentState")
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

// SetFeedbackHandler sets the value of FeedbackHandler for the instance
func (instance *KernelPerfStates) SetPropertyFeedbackHandler(value uint32) (err error) {
	return instance.SetProperty("FeedbackHandler", (value))
}

// GetFeedbackHandler gets the value of FeedbackHandler for the instance
func (instance *KernelPerfStates) GetPropertyFeedbackHandler() (value uint32, err error) {
	retValue, err := instance.GetProperty("FeedbackHandler")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *KernelPerfStates) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *KernelPerfStates) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetLowestPerfState sets the value of LowestPerfState for the instance
func (instance *KernelPerfStates) SetPropertyLowestPerfState(value uint32) (err error) {
	return instance.SetProperty("LowestPerfState", (value))
}

// GetLowestPerfState gets the value of LowestPerfState for the instance
func (instance *KernelPerfStates) GetPropertyLowestPerfState() (value uint32, err error) {
	retValue, err := instance.GetProperty("LowestPerfState")
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

// SetMaxFrequency sets the value of MaxFrequency for the instance
func (instance *KernelPerfStates) SetPropertyMaxFrequency(value uint32) (err error) {
	return instance.SetProperty("MaxFrequency", (value))
}

// GetMaxFrequency gets the value of MaxFrequency for the instance
func (instance *KernelPerfStates) GetPropertyMaxFrequency() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxFrequency")
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

// SetMaxPerfState sets the value of MaxPerfState for the instance
func (instance *KernelPerfStates) SetPropertyMaxPerfState(value uint32) (err error) {
	return instance.SetProperty("MaxPerfState", (value))
}

// GetMaxPerfState gets the value of MaxPerfState for the instance
func (instance *KernelPerfStates) GetPropertyMaxPerfState() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxPerfState")
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

// SetMinPerfState sets the value of MinPerfState for the instance
func (instance *KernelPerfStates) SetPropertyMinPerfState(value uint32) (err error) {
	return instance.SetProperty("MinPerfState", (value))
}

// GetMinPerfState gets the value of MinPerfState for the instance
func (instance *KernelPerfStates) GetPropertyMinPerfState() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinPerfState")
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

// SetPolicyType sets the value of PolicyType for the instance
func (instance *KernelPerfStates) SetPropertyPolicyType(value uint8) (err error) {
	return instance.SetProperty("PolicyType", (value))
}

// GetPolicyType gets the value of PolicyType for the instance
func (instance *KernelPerfStates) GetPropertyPolicyType() (value uint8, err error) {
	retValue, err := instance.GetProperty("PolicyType")
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

// SetPStateContext sets the value of PStateContext for the instance
func (instance *KernelPerfStates) SetPropertyPStateContext(value uint32) (err error) {
	return instance.SetProperty("PStateContext", (value))
}

// GetPStateContext gets the value of PStateContext for the instance
func (instance *KernelPerfStates) GetPropertyPStateContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("PStateContext")
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

// SetPStateHandler sets the value of PStateHandler for the instance
func (instance *KernelPerfStates) SetPropertyPStateHandler(value uint32) (err error) {
	return instance.SetProperty("PStateHandler", (value))
}

// GetPStateHandler gets the value of PStateHandler for the instance
func (instance *KernelPerfStates) GetPropertyPStateHandler() (value uint32, err error) {
	retValue, err := instance.GetProperty("PStateHandler")
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
func (instance *KernelPerfStates) SetPropertyReserved(value uint8) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *KernelPerfStates) GetPropertyReserved() (value uint8, err error) {
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
func (instance *KernelPerfStates) SetPropertyReserved1(value uint32) (err error) {
	return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *KernelPerfStates) GetPropertyReserved1() (value uint32, err error) {
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

// SetReserved2 sets the value of Reserved2 for the instance
func (instance *KernelPerfStates) SetPropertyReserved2(value uint64) (err error) {
	return instance.SetProperty("Reserved2", (value))
}

// GetReserved2 gets the value of Reserved2 for the instance
func (instance *KernelPerfStates) GetPropertyReserved2() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reserved2")
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

// SetState sets the value of State for the instance
func (instance *KernelPerfStates) SetPropertyState(value []KernelPerfState) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *KernelPerfStates) GetPropertyState() (value []KernelPerfState, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(KernelPerfState)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " KernelPerfState is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, KernelPerfState(valuetmp))
	}

	return
}

// SetTargetProcessors sets the value of TargetProcessors for the instance
func (instance *KernelPerfStates) SetPropertyTargetProcessors(value uint64) (err error) {
	return instance.SetProperty("TargetProcessors", (value))
}

// GetTargetProcessors gets the value of TargetProcessors for the instance
func (instance *KernelPerfStates) GetPropertyTargetProcessors() (value uint64, err error) {
	retValue, err := instance.GetProperty("TargetProcessors")
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

// SetThermalConstraint sets the value of ThermalConstraint for the instance
func (instance *KernelPerfStates) SetPropertyThermalConstraint(value uint32) (err error) {
	return instance.SetProperty("ThermalConstraint", (value))
}

// GetThermalConstraint gets the value of ThermalConstraint for the instance
func (instance *KernelPerfStates) GetPropertyThermalConstraint() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThermalConstraint")
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

// SetTimerInterval sets the value of TimerInterval for the instance
func (instance *KernelPerfStates) SetPropertyTimerInterval(value uint32) (err error) {
	return instance.SetProperty("TimerInterval", (value))
}

// GetTimerInterval gets the value of TimerInterval for the instance
func (instance *KernelPerfStates) GetPropertyTimerInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimerInterval")
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

// SetTStateContext sets the value of TStateContext for the instance
func (instance *KernelPerfStates) SetPropertyTStateContext(value uint32) (err error) {
	return instance.SetProperty("TStateContext", (value))
}

// GetTStateContext gets the value of TStateContext for the instance
func (instance *KernelPerfStates) GetPropertyTStateContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("TStateContext")
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

// SetTStateHandler sets the value of TStateHandler for the instance
func (instance *KernelPerfStates) SetPropertyTStateHandler(value uint32) (err error) {
	return instance.SetProperty("TStateHandler", (value))
}

// GetTStateHandler gets the value of TStateHandler for the instance
func (instance *KernelPerfStates) GetPropertyTStateHandler() (value uint32, err error) {
	retValue, err := instance.GetProperty("TStateHandler")
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

// SetType sets the value of Type for the instance
func (instance *KernelPerfStates) SetPropertyType(value uint8) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *KernelPerfStates) GetPropertyType() (value uint8, err error) {
	retValue, err := instance.GetProperty("Type")
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
