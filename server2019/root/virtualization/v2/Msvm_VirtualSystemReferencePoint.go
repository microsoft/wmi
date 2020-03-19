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

// Msvm_VirtualSystemReferencePoint struct
type Msvm_VirtualSystemReferencePoint struct {
	*CIM_ManagedElement

	//
	ConsistencyLevel uint16

	//
	HasAssociatedData bool

	//
	ReferencePointType uint16

	//
	ResilientChangeTrackingIdentifiers []string

	//
	VirtualDiskIdentifiers []string

	//
	VirtualSystemIdentifier string
}

func NewMsvm_VirtualSystemReferencePointEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemReferencePoint, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemReferencePoint{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMsvm_VirtualSystemReferencePointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemReferencePoint, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemReferencePoint{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetConsistencyLevel sets the value of ConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemReferencePoint) SetPropertyConsistencyLevel(value uint16) (err error) {
	return instance.SetProperty("ConsistencyLevel", value)
}

// GetConsistencyLevel gets the value of ConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemReferencePoint) GetPropertyConsistencyLevel() (value uint16, err error) {
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

// SetHasAssociatedData sets the value of HasAssociatedData for the instance
func (instance *Msvm_VirtualSystemReferencePoint) SetPropertyHasAssociatedData(value bool) (err error) {
	return instance.SetProperty("HasAssociatedData", value)
}

// GetHasAssociatedData gets the value of HasAssociatedData for the instance
func (instance *Msvm_VirtualSystemReferencePoint) GetPropertyHasAssociatedData() (value bool, err error) {
	retValue, err := instance.GetProperty("HasAssociatedData")
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
func (instance *Msvm_VirtualSystemReferencePoint) SetPropertyReferencePointType(value uint16) (err error) {
	return instance.SetProperty("ReferencePointType", value)
}

// GetReferencePointType gets the value of ReferencePointType for the instance
func (instance *Msvm_VirtualSystemReferencePoint) GetPropertyReferencePointType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReferencePointType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResilientChangeTrackingIdentifiers sets the value of ResilientChangeTrackingIdentifiers for the instance
func (instance *Msvm_VirtualSystemReferencePoint) SetPropertyResilientChangeTrackingIdentifiers(value []string) (err error) {
	return instance.SetProperty("ResilientChangeTrackingIdentifiers", value)
}

// GetResilientChangeTrackingIdentifiers gets the value of ResilientChangeTrackingIdentifiers for the instance
func (instance *Msvm_VirtualSystemReferencePoint) GetPropertyResilientChangeTrackingIdentifiers() (value []string, err error) {
	retValue, err := instance.GetProperty("ResilientChangeTrackingIdentifiers")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDiskIdentifiers sets the value of VirtualDiskIdentifiers for the instance
func (instance *Msvm_VirtualSystemReferencePoint) SetPropertyVirtualDiskIdentifiers(value []string) (err error) {
	return instance.SetProperty("VirtualDiskIdentifiers", value)
}

// GetVirtualDiskIdentifiers gets the value of VirtualDiskIdentifiers for the instance
func (instance *Msvm_VirtualSystemReferencePoint) GetPropertyVirtualDiskIdentifiers() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualDiskIdentifiers")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemIdentifier sets the value of VirtualSystemIdentifier for the instance
func (instance *Msvm_VirtualSystemReferencePoint) SetPropertyVirtualSystemIdentifier(value string) (err error) {
	return instance.SetProperty("VirtualSystemIdentifier", value)
}

// GetVirtualSystemIdentifier gets the value of VirtualSystemIdentifier for the instance
func (instance *Msvm_VirtualSystemReferencePoint) GetPropertyVirtualSystemIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemIdentifier")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
