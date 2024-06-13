// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_SerialPortSettingData struct
type Msvm_SerialPortSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	DebuggerMode bool
}

func NewMsvm_SerialPortSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_SerialPortSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SerialPortSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_SerialPortSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SerialPortSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SerialPortSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetDebuggerMode sets the value of DebuggerMode for the instance
func (instance *Msvm_SerialPortSettingData) SetPropertyDebuggerMode(value bool) (err error) {
	return instance.SetProperty("DebuggerMode", (value))
}

// GetDebuggerMode gets the value of DebuggerMode for the instance
func (instance *Msvm_SerialPortSettingData) GetPropertyDebuggerMode() (value bool, err error) {
	retValue, err := instance.GetProperty("DebuggerMode")
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
