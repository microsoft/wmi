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
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VirtualSystemExportSettingData struct
type Msvm_VirtualSystemExportSettingData struct {
	*CIM_SettingData

	//
	BackupIntent uint8

	//
	CaptureLiveState uint8

	//
	CopySnapshotConfiguration uint8

	//
	CopyVmRuntimeInformation bool

	//
	CopyVmStorage bool

	//
	CreateVmExportSubdirectory bool

	//
	DifferentialBackupBase string

	//
	DisableDifferentialOfIgnoredStorage bool

	//
	ExcludedVirtualHardDisks []string

	//
	ExportForLiveMigration bool

	//
	GuestDebugStateFilePath string

	//
	SnapshotVirtualSystem string
}

func NewMsvm_VirtualSystemExportSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemExportSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemExportSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_VirtualSystemExportSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemExportSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemExportSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetBackupIntent sets the value of BackupIntent for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyBackupIntent(value uint8) (err error) {
	return instance.SetProperty("BackupIntent", (value))
}

// GetBackupIntent gets the value of BackupIntent for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyBackupIntent() (value uint8, err error) {
	retValue, err := instance.GetProperty("BackupIntent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetCaptureLiveState sets the value of CaptureLiveState for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyCaptureLiveState(value uint8) (err error) {
	return instance.SetProperty("CaptureLiveState", (value))
}

// GetCaptureLiveState gets the value of CaptureLiveState for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyCaptureLiveState() (value uint8, err error) {
	retValue, err := instance.GetProperty("CaptureLiveState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetCopySnapshotConfiguration sets the value of CopySnapshotConfiguration for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyCopySnapshotConfiguration(value uint8) (err error) {
	return instance.SetProperty("CopySnapshotConfiguration", (value))
}

// GetCopySnapshotConfiguration gets the value of CopySnapshotConfiguration for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyCopySnapshotConfiguration() (value uint8, err error) {
	retValue, err := instance.GetProperty("CopySnapshotConfiguration")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetCopyVmRuntimeInformation sets the value of CopyVmRuntimeInformation for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyCopyVmRuntimeInformation(value bool) (err error) {
	return instance.SetProperty("CopyVmRuntimeInformation", (value))
}

// GetCopyVmRuntimeInformation gets the value of CopyVmRuntimeInformation for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyCopyVmRuntimeInformation() (value bool, err error) {
	retValue, err := instance.GetProperty("CopyVmRuntimeInformation")
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

// SetCopyVmStorage sets the value of CopyVmStorage for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyCopyVmStorage(value bool) (err error) {
	return instance.SetProperty("CopyVmStorage", (value))
}

// GetCopyVmStorage gets the value of CopyVmStorage for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyCopyVmStorage() (value bool, err error) {
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

// SetCreateVmExportSubdirectory sets the value of CreateVmExportSubdirectory for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyCreateVmExportSubdirectory(value bool) (err error) {
	return instance.SetProperty("CreateVmExportSubdirectory", (value))
}

// GetCreateVmExportSubdirectory gets the value of CreateVmExportSubdirectory for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyCreateVmExportSubdirectory() (value bool, err error) {
	retValue, err := instance.GetProperty("CreateVmExportSubdirectory")
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
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyDifferentialBackupBase(value string) (err error) {
	return instance.SetProperty("DifferentialBackupBase", (value))
}

// GetDifferentialBackupBase gets the value of DifferentialBackupBase for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyDifferentialBackupBase() (value string, err error) {
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

// SetDisableDifferentialOfIgnoredStorage sets the value of DisableDifferentialOfIgnoredStorage for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyDisableDifferentialOfIgnoredStorage(value bool) (err error) {
	return instance.SetProperty("DisableDifferentialOfIgnoredStorage", (value))
}

// GetDisableDifferentialOfIgnoredStorage gets the value of DisableDifferentialOfIgnoredStorage for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyDisableDifferentialOfIgnoredStorage() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableDifferentialOfIgnoredStorage")
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

// SetExcludedVirtualHardDisks sets the value of ExcludedVirtualHardDisks for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyExcludedVirtualHardDisks(value []string) (err error) {
	return instance.SetProperty("ExcludedVirtualHardDisks", (value))
}

// GetExcludedVirtualHardDisks gets the value of ExcludedVirtualHardDisks for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyExcludedVirtualHardDisks() (value []string, err error) {
	retValue, err := instance.GetProperty("ExcludedVirtualHardDisks")
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

// SetExportForLiveMigration sets the value of ExportForLiveMigration for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyExportForLiveMigration(value bool) (err error) {
	return instance.SetProperty("ExportForLiveMigration", (value))
}

// GetExportForLiveMigration gets the value of ExportForLiveMigration for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyExportForLiveMigration() (value bool, err error) {
	retValue, err := instance.GetProperty("ExportForLiveMigration")
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

// SetGuestDebugStateFilePath sets the value of GuestDebugStateFilePath for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyGuestDebugStateFilePath(value string) (err error) {
	return instance.SetProperty("GuestDebugStateFilePath", (value))
}

// GetGuestDebugStateFilePath gets the value of GuestDebugStateFilePath for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyGuestDebugStateFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("GuestDebugStateFilePath")
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

// SetSnapshotVirtualSystem sets the value of SnapshotVirtualSystem for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertySnapshotVirtualSystem(value string) (err error) {
	return instance.SetProperty("SnapshotVirtualSystem", (value))
}

// GetSnapshotVirtualSystem gets the value of SnapshotVirtualSystem for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertySnapshotVirtualSystem() (value string, err error) {
	retValue, err := instance.GetProperty("SnapshotVirtualSystem")
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
