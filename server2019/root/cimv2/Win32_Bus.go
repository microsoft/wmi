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

// Win32_Bus struct
type Win32_Bus struct {
	*CIM_LogicalDevice

	//
	BusNum uint32

	//
	BusType uint32
}

func NewWin32_BusEx1(instance *cim.WmiInstance) (newInstance *Win32_Bus, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_Bus{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewWin32_BusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_Bus, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_Bus{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetBusNum sets the value of BusNum for the instance
func (instance *Win32_Bus) SetPropertyBusNum(value uint32) (err error) {
	return instance.SetProperty("BusNum", value)
}

// GetBusNum gets the value of BusNum for the instance
func (instance *Win32_Bus) GetPropertyBusNum() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusNum")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBusType sets the value of BusType for the instance
func (instance *Win32_Bus) SetPropertyBusType(value uint32) (err error) {
	return instance.SetProperty("BusType", value)
}

// GetBusType gets the value of BusType for the instance
func (instance *Win32_Bus) GetPropertyBusType() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
