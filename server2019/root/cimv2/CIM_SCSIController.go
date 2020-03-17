// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_SCSIController struct
type CIM_SCSIController struct {
	CIM_Controller

	//
	ControllerTimeouts uint32

	//
	MaxDataWidth uint32

	//
	MaxTransferRate uint64

	//
	ProtectionManagement uint16
}

// SetControllerTimeouts sets the value of ControllerTimeouts for the instance
func (instance *CIM_SCSIController) SetPropertyControllerTimeouts(value uint32) (err error) {
	return instance.SetProperty("ControllerTimeouts", value)
}

// GetControllerTimeouts gets the value of ControllerTimeouts for the instance
func (instance *CIM_SCSIController) GetPropertyControllerTimeouts() (value uint32, err error) {
	retValue, err := instance.GetProperty("ControllerTimeouts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxDataWidth sets the value of MaxDataWidth for the instance
func (instance *CIM_SCSIController) SetPropertyMaxDataWidth(value uint32) (err error) {
	return instance.SetProperty("MaxDataWidth", value)
}

// GetMaxDataWidth gets the value of MaxDataWidth for the instance
func (instance *CIM_SCSIController) GetPropertyMaxDataWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxDataWidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxTransferRate sets the value of MaxTransferRate for the instance
func (instance *CIM_SCSIController) SetPropertyMaxTransferRate(value uint64) (err error) {
	return instance.SetProperty("MaxTransferRate", value)
}

// GetMaxTransferRate gets the value of MaxTransferRate for the instance
func (instance *CIM_SCSIController) GetPropertyMaxTransferRate() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxTransferRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtectionManagement sets the value of ProtectionManagement for the instance
func (instance *CIM_SCSIController) SetPropertyProtectionManagement(value uint16) (err error) {
	return instance.SetProperty("ProtectionManagement", value)
}

// GetProtectionManagement gets the value of ProtectionManagement for the instance
func (instance *CIM_SCSIController) GetPropertyProtectionManagement() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProtectionManagement")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
