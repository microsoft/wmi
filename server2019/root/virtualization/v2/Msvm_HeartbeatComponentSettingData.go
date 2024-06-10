// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_HeartbeatComponentSettingData struct
type Msvm_HeartbeatComponentSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	EnabledState uint16

	//
	ErrorThreshold uint32

	//
	Interval uint32

	//
	Latency uint32
}

func NewMsvm_HeartbeatComponentSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_HeartbeatComponentSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_HeartbeatComponentSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_HeartbeatComponentSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_HeartbeatComponentSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_HeartbeatComponentSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_HeartbeatComponentSettingData) SetPropertyEnabledState(value uint16) (err error) {
	return instance.SetProperty("EnabledState", (value))
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_HeartbeatComponentSettingData) GetPropertyEnabledState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnabledState")
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

// SetErrorThreshold sets the value of ErrorThreshold for the instance
func (instance *Msvm_HeartbeatComponentSettingData) SetPropertyErrorThreshold(value uint32) (err error) {
	return instance.SetProperty("ErrorThreshold", (value))
}

// GetErrorThreshold gets the value of ErrorThreshold for the instance
func (instance *Msvm_HeartbeatComponentSettingData) GetPropertyErrorThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorThreshold")
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

// SetInterval sets the value of Interval for the instance
func (instance *Msvm_HeartbeatComponentSettingData) SetPropertyInterval(value uint32) (err error) {
	return instance.SetProperty("Interval", (value))
}

// GetInterval gets the value of Interval for the instance
func (instance *Msvm_HeartbeatComponentSettingData) GetPropertyInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("Interval")
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

// SetLatency sets the value of Latency for the instance
func (instance *Msvm_HeartbeatComponentSettingData) SetPropertyLatency(value uint32) (err error) {
	return instance.SetProperty("Latency", (value))
}

// GetLatency gets the value of Latency for the instance
func (instance *Msvm_HeartbeatComponentSettingData) GetPropertyLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("Latency")
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
func (instance *Msvm_HeartbeatComponentSettingData) GetRelatedVirtualSystemSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemSettingData")
}

func (instance *Msvm_HeartbeatComponentSettingData) GetRelatedHeartbeatComponent() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_HeartbeatComponent")
}
