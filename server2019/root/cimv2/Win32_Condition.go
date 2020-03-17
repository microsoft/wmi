// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Condition struct
type Win32_Condition struct {
	CIM_Check

	//
	Condition string

	//
	Feature string

	//
	Level uint16
}

// SetCondition sets the value of Condition for the instance
func (instance *Win32_Condition) SetPropertyCondition(value string) (err error) {
	return instance.SetProperty("Condition", value)
}

// GetCondition gets the value of Condition for the instance
func (instance *Win32_Condition) GetPropertyCondition() (value string, err error) {
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

// SetFeature sets the value of Feature for the instance
func (instance *Win32_Condition) SetPropertyFeature(value string) (err error) {
	return instance.SetProperty("Feature", value)
}

// GetFeature gets the value of Feature for the instance
func (instance *Win32_Condition) GetPropertyFeature() (value string, err error) {
	retValue, err := instance.GetProperty("Feature")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLevel sets the value of Level for the instance
func (instance *Win32_Condition) SetPropertyLevel(value uint16) (err error) {
	return instance.SetProperty("Level", value)
}

// GetLevel gets the value of Level for the instance
func (instance *Win32_Condition) GetPropertyLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("Level")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
