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

// MSFT_ExpressionLet struct
type MSFT_ExpressionLet struct {
	*MSFT_Expression

	//
	body MSFT_Expression

	//
	initializers []MSFT_ExpressionAssignment
}

func NewMSFT_ExpressionLetEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionLet, err error) {
	tmp, err := NewMSFT_ExpressionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionLet{
		MSFT_Expression: tmp,
	}
	return
}

func NewMSFT_ExpressionLetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionLet, err error) {
	tmp, err := NewMSFT_ExpressionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionLet{
		MSFT_Expression: tmp,
	}
	return
}

// Setbody sets the value of body for the instance
func (instance *MSFT_ExpressionLet) SetPropertybody(value MSFT_Expression) (err error) {
	return instance.SetProperty("body", (value))
}

// Getbody gets the value of body for the instance
func (instance *MSFT_ExpressionLet) GetPropertybody() (value MSFT_Expression, err error) {
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

// Setinitializers sets the value of initializers for the instance
func (instance *MSFT_ExpressionLet) SetPropertyinitializers(value []MSFT_ExpressionAssignment) (err error) {
	return instance.SetProperty("initializers", (value))
}

// Getinitializers gets the value of initializers for the instance
func (instance *MSFT_ExpressionLet) GetPropertyinitializers() (value []MSFT_ExpressionAssignment, err error) {
	retValue, err := instance.GetProperty("initializers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_ExpressionAssignment)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_ExpressionAssignment is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_ExpressionAssignment(valuetmp))
	}

	return
}
