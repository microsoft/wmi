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

// BcdStringElement struct
type BcdStringElement struct {
	*BcdElement

	// This is the string value of the element.
	String string
}

func NewBcdStringElementEx1(instance *cim.WmiInstance) (newInstance *BcdStringElement, err error) {
	tmp, err := NewBcdElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdStringElement{
		BcdElement: tmp,
	}
	return
}

func NewBcdStringElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdStringElement, err error) {
	tmp, err := NewBcdElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdStringElement{
		BcdElement: tmp,
	}
	return
}

// SetString sets the value of String for the instance
func (instance *BcdStringElement) SetPropertyString(value string) (err error) {
	return instance.SetProperty("String", (value))
}

// GetString gets the value of String for the instance
func (instance *BcdStringElement) GetPropertyString() (value string, err error) {
	retValue, err := instance.GetProperty("String")
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
