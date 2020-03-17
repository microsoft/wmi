// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionKeywordValue struct
type MSFT_ExpressionKeywordValue struct {
	MSFT_Expression

	//
	keyword string
}

// Setkeyword sets the value of keyword for the instance
func (instance *MSFT_ExpressionKeywordValue) SetPropertykeyword(value string) (err error) {
	return instance.SetProperty("keyword", value)
}

// Getkeyword gets the value of keyword for the instance
func (instance *MSFT_ExpressionKeywordValue) GetPropertykeyword() (value string, err error) {
	retValue, err := instance.GetProperty("keyword")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
