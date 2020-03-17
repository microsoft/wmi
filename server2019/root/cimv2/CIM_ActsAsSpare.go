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

// CIM_ActsAsSpare struct
type CIM_ActsAsSpare struct {
	cim.WmiInstance

	//
	Group CIM_SpareGroup

	//
	HotStandby bool

	//
	Spare CIM_ManagedSystemElement
}

// SetGroup sets the value of Group for the instance
func (instance *CIM_ActsAsSpare) SetPropertyGroup(value CIM_SpareGroup) (err error) {
	return instance.SetProperty("Group", value)
}

// GetGroup gets the value of Group for the instance
func (instance *CIM_ActsAsSpare) GetPropertyGroup() (value CIM_SpareGroup, err error) {
	retValue, err := instance.GetProperty("Group")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_SpareGroup)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHotStandby sets the value of HotStandby for the instance
func (instance *CIM_ActsAsSpare) SetPropertyHotStandby(value bool) (err error) {
	return instance.SetProperty("HotStandby", value)
}

// GetHotStandby gets the value of HotStandby for the instance
func (instance *CIM_ActsAsSpare) GetPropertyHotStandby() (value bool, err error) {
	retValue, err := instance.GetProperty("HotStandby")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpare sets the value of Spare for the instance
func (instance *CIM_ActsAsSpare) SetPropertySpare(value CIM_ManagedSystemElement) (err error) {
	return instance.SetProperty("Spare", value)
}

// GetSpare gets the value of Spare for the instance
func (instance *CIM_ActsAsSpare) GetPropertySpare() (value CIM_ManagedSystemElement, err error) {
	retValue, err := instance.GetProperty("Spare")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedSystemElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
