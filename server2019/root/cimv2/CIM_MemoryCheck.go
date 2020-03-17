// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_MemoryCheck struct
type CIM_MemoryCheck struct {
	CIM_Check

	//
	MemorySize uint64
}

// SetMemorySize sets the value of MemorySize for the instance
func (instance *CIM_MemoryCheck) SetPropertyMemorySize(value uint64) (err error) {
	return instance.SetProperty("MemorySize", value)
}

// GetMemorySize gets the value of MemorySize for the instance
func (instance *CIM_MemoryCheck) GetPropertyMemorySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemorySize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
