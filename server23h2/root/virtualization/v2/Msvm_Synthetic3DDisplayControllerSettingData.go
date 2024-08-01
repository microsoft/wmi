// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_Synthetic3DDisplayControllerSettingData struct
type Msvm_Synthetic3DDisplayControllerSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	MaximumMonitors uint8

	//
	MaximumScreenResolution uint8

	// The video memory size for the Virtual Machine
	VRAMSizeBytes uint64
}

func NewMsvm_Synthetic3DDisplayControllerSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_Synthetic3DDisplayControllerSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_Synthetic3DDisplayControllerSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_Synthetic3DDisplayControllerSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_Synthetic3DDisplayControllerSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_Synthetic3DDisplayControllerSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetMaximumMonitors sets the value of MaximumMonitors for the instance
func (instance *Msvm_Synthetic3DDisplayControllerSettingData) SetPropertyMaximumMonitors(value uint8) (err error) {
	return instance.SetProperty("MaximumMonitors", (value))
}

// GetMaximumMonitors gets the value of MaximumMonitors for the instance
func (instance *Msvm_Synthetic3DDisplayControllerSettingData) GetPropertyMaximumMonitors() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaximumMonitors")
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

// SetMaximumScreenResolution sets the value of MaximumScreenResolution for the instance
func (instance *Msvm_Synthetic3DDisplayControllerSettingData) SetPropertyMaximumScreenResolution(value uint8) (err error) {
	return instance.SetProperty("MaximumScreenResolution", (value))
}

// GetMaximumScreenResolution gets the value of MaximumScreenResolution for the instance
func (instance *Msvm_Synthetic3DDisplayControllerSettingData) GetPropertyMaximumScreenResolution() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaximumScreenResolution")
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

// SetVRAMSizeBytes sets the value of VRAMSizeBytes for the instance
func (instance *Msvm_Synthetic3DDisplayControllerSettingData) SetPropertyVRAMSizeBytes(value uint64) (err error) {
	return instance.SetProperty("VRAMSizeBytes", (value))
}

// GetVRAMSizeBytes gets the value of VRAMSizeBytes for the instance
func (instance *Msvm_Synthetic3DDisplayControllerSettingData) GetPropertyVRAMSizeBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRAMSizeBytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
func (instance *Msvm_Synthetic3DDisplayControllerSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
