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

// MSFT_StorageSubSystemToFileShare struct
type MSFT_StorageSubSystemToFileShare struct {
	cim.WmiInstance

	//
	FileShare MSFT_FileShare

	//
	StorageSubSystem MSFT_StorageSubSystem
}

// SetFileShare sets the value of FileShare for the instance
func (instance *MSFT_StorageSubSystemToFileShare) SetPropertyFileShare(value MSFT_FileShare) (err error) {
	return instance.SetProperty("FileShare", value)
}

// GetFileShare gets the value of FileShare for the instance
func (instance *MSFT_StorageSubSystemToFileShare) GetPropertyFileShare() (value MSFT_FileShare, err error) {
	retValue, err := instance.GetProperty("FileShare")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_FileShare)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToFileShare) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToFileShare) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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
