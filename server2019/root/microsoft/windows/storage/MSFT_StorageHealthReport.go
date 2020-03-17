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

// MSFT_StorageHealthReport struct
type MSFT_StorageHealthReport struct {
	cim.WmiInstance

	//
	Records []MSFT_HealthRecord

	//
	ReportedObjectUniqueId string

	//
	StorageSubsystemUniqueId string
}

// SetRecords sets the value of Records for the instance
func (instance *MSFT_StorageHealthReport) SetPropertyRecords(value []MSFT_HealthRecord) (err error) {
	return instance.SetProperty("Records", value)
}

// GetRecords gets the value of Records for the instance
func (instance *MSFT_StorageHealthReport) GetPropertyRecords() (value []MSFT_HealthRecord, err error) {
	retValue, err := instance.GetProperty("Records")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_HealthRecord)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReportedObjectUniqueId sets the value of ReportedObjectUniqueId for the instance
func (instance *MSFT_StorageHealthReport) SetPropertyReportedObjectUniqueId(value string) (err error) {
	return instance.SetProperty("ReportedObjectUniqueId", value)
}

// GetReportedObjectUniqueId gets the value of ReportedObjectUniqueId for the instance
func (instance *MSFT_StorageHealthReport) GetPropertyReportedObjectUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("ReportedObjectUniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubsystemUniqueId sets the value of StorageSubsystemUniqueId for the instance
func (instance *MSFT_StorageHealthReport) SetPropertyStorageSubsystemUniqueId(value string) (err error) {
	return instance.SetProperty("StorageSubsystemUniqueId", value)
}

// GetStorageSubsystemUniqueId gets the value of StorageSubsystemUniqueId for the instance
func (instance *MSFT_StorageHealthReport) GetPropertyStorageSubsystemUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("StorageSubsystemUniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
