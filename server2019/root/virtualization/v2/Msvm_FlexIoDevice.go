// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_FlexIoDevice struct
type Msvm_FlexIoDevice struct {
	*CIM_LogicalDevice

	//
	EmulatorConfiguration []string

	//
	EmulatorId string
}

func NewMsvm_FlexIoDeviceEx1(instance *cim.WmiInstance) (newInstance *Msvm_FlexIoDevice, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_FlexIoDevice{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewMsvm_FlexIoDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_FlexIoDevice, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_FlexIoDevice{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetEmulatorConfiguration sets the value of EmulatorConfiguration for the instance
func (instance *Msvm_FlexIoDevice) SetPropertyEmulatorConfiguration(value []string) (err error) {
	return instance.SetProperty("EmulatorConfiguration", value)
}

// GetEmulatorConfiguration gets the value of EmulatorConfiguration for the instance
func (instance *Msvm_FlexIoDevice) GetPropertyEmulatorConfiguration() (value []string, err error) {
	retValue, err := instance.GetProperty("EmulatorConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEmulatorId sets the value of EmulatorId for the instance
func (instance *Msvm_FlexIoDevice) SetPropertyEmulatorId(value string) (err error) {
	return instance.SetProperty("EmulatorId", value)
}

// GetEmulatorId gets the value of EmulatorId for the instance
func (instance *Msvm_FlexIoDevice) GetPropertyEmulatorId() (value string, err error) {
	retValue, err := instance.GetProperty("EmulatorId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
