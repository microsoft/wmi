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

// MSFT_iSCSITargetToiSCSISession struct
type MSFT_iSCSITargetToiSCSISession struct {
	cim.WmiInstance

	//
	iSCSISession MSFT_iSCSISession

	//
	iSCSITarget MSFT_iSCSITarget
}

// SetiSCSISession sets the value of iSCSISession for the instance
func (instance *MSFT_iSCSITargetToiSCSISession) SetPropertyiSCSISession(value MSFT_iSCSISession) (err error) {
	return instance.SetProperty("iSCSISession", value)
}

// GetiSCSISession gets the value of iSCSISession for the instance
func (instance *MSFT_iSCSITargetToiSCSISession) GetPropertyiSCSISession() (value MSFT_iSCSISession, err error) {
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

// SetiSCSITarget sets the value of iSCSITarget for the instance
func (instance *MSFT_iSCSITargetToiSCSISession) SetPropertyiSCSITarget(value MSFT_iSCSITarget) (err error) {
	return instance.SetProperty("iSCSITarget", value)
}

// GetiSCSITarget gets the value of iSCSITarget for the instance
func (instance *MSFT_iSCSITargetToiSCSISession) GetPropertyiSCSITarget() (value MSFT_iSCSITarget, err error) {
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
