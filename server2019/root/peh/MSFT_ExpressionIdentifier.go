// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionIdentifier struct
type MSFT_ExpressionIdentifier struct {
	MSFT_Expression

	//
	name string
}

// Setname sets the value of name for the instance
func (instance *MSFT_ExpressionIdentifier) SetPropertyname(value string) (err error) {
	return instance.SetProperty("name", value)
}

// Getname gets the value of name for the instance
func (instance *MSFT_ExpressionIdentifier) GetPropertyname() (value string, err error) {
	retValue, err := instance.GetProperty("name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
