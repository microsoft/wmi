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

// CIM_CollectedCollections struct
type CIM_CollectedCollections struct {
	cim.WmiInstance

	//
	Collection CIM_CollectionOfMSEs

	//
	CollectionInCollection CIM_CollectionOfMSEs
}

// SetCollection sets the value of Collection for the instance
func (instance *CIM_CollectedCollections) SetPropertyCollection(value CIM_CollectionOfMSEs) (err error) {
	return instance.SetProperty("Collection", value)
}

// GetCollection gets the value of Collection for the instance
func (instance *CIM_CollectedCollections) GetPropertyCollection() (value CIM_CollectionOfMSEs, err error) {
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

// SetCollectionInCollection sets the value of CollectionInCollection for the instance
func (instance *CIM_CollectedCollections) SetPropertyCollectionInCollection(value CIM_CollectionOfMSEs) (err error) {
	return instance.SetProperty("CollectionInCollection", value)
}

// GetCollectionInCollection gets the value of CollectionInCollection for the instance
func (instance *CIM_CollectedCollections) GetPropertyCollectionInCollection() (value CIM_CollectionOfMSEs, err error) {
	retValue, err := instance.GetProperty("CollectionInCollection")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_CollectionOfMSEs)
	if !ok {
		// TODO: Set an error
	}
	return
}
