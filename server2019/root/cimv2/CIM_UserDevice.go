// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_UserDevice struct
type CIM_UserDevice struct {
	CIM_LogicalDevice

	//
	IsLocked bool
}

// SetIsLocked sets the value of IsLocked for the instance
func (instance *CIM_UserDevice) SetPropertyIsLocked(value bool) (err error) {
	return instance.SetProperty("IsLocked", value)
}

// GetIsLocked gets the value of IsLocked for the instance
func (instance *CIM_UserDevice) GetPropertyIsLocked() (value bool, err error) {
	retValue, err := instance.GetProperty("IsLocked")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
