// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// Msvm_SnapshotCollection struct
type Msvm_SnapshotCollection struct {
	CIM_Collection

	// The unique identification of the Collection object.
	CollectionID string
}

// SetCollectionID sets the value of CollectionID for the instance
func (instance *Msvm_SnapshotCollection) SetPropertyCollectionID(value string) (err error) {
	return instance.SetProperty("CollectionID", value)
}

// GetCollectionID gets the value of CollectionID for the instance
func (instance *Msvm_SnapshotCollection) GetPropertyCollectionID() (value string, err error) {
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
