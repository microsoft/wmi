// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_PhysicalComponent struct
type CIM_PhysicalComponent struct {
	*CIM_PhysicalElement

	//
	HotSwappable bool

	//
	Removable bool

	//
	Replaceable bool
}

func NewCIM_PhysicalComponentEx1(instance *cim.WmiInstance) (newInstance *CIM_PhysicalComponent, err error) {
	tmp, err := NewCIM_PhysicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PhysicalComponent{
		CIM_PhysicalElement: tmp,
	}
	return
}

func NewCIM_PhysicalComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PhysicalComponent, err error) {
	tmp, err := NewCIM_PhysicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PhysicalComponent{
		CIM_PhysicalElement: tmp,
	}
	return
}

// SetHotSwappable sets the value of HotSwappable for the instance
func (instance *CIM_PhysicalComponent) SetPropertyHotSwappable(value bool) (err error) {
	return instance.SetProperty("HotSwappable", value)
}

// GetHotSwappable gets the value of HotSwappable for the instance
func (instance *CIM_PhysicalComponent) GetPropertyHotSwappable() (value bool, err error) {
	retValue, err := instance.GetProperty("HotSwappable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemovable sets the value of Removable for the instance
func (instance *CIM_PhysicalComponent) SetPropertyRemovable(value bool) (err error) {
	return instance.SetProperty("Removable", value)
}

// GetRemovable gets the value of Removable for the instance
func (instance *CIM_PhysicalComponent) GetPropertyRemovable() (value bool, err error) {
	retValue, err := instance.GetProperty("Removable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplaceable sets the value of Replaceable for the instance
func (instance *CIM_PhysicalComponent) SetPropertyReplaceable(value bool) (err error) {
	return instance.SetProperty("Replaceable", value)
}

// GetReplaceable gets the value of Replaceable for the instance
func (instance *CIM_PhysicalComponent) GetPropertyReplaceable() (value bool, err error) {
	retValue, err := instance.GetProperty("Replaceable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
