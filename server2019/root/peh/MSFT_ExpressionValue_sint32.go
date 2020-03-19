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

// MSFT_ExpressionValue_sint32 struct
type MSFT_ExpressionValue_sint32 struct {
	*MSFT_ExpressionValue

	//
	value int32
}

func NewMSFT_ExpressionValue_sint32Ex1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionValue_sint32, err error) {
	tmp, err := NewMSFT_ExpressionValueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionValue_sint32{
		MSFT_ExpressionValue: tmp,
	}
	return
}

func NewMSFT_ExpressionValue_sint32Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionValue_sint32, err error) {
	tmp, err := NewMSFT_ExpressionValueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionValue_sint32{
		MSFT_ExpressionValue: tmp,
	}
	return
}

// Setvalue sets the value of value for the instance
func (instance *MSFT_ExpressionValue_sint32) SetPropertyvalue(value int32) (err error) {
	return instance.SetProperty("value", value)
}

// Getvalue gets the value of value for the instance
func (instance *MSFT_ExpressionValue_sint32) GetPropertyvalue() (value int32, err error) {
	retValue, err := instance.GetProperty("value")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
