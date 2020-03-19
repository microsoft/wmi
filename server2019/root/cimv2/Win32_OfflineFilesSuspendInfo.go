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

// Win32_OfflineFilesSuspendInfo struct
type Win32_OfflineFilesSuspendInfo struct {
	*cim.WmiInstance

	//
	Suspended bool

	//
	SuspendedRoot bool
}

func NewWin32_OfflineFilesSuspendInfoEx1(instance *cim.WmiInstance) (newInstance *Win32_OfflineFilesSuspendInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_OfflineFilesSuspendInfo{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_OfflineFilesSuspendInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_OfflineFilesSuspendInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_OfflineFilesSuspendInfo{
		WmiInstance: tmp,
	}
	return
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
