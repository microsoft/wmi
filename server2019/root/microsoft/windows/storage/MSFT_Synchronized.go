// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_Synchronized struct
type MSFT_Synchronized struct {
	cim.WmiInstance

	//
	CopyMethodology uint16

	//
	CopyPriority uint16

	//
	CopyState uint16

	//
	CopyType uint16

	//
	PercentSynced uint16

	//
	ProgressStatus uint16

	//
	RecoveryPointObjective uint32

	//
	ReplicaType uint16

	//
	RequestedCopyState uint16

	//
	SyncMaintained bool

	//
	SyncMode uint16

	//
	SyncState uint16

	//
	SyncTime string

	//
	SyncType uint16
}

// SetCopyMethodology sets the value of CopyMethodology for the instance
func (instance *MSFT_Synchronized) SetPropertyCopyMethodology(value uint16) (err error) {
	return instance.SetProperty("CopyMethodology", value)
}

// GetCopyMethodology gets the value of CopyMethodology for the instance
func (instance *MSFT_Synchronized) GetPropertyCopyMethodology() (value uint16, err error) {
	retValue, err := instance.GetProperty("CopyMethodology")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCopyPriority sets the value of CopyPriority for the instance
func (instance *MSFT_Synchronized) SetPropertyCopyPriority(value uint16) (err error) {
	return instance.SetProperty("CopyPriority", value)
}

// GetCopyPriority gets the value of CopyPriority for the instance
func (instance *MSFT_Synchronized) GetPropertyCopyPriority() (value uint16, err error) {
	retValue, err := instance.GetProperty("CopyPriority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCopyState sets the value of CopyState for the instance
func (instance *MSFT_Synchronized) SetPropertyCopyState(value uint16) (err error) {
	return instance.SetProperty("CopyState", value)
}

// GetCopyState gets the value of CopyState for the instance
func (instance *MSFT_Synchronized) GetPropertyCopyState() (value uint16, err error) {
	retValue, err := instance.GetProperty("CopyState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCopyType sets the value of CopyType for the instance
func (instance *MSFT_Synchronized) SetPropertyCopyType(value uint16) (err error) {
	return instance.SetProperty("CopyType", value)
}

// GetCopyType gets the value of CopyType for the instance
func (instance *MSFT_Synchronized) GetPropertyCopyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("CopyType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentSynced sets the value of PercentSynced for the instance
func (instance *MSFT_Synchronized) SetPropertyPercentSynced(value uint16) (err error) {
	return instance.SetProperty("PercentSynced", value)
}

// GetPercentSynced gets the value of PercentSynced for the instance
func (instance *MSFT_Synchronized) GetPropertyPercentSynced() (value uint16, err error) {
	retValue, err := instance.GetProperty("PercentSynced")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProgressStatus sets the value of ProgressStatus for the instance
func (instance *MSFT_Synchronized) SetPropertyProgressStatus(value uint16) (err error) {
	return instance.SetProperty("ProgressStatus", value)
}

// GetProgressStatus gets the value of ProgressStatus for the instance
func (instance *MSFT_Synchronized) GetPropertyProgressStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProgressStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecoveryPointObjective sets the value of RecoveryPointObjective for the instance
func (instance *MSFT_Synchronized) SetPropertyRecoveryPointObjective(value uint32) (err error) {
	return instance.SetProperty("RecoveryPointObjective", value)
}

// GetRecoveryPointObjective gets the value of RecoveryPointObjective for the instance
func (instance *MSFT_Synchronized) GetPropertyRecoveryPointObjective() (value uint32, err error) {
	retValue, err := instance.GetProperty("RecoveryPointObjective")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicaType sets the value of ReplicaType for the instance
func (instance *MSFT_Synchronized) SetPropertyReplicaType(value uint16) (err error) {
	return instance.SetProperty("ReplicaType", value)
}

// GetReplicaType gets the value of ReplicaType for the instance
func (instance *MSFT_Synchronized) GetPropertyReplicaType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicaType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestedCopyState sets the value of RequestedCopyState for the instance
func (instance *MSFT_Synchronized) SetPropertyRequestedCopyState(value uint16) (err error) {
	return instance.SetProperty("RequestedCopyState", value)
}

// GetRequestedCopyState gets the value of RequestedCopyState for the instance
func (instance *MSFT_Synchronized) GetPropertyRequestedCopyState() (value uint16, err error) {
	retValue, err := instance.GetProperty("RequestedCopyState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncMaintained sets the value of SyncMaintained for the instance
func (instance *MSFT_Synchronized) SetPropertySyncMaintained(value bool) (err error) {
	return instance.SetProperty("SyncMaintained", value)
}

// GetSyncMaintained gets the value of SyncMaintained for the instance
func (instance *MSFT_Synchronized) GetPropertySyncMaintained() (value bool, err error) {
	retValue, err := instance.GetProperty("SyncMaintained")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncMode sets the value of SyncMode for the instance
func (instance *MSFT_Synchronized) SetPropertySyncMode(value uint16) (err error) {
	return instance.SetProperty("SyncMode", value)
}

// GetSyncMode gets the value of SyncMode for the instance
func (instance *MSFT_Synchronized) GetPropertySyncMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("SyncMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncState sets the value of SyncState for the instance
func (instance *MSFT_Synchronized) SetPropertySyncState(value uint16) (err error) {
	return instance.SetProperty("SyncState", value)
}

// GetSyncState gets the value of SyncState for the instance
func (instance *MSFT_Synchronized) GetPropertySyncState() (value uint16, err error) {
	retValue, err := instance.GetProperty("SyncState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncTime sets the value of SyncTime for the instance
func (instance *MSFT_Synchronized) SetPropertySyncTime(value string) (err error) {
	return instance.SetProperty("SyncTime", value)
}

// GetSyncTime gets the value of SyncTime for the instance
func (instance *MSFT_Synchronized) GetPropertySyncTime() (value string, err error) {
	retValue, err := instance.GetProperty("SyncTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSyncType sets the value of SyncType for the instance
func (instance *MSFT_Synchronized) SetPropertySyncType(value uint16) (err error) {
	return instance.SetProperty("SyncType", value)
}

// GetSyncType gets the value of SyncType for the instance
func (instance *MSFT_Synchronized) GetPropertySyncType() (value uint16, err error) {
	retValue, err := instance.GetProperty("SyncType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
