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

// MSFT_StoragePoolToStorageTier struct
type MSFT_StoragePoolToStorageTier struct {
	cim.WmiInstance

	//
	StoragePool MSFT_StoragePool

	//
	StorageTier MSFT_StorageTier
}

// SetStoragePool sets the value of StoragePool for the instance
func (instance *MSFT_StoragePoolToStorageTier) SetPropertyStoragePool(value MSFT_StoragePool) (err error) {
	return instance.SetProperty("StoragePool", value)
}

// GetStoragePool gets the value of StoragePool for the instance
func (instance *MSFT_StoragePoolToStorageTier) GetPropertyStoragePool() (value MSFT_StoragePool, err error) {
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

// SetStorageTier sets the value of StorageTier for the instance
func (instance *MSFT_StoragePoolToStorageTier) SetPropertyStorageTier(value MSFT_StorageTier) (err error) {
	return instance.SetProperty("StorageTier", value)
}

// GetStorageTier gets the value of StorageTier for the instance
func (instance *MSFT_StoragePoolToStorageTier) GetPropertyStorageTier() (value MSFT_StorageTier, err error) {
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
