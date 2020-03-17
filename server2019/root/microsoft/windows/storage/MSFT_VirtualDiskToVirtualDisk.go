// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_VirtualDiskToVirtualDisk struct
type MSFT_VirtualDiskToVirtualDisk struct {
	MSFT_Synchronized

	//
	SourceVirtualDisk MSFT_VirtualDisk

	//
	TargetVirtualDisk MSFT_VirtualDisk
}

// SetSourceVirtualDisk sets the value of SourceVirtualDisk for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertySourceVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("SourceVirtualDisk", value)
}

// GetSourceVirtualDisk gets the value of SourceVirtualDisk for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertySourceVirtualDisk() (value MSFT_VirtualDisk, err error) {
	retValue, err := instance.GetProperty("SourceVirtualDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_VirtualDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetVirtualDisk sets the value of TargetVirtualDisk for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) SetPropertyTargetVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("TargetVirtualDisk", value)
}

// GetTargetVirtualDisk gets the value of TargetVirtualDisk for the instance
func (instance *MSFT_VirtualDiskToVirtualDisk) GetPropertyTargetVirtualDisk() (value MSFT_VirtualDisk, err error) {
	retValue, err := instance.GetProperty("TargetVirtualDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_VirtualDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}
