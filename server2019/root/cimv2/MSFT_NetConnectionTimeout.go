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

// MSFT_NetConnectionTimeout struct
type MSFT_NetConnectionTimeout struct {
	*MSFT_SCMEventLogEvent

	//
	Milliseconds uint32

	//
	Service string
}

func NewMSFT_NetConnectionTimeoutEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetConnectionTimeout, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConnectionTimeout{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

func NewMSFT_NetConnectionTimeoutEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetConnectionTimeout, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConnectionTimeout{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

// SetMilliseconds sets the value of Milliseconds for the instance
func (instance *MSFT_NetConnectionTimeout) SetPropertyMilliseconds(value uint32) (err error) {
	return instance.SetProperty("Milliseconds", value)
}

// GetMilliseconds gets the value of Milliseconds for the instance
func (instance *MSFT_NetConnectionTimeout) GetPropertyMilliseconds() (value uint32, err error) {
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
func (instance *MSFT_NetConnectionTimeout) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetConnectionTimeout) GetPropertyService() (value string, err error) {
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
