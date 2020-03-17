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

// MSFT_StorageSubSystemToOffloadDataTransferSetting struct
type MSFT_StorageSubSystemToOffloadDataTransferSetting struct {
	cim.WmiInstance

	//
	OffloadDataTransferSetting MSFT_OffloadDataTransferSetting

	//
	StorageSubSystem MSFT_StorageSubSystem
}

// SetOffloadDataTransferSetting sets the value of OffloadDataTransferSetting for the instance
func (instance *MSFT_StorageSubSystemToOffloadDataTransferSetting) SetPropertyOffloadDataTransferSetting(value MSFT_OffloadDataTransferSetting) (err error) {
	return instance.SetProperty("OffloadDataTransferSetting", value)
}

// GetOffloadDataTransferSetting gets the value of OffloadDataTransferSetting for the instance
func (instance *MSFT_StorageSubSystemToOffloadDataTransferSetting) GetPropertyOffloadDataTransferSetting() (value MSFT_OffloadDataTransferSetting, err error) {
	retValue, err := instance.GetProperty("OffloadDataTransferSetting")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_OffloadDataTransferSetting)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToOffloadDataTransferSetting) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToOffloadDataTransferSetting) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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
