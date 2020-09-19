// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ExpressionValue_sint16 struct
type MSFT_ExpressionValue_sint16 struct {
	*MSFT_ExpressionValue

	//
	value int16
}

func NewMSFT_ExpressionValue_sint16Ex1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionValue_sint16, err error) {
	tmp, err := NewMSFT_ExpressionValueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionValue_sint16{
		MSFT_ExpressionValue: tmp,
	}
	return
}

func NewMSFT_ExpressionValue_sint16Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionValue_sint16, err error) {
	tmp, err := NewMSFT_ExpressionValueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionValue_sint16{
		MSFT_ExpressionValue: tmp,
	}
	return
}

// Setvalue sets the value of value for the instance
func (instance *MSFT_ExpressionValue_sint16) SetPropertyvalue(value int16) (err error) {
	return instance.SetProperty("value", (value))
}

// Getvalue gets the value of value for the instance
func (instance *MSFT_ExpressionValue_sint16) GetPropertyvalue() (value int16, err error) {
	retValue, err := instance.GetProperty("value")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int16(valuetmp)

	return
}
