// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageFaultEvent struct
type MSFT_StorageFaultEvent struct {
	*MSFT_StorageEvent

	//
	ChangeType uint16

	//
	FaultId string

	//
	FaultingObjectDescription string

	//
	FaultingObjectLocation string

	//
	FaultingObjectType string

	//
	FaultingObjectUniqueId string

	//
	FaultType string

	//
	Reason string

	//
	RecommendedActions []string

	//
	SourceUniqueId string

	//
	StorageSubsystemUniqueId string
}

func NewMSFT_StorageFaultEventEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageFaultEvent, err error) {
	tmp, err := NewMSFT_StorageEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageFaultEvent{
		MSFT_StorageEvent: tmp,
	}
	return
}

func NewMSFT_StorageFaultEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageFaultEvent, err error) {
	tmp, err := NewMSFT_StorageEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageFaultEvent{
		MSFT_StorageEvent: tmp,
	}
	return
}

// SetChangeType sets the value of ChangeType for the instance
func (instance *MSFT_StorageFaultEvent) SetPropertyChangeType(value uint16) (err error) {
	return instance.SetProperty("ChangeType", value)
}

// GetChangeType gets the value of ChangeType for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertyChangeType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ChangeType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultId sets the value of FaultId for the instance
func (instance *MSFT_StorageFaultEvent) SetPropertyFaultId(value string) (err error) {
	return instance.SetProperty("FaultId", value)
}

// GetFaultId gets the value of FaultId for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertyFaultId() (value string, err error) {
	retValue, err := instance.GetProperty("FaultId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultingObjectDescription sets the value of FaultingObjectDescription for the instance
func (instance *MSFT_StorageFaultEvent) SetPropertyFaultingObjectDescription(value string) (err error) {
	return instance.SetProperty("FaultingObjectDescription", value)
}

// GetFaultingObjectDescription gets the value of FaultingObjectDescription for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertyFaultingObjectDescription() (value string, err error) {
	retValue, err := instance.GetProperty("FaultingObjectDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultingObjectLocation sets the value of FaultingObjectLocation for the instance
func (instance *MSFT_StorageFaultEvent) SetPropertyFaultingObjectLocation(value string) (err error) {
	return instance.SetProperty("FaultingObjectLocation", value)
}

// GetFaultingObjectLocation gets the value of FaultingObjectLocation for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertyFaultingObjectLocation() (value string, err error) {
	retValue, err := instance.GetProperty("FaultingObjectLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultingObjectType sets the value of FaultingObjectType for the instance
func (instance *MSFT_StorageFaultEvent) SetPropertyFaultingObjectType(value string) (err error) {
	return instance.SetProperty("FaultingObjectType", value)
}

// GetFaultingObjectType gets the value of FaultingObjectType for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertyFaultingObjectType() (value string, err error) {
	retValue, err := instance.GetProperty("FaultingObjectType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultingObjectUniqueId sets the value of FaultingObjectUniqueId for the instance
func (instance *MSFT_StorageFaultEvent) SetPropertyFaultingObjectUniqueId(value string) (err error) {
	return instance.SetProperty("FaultingObjectUniqueId", value)
}

// GetFaultingObjectUniqueId gets the value of FaultingObjectUniqueId for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertyFaultingObjectUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("FaultingObjectUniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultType sets the value of FaultType for the instance
func (instance *MSFT_StorageFaultEvent) SetPropertyFaultType(value string) (err error) {
	return instance.SetProperty("FaultType", value)
}

// GetFaultType gets the value of FaultType for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertyFaultType() (value string, err error) {
	retValue, err := instance.GetProperty("FaultType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReason sets the value of Reason for the instance
func (instance *MSFT_StorageFaultEvent) SetPropertyReason(value string) (err error) {
	return instance.SetProperty("Reason", value)
}

// GetReason gets the value of Reason for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertyReason() (value string, err error) {
	retValue, err := instance.GetProperty("Reason")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecommendedActions sets the value of RecommendedActions for the instance
func (instance *MSFT_StorageFaultEvent) SetPropertyRecommendedActions(value []string) (err error) {
	return instance.SetProperty("RecommendedActions", value)
}

// GetRecommendedActions gets the value of RecommendedActions for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertyRecommendedActions() (value []string, err error) {
	retValue, err := instance.GetProperty("RecommendedActions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceUniqueId sets the value of SourceUniqueId for the instance
func (instance *MSFT_StorageFaultEvent) SetPropertySourceUniqueId(value string) (err error) {
	return instance.SetProperty("SourceUniqueId", value)
}

// GetSourceUniqueId gets the value of SourceUniqueId for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertySourceUniqueId() (value string, err error) {
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
func (instance *MSFT_StorageFaultEvent) SetPropertyStorageSubsystemUniqueId(value string) (err error) {
	return instance.SetProperty("StorageSubsystemUniqueId", value)
}

// GetStorageSubsystemUniqueId gets the value of StorageSubsystemUniqueId for the instance
func (instance *MSFT_StorageFaultEvent) GetPropertyStorageSubsystemUniqueId() (value string, err error) {
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
