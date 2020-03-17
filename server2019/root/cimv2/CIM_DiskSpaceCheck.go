// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_DiskSpaceCheck struct
type CIM_DiskSpaceCheck struct {
	CIM_Check

	//
	AvailableDiskSpace uint64
}

// SetAvailableDiskSpace sets the value of AvailableDiskSpace for the instance
func (instance *CIM_DiskSpaceCheck) SetPropertyAvailableDiskSpace(value uint64) (err error) {
	return instance.SetProperty("AvailableDiskSpace", value)
}

// GetAvailableDiskSpace gets the value of AvailableDiskSpace for the instance
func (instance *CIM_DiskSpaceCheck) GetPropertyAvailableDiskSpace() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableDiskSpace")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
