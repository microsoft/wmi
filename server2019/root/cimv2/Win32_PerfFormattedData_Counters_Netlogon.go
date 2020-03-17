// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_Netlogon struct
type Win32_PerfFormattedData_Counters_Netlogon struct {
	Win32_PerfFormattedData

	//
	AverageSemaphoreHoldTime uint32

	//
	LastAuthenticationTime uint32

	//
	SemaphoreAcquires uint64

	//
	SemaphoreHolders uint32

	//
	SemaphoreTimeouts uint64

	//
	SemaphoreWaiters uint32
}

// SetAverageSemaphoreHoldTime sets the value of AverageSemaphoreHoldTime for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) SetPropertyAverageSemaphoreHoldTime(value uint32) (err error) {
	return instance.SetProperty("AverageSemaphoreHoldTime", value)
}

// GetAverageSemaphoreHoldTime gets the value of AverageSemaphoreHoldTime for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) GetPropertyAverageSemaphoreHoldTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("AverageSemaphoreHoldTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastAuthenticationTime sets the value of LastAuthenticationTime for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) SetPropertyLastAuthenticationTime(value uint32) (err error) {
	return instance.SetProperty("LastAuthenticationTime", value)
}

// GetLastAuthenticationTime gets the value of LastAuthenticationTime for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) GetPropertyLastAuthenticationTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastAuthenticationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSemaphoreAcquires sets the value of SemaphoreAcquires for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) SetPropertySemaphoreAcquires(value uint64) (err error) {
	return instance.SetProperty("SemaphoreAcquires", value)
}

// GetSemaphoreAcquires gets the value of SemaphoreAcquires for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) GetPropertySemaphoreAcquires() (value uint64, err error) {
	retValue, err := instance.GetProperty("SemaphoreAcquires")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSemaphoreHolders sets the value of SemaphoreHolders for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) SetPropertySemaphoreHolders(value uint32) (err error) {
	return instance.SetProperty("SemaphoreHolders", value)
}

// GetSemaphoreHolders gets the value of SemaphoreHolders for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) GetPropertySemaphoreHolders() (value uint32, err error) {
	retValue, err := instance.GetProperty("SemaphoreHolders")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSemaphoreTimeouts sets the value of SemaphoreTimeouts for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) SetPropertySemaphoreTimeouts(value uint64) (err error) {
	return instance.SetProperty("SemaphoreTimeouts", value)
}

// GetSemaphoreTimeouts gets the value of SemaphoreTimeouts for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) GetPropertySemaphoreTimeouts() (value uint64, err error) {
	retValue, err := instance.GetProperty("SemaphoreTimeouts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSemaphoreWaiters sets the value of SemaphoreWaiters for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) SetPropertySemaphoreWaiters(value uint32) (err error) {
	return instance.SetProperty("SemaphoreWaiters", value)
}

// GetSemaphoreWaiters gets the value of SemaphoreWaiters for the instance
func (instance *Win32_PerfFormattedData_Counters_Netlogon) GetPropertySemaphoreWaiters() (value uint32, err error) {
	retValue, err := instance.GetProperty("SemaphoreWaiters")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
