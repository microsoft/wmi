// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_LocalSessionManager_TerminalServices struct
type Win32_PerfRawData_LocalSessionManager_TerminalServices struct {
	Win32_PerfRawData

	//
	ActiveSessions uint32

	//
	InactiveSessions uint32

	//
	TotalSessions uint32
}

// SetActiveSessions sets the value of ActiveSessions for the instance
func (instance *Win32_PerfRawData_LocalSessionManager_TerminalServices) SetPropertyActiveSessions(value uint32) (err error) {
	return instance.SetProperty("ActiveSessions", value)
}

// GetActiveSessions gets the value of ActiveSessions for the instance
func (instance *Win32_PerfRawData_LocalSessionManager_TerminalServices) GetPropertyActiveSessions() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveSessions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInactiveSessions sets the value of InactiveSessions for the instance
func (instance *Win32_PerfRawData_LocalSessionManager_TerminalServices) SetPropertyInactiveSessions(value uint32) (err error) {
	return instance.SetProperty("InactiveSessions", value)
}

// GetInactiveSessions gets the value of InactiveSessions for the instance
func (instance *Win32_PerfRawData_LocalSessionManager_TerminalServices) GetPropertyInactiveSessions() (value uint32, err error) {
	retValue, err := instance.GetProperty("InactiveSessions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalSessions sets the value of TotalSessions for the instance
func (instance *Win32_PerfRawData_LocalSessionManager_TerminalServices) SetPropertyTotalSessions(value uint32) (err error) {
	return instance.SetProperty("TotalSessions", value)
}

// GetTotalSessions gets the value of TotalSessions for the instance
func (instance *Win32_PerfRawData_LocalSessionManager_TerminalServices) GetPropertyTotalSessions() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalSessions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
