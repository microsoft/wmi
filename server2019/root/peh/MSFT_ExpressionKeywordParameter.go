// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionKeywordParameter struct
type MSFT_ExpressionKeywordParameter struct {
	MSFT_ExpressionIdentifier

	//
	keywordalias string
}

// Setkeywordalias sets the value of keywordalias for the instance
func (instance *MSFT_ExpressionKeywordParameter) SetPropertykeywordalias(value string) (err error) {
	return instance.SetProperty("keywordalias", value)
}

// Getkeywordalias gets the value of keywordalias for the instance
func (instance *MSFT_ExpressionKeywordParameter) GetPropertykeywordalias() (value string, err error) {
	retValue, err := instance.GetProperty("keywordalias")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
