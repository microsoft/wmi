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

// MSFT_StorageNodeToVolume struct
type MSFT_StorageNodeToVolume struct {
	cim.WmiInstance

	//
	StorageNode MSFT_StorageNode

	//
	Volume MSFT_Volume
}

// SetStorageNode sets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToVolume) SetPropertyStorageNode(value MSFT_StorageNode) (err error) {
	return instance.SetProperty("StorageNode", value)
}

// GetStorageNode gets the value of StorageNode for the instance
func (instance *MSFT_StorageNodeToVolume) GetPropertyStorageNode() (value MSFT_StorageNode, err error) {
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

// SetVolume sets the value of Volume for the instance
func (instance *MSFT_StorageNodeToVolume) SetPropertyVolume(value MSFT_Volume) (err error) {
	return instance.SetProperty("Volume", value)
}

// GetVolume gets the value of Volume for the instance
func (instance *MSFT_StorageNodeToVolume) GetPropertyVolume() (value MSFT_Volume, err error) {
	retValue, err := instance.GetProperty("Volume")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Volume)
	if !ok {
		// TODO: Set an error
	}
	return
}
