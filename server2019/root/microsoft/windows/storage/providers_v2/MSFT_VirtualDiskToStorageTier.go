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

// MSFT_VirtualDiskToStorageTier struct
type MSFT_VirtualDiskToStorageTier struct {
	cim.WmiInstance

	//
	StorageTier MSFT_StorageTier

	//
	VirtualDisk MSFT_VirtualDisk
}

// SetStorageTier sets the value of StorageTier for the instance
func (instance *MSFT_VirtualDiskToStorageTier) SetPropertyStorageTier(value MSFT_StorageTier) (err error) {
	return instance.SetProperty("StorageTier", value)
}

// GetStorageTier gets the value of StorageTier for the instance
func (instance *MSFT_VirtualDiskToStorageTier) GetPropertyStorageTier() (value MSFT_StorageTier, err error) {
	retValue, err := instance.GetProperty("StorageTier")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageTier)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDisk sets the value of VirtualDisk for the instance
func (instance *MSFT_VirtualDiskToStorageTier) SetPropertyVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("VirtualDisk", value)
}

// GetVirtualDisk gets the value of VirtualDisk for the instance
func (instance *MSFT_VirtualDiskToStorageTier) GetPropertyVirtualDisk() (value MSFT_VirtualDisk, err error) {
	retValue, err := instance.GetProperty("VirtualDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_VirtualDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}
