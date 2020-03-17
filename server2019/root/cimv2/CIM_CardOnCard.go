// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_CardOnCard struct
type CIM_CardOnCard struct {
	CIM_Container

	//
	MountOrSlotDescription string
}

// SetMountOrSlotDescription sets the value of MountOrSlotDescription for the instance
func (instance *CIM_CardOnCard) SetPropertyMountOrSlotDescription(value string) (err error) {
	return instance.SetProperty("MountOrSlotDescription", value)
}

// GetMountOrSlotDescription gets the value of MountOrSlotDescription for the instance
func (instance *CIM_CardOnCard) GetPropertyMountOrSlotDescription() (value string, err error) {
	retValue, err := instance.GetProperty("MountOrSlotDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
