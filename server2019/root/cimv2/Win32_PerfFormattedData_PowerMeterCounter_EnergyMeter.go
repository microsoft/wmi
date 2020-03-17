// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_PowerMeterCounter_EnergyMeter struct
type Win32_PerfFormattedData_PowerMeterCounter_EnergyMeter struct {
	Win32_PerfFormattedData

	//
	Energy uint64

	//
	Power uint64

	//
	Time uint64
}

// SetEnergy sets the value of Energy for the instance
func (instance *Win32_PerfFormattedData_PowerMeterCounter_EnergyMeter) SetPropertyEnergy(value uint64) (err error) {
	return instance.SetProperty("Energy", value)
}

// GetEnergy gets the value of Energy for the instance
func (instance *Win32_PerfFormattedData_PowerMeterCounter_EnergyMeter) GetPropertyEnergy() (value uint64, err error) {
	retValue, err := instance.GetProperty("Energy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPower sets the value of Power for the instance
func (instance *Win32_PerfFormattedData_PowerMeterCounter_EnergyMeter) SetPropertyPower(value uint64) (err error) {
	return instance.SetProperty("Power", value)
}

// GetPower gets the value of Power for the instance
func (instance *Win32_PerfFormattedData_PowerMeterCounter_EnergyMeter) GetPropertyPower() (value uint64, err error) {
	retValue, err := instance.GetProperty("Power")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTime sets the value of Time for the instance
func (instance *Win32_PerfFormattedData_PowerMeterCounter_EnergyMeter) SetPropertyTime(value uint64) (err error) {
	return instance.SetProperty("Time", value)
}

// GetTime gets the value of Time for the instance
func (instance *Win32_PerfFormattedData_PowerMeterCounter_EnergyMeter) GetPropertyTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Time")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
