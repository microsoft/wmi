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

// MSFT_StorageSubSystemToReplicationCapabilities struct
type MSFT_StorageSubSystemToReplicationCapabilities struct {
	cim.WmiInstance

	//
	ReplicationCapabilities MSFT_ReplicationCapabilities

	//
	StorageSubSystem MSFT_StorageSubSystem
}

// SetReplicationCapabilities sets the value of ReplicationCapabilities for the instance
func (instance *MSFT_StorageSubSystemToReplicationCapabilities) SetPropertyReplicationCapabilities(value MSFT_ReplicationCapabilities) (err error) {
	return instance.SetProperty("ReplicationCapabilities", value)
}

// GetReplicationCapabilities gets the value of ReplicationCapabilities for the instance
func (instance *MSFT_StorageSubSystemToReplicationCapabilities) GetPropertyReplicationCapabilities() (value MSFT_ReplicationCapabilities, err error) {
	retValue, err := instance.GetProperty("ReplicationCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_ReplicationCapabilities)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToReplicationCapabilities) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToReplicationCapabilities) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
	retValue, err := instance.GetProperty("StorageSubSystem")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageSubSystem)
	if !ok {
		// TODO: Set an error
	}
	return
}
