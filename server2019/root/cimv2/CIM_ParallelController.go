// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_ParallelController struct
type CIM_ParallelController struct {
	CIM_Controller

	//
	Capabilities []uint16

	//
	CapabilityDescriptions []string

	//
	DMASupport bool
}

// SetCapabilities sets the value of Capabilities for the instance
func (instance *CIM_ParallelController) SetPropertyCapabilities(value []uint16) (err error) {
	return instance.SetProperty("Capabilities", value)
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *CIM_ParallelController) GetPropertyCapabilities() (value []uint16, err error) {
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
func (instance *CIM_ParallelController) SetPropertyCapabilityDescriptions(value []string) (err error) {
	return instance.SetProperty("CapabilityDescriptions", value)
}

// GetCapabilityDescriptions gets the value of CapabilityDescriptions for the instance
func (instance *CIM_ParallelController) GetPropertyCapabilityDescriptions() (value []string, err error) {
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

// SetDMASupport sets the value of DMASupport for the instance
func (instance *CIM_ParallelController) SetPropertyDMASupport(value bool) (err error) {
	return instance.SetProperty("DMASupport", value)
}

// GetDMASupport gets the value of DMASupport for the instance
func (instance *CIM_ParallelController) GetPropertyDMASupport() (value bool, err error) {
	retValue, err := instance.GetProperty("DMASupport")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
