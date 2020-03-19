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

// CIM_PhysicalExtent struct
type CIM_PhysicalExtent struct {
	*CIM_StorageExtent

	//
	UnitsBeforeCheckDataInterleave uint64

	//
	UnitsOfCheckData uint64

	//
	UnitsOfUserData uint64
}

func NewCIM_PhysicalExtentEx1(instance *cim.WmiInstance) (newInstance *CIM_PhysicalExtent, err error) {
	tmp, err := NewCIM_StorageExtentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PhysicalExtent{
		CIM_StorageExtent: tmp,
	}
	return
}

func NewCIM_PhysicalExtentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PhysicalExtent, err error) {
	tmp, err := NewCIM_StorageExtentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PhysicalExtent{
		CIM_StorageExtent: tmp,
	}
	return
}

// SetUnitsBeforeCheckDataInterleave sets the value of UnitsBeforeCheckDataInterleave for the instance
func (instance *CIM_PhysicalExtent) SetPropertyUnitsBeforeCheckDataInterleave(value uint64) (err error) {
	return instance.SetProperty("UnitsBeforeCheckDataInterleave", value)
}

// GetUnitsBeforeCheckDataInterleave gets the value of UnitsBeforeCheckDataInterleave for the instance
func (instance *CIM_PhysicalExtent) GetPropertyUnitsBeforeCheckDataInterleave() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnitsBeforeCheckDataInterleave")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnitsOfCheckData sets the value of UnitsOfCheckData for the instance
func (instance *CIM_PhysicalExtent) SetPropertyUnitsOfCheckData(value uint64) (err error) {
	return instance.SetProperty("UnitsOfCheckData", value)
}

// GetUnitsOfCheckData gets the value of UnitsOfCheckData for the instance
func (instance *CIM_PhysicalExtent) GetPropertyUnitsOfCheckData() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnitsOfCheckData")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnitsOfUserData sets the value of UnitsOfUserData for the instance
func (instance *CIM_PhysicalExtent) SetPropertyUnitsOfUserData(value uint64) (err error) {
	return instance.SetProperty("UnitsOfUserData", value)
}

// GetUnitsOfUserData gets the value of UnitsOfUserData for the instance
func (instance *CIM_PhysicalExtent) GetPropertyUnitsOfUserData() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnitsOfUserData")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
