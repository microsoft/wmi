// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ElementCapacity struct
type CIM_ElementCapacity struct {
	*cim.WmiInstance

	//
	Capacity CIM_PhysicalCapacity

	//
	Element CIM_PhysicalElement
}

func NewCIM_ElementCapacityEx1(instance *cim.WmiInstance) (newInstance *CIM_ElementCapacity, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_ElementCapacity{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ElementCapacityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ElementCapacity, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ElementCapacity{
		WmiInstance: tmp,
	}
	return
}

// SetCapacity sets the value of Capacity for the instance
func (instance *CIM_ElementCapacity) SetPropertyCapacity(value CIM_PhysicalCapacity) (err error) {
	return instance.SetProperty("Capacity", value)
}

// GetCapacity gets the value of Capacity for the instance
func (instance *CIM_ElementCapacity) GetPropertyCapacity() (value CIM_PhysicalCapacity, err error) {
	retValue, err := instance.GetProperty("Capacity")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_PhysicalCapacity)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElement sets the value of Element for the instance
func (instance *CIM_ElementCapacity) SetPropertyElement(value CIM_PhysicalElement) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *CIM_ElementCapacity) GetPropertyElement() (value CIM_PhysicalElement, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_PhysicalElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
