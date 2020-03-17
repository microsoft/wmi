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

// Win32_OfflineFilesDiskSpaceLimit struct
type Win32_OfflineFilesDiskSpaceLimit struct {
	cim.WmiInstance

	//
	AutoCacheSizeInMB uint32

	//
	TotalSizeInMB uint32
}

// SetAutoCacheSizeInMB sets the value of AutoCacheSizeInMB for the instance
func (instance *Win32_OfflineFilesDiskSpaceLimit) SetPropertyAutoCacheSizeInMB(value uint32) (err error) {
	return instance.SetProperty("AutoCacheSizeInMB", value)
}

// GetAutoCacheSizeInMB gets the value of AutoCacheSizeInMB for the instance
func (instance *Win32_OfflineFilesDiskSpaceLimit) GetPropertyAutoCacheSizeInMB() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoCacheSizeInMB")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalSizeInMB sets the value of TotalSizeInMB for the instance
func (instance *Win32_OfflineFilesDiskSpaceLimit) SetPropertyTotalSizeInMB(value uint32) (err error) {
	return instance.SetProperty("TotalSizeInMB", value)
}

// GetTotalSizeInMB gets the value of TotalSizeInMB for the instance
func (instance *Win32_OfflineFilesDiskSpaceLimit) GetPropertyTotalSizeInMB() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalSizeInMB")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
