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

// MSFT_StorageSubSystemToTargetPort struct
type MSFT_StorageSubSystemToTargetPort struct {
	cim.WmiInstance

	//
	StorageSubSystem MSFT_StorageSubSystem

	//
	TargetPort MSFT_TargetPort
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToTargetPort) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToTargetPort) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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

// SetTargetPort sets the value of TargetPort for the instance
func (instance *MSFT_StorageSubSystemToTargetPort) SetPropertyTargetPort(value MSFT_TargetPort) (err error) {
	return instance.SetProperty("TargetPort", value)
}

// GetTargetPort gets the value of TargetPort for the instance
func (instance *MSFT_StorageSubSystemToTargetPort) GetPropertyTargetPort() (value MSFT_TargetPort, err error) {
	retValue, err := instance.GetProperty("TargetPort")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_TargetPort)
	if !ok {
		// TODO: Set an error
	}
	return
}
