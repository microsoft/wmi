// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats struct
type Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats struct {
	Win32_PerfFormattedData

	//
	Missingheartbeats uint32

	//
	Missingheartbeatslimit uint32
}

// SetMissingheartbeats sets the value of Missingheartbeats for the instance
func (instance *Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats) SetPropertyMissingheartbeats(value uint32) (err error) {
	return instance.SetProperty("Missingheartbeats", value)
}

// GetMissingheartbeats gets the value of Missingheartbeats for the instance
func (instance *Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats) GetPropertyMissingheartbeats() (value uint32, err error) {
	retValue, err := instance.GetProperty("Missingheartbeats")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMissingheartbeatslimit sets the value of Missingheartbeatslimit for the instance
func (instance *Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats) SetPropertyMissingheartbeatslimit(value uint32) (err error) {
	return instance.SetProperty("Missingheartbeatslimit", value)
}

// GetMissingheartbeatslimit gets the value of Missingheartbeatslimit for the instance
func (instance *Win32_PerfFormattedData_NetFtPerfProvider_ClusterNetFtHeartbeats) GetPropertyMissingheartbeatslimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("Missingheartbeatslimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
