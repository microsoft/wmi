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

// MSFT_ExpressionValue struct
type MSFT_ExpressionValue struct {
	*MSFT_Expression

	//
	hasValue bool
}

func NewMSFT_ExpressionValueEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionValue, err error) {
	tmp, err := NewMSFT_ExpressionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionValue{
		MSFT_Expression: tmp,
	}
	return
}

func NewMSFT_ExpressionValueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionValue, err error) {
	tmp, err := NewMSFT_ExpressionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionValue{
		MSFT_Expression: tmp,
	}
	return
}

// SethasValue sets the value of hasValue for the instance
func (instance *MSFT_ExpressionValue) SetPropertyhasValue(value bool) (err error) {
	return instance.SetProperty("hasValue", (value))
}

// GethasValue gets the value of hasValue for the instance
func (instance *MSFT_ExpressionValue) GetPropertyhasValue() (value bool, err error) {
	retValue, err := instance.GetProperty("hasValue")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
