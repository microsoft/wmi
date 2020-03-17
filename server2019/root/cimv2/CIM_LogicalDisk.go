// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_LogicalDisk struct
type CIM_LogicalDisk struct {
	CIM_StorageExtent

	//
	FreeSpace uint64

	//
	Size uint64
}

// SetFreeSpace sets the value of FreeSpace for the instance
func (instance *CIM_LogicalDisk) SetPropertyFreeSpace(value uint64) (err error) {
	return instance.SetProperty("FreeSpace", value)
}

// GetFreeSpace gets the value of FreeSpace for the instance
func (instance *CIM_LogicalDisk) GetPropertyFreeSpace() (value uint64, err error) {
	retValue, err := instance.GetProperty("FreeSpace")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSize sets the value of Size for the instance
func (instance *CIM_LogicalDisk) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", value)
}

// GetSize gets the value of Size for the instance
func (instance *CIM_LogicalDisk) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
