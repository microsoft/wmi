// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// SDDC_Memory struct
type SDDC_Memory struct {
	cim.WmiInstance

	//
	Manufacturer string

	//
	Model string

	//
	SerialNumber string

	//
	SizeInBytes uint64

	//
	SpeedInMHz uint32
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *SDDC_Memory) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *SDDC_Memory) GetPropertyManufacturer() (value string, err error) {
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

// SetModel sets the value of Model for the instance
func (instance *SDDC_Memory) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", value)
}

// GetModel gets the value of Model for the instance
func (instance *SDDC_Memory) GetPropertyModel() (value string, err error) {
	retValue, err := instance.GetProperty("Model")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *SDDC_Memory) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", value)
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *SDDC_Memory) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSizeInBytes sets the value of SizeInBytes for the instance
func (instance *SDDC_Memory) SetPropertySizeInBytes(value uint64) (err error) {
	return instance.SetProperty("SizeInBytes", value)
}

// GetSizeInBytes gets the value of SizeInBytes for the instance
func (instance *SDDC_Memory) GetPropertySizeInBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("SizeInBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpeedInMHz sets the value of SpeedInMHz for the instance
func (instance *SDDC_Memory) SetPropertySpeedInMHz(value uint32) (err error) {
	return instance.SetProperty("SpeedInMHz", value)
}

// GetSpeedInMHz gets the value of SpeedInMHz for the instance
func (instance *SDDC_Memory) GetPropertySpeedInMHz() (value uint32, err error) {
	retValue, err := instance.GetProperty("SpeedInMHz")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
