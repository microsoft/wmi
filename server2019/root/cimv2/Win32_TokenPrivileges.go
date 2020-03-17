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

// Win32_TokenPrivileges struct
type Win32_TokenPrivileges struct {
	cim.WmiInstance

	//
	PrivilegeCount uint32

	//
	Privileges []Win32_LUIDandAttributes
}

// SetPrivilegeCount sets the value of PrivilegeCount for the instance
func (instance *Win32_TokenPrivileges) SetPropertyPrivilegeCount(value uint32) (err error) {
	return instance.SetProperty("PrivilegeCount", value)
}

// GetPrivilegeCount gets the value of PrivilegeCount for the instance
func (instance *Win32_TokenPrivileges) GetPropertyPrivilegeCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("PrivilegeCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrivileges sets the value of Privileges for the instance
func (instance *Win32_TokenPrivileges) SetPropertyPrivileges(value []Win32_LUIDandAttributes) (err error) {
	return instance.SetProperty("Privileges", value)
}

// GetPrivileges gets the value of Privileges for the instance
func (instance *Win32_TokenPrivileges) GetPropertyPrivileges() (value []Win32_LUIDandAttributes, err error) {
	retValue, err := instance.GetProperty("Privileges")
	if err != nil {
		return
	}
	value, ok := retValue.([]Win32_LUIDandAttributes)
	if !ok {
		// TODO: Set an error
	}
	return
}
