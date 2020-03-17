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

// MSFT_iSCSISessionToiSCSITargetPortal struct
type MSFT_iSCSISessionToiSCSITargetPortal struct {
	cim.WmiInstance

	//
	iSCSISession MSFT_iSCSISession

	//
	iSCSITargetPortal MSFT_iSCSITargetPortal
}

// SetiSCSISession sets the value of iSCSISession for the instance
func (instance *MSFT_iSCSISessionToiSCSITargetPortal) SetPropertyiSCSISession(value MSFT_iSCSISession) (err error) {
	return instance.SetProperty("iSCSISession", value)
}

// GetiSCSISession gets the value of iSCSISession for the instance
func (instance *MSFT_iSCSISessionToiSCSITargetPortal) GetPropertyiSCSISession() (value MSFT_iSCSISession, err error) {
	retValue, err := instance.GetProperty("iSCSISession")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_iSCSISession)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetiSCSITargetPortal sets the value of iSCSITargetPortal for the instance
func (instance *MSFT_iSCSISessionToiSCSITargetPortal) SetPropertyiSCSITargetPortal(value MSFT_iSCSITargetPortal) (err error) {
	return instance.SetProperty("iSCSITargetPortal", value)
}

// GetiSCSITargetPortal gets the value of iSCSITargetPortal for the instance
func (instance *MSFT_iSCSISessionToiSCSITargetPortal) GetPropertyiSCSITargetPortal() (value MSFT_iSCSITargetPortal, err error) {
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
