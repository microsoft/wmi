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

// Win32_PnPDevicePropertySecurityDescriptor struct
type Win32_PnPDevicePropertySecurityDescriptor struct {
	*Win32_PnPDeviceProperty

	//
	Data Win32_SecurityDescriptor
}

func NewWin32_PnPDevicePropertySecurityDescriptorEx1(instance *cim.WmiInstance) (newInstance *Win32_PnPDevicePropertySecurityDescriptor, err error) {
	tmp, err := NewWin32_PnPDevicePropertyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PnPDevicePropertySecurityDescriptor{
		Win32_PnPDeviceProperty: tmp,
	}
	return
}

func NewWin32_PnPDevicePropertySecurityDescriptorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PnPDevicePropertySecurityDescriptor, err error) {
	tmp, err := NewWin32_PnPDevicePropertyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PnPDevicePropertySecurityDescriptor{
		Win32_PnPDeviceProperty: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertySecurityDescriptor) SetPropertyData(value Win32_SecurityDescriptor) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertySecurityDescriptor) GetPropertyData() (value Win32_SecurityDescriptor, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_SecurityDescriptor)
	if !ok {
		// TODO: Set an error
	}
	return
}
