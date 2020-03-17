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

// MSFT_iSCSIConnectionToDisk struct
type MSFT_iSCSIConnectionToDisk struct {
	cim.WmiInstance

	//
	Disk MSFT_Disk

	//
	iSCSIConnection MSFT_iSCSIConnection
}

// SetDisk sets the value of Disk for the instance
func (instance *MSFT_iSCSIConnectionToDisk) SetPropertyDisk(value MSFT_Disk) (err error) {
	return instance.SetProperty("Disk", value)
}

// GetDisk gets the value of Disk for the instance
func (instance *MSFT_iSCSIConnectionToDisk) GetPropertyDisk() (value MSFT_Disk, err error) {
	retValue, err := instance.GetProperty("Disk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Disk)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetiSCSIConnection sets the value of iSCSIConnection for the instance
func (instance *MSFT_iSCSIConnectionToDisk) SetPropertyiSCSIConnection(value MSFT_iSCSIConnection) (err error) {
	return instance.SetProperty("iSCSIConnection", value)
}

// GetiSCSIConnection gets the value of iSCSIConnection for the instance
func (instance *MSFT_iSCSIConnectionToDisk) GetPropertyiSCSIConnection() (value MSFT_iSCSIConnection, err error) {
	retValue, err := instance.GetProperty("iSCSIConnection")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_iSCSIConnection)
	if !ok {
		// TODO: Set an error
	}
	return
}
