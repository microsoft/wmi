// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionValue_sint32 struct
type MSFT_ExpressionValue_sint32 struct {
	MSFT_ExpressionValue

	//
	value int32
}

// Setvalue sets the value of value for the instance
func (instance *MSFT_ExpressionValue_sint32) SetPropertyvalue(value int32) (err error) {
	return instance.SetProperty("value", value)
}

// Getvalue gets the value of value for the instance
func (instance *MSFT_ExpressionValue_sint32) GetPropertyvalue() (value int32, err error) {
	retValue, err := instance.GetProperty("value")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
