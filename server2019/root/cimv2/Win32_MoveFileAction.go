// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_MoveFileAction struct
type Win32_MoveFileAction struct {
	CIM_FileAction

	//
	DestFolder string

	//
	DestName string

	//
	FileKey string

	//
	Options uint16

	//
	SourceFolder string

	//
	SourceName string
}

// SetDestFolder sets the value of DestFolder for the instance
func (instance *Win32_MoveFileAction) SetPropertyDestFolder(value string) (err error) {
	return instance.SetProperty("DestFolder", value)
}

// GetDestFolder gets the value of DestFolder for the instance
func (instance *Win32_MoveFileAction) GetPropertyDestFolder() (value string, err error) {
	retValue, err := instance.GetProperty("DestFolder")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestName sets the value of DestName for the instance
func (instance *Win32_MoveFileAction) SetPropertyDestName(value string) (err error) {
	return instance.SetProperty("DestName", value)
}

// GetDestName gets the value of DestName for the instance
func (instance *Win32_MoveFileAction) GetPropertyDestName() (value string, err error) {
	retValue, err := instance.GetProperty("DestName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileKey sets the value of FileKey for the instance
func (instance *Win32_MoveFileAction) SetPropertyFileKey(value string) (err error) {
	return instance.SetProperty("FileKey", value)
}

// GetFileKey gets the value of FileKey for the instance
func (instance *Win32_MoveFileAction) GetPropertyFileKey() (value string, err error) {
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

// SetOptions sets the value of Options for the instance
func (instance *Win32_MoveFileAction) SetPropertyOptions(value uint16) (err error) {
	return instance.SetProperty("Options", value)
}

// GetOptions gets the value of Options for the instance
func (instance *Win32_MoveFileAction) GetPropertyOptions() (value uint16, err error) {
	retValue, err := instance.GetProperty("Options")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceFolder sets the value of SourceFolder for the instance
func (instance *Win32_MoveFileAction) SetPropertySourceFolder(value string) (err error) {
	return instance.SetProperty("SourceFolder", value)
}

// GetSourceFolder gets the value of SourceFolder for the instance
func (instance *Win32_MoveFileAction) GetPropertySourceFolder() (value string, err error) {
	retValue, err := instance.GetProperty("SourceFolder")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceName sets the value of SourceName for the instance
func (instance *Win32_MoveFileAction) SetPropertySourceName(value string) (err error) {
	return instance.SetProperty("SourceName", value)
}

// GetSourceName gets the value of SourceName for the instance
func (instance *Win32_MoveFileAction) GetPropertySourceName() (value string, err error) {
	retValue, err := instance.GetProperty("SourceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
