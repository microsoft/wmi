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

// MSFT_StoragePoolToPhysicalDisk struct
type MSFT_StoragePoolToPhysicalDisk struct {
	cim.WmiInstance

	//
	PhysicalDisk MSFT_PhysicalDisk

	//
	StoragePool MSFT_StoragePool
}

// SetPhysicalDisk sets the value of PhysicalDisk for the instance
func (instance *MSFT_StoragePoolToPhysicalDisk) SetPropertyPhysicalDisk(value MSFT_PhysicalDisk) (err error) {
	return instance.SetProperty("PhysicalDisk", value)
}

// GetPhysicalDisk gets the value of PhysicalDisk for the instance
func (instance *MSFT_StoragePoolToPhysicalDisk) GetPropertyPhysicalDisk() (value MSFT_PhysicalDisk, err error) {
	retValue, err := instance.GetProperty("PhysicalDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_PhysicalDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStoragePool sets the value of StoragePool for the instance
func (instance *MSFT_StoragePoolToPhysicalDisk) SetPropertyStoragePool(value MSFT_StoragePool) (err error) {
	return instance.SetProperty("StoragePool", value)
}

// GetStoragePool gets the value of StoragePool for the instance
func (instance *MSFT_StoragePoolToPhysicalDisk) GetPropertyStoragePool() (value MSFT_StoragePool, err error) {
	retValue, err := instance.GetProperty("StoragePool")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StoragePool)
	if !ok {
		// TODO: Set an error
	}
	return
}
