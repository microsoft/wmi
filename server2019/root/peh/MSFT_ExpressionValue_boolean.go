// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionValue_boolean struct
type MSFT_ExpressionValue_boolean struct {
	MSFT_ExpressionValue

	//
	value bool
}

// Setvalue sets the value of value for the instance
func (instance *MSFT_ExpressionValue_boolean) SetPropertyvalue(value bool) (err error) {
	return instance.SetProperty("value", value)
}

// Getvalue gets the value of value for the instance
func (instance *MSFT_ExpressionValue_boolean) GetPropertyvalue() (value bool, err error) {
	retValue, err := instance.GetProperty("value")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
