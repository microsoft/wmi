// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetServiceCrashNoAction struct
type MSFT_NetServiceCrashNoAction struct {
	MSFT_SCMEventLogEvent

	//
	Service string

	//
	TimesFailed uint32
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetServiceCrashNoAction) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceCrashNoAction) GetPropertyService() (value string, err error) {
	retValue, err := instance.GetProperty("Service")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimesFailed sets the value of TimesFailed for the instance
func (instance *MSFT_NetServiceCrashNoAction) SetPropertyTimesFailed(value uint32) (err error) {
	return instance.SetProperty("TimesFailed", value)
}

// GetTimesFailed gets the value of TimesFailed for the instance
func (instance *MSFT_NetServiceCrashNoAction) GetPropertyTimesFailed() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimesFailed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
