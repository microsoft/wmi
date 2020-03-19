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

// MSFT_NetTransactTimeout struct
type MSFT_NetTransactTimeout struct {
	*MSFT_SCMEventLogEvent

	//
	Milliseconds uint32

	//
	Service string
}

func NewMSFT_NetTransactTimeoutEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetTransactTimeout, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetTransactTimeout{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

func NewMSFT_NetTransactTimeoutEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetTransactTimeout, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetTransactTimeout{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

// SetMilliseconds sets the value of Milliseconds for the instance
func (instance *MSFT_NetTransactTimeout) SetPropertyMilliseconds(value uint32) (err error) {
	return instance.SetProperty("Milliseconds", value)
}

// GetMilliseconds gets the value of Milliseconds for the instance
func (instance *MSFT_NetTransactTimeout) GetPropertyMilliseconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("Milliseconds")
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
func (instance *MSFT_NetTransactTimeout) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetTransactTimeout) GetPropertyService() (value string, err error) {
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
