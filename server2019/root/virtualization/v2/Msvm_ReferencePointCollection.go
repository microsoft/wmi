// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ReferencePointCollection struct
type Msvm_ReferencePointCollection struct {
	*CIM_Collection

	//
	CollectionID string

	//
	ConsistencyLevel uint16

	//
	HasAssociatedLog bool

	//
	ReferencePointType ReferencePointCollection_ReferencePointType

	//
	VirtualSystemCollectionId string
}

func NewMsvm_ReferencePointCollectionEx1(instance *cim.WmiInstance) (newInstance *Msvm_ReferencePointCollection, err error) {
	tmp, err := NewCIM_CollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReferencePointCollection{
		CIM_Collection: tmp,
	}
	return
}

func NewMsvm_ReferencePointCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ReferencePointCollection, err error) {
	tmp, err := NewCIM_CollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReferencePointCollection{
		CIM_Collection: tmp,
	}
	return
}

// SetCollectionID sets the value of CollectionID for the instance
func (instance *Msvm_ReferencePointCollection) SetPropertyCollectionID(value string) (err error) {
	return instance.SetProperty("CollectionID", value)
}

// GetCollectionID gets the value of CollectionID for the instance
func (instance *Msvm_ReferencePointCollection) GetPropertyCollectionID() (value string, err error) {
	retValue, err := instance.GetProperty("CollectionID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConsistencyLevel sets the value of ConsistencyLevel for the instance
func (instance *Msvm_ReferencePointCollection) SetPropertyConsistencyLevel(value uint16) (err error) {
	return instance.SetProperty("ConsistencyLevel", value)
}

// GetConsistencyLevel gets the value of ConsistencyLevel for the instance
func (instance *Msvm_ReferencePointCollection) GetPropertyConsistencyLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("ConsistencyLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHasAssociatedLog sets the value of HasAssociatedLog for the instance
func (instance *Msvm_ReferencePointCollection) SetPropertyHasAssociatedLog(value bool) (err error) {
	return instance.SetProperty("HasAssociatedLog", value)
}

// GetHasAssociatedLog gets the value of HasAssociatedLog for the instance
func (instance *Msvm_ReferencePointCollection) GetPropertyHasAssociatedLog() (value bool, err error) {
	retValue, err := instance.GetProperty("HasAssociatedLog")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReferencePointType sets the value of ReferencePointType for the instance
func (instance *Msvm_ReferencePointCollection) SetPropertyReferencePointType(value ReferencePointCollection_ReferencePointType) (err error) {
	return instance.SetProperty("ReferencePointType", value)
}

// GetReferencePointType gets the value of ReferencePointType for the instance
func (instance *Msvm_ReferencePointCollection) GetPropertyReferencePointType() (value ReferencePointCollection_ReferencePointType, err error) {
	retValue, err := instance.GetProperty("ReferencePointType")
	if err != nil {
		return
	}
	value, ok := retValue.(ReferencePointCollection_ReferencePointType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemCollectionId sets the value of VirtualSystemCollectionId for the instance
func (instance *Msvm_ReferencePointCollection) SetPropertyVirtualSystemCollectionId(value string) (err error) {
	return instance.SetProperty("VirtualSystemCollectionId", value)
}

// GetVirtualSystemCollectionId gets the value of VirtualSystemCollectionId for the instance
func (instance *Msvm_ReferencePointCollection) GetPropertyVirtualSystemCollectionId() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemCollectionId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
