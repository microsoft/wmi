// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_ClussvcPerfProvider_ClusterResourceControlManager struct
type Win32_PerfFormattedData_ClussvcPerfProvider_ClusterResourceControlManager struct {
	Win32_PerfFormattedData

	//
	GroupsOnline uint64

	//
	RHSProcesses uint64

	//
	RHSRestarts uint64
}

// SetGroupsOnline sets the value of GroupsOnline for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterResourceControlManager) SetPropertyGroupsOnline(value uint64) (err error) {
	return instance.SetProperty("GroupsOnline", value)
}

// GetGroupsOnline gets the value of GroupsOnline for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterResourceControlManager) GetPropertyGroupsOnline() (value uint64, err error) {
	retValue, err := instance.GetProperty("GroupsOnline")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRHSProcesses sets the value of RHSProcesses for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterResourceControlManager) SetPropertyRHSProcesses(value uint64) (err error) {
	return instance.SetProperty("RHSProcesses", value)
}

// GetRHSProcesses gets the value of RHSProcesses for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterResourceControlManager) GetPropertyRHSProcesses() (value uint64, err error) {
	retValue, err := instance.GetProperty("RHSProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRHSRestarts sets the value of RHSRestarts for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterResourceControlManager) SetPropertyRHSRestarts(value uint64) (err error) {
	return instance.SetProperty("RHSRestarts", value)
}

// GetRHSRestarts gets the value of RHSRestarts for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterResourceControlManager) GetPropertyRHSRestarts() (value uint64, err error) {
	retValue, err := instance.GetProperty("RHSRestarts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
