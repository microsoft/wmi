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

// Msvm_VirtualSystemMigrationSettingData struct
type Msvm_VirtualSystemMigrationSettingData struct {
	CIM_VirtualSystemMigrationSettingData

	//
	AllowOverwriteExistingFile bool

	//
	AvoidRemovingVHDs bool

	//
	CancelIfBlackoutThresholdExceeded bool

	//
	CPUCappingMagnitude VirtualSystemMigrationSettingData_CPUCappingMagnitude

	//
	DestinationIPAddressList []string

	//
	DestinationPlannedVirtualSystemId string

	//
	EnableCompression bool

	//
	EnableVhdCompression bool

	//
	RemoveSourceUnmanagedVhds bool

	//
	RetainVhdCopiesOnSource bool

	//
	UnmanagedVhds []string
}

// SetAllowOverwriteExistingFile sets the value of AllowOverwriteExistingFile for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyAllowOverwriteExistingFile(value bool) (err error) {
	return instance.SetProperty("AllowOverwriteExistingFile", value)
}

// GetAllowOverwriteExistingFile gets the value of AllowOverwriteExistingFile for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyAllowOverwriteExistingFile() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowOverwriteExistingFile")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvoidRemovingVHDs sets the value of AvoidRemovingVHDs for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyAvoidRemovingVHDs(value bool) (err error) {
	return instance.SetProperty("AvoidRemovingVHDs", value)
}

// GetAvoidRemovingVHDs gets the value of AvoidRemovingVHDs for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyAvoidRemovingVHDs() (value bool, err error) {
	retValue, err := instance.GetProperty("AvoidRemovingVHDs")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCancelIfBlackoutThresholdExceeded sets the value of CancelIfBlackoutThresholdExceeded for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyCancelIfBlackoutThresholdExceeded(value bool) (err error) {
	return instance.SetProperty("CancelIfBlackoutThresholdExceeded", value)
}

// GetCancelIfBlackoutThresholdExceeded gets the value of CancelIfBlackoutThresholdExceeded for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyCancelIfBlackoutThresholdExceeded() (value bool, err error) {
	retValue, err := instance.GetProperty("CancelIfBlackoutThresholdExceeded")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCPUCappingMagnitude sets the value of CPUCappingMagnitude for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyCPUCappingMagnitude(value VirtualSystemMigrationSettingData_CPUCappingMagnitude) (err error) {
	return instance.SetProperty("CPUCappingMagnitude", value)
}

// GetCPUCappingMagnitude gets the value of CPUCappingMagnitude for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyCPUCappingMagnitude() (value VirtualSystemMigrationSettingData_CPUCappingMagnitude, err error) {
	retValue, err := instance.GetProperty("CPUCappingMagnitude")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualSystemMigrationSettingData_CPUCappingMagnitude)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestinationIPAddressList sets the value of DestinationIPAddressList for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyDestinationIPAddressList(value []string) (err error) {
	return instance.SetProperty("DestinationIPAddressList", value)
}

// GetDestinationIPAddressList gets the value of DestinationIPAddressList for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyDestinationIPAddressList() (value []string, err error) {
	retValue, err := instance.GetProperty("DestinationIPAddressList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestinationPlannedVirtualSystemId sets the value of DestinationPlannedVirtualSystemId for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyDestinationPlannedVirtualSystemId(value string) (err error) {
	return instance.SetProperty("DestinationPlannedVirtualSystemId", value)
}

// GetDestinationPlannedVirtualSystemId gets the value of DestinationPlannedVirtualSystemId for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyDestinationPlannedVirtualSystemId() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationPlannedVirtualSystemId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableCompression sets the value of EnableCompression for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyEnableCompression(value bool) (err error) {
	return instance.SetProperty("EnableCompression", value)
}

// GetEnableCompression gets the value of EnableCompression for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyEnableCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableCompression")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableVhdCompression sets the value of EnableVhdCompression for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyEnableVhdCompression(value bool) (err error) {
	return instance.SetProperty("EnableVhdCompression", value)
}

// GetEnableVhdCompression gets the value of EnableVhdCompression for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyEnableVhdCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableVhdCompression")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoveSourceUnmanagedVhds sets the value of RemoveSourceUnmanagedVhds for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyRemoveSourceUnmanagedVhds(value bool) (err error) {
	return instance.SetProperty("RemoveSourceUnmanagedVhds", value)
}

// GetRemoveSourceUnmanagedVhds gets the value of RemoveSourceUnmanagedVhds for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyRemoveSourceUnmanagedVhds() (value bool, err error) {
	retValue, err := instance.GetProperty("RemoveSourceUnmanagedVhds")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRetainVhdCopiesOnSource sets the value of RetainVhdCopiesOnSource for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyRetainVhdCopiesOnSource(value bool) (err error) {
	return instance.SetProperty("RetainVhdCopiesOnSource", value)
}

// GetRetainVhdCopiesOnSource gets the value of RetainVhdCopiesOnSource for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyRetainVhdCopiesOnSource() (value bool, err error) {
	retValue, err := instance.GetProperty("RetainVhdCopiesOnSource")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnmanagedVhds sets the value of UnmanagedVhds for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) SetPropertyUnmanagedVhds(value []string) (err error) {
	return instance.SetProperty("UnmanagedVhds", value)
}

// GetUnmanagedVhds gets the value of UnmanagedVhds for the instance
func (instance *Msvm_VirtualSystemMigrationSettingData) GetPropertyUnmanagedVhds() (value []string, err error) {
	retValue, err := instance.GetProperty("UnmanagedVhds")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_VirtualSystemMigrationSettingData) GetRelatedVirtualSystemMigrationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemMigrationCapabilities")
}
