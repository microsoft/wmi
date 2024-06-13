// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
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
	return instance.SetProperty("CollectionID", (value))
}

// GetCollectionID gets the value of CollectionID for the instance
func (instance *Msvm_ReferencePointCollection) GetPropertyCollectionID() (value string, err error) {
	retValue, err := instance.GetProperty("CollectionID")
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

// SetConsistencyLevel sets the value of ConsistencyLevel for the instance
func (instance *Msvm_ReferencePointCollection) SetPropertyConsistencyLevel(value uint16) (err error) {
	return instance.SetProperty("ConsistencyLevel", (value))
}

// GetConsistencyLevel gets the value of ConsistencyLevel for the instance
func (instance *Msvm_ReferencePointCollection) GetPropertyConsistencyLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("ConsistencyLevel")
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

// SetHasAssociatedLog sets the value of HasAssociatedLog for the instance
func (instance *Msvm_ReferencePointCollection) SetPropertyHasAssociatedLog(value bool) (err error) {
	return instance.SetProperty("HasAssociatedLog", (value))
}

// GetHasAssociatedLog gets the value of HasAssociatedLog for the instance
func (instance *Msvm_ReferencePointCollection) GetPropertyHasAssociatedLog() (value bool, err error) {
	retValue, err := instance.GetProperty("HasAssociatedLog")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetReferencePointType sets the value of ReferencePointType for the instance
func (instance *Msvm_ReferencePointCollection) SetPropertyReferencePointType(value ReferencePointCollection_ReferencePointType) (err error) {
	return instance.SetProperty("ReferencePointType", (value))
}

// GetReferencePointType gets the value of ReferencePointType for the instance
func (instance *Msvm_ReferencePointCollection) GetPropertyReferencePointType() (value ReferencePointCollection_ReferencePointType, err error) {
	retValue, err := instance.GetProperty("ReferencePointType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ReferencePointCollection_ReferencePointType(valuetmp)

	return
}

// SetVirtualSystemCollectionId sets the value of VirtualSystemCollectionId for the instance
func (instance *Msvm_ReferencePointCollection) SetPropertyVirtualSystemCollectionId(value string) (err error) {
	return instance.SetProperty("VirtualSystemCollectionId", (value))
}

// GetVirtualSystemCollectionId gets the value of VirtualSystemCollectionId for the instance
func (instance *Msvm_ReferencePointCollection) GetPropertyVirtualSystemCollectionId() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemCollectionId")
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
