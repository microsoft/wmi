// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_PowerMeterCounter_EnergyMeter struct
type Win32_PerfRawData_PowerMeterCounter_EnergyMeter struct {
	Win32_PerfRawData

	//
	Energy uint64

	//
	Power uint64

	//
	Power_Base uint32

	//
	Time uint64
}

// SetEnergy sets the value of Energy for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_EnergyMeter) SetPropertyEnergy(value uint64) (err error) {
	return instance.SetProperty("Energy", value)
}

// GetEnergy gets the value of Energy for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_EnergyMeter) GetPropertyEnergy() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_PowerMeterCounter_EnergyMeter) SetPropertyPower(value uint64) (err error) {
	return instance.SetProperty("Power", value)
}

// GetPower gets the value of Power for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_EnergyMeter) GetPropertyPower() (value uint64, err error) {
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

// SetPower_Base sets the value of Power_Base for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_EnergyMeter) SetPropertyPower_Base(value uint32) (err error) {
	return instance.SetProperty("Power_Base", value)
}

// GetPower_Base gets the value of Power_Base for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_EnergyMeter) GetPropertyPower_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Power_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTime sets the value of Time for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_EnergyMeter) SetPropertyTime(value uint64) (err error) {
	return instance.SetProperty("Time", value)
}

// GetTime gets the value of Time for the instance
func (instance *Win32_PerfRawData_PowerMeterCounter_EnergyMeter) GetPropertyTime() (value uint64, err error) {
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
