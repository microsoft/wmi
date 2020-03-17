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

// CIM_ElementSetting struct
type CIM_ElementSetting struct {
	cim.WmiInstance

	//
	Element CIM_ManagedSystemElement

	//
	Setting CIM_Setting
}

// SetElement sets the value of Element for the instance
func (instance *CIM_ElementSetting) SetPropertyElement(value CIM_ManagedSystemElement) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *CIM_ElementSetting) GetPropertyElement() (value CIM_ManagedSystemElement, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedSystemElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetting sets the value of Setting for the instance
func (instance *CIM_ElementSetting) SetPropertySetting(value CIM_Setting) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *CIM_ElementSetting) GetPropertySetting() (value CIM_Setting, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Setting)
	if !ok {
		// TODO: Set an error
	}
	return
}
