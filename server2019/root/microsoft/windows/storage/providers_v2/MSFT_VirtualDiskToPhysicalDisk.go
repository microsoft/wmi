// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_VirtualDiskToPhysicalDisk struct
type MSFT_VirtualDiskToPhysicalDisk struct {
	cim.WmiInstance

	//
	PhysicalDisk MSFT_PhysicalDisk

	//
	VirtualDisk MSFT_VirtualDisk
}

// SetPhysicalDisk sets the value of PhysicalDisk for the instance
func (instance *MSFT_VirtualDiskToPhysicalDisk) SetPropertyPhysicalDisk(value MSFT_PhysicalDisk) (err error) {
	return instance.SetProperty("PhysicalDisk", value)
}

// GetPhysicalDisk gets the value of PhysicalDisk for the instance
func (instance *MSFT_VirtualDiskToPhysicalDisk) GetPropertyPhysicalDisk() (value MSFT_PhysicalDisk, err error) {
	retValue, err := instance.GetProperty("PhysicalDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_PhysicalDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDisk sets the value of VirtualDisk for the instance
func (instance *MSFT_VirtualDiskToPhysicalDisk) SetPropertyVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("VirtualDisk", value)
}

// GetVirtualDisk gets the value of VirtualDisk for the instance
func (instance *MSFT_VirtualDiskToPhysicalDisk) GetPropertyVirtualDisk() (value MSFT_VirtualDisk, err error) {
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
