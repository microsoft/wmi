// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_SerialController struct
type CIM_SerialController struct {
	CIM_Controller

	//
	Capabilities []uint16

	//
	CapabilityDescriptions []string

	//
	MaxBaudRate uint32
}

// SetCapabilities sets the value of Capabilities for the instance
func (instance *CIM_SerialController) SetPropertyCapabilities(value []uint16) (err error) {
	return instance.SetProperty("Capabilities", value)
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *CIM_SerialController) GetPropertyCapabilities() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Capabilities")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCapabilityDescriptions sets the value of CapabilityDescriptions for the instance
func (instance *CIM_SerialController) SetPropertyCapabilityDescriptions(value []string) (err error) {
	return instance.SetProperty("CapabilityDescriptions", value)
}

// GetCapabilityDescriptions gets the value of CapabilityDescriptions for the instance
func (instance *CIM_SerialController) GetPropertyCapabilityDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("CapabilityDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxBaudRate sets the value of MaxBaudRate for the instance
func (instance *CIM_SerialController) SetPropertyMaxBaudRate(value uint32) (err error) {
	return instance.SetProperty("MaxBaudRate", value)
}

// GetMaxBaudRate gets the value of MaxBaudRate for the instance
func (instance *CIM_SerialController) GetPropertyMaxBaudRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxBaudRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
