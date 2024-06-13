// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ExpressionIf struct
type MSFT_ExpressionIf struct {
	*MSFT_Expression

	//
	condition MSFT_Expression

	//
	falsecase MSFT_Expression

	//
	truecase MSFT_Expression
}

func NewMSFT_ExpressionIfEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionIf, err error) {
	tmp, err := NewMSFT_ExpressionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionIf{
		MSFT_Expression: tmp,
	}
	return
}

func NewMSFT_ExpressionIfEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionIf, err error) {
	tmp, err := NewMSFT_ExpressionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionIf{
		MSFT_Expression: tmp,
	}
	return
}

// Setcondition sets the value of condition for the instance
func (instance *MSFT_ExpressionIf) SetPropertycondition(value MSFT_Expression) (err error) {
	return instance.SetProperty("condition", (value))
}

// Getcondition gets the value of condition for the instance
func (instance *MSFT_ExpressionIf) GetPropertycondition() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("condition")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_Expression)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_Expression is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_Expression(valuetmp)

	return
}

// Setfalsecase sets the value of falsecase for the instance
func (instance *MSFT_ExpressionIf) SetPropertyfalsecase(value MSFT_Expression) (err error) {
	return instance.SetProperty("falsecase", (value))
}

// Getfalsecase gets the value of falsecase for the instance
func (instance *MSFT_ExpressionIf) GetPropertyfalsecase() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("falsecase")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_Expression)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_Expression is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_Expression(valuetmp)

	return
}

// Settruecase sets the value of truecase for the instance
func (instance *MSFT_ExpressionIf) SetPropertytruecase(value MSFT_Expression) (err error) {
	return instance.SetProperty("truecase", (value))
}

// Gettruecase gets the value of truecase for the instance
func (instance *MSFT_ExpressionIf) GetPropertytruecase() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("truecase")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_Expression)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_Expression is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_Expression(valuetmp)

	return
}
