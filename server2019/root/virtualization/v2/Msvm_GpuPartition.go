// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_GpuPartition struct
type Msvm_GpuPartition struct {
	CIM_LogicalDevice

	//
	DeviceInstancePath string

	//
	PartitionId uint16

	//
	PartitionVfLuid string
}

// SetDeviceInstancePath sets the value of DeviceInstancePath for the instance
func (instance *Msvm_GpuPartition) SetPropertyDeviceInstancePath(value string) (err error) {
	return instance.SetProperty("DeviceInstancePath", value)
}

// GetDeviceInstancePath gets the value of DeviceInstancePath for the instance
func (instance *Msvm_GpuPartition) GetPropertyDeviceInstancePath() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceInstancePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartitionId sets the value of PartitionId for the instance
func (instance *Msvm_GpuPartition) SetPropertyPartitionId(value uint16) (err error) {
	return instance.SetProperty("PartitionId", value)
}

// GetPartitionId gets the value of PartitionId for the instance
func (instance *Msvm_GpuPartition) GetPropertyPartitionId() (value uint16, err error) {
	retValue, err := instance.GetProperty("PartitionId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartitionVfLuid sets the value of PartitionVfLuid for the instance
func (instance *Msvm_GpuPartition) SetPropertyPartitionVfLuid(value string) (err error) {
	return instance.SetProperty("PartitionVfLuid", value)
}

// GetPartitionVfLuid gets the value of PartitionVfLuid for the instance
func (instance *Msvm_GpuPartition) GetPropertyPartitionVfLuid() (value string, err error) {
	retValue, err := instance.GetProperty("PartitionVfLuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
