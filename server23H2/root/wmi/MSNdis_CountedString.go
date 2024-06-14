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

// MSNdis_CountedString struct
type MSNdis_CountedString struct {
	*MSNdis

	//
	Length uint16

	//
	String []byte
}

func NewMSNdis_CountedStringEx1(instance *cim.WmiInstance) (newInstance *MSNdis_CountedString, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_CountedString{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_CountedStringEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_CountedString, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_CountedString{
		MSNdis: tmp,
	}
	return
}

// SetLength sets the value of Length for the instance
func (instance *MSNdis_CountedString) SetPropertyLength(value uint16) (err error) {
	return instance.SetProperty("Length", (value))
}

// GetLength gets the value of Length for the instance
func (instance *MSNdis_CountedString) GetPropertyLength() (value uint16, err error) {
	retValue, err := instance.GetProperty("Length")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetString sets the value of String for the instance
func (instance *MSNdis_CountedString) SetPropertyString(value []byte) (err error) {
	return instance.SetProperty("String", (value))
}

// GetString gets the value of String for the instance
func (instance *MSNdis_CountedString) GetPropertyString() (value []byte, err error) {
	retValue, err := instance.GetProperty("String")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}
