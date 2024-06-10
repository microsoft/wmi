// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Hardware
//
// ////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Sensor struct
type CIM_Sensor struct {
	*CIM_LogicalDevice

	//
	CurrentState string

	//
	OtherSensorTypeDescription string

	//
	PollingInterval uint64

	//
	PossibleStates []string

	//
	SensorType uint16
}

func NewCIM_SensorEx1(instance *cim.WmiInstance) (newInstance *CIM_Sensor, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Sensor{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_SensorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Sensor, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Sensor{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetCurrentState sets the value of CurrentState for the instance
func (instance *CIM_Sensor) SetPropertyCurrentState(value string) (err error) {
	return instance.SetProperty("CurrentState", (value))
}

// GetCurrentState gets the value of CurrentState for the instance
func (instance *CIM_Sensor) GetPropertyCurrentState() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentState")
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

// SetOtherSensorTypeDescription sets the value of OtherSensorTypeDescription for the instance
func (instance *CIM_Sensor) SetPropertyOtherSensorTypeDescription(value string) (err error) {
	return instance.SetProperty("OtherSensorTypeDescription", (value))
}

// GetOtherSensorTypeDescription gets the value of OtherSensorTypeDescription for the instance
func (instance *CIM_Sensor) GetPropertyOtherSensorTypeDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherSensorTypeDescription")
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

// SetPollingInterval sets the value of PollingInterval for the instance
func (instance *CIM_Sensor) SetPropertyPollingInterval(value uint64) (err error) {
	return instance.SetProperty("PollingInterval", (value))
}

// GetPollingInterval gets the value of PollingInterval for the instance
func (instance *CIM_Sensor) GetPropertyPollingInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("PollingInterval")
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

// SetPossibleStates sets the value of PossibleStates for the instance
func (instance *CIM_Sensor) SetPropertyPossibleStates(value []string) (err error) {
	return instance.SetProperty("PossibleStates", (value))
}

// GetPossibleStates gets the value of PossibleStates for the instance
func (instance *CIM_Sensor) GetPropertyPossibleStates() (value []string, err error) {
	retValue, err := instance.GetProperty("PossibleStates")
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

// SetSensorType sets the value of SensorType for the instance
func (instance *CIM_Sensor) SetPropertySensorType(value uint16) (err error) {
	return instance.SetProperty("SensorType", (value))
}

// GetSensorType gets the value of SensorType for the instance
func (instance *CIM_Sensor) GetPropertySensorType() (value uint16, err error) {
	retValue, err := instance.GetProperty("SensorType")
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
