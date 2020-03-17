// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NCProvCancelQuery struct
type MSFT_NCProvCancelQuery struct {
	MSFT_NCProvEvent

	//
	ID uint32
}

// SetID sets the value of ID for the instance
func (instance *MSFT_NCProvCancelQuery) SetPropertyID(value uint32) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *MSFT_NCProvCancelQuery) GetPropertyID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
