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

// MSFT_DiskToStorageReliabilityCounter struct
type MSFT_DiskToStorageReliabilityCounter struct {
	cim.WmiInstance

	//
	Disk MSFT_Disk

	//
	StorageReliabilityCounter MSFT_StorageReliabilityCounter
}

// SetDisk sets the value of Disk for the instance
func (instance *MSFT_DiskToStorageReliabilityCounter) SetPropertyDisk(value MSFT_Disk) (err error) {
	return instance.SetProperty("Disk", value)
}

// GetDisk gets the value of Disk for the instance
func (instance *MSFT_DiskToStorageReliabilityCounter) GetPropertyDisk() (value MSFT_Disk, err error) {
	retValue, err := instance.GetProperty("Disk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Disk)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageReliabilityCounter sets the value of StorageReliabilityCounter for the instance
func (instance *MSFT_DiskToStorageReliabilityCounter) SetPropertyStorageReliabilityCounter(value MSFT_StorageReliabilityCounter) (err error) {
	return instance.SetProperty("StorageReliabilityCounter", value)
}

// GetStorageReliabilityCounter gets the value of StorageReliabilityCounter for the instance
func (instance *MSFT_DiskToStorageReliabilityCounter) GetPropertyStorageReliabilityCounter() (value MSFT_StorageReliabilityCounter, err error) {
	retValue, err := instance.GetProperty("StorageReliabilityCounter")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageReliabilityCounter)
	if !ok {
		// TODO: Set an error
	}
	return
}
