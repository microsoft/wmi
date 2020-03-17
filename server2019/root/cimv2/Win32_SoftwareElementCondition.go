// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_SoftwareElementCondition struct
type Win32_SoftwareElementCondition struct {
	CIM_Check

	//
	Condition string
}

// SetCondition sets the value of Condition for the instance
func (instance *Win32_SoftwareElementCondition) SetPropertyCondition(value string) (err error) {
	return instance.SetProperty("Condition", value)
}

// GetCondition gets the value of Condition for the instance
func (instance *Win32_SoftwareElementCondition) GetPropertyCondition() (value string, err error) {
	retValue, err := instance.GetProperty("Condition")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
