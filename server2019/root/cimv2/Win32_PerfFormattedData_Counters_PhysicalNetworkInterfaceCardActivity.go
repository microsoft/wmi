// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_PhysicalNetworkInterfaceCardActivity struct
type Win32_PerfFormattedData_Counters_PhysicalNetworkInterfaceCardActivity struct {
	Win32_PerfFormattedData

	//
	DevicePowerState uint32

	//
	LowPowerTransitionsLifetime uint32

	//
	PercentTimeSuspendedInstantaneous uint64

	//
	PercentTimeSuspendedLifetime uint64
}

// SetDevicePowerState sets the value of DevicePowerState for the instance
func (instance *Win32_PerfFormattedData_Counters_PhysicalNetworkInterfaceCardActivity) SetPropertyDevicePowerState(value uint32) (err error) {
	return instance.SetProperty("DevicePowerState", value)
}

// GetDevicePowerState gets the value of DevicePowerState for the instance
func (instance *Win32_PerfFormattedData_Counters_PhysicalNetworkInterfaceCardActivity) GetPropertyDevicePowerState() (value uint32, err error) {
	retValue, err := instance.GetProperty("DevicePowerState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLowPowerTransitionsLifetime sets the value of LowPowerTransitionsLifetime for the instance
func (instance *Win32_PerfFormattedData_Counters_PhysicalNetworkInterfaceCardActivity) SetPropertyLowPowerTransitionsLifetime(value uint32) (err error) {
	return instance.SetProperty("LowPowerTransitionsLifetime", value)
}

// GetLowPowerTransitionsLifetime gets the value of LowPowerTransitionsLifetime for the instance
func (instance *Win32_PerfFormattedData_Counters_PhysicalNetworkInterfaceCardActivity) GetPropertyLowPowerTransitionsLifetime() (value uint32, err error) {
	retValue, err := instance.GetProperty("LowPowerTransitionsLifetime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTimeSuspendedInstantaneous sets the value of PercentTimeSuspendedInstantaneous for the instance
func (instance *Win32_PerfFormattedData_Counters_PhysicalNetworkInterfaceCardActivity) SetPropertyPercentTimeSuspendedInstantaneous(value uint64) (err error) {
	return instance.SetProperty("PercentTimeSuspendedInstantaneous", value)
}

// GetPercentTimeSuspendedInstantaneous gets the value of PercentTimeSuspendedInstantaneous for the instance
func (instance *Win32_PerfFormattedData_Counters_PhysicalNetworkInterfaceCardActivity) GetPropertyPercentTimeSuspendedInstantaneous() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentTimeSuspendedInstantaneous")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTimeSuspendedLifetime sets the value of PercentTimeSuspendedLifetime for the instance
func (instance *Win32_PerfFormattedData_Counters_PhysicalNetworkInterfaceCardActivity) SetPropertyPercentTimeSuspendedLifetime(value uint64) (err error) {
	return instance.SetProperty("PercentTimeSuspendedLifetime", value)
}

// GetPercentTimeSuspendedLifetime gets the value of PercentTimeSuspendedLifetime for the instance
func (instance *Win32_PerfFormattedData_Counters_PhysicalNetworkInterfaceCardActivity) GetPropertyPercentTimeSuspendedLifetime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentTimeSuspendedLifetime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
