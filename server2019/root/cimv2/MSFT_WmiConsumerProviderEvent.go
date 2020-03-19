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

// MSFT_WmiConsumerProviderEvent struct
type MSFT_WmiConsumerProviderEvent struct {
	*MSFT_WmiProviderEvent

	//
	Machine string
}

func NewMSFT_WmiConsumerProviderEventEx1(instance *cim.WmiInstance) (newInstance *MSFT_WmiConsumerProviderEvent, err error) {
	tmp, err := NewMSFT_WmiProviderEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_WmiConsumerProviderEvent{
		MSFT_WmiProviderEvent: tmp,
	}
	return
}

func NewMSFT_WmiConsumerProviderEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_WmiConsumerProviderEvent, err error) {
	tmp, err := NewMSFT_WmiProviderEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_WmiConsumerProviderEvent{
		MSFT_WmiProviderEvent: tmp,
	}
	return
}

// SetMachine sets the value of Machine for the instance
func (instance *MSFT_WmiConsumerProviderEvent) SetPropertyMachine(value string) (err error) {
	return instance.SetProperty("Machine", value)
}

// GetMachine gets the value of Machine for the instance
func (instance *MSFT_WmiConsumerProviderEvent) GetPropertyMachine() (value string, err error) {
	retValue, err := instance.GetProperty("Machine")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
