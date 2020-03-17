// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_DVDDrive struct
type CIM_DVDDrive struct {
	CIM_MediaAccessDevice

	// The CD and DVD formats that are supported by this Device. For example, the Drive may support "CD-ROM" and "DVD-RAM". In this case, the values 16 and 24 would be written to the array. This property's values align with those defined in PhysicalMedia.MediaType.
	FormatsSupported []DVDDrive_FormatsSupported
}

// SetFormatsSupported sets the value of FormatsSupported for the instance
func (instance *CIM_DVDDrive) SetPropertyFormatsSupported(value []DVDDrive_FormatsSupported) (err error) {
	return instance.SetProperty("FormatsSupported", value)
}

// GetFormatsSupported gets the value of FormatsSupported for the instance
func (instance *CIM_DVDDrive) GetPropertyFormatsSupported() (value []DVDDrive_FormatsSupported, err error) {
	retValue, err := instance.GetProperty("FormatsSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]DVDDrive_FormatsSupported)
	if !ok {
		// TODO: Set an error
	}
	return
}
