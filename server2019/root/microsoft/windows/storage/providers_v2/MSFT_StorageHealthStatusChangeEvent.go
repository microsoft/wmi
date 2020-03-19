// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageHealthStatusChangeEvent struct
type MSFT_StorageHealthStatusChangeEvent struct {
	*MSFT_StorageEvent

	// Denotes the current health status of the source object.
	/// 0 - 'Healthy': TBD.
	///1 - 'Warning': TBD.
	///2 - 'Unhealthy': TBD.
	CurrentHealthStatus StorageHealthStatusChangeEvent_CurrentHealthStatus

	// Denotes the previous health status of the source object.
	/// 0 - 'Healthy': TBD.
	///1 - 'Warning': TBD.
	///2 - 'Unhealthy': TBD.
	PreviousHealthStatus StorageHealthStatusChangeEvent_PreviousHealthStatus

	// A unique identifier for the source object.
	SourceUniqueId string

	// A globally unique identifier for the storage subsystem
	StorageSubsystemUniqueId string
}

func NewMSFT_StorageHealthStatusChangeEventEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageHealthStatusChangeEvent, err error) {
	tmp, err := NewMSFT_StorageEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageHealthStatusChangeEvent{
		MSFT_StorageEvent: tmp,
	}
	return
}

func NewMSFT_StorageHealthStatusChangeEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageHealthStatusChangeEvent, err error) {
	tmp, err := NewMSFT_StorageEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageHealthStatusChangeEvent{
		MSFT_StorageEvent: tmp,
	}
	return
}

// SetCurrentHealthStatus sets the value of CurrentHealthStatus for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) SetPropertyCurrentHealthStatus(value StorageHealthStatusChangeEvent_CurrentHealthStatus) (err error) {
	return instance.SetProperty("CurrentHealthStatus", value)
}

// GetCurrentHealthStatus gets the value of CurrentHealthStatus for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) GetPropertyCurrentHealthStatus() (value StorageHealthStatusChangeEvent_CurrentHealthStatus, err error) {
	retValue, err := instance.GetProperty("CurrentHealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageHealthStatusChangeEvent_CurrentHealthStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreviousHealthStatus sets the value of PreviousHealthStatus for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) SetPropertyPreviousHealthStatus(value StorageHealthStatusChangeEvent_PreviousHealthStatus) (err error) {
	return instance.SetProperty("PreviousHealthStatus", value)
}

// GetPreviousHealthStatus gets the value of PreviousHealthStatus for the instance
func (instance *MSFT_StorageHealthStatusChangeEvent) GetPropertyPreviousHealthStatus() (value StorageHealthStatusChangeEvent_PreviousHealthStatus, err error) {
	retValue, err := instance.GetProperty("PreviousHealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageHealthStatusChangeEvent_PreviousHealthStatus)
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
