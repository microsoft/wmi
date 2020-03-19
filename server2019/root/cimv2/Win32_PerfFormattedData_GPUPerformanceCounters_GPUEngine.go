// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine struct
type Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine struct {
	*Win32_PerfFormattedData

	//
	RunningTime uint64

	//
	UtilizationPercentage uint64
}

func NewWin32_PerfFormattedData_GPUPerformanceCounters_GPUEngineEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_GPUPerformanceCounters_GPUEngineEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetRunningTime sets the value of RunningTime for the instance
func (instance *Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine) SetPropertyRunningTime(value uint64) (err error) {
	return instance.SetProperty("RunningTime", value)
}

// GetRunningTime gets the value of RunningTime for the instance
func (instance *Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine) GetPropertyRunningTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("RunningTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUtilizationPercentage sets the value of UtilizationPercentage for the instance
func (instance *Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine) SetPropertyUtilizationPercentage(value uint64) (err error) {
	return instance.SetProperty("UtilizationPercentage", value)
}

// GetUtilizationPercentage gets the value of UtilizationPercentage for the instance
func (instance *Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine) GetPropertyUtilizationPercentage() (value uint64, err error) {
	retValue, err := instance.GetProperty("UtilizationPercentage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
