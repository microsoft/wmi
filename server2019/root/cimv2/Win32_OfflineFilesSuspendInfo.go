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

// Win32_OfflineFilesSuspendInfo struct
type Win32_OfflineFilesSuspendInfo struct {
	cim.WmiInstance

	//
	Suspended bool

	//
	SuspendedRoot bool
}

// SetSuspended sets the value of Suspended for the instance
func (instance *Win32_OfflineFilesSuspendInfo) SetPropertySuspended(value bool) (err error) {
	return instance.SetProperty("Suspended", value)
}

// GetSuspended gets the value of Suspended for the instance
func (instance *Win32_OfflineFilesSuspendInfo) GetPropertySuspended() (value bool, err error) {
	retValue, err := instance.GetProperty("Suspended")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuspendedRoot sets the value of SuspendedRoot for the instance
func (instance *Win32_OfflineFilesSuspendInfo) SetPropertySuspendedRoot(value bool) (err error) {
	return instance.SetProperty("SuspendedRoot", value)
}

// GetSuspendedRoot gets the value of SuspendedRoot for the instance
func (instance *Win32_OfflineFilesSuspendInfo) GetPropertySuspendedRoot() (value bool, err error) {
	retValue, err := instance.GetProperty("SuspendedRoot")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
