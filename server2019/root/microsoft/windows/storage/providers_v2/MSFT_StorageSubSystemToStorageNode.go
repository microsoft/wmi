// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageSubSystemToStorageNode struct
type MSFT_StorageSubSystemToStorageNode struct {
	cim.WmiInstance

	//
	StorageNode MSFT_StorageNode

	//
	StorageSubSystem MSFT_StorageSubSystem
}

// SetStorageNode sets the value of StorageNode for the instance
func (instance *MSFT_StorageSubSystemToStorageNode) SetPropertyStorageNode(value MSFT_StorageNode) (err error) {
	return instance.SetProperty("StorageNode", value)
}

// GetStorageNode gets the value of StorageNode for the instance
func (instance *MSFT_StorageSubSystemToStorageNode) GetPropertyStorageNode() (value MSFT_StorageNode, err error) {
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

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageNode) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageNode) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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
