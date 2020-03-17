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

// Msvm_GpuPartitionSettingData struct
type Msvm_GpuPartitionSettingData struct {
	CIM_ResourceAllocationSettingData

	//
	MaxPartitionCompute uint64

	//
	MaxPartitionDecode uint64

	//
	MaxPartitionEncode uint64

	//
	MaxPartitionVRAM uint64

	//
	MinPartitionCompute uint64

	//
	MinPartitionDecode uint64

	//
	MinPartitionEncode uint64

	//
	MinPartitionVRAM uint64

	//
	OptimalPartitionCompute uint64

	//
	OptimalPartitionDecode uint64

	//
	OptimalPartitionEncode uint64

	//
	OptimalPartitionVRAM uint64

	//
	VirtualSystemIdentifiers []string
}

// SetMaxPartitionCompute sets the value of MaxPartitionCompute for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyMaxPartitionCompute(value uint64) (err error) {
	return instance.SetProperty("MaxPartitionCompute", value)
}

// GetMaxPartitionCompute gets the value of MaxPartitionCompute for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyMaxPartitionCompute() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxPartitionCompute")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxPartitionDecode sets the value of MaxPartitionDecode for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyMaxPartitionDecode(value uint64) (err error) {
	return instance.SetProperty("MaxPartitionDecode", value)
}

// GetMaxPartitionDecode gets the value of MaxPartitionDecode for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyMaxPartitionDecode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxPartitionDecode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxPartitionEncode sets the value of MaxPartitionEncode for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyMaxPartitionEncode(value uint64) (err error) {
	return instance.SetProperty("MaxPartitionEncode", value)
}

// GetMaxPartitionEncode gets the value of MaxPartitionEncode for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyMaxPartitionEncode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxPartitionEncode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxPartitionVRAM sets the value of MaxPartitionVRAM for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyMaxPartitionVRAM(value uint64) (err error) {
	return instance.SetProperty("MaxPartitionVRAM", value)
}

// GetMaxPartitionVRAM gets the value of MaxPartitionVRAM for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyMaxPartitionVRAM() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxPartitionVRAM")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinPartitionCompute sets the value of MinPartitionCompute for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyMinPartitionCompute(value uint64) (err error) {
	return instance.SetProperty("MinPartitionCompute", value)
}

// GetMinPartitionCompute gets the value of MinPartitionCompute for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyMinPartitionCompute() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinPartitionCompute")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinPartitionDecode sets the value of MinPartitionDecode for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyMinPartitionDecode(value uint64) (err error) {
	return instance.SetProperty("MinPartitionDecode", value)
}

// GetMinPartitionDecode gets the value of MinPartitionDecode for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyMinPartitionDecode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinPartitionDecode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinPartitionEncode sets the value of MinPartitionEncode for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyMinPartitionEncode(value uint64) (err error) {
	return instance.SetProperty("MinPartitionEncode", value)
}

// GetMinPartitionEncode gets the value of MinPartitionEncode for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyMinPartitionEncode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinPartitionEncode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinPartitionVRAM sets the value of MinPartitionVRAM for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyMinPartitionVRAM(value uint64) (err error) {
	return instance.SetProperty("MinPartitionVRAM", value)
}

// GetMinPartitionVRAM gets the value of MinPartitionVRAM for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyMinPartitionVRAM() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinPartitionVRAM")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptimalPartitionCompute sets the value of OptimalPartitionCompute for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyOptimalPartitionCompute(value uint64) (err error) {
	return instance.SetProperty("OptimalPartitionCompute", value)
}

// GetOptimalPartitionCompute gets the value of OptimalPartitionCompute for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyOptimalPartitionCompute() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimalPartitionCompute")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptimalPartitionDecode sets the value of OptimalPartitionDecode for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyOptimalPartitionDecode(value uint64) (err error) {
	return instance.SetProperty("OptimalPartitionDecode", value)
}

// GetOptimalPartitionDecode gets the value of OptimalPartitionDecode for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyOptimalPartitionDecode() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimalPartitionDecode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptimalPartitionEncode sets the value of OptimalPartitionEncode for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyOptimalPartitionEncode(value uint64) (err error) {
	return instance.SetProperty("OptimalPartitionEncode", value)
}

// GetOptimalPartitionEncode gets the value of OptimalPartitionEncode for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyOptimalPartitionEncode() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimalPartitionEncode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptimalPartitionVRAM sets the value of OptimalPartitionVRAM for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyOptimalPartitionVRAM(value uint64) (err error) {
	return instance.SetProperty("OptimalPartitionVRAM", value)
}

// GetOptimalPartitionVRAM gets the value of OptimalPartitionVRAM for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyOptimalPartitionVRAM() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimalPartitionVRAM")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemIdentifiers sets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_GpuPartitionSettingData) SetPropertyVirtualSystemIdentifiers(value []string) (err error) {
	return instance.SetProperty("VirtualSystemIdentifiers", value)
}

// GetVirtualSystemIdentifiers gets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_GpuPartitionSettingData) GetPropertyVirtualSystemIdentifiers() (value []string, err error) {
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
func (instance *Msvm_GpuPartitionSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
