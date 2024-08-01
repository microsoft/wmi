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

// MSMCAInfo_Entry struct
type MSMCAInfo_Entry struct {
	*MSMCAInfo

	//
	Data []uint8

	//
	Length uint32
}

func NewMSMCAInfo_EntryEx1(instance *cim.WmiInstance) (newInstance *MSMCAInfo_Entry, err error) {
	tmp, err := NewMSMCAInfoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMCAInfo_Entry{
		MSMCAInfo: tmp,
	}
	return
}

func NewMSMCAInfo_EntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMCAInfo_Entry, err error) {
	tmp, err := NewMSMCAInfoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMCAInfo_Entry{
		MSMCAInfo: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *MSMCAInfo_Entry) SetPropertyData(value []uint8) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *MSMCAInfo_Entry) GetPropertyData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetLength sets the value of Length for the instance
func (instance *MSMCAInfo_Entry) SetPropertyLength(value uint32) (err error) {
	return instance.SetProperty("Length", (value))
}

// GetLength gets the value of Length for the instance
func (instance *MSMCAInfo_Entry) GetPropertyLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("Length")
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
