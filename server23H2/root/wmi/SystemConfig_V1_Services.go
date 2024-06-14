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

// SystemConfig_V1_Services struct
type SystemConfig_V1_Services struct {
	*SystemConfig_V1

	//
	DisplayName []byte

	//
	ProcessId uint32

	//
	ProcessName []byte

	//
	ServiceName []byte
}

func NewSystemConfig_V1_ServicesEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V1_Services, err error) {
	tmp, err := NewSystemConfig_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V1_Services{
		SystemConfig_V1: tmp,
	}
	return
}

func NewSystemConfig_V1_ServicesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V1_Services, err error) {
	tmp, err := NewSystemConfig_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V1_Services{
		SystemConfig_V1: tmp,
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *SystemConfig_V1_Services) SetPropertyDisplayName(value []byte) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *SystemConfig_V1_Services) GetPropertyDisplayName() (value []byte, err error) {
	retValue, err := instance.GetProperty("DisplayName")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *SystemConfig_V1_Services) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *SystemConfig_V1_Services) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetProcessName sets the value of ProcessName for the instance
func (instance *SystemConfig_V1_Services) SetPropertyProcessName(value []byte) (err error) {
	return instance.SetProperty("ProcessName", (value))
}

// GetProcessName gets the value of ProcessName for the instance
func (instance *SystemConfig_V1_Services) GetPropertyProcessName() (value []byte, err error) {
	retValue, err := instance.GetProperty("ProcessName")
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

// SetServiceName sets the value of ServiceName for the instance
func (instance *SystemConfig_V1_Services) SetPropertyServiceName(value []byte) (err error) {
	return instance.SetProperty("ServiceName", (value))
}

// GetServiceName gets the value of ServiceName for the instance
func (instance *SystemConfig_V1_Services) GetPropertyServiceName() (value []byte, err error) {
	retValue, err := instance.GetProperty("ServiceName")
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
