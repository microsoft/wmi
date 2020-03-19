// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageProviderToStorageSubSystem struct
type MSFT_StorageProviderToStorageSubSystem struct {
	*cim.WmiInstance

	//
	StorageProvider MSFT_StorageProvider

	//
	StorageSubSystem MSFT_StorageSubSystem
}

func NewMSFT_StorageProviderToStorageSubSystemEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageProviderToStorageSubSystem, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageProviderToStorageSubSystem{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_StorageProviderToStorageSubSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageProviderToStorageSubSystem, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageProviderToStorageSubSystem{
		WmiInstance: tmp,
	}
	return
}

// SetStorageProvider sets the value of StorageProvider for the instance
func (instance *MSFT_StorageProviderToStorageSubSystem) SetPropertyStorageProvider(value MSFT_StorageProvider) (err error) {
	return instance.SetProperty("StorageProvider", value)
}

// GetStorageProvider gets the value of StorageProvider for the instance
func (instance *MSFT_StorageProviderToStorageSubSystem) GetPropertyStorageProvider() (value MSFT_StorageProvider, err error) {
	retValue, err := instance.GetProperty("StorageProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageProvider)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageProviderToStorageSubSystem) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageProviderToStorageSubSystem) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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
