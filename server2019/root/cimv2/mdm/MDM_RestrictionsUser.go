// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_RestrictionsUser struct
type MDM_RestrictionsUser struct {
	cim.WmiInstance

	//
	key uint32

	//
	PCSettingsMeteredNetworkSyncEnabled bool

	//
	PCSettingsPasswordSyncEnabled bool

	//
	PCSettingsSyncEnabled bool
}

// Setkey sets the value of key for the instance
func (instance *MDM_RestrictionsUser) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", value)
}

// Getkey gets the value of key for the instance
func (instance *MDM_RestrictionsUser) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPCSettingsMeteredNetworkSyncEnabled sets the value of PCSettingsMeteredNetworkSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) SetPropertyPCSettingsMeteredNetworkSyncEnabled(value bool) (err error) {
	return instance.SetProperty("PCSettingsMeteredNetworkSyncEnabled", value)
}

// GetPCSettingsMeteredNetworkSyncEnabled gets the value of PCSettingsMeteredNetworkSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) GetPropertyPCSettingsMeteredNetworkSyncEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PCSettingsMeteredNetworkSyncEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPCSettingsPasswordSyncEnabled sets the value of PCSettingsPasswordSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) SetPropertyPCSettingsPasswordSyncEnabled(value bool) (err error) {
	return instance.SetProperty("PCSettingsPasswordSyncEnabled", value)
}

// GetPCSettingsPasswordSyncEnabled gets the value of PCSettingsPasswordSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) GetPropertyPCSettingsPasswordSyncEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PCSettingsPasswordSyncEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPCSettingsSyncEnabled sets the value of PCSettingsSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) SetPropertyPCSettingsSyncEnabled(value bool) (err error) {
	return instance.SetProperty("PCSettingsSyncEnabled", value)
}

// GetPCSettingsSyncEnabled gets the value of PCSettingsSyncEnabled for the instance
func (instance *MDM_RestrictionsUser) GetPropertyPCSettingsSyncEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PCSettingsSyncEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
