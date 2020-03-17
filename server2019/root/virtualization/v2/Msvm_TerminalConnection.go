// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_TerminalConnection struct
type Msvm_TerminalConnection struct {
	CIM_EnabledLogicalElement

	//
	ConnectionID string
}

// SetConnectionID sets the value of ConnectionID for the instance
func (instance *Msvm_TerminalConnection) SetPropertyConnectionID(value string) (err error) {
	return instance.SetProperty("ConnectionID", value)
}

// GetConnectionID gets the value of ConnectionID for the instance
func (instance *Msvm_TerminalConnection) GetPropertyConnectionID() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
