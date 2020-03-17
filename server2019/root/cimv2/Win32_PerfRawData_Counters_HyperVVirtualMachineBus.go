// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_HyperVVirtualMachineBus struct
type Win32_PerfRawData_Counters_HyperVVirtualMachineBus struct {
	Win32_PerfRawData

	//
	InterruptsReceivedPersec uint64

	//
	InterruptsSentPersec uint64

	//
	ThrottleEvents uint64
}

// SetInterruptsReceivedPersec sets the value of InterruptsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBus) SetPropertyInterruptsReceivedPersec(value uint64) (err error) {
	return instance.SetProperty("InterruptsReceivedPersec", value)
}

// GetInterruptsReceivedPersec gets the value of InterruptsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBus) GetPropertyInterruptsReceivedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterruptsReceivedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterruptsSentPersec sets the value of InterruptsSentPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBus) SetPropertyInterruptsSentPersec(value uint64) (err error) {
	return instance.SetProperty("InterruptsSentPersec", value)
}

// GetInterruptsSentPersec gets the value of InterruptsSentPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBus) GetPropertyInterruptsSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterruptsSentPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThrottleEvents sets the value of ThrottleEvents for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBus) SetPropertyThrottleEvents(value uint64) (err error) {
	return instance.SetProperty("ThrottleEvents", value)
}

// GetThrottleEvents gets the value of ThrottleEvents for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualMachineBus) GetPropertyThrottleEvents() (value uint64, err error) {
	retValue, err := instance.GetProperty("ThrottleEvents")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
