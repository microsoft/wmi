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

// Msvm_GpuPartition struct
type Msvm_GpuPartition struct {
	*CIM_LogicalDevice

	//
	CurrentCompute uint64

	//
	CurrentDecode uint64

	//
	CurrentEncode uint64

	//
	CurrentVRAM uint64

	//
	DeviceInstancePath string

	//
	PartitionId uint16

	//
	PartitionVfLuid string
}

func NewMsvm_GpuPartitionEx1(instance *cim.WmiInstance) (newInstance *Msvm_GpuPartition, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_GpuPartition{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewMsvm_GpuPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_GpuPartition, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_GpuPartition{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetCurrentCompute sets the value of CurrentCompute for the instance
func (instance *Msvm_GpuPartition) SetPropertyCurrentCompute(value uint64) (err error) {
	return instance.SetProperty("CurrentCompute", (value))
}

// GetCurrentCompute gets the value of CurrentCompute for the instance
func (instance *Msvm_GpuPartition) GetPropertyCurrentCompute() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentCompute")
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

// SetCurrentDecode sets the value of CurrentDecode for the instance
func (instance *Msvm_GpuPartition) SetPropertyCurrentDecode(value uint64) (err error) {
	return instance.SetProperty("CurrentDecode", (value))
}

// GetCurrentDecode gets the value of CurrentDecode for the instance
func (instance *Msvm_GpuPartition) GetPropertyCurrentDecode() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentDecode")
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

// SetCurrentEncode sets the value of CurrentEncode for the instance
func (instance *Msvm_GpuPartition) SetPropertyCurrentEncode(value uint64) (err error) {
	return instance.SetProperty("CurrentEncode", (value))
}

// GetCurrentEncode gets the value of CurrentEncode for the instance
func (instance *Msvm_GpuPartition) GetPropertyCurrentEncode() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentEncode")
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

// SetCurrentVRAM sets the value of CurrentVRAM for the instance
func (instance *Msvm_GpuPartition) SetPropertyCurrentVRAM(value uint64) (err error) {
	return instance.SetProperty("CurrentVRAM", (value))
}

// GetCurrentVRAM gets the value of CurrentVRAM for the instance
func (instance *Msvm_GpuPartition) GetPropertyCurrentVRAM() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentVRAM")
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

// SetDeviceInstancePath sets the value of DeviceInstancePath for the instance
func (instance *Msvm_GpuPartition) SetPropertyDeviceInstancePath(value string) (err error) {
	return instance.SetProperty("DeviceInstancePath", (value))
}

// GetDeviceInstancePath gets the value of DeviceInstancePath for the instance
func (instance *Msvm_GpuPartition) GetPropertyDeviceInstancePath() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceInstancePath")
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

// SetPartitionId sets the value of PartitionId for the instance
func (instance *Msvm_GpuPartition) SetPropertyPartitionId(value uint16) (err error) {
	return instance.SetProperty("PartitionId", (value))
}

// GetPartitionId gets the value of PartitionId for the instance
func (instance *Msvm_GpuPartition) GetPropertyPartitionId() (value uint16, err error) {
	retValue, err := instance.GetProperty("PartitionId")
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

// SetPartitionVfLuid sets the value of PartitionVfLuid for the instance
func (instance *Msvm_GpuPartition) SetPropertyPartitionVfLuid(value string) (err error) {
	return instance.SetProperty("PartitionVfLuid", (value))
}

// GetPartitionVfLuid gets the value of PartitionVfLuid for the instance
func (instance *Msvm_GpuPartition) GetPropertyPartitionVfLuid() (value string, err error) {
	retValue, err := instance.GetProperty("PartitionVfLuid")
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
