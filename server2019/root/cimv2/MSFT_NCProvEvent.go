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

// MSFT_NCProvEvent struct
type MSFT_NCProvEvent struct {
	*__ExtrinsicEvent

	//
	Namespace string

	//
	ProviderName string

	//
	Result uint32
}

func NewMSFT_NCProvEventEx1(instance *cim.WmiInstance) (newInstance *MSFT_NCProvEvent, err error) {
	tmp, err := New__ExtrinsicEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NCProvEvent{
		__ExtrinsicEvent: tmp,
	}
	return
}

func NewMSFT_NCProvEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NCProvEvent, err error) {
	tmp, err := New__ExtrinsicEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NCProvEvent{
		__ExtrinsicEvent: tmp,
	}
	return
}

// SetNamespace sets the value of Namespace for the instance
func (instance *MSFT_NCProvEvent) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", value)
}

// GetNamespace gets the value of Namespace for the instance
func (instance *MSFT_NCProvEvent) GetPropertyNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("Namespace")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProviderName sets the value of ProviderName for the instance
func (instance *MSFT_NCProvEvent) SetPropertyProviderName(value string) (err error) {
	return instance.SetProperty("ProviderName", value)
}

// GetProviderName gets the value of ProviderName for the instance
func (instance *MSFT_NCProvEvent) GetPropertyProviderName() (value string, err error) {
	retValue, err := instance.GetProperty("ProviderName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResult sets the value of Result for the instance
func (instance *MSFT_NCProvEvent) SetPropertyResult(value uint32) (err error) {
	return instance.SetProperty("Result", value)
}

// GetResult gets the value of Result for the instance
func (instance *MSFT_NCProvEvent) GetPropertyResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("Result")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
