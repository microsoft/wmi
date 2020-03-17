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

// MSFT_FileServerToFileShare struct
type MSFT_FileServerToFileShare struct {
	cim.WmiInstance

	//
	FileServer MSFT_FileServer

	//
	FileShare MSFT_FileShare
}

// SetFileServer sets the value of FileServer for the instance
func (instance *MSFT_FileServerToFileShare) SetPropertyFileServer(value MSFT_FileServer) (err error) {
	return instance.SetProperty("FileServer", value)
}

// GetFileServer gets the value of FileServer for the instance
func (instance *MSFT_FileServerToFileShare) GetPropertyFileServer() (value MSFT_FileServer, err error) {
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

// SetFileShare sets the value of FileShare for the instance
func (instance *MSFT_FileServerToFileShare) SetPropertyFileShare(value MSFT_FileShare) (err error) {
	return instance.SetProperty("FileShare", value)
}

// GetFileShare gets the value of FileShare for the instance
func (instance *MSFT_FileServerToFileShare) GetPropertyFileShare() (value MSFT_FileShare, err error) {
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
