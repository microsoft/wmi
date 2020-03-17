// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_IRQResource struct
type Win32_IRQResource struct {
	CIM_IRQ

	//
	Hardware bool

	//
	Vector uint32
}

// SetHardware sets the value of Hardware for the instance
func (instance *Win32_IRQResource) SetPropertyHardware(value bool) (err error) {
	return instance.SetProperty("Hardware", value)
}

// GetHardware gets the value of Hardware for the instance
func (instance *Win32_IRQResource) GetPropertyHardware() (value bool, err error) {
	retValue, err := instance.GetProperty("Hardware")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVector sets the value of Vector for the instance
func (instance *Win32_IRQResource) SetPropertyVector(value uint32) (err error) {
	return instance.SetProperty("Vector", value)
}

// GetVector gets the value of Vector for the instance
func (instance *Win32_IRQResource) GetPropertyVector() (value uint32, err error) {
	retValue, err := instance.GetProperty("Vector")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
