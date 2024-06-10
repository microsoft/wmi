// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_HealthActionEvent struct
type MSFT_HealthActionEvent struct {
	*MSFT_StorageEvent

	// The state change of the alert.
	ChangeType uint16

	// A unique identifier for the health action instance.
	HealthActionId string

	// A string that uniquely identifies the type of health action.
	HealthActionType string

	// The percentage of the action that has completed at the time that this value is requested.
	PercentComplete uint16

	// The short summary description of the health action.
	Reason string

	// The time that the action was started.
	StartTime string

	// Current high level state of the action.
	State uint16

	// A free-form string that provides implementation-specific status of the action.
	Status string

	// A globally unique identifier for the storage subsystem
	StorageSubsystemUniqueId string
}

func NewMSFT_HealthActionEventEx1(instance *cim.WmiInstance) (newInstance *MSFT_HealthActionEvent, err error) {
	tmp, err := NewMSFT_StorageEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_HealthActionEvent{
		MSFT_StorageEvent: tmp,
	}
	return
}

func NewMSFT_HealthActionEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_HealthActionEvent, err error) {
	tmp, err := NewMSFT_StorageEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_HealthActionEvent{
		MSFT_StorageEvent: tmp,
	}
	return
}

// SetChangeType sets the value of ChangeType for the instance
func (instance *MSFT_HealthActionEvent) SetPropertyChangeType(value uint16) (err error) {
	return instance.SetProperty("ChangeType", (value))
}

// GetChangeType gets the value of ChangeType for the instance
func (instance *MSFT_HealthActionEvent) GetPropertyChangeType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ChangeType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetHealthActionId sets the value of HealthActionId for the instance
func (instance *MSFT_HealthActionEvent) SetPropertyHealthActionId(value string) (err error) {
	return instance.SetProperty("HealthActionId", (value))
}

// GetHealthActionId gets the value of HealthActionId for the instance
func (instance *MSFT_HealthActionEvent) GetPropertyHealthActionId() (value string, err error) {
	retValue, err := instance.GetProperty("HealthActionId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetHealthActionType sets the value of HealthActionType for the instance
func (instance *MSFT_HealthActionEvent) SetPropertyHealthActionType(value string) (err error) {
	return instance.SetProperty("HealthActionType", (value))
}

// GetHealthActionType gets the value of HealthActionType for the instance
func (instance *MSFT_HealthActionEvent) GetPropertyHealthActionType() (value string, err error) {
	retValue, err := instance.GetProperty("HealthActionType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *MSFT_HealthActionEvent) SetPropertyPercentComplete(value uint16) (err error) {
	return instance.SetProperty("PercentComplete", (value))
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *MSFT_HealthActionEvent) GetPropertyPercentComplete() (value uint16, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReason sets the value of Reason for the instance
func (instance *MSFT_HealthActionEvent) SetPropertyReason(value string) (err error) {
	return instance.SetProperty("Reason", (value))
}

// GetReason gets the value of Reason for the instance
func (instance *MSFT_HealthActionEvent) GetPropertyReason() (value string, err error) {
	retValue, err := instance.GetProperty("Reason")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetStartTime sets the value of StartTime for the instance
func (instance *MSFT_HealthActionEvent) SetPropertyStartTime(value string) (err error) {
	return instance.SetProperty("StartTime", (value))
}

// GetStartTime gets the value of StartTime for the instance
func (instance *MSFT_HealthActionEvent) GetPropertyStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_HealthActionEvent) SetPropertyState(value uint16) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSFT_HealthActionEvent) GetPropertyState() (value uint16, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_HealthActionEvent) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_HealthActionEvent) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetStorageSubsystemUniqueId sets the value of StorageSubsystemUniqueId for the instance
func (instance *MSFT_HealthActionEvent) SetPropertyStorageSubsystemUniqueId(value string) (err error) {
	return instance.SetProperty("StorageSubsystemUniqueId", (value))
}

// GetStorageSubsystemUniqueId gets the value of StorageSubsystemUniqueId for the instance
func (instance *MSFT_HealthActionEvent) GetPropertyStorageSubsystemUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("StorageSubsystemUniqueId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
