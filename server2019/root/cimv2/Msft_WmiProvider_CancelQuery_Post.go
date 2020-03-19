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

// Msft_WmiProvider_CancelQuery_Post struct
type Msft_WmiProvider_CancelQuery_Post struct {
	*Msft_WmiProvider_OperationEvent_Post

	//
	QueryId uint32

	//
	Result uint32
}

func NewMsft_WmiProvider_CancelQuery_PostEx1(instance *cim.WmiInstance) (newInstance *Msft_WmiProvider_CancelQuery_Post, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PostEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_CancelQuery_Post{
		Msft_WmiProvider_OperationEvent_Post: tmp,
	}
	return
}

func NewMsft_WmiProvider_CancelQuery_PostEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_WmiProvider_CancelQuery_Post, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PostEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_CancelQuery_Post{
		Msft_WmiProvider_OperationEvent_Post: tmp,
	}
	return
}

// SetQueryId sets the value of QueryId for the instance
func (instance *Msft_WmiProvider_CancelQuery_Post) SetPropertyQueryId(value uint32) (err error) {
	return instance.SetProperty("QueryId", value)
}

// GetQueryId gets the value of QueryId for the instance
func (instance *Msft_WmiProvider_CancelQuery_Post) GetPropertyQueryId() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueryId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResult sets the value of Result for the instance
func (instance *Msft_WmiProvider_CancelQuery_Post) SetPropertyResult(value uint32) (err error) {
	return instance.SetProperty("Result", value)
}

// GetResult gets the value of Result for the instance
func (instance *Msft_WmiProvider_CancelQuery_Post) GetPropertyResult() (value uint32, err error) {
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
