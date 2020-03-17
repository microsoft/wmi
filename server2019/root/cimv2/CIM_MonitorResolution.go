// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_MonitorResolution struct
type CIM_MonitorResolution struct {
	CIM_Setting

	//
	HorizontalResolution uint32

	//
	MaxRefreshRate uint32

	//
	MinRefreshRate uint32

	//
	RefreshRate uint32

	//
	ScanMode uint16

	//
	VerticalResolution uint32
}

// SetHorizontalResolution sets the value of HorizontalResolution for the instance
func (instance *CIM_MonitorResolution) SetPropertyHorizontalResolution(value uint32) (err error) {
	return instance.SetProperty("HorizontalResolution", value)
}

// GetHorizontalResolution gets the value of HorizontalResolution for the instance
func (instance *CIM_MonitorResolution) GetPropertyHorizontalResolution() (value uint32, err error) {
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

// SetMaxRefreshRate sets the value of MaxRefreshRate for the instance
func (instance *CIM_MonitorResolution) SetPropertyMaxRefreshRate(value uint32) (err error) {
	return instance.SetProperty("MaxRefreshRate", value)
}

// GetMaxRefreshRate gets the value of MaxRefreshRate for the instance
func (instance *CIM_MonitorResolution) GetPropertyMaxRefreshRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxRefreshRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinRefreshRate sets the value of MinRefreshRate for the instance
func (instance *CIM_MonitorResolution) SetPropertyMinRefreshRate(value uint32) (err error) {
	return instance.SetProperty("MinRefreshRate", value)
}

// GetMinRefreshRate gets the value of MinRefreshRate for the instance
func (instance *CIM_MonitorResolution) GetPropertyMinRefreshRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinRefreshRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRefreshRate sets the value of RefreshRate for the instance
func (instance *CIM_MonitorResolution) SetPropertyRefreshRate(value uint32) (err error) {
	return instance.SetProperty("RefreshRate", value)
}

// GetRefreshRate gets the value of RefreshRate for the instance
func (instance *CIM_MonitorResolution) GetPropertyRefreshRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("RefreshRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScanMode sets the value of ScanMode for the instance
func (instance *CIM_MonitorResolution) SetPropertyScanMode(value uint16) (err error) {
	return instance.SetProperty("ScanMode", value)
}

// GetScanMode gets the value of ScanMode for the instance
func (instance *CIM_MonitorResolution) GetPropertyScanMode() (value uint16, err error) {
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

// SetVerticalResolution sets the value of VerticalResolution for the instance
func (instance *CIM_MonitorResolution) SetPropertyVerticalResolution(value uint32) (err error) {
	return instance.SetProperty("VerticalResolution", value)
}

// GetVerticalResolution gets the value of VerticalResolution for the instance
func (instance *CIM_MonitorResolution) GetPropertyVerticalResolution() (value uint32, err error) {
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
