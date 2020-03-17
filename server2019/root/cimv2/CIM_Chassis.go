// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_Chassis struct
type CIM_Chassis struct {
	CIM_PhysicalFrame

	//
	ChassisTypes []uint16

	//
	CurrentRequiredOrProduced int16

	//
	HeatGeneration uint16

	//
	NumberOfPowerCords uint16

	//
	TypeDescriptions []string
}

// SetChassisTypes sets the value of ChassisTypes for the instance
func (instance *CIM_Chassis) SetPropertyChassisTypes(value []uint16) (err error) {
	return instance.SetProperty("ChassisTypes", value)
}

// GetChassisTypes gets the value of ChassisTypes for the instance
func (instance *CIM_Chassis) GetPropertyChassisTypes() (value []uint16, err error) {
	retValue, err := instance.GetProperty("ChassisTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentRequiredOrProduced sets the value of CurrentRequiredOrProduced for the instance
func (instance *CIM_Chassis) SetPropertyCurrentRequiredOrProduced(value int16) (err error) {
	return instance.SetProperty("CurrentRequiredOrProduced", value)
}

// GetCurrentRequiredOrProduced gets the value of CurrentRequiredOrProduced for the instance
func (instance *CIM_Chassis) GetPropertyCurrentRequiredOrProduced() (value int16, err error) {
	retValue, err := instance.GetProperty("CurrentRequiredOrProduced")
	if err != nil {
		return
	}
	value, ok := retValue.(int16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHeatGeneration sets the value of HeatGeneration for the instance
func (instance *CIM_Chassis) SetPropertyHeatGeneration(value uint16) (err error) {
	return instance.SetProperty("HeatGeneration", value)
}

// GetHeatGeneration gets the value of HeatGeneration for the instance
func (instance *CIM_Chassis) GetPropertyHeatGeneration() (value uint16, err error) {
	retValue, err := instance.GetProperty("HeatGeneration")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfPowerCords sets the value of NumberOfPowerCords for the instance
func (instance *CIM_Chassis) SetPropertyNumberOfPowerCords(value uint16) (err error) {
	return instance.SetProperty("NumberOfPowerCords", value)
}

// GetNumberOfPowerCords gets the value of NumberOfPowerCords for the instance
func (instance *CIM_Chassis) GetPropertyNumberOfPowerCords() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfPowerCords")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTypeDescriptions sets the value of TypeDescriptions for the instance
func (instance *CIM_Chassis) SetPropertyTypeDescriptions(value []string) (err error) {
	return instance.SetProperty("TypeDescriptions", value)
}

// GetTypeDescriptions gets the value of TypeDescriptions for the instance
func (instance *CIM_Chassis) GetPropertyTypeDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("TypeDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
