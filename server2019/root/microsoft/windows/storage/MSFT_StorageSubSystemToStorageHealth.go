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

// MSFT_StorageSubSystemToStorageHealth struct
type MSFT_StorageSubSystemToStorageHealth struct {
	cim.WmiInstance

	//
	StorageHealth MSFT_StorageHealth

	//
	StorageSubSystem MSFT_StorageSubSystem
}

// SetStorageHealth sets the value of StorageHealth for the instance
func (instance *MSFT_StorageSubSystemToStorageHealth) SetPropertyStorageHealth(value MSFT_StorageHealth) (err error) {
	return instance.SetProperty("StorageHealth", value)
}

// GetStorageHealth gets the value of StorageHealth for the instance
func (instance *MSFT_StorageSubSystemToStorageHealth) GetPropertyStorageHealth() (value MSFT_StorageHealth, err error) {
	retValue, err := instance.GetProperty("StorageHealth")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageHealth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageHealth) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageHealth) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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
