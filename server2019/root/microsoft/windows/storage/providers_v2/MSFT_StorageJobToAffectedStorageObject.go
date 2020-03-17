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

// MSFT_StorageJobToAffectedStorageObject struct
type MSFT_StorageJobToAffectedStorageObject struct {
	cim.WmiInstance

	//
	AffectedStorageObject MSFT_StorageObject

	//
	StorageJob MSFT_StorageJob
}

// SetAffectedStorageObject sets the value of AffectedStorageObject for the instance
func (instance *MSFT_StorageJobToAffectedStorageObject) SetPropertyAffectedStorageObject(value MSFT_StorageObject) (err error) {
	return instance.SetProperty("AffectedStorageObject", value)
}

// GetAffectedStorageObject gets the value of AffectedStorageObject for the instance
func (instance *MSFT_StorageJobToAffectedStorageObject) GetPropertyAffectedStorageObject() (value MSFT_StorageObject, err error) {
	retValue, err := instance.GetProperty("AffectedStorageObject")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageObject)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageJob sets the value of StorageJob for the instance
func (instance *MSFT_StorageJobToAffectedStorageObject) SetPropertyStorageJob(value MSFT_StorageJob) (err error) {
	return instance.SetProperty("StorageJob", value)
}

// GetStorageJob gets the value of StorageJob for the instance
func (instance *MSFT_StorageJobToAffectedStorageObject) GetPropertyStorageJob() (value MSFT_StorageJob, err error) {
	retValue, err := instance.GetProperty("StorageJob")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageJob)
	if !ok {
		// TODO: Set an error
	}
	return
}
