// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_LogicalIdentity struct
type CIM_LogicalIdentity struct {
	*cim.WmiInstance

	//
	SameElement CIM_ManagedElement

	//
	SystemElement CIM_ManagedElement
}

func NewCIM_LogicalIdentityEx1(instance *cim.WmiInstance) (newInstance *CIM_LogicalIdentity, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_LogicalIdentity{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_LogicalIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_LogicalIdentity, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_LogicalIdentity{
		WmiInstance: tmp,
	}
	return
}

// SetSameElement sets the value of SameElement for the instance
func (instance *CIM_LogicalIdentity) SetPropertySameElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("SameElement", value)
}

// GetSameElement gets the value of SameElement for the instance
func (instance *CIM_LogicalIdentity) GetPropertySameElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("SameElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemElement sets the value of SystemElement for the instance
func (instance *CIM_LogicalIdentity) SetPropertySystemElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("SystemElement", value)
}

// GetSystemElement gets the value of SystemElement for the instance
func (instance *CIM_LogicalIdentity) GetPropertySystemElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("SystemElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
