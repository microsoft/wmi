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

// CIM_CardOnCard struct
type CIM_CardOnCard struct {
	*CIM_Container

	//
	MountOrSlotDescription string
}

func NewCIM_CardOnCardEx1(instance *cim.WmiInstance) (newInstance *CIM_CardOnCard, err error) {
	tmp, err := NewCIM_ContainerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_CardOnCard{
		CIM_Container: tmp,
	}
	return
}

func NewCIM_CardOnCardEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_CardOnCard, err error) {
	tmp, err := NewCIM_ContainerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_CardOnCard{
		CIM_Container: tmp,
	}
	return
}

// SetMountOrSlotDescription sets the value of MountOrSlotDescription for the instance
func (instance *CIM_CardOnCard) SetPropertyMountOrSlotDescription(value string) (err error) {
	return instance.SetProperty("MountOrSlotDescription", value)
}

// GetMountOrSlotDescription gets the value of MountOrSlotDescription for the instance
func (instance *CIM_CardOnCard) GetPropertyMountOrSlotDescription() (value string, err error) {
	retValue, err := instance.GetProperty("MountOrSlotDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
