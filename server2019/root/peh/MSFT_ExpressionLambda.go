// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

// MSFT_ExpressionLambda struct
type MSFT_ExpressionLambda struct {
	MSFT_Expression

	//
	body MSFT_Expression

	//
	parameters []MSFT_ExpressionIdentifier

	//
	pipeline MSFT_ExpressionIdentifier
}

// Setbody sets the value of body for the instance
func (instance *MSFT_ExpressionLambda) SetPropertybody(value MSFT_Expression) (err error) {
	return instance.SetProperty("body", value)
}

// Getbody gets the value of body for the instance
func (instance *MSFT_ExpressionLambda) GetPropertybody() (value MSFT_Expression, err error) {
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

// Setparameters sets the value of parameters for the instance
func (instance *MSFT_ExpressionLambda) SetPropertyparameters(value []MSFT_ExpressionIdentifier) (err error) {
	return instance.SetProperty("parameters", value)
}

// Getparameters gets the value of parameters for the instance
func (instance *MSFT_ExpressionLambda) GetPropertyparameters() (value []MSFT_ExpressionIdentifier, err error) {
	retValue, err := instance.GetProperty("parameters")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_ExpressionIdentifier)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setpipeline sets the value of pipeline for the instance
func (instance *MSFT_ExpressionLambda) SetPropertypipeline(value MSFT_ExpressionIdentifier) (err error) {
	return instance.SetProperty("pipeline", value)
}

// Getpipeline gets the value of pipeline for the instance
func (instance *MSFT_ExpressionLambda) GetPropertypipeline() (value MSFT_ExpressionIdentifier, err error) {
	retValue, err := instance.GetProperty("pipeline")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_ExpressionIdentifier)
	if !ok {
		// TODO: Set an error
	}
	return
}
