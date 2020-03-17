// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DesktopMonitor struct
type Win32_DesktopMonitor struct {
	CIM_DesktopMonitor

	//
	MonitorManufacturer string

	//
	MonitorType string

	//
	PixelsPerXLogicalInch uint32

	//
	PixelsPerYLogicalInch uint32
}

// SetMonitorManufacturer sets the value of MonitorManufacturer for the instance
func (instance *Win32_DesktopMonitor) SetPropertyMonitorManufacturer(value string) (err error) {
	return instance.SetProperty("MonitorManufacturer", value)
}

// GetMonitorManufacturer gets the value of MonitorManufacturer for the instance
func (instance *Win32_DesktopMonitor) GetPropertyMonitorManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("MonitorManufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonitorType sets the value of MonitorType for the instance
func (instance *Win32_DesktopMonitor) SetPropertyMonitorType(value string) (err error) {
	return instance.SetProperty("MonitorType", value)
}

// GetMonitorType gets the value of MonitorType for the instance
func (instance *Win32_DesktopMonitor) GetPropertyMonitorType() (value string, err error) {
	retValue, err := instance.GetProperty("MonitorType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPixelsPerXLogicalInch sets the value of PixelsPerXLogicalInch for the instance
func (instance *Win32_DesktopMonitor) SetPropertyPixelsPerXLogicalInch(value uint32) (err error) {
	return instance.SetProperty("PixelsPerXLogicalInch", value)
}

// GetPixelsPerXLogicalInch gets the value of PixelsPerXLogicalInch for the instance
func (instance *Win32_DesktopMonitor) GetPropertyPixelsPerXLogicalInch() (value uint32, err error) {
	retValue, err := instance.GetProperty("PixelsPerXLogicalInch")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPixelsPerYLogicalInch sets the value of PixelsPerYLogicalInch for the instance
func (instance *Win32_DesktopMonitor) SetPropertyPixelsPerYLogicalInch(value uint32) (err error) {
	return instance.SetProperty("PixelsPerYLogicalInch", value)
}

// GetPixelsPerYLogicalInch gets the value of PixelsPerYLogicalInch for the instance
func (instance *Win32_DesktopMonitor) GetPropertyPixelsPerYLogicalInch() (value uint32, err error) {
	retValue, err := instance.GetProperty("PixelsPerYLogicalInch")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
