// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

// MSFT_MTDiskVolume struct
type MSFT_MTDiskVolume struct {
	CIM_ManagedElement

	//
	FormattedSize uint64

	//
	PageFile bool

	//
	SystemDisk bool

	//
	VolumePath string
}

// SetFormattedSize sets the value of FormattedSize for the instance
func (instance *MSFT_MTDiskVolume) SetPropertyFormattedSize(value uint64) (err error) {
	return instance.SetProperty("FormattedSize", value)
}

// GetFormattedSize gets the value of FormattedSize for the instance
func (instance *MSFT_MTDiskVolume) GetPropertyFormattedSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("FormattedSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPageFile sets the value of PageFile for the instance
func (instance *MSFT_MTDiskVolume) SetPropertyPageFile(value bool) (err error) {
	return instance.SetProperty("PageFile", value)
}

// GetPageFile gets the value of PageFile for the instance
func (instance *MSFT_MTDiskVolume) GetPropertyPageFile() (value bool, err error) {
	retValue, err := instance.GetProperty("PageFile")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemDisk sets the value of SystemDisk for the instance
func (instance *MSFT_MTDiskVolume) SetPropertySystemDisk(value bool) (err error) {
	return instance.SetProperty("SystemDisk", value)
}

// GetSystemDisk gets the value of SystemDisk for the instance
func (instance *MSFT_MTDiskVolume) GetPropertySystemDisk() (value bool, err error) {
	retValue, err := instance.GetProperty("SystemDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolumePath sets the value of VolumePath for the instance
func (instance *MSFT_MTDiskVolume) SetPropertyVolumePath(value string) (err error) {
	return instance.SetProperty("VolumePath", value)
}

// GetVolumePath gets the value of VolumePath for the instance
func (instance *MSFT_MTDiskVolume) GetPropertyVolumePath() (value string, err error) {
	retValue, err := instance.GetProperty("VolumePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
