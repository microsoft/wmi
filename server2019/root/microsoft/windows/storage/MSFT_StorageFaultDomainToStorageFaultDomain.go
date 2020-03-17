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

// MSFT_StorageFaultDomainToStorageFaultDomain struct
type MSFT_StorageFaultDomainToStorageFaultDomain struct {
	cim.WmiInstance

	//
	SourceStorageFaultDomain MSFT_StorageFaultDomain

	//
	TargetStorageFaultDomain MSFT_StorageFaultDomain
}

// SetSourceStorageFaultDomain sets the value of SourceStorageFaultDomain for the instance
func (instance *MSFT_StorageFaultDomainToStorageFaultDomain) SetPropertySourceStorageFaultDomain(value MSFT_StorageFaultDomain) (err error) {
	return instance.SetProperty("SourceStorageFaultDomain", value)
}

// GetSourceStorageFaultDomain gets the value of SourceStorageFaultDomain for the instance
func (instance *MSFT_StorageFaultDomainToStorageFaultDomain) GetPropertySourceStorageFaultDomain() (value MSFT_StorageFaultDomain, err error) {
	retValue, err := instance.GetProperty("SourceStorageFaultDomain")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageFaultDomain)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetStorageFaultDomain sets the value of TargetStorageFaultDomain for the instance
func (instance *MSFT_StorageFaultDomainToStorageFaultDomain) SetPropertyTargetStorageFaultDomain(value MSFT_StorageFaultDomain) (err error) {
	return instance.SetProperty("TargetStorageFaultDomain", value)
}

// GetTargetStorageFaultDomain gets the value of TargetStorageFaultDomain for the instance
func (instance *MSFT_StorageFaultDomainToStorageFaultDomain) GetPropertyTargetStorageFaultDomain() (value MSFT_StorageFaultDomain, err error) {
	retValue, err := instance.GetProperty("TargetStorageFaultDomain")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageFaultDomain)
	if !ok {
		// TODO: Set an error
	}
	return
}
