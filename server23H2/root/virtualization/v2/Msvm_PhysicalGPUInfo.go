// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_PhysicalGPUInfo struct
type Msvm_PhysicalGPUInfo struct {
	*CIM_ManagedElement

	//
	AvailableVideoMemory uint64

	//
	ID string

	//
	Name string

	//
	TotalVideoMemory uint64
}

func NewMsvm_PhysicalGPUInfoEx1(instance *cim.WmiInstance) (newInstance *Msvm_PhysicalGPUInfo, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_PhysicalGPUInfo{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMsvm_PhysicalGPUInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_PhysicalGPUInfo, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_PhysicalGPUInfo{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetAvailableVideoMemory sets the value of AvailableVideoMemory for the instance
func (instance *Msvm_PhysicalGPUInfo) SetPropertyAvailableVideoMemory(value uint64) (err error) {
	return instance.SetProperty("AvailableVideoMemory", (value))
}

// GetAvailableVideoMemory gets the value of AvailableVideoMemory for the instance
func (instance *Msvm_PhysicalGPUInfo) GetPropertyAvailableVideoMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableVideoMemory")
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

// SetID sets the value of ID for the instance
func (instance *Msvm_PhysicalGPUInfo) SetPropertyID(value string) (err error) {
	return instance.SetProperty("ID", (value))
}

// GetID gets the value of ID for the instance
func (instance *Msvm_PhysicalGPUInfo) GetPropertyID() (value string, err error) {
	retValue, err := instance.GetProperty("ID")
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

// SetName sets the value of Name for the instance
func (instance *Msvm_PhysicalGPUInfo) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Msvm_PhysicalGPUInfo) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetTotalVideoMemory sets the value of TotalVideoMemory for the instance
func (instance *Msvm_PhysicalGPUInfo) SetPropertyTotalVideoMemory(value uint64) (err error) {
	return instance.SetProperty("TotalVideoMemory", (value))
}

// GetTotalVideoMemory gets the value of TotalVideoMemory for the instance
func (instance *Msvm_PhysicalGPUInfo) GetPropertyTotalVideoMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalVideoMemory")
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
