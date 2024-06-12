// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_FlexIoDeviceSettingData struct
type Msvm_FlexIoDeviceSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	EmulatorConfiguration []string

	//
	EmulatorId string

	//
	PhysicalNumaNode uint16

	//
	TargetVtl uint8

	//
	VirtualSystemIdentifiers []string
}

func NewMsvm_FlexIoDeviceSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_FlexIoDeviceSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_FlexIoDeviceSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_FlexIoDeviceSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_FlexIoDeviceSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_FlexIoDeviceSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetEmulatorConfiguration sets the value of EmulatorConfiguration for the instance
func (instance *Msvm_FlexIoDeviceSettingData) SetPropertyEmulatorConfiguration(value []string) (err error) {
	return instance.SetProperty("EmulatorConfiguration", (value))
}

// GetEmulatorConfiguration gets the value of EmulatorConfiguration for the instance
func (instance *Msvm_FlexIoDeviceSettingData) GetPropertyEmulatorConfiguration() (value []string, err error) {
	retValue, err := instance.GetProperty("EmulatorConfiguration")
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

// SetEmulatorId sets the value of EmulatorId for the instance
func (instance *Msvm_FlexIoDeviceSettingData) SetPropertyEmulatorId(value string) (err error) {
	return instance.SetProperty("EmulatorId", (value))
}

// GetEmulatorId gets the value of EmulatorId for the instance
func (instance *Msvm_FlexIoDeviceSettingData) GetPropertyEmulatorId() (value string, err error) {
	retValue, err := instance.GetProperty("EmulatorId")
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

// SetPhysicalNumaNode sets the value of PhysicalNumaNode for the instance
func (instance *Msvm_FlexIoDeviceSettingData) SetPropertyPhysicalNumaNode(value uint16) (err error) {
	return instance.SetProperty("PhysicalNumaNode", (value))
}

// GetPhysicalNumaNode gets the value of PhysicalNumaNode for the instance
func (instance *Msvm_FlexIoDeviceSettingData) GetPropertyPhysicalNumaNode() (value uint16, err error) {
	retValue, err := instance.GetProperty("PhysicalNumaNode")
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

// SetTargetVtl sets the value of TargetVtl for the instance
func (instance *Msvm_FlexIoDeviceSettingData) SetPropertyTargetVtl(value uint8) (err error) {
	return instance.SetProperty("TargetVtl", (value))
}

// GetTargetVtl gets the value of TargetVtl for the instance
func (instance *Msvm_FlexIoDeviceSettingData) GetPropertyTargetVtl() (value uint8, err error) {
	retValue, err := instance.GetProperty("TargetVtl")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetVirtualSystemIdentifiers sets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_FlexIoDeviceSettingData) SetPropertyVirtualSystemIdentifiers(value []string) (err error) {
	return instance.SetProperty("VirtualSystemIdentifiers", (value))
}

// GetVirtualSystemIdentifiers gets the value of VirtualSystemIdentifiers for the instance
func (instance *Msvm_FlexIoDeviceSettingData) GetPropertyVirtualSystemIdentifiers() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemIdentifiers")
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
func (instance *Msvm_FlexIoDeviceSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
