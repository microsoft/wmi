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

// Msft_WmiProvider_PutInstanceAsyncEvent_Pre struct
type Msft_WmiProvider_PutInstanceAsyncEvent_Pre struct {
	*Msft_WmiProvider_OperationEvent_Pre

	//
	Flags uint32

	//
	InstanceObject interface{}
}

func NewMsft_WmiProvider_PutInstanceAsyncEvent_PreEx1(instance *cim.WmiInstance) (newInstance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PreEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_PutInstanceAsyncEvent_Pre{
		Msft_WmiProvider_OperationEvent_Pre: tmp,
	}
	return
}

func NewMsft_WmiProvider_PutInstanceAsyncEvent_PreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PreEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_PutInstanceAsyncEvent_Pre{
		Msft_WmiProvider_OperationEvent_Pre: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceObject sets the value of InstanceObject for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre) SetPropertyInstanceObject(value interface{}) (err error) {
	return instance.SetProperty("InstanceObject", value)
}

// GetInstanceObject gets the value of InstanceObject for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Pre) GetPropertyInstanceObject() (value interface{}, err error) {
	retValue, err := instance.GetProperty("InstanceObject")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
