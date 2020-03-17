// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_NonVolatileStorage struct
type CIM_NonVolatileStorage struct {
	CIM_Memory

	//
	IsWriteable bool
}

// SetIsWriteable sets the value of IsWriteable for the instance
func (instance *CIM_NonVolatileStorage) SetPropertyIsWriteable(value bool) (err error) {
	return instance.SetProperty("IsWriteable", value)
}

// GetIsWriteable gets the value of IsWriteable for the instance
func (instance *CIM_NonVolatileStorage) GetPropertyIsWriteable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsWriteable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
