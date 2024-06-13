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

// Msvm_CopyFileToGuestSettingData struct
type Msvm_CopyFileToGuestSettingData struct {
	*CIM_SettingData

	//
	CreateFullPath bool

	//
	DestinationPath string

	//
	OverwriteExisting bool

	//
	SourcePath string
}

func NewMsvm_CopyFileToGuestSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_CopyFileToGuestSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CopyFileToGuestSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_CopyFileToGuestSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CopyFileToGuestSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CopyFileToGuestSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetCreateFullPath sets the value of CreateFullPath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) SetPropertyCreateFullPath(value bool) (err error) {
	return instance.SetProperty("CreateFullPath", (value))
}

// GetCreateFullPath gets the value of CreateFullPath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) GetPropertyCreateFullPath() (value bool, err error) {
	retValue, err := instance.GetProperty("CreateFullPath")
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

// SetDestinationPath sets the value of DestinationPath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) SetPropertyDestinationPath(value string) (err error) {
	return instance.SetProperty("DestinationPath", (value))
}

// GetDestinationPath gets the value of DestinationPath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) GetPropertyDestinationPath() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationPath")
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

// SetOverwriteExisting sets the value of OverwriteExisting for the instance
func (instance *Msvm_CopyFileToGuestSettingData) SetPropertyOverwriteExisting(value bool) (err error) {
	return instance.SetProperty("OverwriteExisting", (value))
}

// GetOverwriteExisting gets the value of OverwriteExisting for the instance
func (instance *Msvm_CopyFileToGuestSettingData) GetPropertyOverwriteExisting() (value bool, err error) {
	retValue, err := instance.GetProperty("OverwriteExisting")
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

// SetSourcePath sets the value of SourcePath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) SetPropertySourcePath(value string) (err error) {
	return instance.SetProperty("SourcePath", (value))
}

// GetSourcePath gets the value of SourcePath for the instance
func (instance *Msvm_CopyFileToGuestSettingData) GetPropertySourcePath() (value string, err error) {
	retValue, err := instance.GetProperty("SourcePath")
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
