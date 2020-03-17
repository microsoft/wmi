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

// MSFT_StorageProviderToStorageSubSystem struct
type MSFT_StorageProviderToStorageSubSystem struct {
	cim.WmiInstance

	//
	StorageProvider MSFT_StorageProvider

	//
	StorageSubSystem MSFT_StorageSubSystem
}

// SetStorageProvider sets the value of StorageProvider for the instance
func (instance *MSFT_StorageProviderToStorageSubSystem) SetPropertyStorageProvider(value MSFT_StorageProvider) (err error) {
	return instance.SetProperty("StorageProvider", value)
}

// GetStorageProvider gets the value of StorageProvider for the instance
func (instance *MSFT_StorageProviderToStorageSubSystem) GetPropertyStorageProvider() (value MSFT_StorageProvider, err error) {
	retValue, err := instance.GetProperty("StorageProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageProviderToStorageSubSystem) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageProviderToStorageSubSystem) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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
