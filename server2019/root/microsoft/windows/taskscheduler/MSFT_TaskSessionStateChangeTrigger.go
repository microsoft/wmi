// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_TaskSessionStateChangeTrigger struct
type MSFT_TaskSessionStateChangeTrigger struct {
	*MSFT_TaskTrigger

	//
	Delay string

	//
	StateChange uint32

	//
	UserId string
}

func NewMSFT_TaskSessionStateChangeTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskSessionStateChangeTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskSessionStateChangeTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

func NewMSFT_TaskSessionStateChangeTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskSessionStateChangeTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskSessionStateChangeTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

// SetDelay sets the value of Delay for the instance
func (instance *MSFT_TaskSessionStateChangeTrigger) SetPropertyDelay(value string) (err error) {
	return instance.SetProperty("Delay", value)
}

// GetDelay gets the value of Delay for the instance
func (instance *MSFT_TaskSessionStateChangeTrigger) GetPropertyDelay() (value string, err error) {
	retValue, err := instance.GetProperty("Delay")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStateChange sets the value of StateChange for the instance
func (instance *MSFT_TaskSessionStateChangeTrigger) SetPropertyStateChange(value uint32) (err error) {
	return instance.SetProperty("StateChange", value)
}

// GetStateChange gets the value of StateChange for the instance
func (instance *MSFT_TaskSessionStateChangeTrigger) GetPropertyStateChange() (value uint32, err error) {
	retValue, err := instance.GetProperty("StateChange")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserId sets the value of UserId for the instance
func (instance *MSFT_TaskSessionStateChangeTrigger) SetPropertyUserId(value string) (err error) {
	return instance.SetProperty("UserId", value)
}

// GetUserId gets the value of UserId for the instance
func (instance *MSFT_TaskSessionStateChangeTrigger) GetPropertyUserId() (value string, err error) {
	retValue, err := instance.GetProperty("UserId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
