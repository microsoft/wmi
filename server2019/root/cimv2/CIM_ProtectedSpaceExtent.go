// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_ProtectedSpaceExtent struct
type CIM_ProtectedSpaceExtent struct {
	CIM_StorageExtent

	//
	UserDataStripeDepth uint64
}

// SetUserDataStripeDepth sets the value of UserDataStripeDepth for the instance
func (instance *CIM_ProtectedSpaceExtent) SetPropertyUserDataStripeDepth(value uint64) (err error) {
	return instance.SetProperty("UserDataStripeDepth", value)
}

// GetUserDataStripeDepth gets the value of UserDataStripeDepth for the instance
func (instance *CIM_ProtectedSpaceExtent) GetPropertyUserDataStripeDepth() (value uint64, err error) {
	retValue, err := instance.GetProperty("UserDataStripeDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
