// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_BinarySensor struct
type CIM_BinarySensor struct {
	CIM_Sensor

	//
	CurrentReading bool

	//
	ExpectedReading bool

	//
	InterpretationOfFalse string

	//
	InterpretationOfTrue string
}

// SetCurrentReading sets the value of CurrentReading for the instance
func (instance *CIM_BinarySensor) SetPropertyCurrentReading(value bool) (err error) {
	return instance.SetProperty("CurrentReading", value)
}

// GetCurrentReading gets the value of CurrentReading for the instance
func (instance *CIM_BinarySensor) GetPropertyCurrentReading() (value bool, err error) {
	retValue, err := instance.GetProperty("CurrentReading")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpectedReading sets the value of ExpectedReading for the instance
func (instance *CIM_BinarySensor) SetPropertyExpectedReading(value bool) (err error) {
	return instance.SetProperty("ExpectedReading", value)
}

// GetExpectedReading gets the value of ExpectedReading for the instance
func (instance *CIM_BinarySensor) GetPropertyExpectedReading() (value bool, err error) {
	retValue, err := instance.GetProperty("ExpectedReading")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterpretationOfFalse sets the value of InterpretationOfFalse for the instance
func (instance *CIM_BinarySensor) SetPropertyInterpretationOfFalse(value string) (err error) {
	return instance.SetProperty("InterpretationOfFalse", value)
}

// GetInterpretationOfFalse gets the value of InterpretationOfFalse for the instance
func (instance *CIM_BinarySensor) GetPropertyInterpretationOfFalse() (value string, err error) {
	retValue, err := instance.GetProperty("InterpretationOfFalse")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterpretationOfTrue sets the value of InterpretationOfTrue for the instance
func (instance *CIM_BinarySensor) SetPropertyInterpretationOfTrue(value string) (err error) {
	return instance.SetProperty("InterpretationOfTrue", value)
}

// GetInterpretationOfTrue gets the value of InterpretationOfTrue for the instance
func (instance *CIM_BinarySensor) GetPropertyInterpretationOfTrue() (value string, err error) {
	retValue, err := instance.GetProperty("InterpretationOfTrue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
