// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_ChassisInRack struct
type CIM_ChassisInRack struct {
	CIM_Container

	//
	BottomU uint16
}

// SetBottomU sets the value of BottomU for the instance
func (instance *CIM_ChassisInRack) SetPropertyBottomU(value uint16) (err error) {
	return instance.SetProperty("BottomU", value)
}

// GetBottomU gets the value of BottomU for the instance
func (instance *CIM_ChassisInRack) GetPropertyBottomU() (value uint16, err error) {
	retValue, err := instance.GetProperty("BottomU")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
