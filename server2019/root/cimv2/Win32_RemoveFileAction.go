// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_RemoveFileAction struct
type Win32_RemoveFileAction struct {
	CIM_RemoveFileAction

	//
	DirProperty string

	//
	FileKey string

	//
	FileName string

	//
	InstallMode uint16
}

// SetDirProperty sets the value of DirProperty for the instance
func (instance *Win32_RemoveFileAction) SetPropertyDirProperty(value string) (err error) {
	return instance.SetProperty("DirProperty", value)
}

// GetDirProperty gets the value of DirProperty for the instance
func (instance *Win32_RemoveFileAction) GetPropertyDirProperty() (value string, err error) {
	retValue, err := instance.GetProperty("DirProperty")
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
func (instance *Win32_RemoveFileAction) SetPropertyFileKey(value string) (err error) {
	return instance.SetProperty("FileKey", value)
}

// GetFileKey gets the value of FileKey for the instance
func (instance *Win32_RemoveFileAction) GetPropertyFileKey() (value string, err error) {
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

// SetFileName sets the value of FileName for the instance
func (instance *Win32_RemoveFileAction) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", value)
}

// GetFileName gets the value of FileName for the instance
func (instance *Win32_RemoveFileAction) GetPropertyFileName() (value string, err error) {
	retValue, err := instance.GetProperty("FileName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstallMode sets the value of InstallMode for the instance
func (instance *Win32_RemoveFileAction) SetPropertyInstallMode(value uint16) (err error) {
	return instance.SetProperty("InstallMode", value)
}

// GetInstallMode gets the value of InstallMode for the instance
func (instance *Win32_RemoveFileAction) GetPropertyInstallMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("InstallMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
