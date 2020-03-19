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

// MSFT_ExpressionLoop struct
type MSFT_ExpressionLoop struct {
	*MSFT_Expression

	//
	body MSFT_Expression

	//
	condition MSFT_Expression
}

func NewMSFT_ExpressionLoopEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionLoop, err error) {
	tmp, err := NewMSFT_ExpressionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionLoop{
		MSFT_Expression: tmp,
	}
	return
}

func NewMSFT_ExpressionLoopEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionLoop, err error) {
	tmp, err := NewMSFT_ExpressionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionLoop{
		MSFT_Expression: tmp,
	}
	return
}

// Setbody sets the value of body for the instance
func (instance *MSFT_ExpressionLoop) SetPropertybody(value MSFT_Expression) (err error) {
	return instance.SetProperty("body", value)
}

// Getbody gets the value of body for the instance
func (instance *MSFT_ExpressionLoop) GetPropertybody() (value MSFT_Expression, err error) {
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

// Setcondition sets the value of condition for the instance
func (instance *MSFT_ExpressionLoop) SetPropertycondition(value MSFT_Expression) (err error) {
	return instance.SetProperty("condition", value)
}

// Getcondition gets the value of condition for the instance
func (instance *MSFT_ExpressionLoop) GetPropertycondition() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("condition")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Expression)
	if !ok {
		// TODO: Set an error
	}
	return
}
