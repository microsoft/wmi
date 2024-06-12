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

// BcdObjectListElement struct
type BcdObjectListElement struct {
	*BcdElement

	// This is the array of object ids this element refers to.
	Ids []string
}

func NewBcdObjectListElementEx1(instance *cim.WmiInstance) (newInstance *BcdObjectListElement, err error) {
	tmp, err := NewBcdElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdObjectListElement{
		BcdElement: tmp,
	}
	return
}

func NewBcdObjectListElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdObjectListElement, err error) {
	tmp, err := NewBcdElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdObjectListElement{
		BcdElement: tmp,
	}
	return
}

// SetIds sets the value of Ids for the instance
func (instance *BcdObjectListElement) SetPropertyIds(value []string) (err error) {
	return instance.SetProperty("Ids", (value))
}

// GetIds gets the value of Ids for the instance
func (instance *BcdObjectListElement) GetPropertyIds() (value []string, err error) {
	retValue, err := instance.GetProperty("Ids")
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
