// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SystemConfig_V3_Services struct
type SystemConfig_V3_Services struct {
	*SystemConfig_V3

	//
	DisplayName string

	//
	LoadOrderGroup string

	//
	ProcessId uint32

	//
	ProcessName string

	//
	ServiceName string

	//
	ServiceState uint32

	//
	SubProcessTag uint32

	//
	SvchostGroup string
}

func NewSystemConfig_V3_ServicesEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V3_Services, err error) {
	tmp, err := NewSystemConfig_V3Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V3_Services{
		SystemConfig_V3: tmp,
	}
	return
}

func NewSystemConfig_V3_ServicesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V3_Services, err error) {
	tmp, err := NewSystemConfig_V3Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V3_Services{
		SystemConfig_V3: tmp,
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *SystemConfig_V3_Services) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *SystemConfig_V3_Services) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
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

// SetLoadOrderGroup sets the value of LoadOrderGroup for the instance
func (instance *SystemConfig_V3_Services) SetPropertyLoadOrderGroup(value string) (err error) {
	return instance.SetProperty("LoadOrderGroup", (value))
}

// GetLoadOrderGroup gets the value of LoadOrderGroup for the instance
func (instance *SystemConfig_V3_Services) GetPropertyLoadOrderGroup() (value string, err error) {
	retValue, err := instance.GetProperty("LoadOrderGroup")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *SystemConfig_V3_Services) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *SystemConfig_V3_Services) GetPropertyProcessId() (value uint32, err error) {
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
func (instance *SystemConfig_V3_Services) SetPropertyProcessName(value string) (err error) {
	return instance.SetProperty("ProcessName", (value))
}

// GetProcessName gets the value of ProcessName for the instance
func (instance *SystemConfig_V3_Services) GetPropertyProcessName() (value string, err error) {
	retValue, err := instance.GetProperty("ProcessName")
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

// SetServiceName sets the value of ServiceName for the instance
func (instance *SystemConfig_V3_Services) SetPropertyServiceName(value string) (err error) {
	return instance.SetProperty("ServiceName", (value))
}

// GetServiceName gets the value of ServiceName for the instance
func (instance *SystemConfig_V3_Services) GetPropertyServiceName() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceName")
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

// SetServiceState sets the value of ServiceState for the instance
func (instance *SystemConfig_V3_Services) SetPropertyServiceState(value uint32) (err error) {
	return instance.SetProperty("ServiceState", (value))
}

// GetServiceState gets the value of ServiceState for the instance
func (instance *SystemConfig_V3_Services) GetPropertyServiceState() (value uint32, err error) {
	retValue, err := instance.GetProperty("ServiceState")
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

// SetSubProcessTag sets the value of SubProcessTag for the instance
func (instance *SystemConfig_V3_Services) SetPropertySubProcessTag(value uint32) (err error) {
	return instance.SetProperty("SubProcessTag", (value))
}

// GetSubProcessTag gets the value of SubProcessTag for the instance
func (instance *SystemConfig_V3_Services) GetPropertySubProcessTag() (value uint32, err error) {
	retValue, err := instance.GetProperty("SubProcessTag")
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

// SetSvchostGroup sets the value of SvchostGroup for the instance
func (instance *SystemConfig_V3_Services) SetPropertySvchostGroup(value string) (err error) {
	return instance.SetProperty("SvchostGroup", (value))
}

// GetSvchostGroup gets the value of SvchostGroup for the instance
func (instance *SystemConfig_V3_Services) GetPropertySvchostGroup() (value string, err error) {
	retValue, err := instance.GetProperty("SvchostGroup")
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
