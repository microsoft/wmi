// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_CollectionOfMSEs struct
type CIM_CollectionOfMSEs struct {
	CIM_Collection

	// The identification of the Collection object. When subclassed, the CollectionID property can be overridden to be a Key property.
	CollectionID string
}

// SetCollectionID sets the value of CollectionID for the instance
func (instance *CIM_CollectionOfMSEs) SetPropertyCollectionID(value string) (err error) {
	return instance.SetProperty("CollectionID", value)
}

// GetCollectionID gets the value of CollectionID for the instance
func (instance *CIM_CollectionOfMSEs) GetPropertyCollectionID() (value string, err error) {
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
