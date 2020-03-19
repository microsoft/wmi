// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("BackupIntent", value)
}

// GetBackupIntent gets the value of BackupIntent for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) GetPropertyBackupIntent() (value uint16, err error) {
	retValue, err := instance.GetProperty("BackupIntent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCopyVmStorage sets the value of CopyVmStorage for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) SetPropertyCopyVmStorage(value bool) (err error) {
	return instance.SetProperty("CopyVmStorage", value)
}

// GetCopyVmStorage gets the value of CopyVmStorage for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) GetPropertyCopyVmStorage() (value bool, err error) {
	retValue, err := instance.GetProperty("CopyVmStorage")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDifferentialBackupBase sets the value of DifferentialBackupBase for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) SetPropertyDifferentialBackupBase(value string) (err error) {
	return instance.SetProperty("DifferentialBackupBase", value)
}

// GetDifferentialBackupBase gets the value of DifferentialBackupBase for the instance
func (instance *Msvm_CollectionSnapshotExportSettingData) GetPropertyDifferentialBackupBase() (value string, err error) {
	retValue, err := instance.GetProperty("DifferentialBackupBase")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
