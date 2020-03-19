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

// CIM_Container struct
type CIM_Container struct {
	*CIM_Component

	//
	LocationWithinContainer string
}

func NewCIM_ContainerEx1(instance *cim.WmiInstance) (newInstance *CIM_Container, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Container{
		CIM_Component: tmp,
	}
	return
}

func NewCIM_ContainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Container, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Container{
		CIM_Component: tmp,
	}
	return
}

// SetLocationWithinContainer sets the value of LocationWithinContainer for the instance
func (instance *CIM_Container) SetPropertyLocationWithinContainer(value string) (err error) {
	return instance.SetProperty("LocationWithinContainer", value)
}

// GetLocationWithinContainer gets the value of LocationWithinContainer for the instance
func (instance *CIM_Container) GetPropertyLocationWithinContainer() (value string, err error) {
	retValue, err := instance.GetProperty("LocationWithinContainer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
