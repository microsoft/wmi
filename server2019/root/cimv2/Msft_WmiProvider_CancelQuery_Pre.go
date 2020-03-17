// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_CancelQuery_Pre struct
type Msft_WmiProvider_CancelQuery_Pre struct {
	Msft_WmiProvider_OperationEvent_Pre

	//
	QueryId uint32
}

// SetQueryId sets the value of QueryId for the instance
func (instance *Msft_WmiProvider_CancelQuery_Pre) SetPropertyQueryId(value uint32) (err error) {
	return instance.SetProperty("QueryId", value)
}

// GetQueryId gets the value of QueryId for the instance
func (instance *Msft_WmiProvider_CancelQuery_Pre) GetPropertyQueryId() (value uint32, err error) {
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
