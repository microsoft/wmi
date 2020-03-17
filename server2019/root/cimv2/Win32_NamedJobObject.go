// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_NamedJobObject struct
type Win32_NamedJobObject struct {
	CIM_CollectionOfMSEs

	//
	BasicUIRestrictions uint32
}

// SetBasicUIRestrictions sets the value of BasicUIRestrictions for the instance
func (instance *Win32_NamedJobObject) SetPropertyBasicUIRestrictions(value uint32) (err error) {
	return instance.SetProperty("BasicUIRestrictions", value)
}

// GetBasicUIRestrictions gets the value of BasicUIRestrictions for the instance
func (instance *Win32_NamedJobObject) GetPropertyBasicUIRestrictions() (value uint32, err error) {
	retValue, err := instance.GetProperty("BasicUIRestrictions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
