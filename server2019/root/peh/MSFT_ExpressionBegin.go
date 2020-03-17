// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionBegin struct
type MSFT_ExpressionBegin struct {
	MSFT_Expression

	//
	body []MSFT_Expression
}

// Setbody sets the value of body for the instance
func (instance *MSFT_ExpressionBegin) SetPropertybody(value []MSFT_Expression) (err error) {
	return instance.SetProperty("body", value)
}

// Getbody gets the value of body for the instance
func (instance *MSFT_ExpressionBegin) GetPropertybody() (value []MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("body")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_Expression)
	if !ok {
		// TODO: Set an error
	}
	return
}
