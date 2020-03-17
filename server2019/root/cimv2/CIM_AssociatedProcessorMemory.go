// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_AssociatedProcessorMemory struct
type CIM_AssociatedProcessorMemory struct {
	CIM_AssociatedMemory

	//
	BusSpeed uint32
}

// SetBusSpeed sets the value of BusSpeed for the instance
func (instance *CIM_AssociatedProcessorMemory) SetPropertyBusSpeed(value uint32) (err error) {
	return instance.SetProperty("BusSpeed", value)
}

// GetBusSpeed gets the value of BusSpeed for the instance
func (instance *CIM_AssociatedProcessorMemory) GetPropertyBusSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
