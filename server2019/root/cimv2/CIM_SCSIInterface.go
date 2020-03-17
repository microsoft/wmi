// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_SCSIInterface struct
type CIM_SCSIInterface struct {
	CIM_ControlledBy

	//
	SCSIRetries uint32

	//
	SCSITimeouts uint32
}

// SetSCSIRetries sets the value of SCSIRetries for the instance
func (instance *CIM_SCSIInterface) SetPropertySCSIRetries(value uint32) (err error) {
	return instance.SetProperty("SCSIRetries", value)
}

// GetSCSIRetries gets the value of SCSIRetries for the instance
func (instance *CIM_SCSIInterface) GetPropertySCSIRetries() (value uint32, err error) {
	retValue, err := instance.GetProperty("SCSIRetries")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSCSITimeouts sets the value of SCSITimeouts for the instance
func (instance *CIM_SCSIInterface) SetPropertySCSITimeouts(value uint32) (err error) {
	return instance.SetProperty("SCSITimeouts", value)
}

// GetSCSITimeouts gets the value of SCSITimeouts for the instance
func (instance *CIM_SCSIInterface) GetPropertySCSITimeouts() (value uint32, err error) {
	retValue, err := instance.GetProperty("SCSITimeouts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
