// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_VmmsVirtualMachineStats_HyperVVirtualMachineHealthSummary struct
type Win32_PerfRawData_VmmsVirtualMachineStats_HyperVVirtualMachineHealthSummary struct {
	Win32_PerfRawData

	//
	HealthCritical uint32

	//
	HealthOk uint32
}

// SetHealthCritical sets the value of HealthCritical for the instance
func (instance *Win32_PerfRawData_VmmsVirtualMachineStats_HyperVVirtualMachineHealthSummary) SetPropertyHealthCritical(value uint32) (err error) {
	return instance.SetProperty("HealthCritical", value)
}

// GetHealthCritical gets the value of HealthCritical for the instance
func (instance *Win32_PerfRawData_VmmsVirtualMachineStats_HyperVVirtualMachineHealthSummary) GetPropertyHealthCritical() (value uint32, err error) {
	retValue, err := instance.GetProperty("HealthCritical")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthOk sets the value of HealthOk for the instance
func (instance *Win32_PerfRawData_VmmsVirtualMachineStats_HyperVVirtualMachineHealthSummary) SetPropertyHealthOk(value uint32) (err error) {
	return instance.SetProperty("HealthOk", value)
}

// GetHealthOk gets the value of HealthOk for the instance
func (instance *Win32_PerfRawData_VmmsVirtualMachineStats_HyperVVirtualMachineHealthSummary) GetPropertyHealthOk() (value uint32, err error) {
	retValue, err := instance.GetProperty("HealthOk")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
