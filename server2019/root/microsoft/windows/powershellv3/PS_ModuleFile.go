// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Powershellv3
//////////////////////////////////////////////
package powershellv3

// PS_ModuleFile struct
type PS_ModuleFile struct {
	CIM_ManagedElement

	//
	FileData []uint8

	//
	FileName string
}

// SetFileData sets the value of FileData for the instance
func (instance *PS_ModuleFile) SetPropertyFileData(value []uint8) (err error) {
	return instance.SetProperty("FileData", value)
}

// GetFileData gets the value of FileData for the instance
func (instance *PS_ModuleFile) GetPropertyFileData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("FileData")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileName sets the value of FileName for the instance
func (instance *PS_ModuleFile) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", value)
}

// GetFileName gets the value of FileName for the instance
func (instance *PS_ModuleFile) GetPropertyFileName() (value string, err error) {
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
