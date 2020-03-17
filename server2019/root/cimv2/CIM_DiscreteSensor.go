// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_DiscreteSensor struct
type CIM_DiscreteSensor struct {
	CIM_Sensor

	//
	AcceptableValues []string

	//
	CurrentReading string

	//
	PossibleValues []string
}

// SetAcceptableValues sets the value of AcceptableValues for the instance
func (instance *CIM_DiscreteSensor) SetPropertyAcceptableValues(value []string) (err error) {
	return instance.SetProperty("AcceptableValues", value)
}

// GetAcceptableValues gets the value of AcceptableValues for the instance
func (instance *CIM_DiscreteSensor) GetPropertyAcceptableValues() (value []string, err error) {
	retValue, err := instance.GetProperty("AcceptableValues")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentReading sets the value of CurrentReading for the instance
func (instance *CIM_DiscreteSensor) SetPropertyCurrentReading(value string) (err error) {
	return instance.SetProperty("CurrentReading", value)
}

// GetCurrentReading gets the value of CurrentReading for the instance
func (instance *CIM_DiscreteSensor) GetPropertyCurrentReading() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentReading")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPossibleValues sets the value of PossibleValues for the instance
func (instance *CIM_DiscreteSensor) SetPropertyPossibleValues(value []string) (err error) {
	return instance.SetProperty("PossibleValues", value)
}

// GetPossibleValues gets the value of PossibleValues for the instance
func (instance *CIM_DiscreteSensor) GetPropertyPossibleValues() (value []string, err error) {
	retValue, err := instance.GetProperty("PossibleValues")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
