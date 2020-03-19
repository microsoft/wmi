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

// CIM_MemoryCapacity struct
type CIM_MemoryCapacity struct {
	*CIM_PhysicalCapacity

	//
	MaximumMemoryCapacity uint64

	//
	MemoryType uint16

	//
	MinimumMemoryCapacity uint64
}

func NewCIM_MemoryCapacityEx1(instance *cim.WmiInstance) (newInstance *CIM_MemoryCapacity, err error) {
	tmp, err := NewCIM_PhysicalCapacityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_MemoryCapacity{
		CIM_PhysicalCapacity: tmp,
	}
	return
}

func NewCIM_MemoryCapacityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_MemoryCapacity, err error) {
	tmp, err := NewCIM_PhysicalCapacityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_MemoryCapacity{
		CIM_PhysicalCapacity: tmp,
	}
	return
}

// SetMaximumMemoryCapacity sets the value of MaximumMemoryCapacity for the instance
func (instance *CIM_MemoryCapacity) SetPropertyMaximumMemoryCapacity(value uint64) (err error) {
	return instance.SetProperty("MaximumMemoryCapacity", value)
}

// GetMaximumMemoryCapacity gets the value of MaximumMemoryCapacity for the instance
func (instance *CIM_MemoryCapacity) GetPropertyMaximumMemoryCapacity() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaximumMemoryCapacity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemoryType sets the value of MemoryType for the instance
func (instance *CIM_MemoryCapacity) SetPropertyMemoryType(value uint16) (err error) {
	return instance.SetProperty("MemoryType", value)
}

// GetMemoryType gets the value of MemoryType for the instance
func (instance *CIM_MemoryCapacity) GetPropertyMemoryType() (value uint16, err error) {
	retValue, err := instance.GetProperty("MemoryType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinimumMemoryCapacity sets the value of MinimumMemoryCapacity for the instance
func (instance *CIM_MemoryCapacity) SetPropertyMinimumMemoryCapacity(value uint64) (err error) {
	return instance.SetProperty("MinimumMemoryCapacity", value)
}

// GetMinimumMemoryCapacity gets the value of MinimumMemoryCapacity for the instance
func (instance *CIM_MemoryCapacity) GetPropertyMinimumMemoryCapacity() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinimumMemoryCapacity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
