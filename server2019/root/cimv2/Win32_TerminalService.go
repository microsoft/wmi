// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_TerminalService struct
type Win32_TerminalService struct {
	Win32_Service

	//
	DisconnectedSessions uint32

	//
	TotalSessions uint32
}

// SetDisconnectedSessions sets the value of DisconnectedSessions for the instance
func (instance *Win32_TerminalService) SetPropertyDisconnectedSessions(value uint32) (err error) {
	return instance.SetProperty("DisconnectedSessions", value)
}

// GetDisconnectedSessions gets the value of DisconnectedSessions for the instance
func (instance *Win32_TerminalService) GetPropertyDisconnectedSessions() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisconnectedSessions")
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
func (instance *Win32_TerminalService) SetPropertyTotalSessions(value uint32) (err error) {
	return instance.SetProperty("TotalSessions", value)
}

// GetTotalSessions gets the value of TotalSessions for the instance
func (instance *Win32_TerminalService) GetPropertyTotalSessions() (value uint32, err error) {
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
