// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_MoveUnmanagedVhd struct
type Msvm_MoveUnmanagedVhd struct {
	*CIM_ManagedElement

	//
	VhdDestinationPath string

	//
	VhdSourcePath string
}

func NewMsvm_MoveUnmanagedVhdEx1(instance *cim.WmiInstance) (newInstance *Msvm_MoveUnmanagedVhd, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MoveUnmanagedVhd{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMsvm_MoveUnmanagedVhdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MoveUnmanagedVhd, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MoveUnmanagedVhd{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetVhdDestinationPath sets the value of VhdDestinationPath for the instance
func (instance *Msvm_MoveUnmanagedVhd) SetPropertyVhdDestinationPath(value string) (err error) {
	return instance.SetProperty("VhdDestinationPath", (value))
}

// GetVhdDestinationPath gets the value of VhdDestinationPath for the instance
func (instance *Msvm_MoveUnmanagedVhd) GetPropertyVhdDestinationPath() (value string, err error) {
	retValue, err := instance.GetProperty("VhdDestinationPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetVhdSourcePath sets the value of VhdSourcePath for the instance
func (instance *Msvm_MoveUnmanagedVhd) SetPropertyVhdSourcePath(value string) (err error) {
	return instance.SetProperty("VhdSourcePath", (value))
}

// GetVhdSourcePath gets the value of VhdSourcePath for the instance
func (instance *Msvm_MoveUnmanagedVhd) GetPropertyVhdSourcePath() (value string, err error) {
	retValue, err := instance.GetProperty("VhdSourcePath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
