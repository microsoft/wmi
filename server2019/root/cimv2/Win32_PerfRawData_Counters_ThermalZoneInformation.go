// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_ThermalZoneInformation struct
type Win32_PerfRawData_Counters_ThermalZoneInformation struct {
	Win32_PerfRawData

	//
	HighPrecisionTemperature uint32

	//
	PercentPassiveLimit uint32

	//
	Temperature uint32

	//
	ThrottleReasons uint32
}

// SetHighPrecisionTemperature sets the value of HighPrecisionTemperature for the instance
func (instance *Win32_PerfRawData_Counters_ThermalZoneInformation) SetPropertyHighPrecisionTemperature(value uint32) (err error) {
	return instance.SetProperty("HighPrecisionTemperature", value)
}

// GetHighPrecisionTemperature gets the value of HighPrecisionTemperature for the instance
func (instance *Win32_PerfRawData_Counters_ThermalZoneInformation) GetPropertyHighPrecisionTemperature() (value uint32, err error) {
	retValue, err := instance.GetProperty("HighPrecisionTemperature")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentPassiveLimit sets the value of PercentPassiveLimit for the instance
func (instance *Win32_PerfRawData_Counters_ThermalZoneInformation) SetPropertyPercentPassiveLimit(value uint32) (err error) {
	return instance.SetProperty("PercentPassiveLimit", value)
}

// GetPercentPassiveLimit gets the value of PercentPassiveLimit for the instance
func (instance *Win32_PerfRawData_Counters_ThermalZoneInformation) GetPropertyPercentPassiveLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentPassiveLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemperature sets the value of Temperature for the instance
func (instance *Win32_PerfRawData_Counters_ThermalZoneInformation) SetPropertyTemperature(value uint32) (err error) {
	return instance.SetProperty("Temperature", value)
}

// GetTemperature gets the value of Temperature for the instance
func (instance *Win32_PerfRawData_Counters_ThermalZoneInformation) GetPropertyTemperature() (value uint32, err error) {
	retValue, err := instance.GetProperty("Temperature")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThrottleReasons sets the value of ThrottleReasons for the instance
func (instance *Win32_PerfRawData_Counters_ThermalZoneInformation) SetPropertyThrottleReasons(value uint32) (err error) {
	return instance.SetProperty("ThrottleReasons", value)
}

// GetThrottleReasons gets the value of ThrottleReasons for the instance
func (instance *Win32_PerfRawData_Counters_ThermalZoneInformation) GetPropertyThrottleReasons() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThrottleReasons")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
