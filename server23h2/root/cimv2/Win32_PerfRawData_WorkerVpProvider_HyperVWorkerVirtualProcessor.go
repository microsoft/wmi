// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor struct
type Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor struct {
	*Win32_PerfRawData

	//
	InitialAPICID uint64

	//
	InterceptDelayTimems uint64

	//
	InterceptsDelayed uint64

	//
	MPIDR uint64

	//
	TargetSubnode uint64
}

func NewWin32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessorEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetInitialAPICID sets the value of InitialAPICID for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) SetPropertyInitialAPICID(value uint64) (err error) {
	return instance.SetProperty("InitialAPICID", (value))
}

// GetInitialAPICID gets the value of InitialAPICID for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) GetPropertyInitialAPICID() (value uint64, err error) {
	retValue, err := instance.GetProperty("InitialAPICID")
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

// SetInterceptDelayTimems sets the value of InterceptDelayTimems for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) SetPropertyInterceptDelayTimems(value uint64) (err error) {
	return instance.SetProperty("InterceptDelayTimems", (value))
}

// GetInterceptDelayTimems gets the value of InterceptDelayTimems for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) GetPropertyInterceptDelayTimems() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterceptDelayTimems")
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

// SetInterceptsDelayed sets the value of InterceptsDelayed for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) SetPropertyInterceptsDelayed(value uint64) (err error) {
	return instance.SetProperty("InterceptsDelayed", (value))
}

// GetInterceptsDelayed gets the value of InterceptsDelayed for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) GetPropertyInterceptsDelayed() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterceptsDelayed")
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

// SetMPIDR sets the value of MPIDR for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) SetPropertyMPIDR(value uint64) (err error) {
	return instance.SetProperty("MPIDR", (value))
}

// GetMPIDR gets the value of MPIDR for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) GetPropertyMPIDR() (value uint64, err error) {
	retValue, err := instance.GetProperty("MPIDR")
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

// SetTargetSubnode sets the value of TargetSubnode for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) SetPropertyTargetSubnode(value uint64) (err error) {
	return instance.SetProperty("TargetSubnode", (value))
}

// GetTargetSubnode gets the value of TargetSubnode for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) GetPropertyTargetSubnode() (value uint64, err error) {
	retValue, err := instance.GetProperty("TargetSubnode")
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
