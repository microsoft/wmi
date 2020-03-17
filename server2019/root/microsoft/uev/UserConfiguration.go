// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Uev
//////////////////////////////////////////////
package uev

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// UserConfiguration struct
type UserConfiguration struct {
	cim.WmiInstance

	// Flag for enabling / disabling sync settings for Windows apps
	DontSyncWindows8AppSettings bool

	// Max package size (in bytes)
	MaxPackageSizeInBytes uint32

	// Delay in seconds before displaying the notification on settings import
	SettingsImportNotifyDelayInSeconds uint32

	// Flag for displaying the notification on settings import
	SettingsImportNotifyEnabled bool

	// Absolute path to the settings storage path
	SettingsStoragePath string

	// Sync enablement flag
	SyncEnabled bool

	// Synchronization method
	SyncMethod string

	// Sync over metered network flag
	SyncOverMeteredNetwork bool

	// Sync over metered network when roaming flag
	SyncOverMeteredNetworkWhenRoaming bool

	// Enable ping of the sync provider
	SyncProviderPingEnabled bool

	// Timeout for synchronization from the settings repository (in milliseconds)
	SyncTimeoutInMilliseconds uint32

	// Flag for enabling / disabling default sync settings for Windows apps
	SyncUnlistedWindows8Apps bool

	// VDI collection name for current computer
	VdiCollectionName string

	// Wait for sync when starting an application
	WaitForSyncOnApplicationStart bool

	// Wait for sync when logging on
	WaitForSyncOnLogon bool

	// Wait timeout for synchronization from the settings repository (in milliseconds)
	WaitForSyncTimeoutInMilliseconds uint32
}

// SetDontSyncWindows8AppSettings sets the value of DontSyncWindows8AppSettings for the instance
func (instance *UserConfiguration) SetPropertyDontSyncWindows8AppSettings(value bool) (err error) {
	return instance.SetProperty("DontSyncWindows8AppSettings", value)
}

