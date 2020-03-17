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

// MSFT_StorageNodeToVirtualDisk struct
type MSFT_StorageNodeToVirtualDisk struct {
	cim.WmiInstance

	//
	StorageNode MSFT_StorageNode

	//
	VirtualDisk MSFT_VirtualDisk
}

// SetStorageNode sets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToVirtualDisk) SetPropertyStorageNode(value MSFT_StorageNode) (err error) {
	return instance.SetProperty("StorageNode", value)
}

// GetStorageNode gets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToVirtualDisk) GetPropertyStorageNode() (value MSFT_StorageNode, err error) {
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

// SetVirtualDisk sets the value of VirtualDisk for the instance
func (instance *MSFT_StorageNodeToVirtualDisk) SetPropertyVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("VirtualDisk", value)
}

// GetVirtualDisk gets the value of VirtualDisk for the instance
func (instance *MSFT_StorageNodeToVirtualDisk) GetPropertyVirtualDisk() (value MSFT_VirtualDisk, err error) {
	retValue, err := instance.GetProperty("VirtualDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_VirtualDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}
