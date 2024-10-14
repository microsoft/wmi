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

// Mark_V0 struct
type Mark_V0 struct {
	*PerfInfo_V0

	//
	Message string

	//
	Padding []byte
}

func NewMark_V0Ex1(instance *cim.WmiInstance) (newInstance *Mark_V0, err error) {
	tmp, err := NewPerfInfo_V0Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Mark_V0{
		PerfInfo_V0: tmp,
	}
	return
}

func NewMark_V0Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Mark_V0, err error) {
	tmp, err := NewPerfInfo_V0Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Mark_V0{
		PerfInfo_V0: tmp,
	}
	return
}

// SetMessage sets the value of Message for the instance
func (instance *Mark_V0) SetPropertyMessage(value string) (err error) {
	return instance.SetProperty("Message", (value))
}

// GetMessage gets the value of Message for the instance
func (instance *Mark_V0) GetPropertyMessage() (value string, err error) {
	retValue, err := instance.GetProperty("Message")
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

// SetPadding sets the value of Padding for the instance
func (instance *Mark_V0) SetPropertyPadding(value []byte) (err error) {
	return instance.SetProperty("Padding", (value))
}

// GetPadding gets the value of Padding for the instance
func (instance *Mark_V0) GetPropertyPadding() (value []byte, err error) {
	retValue, err := instance.GetProperty("Padding")
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
