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

// MSFT_iSCSISessionToiSCSIConnection struct
type MSFT_iSCSISessionToiSCSIConnection struct {
	cim.WmiInstance

	//
	iSCSIConnection MSFT_iSCSIConnection

	//
	iSCSISession MSFT_iSCSISession
}

// SetiSCSIConnection sets the value of iSCSIConnection for the instance
func (instance *MSFT_iSCSISessionToiSCSIConnection) SetPropertyiSCSIConnection(value MSFT_iSCSIConnection) (err error) {
	return instance.SetProperty("iSCSIConnection", value)
}

// GetiSCSIConnection gets the value of iSCSIConnection for the instance
func (instance *MSFT_iSCSISessionToiSCSIConnection) GetPropertyiSCSIConnection() (value MSFT_iSCSIConnection, err error) {
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

// SetiSCSISession sets the value of iSCSISession for the instance
func (instance *MSFT_iSCSISessionToiSCSIConnection) SetPropertyiSCSISession(value MSFT_iSCSISession) (err error) {
	return instance.SetProperty("iSCSISession", value)
}

// GetiSCSISession gets the value of iSCSISession for the instance
func (instance *MSFT_iSCSISessionToiSCSIConnection) GetPropertyiSCSISession() (value MSFT_iSCSISession, err error) {
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
