// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_Synthetic3DDisplayController struct
type Msvm_Synthetic3DDisplayController struct {
	CIM_DisplayController

	//
	AllocatedGPU string
}

// SetAllocatedGPU sets the value of AllocatedGPU for the instance
func (instance *Msvm_Synthetic3DDisplayController) SetPropertyAllocatedGPU(value string) (err error) {
	return instance.SetProperty("AllocatedGPU", value)
}

// GetAllocatedGPU gets the value of AllocatedGPU for the instance
func (instance *Msvm_Synthetic3DDisplayController) GetPropertyAllocatedGPU() (value string, err error) {
	retValue, err := instance.GetProperty("AllocatedGPU")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
