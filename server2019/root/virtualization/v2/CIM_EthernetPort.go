// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_EthernetPort struct
type CIM_EthernetPort struct {
	CIM_NetworkPort

	// Capabilities of the EthernetPort. For example, the Device might support AlertOnLan, WakeOnLan, Load Balancing, or FailOver. If failover or load balancing capabilities are listed, a SpareGroup (failover) or ExtraCapacityGroup (load balancing) should also be defined to completely describe the capability.
	Capabilities []EthernetPort_Capabilities

	// An array of free-form strings that provides more detailed explanations for any of the EthernetPort features that are indicated in the Capabilities array. Note, each entry of this array is related to the entry in the Capabilities array that is located at the same index.
	CapabilityDescriptions []string

	// Specifies which capabilities are enabled from the list of all supported ones, which are defined in the Capabilities array.
	EnabledCapabilities []EthernetPort_EnabledCapabilities

	// The maximum size of the INFO (non-MAC) field that will be received or transmitted.
	MaxDataSize uint32

	// An array of free-form strings that provides more detailed explanations for any of the enabled capabilities that are specified as 'Other'.
	OtherEnabledCapabilities []string
}

// SetCapabilities sets the value of Capabilities for the instance
func (instance *CIM_EthernetPort) SetPropertyCapabilities(value []EthernetPort_Capabilities) (err error) {
	return instance.SetProperty("Capabilities", value)
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *CIM_EthernetPort) GetPropertyCapabilities() (value []EthernetPort_Capabilities, err error) {
	retValue, err := instance.GetProperty("Capabilities")
	if err != nil {
		return
	}
	value, ok := retValue.([]EthernetPort_Capabilities)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCapabilityDescriptions sets the value of CapabilityDescriptions for the instance
func (instance *CIM_EthernetPort) SetPropertyCapabilityDescriptions(value []string) (err error) {
	return instance.SetProperty("CapabilityDescriptions", value)
}

// GetCapabilityDescriptions gets the value of CapabilityDescriptions for the instance
func (instance *CIM_EthernetPort) GetPropertyCapabilityDescriptions() (value []string, err error) {
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

// SetEnabledCapabilities sets the value of EnabledCapabilities for the instance
func (instance *CIM_EthernetPort) SetPropertyEnabledCapabilities(value []EthernetPort_EnabledCapabilities) (err error) {
	return instance.SetProperty("EnabledCapabilities", value)
}

// GetEnabledCapabilities gets the value of EnabledCapabilities for the instance
func (instance *CIM_EthernetPort) GetPropertyEnabledCapabilities() (value []EthernetPort_EnabledCapabilities, err error) {
	retValue, err := instance.GetProperty("EnabledCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.([]EthernetPort_EnabledCapabilities)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxDataSize sets the value of MaxDataSize for the instance
func (instance *CIM_EthernetPort) SetPropertyMaxDataSize(value uint32) (err error) {
	return instance.SetProperty("MaxDataSize", value)
}

// GetMaxDataSize gets the value of MaxDataSize for the instance
func (instance *CIM_EthernetPort) GetPropertyMaxDataSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxDataSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherEnabledCapabilities sets the value of OtherEnabledCapabilities for the instance
func (instance *CIM_EthernetPort) SetPropertyOtherEnabledCapabilities(value []string) (err error) {
	return instance.SetProperty("OtherEnabledCapabilities", value)
}

// GetOtherEnabledCapabilities gets the value of OtherEnabledCapabilities for the instance
func (instance *CIM_EthernetPort) GetPropertyOtherEnabledCapabilities() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherEnabledCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
