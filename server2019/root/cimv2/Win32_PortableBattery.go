// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PortableBattery struct
type Win32_PortableBattery struct {
	CIM_Battery

	//
	CapacityMultiplier uint16

	//
	Location string

	//
	ManufactureDate string

	//
	Manufacturer string

	//
	MaxBatteryError uint16
}

// SetCapacityMultiplier sets the value of CapacityMultiplier for the instance
func (instance *Win32_PortableBattery) SetPropertyCapacityMultiplier(value uint16) (err error) {
	return instance.SetProperty("CapacityMultiplier", value)
}

// GetCapacityMultiplier gets the value of CapacityMultiplier for the instance
func (instance *Win32_PortableBattery) GetPropertyCapacityMultiplier() (value uint16, err error) {
	retValue, err := instance.GetProperty("CapacityMultiplier")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocation sets the value of Location for the instance
func (instance *Win32_PortableBattery) SetPropertyLocation(value string) (err error) {
	return instance.SetProperty("Location", value)
}

// GetLocation gets the value of Location for the instance
func (instance *Win32_PortableBattery) GetPropertyLocation() (value string, err error) {
	retValue, err := instance.GetProperty("Location")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufactureDate sets the value of ManufactureDate for the instance
func (instance *Win32_PortableBattery) SetPropertyManufactureDate(value string) (err error) {
	return instance.SetProperty("ManufactureDate", value)
}

// GetManufactureDate gets the value of ManufactureDate for the instance
func (instance *Win32_PortableBattery) GetPropertyManufactureDate() (value string, err error) {
	retValue, err := instance.GetProperty("ManufactureDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *Win32_PortableBattery) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *Win32_PortableBattery) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxBatteryError sets the value of MaxBatteryError for the instance
func (instance *Win32_PortableBattery) SetPropertyMaxBatteryError(value uint16) (err error) {
	return instance.SetProperty("MaxBatteryError", value)
}

// GetMaxBatteryError gets the value of MaxBatteryError for the instance
func (instance *Win32_PortableBattery) GetPropertyMaxBatteryError() (value uint16, err error) {
	retValue, err := instance.GetProperty("MaxBatteryError")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
