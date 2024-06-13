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

// MuiTraceData_String struct
type MuiTraceData_String struct {
	*MuiTraceData

	//
	MuiLoadString string
}

func NewMuiTraceData_StringEx1(instance *cim.WmiInstance) (newInstance *MuiTraceData_String, err error) {
	tmp, err := NewMuiTraceDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MuiTraceData_String{
		MuiTraceData: tmp,
	}
	return
}

func NewMuiTraceData_StringEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MuiTraceData_String, err error) {
	tmp, err := NewMuiTraceDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MuiTraceData_String{
		MuiTraceData: tmp,
	}
	return
}

// SetMuiLoadString sets the value of MuiLoadString for the instance
func (instance *MuiTraceData_String) SetPropertyMuiLoadString(value string) (err error) {
	return instance.SetProperty("MuiLoadString", (value))
}

// GetMuiLoadString gets the value of MuiLoadString for the instance
func (instance *MuiTraceData_String) GetPropertyMuiLoadString() (value string, err error) {
	retValue, err := instance.GetProperty("MuiLoadString")
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
