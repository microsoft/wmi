// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_OfflineFilesHealth struct
type Win32_OfflineFilesHealth struct {
	cim.WmiInstance

	// A DATETIME value, in string format, that represents the last time this folder was successfully synchronized to the Offline Files cache.
	LastSuccessfulSyncTime string

	// The status of the last attempt to synchronize this folder to the Offline Files cache.
	LastSyncStatus uint8

	// A DATETIME value, in string format, that represents the last time an attempt was made to synchronized this folder to the Offline Files cache, even if it was unsuccessful.
	LastSyncTime string

	// If true, the Offline Files feature is enabled for this folder.
	OfflineAccessEnabled bool

	// If true, the share is working in Online mode
	OnlineMode bool
}

// SetLastSuccessfulSyncTime sets the value of LastSuccessfulSyncTime for the instance
func (instance *Win32_OfflineFilesHealth) SetPropertyLastSuccessfulSyncTime(value string) (err error) {
	return instance.SetProperty("LastSuccessfulSyncTime", value)
}

// GetLastSuccessfulSyncTime gets the value of LastSuccessfulSyncTime for the instance
func (instance *Win32_OfflineFilesHealth) GetPropertyLastSuccessfulSyncTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastSuccessfulSyncTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastSyncStatus sets the value of LastSyncStatus for the instance
func (instance *Win32_OfflineFilesHealth) SetPropertyLastSyncStatus(value uint8) (err error) {
	return instance.SetProperty("LastSyncStatus", value)
}

// GetLastSyncStatus gets the value of LastSyncStatus for the instance
func (instance *Win32_OfflineFilesHealth) GetPropertyLastSyncStatus() (value uint8, err error) {
	retValue, err := instance.GetProperty("LastSyncStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastSyncTime sets the value of LastSyncTime for the instance
func (instance *Win32_OfflineFilesHealth) SetPropertyLastSyncTime(value string) (err error) {
	return instance.SetProperty("LastSyncTime", value)
}

// GetLastSyncTime gets the value of LastSyncTime for the instance
func (instance *Win32_OfflineFilesHealth) GetPropertyLastSyncTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastSyncTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOfflineAccessEnabled sets the value of OfflineAccessEnabled for the instance
func (instance *Win32_OfflineFilesHealth) SetPropertyOfflineAccessEnabled(value bool) (err error) {
	return instance.SetProperty("OfflineAccessEnabled", value)
}

// GetOfflineAccessEnabled gets the value of OfflineAccessEnabled for the instance
func (instance *Win32_OfflineFilesHealth) GetPropertyOfflineAccessEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("OfflineAccessEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOnlineMode sets the value of OnlineMode for the instance
func (instance *Win32_OfflineFilesHealth) SetPropertyOnlineMode(value bool) (err error) {
	return instance.SetProperty("OnlineMode", value)
}

// GetOnlineMode gets the value of OnlineMode for the instance
func (instance *Win32_OfflineFilesHealth) GetPropertyOnlineMode() (value bool, err error) {
	retValue, err := instance.GetProperty("OnlineMode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
