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

// MSFT_StorageSubSystemToStorageEnclosure struct
type MSFT_StorageSubSystemToStorageEnclosure struct {
	cim.WmiInstance

	//
	StorageEnclosure MSFT_StorageEnclosure

	//
	StorageSubSystem MSFT_StorageSubSystem
}

// SetStorageEnclosure sets the value of StorageEnclosure for the instance
func (instance *MSFT_StorageSubSystemToStorageEnclosure) SetPropertyStorageEnclosure(value MSFT_StorageEnclosure) (err error) {
	return instance.SetProperty("StorageEnclosure", value)
}

// GetStorageEnclosure gets the value of StorageEnclosure for the instance
func (instance *MSFT_StorageSubSystemToStorageEnclosure) GetPropertyStorageEnclosure() (value MSFT_StorageEnclosure, err error) {
	retValue, err := instance.GetProperty("StorageEnclosure")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageEnclosure)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageEnclosure) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageEnclosure) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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
