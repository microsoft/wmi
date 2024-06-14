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

// SystemConfig_V0_Video struct
type SystemConfig_V0_Video struct {
	*SystemConfig_V0

	//
	AdapterString []byte

	//
	BiosString []byte

	//
	BitsPerPixel uint32

	//
	ChipType []byte

	//
	DACType []byte

	//
	DeviceId []byte

	//
	MemorySize uint32

	//
	StateFlags uint32

	//
	VRefresh uint32

	//
	XResolution uint32

	//
	YResolution uint32
}

func NewSystemConfig_V0_VideoEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V0_Video, err error) {
	tmp, err := NewSystemConfig_V0Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V0_Video{
		SystemConfig_V0: tmp,
	}
	return
}

func NewSystemConfig_V0_VideoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V0_Video, err error) {
	tmp, err := NewSystemConfig_V0Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V0_Video{
		SystemConfig_V0: tmp,
	}
	return
}

// SetAdapterString sets the value of AdapterString for the instance
func (instance *SystemConfig_V0_Video) SetPropertyAdapterString(value []byte) (err error) {
	return instance.SetProperty("AdapterString", (value))
}

// GetAdapterString gets the value of AdapterString for the instance
func (instance *SystemConfig_V0_Video) GetPropertyAdapterString() (value []byte, err error) {
	retValue, err := instance.GetProperty("AdapterString")
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

// SetBiosString sets the value of BiosString for the instance
func (instance *SystemConfig_V0_Video) SetPropertyBiosString(value []byte) (err error) {
	return instance.SetProperty("BiosString", (value))
}

// GetBiosString gets the value of BiosString for the instance
func (instance *SystemConfig_V0_Video) GetPropertyBiosString() (value []byte, err error) {
	retValue, err := instance.GetProperty("BiosString")
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

// SetBitsPerPixel sets the value of BitsPerPixel for the instance
func (instance *SystemConfig_V0_Video) SetPropertyBitsPerPixel(value uint32) (err error) {
	return instance.SetProperty("BitsPerPixel", (value))
}

// GetBitsPerPixel gets the value of BitsPerPixel for the instance
func (instance *SystemConfig_V0_Video) GetPropertyBitsPerPixel() (value uint32, err error) {
	retValue, err := instance.GetProperty("BitsPerPixel")
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

// SetChipType sets the value of ChipType for the instance
func (instance *SystemConfig_V0_Video) SetPropertyChipType(value []byte) (err error) {
	return instance.SetProperty("ChipType", (value))
}

// GetChipType gets the value of ChipType for the instance
func (instance *SystemConfig_V0_Video) GetPropertyChipType() (value []byte, err error) {
	retValue, err := instance.GetProperty("ChipType")
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

// SetDACType sets the value of DACType for the instance
func (instance *SystemConfig_V0_Video) SetPropertyDACType(value []byte) (err error) {
	return instance.SetProperty("DACType", (value))
}

// GetDACType gets the value of DACType for the instance
func (instance *SystemConfig_V0_Video) GetPropertyDACType() (value []byte, err error) {
	retValue, err := instance.GetProperty("DACType")
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

// SetDeviceId sets the value of DeviceId for the instance
func (instance *SystemConfig_V0_Video) SetPropertyDeviceId(value []byte) (err error) {
	return instance.SetProperty("DeviceId", (value))
}

// GetDeviceId gets the value of DeviceId for the instance
func (instance *SystemConfig_V0_Video) GetPropertyDeviceId() (value []byte, err error) {
	retValue, err := instance.GetProperty("DeviceId")
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

// SetMemorySize sets the value of MemorySize for the instance
func (instance *SystemConfig_V0_Video) SetPropertyMemorySize(value uint32) (err error) {
	return instance.SetProperty("MemorySize", (value))
}

// GetMemorySize gets the value of MemorySize for the instance
func (instance *SystemConfig_V0_Video) GetPropertyMemorySize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MemorySize")
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

// SetStateFlags sets the value of StateFlags for the instance
func (instance *SystemConfig_V0_Video) SetPropertyStateFlags(value uint32) (err error) {
	return instance.SetProperty("StateFlags", (value))
}

// GetStateFlags gets the value of StateFlags for the instance
func (instance *SystemConfig_V0_Video) GetPropertyStateFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("StateFlags")
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

// SetVRefresh sets the value of VRefresh for the instance
func (instance *SystemConfig_V0_Video) SetPropertyVRefresh(value uint32) (err error) {
	return instance.SetProperty("VRefresh", (value))
}

// GetVRefresh gets the value of VRefresh for the instance
func (instance *SystemConfig_V0_Video) GetPropertyVRefresh() (value uint32, err error) {
	retValue, err := instance.GetProperty("VRefresh")
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

// SetXResolution sets the value of XResolution for the instance
func (instance *SystemConfig_V0_Video) SetPropertyXResolution(value uint32) (err error) {
	return instance.SetProperty("XResolution", (value))
}

// GetXResolution gets the value of XResolution for the instance
func (instance *SystemConfig_V0_Video) GetPropertyXResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("XResolution")
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

// SetYResolution sets the value of YResolution for the instance
func (instance *SystemConfig_V0_Video) SetPropertyYResolution(value uint32) (err error) {
	return instance.SetProperty("YResolution", (value))
}

// GetYResolution gets the value of YResolution for the instance
func (instance *SystemConfig_V0_Video) GetPropertyYResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("YResolution")
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
