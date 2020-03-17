// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_PointingDevice struct
type CIM_PointingDevice struct {
	CIM_UserDevice

	//
	Handedness uint16

	//
	NumberOfButtons uint8

	//
	PointingType uint16

	//
	Resolution uint32
}

// SetHandedness sets the value of Handedness for the instance
func (instance *CIM_PointingDevice) SetPropertyHandedness(value uint16) (err error) {
	return instance.SetProperty("Handedness", value)
}

// GetHandedness gets the value of Handedness for the instance
func (instance *CIM_PointingDevice) GetPropertyHandedness() (value uint16, err error) {
	retValue, err := instance.GetProperty("Handedness")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfButtons sets the value of NumberOfButtons for the instance
func (instance *CIM_PointingDevice) SetPropertyNumberOfButtons(value uint8) (err error) {
	return instance.SetProperty("NumberOfButtons", value)
}

// GetNumberOfButtons gets the value of NumberOfButtons for the instance
func (instance *CIM_PointingDevice) GetPropertyNumberOfButtons() (value uint8, err error) {
	retValue, err := instance.GetProperty("NumberOfButtons")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPointingType sets the value of PointingType for the instance
func (instance *CIM_PointingDevice) SetPropertyPointingType(value uint16) (err error) {
	return instance.SetProperty("PointingType", value)
}

// GetPointingType gets the value of PointingType for the instance
func (instance *CIM_PointingDevice) GetPropertyPointingType() (value uint16, err error) {
	retValue, err := instance.GetProperty("PointingType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResolution sets the value of Resolution for the instance
func (instance *CIM_PointingDevice) SetPropertyResolution(value uint32) (err error) {
	return instance.SetProperty("Resolution", value)
}

// GetResolution gets the value of Resolution for the instance
func (instance *CIM_PointingDevice) GetPropertyResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("Resolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
