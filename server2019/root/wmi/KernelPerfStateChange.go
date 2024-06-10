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

// KernelPerfStateChange struct
type KernelPerfStateChange struct {
	*WMIEvent

	//
	Active bool

	//
	InstanceName string

	//
	Latency uint32

	//
	Processor uint32

	//
	Speed uint32

	//
	State uint32

	//
	Status uint32
}

func NewKernelPerfStateChangeEx1(instance *cim.WmiInstance) (newInstance *KernelPerfStateChange, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KernelPerfStateChange{
		WMIEvent: tmp,
	}
	return
}

func NewKernelPerfStateChangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KernelPerfStateChange, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KernelPerfStateChange{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *KernelPerfStateChange) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *KernelPerfStateChange) GetPropertyActive() (value bool, err error) {
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
func (instance *KernelPerfStateChange) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *KernelPerfStateChange) GetPropertyInstanceName() (value string, err error) {
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
func (instance *KernelPerfStateChange) SetPropertyLatency(value uint32) (err error) {
	return instance.SetProperty("Latency", (value))
}

// GetLatency gets the value of Latency for the instance
func (instance *KernelPerfStateChange) GetPropertyLatency() (value uint32, err error) {
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

// SetProcessor sets the value of Processor for the instance
func (instance *KernelPerfStateChange) SetPropertyProcessor(value uint32) (err error) {
	return instance.SetProperty("Processor", (value))
}

// GetProcessor gets the value of Processor for the instance
func (instance *KernelPerfStateChange) GetPropertyProcessor() (value uint32, err error) {
	retValue, err := instance.GetProperty("Processor")
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

// SetSpeed sets the value of Speed for the instance
func (instance *KernelPerfStateChange) SetPropertySpeed(value uint32) (err error) {
	return instance.SetProperty("Speed", (value))
}

// GetSpeed gets the value of Speed for the instance
func (instance *KernelPerfStateChange) GetPropertySpeed() (value uint32, err error) {
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
func (instance *KernelPerfStateChange) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *KernelPerfStateChange) GetPropertyState() (value uint32, err error) {
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

// SetStatus sets the value of Status for the instance
func (instance *KernelPerfStateChange) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *KernelPerfStateChange) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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
