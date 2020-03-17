// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_AuthorizationManager_AuthorizationManagerApplications struct
type Win32_PerfFormattedData_AuthorizationManager_AuthorizationManagerApplications struct {
	Win32_PerfFormattedData

	//
	NumberofScopesloadedinmemory uint32

	//
	Totalnumberofscopes uint32
}

// SetNumberofScopesloadedinmemory sets the value of NumberofScopesloadedinmemory for the instance
func (instance *Win32_PerfFormattedData_AuthorizationManager_AuthorizationManagerApplications) SetPropertyNumberofScopesloadedinmemory(value uint32) (err error) {
	return instance.SetProperty("NumberofScopesloadedinmemory", value)
}

// GetNumberofScopesloadedinmemory gets the value of NumberofScopesloadedinmemory for the instance
func (instance *Win32_PerfFormattedData_AuthorizationManager_AuthorizationManagerApplications) GetPropertyNumberofScopesloadedinmemory() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofScopesloadedinmemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalnumberofscopes sets the value of Totalnumberofscopes for the instance
func (instance *Win32_PerfFormattedData_AuthorizationManager_AuthorizationManagerApplications) SetPropertyTotalnumberofscopes(value uint32) (err error) {
	return instance.SetProperty("Totalnumberofscopes", value)
}

// GetTotalnumberofscopes gets the value of Totalnumberofscopes for the instance
func (instance *Win32_PerfFormattedData_AuthorizationManager_AuthorizationManagerApplications) GetPropertyTotalnumberofscopes() (value uint32, err error) {
	retValue, err := instance.GetProperty("Totalnumberofscopes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
