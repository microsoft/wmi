// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionLet struct
type MSFT_ExpressionLet struct {
	MSFT_Expression

	//
	body MSFT_Expression

	//
	initializers []MSFT_ExpressionAssignment
}

// Setbody sets the value of body for the instance
func (instance *MSFT_ExpressionLet) SetPropertybody(value MSFT_Expression) (err error) {
	return instance.SetProperty("body", value)
}

// Getbody gets the value of body for the instance
func (instance *MSFT_ExpressionLet) GetPropertybody() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("body")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Expression)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setinitializers sets the value of initializers for the instance
func (instance *MSFT_ExpressionLet) SetPropertyinitializers(value []MSFT_ExpressionAssignment) (err error) {
	return instance.SetProperty("initializers", value)
}

// Getinitializers gets the value of initializers for the instance
func (instance *MSFT_ExpressionLet) GetPropertyinitializers() (value []MSFT_ExpressionAssignment, err error) {
	retValue, err := instance.GetProperty("initializers")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_ExpressionAssignment)
	if !ok {
		// TODO: Set an error
	}
	return
}
