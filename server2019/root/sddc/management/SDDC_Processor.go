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

// SDDC_Processor struct
type SDDC_Processor struct {
	cim.WmiInstance

	//
	Manufacturer string

	//
	MaxClockSpeedInMHz uint32

	//
	Model string

	//
	NumberOfCores uint32

	//
	NumberOfLogicalProcessors uint32
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *SDDC_Processor) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *SDDC_Processor) GetPropertyManufacturer() (value string, err error) {
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

// SetMaxClockSpeedInMHz sets the value of MaxClockSpeedInMHz for the instance
func (instance *SDDC_Processor) SetPropertyMaxClockSpeedInMHz(value uint32) (err error) {
	return instance.SetProperty("MaxClockSpeedInMHz", value)
}

// GetMaxClockSpeedInMHz gets the value of MaxClockSpeedInMHz for the instance
func (instance *SDDC_Processor) GetPropertyMaxClockSpeedInMHz() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxClockSpeedInMHz")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModel sets the value of Model for the instance
func (instance *SDDC_Processor) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", value)
}

// GetModel gets the value of Model for the instance
func (instance *SDDC_Processor) GetPropertyModel() (value string, err error) {
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

// SetNumberOfCores sets the value of NumberOfCores for the instance
func (instance *SDDC_Processor) SetPropertyNumberOfCores(value uint32) (err error) {
	return instance.SetProperty("NumberOfCores", value)
}

// GetNumberOfCores gets the value of NumberOfCores for the instance
func (instance *SDDC_Processor) GetPropertyNumberOfCores() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfCores")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfLogicalProcessors sets the value of NumberOfLogicalProcessors for the instance
func (instance *SDDC_Processor) SetPropertyNumberOfLogicalProcessors(value uint32) (err error) {
	return instance.SetProperty("NumberOfLogicalProcessors", value)
}

// GetNumberOfLogicalProcessors gets the value of NumberOfLogicalProcessors for the instance
func (instance *SDDC_Processor) GetPropertyNumberOfLogicalProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfLogicalProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
