// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_ExtraCapacityGroup struct
type CIM_ExtraCapacityGroup struct {
	CIM_RedundancyGroup

	//
	MinNumberNeeded uint32
}

// SetMinNumberNeeded sets the value of MinNumberNeeded for the instance
func (instance *CIM_ExtraCapacityGroup) SetPropertyMinNumberNeeded(value uint32) (err error) {
	return instance.SetProperty("MinNumberNeeded", value)
}

// GetMinNumberNeeded gets the value of MinNumberNeeded for the instance
func (instance *CIM_ExtraCapacityGroup) GetPropertyMinNumberNeeded() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinNumberNeeded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
