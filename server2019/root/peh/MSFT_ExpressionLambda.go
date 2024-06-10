// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.PEH
//
// ////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ExpressionLambda struct
type MSFT_ExpressionLambda struct {
	*MSFT_Expression

	//
	body MSFT_Expression

	//
	parameters []MSFT_ExpressionIdentifier

	//
	pipeline MSFT_ExpressionIdentifier
}

func NewMSFT_ExpressionLambdaEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionLambda, err error) {
	tmp, err := NewMSFT_ExpressionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionLambda{
		MSFT_Expression: tmp,
	}
	return
}

func NewMSFT_ExpressionLambdaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionLambda, err error) {
	tmp, err := NewMSFT_ExpressionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionLambda{
		MSFT_Expression: tmp,
	}
	return
}

// Setbody sets the value of body for the instance
func (instance *MSFT_ExpressionLambda) SetPropertybody(value MSFT_Expression) (err error) {
	return instance.SetProperty("body", (value))
}

// Getbody gets the value of body for the instance
func (instance *MSFT_ExpressionLambda) GetPropertybody() (value MSFT_Expression, err error) {
	retValue, err := instance.GetProperty("body")
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

// Setparameters sets the value of parameters for the instance
func (instance *MSFT_ExpressionLambda) SetPropertyparameters(value []MSFT_ExpressionIdentifier) (err error) {
	return instance.SetProperty("parameters", (value))
}

// Getparameters gets the value of parameters for the instance
func (instance *MSFT_ExpressionLambda) GetPropertyparameters() (value []MSFT_ExpressionIdentifier, err error) {
	retValue, err := instance.GetProperty("parameters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_ExpressionIdentifier)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_ExpressionIdentifier is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_ExpressionIdentifier(valuetmp))
	}

	return
}

// Setpipeline sets the value of pipeline for the instance
func (instance *MSFT_ExpressionLambda) SetPropertypipeline(value MSFT_ExpressionIdentifier) (err error) {
	return instance.SetProperty("pipeline", (value))
}

// Getpipeline gets the value of pipeline for the instance
func (instance *MSFT_ExpressionLambda) GetPropertypipeline() (value MSFT_ExpressionIdentifier, err error) {
	retValue, err := instance.GetProperty("pipeline")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_ExpressionIdentifier)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_ExpressionIdentifier is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_ExpressionIdentifier(valuetmp)

	return
}
