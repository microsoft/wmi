// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ExpressionCall struct
type MSFT_ExpressionCall struct {
	*MSFT_Expression

	//
	arguments []MSFT_Expression

	//
	function MSFT_Expression

	//
	pipeline MSFT_Expression
}

func NewMSFT_ExpressionCallEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionCall, err error) {
	tmp, err := NewMSFT_ExpressionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionCall{
		MSFT_Expression: tmp,
	}
	return
}

func NewMSFT_ExpressionCallEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionCall, err error) {
	tmp, err := NewMSFT_ExpressionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionCall{
		MSFT_Expression: tmp,
	}
	return
}

// Setarguments sets the value of arguments for the instance
func (instance *MSFT_ExpressionCall) SetPropertyarguments(value []MSFT_Expression) (err error) {
	return instance.SetProperty("arguments", value)
}

// Getarguments gets the value of arguments for the instance
func (instance *MSFT_ExpressionCall) GetPropertyarguments() (value []MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("arguments")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_Expression)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setfunction sets the value of function for the instance
func (instance *MSFT_ExpressionCall) SetPropertyfunction(value MSFT_Expression) (err error) {
	return instance.SetProperty("function", value)
}

// Getfunction gets the value of function for the instance
func (instance *MSFT_ExpressionCall) GetPropertyfunction() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("function")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Expression)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setpipeline sets the value of pipeline for the instance
func (instance *MSFT_ExpressionCall) SetPropertypipeline(value MSFT_Expression) (err error) {
	return instance.SetProperty("pipeline", value)
}

// Getpipeline gets the value of pipeline for the instance
func (instance *MSFT_ExpressionCall) GetPropertypipeline() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("pipeline")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Expression)
	if !ok {
		// TODO: Set an error
	}
	return
}
