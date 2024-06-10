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

// KernelPerfStateDomainChange struct
type KernelPerfStateDomainChange struct {
	*WMIEvent

	//
	Active bool

	//
	InstanceName string

	//
	Latency uint32

	//
	Processors uint64

	//
	Speed uint32

	//
	State uint32
}

func NewKernelPerfStateDomainChangeEx1(instance *cim.WmiInstance) (newInstance *KernelPerfStateDomainChange, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KernelPerfStateDomainChange{
		WMIEvent: tmp,
	}
	return
}

func NewKernelPerfStateDomainChangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KernelPerfStateDomainChange, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KernelPerfStateDomainChange{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *KernelPerfStateDomainChange) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *KernelPerfStateDomainChange) GetPropertyActive() (value bool, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *KernelPerfStateDomainChange) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *KernelPerfStateDomainChange) GetPropertyInstanceName() (value string, err error) {
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

// SetLatency sets the value of Latency for the instance
func (instance *KernelPerfStateDomainChange) SetPropertyLatency(value uint32) (err error) {
	return instance.SetProperty("Latency", (value))
}

// GetLatency gets the value of Latency for the instance
func (instance *KernelPerfStateDomainChange) GetPropertyLatency() (value uint32, err error) {
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

// SetProcessors sets the value of Processors for the instance
func (instance *KernelPerfStateDomainChange) SetPropertyProcessors(value uint64) (err error) {
	return instance.SetProperty("Processors", (value))
}

// GetProcessors gets the value of Processors for the instance
func (instance *KernelPerfStateDomainChange) GetPropertyProcessors() (value uint64, err error) {
	retValue, err := instance.GetProperty("Processors")
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

// SetSpeed sets the value of Speed for the instance
func (instance *KernelPerfStateDomainChange) SetPropertySpeed(value uint32) (err error) {
	return instance.SetProperty("Speed", (value))
}

// GetSpeed gets the value of Speed for the instance
func (instance *KernelPerfStateDomainChange) GetPropertySpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("Speed")
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
func (instance *KernelPerfStateDomainChange) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *KernelPerfStateDomainChange) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
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
