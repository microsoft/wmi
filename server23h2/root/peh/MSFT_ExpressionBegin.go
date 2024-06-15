// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ExpressionBegin struct
type MSFT_ExpressionBegin struct {
	*MSFT_Expression

	//
	body []MSFT_Expression
}

func NewMSFT_ExpressionBeginEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionBegin, err error) {
	tmp, err := NewMSFT_ExpressionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionBegin{
		MSFT_Expression: tmp,
	}
	return
}

func NewMSFT_ExpressionBeginEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionBegin, err error) {
	tmp, err := NewMSFT_ExpressionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionBegin{
		MSFT_Expression: tmp,
	}
	return
}

// Setbody sets the value of body for the instance
func (instance *MSFT_ExpressionBegin) SetPropertybody(value []MSFT_Expression) (err error) {
	return instance.SetProperty("body", (value))
}

// Getbody gets the value of body for the instance
func (instance *MSFT_ExpressionBegin) GetPropertybody() (value []MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("body")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_Expression)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_Expression is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_Expression(valuetmp))
	}

	return
}
