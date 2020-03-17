// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VHDSnapshotInformation struct
type Msvm_VHDSnapshotInformation struct {
	cim.WmiInstance

	//
	CreationTime string

	//
	FilePath string

	//
	ParentPathsList []string

	//
	ResilientChangeTrackingId string

	//
	SnapshotId string

	//
	SnapshotPath string
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", value)
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFilePath sets the value of FilePath for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertyFilePath(value string) (err error) {
	return instance.SetProperty("FilePath", value)
}

// GetFilePath gets the value of FilePath for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertyFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("FilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentPathsList sets the value of ParentPathsList for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertyParentPathsList(value []string) (err error) {
	return instance.SetProperty("ParentPathsList", value)
}

// GetParentPathsList gets the value of ParentPathsList for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertyParentPathsList() (value []string, err error) {
	retValue, err := instance.GetProperty("ParentPathsList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResilientChangeTrackingId sets the value of ResilientChangeTrackingId for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertyResilientChangeTrackingId(value string) (err error) {
	return instance.SetProperty("ResilientChangeTrackingId", value)
}

// GetResilientChangeTrackingId gets the value of ResilientChangeTrackingId for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertyResilientChangeTrackingId() (value string, err error) {
	retValue, err := instance.GetProperty("ResilientChangeTrackingId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSnapshotId sets the value of SnapshotId for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertySnapshotId(value string) (err error) {
	return instance.SetProperty("SnapshotId", value)
}

// GetSnapshotId gets the value of SnapshotId for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertySnapshotId() (value string, err error) {
	retValue, err := instance.GetProperty("SnapshotId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSnapshotPath sets the value of SnapshotPath for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertySnapshotPath(value string) (err error) {
	return instance.SetProperty("SnapshotPath", value)
}

// GetSnapshotPath gets the value of SnapshotPath for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertySnapshotPath() (value string, err error) {
	retValue, err := instance.GetProperty("SnapshotPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
