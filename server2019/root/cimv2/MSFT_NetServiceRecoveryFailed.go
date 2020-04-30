// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetServiceRecoveryFailed struct
type MSFT_NetServiceRecoveryFailed struct {
	*MSFT_SCMEventLogEvent

	//
	Action string

	//
	ActionType uint32

	//
	Error uint32

	//
	Service string
}

func NewMSFT_NetServiceRecoveryFailedEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetServiceRecoveryFailed, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetServiceRecoveryFailed{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

func NewMSFT_NetServiceRecoveryFailedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetServiceRecoveryFailed, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetServiceRecoveryFailed{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

// SetAction sets the value of Action for the instance
func (instance *MSFT_NetServiceRecoveryFailed) SetPropertyAction(value string) (err error) {
	return instance.SetProperty("Action", value)
}

// GetAction gets the value of Action for the instance
func (instance *MSFT_NetServiceRecoveryFailed) GetPropertyAction() (value string, err error) {
	retValue, err := instance.GetProperty("Action")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActionType sets the value of ActionType for the instance
func (instance *MSFT_NetServiceRecoveryFailed) SetPropertyActionType(value uint32) (err error) {
	return instance.SetProperty("ActionType", value)
}

// GetActionType gets the value of ActionType for the instance
func (instance *MSFT_NetServiceRecoveryFailed) GetPropertyActionType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActionType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetError sets the value of Error for the instance
func (instance *MSFT_NetServiceRecoveryFailed) SetPropertyError(value uint32) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MSFT_NetServiceRecoveryFailed) GetPropertyError() (value uint32, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetServiceRecoveryFailed) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceRecoveryFailed) GetPropertyService() (value string, err error) {
	retValue, err := instance.GetProperty("Service")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}