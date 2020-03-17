// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionValue_datetime struct
type MSFT_ExpressionValue_datetime struct {
	MSFT_ExpressionValue

	//
	value string
}

// Setvalue sets the value of value for the instance
func (instance *MSFT_ExpressionValue_datetime) SetPropertyvalue(value string) (err error) {
	return instance.SetProperty("value", value)
}

// Getvalue gets the value of value for the instance
func (instance *MSFT_ExpressionValue_datetime) GetPropertyvalue() (value string, err error) {
	retValue, err := instance.GetProperty("value")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
