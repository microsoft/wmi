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

// Win32_InfraredDevice struct
type Win32_InfraredDevice struct {
	*CIM_InfraredController

	//
	Manufacturer string
}

func NewWin32_InfraredDeviceEx1(instance *cim.WmiInstance) (newInstance *Win32_InfraredDevice, err error) {
	tmp, err := NewCIM_InfraredControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_InfraredDevice{
		CIM_InfraredController: tmp,
	}
	return
}

func NewWin32_InfraredDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_InfraredDevice, err error) {
	tmp, err := NewCIM_InfraredControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_InfraredDevice{
		CIM_InfraredController: tmp,
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *Win32_InfraredDevice) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *Win32_InfraredDevice) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
