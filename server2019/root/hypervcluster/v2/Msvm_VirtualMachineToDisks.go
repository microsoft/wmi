// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualMachineToDisks struct
type Msvm_VirtualMachineToDisks struct {
	cim.WmiInstance

	//
	DisksToExport []string

	//
	VirtualMachineId string
}

// SetDisksToExport sets the value of DisksToExport for the instance
func (instance *Msvm_VirtualMachineToDisks) SetPropertyDisksToExport(value []string) (err error) {
	return instance.SetProperty("DisksToExport", value)
}

// GetDisksToExport gets the value of DisksToExport for the instance
func (instance *Msvm_VirtualMachineToDisks) GetPropertyDisksToExport() (value []string, err error) {
	retValue, err := instance.GetProperty("DisksToExport")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualMachineId sets the value of VirtualMachineId for the instance
func (instance *Msvm_VirtualMachineToDisks) SetPropertyVirtualMachineId(value string) (err error) {
	return instance.SetProperty("VirtualMachineId", value)
}

// GetVirtualMachineId gets the value of VirtualMachineId for the instance
func (instance *Msvm_VirtualMachineToDisks) GetPropertyVirtualMachineId() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualMachineId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
