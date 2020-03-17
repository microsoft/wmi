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

// Win32_TokenGroups struct
type Win32_TokenGroups struct {
	cim.WmiInstance

	//
	GroupCount uint32

	//
	Groups []Win32_SIDandAttributes
}

// SetGroupCount sets the value of GroupCount for the instance
func (instance *Win32_TokenGroups) SetPropertyGroupCount(value uint32) (err error) {
	return instance.SetProperty("GroupCount", value)
}

// GetGroupCount gets the value of GroupCount for the instance
func (instance *Win32_TokenGroups) GetPropertyGroupCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGroups sets the value of Groups for the instance
func (instance *Win32_TokenGroups) SetPropertyGroups(value []Win32_SIDandAttributes) (err error) {
	return instance.SetProperty("Groups", value)
}

// GetGroups gets the value of Groups for the instance
func (instance *Win32_TokenGroups) GetPropertyGroups() (value []Win32_SIDandAttributes, err error) {
	retValue, err := instance.GetProperty("Groups")
	if err != nil {
		return
	}
	value, ok := retValue.([]Win32_SIDandAttributes)
	if !ok {
		// TODO: Set an error
	}
	return
}
