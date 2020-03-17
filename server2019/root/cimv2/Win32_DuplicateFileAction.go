// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DuplicateFileAction struct
type Win32_DuplicateFileAction struct {
	CIM_CopyFileAction

	//
	FileKey string
}

// SetFileKey sets the value of FileKey for the instance
func (instance *Win32_DuplicateFileAction) SetPropertyFileKey(value string) (err error) {
	return instance.SetProperty("FileKey", value)
}

// GetFileKey gets the value of FileKey for the instance
func (instance *Win32_DuplicateFileAction) GetPropertyFileKey() (value string, err error) {
	retValue, err := instance.GetProperty("FileKey")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
