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

// Win32_ShareToDirectory struct
type Win32_ShareToDirectory struct {
	cim.WmiInstance

	//
	Share Win32_Share

	//
	SharedElement CIM_Directory
}

// SetShare sets the value of Share for the instance
func (instance *Win32_ShareToDirectory) SetPropertyShare(value Win32_Share) (err error) {
	return instance.SetProperty("Share", value)
}

// GetShare gets the value of Share for the instance
func (instance *Win32_ShareToDirectory) GetPropertyShare() (value Win32_Share, err error) {
	retValue, err := instance.GetProperty("Share")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Share)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSharedElement sets the value of SharedElement for the instance
func (instance *Win32_ShareToDirectory) SetPropertySharedElement(value CIM_Directory) (err error) {
	return instance.SetProperty("SharedElement", value)
}

// GetSharedElement gets the value of SharedElement for the instance
func (instance *Win32_ShareToDirectory) GetPropertySharedElement() (value CIM_Directory, err error) {
	retValue, err := instance.GetProperty("SharedElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Directory)
	if !ok {
		// TODO: Set an error
	}
	return
}
