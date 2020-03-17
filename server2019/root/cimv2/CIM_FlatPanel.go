// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_FlatPanel struct
type CIM_FlatPanel struct {
	CIM_Display

	//
	DisplayType uint16

	//
	HorizontalResolution uint32

	//
	LightSource uint16

	//
	ScanMode uint16

	//
	SupportsColor bool

	//
	VerticalResolution uint32
}

// SetDisplayType sets the value of DisplayType for the instance
func (instance *CIM_FlatPanel) SetPropertyDisplayType(value uint16) (err error) {
	return instance.SetProperty("DisplayType", value)
}

// GetDisplayType gets the value of DisplayType for the instance
func (instance *CIM_FlatPanel) GetPropertyDisplayType() (value uint16, err error) {
	retValue, err := instance.GetProperty("DisplayType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHorizontalResolution sets the value of HorizontalResolution for the instance
func (instance *CIM_FlatPanel) SetPropertyHorizontalResolution(value uint32) (err error) {
	return instance.SetProperty("HorizontalResolution", value)
}

// GetHorizontalResolution gets the value of HorizontalResolution for the instance
func (instance *CIM_FlatPanel) GetPropertyHorizontalResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("HorizontalResolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLightSource sets the value of LightSource for the instance
func (instance *CIM_FlatPanel) SetPropertyLightSource(value uint16) (err error) {
	return instance.SetProperty("LightSource", value)
}

// GetLightSource gets the value of LightSource for the instance
func (instance *CIM_FlatPanel) GetPropertyLightSource() (value uint16, err error) {
	retValue, err := instance.GetProperty("LightSource")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScanMode sets the value of ScanMode for the instance
func (instance *CIM_FlatPanel) SetPropertyScanMode(value uint16) (err error) {
	return instance.SetProperty("ScanMode", value)
}

// GetScanMode gets the value of ScanMode for the instance
func (instance *CIM_FlatPanel) GetPropertyScanMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ScanMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsColor sets the value of SupportsColor for the instance
func (instance *CIM_FlatPanel) SetPropertySupportsColor(value bool) (err error) {
	return instance.SetProperty("SupportsColor", value)
}

// GetSupportsColor gets the value of SupportsColor for the instance
func (instance *CIM_FlatPanel) GetPropertySupportsColor() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsColor")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVerticalResolution sets the value of VerticalResolution for the instance
func (instance *CIM_FlatPanel) SetPropertyVerticalResolution(value uint32) (err error) {
	return instance.SetProperty("VerticalResolution", value)
}

// GetVerticalResolution gets the value of VerticalResolution for the instance
func (instance *CIM_FlatPanel) GetPropertyVerticalResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("VerticalResolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
