// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_TargetPortToVirtualDisk struct
type MSFT_TargetPortToVirtualDisk struct {
	*cim.WmiInstance

	//
	TargetPort MSFT_TargetPort

	//
	VirtualDisk MSFT_VirtualDisk
}

func NewMSFT_TargetPortToVirtualDiskEx1(instance *cim.WmiInstance) (newInstance *MSFT_TargetPortToVirtualDisk, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_TargetPortToVirtualDisk{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_TargetPortToVirtualDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TargetPortToVirtualDisk, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TargetPortToVirtualDisk{
		WmiInstance: tmp,
	}
	return
}

// SetTargetPort sets the value of TargetPort for the instance
func (instance *MSFT_TargetPortToVirtualDisk) SetPropertyTargetPort(value MSFT_TargetPort) (err error) {
	return instance.SetProperty("TargetPort", value)
}

// GetTargetPort gets the value of TargetPort for the instance
func (instance *MSFT_TargetPortToVirtualDisk) GetPropertyTargetPort() (value MSFT_TargetPort, err error) {
	retValue, err := instance.GetProperty("TargetPort")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_TargetPort)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDisk sets the value of VirtualDisk for the instance
func (instance *MSFT_TargetPortToVirtualDisk) SetPropertyVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("VirtualDisk", value)
}

// GetVirtualDisk gets the value of VirtualDisk for the instance
func (instance *MSFT_TargetPortToVirtualDisk) GetPropertyVirtualDisk() (value MSFT_VirtualDisk, err error) {
	retValue, err := instance.GetProperty("VirtualDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_VirtualDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}
