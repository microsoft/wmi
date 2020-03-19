// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_MemberOfCollection struct
type CIM_MemberOfCollection struct {
	*cim.WmiInstance

	// The Collection that aggregates members.
	Collection CIM_Collection

	// The aggregated member of the Collection.
	Member CIM_ManagedElement
}

func NewCIM_MemberOfCollectionEx1(instance *cim.WmiInstance) (newInstance *CIM_MemberOfCollection, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_MemberOfCollection{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_MemberOfCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_MemberOfCollection, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_MemberOfCollection{
		WmiInstance: tmp,
	}
	return
}

// SetCollection sets the value of Collection for the instance
func (instance *CIM_MemberOfCollection) SetPropertyCollection(value CIM_Collection) (err error) {
	return instance.SetProperty("Collection", value)
}

// GetCollection gets the value of Collection for the instance
func (instance *CIM_MemberOfCollection) GetPropertyCollection() (value CIM_Collection, err error) {
	retValue, err := instance.GetProperty("Collection")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Collection)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMember sets the value of Member for the instance
func (instance *CIM_MemberOfCollection) SetPropertyMember(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("Member", value)
}

// GetMember gets the value of Member for the instance
func (instance *CIM_MemberOfCollection) GetPropertyMember() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("Member")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
