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

// MSFT_iSCSITargetToiSCSIConnection struct
type MSFT_iSCSITargetToiSCSIConnection struct {
	cim.WmiInstance

	//
	iSCSIConnection MSFT_iSCSIConnection

	//
	iSCSITarget MSFT_iSCSITarget
}

// SetiSCSIConnection sets the value of iSCSIConnection for the instance
func (instance *MSFT_iSCSITargetToiSCSIConnection) SetPropertyiSCSIConnection(value MSFT_iSCSIConnection) (err error) {
	return instance.SetProperty("iSCSIConnection", value)
}

// GetiSCSIConnection gets the value of iSCSIConnection for the instance
func (instance *MSFT_iSCSITargetToiSCSIConnection) GetPropertyiSCSIConnection() (value MSFT_iSCSIConnection, err error) {
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

// SetiSCSITarget sets the value of iSCSITarget for the instance
func (instance *MSFT_iSCSITargetToiSCSIConnection) SetPropertyiSCSITarget(value MSFT_iSCSITarget) (err error) {
	return instance.SetProperty("iSCSITarget", value)
}

// GetiSCSITarget gets the value of iSCSITarget for the instance
func (instance *MSFT_iSCSITargetToiSCSIConnection) GetPropertyiSCSITarget() (value MSFT_iSCSITarget, err error) {
	retValue, err := instance.GetProperty("iSCSITarget")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_iSCSITarget)
	if !ok {
		// TODO: Set an error
	}
	return
}
