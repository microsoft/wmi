// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_DiskPartition struct
type CIM_DiskPartition struct {
	CIM_StorageExtent

	//
	Bootable bool

	//
	PrimaryPartition bool
}

// SetBootable sets the value of Bootable for the instance
func (instance *CIM_DiskPartition) SetPropertyBootable(value bool) (err error) {
	return instance.SetProperty("Bootable", value)
}

// GetBootable gets the value of Bootable for the instance
func (instance *CIM_DiskPartition) GetPropertyBootable() (value bool, err error) {
	retValue, err := instance.GetProperty("Bootable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimaryPartition sets the value of PrimaryPartition for the instance
func (instance *CIM_DiskPartition) SetPropertyPrimaryPartition(value bool) (err error) {
	return instance.SetProperty("PrimaryPartition", value)
}

// GetPrimaryPartition gets the value of PrimaryPartition for the instance
func (instance *CIM_DiskPartition) GetPropertyPrimaryPartition() (value bool, err error) {
	retValue, err := instance.GetProperty("PrimaryPartition")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
