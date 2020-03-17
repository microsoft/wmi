// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_CollectedMSEs struct
type CIM_CollectedMSEs struct {
	cim.WmiInstance

	//
	Collection CIM_CollectionOfMSEs

	//
	Member CIM_ManagedSystemElement
}

// SetCollection sets the value of Collection for the instance
func (instance *CIM_CollectedMSEs) SetPropertyCollection(value CIM_CollectionOfMSEs) (err error) {
	return instance.SetProperty("Collection", value)
}

// GetCollection gets the value of Collection for the instance
func (instance *CIM_CollectedMSEs) GetPropertyCollection() (value CIM_CollectionOfMSEs, err error) {
	retValue, err := instance.GetProperty("Collection")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_CollectionOfMSEs)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMember sets the value of Member for the instance
func (instance *CIM_CollectedMSEs) SetPropertyMember(value CIM_ManagedSystemElement) (err error) {
	return instance.SetProperty("Member", value)
}

// GetMember gets the value of Member for the instance
func (instance *CIM_CollectedMSEs) GetPropertyMember() (value CIM_ManagedSystemElement, err error) {
	retValue, err := instance.GetProperty("Member")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedSystemElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
