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

// Win32_LUIDandAttributes struct
type Win32_LUIDandAttributes struct {
	cim.WmiInstance

	//
	Attributes uint32

	//
	LUID Win32_LUID
}

// SetAttributes sets the value of Attributes for the instance
func (instance *Win32_LUIDandAttributes) SetPropertyAttributes(value uint32) (err error) {
	return instance.SetProperty("Attributes", value)
}

// GetAttributes gets the value of Attributes for the instance
func (instance *Win32_LUIDandAttributes) GetPropertyAttributes() (value uint32, err error) {
	retValue, err := instance.GetProperty("Attributes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLUID sets the value of LUID for the instance
func (instance *Win32_LUIDandAttributes) SetPropertyLUID(value Win32_LUID) (err error) {
	return instance.SetProperty("LUID", value)
}

// GetLUID gets the value of LUID for the instance
func (instance *Win32_LUIDandAttributes) GetPropertyLUID() (value Win32_LUID, err error) {
	retValue, err := instance.GetProperty("LUID")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_LUID)
	if !ok {
		// TODO: Set an error
	}
	return
}
