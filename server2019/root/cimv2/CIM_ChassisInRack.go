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

// CIM_ChassisInRack struct
type CIM_ChassisInRack struct {
	*CIM_Container

	//
	BottomU uint16
}

func NewCIM_ChassisInRackEx1(instance *cim.WmiInstance) (newInstance *CIM_ChassisInRack, err error) {
	tmp, err := NewCIM_ContainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ChassisInRack{
		CIM_Container: tmp,
	}
	return
}

func NewCIM_ChassisInRackEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ChassisInRack, err error) {
	tmp, err := NewCIM_ContainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ChassisInRack{
		CIM_Container: tmp,
	}
	return
}

// SetBottomU sets the value of BottomU for the instance
func (instance *CIM_ChassisInRack) SetPropertyBottomU(value uint16) (err error) {
	return instance.SetProperty("BottomU", value)
}

// GetBottomU gets the value of BottomU for the instance
func (instance *CIM_ChassisInRack) GetPropertyBottomU() (value uint16, err error) {
	retValue, err := instance.GetProperty("BottomU")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
