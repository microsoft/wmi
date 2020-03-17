// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_DiskDrive struct
type Msvm_DiskDrive struct {
	CIM_DiskDrive

	//
	DriveNumber uint32
}

// SetDriveNumber sets the value of DriveNumber for the instance
func (instance *Msvm_DiskDrive) SetPropertyDriveNumber(value uint32) (err error) {
	return instance.SetProperty("DriveNumber", value)
}

// GetDriveNumber gets the value of DriveNumber for the instance
func (instance *Msvm_DiskDrive) GetPropertyDriveNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DriveNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
