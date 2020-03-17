// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_SerialController struct
type CIM_SerialController struct {
	CIM_Controller

	// The Capabilities property defines chip level compatibility for the SerialController. Therefore, this property describes the buffering and other capabilities of the SerialController that might be inherent in the chip hardware. The property is an enumerated integer.
	Capabilities []SerialController_Capabilities

	// An array of free-form strings that provides more detailed explanations for any of the SerialController features that are indicated in the Capabilities array. Note, each entry of this array is related to the entry in the Capabilities array that is located at the same index.
	CapabilityDescriptions []string

	// Maximum baud rate in Bits per Second that is supported by the SerialController.
	MaxBaudRate uint32

	// An enumeration that indicates the operational security for the Controller. For example, information that the external interface of the Device is locked out (value=4) or "Boot Bypass" (value=6) can be described using this property.
	Security SerialController_Security
}

// SetCapabilities sets the value of Capabilities for the instance
func (instance *CIM_SerialController) SetPropertyCapabilities(value []SerialController_Capabilities) (err error) {
	return instance.SetProperty("Capabilities", value)
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *CIM_SerialController) GetPropertyCapabilities() (value []SerialController_Capabilities, err error) {
	retValue, err := instance.GetProperty("Capabilities")
	if err != nil {
		return
	}
	value, ok := retValue.([]SerialController_Capabilities)
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

// SetSecurity sets the value of Security for the instance
func (instance *CIM_SerialController) SetPropertySecurity(value SerialController_Security) (err error) {
	return instance.SetProperty("Security", value)
}

// GetSecurity gets the value of Security for the instance
func (instance *CIM_SerialController) GetPropertySecurity() (value SerialController_Security, err error) {
	retValue, err := instance.GetProperty("Security")
	if err != nil {
		return
	}
	value, ok := retValue.(SerialController_Security)
	if !ok {
		// TODO: Set an error
	}
	return
}
