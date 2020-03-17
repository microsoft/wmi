// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_SerialInterface struct
type CIM_SerialInterface struct {
	CIM_ControlledBy

	//
	FlowControlInfo uint16

	//
	NumberOfStopBits uint16

	//
	ParityInfo uint16
}

// SetFlowControlInfo sets the value of FlowControlInfo for the instance
func (instance *CIM_SerialInterface) SetPropertyFlowControlInfo(value uint16) (err error) {
	return instance.SetProperty("FlowControlInfo", value)
}

// GetFlowControlInfo gets the value of FlowControlInfo for the instance
func (instance *CIM_SerialInterface) GetPropertyFlowControlInfo() (value uint16, err error) {
	retValue, err := instance.GetProperty("FlowControlInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfStopBits sets the value of NumberOfStopBits for the instance
func (instance *CIM_SerialInterface) SetPropertyNumberOfStopBits(value uint16) (err error) {
	return instance.SetProperty("NumberOfStopBits", value)
}

// GetNumberOfStopBits gets the value of NumberOfStopBits for the instance
func (instance *CIM_SerialInterface) GetPropertyNumberOfStopBits() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfStopBits")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParityInfo sets the value of ParityInfo for the instance
func (instance *CIM_SerialInterface) SetPropertyParityInfo(value uint16) (err error) {
	return instance.SetProperty("ParityInfo", value)
}

// GetParityInfo gets the value of ParityInfo for the instance
func (instance *CIM_SerialInterface) GetPropertyParityInfo() (value uint16, err error) {
	retValue, err := instance.GetProperty("ParityInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
