// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_SwapSpaceCheck struct
type CIM_SwapSpaceCheck struct {
	CIM_Check

	//
	SwapSpaceSize uint64
}

// SetSwapSpaceSize sets the value of SwapSpaceSize for the instance
func (instance *CIM_SwapSpaceCheck) SetPropertySwapSpaceSize(value uint64) (err error) {
	return instance.SetProperty("SwapSpaceSize", value)
}

// GetSwapSpaceSize gets the value of SwapSpaceSize for the instance
func (instance *CIM_SwapSpaceCheck) GetPropertySwapSpaceSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("SwapSpaceSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
