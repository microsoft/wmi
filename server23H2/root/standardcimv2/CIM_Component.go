// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Component struct
type CIM_Component struct {
	*cim.WmiInstance

	//
	GroupComponent CIM_ManagedElement

	//
	PartComponent CIM_ManagedElement
}

func NewCIM_ComponentEx1(instance *cim.WmiInstance) (newInstance *CIM_Component, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_Component{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Component, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Component{
		WmiInstance: tmp,
	}
	return
}

// SetGroupComponent sets the value of GroupComponent for the instance
func (instance *CIM_Component) SetPropertyGroupComponent(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("GroupComponent", (value))
}

// GetGroupComponent gets the value of GroupComponent for the instance
func (instance *CIM_Component) GetPropertyGroupComponent() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("GroupComponent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_ManagedElement)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_ManagedElement is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_ManagedElement(valuetmp)

	return
}

// SetPartComponent sets the value of PartComponent for the instance
func (instance *CIM_Component) SetPropertyPartComponent(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("PartComponent", (value))
}

// GetPartComponent gets the value of PartComponent for the instance
func (instance *CIM_Component) GetPropertyPartComponent() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("PartComponent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_ManagedElement)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_ManagedElement is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_ManagedElement(valuetmp)

	return
}
