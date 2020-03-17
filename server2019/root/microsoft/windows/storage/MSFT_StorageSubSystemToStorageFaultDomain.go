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

// MSFT_StorageSubSystemToStorageFaultDomain struct
type MSFT_StorageSubSystemToStorageFaultDomain struct {
	cim.WmiInstance

	//
	StorageFaultDomain MSFT_StorageFaultDomain

	//
	StorageSubSystem MSFT_StorageSubSystem
}

// SetStorageFaultDomain sets the value of StorageFaultDomain for the instance
func (instance *MSFT_StorageSubSystemToStorageFaultDomain) SetPropertyStorageFaultDomain(value MSFT_StorageFaultDomain) (err error) {
	return instance.SetProperty("StorageFaultDomain", value)
}

// GetStorageFaultDomain gets the value of StorageFaultDomain for the instance
func (instance *MSFT_StorageSubSystemToStorageFaultDomain) GetPropertyStorageFaultDomain() (value MSFT_StorageFaultDomain, err error) {
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

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageFaultDomain) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageFaultDomain) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
	retValue, err := instance.GetProperty("StorageSubSystem")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageSubSystem)
	if !ok {
		// TODO: Set an error
	}
	return
}
