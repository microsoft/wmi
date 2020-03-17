// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_DesktopMonitor struct
type CIM_DesktopMonitor struct {
	CIM_Display

	// Monitor's bandwidth in MHertz. If unknown, enter 0.
	Bandwidth uint32

	// The type of DesktopMonitor or CRT. For example, multiscan color or monochrome monitors (values 2 or 3, respectively) can be indicated in this property.
	DisplayType DesktopMonitor_DisplayType

	// The logical height of the Display in screen coordinates.
	ScreenHeight uint32

	// The logical width of the Display in screen coordinates.
	ScreenWidth uint32
}

// SetBandwidth sets the value of Bandwidth for the instance
func (instance *CIM_DesktopMonitor) SetPropertyBandwidth(value uint32) (err error) {
	return instance.SetProperty("Bandwidth", value)
}

// GetBandwidth gets the value of Bandwidth for the instance
func (instance *CIM_DesktopMonitor) GetPropertyBandwidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("Bandwidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayType sets the value of DisplayType for the instance
func (instance *CIM_DesktopMonitor) SetPropertyDisplayType(value DesktopMonitor_DisplayType) (err error) {
	return instance.SetProperty("DisplayType", value)
}

// GetDisplayType gets the value of DisplayType for the instance
func (instance *CIM_DesktopMonitor) GetPropertyDisplayType() (value DesktopMonitor_DisplayType, err error) {
	retValue, err := instance.GetProperty("DisplayType")
	if err != nil {
		return
	}
	value, ok := retValue.(DesktopMonitor_DisplayType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScreenHeight sets the value of ScreenHeight for the instance
func (instance *CIM_DesktopMonitor) SetPropertyScreenHeight(value uint32) (err error) {
	return instance.SetProperty("ScreenHeight", value)
}

// GetScreenHeight gets the value of ScreenHeight for the instance
func (instance *CIM_DesktopMonitor) GetPropertyScreenHeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScreenHeight")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScreenWidth sets the value of ScreenWidth for the instance
func (instance *CIM_DesktopMonitor) SetPropertyScreenWidth(value uint32) (err error) {
	return instance.SetProperty("ScreenWidth", value)
}

// GetScreenWidth gets the value of ScreenWidth for the instance
func (instance *CIM_DesktopMonitor) GetPropertyScreenWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScreenWidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
