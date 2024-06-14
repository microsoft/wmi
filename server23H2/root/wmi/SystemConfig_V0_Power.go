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

// SystemConfig_V0_Power struct
type SystemConfig_V0_Power struct {
	*SystemConfig_V0

	//
	Pad1 uint8

	//
	Pad2 uint8

	//
	Pad3 uint8

	//
	S1 uint8

	//
	S2 uint8

	//
	S3 uint8

	//
	S4 uint8

	//
	S5 uint8
}

func NewSystemConfig_V0_PowerEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V0_Power, err error) {
	tmp, err := NewSystemConfig_V0Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V0_Power{
		SystemConfig_V0: tmp,
	}
	return
}

func NewSystemConfig_V0_PowerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V0_Power, err error) {
	tmp, err := NewSystemConfig_V0Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V0_Power{
		SystemConfig_V0: tmp,
	}
	return
}

// SetPad1 sets the value of Pad1 for the instance
func (instance *SystemConfig_V0_Power) SetPropertyPad1(value uint8) (err error) {
	return instance.SetProperty("Pad1", (value))
}

// GetPad1 gets the value of Pad1 for the instance
func (instance *SystemConfig_V0_Power) GetPropertyPad1() (value uint8, err error) {
	retValue, err := instance.GetProperty("Pad1")
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

// SetPad2 sets the value of Pad2 for the instance
func (instance *SystemConfig_V0_Power) SetPropertyPad2(value uint8) (err error) {
	return instance.SetProperty("Pad2", (value))
}

// GetPad2 gets the value of Pad2 for the instance
func (instance *SystemConfig_V0_Power) GetPropertyPad2() (value uint8, err error) {
	retValue, err := instance.GetProperty("Pad2")
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

// SetPad3 sets the value of Pad3 for the instance
func (instance *SystemConfig_V0_Power) SetPropertyPad3(value uint8) (err error) {
	return instance.SetProperty("Pad3", (value))
}

// GetPad3 gets the value of Pad3 for the instance
func (instance *SystemConfig_V0_Power) GetPropertyPad3() (value uint8, err error) {
	retValue, err := instance.GetProperty("Pad3")
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

// SetS1 sets the value of S1 for the instance
func (instance *SystemConfig_V0_Power) SetPropertyS1(value uint8) (err error) {
	return instance.SetProperty("S1", (value))
}

// GetS1 gets the value of S1 for the instance
func (instance *SystemConfig_V0_Power) GetPropertyS1() (value uint8, err error) {
	retValue, err := instance.GetProperty("S1")
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

// SetS2 sets the value of S2 for the instance
func (instance *SystemConfig_V0_Power) SetPropertyS2(value uint8) (err error) {
	return instance.SetProperty("S2", (value))
}

// GetS2 gets the value of S2 for the instance
func (instance *SystemConfig_V0_Power) GetPropertyS2() (value uint8, err error) {
	retValue, err := instance.GetProperty("S2")
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

// SetS3 sets the value of S3 for the instance
func (instance *SystemConfig_V0_Power) SetPropertyS3(value uint8) (err error) {
	return instance.SetProperty("S3", (value))
}

// GetS3 gets the value of S3 for the instance
func (instance *SystemConfig_V0_Power) GetPropertyS3() (value uint8, err error) {
	retValue, err := instance.GetProperty("S3")
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

// SetS4 sets the value of S4 for the instance
func (instance *SystemConfig_V0_Power) SetPropertyS4(value uint8) (err error) {
	return instance.SetProperty("S4", (value))
}

// GetS4 gets the value of S4 for the instance
func (instance *SystemConfig_V0_Power) GetPropertyS4() (value uint8, err error) {
	retValue, err := instance.GetProperty("S4")
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

// SetS5 sets the value of S5 for the instance
func (instance *SystemConfig_V0_Power) SetPropertyS5(value uint8) (err error) {
	return instance.SetProperty("S5", (value))
}

// GetS5 gets the value of S5 for the instance
func (instance *SystemConfig_V0_Power) GetPropertyS5() (value uint8, err error) {
	retValue, err := instance.GetProperty("S5")
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
