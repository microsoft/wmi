// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BcdIntegerElement struct
type BcdIntegerElement struct {
	*BcdElement

	// This is the integer value of the element.
	Integer uint64
}

func NewBcdIntegerElementEx1(instance *cim.WmiInstance) (newInstance *BcdIntegerElement, err error) {
	tmp, err := NewBcdElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdIntegerElement{
		BcdElement: tmp,
	}
	return
}

func NewBcdIntegerElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdIntegerElement, err error) {
	tmp, err := NewBcdElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdIntegerElement{
		BcdElement: tmp,
	}
	return
}

// SetInteger sets the value of Integer for the instance
func (instance *BcdIntegerElement) SetPropertyInteger(value uint64) (err error) {
	return instance.SetProperty("Integer", (value))
}

// GetInteger gets the value of Integer for the instance
func (instance *BcdIntegerElement) GetPropertyInteger() (value uint64, err error) {
	retValue, err := instance.GetProperty("Integer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
