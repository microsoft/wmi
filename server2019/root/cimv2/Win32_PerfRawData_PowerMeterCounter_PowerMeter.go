// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_PowerMeterCounter_PowerMeter struct
type Win32_PerfRawData_PowerMeterCounter_PowerMeter struct {
	Win32_PerfRawData

	//
	Power uint32

	//
	PowerBudget uint32
}

// SetPower sets the value of Power for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_PowerMeter) SetPropertyPower(value uint32) (err error) {
	return instance.SetProperty("Power", value)
}

// GetPower gets the value of Power for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_PowerMeter) GetPropertyPower() (value uint32, err error) {
	retValue, err := instance.GetProperty("Power")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerBudget sets the value of PowerBudget for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_PowerMeter) SetPropertyPowerBudget(value uint32) (err error) {
	return instance.SetProperty("PowerBudget", value)
}

// GetPowerBudget gets the value of PowerBudget for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_PowerMeter) GetPropertyPowerBudget() (value uint32, err error) {
	retValue, err := instance.GetProperty("PowerBudget")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
