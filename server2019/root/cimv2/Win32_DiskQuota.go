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

// Win32_DiskQuota struct
type Win32_DiskQuota struct {
	cim.WmiInstance

	//
	DiskSpaceUsed uint64

	//
	Limit uint64

	//
	QuotaVolume Win32_LogicalDisk

	//
	Status uint32

	//
	User Win32_Account

	//
	WarningLimit uint64
}

// SetDiskSpaceUsed sets the value of DiskSpaceUsed for the instance
func (instance *Win32_DiskQuota) SetPropertyDiskSpaceUsed(value uint64) (err error) {
	return instance.SetProperty("DiskSpaceUsed", value)
}

// GetDiskSpaceUsed gets the value of DiskSpaceUsed for the instance
func (instance *Win32_DiskQuota) GetPropertyDiskSpaceUsed() (value uint64, err error) {
	retValue, err := instance.GetProperty("DiskSpaceUsed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLimit sets the value of Limit for the instance
func (instance *Win32_DiskQuota) SetPropertyLimit(value uint64) (err error) {
	return instance.SetProperty("Limit", value)
}

// GetLimit gets the value of Limit for the instance
func (instance *Win32_DiskQuota) GetPropertyLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("Limit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuotaVolume sets the value of QuotaVolume for the instance
func (instance *Win32_DiskQuota) SetPropertyQuotaVolume(value Win32_LogicalDisk) (err error) {
	return instance.SetProperty("QuotaVolume", value)
}

// GetQuotaVolume gets the value of QuotaVolume for the instance
func (instance *Win32_DiskQuota) GetPropertyQuotaVolume() (value Win32_LogicalDisk, err error) {
	retValue, err := instance.GetProperty("QuotaVolume")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_LogicalDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *Win32_DiskQuota) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *Win32_DiskQuota) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUser sets the value of User for the instance
func (instance *Win32_DiskQuota) SetPropertyUser(value Win32_Account) (err error) {
	return instance.SetProperty("User", value)
}

// GetUser gets the value of User for the instance
func (instance *Win32_DiskQuota) GetPropertyUser() (value Win32_Account, err error) {
	retValue, err := instance.GetProperty("User")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Account)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWarningLimit sets the value of WarningLimit for the instance
func (instance *Win32_DiskQuota) SetPropertyWarningLimit(value uint64) (err error) {
	return instance.SetProperty("WarningLimit", value)
}

// GetWarningLimit gets the value of WarningLimit for the instance
func (instance *Win32_DiskQuota) GetPropertyWarningLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("WarningLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
