// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_SnapshotCollection struct
type Msvm_SnapshotCollection struct {
	*CIM_Collection

	// The unique identification of the Collection object.
	CollectionID string
}

func NewMsvm_SnapshotCollectionEx1(instance *cim.WmiInstance) (newInstance *Msvm_SnapshotCollection, err error) {
	tmp, err := NewCIM_CollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SnapshotCollection{
		CIM_Collection: tmp,
	}
	return
}

func NewMsvm_SnapshotCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SnapshotCollection, err error) {
	tmp, err := NewCIM_CollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SnapshotCollection{
		CIM_Collection: tmp,
	}
	return
}

// SetCollectionID sets the value of CollectionID for the instance
func (instance *Msvm_SnapshotCollection) SetPropertyCollectionID(value string) (err error) {
	return instance.SetProperty("CollectionID", (value))
}

// GetCollectionID gets the value of CollectionID for the instance
func (instance *Msvm_SnapshotCollection) GetPropertyCollectionID() (value string, err error) {
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
