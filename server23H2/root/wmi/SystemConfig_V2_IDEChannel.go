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

// SystemConfig_V2_IDEChannel struct
type SystemConfig_V2_IDEChannel struct {
	*SystemConfig_V2

	//
	DeviceTimingMode uint32

	//
	DeviceType uint32

	//
	LocationInformation string

	//
	LocationInformationLen uint32

	//
	TargetId uint32
}

func NewSystemConfig_V2_IDEChannelEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_IDEChannel, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_IDEChannel{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_IDEChannelEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_IDEChannel, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_IDEChannel{
		SystemConfig_V2: tmp,
	}
	return
}

// SetDeviceTimingMode sets the value of DeviceTimingMode for the instance
func (instance *SystemConfig_V2_IDEChannel) SetPropertyDeviceTimingMode(value uint32) (err error) {
	return instance.SetProperty("DeviceTimingMode", (value))
}

// GetDeviceTimingMode gets the value of DeviceTimingMode for the instance
func (instance *SystemConfig_V2_IDEChannel) GetPropertyDeviceTimingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceTimingMode")
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

// SetDeviceType sets the value of DeviceType for the instance
func (instance *SystemConfig_V2_IDEChannel) SetPropertyDeviceType(value uint32) (err error) {
	return instance.SetProperty("DeviceType", (value))
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *SystemConfig_V2_IDEChannel) GetPropertyDeviceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceType")
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

// SetLocationInformation sets the value of LocationInformation for the instance
func (instance *SystemConfig_V2_IDEChannel) SetPropertyLocationInformation(value string) (err error) {
	return instance.SetProperty("LocationInformation", (value))
}

// GetLocationInformation gets the value of LocationInformation for the instance
func (instance *SystemConfig_V2_IDEChannel) GetPropertyLocationInformation() (value string, err error) {
	retValue, err := instance.GetProperty("LocationInformation")
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

// SetLocationInformationLen sets the value of LocationInformationLen for the instance
func (instance *SystemConfig_V2_IDEChannel) SetPropertyLocationInformationLen(value uint32) (err error) {
	return instance.SetProperty("LocationInformationLen", (value))
}

// GetLocationInformationLen gets the value of LocationInformationLen for the instance
func (instance *SystemConfig_V2_IDEChannel) GetPropertyLocationInformationLen() (value uint32, err error) {
	retValue, err := instance.GetProperty("LocationInformationLen")
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

// SetTargetId sets the value of TargetId for the instance
func (instance *SystemConfig_V2_IDEChannel) SetPropertyTargetId(value uint32) (err error) {
	return instance.SetProperty("TargetId", (value))
}

// GetTargetId gets the value of TargetId for the instance
func (instance *SystemConfig_V2_IDEChannel) GetPropertyTargetId() (value uint32, err error) {
	retValue, err := instance.GetProperty("TargetId")
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
