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

// Win32_PerfRawData_AuthorizationManager_AuthorizationManagerApplications struct
type Win32_PerfRawData_AuthorizationManager_AuthorizationManagerApplications struct {
	*Win32_PerfRawData

	//
	NumberofScopesloadedinmemory uint32

	//
	Totalnumberofscopes uint32
}

func NewWin32_PerfRawData_AuthorizationManager_AuthorizationManagerApplicationsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_AuthorizationManager_AuthorizationManagerApplications, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_AuthorizationManager_AuthorizationManagerApplications{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_AuthorizationManager_AuthorizationManagerApplicationsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_AuthorizationManager_AuthorizationManagerApplications, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_AuthorizationManager_AuthorizationManagerApplications{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetNumberofScopesloadedinmemory sets the value of NumberofScopesloadedinmemory for the instance
func (instance *Win32_PerfRawData_AuthorizationManager_AuthorizationManagerApplications) SetPropertyNumberofScopesloadedinmemory(value uint32) (err error) {
	return instance.SetProperty("NumberofScopesloadedinmemory", value)
}

// GetNumberofScopesloadedinmemory gets the value of NumberofScopesloadedinmemory for the instance
func (instance *Win32_PerfRawData_AuthorizationManager_AuthorizationManagerApplications) GetPropertyNumberofScopesloadedinmemory() (value uint32, err error) {
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
func (instance *Win32_PerfRawData_AuthorizationManager_AuthorizationManagerApplications) SetPropertyTotalnumberofscopes(value uint32) (err error) {
	return instance.SetProperty("Totalnumberofscopes", value)
}

// GetTotalnumberofscopes gets the value of Totalnumberofscopes for the instance
func (instance *Win32_PerfRawData_AuthorizationManager_AuthorizationManagerApplications) GetPropertyTotalnumberofscopes() (value uint32, err error) {
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
