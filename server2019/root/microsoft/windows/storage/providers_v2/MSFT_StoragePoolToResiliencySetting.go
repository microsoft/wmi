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

// MSFT_StoragePoolToResiliencySetting struct
type MSFT_StoragePoolToResiliencySetting struct {
	cim.WmiInstance

	//
	ResiliencySetting MSFT_ResiliencySetting

	//
	StoragePool MSFT_StoragePool
}

// SetResiliencySetting sets the value of ResiliencySetting for the instance
func (instance *MSFT_StoragePoolToResiliencySetting) SetPropertyResiliencySetting(value MSFT_ResiliencySetting) (err error) {
	return instance.SetProperty("ResiliencySetting", value)
}

// GetResiliencySetting gets the value of ResiliencySetting for the instance
func (instance *MSFT_StoragePoolToResiliencySetting) GetPropertyResiliencySetting() (value MSFT_ResiliencySetting, err error) {
	retValue, err := instance.GetProperty("ResiliencySetting")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_ResiliencySetting)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStoragePool sets the value of StoragePool for the instance
func (instance *MSFT_StoragePoolToResiliencySetting) SetPropertyStoragePool(value MSFT_StoragePool) (err error) {
	return instance.SetProperty("StoragePool", value)
}

// GetStoragePool gets the value of StoragePool for the instance
func (instance *MSFT_StoragePoolToResiliencySetting) GetPropertyStoragePool() (value MSFT_StoragePool, err error) {
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
