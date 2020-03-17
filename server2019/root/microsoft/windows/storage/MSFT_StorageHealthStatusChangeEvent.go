// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_StorageHealthStatusChangeEvent struct
type MSFT_StorageHealthStatusChangeEvent struct {
	MSFT_StorageEvent

	//
	CurrentHealthStatus uint16

	//
	PreviousHealthStatus uint16

	//
	SourceUniqueId string

	//
	StorageSubsystemUniqueId string
}

// SetCurrentHealthStatus sets the value of CurrentHealthStatus for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) SetPropertyCurrentHealthStatus(value uint16) (err error) {
	return instance.SetProperty("CurrentHealthStatus", value)
}

// GetCurrentHealthStatus gets the value of CurrentHealthStatus for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) GetPropertyCurrentHealthStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentHealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreviousHealthStatus sets the value of PreviousHealthStatus for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) SetPropertyPreviousHealthStatus(value uint16) (err error) {
	return instance.SetProperty("PreviousHealthStatus", value)
}

// GetPreviousHealthStatus gets the value of PreviousHealthStatus for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) GetPropertyPreviousHealthStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("PreviousHealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceUniqueId sets the value of SourceUniqueId for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) SetPropertySourceUniqueId(value string) (err error) {
	return instance.SetProperty("SourceUniqueId", value)
}

// GetSourceUniqueId gets the value of SourceUniqueId for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) GetPropertySourceUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("SourceUniqueId")
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
func (instance *MSFT_StorageHealthStatusChangeEvent) SetPropertyStorageSubsystemUniqueId(value string) (err error) {
	return instance.SetProperty("StorageSubsystemUniqueId", value)
}

// GetStorageSubsystemUniqueId gets the value of StorageSubsystemUniqueId for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) GetPropertyStorageSubsystemUniqueId() (value string, err error) {
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
