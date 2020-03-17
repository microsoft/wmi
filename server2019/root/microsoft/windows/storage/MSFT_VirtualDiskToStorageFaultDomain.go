// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_VirtualDiskToStorageFaultDomain struct
type MSFT_VirtualDiskToStorageFaultDomain struct {
	cim.WmiInstance

	//
	StorageFaultDomain MSFT_StorageFaultDomain

	//
	VirtualDisk MSFT_VirtualDisk
}

// SetStorageFaultDomain sets the value of StorageFaultDomain for the instance
func (instance *MSFT_VirtualDiskToStorageFaultDomain) SetPropertyStorageFaultDomain(value MSFT_StorageFaultDomain) (err error) {
	return instance.SetProperty("StorageFaultDomain", value)
}

// GetStorageFaultDomain gets the value of StorageFaultDomain for the instance
func (instance *MSFT_VirtualDiskToStorageFaultDomain) GetPropertyStorageFaultDomain() (value MSFT_StorageFaultDomain, err error) {
	retValue, err := instance.GetProperty("StorageFaultDomain")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageFaultDomain)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDisk sets the value of VirtualDisk for the instance
func (instance *MSFT_VirtualDiskToStorageFaultDomain) SetPropertyVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("VirtualDisk", value)
}

// GetVirtualDisk gets the value of VirtualDisk for the instance
func (instance *MSFT_VirtualDiskToStorageFaultDomain) GetPropertyVirtualDisk() (value MSFT_VirtualDisk, err error) {
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
