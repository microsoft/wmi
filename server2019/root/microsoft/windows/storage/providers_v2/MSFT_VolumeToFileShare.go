// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_VolumeToFileShare struct
type MSFT_VolumeToFileShare struct {
	cim.WmiInstance

	//
	FileShare MSFT_FileShare

	//
	Volume MSFT_Volume
}

// SetFileShare sets the value of FileShare for the instance
func (instance *MSFT_VolumeToFileShare) SetPropertyFileShare(value MSFT_FileShare) (err error) {
	return instance.SetProperty("FileShare", value)
}

// GetFileShare gets the value of FileShare for the instance
func (instance *MSFT_VolumeToFileShare) GetPropertyFileShare() (value MSFT_FileShare, err error) {
	retValue, err := instance.GetProperty("FileShare")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_FileShare)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolume sets the value of Volume for the instance
func (instance *MSFT_VolumeToFileShare) SetPropertyVolume(value MSFT_Volume) (err error) {
	return instance.SetProperty("Volume", value)
}

// GetVolume gets the value of Volume for the instance
func (instance *MSFT_VolumeToFileShare) GetPropertyVolume() (value MSFT_Volume, err error) {
	retValue, err := instance.GetProperty("Volume")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Volume)
	if !ok {
		// TODO: Set an error
	}
	return
}
