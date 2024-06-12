// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BcdBooleanElement struct
type BcdBooleanElement struct {
	*BcdElement

	// This is the boolean value of the element.
	Boolean bool
}

func NewBcdBooleanElementEx1(instance *cim.WmiInstance) (newInstance *BcdBooleanElement, err error) {
	tmp, err := NewBcdElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdBooleanElement{
		BcdElement: tmp,
	}
	return
}

func NewBcdBooleanElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdBooleanElement, err error) {
	tmp, err := NewBcdElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdBooleanElement{
		BcdElement: tmp,
	}
	return
}

// SetBoolean sets the value of Boolean for the instance
func (instance *BcdBooleanElement) SetPropertyBoolean(value bool) (err error) {
	return instance.SetProperty("Boolean", (value))
}

// GetBoolean gets the value of Boolean for the instance
func (instance *BcdBooleanElement) GetPropertyBoolean() (value bool, err error) {
	retValue, err := instance.GetProperty("Boolean")
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