// GetDontSyncWindows8AppSettings gets the value of DontSyncWindows8AppSettings for the instance
func (instance *UserConfiguration) GetPropertyDontSyncWindows8AppSettings() (value bool, err error) {
	retValue, err := instance.GetProperty("DontSyncWindows8AppSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxPackageSizeInBytes sets the value of MaxPackageSizeInBytes for the instance
func (instance *UserConfiguration) SetPropertyMaxPackageSizeInBytes(value uint32) (err error) {
	return instance.SetProperty("MaxPackageSizeInBytes", value)
}

// GetMaxPackageSizeInBytes gets the value of MaxPackageSizeInBytes for the instance
func (instance *UserConfiguration) GetPropertyMaxPackageSizeInBytes() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxPackageSizeInBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingsImportNotifyDelayInSeconds sets the value of SettingsImportNotifyDelayInSeconds for the instance
func (instance *UserConfiguration) SetPropertySettingsImportNotifyDelayInSeconds(value uint32) (err error) {
	return instance.SetProperty("SettingsImportNotifyDelayInSeconds", value)
}

// GetSettingsImportNotifyDelayInSeconds gets the value of SettingsImportNotifyDelayInSeconds for the instance
func (instance *UserConfiguration) GetPropertySettingsImportNotifyDelayInSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("SettingsImportNotifyDelayInSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingsImportNotifyEnabled sets the value of SettingsImportNotifyEnabled for the instance
func (instance *UserConfiguration) SetPropertySettingsImportNotifyEnabled(value bool) (err error) {
	return instance.SetProperty("SettingsImportNotifyEnabled", value)
}

// GetSettingsImportNotifyEnabled gets the value of SettingsImportNotifyEnabled for the instance
func (instance *UserConfiguration) GetPropertySettingsImportNotifyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("SettingsImportNotifyEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingsStoragePath sets the value of SettingsStoragePath for the instance
func (instance *UserConfiguration) SetPropertySettingsStoragePath(value string) (err error) {
	return instance.SetProperty("SettingsStoragePath", value)
}

// GetSettingsStoragePath gets the value of SettingsStoragePath for the instance
func (instance *UserConfiguration) GetPropertySettingsStoragePath() (value string, err error) {
	retValue, err := instance.GetProperty("SettingsStoragePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncEnabled sets the value of SyncEnabled for the instance
func (instance *UserConfiguration) SetPropertySyncEnabled(value bool) (err error) {
	return instance.SetProperty("SyncEnabled", value)
}

// GetSyncEnabled gets the value of SyncEnabled for the instance
func (instance *UserConfiguration) GetPropertySyncEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("SyncEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncMethod sets the value of SyncMethod for the instance
func (instance *UserConfiguration) SetPropertySyncMethod(value string) (err error) {
	return instance.SetProperty("SyncMethod", value)
}

// GetSyncMethod gets the value of SyncMethod for the instance
func (instance *UserConfiguration) GetPropertySyncMethod() (value string, err error) {
	retValue, err := instance.GetProperty("SyncMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncOverMeteredNetwork sets the value of SyncOverMeteredNetwork for the instance
func (instance *UserConfiguration) SetPropertySyncOverMeteredNetwork(value bool) (err error) {
	return instance.SetProperty("SyncOverMeteredNetwork", value)
}

// GetSyncOverMeteredNetwork gets the value of SyncOverMeteredNetwork for the instance
func (instance *UserConfiguration) GetPropertySyncOverMeteredNetwork() (value bool, err error) {
	retValue, err := instance.GetProperty("SyncOverMeteredNetwork")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncOverMeteredNetworkWhenRoaming sets the value of SyncOverMeteredNetworkWhenRoaming for the instance
func (instance *UserConfiguration) SetPropertySyncOverMeteredNetworkWhenRoaming(value bool) (err error) {
	return instance.SetProperty("SyncOverMeteredNetworkWhenRoaming", value)
}

// GetSyncOverMeteredNetworkWhenRoaming gets the value of SyncOverMeteredNetworkWhenRoaming for the instance
func (instance *UserConfiguration) GetPropertySyncOverMeteredNetworkWhenRoaming() (value bool, err error) {
	retValue, err := instance.GetProperty("SyncOverMeteredNetworkWhenRoaming")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncProviderPingEnabled sets the value of SyncProviderPingEnabled for the instance
func (instance *UserConfiguration) SetPropertySyncProviderPingEnabled(value bool) (err error) {
	return instance.SetProperty("SyncProviderPingEnabled", value)
}

// GetSyncProviderPingEnabled gets the value of SyncProviderPingEnabled for the instance
func (instance *UserConfiguration) GetPropertySyncProviderPingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("SyncProviderPingEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncTimeoutInMilliseconds sets the value of SyncTimeoutInMilliseconds for the instance
func (instance *UserConfiguration) SetPropertySyncTimeoutInMilliseconds(value uint32) (err error) {
	return instance.SetProperty("SyncTimeoutInMilliseconds", value)
}

// GetSyncTimeoutInMilliseconds gets the value of SyncTimeoutInMilliseconds for the instance
func (instance *UserConfiguration) GetPropertySyncTimeoutInMilliseconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("SyncTimeoutInMilliseconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncUnlistedWindows8Apps sets the value of SyncUnlistedWindows8Apps for the instance
func (instance *UserConfiguration) SetPropertySyncUnlistedWindows8Apps(value bool) (err error) {
	return instance.SetProperty("SyncUnlistedWindows8Apps", value)
}

// GetSyncUnlistedWindows8Apps gets the value of SyncUnlistedWindows8Apps for the instance
func (instance *UserConfiguration) GetPropertySyncUnlistedWindows8Apps() (value bool, err error) {
	retValue, err := instance.GetProperty("SyncUnlistedWindows8Apps")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVdiCollectionName sets the value of VdiCollectionName for the instance
func (instance *UserConfiguration) SetPropertyVdiCollectionName(value string) (err error) {
	return instance.SetProperty("VdiCollectionName", value)
}

// GetVdiCollectionName gets the value of VdiCollectionName for the instance
func (instance *UserConfiguration) GetPropertyVdiCollectionName() (value string, err error) {
	retValue, err := instance.GetProperty("VdiCollectionName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWaitForSyncOnApplicationStart sets the value of WaitForSyncOnApplicationStart for the instance
func (instance *UserConfiguration) SetPropertyWaitForSyncOnApplicationStart(value bool) (err error) {
	return instance.SetProperty("WaitForSyncOnApplicationStart", value)
}

// GetWaitForSyncOnApplicationStart gets the value of WaitForSyncOnApplicationStart for the instance
func (instance *UserConfiguration) GetPropertyWaitForSyncOnApplicationStart() (value bool, err error) {
	retValue, err := instance.GetProperty("WaitForSyncOnApplicationStart")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWaitForSyncOnLogon sets the value of WaitForSyncOnLogon for the instance
func (instance *UserConfiguration) SetPropertyWaitForSyncOnLogon(value bool) (err error) {
	return instance.SetProperty("WaitForSyncOnLogon", value)
}

// GetWaitForSyncOnLogon gets the value of WaitForSyncOnLogon for the instance
func (instance *UserConfiguration) GetPropertyWaitForSyncOnLogon() (value bool, err error) {
	retValue, err := instance.GetProperty("WaitForSyncOnLogon")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWaitForSyncTimeoutInMilliseconds sets the value of WaitForSyncTimeoutInMilliseconds for the instance
func (instance *UserConfiguration) SetPropertyWaitForSyncTimeoutInMilliseconds(value uint32) (err error) {
	return instance.SetProperty("WaitForSyncTimeoutInMilliseconds", value)
}

// GetWaitForSyncTimeoutInMilliseconds gets the value of WaitForSyncTimeoutInMilliseconds for the instance
func (instance *UserConfiguration) GetPropertyWaitForSyncTimeoutInMilliseconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("WaitForSyncTimeoutInMilliseconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
