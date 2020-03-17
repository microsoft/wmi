// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionValue_sint16 struct
type MSFT_ExpressionValue_sint16 struct {
	MSFT_ExpressionValue

	//
	value int16
}

// Setvalue sets the value of value for the instance
func (instance *MSFT_ExpressionValue_sint16) SetPropertyvalue(value int16) (err error) {
	return instance.SetProperty("value", value)
}

// Getvalue gets the value of value for the instance
func (instance *MSFT_ExpressionValue_sint16) GetPropertyvalue() (value int16, err error) {
	retValue, err := instance.GetProperty("value")
	if err != nil {
		return
	}
	value, ok := retValue.(int16)
	if !ok {
		// TODO: Set an error
	}
	return
}
