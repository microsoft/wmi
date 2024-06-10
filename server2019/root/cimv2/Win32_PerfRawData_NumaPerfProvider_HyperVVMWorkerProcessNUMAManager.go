// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager struct
type Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager struct {
	*Win32_PerfRawData

	//
	InitialMemoryAssignedPerNodeMB uint64

	//
	MappedPhysicalNUMANode uint64

	//
	NUMASpanningAllowed uint32

	//
	VirtualProcessorCountPerNode uint64
}

func NewWin32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManagerEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetInitialMemoryAssignedPerNodeMB sets the value of InitialMemoryAssignedPerNodeMB for the instance
func (instance *Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager) SetPropertyInitialMemoryAssignedPerNodeMB(value uint64) (err error) {
	return instance.SetProperty("InitialMemoryAssignedPerNodeMB", (value))
}

// GetInitialMemoryAssignedPerNodeMB gets the value of InitialMemoryAssignedPerNodeMB for the instance
func (instance *Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager) GetPropertyInitialMemoryAssignedPerNodeMB() (value uint64, err error) {
	retValue, err := instance.GetProperty("InitialMemoryAssignedPerNodeMB")
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

// SetMappedPhysicalNUMANode sets the value of MappedPhysicalNUMANode for the instance
func (instance *Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager) SetPropertyMappedPhysicalNUMANode(value uint64) (err error) {
	return instance.SetProperty("MappedPhysicalNUMANode", (value))
}

// GetMappedPhysicalNUMANode gets the value of MappedPhysicalNUMANode for the instance
func (instance *Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager) GetPropertyMappedPhysicalNUMANode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MappedPhysicalNUMANode")
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

// SetNUMASpanningAllowed sets the value of NUMASpanningAllowed for the instance
func (instance *Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager) SetPropertyNUMASpanningAllowed(value uint32) (err error) {
	return instance.SetProperty("NUMASpanningAllowed", (value))
}

// GetNUMASpanningAllowed gets the value of NUMASpanningAllowed for the instance
func (instance *Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager) GetPropertyNUMASpanningAllowed() (value uint32, err error) {
	retValue, err := instance.GetProperty("NUMASpanningAllowed")
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

// SetVirtualProcessorCountPerNode sets the value of VirtualProcessorCountPerNode for the instance
func (instance *Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager) SetPropertyVirtualProcessorCountPerNode(value uint64) (err error) {
	return instance.SetProperty("VirtualProcessorCountPerNode", (value))
}

// GetVirtualProcessorCountPerNode gets the value of VirtualProcessorCountPerNode for the instance
func (instance *Win32_PerfRawData_NumaPerfProvider_HyperVVMWorkerProcessNUMAManager) GetPropertyVirtualProcessorCountPerNode() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualProcessorCountPerNode")
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
