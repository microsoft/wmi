// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BcdUnknownElement struct
type BcdUnknownElement struct {
	*BcdElement

	// This is the actual value of the element inside the BCD store.
	ActualType uint32
}

func NewBcdUnknownElementEx1(instance *cim.WmiInstance) (newInstance *BcdUnknownElement, err error) {
	tmp, err := NewBcdElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdUnknownElement{
		BcdElement: tmp,
	}
	return
}

func NewBcdUnknownElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdUnknownElement, err error) {
	tmp, err := NewBcdElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdUnknownElement{
		BcdElement: tmp,
	}
	return
}

// SetActualType sets the value of ActualType for the instance
func (instance *BcdUnknownElement) SetPropertyActualType(value uint32) (err error) {
	return instance.SetProperty("ActualType", (value))
}

// GetActualType gets the value of ActualType for the instance
func (instance *BcdUnknownElement) GetPropertyActualType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActualType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
