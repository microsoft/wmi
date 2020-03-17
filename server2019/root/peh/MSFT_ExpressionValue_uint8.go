// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionValue_uint8 struct
type MSFT_ExpressionValue_uint8 struct {
	MSFT_ExpressionValue

	//
	value uint8
}

// Setvalue sets the value of value for the instance
func (instance *MSFT_ExpressionValue_uint8) SetPropertyvalue(value uint8) (err error) {
	return instance.SetProperty("value", value)
}

// Getvalue gets the value of value for the instance
func (instance *MSFT_ExpressionValue_uint8) GetPropertyvalue() (value uint8, err error) {
	retValue, err := instance.GetProperty("value")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
