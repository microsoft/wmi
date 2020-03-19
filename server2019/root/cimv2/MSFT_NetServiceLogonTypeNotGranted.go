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

// MSFT_NetServiceLogonTypeNotGranted struct
type MSFT_NetServiceLogonTypeNotGranted struct {
	*MSFT_SCMEventLogEvent

	//
	Account string

	//
	Error uint32

	//
	Service string
}

func NewMSFT_NetServiceLogonTypeNotGrantedEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetServiceLogonTypeNotGranted, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetServiceLogonTypeNotGranted{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

func NewMSFT_NetServiceLogonTypeNotGrantedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetServiceLogonTypeNotGranted, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetServiceLogonTypeNotGranted{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

// SetAccount sets the value of Account for the instance
func (instance *MSFT_NetServiceLogonTypeNotGranted) SetPropertyAccount(value string) (err error) {
	return instance.SetProperty("Account", value)
}

// GetAccount gets the value of Account for the instance
func (instance *MSFT_NetServiceLogonTypeNotGranted) GetPropertyAccount() (value string, err error) {
	retValue, err := instance.GetProperty("Account")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetError sets the value of Error for the instance
func (instance *MSFT_NetServiceLogonTypeNotGranted) SetPropertyError(value uint32) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MSFT_NetServiceLogonTypeNotGranted) GetPropertyError() (value uint32, err error) {
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
func (instance *MSFT_NetServiceLogonTypeNotGranted) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceLogonTypeNotGranted) GetPropertyService() (value string, err error) {
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
