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

// KernelIdleStates struct
type KernelIdleStates struct {
	*MSKernelPpmClass

	//
	Active bool

	//
	Count uint32

	//
	InstanceName string

	//
	OldState uint32

	//
	State []KernelIdleState

	//
	TargetProcessors uint64

	//
	TargetState uint32

	//
	Type uint32
}

func NewKernelIdleStatesEx1(instance *cim.WmiInstance) (newInstance *KernelIdleStates, err error) {
	tmp, err := NewMSKernelPpmClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KernelIdleStates{
		MSKernelPpmClass: tmp,
	}
	return
}

func NewKernelIdleStatesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KernelIdleStates, err error) {
	tmp, err := NewMSKernelPpmClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KernelIdleStates{
		MSKernelPpmClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *KernelIdleStates) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *KernelIdleStates) GetPropertyActive() (value bool, err error) {
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

// SetCount sets the value of Count for the instance
func (instance *KernelIdleStates) SetPropertyCount(value uint32) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *KernelIdleStates) GetPropertyCount() (value uint32, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *KernelIdleStates) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *KernelIdleStates) GetPropertyInstanceName() (value string, err error) {
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

// SetOldState sets the value of OldState for the instance
func (instance *KernelIdleStates) SetPropertyOldState(value uint32) (err error) {
	return instance.SetProperty("OldState", (value))
}

// GetOldState gets the value of OldState for the instance
func (instance *KernelIdleStates) GetPropertyOldState() (value uint32, err error) {
	retValue, err := instance.GetProperty("OldState")
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

// SetState sets the value of State for the instance
func (instance *KernelIdleStates) SetPropertyState(value []KernelIdleState) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *KernelIdleStates) GetPropertyState() (value []KernelIdleState, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(KernelIdleState)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " KernelIdleState is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, KernelIdleState(valuetmp))
	}

	return
}

// SetTargetProcessors sets the value of TargetProcessors for the instance
func (instance *KernelIdleStates) SetPropertyTargetProcessors(value uint64) (err error) {
	return instance.SetProperty("TargetProcessors", (value))
}

// GetTargetProcessors gets the value of TargetProcessors for the instance
func (instance *KernelIdleStates) GetPropertyTargetProcessors() (value uint64, err error) {
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

// SetTargetState sets the value of TargetState for the instance
func (instance *KernelIdleStates) SetPropertyTargetState(value uint32) (err error) {
	return instance.SetProperty("TargetState", (value))
}

// GetTargetState gets the value of TargetState for the instance
func (instance *KernelIdleStates) GetPropertyTargetState() (value uint32, err error) {
	retValue, err := instance.GetProperty("TargetState")
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
func (instance *KernelIdleStates) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *KernelIdleStates) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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
