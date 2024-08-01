// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_Expression struct
type MSFT_Expression struct {
	*cim.WmiInstance

	//
	SourceInfo string

	//
	SourceLines []string
}

func NewMSFT_ExpressionEx1(instance *cim.WmiInstance) (newInstance *MSFT_Expression, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_Expression{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ExpressionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_Expression, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_Expression{
		WmiInstance: tmp,
	}
	return
}

// SetSourceInfo sets the value of SourceInfo for the instance
func (instance *MSFT_Expression) SetPropertySourceInfo(value string) (err error) {
	return instance.SetProperty("SourceInfo", (value))
}

// GetSourceInfo gets the value of SourceInfo for the instance
func (instance *MSFT_Expression) GetPropertySourceInfo() (value string, err error) {
	retValue, err := instance.GetProperty("SourceInfo")
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

// SetSourceLines sets the value of SourceLines for the instance
func (instance *MSFT_Expression) SetPropertySourceLines(value []string) (err error) {
	return instance.SetProperty("SourceLines", (value))
}

// GetSourceLines gets the value of SourceLines for the instance
func (instance *MSFT_Expression) GetPropertySourceLines() (value []string, err error) {
	retValue, err := instance.GetProperty("SourceLines")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
