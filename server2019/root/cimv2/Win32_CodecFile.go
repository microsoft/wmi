// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_CodecFile struct
type Win32_CodecFile struct {
	CIM_DataFile

	//
	Group string
}

// SetGroup sets the value of Group for the instance
func (instance *Win32_CodecFile) SetPropertyGroup(value string) (err error) {
	return instance.SetProperty("Group", value)
}

// GetGroup gets the value of Group for the instance
func (instance *Win32_CodecFile) GetPropertyGroup() (value string, err error) {
	retValue, err := instance.GetProperty("Group")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
