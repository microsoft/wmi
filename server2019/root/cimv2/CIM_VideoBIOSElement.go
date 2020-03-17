// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_VideoBIOSElement struct
type CIM_VideoBIOSElement struct {
	CIM_SoftwareElement

	//
	IsShadowed bool
}

// SetIsShadowed sets the value of IsShadowed for the instance
func (instance *CIM_VideoBIOSElement) SetPropertyIsShadowed(value bool) (err error) {
	return instance.SetProperty("IsShadowed", value)
}

// GetIsShadowed gets the value of IsShadowed for the instance
func (instance *CIM_VideoBIOSElement) GetPropertyIsShadowed() (value bool, err error) {
	retValue, err := instance.GetProperty("IsShadowed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
