// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PerfFormattedData_NETFramework_NETCLRSecurity struct
type Win32_PerfFormattedData_NETFramework_NETCLRSecurity struct {
	*Win32_PerfFormattedData

	//
	NumberLinkTimeChecks uint32

	//
	PercentTimeinRTchecks uint32

	//
	PercentTimeSigAuthenticating uint64

	//
	StackWalkDepth uint32

	//
	TotalRuntimeChecks uint32
}

func NewWin32_PerfFormattedData_NETFramework_NETCLRSecurityEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_NETFramework_NETCLRSecurity{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_NETFramework_NETCLRSecurityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_NETFramework_NETCLRSecurity{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetNumberLinkTimeChecks sets the value of NumberLinkTimeChecks for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity) SetPropertyNumberLinkTimeChecks(value uint32) (err error) {
	return instance.SetProperty("NumberLinkTimeChecks", value)
}

// GetNumberLinkTimeChecks gets the value of NumberLinkTimeChecks for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity) GetPropertyNumberLinkTimeChecks() (value uint32, err error) {
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
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity) SetPropertyPercentTimeinRTchecks(value uint32) (err error) {
	return instance.SetProperty("PercentTimeinRTchecks", value)
}

// GetPercentTimeinRTchecks gets the value of PercentTimeinRTchecks for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity) GetPropertyPercentTimeinRTchecks() (value uint32, err error) {
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

// SetPercentTimeSigAuthenticating sets the value of PercentTimeSigAuthenticating for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity) SetPropertyPercentTimeSigAuthenticating(value uint64) (err error) {
	return instance.SetProperty("PercentTimeSigAuthenticating", value)
}

// GetPercentTimeSigAuthenticating gets the value of PercentTimeSigAuthenticating for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity) GetPropertyPercentTimeSigAuthenticating() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity) SetPropertyStackWalkDepth(value uint32) (err error) {
	return instance.SetProperty("StackWalkDepth", value)
}

// GetStackWalkDepth gets the value of StackWalkDepth for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity) GetPropertyStackWalkDepth() (value uint32, err error) {
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
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity) SetPropertyTotalRuntimeChecks(value uint32) (err error) {
	return instance.SetProperty("TotalRuntimeChecks", value)
}

// GetTotalRuntimeChecks gets the value of TotalRuntimeChecks for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRSecurity) GetPropertyTotalRuntimeChecks() (value uint32, err error) {
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
