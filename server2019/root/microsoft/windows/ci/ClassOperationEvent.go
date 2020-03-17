// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.CI
//////////////////////////////////////////////
package ci

// __ClassOperationEvent struct
type __ClassOperationEvent struct {
	__Event

	//
	TargetClass interface{}
}

// SetTargetClass sets the value of TargetClass for the instance
func (instance *__ClassOperationEvent) SetPropertyTargetClass(value interface{}) (err error) {
	return instance.SetProperty("TargetClass", value)
}

// GetTargetClass gets the value of TargetClass for the instance
func (instance *__ClassOperationEvent) GetPropertyTargetClass() (value interface{}, err error) {
	retValue, err := instance.GetProperty("TargetClass")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
