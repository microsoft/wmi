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

// CIM_Component struct
type CIM_Component struct {
	cim.WmiInstance

	//
	GroupComponent CIM_ManagedSystemElement

	//
	PartComponent CIM_ManagedSystemElement
}

// SetGroupComponent sets the value of GroupComponent for the instance
func (instance *CIM_Component) SetPropertyGroupComponent(value CIM_ManagedSystemElement) (err error) {
	return instance.SetProperty("GroupComponent", value)
}

// GetGroupComponent gets the value of GroupComponent for the instance
func (instance *CIM_Component) GetPropertyGroupComponent() (value CIM_ManagedSystemElement, err error) {
	retValue, err := instance.GetProperty("GroupComponent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedSystemElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartComponent sets the value of PartComponent for the instance
func (instance *CIM_Component) SetPropertyPartComponent(value CIM_ManagedSystemElement) (err error) {
	return instance.SetProperty("PartComponent", value)
}

// GetPartComponent gets the value of PartComponent for the instance
func (instance *CIM_Component) GetPropertyPartComponent() (value CIM_ManagedSystemElement, err error) {
	retValue, err := instance.GetProperty("PartComponent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedSystemElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
