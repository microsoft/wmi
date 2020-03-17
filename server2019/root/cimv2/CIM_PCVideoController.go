// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_PCVideoController struct
type CIM_PCVideoController struct {
	CIM_VideoController

	//
	NumberOfColorPlanes uint16

	//
	VideoArchitecture uint16

	//
	VideoMode uint16
}

// SetNumberOfColorPlanes sets the value of NumberOfColorPlanes for the instance
func (instance *CIM_PCVideoController) SetPropertyNumberOfColorPlanes(value uint16) (err error) {
	return instance.SetProperty("NumberOfColorPlanes", value)
}

// GetNumberOfColorPlanes gets the value of NumberOfColorPlanes for the instance
func (instance *CIM_PCVideoController) GetPropertyNumberOfColorPlanes() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfColorPlanes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVideoArchitecture sets the value of VideoArchitecture for the instance
func (instance *CIM_PCVideoController) SetPropertyVideoArchitecture(value uint16) (err error) {
	return instance.SetProperty("VideoArchitecture", value)
}

// GetVideoArchitecture gets the value of VideoArchitecture for the instance
func (instance *CIM_PCVideoController) GetPropertyVideoArchitecture() (value uint16, err error) {
	retValue, err := instance.GetProperty("VideoArchitecture")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVideoMode sets the value of VideoMode for the instance
func (instance *CIM_PCVideoController) SetPropertyVideoMode(value uint16) (err error) {
	return instance.SetProperty("VideoMode", value)
}

// GetVideoMode gets the value of VideoMode for the instance
func (instance *CIM_PCVideoController) GetPropertyVideoMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("VideoMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
