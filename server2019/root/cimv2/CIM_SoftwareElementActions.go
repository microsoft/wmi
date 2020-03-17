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

// CIM_SoftwareElementActions struct
type CIM_SoftwareElementActions struct {
	cim.WmiInstance

	//
	Action CIM_Action

	//
	Element CIM_SoftwareElement
}

// SetAction sets the value of Action for the instance
func (instance *CIM_SoftwareElementActions) SetPropertyAction(value CIM_Action) (err error) {
	return instance.SetProperty("Action", value)
}

// GetAction gets the value of Action for the instance
func (instance *CIM_SoftwareElementActions) GetPropertyAction() (value CIM_Action, err error) {
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
func (instance *CIM_SoftwareElementActions) SetPropertyElement(value CIM_SoftwareElement) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *CIM_SoftwareElementActions) GetPropertyElement() (value CIM_SoftwareElement, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_SoftwareElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
