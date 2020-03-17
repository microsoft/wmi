// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_BalancerStats_HyperVDynamicMemoryBalancer struct
type Win32_PerfFormattedData_BalancerStats_HyperVDynamicMemoryBalancer struct {
	Win32_PerfFormattedData

	//
	AvailableMemory uint32

	//
	AvailableMemoryForBalancing uint32

	//
	AveragePressure uint32

	//
	SystemCurrentPressure uint32
}

// SetAvailableMemory sets the value of AvailableMemory for the instance
func (instance *Win32_PerfFormattedData_BalancerStats_HyperVDynamicMemoryBalancer) SetPropertyAvailableMemory(value uint32) (err error) {
	return instance.SetProperty("AvailableMemory", value)
}

// GetAvailableMemory gets the value of AvailableMemory for the instance
func (instance *Win32_PerfFormattedData_BalancerStats_HyperVDynamicMemoryBalancer) GetPropertyAvailableMemory() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvailableMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvailableMemoryForBalancing sets the value of AvailableMemoryForBalancing for the instance
func (instance *Win32_PerfFormattedData_BalancerStats_HyperVDynamicMemoryBalancer) SetPropertyAvailableMemoryForBalancing(value uint32) (err error) {
	return instance.SetProperty("AvailableMemoryForBalancing", value)
}

// GetAvailableMemoryForBalancing gets the value of AvailableMemoryForBalancing for the instance
func (instance *Win32_PerfFormattedData_BalancerStats_HyperVDynamicMemoryBalancer) GetPropertyAvailableMemoryForBalancing() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvailableMemoryForBalancing")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAveragePressure sets the value of AveragePressure for the instance
func (instance *Win32_PerfFormattedData_BalancerStats_HyperVDynamicMemoryBalancer) SetPropertyAveragePressure(value uint32) (err error) {
	return instance.SetProperty("AveragePressure", value)
}

// GetAveragePressure gets the value of AveragePressure for the instance
func (instance *Win32_PerfFormattedData_BalancerStats_HyperVDynamicMemoryBalancer) GetPropertyAveragePressure() (value uint32, err error) {
	retValue, err := instance.GetProperty("AveragePressure")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCurrentPressure sets the value of SystemCurrentPressure for the instance
func (instance *Win32_PerfFormattedData_BalancerStats_HyperVDynamicMemoryBalancer) SetPropertySystemCurrentPressure(value uint32) (err error) {
	return instance.SetProperty("SystemCurrentPressure", value)
}

// GetSystemCurrentPressure gets the value of SystemCurrentPressure for the instance
func (instance *Win32_PerfFormattedData_BalancerStats_HyperVDynamicMemoryBalancer) GetPropertySystemCurrentPressure() (value uint32, err error) {
	retValue, err := instance.GetProperty("SystemCurrentPressure")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
