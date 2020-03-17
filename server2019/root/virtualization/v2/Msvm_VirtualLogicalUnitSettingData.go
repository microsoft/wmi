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

// Msvm_VirtualLogicalUnitSettingData struct
type Msvm_VirtualLogicalUnitSettingData struct {
	CIM_ResourceAllocationSettingData

	//
	StorageSubsystemType string
}

// SetStorageSubsystemType sets the value of StorageSubsystemType for the instance
func (instance *Msvm_VirtualLogicalUnitSettingData) SetPropertyStorageSubsystemType(value string) (err error) {
	return instance.SetProperty("StorageSubsystemType", value)
}

// GetStorageSubsystemType gets the value of StorageSubsystemType for the instance
func (instance *Msvm_VirtualLogicalUnitSettingData) GetPropertyStorageSubsystemType() (value string, err error) {
	retValue, err := instance.GetProperty("StorageSubsystemType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_VirtualLogicalUnitSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
