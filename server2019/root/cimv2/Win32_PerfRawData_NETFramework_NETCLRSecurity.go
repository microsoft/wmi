// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_NETFramework_NETCLRSecurity struct
type Win32_PerfRawData_NETFramework_NETCLRSecurity struct {
	Win32_PerfRawData

	//
	NumberLinkTimeChecks uint32

	//
	PercentTimeinRTchecks uint32

	//
	PercentTimeinRTchecks_Base uint32

	//
	PercentTimeSigAuthenticating uint64

	//
	StackWalkDepth uint32

	//
	TotalRuntimeChecks uint32
}

// SetNumberLinkTimeChecks sets the value of NumberLinkTimeChecks for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) SetPropertyNumberLinkTimeChecks(value uint32) (err error) {
	return instance.SetProperty("NumberLinkTimeChecks", value)
}

// GetNumberLinkTimeChecks gets the value of NumberLinkTimeChecks for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) GetPropertyNumberLinkTimeChecks() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberLinkTimeChecks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTimeinRTchecks sets the value of PercentTimeinRTchecks for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) SetPropertyPercentTimeinRTchecks(value uint32) (err error) {
	return instance.SetProperty("PercentTimeinRTchecks", value)
}

// GetPercentTimeinRTchecks gets the value of PercentTimeinRTchecks for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) GetPropertyPercentTimeinRTchecks() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentTimeinRTchecks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTimeinRTchecks_Base sets the value of PercentTimeinRTchecks_Base for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) SetPropertyPercentTimeinRTchecks_Base(value uint32) (err error) {
	return instance.SetProperty("PercentTimeinRTchecks_Base", value)
}

// GetPercentTimeinRTchecks_Base gets the value of PercentTimeinRTchecks_Base for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) GetPropertyPercentTimeinRTchecks_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentTimeinRTchecks_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTimeSigAuthenticating sets the value of PercentTimeSigAuthenticating for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) SetPropertyPercentTimeSigAuthenticating(value uint64) (err error) {
	return instance.SetProperty("PercentTimeSigAuthenticating", value)
}

// GetPercentTimeSigAuthenticating gets the value of PercentTimeSigAuthenticating for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) GetPropertyPercentTimeSigAuthenticating() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentTimeSigAuthenticating")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStackWalkDepth sets the value of StackWalkDepth for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) SetPropertyStackWalkDepth(value uint32) (err error) {
	return instance.SetProperty("StackWalkDepth", value)
}

// GetStackWalkDepth gets the value of StackWalkDepth for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) GetPropertyStackWalkDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("StackWalkDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalRuntimeChecks sets the value of TotalRuntimeChecks for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) SetPropertyTotalRuntimeChecks(value uint32) (err error) {
	return instance.SetProperty("TotalRuntimeChecks", value)
}

// GetTotalRuntimeChecks gets the value of TotalRuntimeChecks for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRSecurity) GetPropertyTotalRuntimeChecks() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalRuntimeChecks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
