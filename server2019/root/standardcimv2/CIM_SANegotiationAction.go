// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_SANegotiationAction struct
type CIM_SANegotiationAction struct {
	CIM_SAAction

	//
	IdleDurationSeconds uint64

	//
	MinLifetimeKilobytes uint64

	//
	MinLifetimeSeconds uint64
}

// SetIdleDurationSeconds sets the value of IdleDurationSeconds for the instance
func (instance *CIM_SANegotiationAction) SetPropertyIdleDurationSeconds(value uint64) (err error) {
	return instance.SetProperty("IdleDurationSeconds", value)
}

// GetIdleDurationSeconds gets the value of IdleDurationSeconds for the instance
func (instance *CIM_SANegotiationAction) GetPropertyIdleDurationSeconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("IdleDurationSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinLifetimeKilobytes sets the value of MinLifetimeKilobytes for the instance
func (instance *CIM_SANegotiationAction) SetPropertyMinLifetimeKilobytes(value uint64) (err error) {
	return instance.SetProperty("MinLifetimeKilobytes", value)
}

// GetMinLifetimeKilobytes gets the value of MinLifetimeKilobytes for the instance
func (instance *CIM_SANegotiationAction) GetPropertyMinLifetimeKilobytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinLifetimeKilobytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinLifetimeSeconds sets the value of MinLifetimeSeconds for the instance
func (instance *CIM_SANegotiationAction) SetPropertyMinLifetimeSeconds(value uint64) (err error) {
	return instance.SetProperty("MinLifetimeSeconds", value)
}

// GetMinLifetimeSeconds gets the value of MinLifetimeSeconds for the instance
func (instance *CIM_SANegotiationAction) GetPropertyMinLifetimeSeconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinLifetimeSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
