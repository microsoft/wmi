// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_Controller struct
type CIM_Controller struct {
	CIM_LogicalDevice

	// Maximum number of directly addressable entities that are supported by this Controller. A value of 0 should be used if the number is unknown or unlimited.
	MaxNumberControlled uint32

	// A free-form string that provides more information that is related to the ProtocolSupported by the Controller.
	ProtocolDescription string

	// The protocol used by the Controller to access controlled Devices.
	ProtocolSupported Controller_ProtocolSupported

	// Time of last reset of the Controller.
	TimeOfLastReset string
}

// SetMaxNumberControlled sets the value of MaxNumberControlled for the instance
func (instance *CIM_Controller) SetPropertyMaxNumberControlled(value uint32) (err error) {
	return instance.SetProperty("MaxNumberControlled", value)
}

// GetMaxNumberControlled gets the value of MaxNumberControlled for the instance
func (instance *CIM_Controller) GetPropertyMaxNumberControlled() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumberControlled")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolDescription sets the value of ProtocolDescription for the instance
func (instance *CIM_Controller) SetPropertyProtocolDescription(value string) (err error) {
	return instance.SetProperty("ProtocolDescription", value)
}

// GetProtocolDescription gets the value of ProtocolDescription for the instance
func (instance *CIM_Controller) GetPropertyProtocolDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ProtocolDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolSupported sets the value of ProtocolSupported for the instance
func (instance *CIM_Controller) SetPropertyProtocolSupported(value Controller_ProtocolSupported) (err error) {
	return instance.SetProperty("ProtocolSupported", value)
}

// GetProtocolSupported gets the value of ProtocolSupported for the instance
func (instance *CIM_Controller) GetPropertyProtocolSupported() (value Controller_ProtocolSupported, err error) {
	retValue, err := instance.GetProperty("ProtocolSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(Controller_ProtocolSupported)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeOfLastReset sets the value of TimeOfLastReset for the instance
func (instance *CIM_Controller) SetPropertyTimeOfLastReset(value string) (err error) {
	return instance.SetProperty("TimeOfLastReset", value)
}

// GetTimeOfLastReset gets the value of TimeOfLastReset for the instance
func (instance *CIM_Controller) GetPropertyTimeOfLastReset() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfLastReset")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
