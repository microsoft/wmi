// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_CollectionSnapshotExportSettingData struct
type Msvm_CollectionSnapshotExportSettingData struct {
	*CIM_SettingData

	//
	BackupIntent uint16

	//
	CopyVmStorage bool

	//
	DifferentialBackupBase string
}

func NewMsvm_CollectionSnapshotExportSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectionSnapshotExportSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionSnapshotExportSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_CollectionSnapshotExportSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CollectionSnapshotExportSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionSnapshotExportSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetBackupIntent sets the value of BackupIntent for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) SetPropertyBackupIntent(value uint16) (err error) {
	return instance.SetProperty("BackupIntent", (value))
}

// GetBackupIntent gets the value of BackupIntent for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) GetPropertyBackupIntent() (value uint16, err error) {
	retValue, err := instance.GetProperty("BackupIntent")
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

// SetCopyVmStorage sets the value of CopyVmStorage for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) SetPropertyCopyVmStorage(value bool) (err error) {
	return instance.SetProperty("CopyVmStorage", (value))
}

// GetCopyVmStorage gets the value of CopyVmStorage for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) GetPropertyCopyVmStorage() (value bool, err error) {
	retValue, err := instance.GetProperty("CopyVmStorage")
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

// SetDifferentialBackupBase sets the value of DifferentialBackupBase for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) SetPropertyDifferentialBackupBase(value string) (err error) {
	return instance.SetProperty("DifferentialBackupBase", (value))
}

// GetDifferentialBackupBase gets the value of DifferentialBackupBase for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) GetPropertyDifferentialBackupBase() (value string, err error) {
	retValue, err := instance.GetProperty("DifferentialBackupBase")
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
