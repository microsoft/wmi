// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.MSCluster
//
// ////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ParticipatingCS struct
type CIM_ParticipatingCS struct {
	*CIM_Dependency

	//
	RoleOfNode uint16

	//
	StateOfNode uint16
}

func NewCIM_ParticipatingCSEx1(instance *cim.WmiInstance) (newInstance *CIM_ParticipatingCS, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ParticipatingCS{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_ParticipatingCSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ParticipatingCS, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ParticipatingCS{
		CIM_Dependency: tmp,
	}
	return
}

// SetRoleOfNode sets the value of RoleOfNode for the instance
func (instance *CIM_ParticipatingCS) SetPropertyRoleOfNode(value uint16) (err error) {
	return instance.SetProperty("RoleOfNode", (value))
}

// GetRoleOfNode gets the value of RoleOfNode for the instance
func (instance *CIM_ParticipatingCS) GetPropertyRoleOfNode() (value uint16, err error) {
	retValue, err := instance.GetProperty("RoleOfNode")
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

// SetStateOfNode sets the value of StateOfNode for the instance
func (instance *CIM_ParticipatingCS) SetPropertyStateOfNode(value uint16) (err error) {
	return instance.SetProperty("StateOfNode", (value))
}

// GetStateOfNode gets the value of StateOfNode for the instance
func (instance *CIM_ParticipatingCS) GetPropertyStateOfNode() (value uint16, err error) {
	retValue, err := instance.GetProperty("StateOfNode")
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
