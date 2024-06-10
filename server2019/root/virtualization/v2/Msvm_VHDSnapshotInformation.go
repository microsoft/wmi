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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VHDSnapshotInformation struct
type Msvm_VHDSnapshotInformation struct {
	*cim.WmiInstance

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

func NewMsvm_VHDSnapshotInformationEx1(instance *cim.WmiInstance) (newInstance *Msvm_VHDSnapshotInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msvm_VHDSnapshotInformation{
		WmiInstance: tmp,
	}
	return
}

func NewMsvm_VHDSnapshotInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VHDSnapshotInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VHDSnapshotInformation{
		WmiInstance: tmp,
	}
	return
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", (value))
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
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

// SetFilePath sets the value of FilePath for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertyFilePath(value string) (err error) {
	return instance.SetProperty("FilePath", (value))
}

// GetFilePath gets the value of FilePath for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertyFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("FilePath")
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

// SetParentPathsList sets the value of ParentPathsList for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertyParentPathsList(value []string) (err error) {
	return instance.SetProperty("ParentPathsList", (value))
}

// GetParentPathsList gets the value of ParentPathsList for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertyParentPathsList() (value []string, err error) {
	retValue, err := instance.GetProperty("ParentPathsList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetResilientChangeTrackingId sets the value of ResilientChangeTrackingId for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertyResilientChangeTrackingId(value string) (err error) {
	return instance.SetProperty("ResilientChangeTrackingId", (value))
}

// GetResilientChangeTrackingId gets the value of ResilientChangeTrackingId for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertyResilientChangeTrackingId() (value string, err error) {
	retValue, err := instance.GetProperty("ResilientChangeTrackingId")
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

// SetSnapshotId sets the value of SnapshotId for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertySnapshotId(value string) (err error) {
	return instance.SetProperty("SnapshotId", (value))
}

// GetSnapshotId gets the value of SnapshotId for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertySnapshotId() (value string, err error) {
	retValue, err := instance.GetProperty("SnapshotId")
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

// SetSnapshotPath sets the value of SnapshotPath for the instance
func (instance *Msvm_VHDSnapshotInformation) SetPropertySnapshotPath(value string) (err error) {
	return instance.SetProperty("SnapshotPath", (value))
}

// GetSnapshotPath gets the value of SnapshotPath for the instance
func (instance *Msvm_VHDSnapshotInformation) GetPropertySnapshotPath() (value string, err error) {
	retValue, err := instance.GetProperty("SnapshotPath")
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
