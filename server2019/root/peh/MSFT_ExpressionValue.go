// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionValue struct
type MSFT_ExpressionValue struct {
	MSFT_Expression

	//
	hasValue bool
}

// SethasValue sets the value of hasValue for the instance
func (instance *MSFT_ExpressionValue) SetPropertyhasValue(value bool) (err error) {
	return instance.SetProperty("hasValue", value)
}

// GethasValue gets the value of hasValue for the instance
func (instance *MSFT_ExpressionValue) GetPropertyhasValue() (value bool, err error) {
	retValue, err := instance.GetProperty("hasValue")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
