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

// CIM_CoolingDevice struct
type CIM_CoolingDevice struct {
	*CIM_LogicalDevice

	//
	ActiveCooling bool
}

func NewCIM_CoolingDeviceEx1(instance *cim.WmiInstance) (newInstance *CIM_CoolingDevice, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_CoolingDevice{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_CoolingDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_CoolingDevice, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_CoolingDevice{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetActiveCooling sets the value of ActiveCooling for the instance
func (instance *CIM_CoolingDevice) SetPropertyActiveCooling(value bool) (err error) {
	return instance.SetProperty("ActiveCooling", value)
}

// GetActiveCooling gets the value of ActiveCooling for the instance
func (instance *CIM_CoolingDevice) GetPropertyActiveCooling() (value bool, err error) {
	retValue, err := instance.GetProperty("ActiveCooling")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
