// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ProgIDSpecification struct
type Win32_ProgIDSpecification struct {
	CIM_Check

	//
	Parent string

	//
	ProgID string
}

// SetParent sets the value of Parent for the instance
func (instance *Win32_ProgIDSpecification) SetPropertyParent(value string) (err error) {
	return instance.SetProperty("Parent", value)
}

// GetParent gets the value of Parent for the instance
func (instance *Win32_ProgIDSpecification) GetPropertyParent() (value string, err error) {
	retValue, err := instance.GetProperty("Parent")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProgID sets the value of ProgID for the instance
func (instance *Win32_ProgIDSpecification) SetPropertyProgID(value string) (err error) {
	return instance.SetProperty("ProgID", value)
}

// GetProgID gets the value of ProgID for the instance
func (instance *Win32_ProgIDSpecification) GetPropertyProgID() (value string, err error) {
	retValue, err := instance.GetProperty("ProgID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
