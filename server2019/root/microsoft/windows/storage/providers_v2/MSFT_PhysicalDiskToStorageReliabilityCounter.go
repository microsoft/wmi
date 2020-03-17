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

// MSFT_PhysicalDiskToStorageReliabilityCounter struct
type MSFT_PhysicalDiskToStorageReliabilityCounter struct {
	cim.WmiInstance

	//
	PhysicalDisk MSFT_PhysicalDisk

	//
	StorageReliabilityCounter MSFT_StorageReliabilityCounter
}

// SetPhysicalDisk sets the value of PhysicalDisk for the instance
func (instance *MSFT_PhysicalDiskToStorageReliabilityCounter) SetPropertyPhysicalDisk(value MSFT_PhysicalDisk) (err error) {
	return instance.SetProperty("PhysicalDisk", value)
}

// GetPhysicalDisk gets the value of PhysicalDisk for the instance
func (instance *MSFT_PhysicalDiskToStorageReliabilityCounter) GetPropertyPhysicalDisk() (value MSFT_PhysicalDisk, err error) {
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

// SetStorageReliabilityCounter sets the value of StorageReliabilityCounter for the instance
func (instance *MSFT_PhysicalDiskToStorageReliabilityCounter) SetPropertyStorageReliabilityCounter(value MSFT_StorageReliabilityCounter) (err error) {
	return instance.SetProperty("StorageReliabilityCounter", value)
}

// GetStorageReliabilityCounter gets the value of StorageReliabilityCounter for the instance
func (instance *MSFT_PhysicalDiskToStorageReliabilityCounter) GetPropertyStorageReliabilityCounter() (value MSFT_StorageReliabilityCounter, err error) {
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
