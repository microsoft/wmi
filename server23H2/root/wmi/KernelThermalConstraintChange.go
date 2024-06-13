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

// KernelThermalConstraintChange struct
type KernelThermalConstraintChange struct {
	*WMIEvent

	//
	Active bool

	//
	InstanceName string

	//
	Processors uint64

	//
	ThermalConstraint uint32
}

func NewKernelThermalConstraintChangeEx1(instance *cim.WmiInstance) (newInstance *KernelThermalConstraintChange, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KernelThermalConstraintChange{
		WMIEvent: tmp,
	}
	return
}

func NewKernelThermalConstraintChangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KernelThermalConstraintChange, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KernelThermalConstraintChange{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *KernelThermalConstraintChange) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *KernelThermalConstraintChange) GetPropertyActive() (value bool, err error) {
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
func (instance *KernelThermalConstraintChange) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *KernelThermalConstraintChange) GetPropertyInstanceName() (value string, err error) {
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

// SetProcessors sets the value of Processors for the instance
func (instance *KernelThermalConstraintChange) SetPropertyProcessors(value uint64) (err error) {
	return instance.SetProperty("Processors", (value))
}

// GetProcessors gets the value of Processors for the instance
func (instance *KernelThermalConstraintChange) GetPropertyProcessors() (value uint64, err error) {
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

// SetThermalConstraint sets the value of ThermalConstraint for the instance
func (instance *KernelThermalConstraintChange) SetPropertyThermalConstraint(value uint32) (err error) {
	return instance.SetProperty("ThermalConstraint", (value))
}

// GetThermalConstraint gets the value of ThermalConstraint for the instance
func (instance *KernelThermalConstraintChange) GetPropertyThermalConstraint() (value uint32, err error) {
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
