// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_SerialPortSettingData struct
type Msvm_SerialPortSettingData struct {
	CIM_ResourceAllocationSettingData

	//
	DebuggerMode bool
}

// SetDebuggerMode sets the value of DebuggerMode for the instance
func (instance *Msvm_SerialPortSettingData) SetPropertyDebuggerMode(value bool) (err error) {
	return instance.SetProperty("DebuggerMode", value)
}

// GetDebuggerMode gets the value of DebuggerMode for the instance
func (instance *Msvm_SerialPortSettingData) GetPropertyDebuggerMode() (value bool, err error) {
	retValue, err := instance.GetProperty("DebuggerMode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
