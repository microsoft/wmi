// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_MintPipelineDefinition struct
type MSFT_MintPipelineDefinition struct {
	MSFT_PipelineDefinition

	//
	expression MSFT_Expression
}

// Setexpression sets the value of expression for the instance
func (instance *MSFT_MintPipelineDefinition) SetPropertyexpression(value MSFT_Expression) (err error) {
	return instance.SetProperty("expression", value)
}

// Getexpression gets the value of expression for the instance
func (instance *MSFT_MintPipelineDefinition) GetPropertyexpression() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("expression")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Expression)
	if !ok {
		// TODO: Set an error
	}
	return
}
