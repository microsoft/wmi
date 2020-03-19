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

// MSFT_ExpressionKeywordValue struct
type MSFT_ExpressionKeywordValue struct {
	*MSFT_Expression

	//
	keyword string
}

func NewMSFT_ExpressionKeywordValueEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionKeywordValue, err error) {
	tmp, err := NewMSFT_ExpressionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionKeywordValue{
		MSFT_Expression: tmp,
	}
	return
}

func NewMSFT_ExpressionKeywordValueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionKeywordValue, err error) {
	tmp, err := NewMSFT_ExpressionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionKeywordValue{
		MSFT_Expression: tmp,
	}
	return
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
