// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_CopyFileToGuestSettingData struct
type Msvm_CopyFileToGuestSettingData struct {
	CIM_SettingData

	//
	CreateFullPath bool

	//
	DestinationPath string

	//
	OverwriteExisting bool

	//
	SourcePath string
}

// SetCreateFullPath sets the value of CreateFullPath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) SetPropertyCreateFullPath(value bool) (err error) {
	return instance.SetProperty("CreateFullPath", value)
}

// GetCreateFullPath gets the value of CreateFullPath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) GetPropertyCreateFullPath() (value bool, err error) {
	retValue, err := instance.GetProperty("CreateFullPath")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestinationPath sets the value of DestinationPath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) SetPropertyDestinationPath(value string) (err error) {
	return instance.SetProperty("DestinationPath", value)
}

// GetDestinationPath gets the value of DestinationPath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) GetPropertyDestinationPath() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOverwriteExisting sets the value of OverwriteExisting for the instance
func (instance *Msvm_CopyFileToGuestSettingData) SetPropertyOverwriteExisting(value bool) (err error) {
	return instance.SetProperty("OverwriteExisting", value)
}

// GetOverwriteExisting gets the value of OverwriteExisting for the instance
func (instance *Msvm_CopyFileToGuestSettingData) GetPropertyOverwriteExisting() (value bool, err error) {
	retValue, err := instance.GetProperty("OverwriteExisting")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourcePath sets the value of SourcePath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) SetPropertySourcePath(value string) (err error) {
	return instance.SetProperty("SourcePath", value)
}

// GetSourcePath gets the value of SourcePath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) GetPropertySourcePath() (value string, err error) {
	retValue, err := instance.GetProperty("SourcePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
