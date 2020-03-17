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

// Win32_OfflineFilesDirtyInfo struct
type Win32_OfflineFilesDirtyInfo struct {
	cim.WmiInstance

	//
	LocalDirtyByteCount int64

	//
	RemoteDirtyByteCount int64
}

// SetLocalDirtyByteCount sets the value of LocalDirtyByteCount for the instance
func (instance *Win32_OfflineFilesDirtyInfo) SetPropertyLocalDirtyByteCount(value int64) (err error) {
	return instance.SetProperty("LocalDirtyByteCount", value)
}

// GetLocalDirtyByteCount gets the value of LocalDirtyByteCount for the instance
func (instance *Win32_OfflineFilesDirtyInfo) GetPropertyLocalDirtyByteCount() (value int64, err error) {
	retValue, err := instance.GetProperty("LocalDirtyByteCount")
	if err != nil {
		return
	}
	value, ok := retValue.(int64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteDirtyByteCount sets the value of RemoteDirtyByteCount for the instance
func (instance *Win32_OfflineFilesDirtyInfo) SetPropertyRemoteDirtyByteCount(value int64) (err error) {
	return instance.SetProperty("RemoteDirtyByteCount", value)
}

// GetRemoteDirtyByteCount gets the value of RemoteDirtyByteCount for the instance
func (instance *Win32_OfflineFilesDirtyInfo) GetPropertyRemoteDirtyByteCount() (value int64, err error) {
	retValue, err := instance.GetProperty("RemoteDirtyByteCount")
	if err != nil {
		return
	}
	value, ok := retValue.(int64)
	if !ok {
		// TODO: Set an error
	}
	return
}
