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

// MSFT_ExpressionAssignment struct
type MSFT_ExpressionAssignment struct {
	*MSFT_Expression

	//
	lvalue MSFT_ExpressionIdentifier

	//
	rvalue MSFT_Expression
}

func NewMSFT_ExpressionAssignmentEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionAssignment, err error) {
	tmp, err := NewMSFT_ExpressionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionAssignment{
		MSFT_Expression: tmp,
	}
	return
}

func NewMSFT_ExpressionAssignmentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionAssignment, err error) {
	tmp, err := NewMSFT_ExpressionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionAssignment{
		MSFT_Expression: tmp,
	}
	return
}

// Setlvalue sets the value of lvalue for the instance
func (instance *MSFT_ExpressionAssignment) SetPropertylvalue(value MSFT_ExpressionIdentifier) (err error) {
	return instance.SetProperty("lvalue", value)
}

// Getlvalue gets the value of lvalue for the instance
func (instance *MSFT_ExpressionAssignment) GetPropertylvalue() (value MSFT_ExpressionIdentifier, err error) {
	retValue, err := instance.GetProperty("lvalue")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_ExpressionIdentifier)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setrvalue sets the value of rvalue for the instance
func (instance *MSFT_ExpressionAssignment) SetPropertyrvalue(value MSFT_Expression) (err error) {
	return instance.SetProperty("rvalue", value)
}

// Getrvalue gets the value of rvalue for the instance
func (instance *MSFT_ExpressionAssignment) GetPropertyrvalue() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("rvalue")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Expression)
	if !ok {
		// TODO: Set an error
	}
	return
}
