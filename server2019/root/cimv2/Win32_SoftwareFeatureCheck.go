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

// Win32_SoftwareFeatureCheck struct
type Win32_SoftwareFeatureCheck struct {
	cim.WmiInstance

	//
	Check CIM_Check

	//
	Element Win32_SoftwareFeature
}

// SetCheck sets the value of Check for the instance
func (instance *Win32_SoftwareFeatureCheck) SetPropertyCheck(value CIM_Check) (err error) {
	return instance.SetProperty("Check", value)
}

// GetCheck gets the value of Check for the instance
func (instance *Win32_SoftwareFeatureCheck) GetPropertyCheck() (value CIM_Check, err error) {
	retValue, err := instance.GetProperty("Check")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Check)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElement sets the value of Element for the instance
func (instance *Win32_SoftwareFeatureCheck) SetPropertyElement(value Win32_SoftwareFeature) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *Win32_SoftwareFeatureCheck) GetPropertyElement() (value Win32_SoftwareFeature, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_SoftwareFeature)
	if !ok {
		// TODO: Set an error
	}
	return
}
