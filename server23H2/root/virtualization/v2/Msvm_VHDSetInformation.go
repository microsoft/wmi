// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VHDSetInformation struct
type Msvm_VHDSetInformation struct {
	*cim.WmiInstance

	//
	AllPaths []string

	//
	Path string

	//
	SnapshotIdList []string
}

func NewMsvm_VHDSetInformationEx1(instance *cim.WmiInstance) (newInstance *Msvm_VHDSetInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msvm_VHDSetInformation{
		WmiInstance: tmp,
	}
	return
}

func NewMsvm_VHDSetInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VHDSetInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VHDSetInformation{
		WmiInstance: tmp,
	}
	return
}

// SetAllPaths sets the value of AllPaths for the instance
func (instance *Msvm_VHDSetInformation) SetPropertyAllPaths(value []string) (err error) {
	return instance.SetProperty("AllPaths", (value))
}

// GetAllPaths gets the value of AllPaths for the instance
func (instance *Msvm_VHDSetInformation) GetPropertyAllPaths() (value []string, err error) {
	retValue, err := instance.GetProperty("AllPaths")
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

// SetPath sets the value of Path for the instance
func (instance *Msvm_VHDSetInformation) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *Msvm_VHDSetInformation) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetSnapshotIdList sets the value of SnapshotIdList for the instance
func (instance *Msvm_VHDSetInformation) SetPropertySnapshotIdList(value []string) (err error) {
	return instance.SetProperty("SnapshotIdList", (value))
}

// GetSnapshotIdList gets the value of SnapshotIdList for the instance
func (instance *Msvm_VHDSetInformation) GetPropertySnapshotIdList() (value []string, err error) {
	retValue, err := instance.GetProperty("SnapshotIdList")
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
