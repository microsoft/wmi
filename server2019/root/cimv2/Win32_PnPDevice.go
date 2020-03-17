// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PnPDevice struct
type Win32_PnPDevice struct {
	cim.WmiInstance

	//
	SameElement CIM_LogicalDevice

	//
	SystemElement Win32_PnPEntity
}

// SetSameElement sets the value of SameElement for the instance
func (instance *Win32_PnPDevice) SetPropertySameElement(value CIM_LogicalDevice) (err error) {
	return instance.SetProperty("SameElement", value)
}

// GetSameElement gets the value of SameElement for the instance
func (instance *Win32_PnPDevice) GetPropertySameElement() (value CIM_LogicalDevice, err error) {
	retValue, err := instance.GetProperty("SameElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_LogicalDevice)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemElement sets the value of SystemElement for the instance
func (instance *Win32_PnPDevice) SetPropertySystemElement(value Win32_PnPEntity) (err error) {
	return instance.SetProperty("SystemElement", value)
}

// GetSystemElement gets the value of SystemElement for the instance
func (instance *Win32_PnPDevice) GetPropertySystemElement() (value Win32_PnPEntity, err error) {
	retValue, err := instance.GetProperty("SystemElement")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_PnPEntity)
	if !ok {
		// TODO: Set an error
	}
	return
}
