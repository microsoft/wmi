// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_OfflineFilesConnectionInfo struct
type Win32_OfflineFilesConnectionInfo struct {
	cim.WmiInstance

	//
	ConnectState uint32

	//
	OfflineReason uint32
}

// SetConnectState sets the value of ConnectState for the instance
func (instance *Win32_OfflineFilesConnectionInfo) SetPropertyConnectState(value uint32) (err error) {
	return instance.SetProperty("ConnectState", value)
}

// GetConnectState gets the value of ConnectState for the instance
func (instance *Win32_OfflineFilesConnectionInfo) GetPropertyConnectState() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOfflineReason sets the value of OfflineReason for the instance
func (instance *Win32_OfflineFilesConnectionInfo) SetPropertyOfflineReason(value uint32) (err error) {
	return instance.SetProperty("OfflineReason", value)
}

// GetOfflineReason gets the value of OfflineReason for the instance
func (instance *Win32_OfflineFilesConnectionInfo) GetPropertyOfflineReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("OfflineReason")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
