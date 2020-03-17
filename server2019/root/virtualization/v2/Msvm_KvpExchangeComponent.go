// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_KvpExchangeComponent struct
type Msvm_KvpExchangeComponent struct {
	CIM_LogicalDevice

	//
	GuestExchangeItems []string

	//
	GuestIntrinsicExchangeItems []string
}

// SetGuestExchangeItems sets the value of GuestExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponent) SetPropertyGuestExchangeItems(value []string) (err error) {
	return instance.SetProperty("GuestExchangeItems", value)
}

// GetGuestExchangeItems gets the value of GuestExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponent) GetPropertyGuestExchangeItems() (value []string, err error) {
	retValue, err := instance.GetProperty("GuestExchangeItems")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuestIntrinsicExchangeItems sets the value of GuestIntrinsicExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponent) SetPropertyGuestIntrinsicExchangeItems(value []string) (err error) {
	return instance.SetProperty("GuestIntrinsicExchangeItems", value)
}

// GetGuestIntrinsicExchangeItems gets the value of GuestIntrinsicExchangeItems for the instance
func (instance *Msvm_KvpExchangeComponent) GetPropertyGuestIntrinsicExchangeItems() (value []string, err error) {
	retValue, err := instance.GetProperty("GuestIntrinsicExchangeItems")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
