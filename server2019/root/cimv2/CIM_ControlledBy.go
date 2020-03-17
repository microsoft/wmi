// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_ControlledBy struct
type CIM_ControlledBy struct {
	CIM_DeviceConnection

	//
	AccessState uint16

	//
	NumberOfHardResets uint32

	//
	NumberOfSoftResets uint32
}

// SetAccessState sets the value of AccessState for the instance
func (instance *CIM_ControlledBy) SetPropertyAccessState(value uint16) (err error) {
	return instance.SetProperty("AccessState", value)
}

// GetAccessState gets the value of AccessState for the instance
func (instance *CIM_ControlledBy) GetPropertyAccessState() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfHardResets sets the value of NumberOfHardResets for the instance
func (instance *CIM_ControlledBy) SetPropertyNumberOfHardResets(value uint32) (err error) {
	return instance.SetProperty("NumberOfHardResets", value)
}

// GetNumberOfHardResets gets the value of NumberOfHardResets for the instance
func (instance *CIM_ControlledBy) GetPropertyNumberOfHardResets() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfHardResets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfSoftResets sets the value of NumberOfSoftResets for the instance
func (instance *CIM_ControlledBy) SetPropertyNumberOfSoftResets(value uint32) (err error) {
	return instance.SetProperty("NumberOfSoftResets", value)
}

// GetNumberOfSoftResets gets the value of NumberOfSoftResets for the instance
func (instance *CIM_ControlledBy) GetPropertyNumberOfSoftResets() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfSoftResets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
