// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_VolumeUserQuota struct
type Win32_VolumeUserQuota struct {
	*cim.WmiInstance

	//
	Account Win32_Account

	//
	DiskSpaceUsed uint64

	//
	Limit uint64

	//
	Status uint32

	//
	Volume Win32_Volume

	//
	WarningLimit uint64
}

func NewWin32_VolumeUserQuotaEx1(instance *cim.WmiInstance) (newInstance *Win32_VolumeUserQuota, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_VolumeUserQuota{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_VolumeUserQuotaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_VolumeUserQuota, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_VolumeUserQuota{
		WmiInstance: tmp,
	}
	return
}

// SetAccount sets the value of Account for the instance
func (instance *Win32_VolumeUserQuota) SetPropertyAccount(value Win32_Account) (err error) {
	return instance.SetProperty("Account", value)
}

// GetAccount gets the value of Account for the instance
func (instance *Win32_VolumeUserQuota) GetPropertyAccount() (value Win32_Account, err error) {
	retValue, err := instance.GetProperty("Account")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Account)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiskSpaceUsed sets the value of DiskSpaceUsed for the instance
func (instance *Win32_VolumeUserQuota) SetPropertyDiskSpaceUsed(value uint64) (err error) {
	return instance.SetProperty("DiskSpaceUsed", value)
}

// GetDiskSpaceUsed gets the value of DiskSpaceUsed for the instance
func (instance *Win32_VolumeUserQuota) GetPropertyDiskSpaceUsed() (value uint64, err error) {
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
func (instance *Win32_VolumeUserQuota) SetPropertyLimit(value uint64) (err error) {
	return instance.SetProperty("Limit", value)
}

// GetLimit gets the value of Limit for the instance
func (instance *Win32_VolumeUserQuota) GetPropertyLimit() (value uint64, err error) {
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

// SetStatus sets the value of Status for the instance
func (instance *Win32_VolumeUserQuota) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *Win32_VolumeUserQuota) GetPropertyStatus() (value uint32, err error) {
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

// SetVolume sets the value of Volume for the instance
func (instance *Win32_VolumeUserQuota) SetPropertyVolume(value Win32_Volume) (err error) {
	return instance.SetProperty("Volume", value)
}

// GetVolume gets the value of Volume for the instance
func (instance *Win32_VolumeUserQuota) GetPropertyVolume() (value Win32_Volume, err error) {
	retValue, err := instance.GetProperty("Volume")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Volume)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWarningLimit sets the value of WarningLimit for the instance
func (instance *Win32_VolumeUserQuota) SetPropertyWarningLimit(value uint64) (err error) {
	return instance.SetProperty("WarningLimit", value)
}

// GetWarningLimit gets the value of WarningLimit for the instance
func (instance *Win32_VolumeUserQuota) GetPropertyWarningLimit() (value uint64, err error) {
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
