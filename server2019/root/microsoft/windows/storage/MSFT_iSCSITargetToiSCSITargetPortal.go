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

// MSFT_iSCSITargetToiSCSITargetPortal struct
type MSFT_iSCSITargetToiSCSITargetPortal struct {
	cim.WmiInstance

	//
	iSCSITarget MSFT_iSCSITarget

	//
	iSCSITargetPortal MSFT_iSCSITargetPortal
}

// SetiSCSITarget sets the value of iSCSITarget for the instance
func (instance *MSFT_iSCSITargetToiSCSITargetPortal) SetPropertyiSCSITarget(value MSFT_iSCSITarget) (err error) {
	return instance.SetProperty("iSCSITarget", value)
}

// GetiSCSITarget gets the value of iSCSITarget for the instance
func (instance *MSFT_iSCSITargetToiSCSITargetPortal) GetPropertyiSCSITarget() (value MSFT_iSCSITarget, err error) {
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

// SetiSCSITargetPortal sets the value of iSCSITargetPortal for the instance
func (instance *MSFT_iSCSITargetToiSCSITargetPortal) SetPropertyiSCSITargetPortal(value MSFT_iSCSITargetPortal) (err error) {
	return instance.SetProperty("iSCSITargetPortal", value)
}

// GetiSCSITargetPortal gets the value of iSCSITargetPortal for the instance
func (instance *MSFT_iSCSITargetToiSCSITargetPortal) GetPropertyiSCSITargetPortal() (value MSFT_iSCSITargetPortal, err error) {
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
