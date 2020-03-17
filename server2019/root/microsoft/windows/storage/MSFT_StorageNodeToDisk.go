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

// MSFT_StorageNodeToDisk struct
type MSFT_StorageNodeToDisk struct {
	cim.WmiInstance

	//
	Disk MSFT_Disk

	//
	DiskNumber uint32

	//
	HealthStatus uint16

	//
	IsOffline bool

	//
	IsReadOnly bool

	//
	OfflineReason uint16

	//
	OperationalStatus []uint16

	//
	StorageNode MSFT_StorageNode
}

// SetDisk sets the value of Disk for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyDisk(value MSFT_Disk) (err error) {
	return instance.SetProperty("Disk", value)
}

// GetDisk gets the value of Disk for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyDisk() (value MSFT_Disk, err error) {
	retValue, err := instance.GetProperty("Disk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Disk)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiskNumber sets the value of DiskNumber for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyDiskNumber(value uint32) (err error) {
	return instance.SetProperty("DiskNumber", value)
}

// GetDiskNumber gets the value of DiskNumber for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyDiskNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DiskNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyHealthStatus(value uint16) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyHealthStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsOffline sets the value of IsOffline for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyIsOffline(value bool) (err error) {
	return instance.SetProperty("IsOffline", value)
}

// GetIsOffline gets the value of IsOffline for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyIsOffline() (value bool, err error) {
	retValue, err := instance.GetProperty("IsOffline")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsReadOnly sets the value of IsReadOnly for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyIsReadOnly(value bool) (err error) {
	return instance.SetProperty("IsReadOnly", value)
}

// GetIsReadOnly gets the value of IsReadOnly for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyIsReadOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("IsReadOnly")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOfflineReason sets the value of OfflineReason for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyOfflineReason(value uint16) (err error) {
	return instance.SetProperty("OfflineReason", value)
}

// GetOfflineReason gets the value of OfflineReason for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyOfflineReason() (value uint16, err error) {
	retValue, err := instance.GetProperty("OfflineReason")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyOperationalStatus(value []uint16) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyOperationalStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageNode sets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToDisk) SetPropertyStorageNode(value MSFT_StorageNode) (err error) {
	return instance.SetProperty("StorageNode", value)
}

// GetStorageNode gets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToDisk) GetPropertyStorageNode() (value MSFT_StorageNode, err error) {
	retValue, err := instance.GetProperty("StorageNode")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageNode)
	if !ok {
		// TODO: Set an error
	}
	return
}
