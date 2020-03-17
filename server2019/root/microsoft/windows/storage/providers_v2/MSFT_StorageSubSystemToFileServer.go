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

// MSFT_StorageSubSystemToFileServer struct
type MSFT_StorageSubSystemToFileServer struct {
	cim.WmiInstance

	//
	FileServer MSFT_FileServer

	//
	StorageSubSystem MSFT_StorageSubSystem
}

// SetFileServer sets the value of FileServer for the instance
func (instance *MSFT_StorageSubSystemToFileServer) SetPropertyFileServer(value MSFT_FileServer) (err error) {
	return instance.SetProperty("FileServer", value)
}

// GetFileServer gets the value of FileServer for the instance
func (instance *MSFT_StorageSubSystemToFileServer) GetPropertyFileServer() (value MSFT_FileServer, err error) {
	retValue, err := instance.GetProperty("FileServer")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_FileServer)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToFileServer) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToFileServer) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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
