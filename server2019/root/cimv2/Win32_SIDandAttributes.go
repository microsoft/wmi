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

// Win32_SIDandAttributes struct
type Win32_SIDandAttributes struct {
	cim.WmiInstance

	//
	Attributes uint32

	//
	SID Win32_SID
}

// SetAttributes sets the value of Attributes for the instance
func (instance *Win32_SIDandAttributes) SetPropertyAttributes(value uint32) (err error) {
	return instance.SetProperty("Attributes", value)
}

// GetAttributes gets the value of Attributes for the instance
func (instance *Win32_SIDandAttributes) GetPropertyAttributes() (value uint32, err error) {
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

// SetSID sets the value of SID for the instance
func (instance *Win32_SIDandAttributes) SetPropertySID(value Win32_SID) (err error) {
	return instance.SetProperty("SID", value)
}

// GetSID gets the value of SID for the instance
func (instance *Win32_SIDandAttributes) GetPropertySID() (value Win32_SID, err error) {
	retValue, err := instance.GetProperty("SID")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_SID)
	if !ok {
		// TODO: Set an error
	}
	return
}
