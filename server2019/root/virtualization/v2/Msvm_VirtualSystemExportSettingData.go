// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_VirtualSystemExportSettingData struct
type Msvm_VirtualSystemExportSettingData struct {
	CIM_SettingData

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

// SetBackupIntent sets the value of BackupIntent for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyBackupIntent(value uint8) (err error) {
	return instance.SetProperty("BackupIntent", value)
}

// GetBackupIntent gets the value of BackupIntent for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyBackupIntent() (value uint8, err error) {
	retValue, err := instance.GetProperty("BackupIntent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCaptureLiveState sets the value of CaptureLiveState for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyCaptureLiveState(value uint8) (err error) {
	return instance.SetProperty("CaptureLiveState", value)
}

// GetCaptureLiveState gets the value of CaptureLiveState for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyCaptureLiveState() (value uint8, err error) {
	retValue, err := instance.GetProperty("CaptureLiveState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCopySnapshotConfiguration sets the value of CopySnapshotConfiguration for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyCopySnapshotConfiguration(value uint8) (err error) {
	return instance.SetProperty("CopySnapshotConfiguration", value)
}

// GetCopySnapshotConfiguration gets the value of CopySnapshotConfiguration for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyCopySnapshotConfiguration() (value uint8, err error) {
	retValue, err := instance.GetProperty("CopySnapshotConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCopyVmRuntimeInformation sets the value of CopyVmRuntimeInformation for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyCopyVmRuntimeInformation(value bool) (err error) {
	return instance.SetProperty("CopyVmRuntimeInformation", value)
}

// GetCopyVmRuntimeInformation gets the value of CopyVmRuntimeInformation for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyCopyVmRuntimeInformation() (value bool, err error) {
	retValue, err := instance.GetProperty("CopyVmRuntimeInformation")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCopyVmStorage sets the value of CopyVmStorage for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyCopyVmStorage(value bool) (err error) {
	return instance.SetProperty("CopyVmStorage", value)
}

// GetCopyVmStorage gets the value of CopyVmStorage for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyCopyVmStorage() (value bool, err error) {
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

// SetCreateVmExportSubdirectory sets the value of CreateVmExportSubdirectory for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyCreateVmExportSubdirectory(value bool) (err error) {
	return instance.SetProperty("CreateVmExportSubdirectory", value)
}

// GetCreateVmExportSubdirectory gets the value of CreateVmExportSubdirectory for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyCreateVmExportSubdirectory() (value bool, err error) {
	retValue, err := instance.GetProperty("CreateVmExportSubdirectory")
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
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyDifferentialBackupBase(value string) (err error) {
	return instance.SetProperty("DifferentialBackupBase", value)
}

// GetDifferentialBackupBase gets the value of DifferentialBackupBase for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyDifferentialBackupBase() (value string, err error) {
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

// SetDisableDifferentialOfIgnoredStorage sets the value of DisableDifferentialOfIgnoredStorage for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyDisableDifferentialOfIgnoredStorage(value bool) (err error) {
	return instance.SetProperty("DisableDifferentialOfIgnoredStorage", value)
}

// GetDisableDifferentialOfIgnoredStorage gets the value of DisableDifferentialOfIgnoredStorage for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyDisableDifferentialOfIgnoredStorage() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableDifferentialOfIgnoredStorage")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExcludedVirtualHardDisks sets the value of ExcludedVirtualHardDisks for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyExcludedVirtualHardDisks(value []string) (err error) {
	return instance.SetProperty("ExcludedVirtualHardDisks", value)
}

// GetExcludedVirtualHardDisks gets the value of ExcludedVirtualHardDisks for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyExcludedVirtualHardDisks() (value []string, err error) {
	retValue, err := instance.GetProperty("ExcludedVirtualHardDisks")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExportForLiveMigration sets the value of ExportForLiveMigration for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyExportForLiveMigration(value bool) (err error) {
	return instance.SetProperty("ExportForLiveMigration", value)
}

// GetExportForLiveMigration gets the value of ExportForLiveMigration for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyExportForLiveMigration() (value bool, err error) {
	retValue, err := instance.GetProperty("ExportForLiveMigration")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuestDebugStateFilePath sets the value of GuestDebugStateFilePath for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertyGuestDebugStateFilePath(value string) (err error) {
	return instance.SetProperty("GuestDebugStateFilePath", value)
}

// GetGuestDebugStateFilePath gets the value of GuestDebugStateFilePath for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertyGuestDebugStateFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("GuestDebugStateFilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSnapshotVirtualSystem sets the value of SnapshotVirtualSystem for the instance
func (instance *Msvm_VirtualSystemExportSettingData) SetPropertySnapshotVirtualSystem(value string) (err error) {
	return instance.SetProperty("SnapshotVirtualSystem", value)
}

// GetSnapshotVirtualSystem gets the value of SnapshotVirtualSystem for the instance
func (instance *Msvm_VirtualSystemExportSettingData) GetPropertySnapshotVirtualSystem() (value string, err error) {
	retValue, err := instance.GetProperty("SnapshotVirtualSystem")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
