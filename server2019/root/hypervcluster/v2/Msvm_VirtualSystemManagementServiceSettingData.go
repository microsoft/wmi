// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// Msvm_VirtualSystemManagementServiceSettingData struct
type Msvm_VirtualSystemManagementServiceSettingData struct {
	CIM_SettingData

	//
	BiosLockString string

	//
	CurrentWWNNAddress string

	//
	DefaultExternalDataRoot string

	//
	DefaultVirtualHardDiskCachingMode uint16

	//
	DefaultVirtualHardDiskPath string

	//
	EnhancedSessionModeEnabled bool

	//
	HbaLunTimeout uint32

	//
	HypervisorRootSchedulerEnabled bool

	//
	MaximumMacAddress string

	//
	MaximumWWPNAddress string

	//
	MinimumMacAddress string

	//
	MinimumWWPNAddress string

	//
	NumaSpanningEnabled bool

	//
	PrimaryOwnerContact string

	//
	PrimaryOwnerName string
}

// SetBiosLockString sets the value of BiosLockString for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyBiosLockString(value string) (err error) {
	return instance.SetProperty("BiosLockString", value)
}

// GetBiosLockString gets the value of BiosLockString for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyBiosLockString() (value string, err error) {
	retValue, err := instance.GetProperty("BiosLockString")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentWWNNAddress sets the value of CurrentWWNNAddress for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyCurrentWWNNAddress(value string) (err error) {
	return instance.SetProperty("CurrentWWNNAddress", value)
}

// GetCurrentWWNNAddress gets the value of CurrentWWNNAddress for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyCurrentWWNNAddress() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentWWNNAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultExternalDataRoot sets the value of DefaultExternalDataRoot for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyDefaultExternalDataRoot(value string) (err error) {
	return instance.SetProperty("DefaultExternalDataRoot", value)
}

// GetDefaultExternalDataRoot gets the value of DefaultExternalDataRoot for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyDefaultExternalDataRoot() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultExternalDataRoot")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultVirtualHardDiskCachingMode sets the value of DefaultVirtualHardDiskCachingMode for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyDefaultVirtualHardDiskCachingMode(value uint16) (err error) {
	return instance.SetProperty("DefaultVirtualHardDiskCachingMode", value)
}

// GetDefaultVirtualHardDiskCachingMode gets the value of DefaultVirtualHardDiskCachingMode for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyDefaultVirtualHardDiskCachingMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("DefaultVirtualHardDiskCachingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultVirtualHardDiskPath sets the value of DefaultVirtualHardDiskPath for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyDefaultVirtualHardDiskPath(value string) (err error) {
	return instance.SetProperty("DefaultVirtualHardDiskPath", value)
}

// GetDefaultVirtualHardDiskPath gets the value of DefaultVirtualHardDiskPath for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyDefaultVirtualHardDiskPath() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultVirtualHardDiskPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnhancedSessionModeEnabled sets the value of EnhancedSessionModeEnabled for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyEnhancedSessionModeEnabled(value bool) (err error) {
	return instance.SetProperty("EnhancedSessionModeEnabled", value)
}

// GetEnhancedSessionModeEnabled gets the value of EnhancedSessionModeEnabled for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyEnhancedSessionModeEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("EnhancedSessionModeEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHbaLunTimeout sets the value of HbaLunTimeout for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyHbaLunTimeout(value uint32) (err error) {
	return instance.SetProperty("HbaLunTimeout", value)
}

// GetHbaLunTimeout gets the value of HbaLunTimeout for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyHbaLunTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("HbaLunTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHypervisorRootSchedulerEnabled sets the value of HypervisorRootSchedulerEnabled for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyHypervisorRootSchedulerEnabled(value bool) (err error) {
	return instance.SetProperty("HypervisorRootSchedulerEnabled", value)
}

// GetHypervisorRootSchedulerEnabled gets the value of HypervisorRootSchedulerEnabled for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyHypervisorRootSchedulerEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HypervisorRootSchedulerEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumMacAddress sets the value of MaximumMacAddress for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyMaximumMacAddress(value string) (err error) {
	return instance.SetProperty("MaximumMacAddress", value)
}

// GetMaximumMacAddress gets the value of MaximumMacAddress for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyMaximumMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MaximumMacAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumWWPNAddress sets the value of MaximumWWPNAddress for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyMaximumWWPNAddress(value string) (err error) {
	return instance.SetProperty("MaximumWWPNAddress", value)
}

// GetMaximumWWPNAddress gets the value of MaximumWWPNAddress for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyMaximumWWPNAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MaximumWWPNAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinimumMacAddress sets the value of MinimumMacAddress for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyMinimumMacAddress(value string) (err error) {
	return instance.SetProperty("MinimumMacAddress", value)
}

// GetMinimumMacAddress gets the value of MinimumMacAddress for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyMinimumMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MinimumMacAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinimumWWPNAddress sets the value of MinimumWWPNAddress for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyMinimumWWPNAddress(value string) (err error) {
	return instance.SetProperty("MinimumWWPNAddress", value)
}

// GetMinimumWWPNAddress gets the value of MinimumWWPNAddress for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyMinimumWWPNAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MinimumWWPNAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumaSpanningEnabled sets the value of NumaSpanningEnabled for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyNumaSpanningEnabled(value bool) (err error) {
	return instance.SetProperty("NumaSpanningEnabled", value)
}

// GetNumaSpanningEnabled gets the value of NumaSpanningEnabled for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyNumaSpanningEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("NumaSpanningEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimaryOwnerContact sets the value of PrimaryOwnerContact for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyPrimaryOwnerContact(value string) (err error) {
	return instance.SetProperty("PrimaryOwnerContact", value)
}

// GetPrimaryOwnerContact gets the value of PrimaryOwnerContact for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyPrimaryOwnerContact() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryOwnerContact")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimaryOwnerName sets the value of PrimaryOwnerName for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) SetPropertyPrimaryOwnerName(value string) (err error) {
	return instance.SetProperty("PrimaryOwnerName", value)
}

// GetPrimaryOwnerName gets the value of PrimaryOwnerName for the instance
func (instance *Msvm_VirtualSystemManagementServiceSettingData) GetPropertyPrimaryOwnerName() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryOwnerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
