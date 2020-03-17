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

// CIM_SoftwareElementChecks struct
type CIM_SoftwareElementChecks struct {
	cim.WmiInstance

	//
	Check CIM_Check

	//
	Element CIM_SoftwareElement

	//
	Phase uint16
}

// SetCheck sets the value of Check for the instance
func (instance *CIM_SoftwareElementChecks) SetPropertyCheck(value CIM_Check) (err error) {
	return instance.SetProperty("Check", value)
}

// GetCheck gets the value of Check for the instance
func (instance *CIM_SoftwareElementChecks) GetPropertyCheck() (value CIM_Check, err error) {
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
func (instance *CIM_SoftwareElementChecks) SetPropertyElement(value CIM_SoftwareElement) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *CIM_SoftwareElementChecks) GetPropertyElement() (value CIM_SoftwareElement, err error) {
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

// SetPhase sets the value of Phase for the instance
func (instance *CIM_SoftwareElementChecks) SetPropertyPhase(value uint16) (err error) {
	return instance.SetProperty("Phase", value)
}

// GetPhase gets the value of Phase for the instance
func (instance *CIM_SoftwareElementChecks) GetPropertyPhase() (value uint16, err error) {
	retValue, err := instance.GetProperty("Phase")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
