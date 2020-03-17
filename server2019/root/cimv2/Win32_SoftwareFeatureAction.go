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

// Win32_SoftwareFeatureAction struct
type Win32_SoftwareFeatureAction struct {
	cim.WmiInstance

	//
	Action CIM_Action

	//
	Element Win32_SoftwareFeature
}

// SetAction sets the value of Action for the instance
func (instance *Win32_SoftwareFeatureAction) SetPropertyAction(value CIM_Action) (err error) {
	return instance.SetProperty("Action", value)
}

// GetAction gets the value of Action for the instance
func (instance *Win32_SoftwareFeatureAction) GetPropertyAction() (value CIM_Action, err error) {
	retValue, err := instance.GetProperty("Action")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Action)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElement sets the value of Element for the instance
func (instance *Win32_SoftwareFeatureAction) SetPropertyElement(value Win32_SoftwareFeature) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *Win32_SoftwareFeatureAction) GetPropertyElement() (value Win32_SoftwareFeature, err error) {
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
