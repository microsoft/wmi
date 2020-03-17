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

// MSFT_StorageSubSystemToInitiatorId struct
type MSFT_StorageSubSystemToInitiatorId struct {
	cim.WmiInstance

	//
	InitiatorId MSFT_InitiatorId

	//
	StorageSubSystem MSFT_StorageSubSystem
}

// SetInitiatorId sets the value of InitiatorId for the instance
func (instance *MSFT_StorageSubSystemToInitiatorId) SetPropertyInitiatorId(value MSFT_InitiatorId) (err error) {
	return instance.SetProperty("InitiatorId", value)
}

// GetInitiatorId gets the value of InitiatorId for the instance
func (instance *MSFT_StorageSubSystemToInitiatorId) GetPropertyInitiatorId() (value MSFT_InitiatorId, err error) {
	retValue, err := instance.GetProperty("InitiatorId")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_InitiatorId)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToInitiatorId) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToInitiatorId) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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
