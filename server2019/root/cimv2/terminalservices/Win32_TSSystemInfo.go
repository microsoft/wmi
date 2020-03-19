// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_TSSystemInfo struct
type Win32_TSSystemInfo struct {
	*CIM_LogicalElement

	// The version number of this WMI Provider
	ProviderVersion uint32

	// The Remote Desktop Users group, in SDDL format
	RDUGroup string
}

func NewWin32_TSSystemInfoEx1(instance *cim.WmiInstance) (newInstance *Win32_TSSystemInfo, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSSystemInfo{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewWin32_TSSystemInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSSystemInfo, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSSystemInfo{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetProviderVersion sets the value of ProviderVersion for the instance
func (instance *Win32_TSSystemInfo) SetPropertyProviderVersion(value uint32) (err error) {
	return instance.SetProperty("ProviderVersion", value)
}

// GetProviderVersion gets the value of ProviderVersion for the instance
func (instance *Win32_TSSystemInfo) GetPropertyProviderVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProviderVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRDUGroup sets the value of RDUGroup for the instance
func (instance *Win32_TSSystemInfo) SetPropertyRDUGroup(value string) (err error) {
	return instance.SetProperty("RDUGroup", value)
}

// GetRDUGroup gets the value of RDUGroup for the instance
func (instance *Win32_TSSystemInfo) GetPropertyRDUGroup() (value string, err error) {
	retValue, err := instance.GetProperty("RDUGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
