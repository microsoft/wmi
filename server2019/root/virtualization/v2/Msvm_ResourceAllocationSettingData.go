// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ResourceAllocationSettingData struct
type Msvm_ResourceAllocationSettingData struct {
	CIM_ResourceAllocationSettingData

	//
	VirtualSystemIdentifiers []string
}

// SetVirtualSystemIdentifiers sets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_ResourceAllocationSettingData) SetPropertyVirtualSystemIdentifiers(value []string) (err error) {
	return instance.SetProperty("VirtualSystemIdentifiers", value)
}

// GetVirtualSystemIdentifiers gets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_ResourceAllocationSettingData) GetPropertyVirtualSystemIdentifiers() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemIdentifiers")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_ResourceAllocationSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
