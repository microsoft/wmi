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

// MSFT_iSCSIConnectionToiSCSITargetPortal struct
type MSFT_iSCSIConnectionToiSCSITargetPortal struct {
	cim.WmiInstance

	//
	iSCSIConnection MSFT_iSCSIConnection

	//
	iSCSITargetPortal MSFT_iSCSITargetPortal
}

// SetiSCSIConnection sets the value of iSCSIConnection for the instance
func (instance *MSFT_iSCSIConnectionToiSCSITargetPortal) SetPropertyiSCSIConnection(value MSFT_iSCSIConnection) (err error) {
	return instance.SetProperty("iSCSIConnection", value)
}

// GetiSCSIConnection gets the value of iSCSIConnection for the instance
func (instance *MSFT_iSCSIConnectionToiSCSITargetPortal) GetPropertyiSCSIConnection() (value MSFT_iSCSIConnection, err error) {
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

// SetiSCSITargetPortal sets the value of iSCSITargetPortal for the instance
func (instance *MSFT_iSCSIConnectionToiSCSITargetPortal) SetPropertyiSCSITargetPortal(value MSFT_iSCSITargetPortal) (err error) {
	return instance.SetProperty("iSCSITargetPortal", value)
}

// GetiSCSITargetPortal gets the value of iSCSITargetPortal for the instance
func (instance *MSFT_iSCSIConnectionToiSCSITargetPortal) GetPropertyiSCSITargetPortal() (value MSFT_iSCSITargetPortal, err error) {
	retValue, err := instance.GetProperty("iSCSITargetPortal")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_iSCSITargetPortal)
	if !ok {
		// TODO: Set an error
	}
	return
}
