// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("CurrentState", value)
}

// GetCurrentState gets the value of CurrentState for the instance
func (instance *CIM_Sensor) GetPropertyCurrentState() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentState")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherSensorTypeDescription sets the value of OtherSensorTypeDescription for the instance
func (instance *CIM_Sensor) SetPropertyOtherSensorTypeDescription(value string) (err error) {
	return instance.SetProperty("OtherSensorTypeDescription", value)
}

// GetOtherSensorTypeDescription gets the value of OtherSensorTypeDescription for the instance
func (instance *CIM_Sensor) GetPropertyOtherSensorTypeDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherSensorTypeDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPollingInterval sets the value of PollingInterval for the instance
func (instance *CIM_Sensor) SetPropertyPollingInterval(value uint64) (err error) {
	return instance.SetProperty("PollingInterval", value)
}

// GetPollingInterval gets the value of PollingInterval for the instance
func (instance *CIM_Sensor) GetPropertyPollingInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("PollingInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPossibleStates sets the value of PossibleStates for the instance
func (instance *CIM_Sensor) SetPropertyPossibleStates(value []string) (err error) {
	return instance.SetProperty("PossibleStates", value)
}

// GetPossibleStates gets the value of PossibleStates for the instance
func (instance *CIM_Sensor) GetPropertyPossibleStates() (value []string, err error) {
	retValue, err := instance.GetProperty("PossibleStates")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSensorType sets the value of SensorType for the instance
func (instance *CIM_Sensor) SetPropertySensorType(value uint16) (err error) {
	return instance.SetProperty("SensorType", value)
}

// GetSensorType gets the value of SensorType for the instance
func (instance *CIM_Sensor) GetPropertySensorType() (value uint16, err error) {
	retValue, err := instance.GetProperty("SensorType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
