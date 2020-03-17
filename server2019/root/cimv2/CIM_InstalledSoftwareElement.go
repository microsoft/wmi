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

// CIM_InstalledSoftwareElement struct
type CIM_InstalledSoftwareElement struct {
	cim.WmiInstance

	//
	Software CIM_SoftwareElement

	//
	System CIM_ComputerSystem
}

// SetSoftware sets the value of Software for the instance
func (instance *CIM_InstalledSoftwareElement) SetPropertySoftware(value CIM_SoftwareElement) (err error) {
	return instance.SetProperty("Software", value)
}

// GetSoftware gets the value of Software for the instance
func (instance *CIM_InstalledSoftwareElement) GetPropertySoftware() (value CIM_SoftwareElement, err error) {
	retValue, err := instance.GetProperty("Software")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_SoftwareElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystem sets the value of System for the instance
func (instance *CIM_InstalledSoftwareElement) SetPropertySystem(value CIM_ComputerSystem) (err error) {
	return instance.SetProperty("System", value)
}

// GetSystem gets the value of System for the instance
func (instance *CIM_InstalledSoftwareElement) GetPropertySystem() (value CIM_ComputerSystem, err error) {
	retValue, err := instance.GetProperty("System")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ComputerSystem)
	if !ok {
		// TODO: Set an error
	}
	return
}
