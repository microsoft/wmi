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

// Win32_LoadOrderGroup struct
type Win32_LoadOrderGroup struct {
	*CIM_LogicalElement

	//
	DriverEnabled bool

	//
	GroupOrder uint32
}

func NewWin32_LoadOrderGroupEx1(instance *cim.WmiInstance) (newInstance *Win32_LoadOrderGroup, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_LoadOrderGroup{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewWin32_LoadOrderGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_LoadOrderGroup, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_LoadOrderGroup{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetDriverEnabled sets the value of DriverEnabled for the instance
func (instance *Win32_LoadOrderGroup) SetPropertyDriverEnabled(value bool) (err error) {
	return instance.SetProperty("DriverEnabled", value)
}

// GetDriverEnabled gets the value of DriverEnabled for the instance
func (instance *Win32_LoadOrderGroup) GetPropertyDriverEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DriverEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGroupOrder sets the value of GroupOrder for the instance
func (instance *Win32_LoadOrderGroup) SetPropertyGroupOrder(value uint32) (err error) {
	return instance.SetProperty("GroupOrder", value)
}

// GetGroupOrder gets the value of GroupOrder for the instance
func (instance *Win32_LoadOrderGroup) GetPropertyGroupOrder() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
