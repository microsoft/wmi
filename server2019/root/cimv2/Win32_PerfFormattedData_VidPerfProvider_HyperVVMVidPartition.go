// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidPartition struct
type Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidPartition struct {
	Win32_PerfFormattedData

	//
	PhysicalPagesAllocated uint64

	//
	PreferredNUMANodeIndex uint64

	//
	RemotePhysicalPages uint64
}

// SetPhysicalPagesAllocated sets the value of PhysicalPagesAllocated for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidPartition) SetPropertyPhysicalPagesAllocated(value uint64) (err error) {
	return instance.SetProperty("PhysicalPagesAllocated", value)
}

// GetPhysicalPagesAllocated gets the value of PhysicalPagesAllocated for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidPartition) GetPropertyPhysicalPagesAllocated() (value uint64, err error) {
	retValue, err := instance.GetProperty("PhysicalPagesAllocated")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreferredNUMANodeIndex sets the value of PreferredNUMANodeIndex for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidPartition) SetPropertyPreferredNUMANodeIndex(value uint64) (err error) {
	return instance.SetProperty("PreferredNUMANodeIndex", value)
}

// GetPreferredNUMANodeIndex gets the value of PreferredNUMANodeIndex for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidPartition) GetPropertyPreferredNUMANodeIndex() (value uint64, err error) {
	retValue, err := instance.GetProperty("PreferredNUMANodeIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemotePhysicalPages sets the value of RemotePhysicalPages for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidPartition) SetPropertyRemotePhysicalPages(value uint64) (err error) {
	return instance.SetProperty("RemotePhysicalPages", value)
}

// GetRemotePhysicalPages gets the value of RemotePhysicalPages for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidPartition) GetPropertyRemotePhysicalPages() (value uint64, err error) {
	retValue, err := instance.GetProperty("RemotePhysicalPages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
