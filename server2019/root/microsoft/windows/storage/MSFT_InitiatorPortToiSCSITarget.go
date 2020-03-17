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

// MSFT_InitiatorPortToiSCSITarget struct
type MSFT_InitiatorPortToiSCSITarget struct {
	cim.WmiInstance

	//
	InitiatorPort MSFT_InitiatorPort

	//
	iSCSITarget MSFT_iSCSITarget
}

// SetInitiatorPort sets the value of InitiatorPort for the instance
func (instance *MSFT_InitiatorPortToiSCSITarget) SetPropertyInitiatorPort(value MSFT_InitiatorPort) (err error) {
	return instance.SetProperty("InitiatorPort", value)
}

// GetInitiatorPort gets the value of InitiatorPort for the instance
func (instance *MSFT_InitiatorPortToiSCSITarget) GetPropertyInitiatorPort() (value MSFT_InitiatorPort, err error) {
	retValue, err := instance.GetProperty("InitiatorPort")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_InitiatorPort)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetiSCSITarget sets the value of iSCSITarget for the instance
func (instance *MSFT_InitiatorPortToiSCSITarget) SetPropertyiSCSITarget(value MSFT_iSCSITarget) (err error) {
	return instance.SetProperty("iSCSITarget", value)
}

// GetiSCSITarget gets the value of iSCSITarget for the instance
func (instance *MSFT_InitiatorPortToiSCSITarget) GetPropertyiSCSITarget() (value MSFT_iSCSITarget, err error) {
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
