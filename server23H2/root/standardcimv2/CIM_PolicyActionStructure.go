// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_PolicyActionStructure struct
type CIM_PolicyActionStructure struct {
	*CIM_PolicyComponent

	//
	ActionOrder uint16
}

func NewCIM_PolicyActionStructureEx1(instance *cim.WmiInstance) (newInstance *CIM_PolicyActionStructure, err error) {
	tmp, err := NewCIM_PolicyComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PolicyActionStructure{
		CIM_PolicyComponent: tmp,
	}
	return
}

func NewCIM_PolicyActionStructureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PolicyActionStructure, err error) {
	tmp, err := NewCIM_PolicyComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PolicyActionStructure{
		CIM_PolicyComponent: tmp,
	}
	return
}

// SetActionOrder sets the value of ActionOrder for the instance
func (instance *CIM_PolicyActionStructure) SetPropertyActionOrder(value uint16) (err error) {
	return instance.SetProperty("ActionOrder", (value))
}

// GetActionOrder gets the value of ActionOrder for the instance
func (instance *CIM_PolicyActionStructure) GetPropertyActionOrder() (value uint16, err error) {
	retValue, err := instance.GetProperty("ActionOrder")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
