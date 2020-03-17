// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_SecurityPerProcessStatistics struct
type Win32_PerfFormattedData_Counters_SecurityPerProcessStatistics struct {
	Win32_PerfFormattedData

	//
	ContextHandles uint32

	//
	CredentialHandles uint32
}

// SetContextHandles sets the value of ContextHandles for the instance
func (instance *Win32_PerfFormattedData_Counters_SecurityPerProcessStatistics) SetPropertyContextHandles(value uint32) (err error) {
	return instance.SetProperty("ContextHandles", value)
}

// GetContextHandles gets the value of ContextHandles for the instance
func (instance *Win32_PerfFormattedData_Counters_SecurityPerProcessStatistics) GetPropertyContextHandles() (value uint32, err error) {
	retValue, err := instance.GetProperty("ContextHandles")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCredentialHandles sets the value of CredentialHandles for the instance
func (instance *Win32_PerfFormattedData_Counters_SecurityPerProcessStatistics) SetPropertyCredentialHandles(value uint32) (err error) {
	return instance.SetProperty("CredentialHandles", value)
}

// GetCredentialHandles gets the value of CredentialHandles for the instance
func (instance *Win32_PerfFormattedData_Counters_SecurityPerProcessStatistics) GetPropertyCredentialHandles() (value uint32, err error) {
	retValue, err := instance.GetProperty("CredentialHandles")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
