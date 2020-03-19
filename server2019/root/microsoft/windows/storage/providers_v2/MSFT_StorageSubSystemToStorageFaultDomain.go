// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageSubSystemToStorageFaultDomain struct
type MSFT_StorageSubSystemToStorageFaultDomain struct {
	*cim.WmiInstance

	//
	StorageFaultDomain MSFT_StorageFaultDomain

	//
	StorageSubSystem MSFT_StorageSubSystem
}

func NewMSFT_StorageSubSystemToStorageFaultDomainEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageSubSystemToStorageFaultDomain, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageSubSystemToStorageFaultDomain{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_StorageSubSystemToStorageFaultDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageSubSystemToStorageFaultDomain, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageSubSystemToStorageFaultDomain{
		WmiInstance: tmp,
	}
	return
}

// SetStorageFaultDomain sets the value of StorageFaultDomain for the instance
func (instance *MSFT_StorageSubSystemToStorageFaultDomain) SetPropertyStorageFaultDomain(value MSFT_StorageFaultDomain) (err error) {
	return instance.SetProperty("StorageFaultDomain", value)
}

// GetStorageFaultDomain gets the value of StorageFaultDomain for the instance
func (instance *MSFT_StorageSubSystemToStorageFaultDomain) GetPropertyStorageFaultDomain() (value MSFT_StorageFaultDomain, err error) {
	retValue, err := instance.GetProperty("StorageFaultDomain")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageFaultDomain)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageFaultDomain) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToStorageFaultDomain) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
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
