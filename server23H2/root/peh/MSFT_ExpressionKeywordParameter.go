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

// MSFT_ExpressionKeywordParameter struct
type MSFT_ExpressionKeywordParameter struct {
	*MSFT_ExpressionIdentifier

	//
	keywordalias string
}

func NewMSFT_ExpressionKeywordParameterEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExpressionKeywordParameter, err error) {
	tmp, err := NewMSFT_ExpressionIdentifierEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionKeywordParameter{
		MSFT_ExpressionIdentifier: tmp,
	}
	return
}

func NewMSFT_ExpressionKeywordParameterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExpressionKeywordParameter, err error) {
	tmp, err := NewMSFT_ExpressionIdentifierEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExpressionKeywordParameter{
		MSFT_ExpressionIdentifier: tmp,
	}
	return
}

// Setkeywordalias sets the value of keywordalias for the instance
func (instance *MSFT_ExpressionKeywordParameter) SetPropertykeywordalias(value string) (err error) {
	return instance.SetProperty("keywordalias", (value))
}

// Getkeywordalias gets the value of keywordalias for the instance
func (instance *MSFT_ExpressionKeywordParameter) GetPropertykeywordalias() (value string, err error) {
	retValue, err := instance.GetProperty("keywordalias")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
