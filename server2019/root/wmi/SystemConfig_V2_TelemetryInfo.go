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

// SystemConfig_V2_TelemetryInfo struct
type SystemConfig_V2_TelemetryInfo struct {
	*SystemConfig_V2

	//
	MachineId interface{}
}

func NewSystemConfig_V2_TelemetryInfoEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_TelemetryInfo, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_TelemetryInfo{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_TelemetryInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_TelemetryInfo, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_TelemetryInfo{
		SystemConfig_V2: tmp,
	}
	return
}

// SetMachineId sets the value of MachineId for the instance
func (instance *SystemConfig_V2_TelemetryInfo) SetPropertyMachineId(value interface{}) (err error) {
	return instance.SetProperty("MachineId", (value))
}

// GetMachineId gets the value of MachineId for the instance
func (instance *SystemConfig_V2_TelemetryInfo) GetPropertyMachineId() (value interface{}, err error) {
	retValue, err := instance.GetProperty("MachineId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
