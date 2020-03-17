// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PageFile struct
type Win32_PageFile struct {
	CIM_DataFile

	//
	FreeSpace uint32

	//
	InitialSize uint32

	//
	MaximumSize uint32
}

// SetFreeSpace sets the value of FreeSpace for the instance
func (instance *Win32_PageFile) SetPropertyFreeSpace(value uint32) (err error) {
	return instance.SetProperty("FreeSpace", value)
}

// GetFreeSpace gets the value of FreeSpace for the instance
func (instance *Win32_PageFile) GetPropertyFreeSpace() (value uint32, err error) {
	retValue, err := instance.GetProperty("FreeSpace")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitialSize sets the value of InitialSize for the instance
func (instance *Win32_PageFile) SetPropertyInitialSize(value uint32) (err error) {
	return instance.SetProperty("InitialSize", value)
}

// GetInitialSize gets the value of InitialSize for the instance
func (instance *Win32_PageFile) GetPropertyInitialSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("InitialSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumSize sets the value of MaximumSize for the instance
func (instance *Win32_PageFile) SetPropertyMaximumSize(value uint32) (err error) {
	return instance.SetProperty("MaximumSize", value)
}

// GetMaximumSize gets the value of MaximumSize for the instance
func (instance *Win32_PageFile) GetPropertyMaximumSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
