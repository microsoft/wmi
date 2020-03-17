// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DisplayControllerConfiguration struct
type Win32_DisplayControllerConfiguration struct {
	CIM_Setting

	//
	BitsPerPixel uint32

	//
	ColorPlanes uint32

	//
	DeviceEntriesInAColorTable uint32

	//
	DeviceSpecificPens uint32

	//
	HorizontalResolution uint32

	//
	Name string

	//
	RefreshRate int32

	//
	ReservedSystemPaletteEntries uint32

	//
	SystemPaletteEntries uint32

	//
	VerticalResolution uint32

	//
	VideoMode string
}

// SetBitsPerPixel sets the value of BitsPerPixel for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertyBitsPerPixel(value uint32) (err error) {
	return instance.SetProperty("BitsPerPixel", value)
}

// GetBitsPerPixel gets the value of BitsPerPixel for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertyBitsPerPixel() (value uint32, err error) {
	retValue, err := instance.GetProperty("BitsPerPixel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetColorPlanes sets the value of ColorPlanes for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertyColorPlanes(value uint32) (err error) {
	return instance.SetProperty("ColorPlanes", value)
}

// GetColorPlanes gets the value of ColorPlanes for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertyColorPlanes() (value uint32, err error) {
	retValue, err := instance.GetProperty("ColorPlanes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceEntriesInAColorTable sets the value of DeviceEntriesInAColorTable for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertyDeviceEntriesInAColorTable(value uint32) (err error) {
	return instance.SetProperty("DeviceEntriesInAColorTable", value)
}

// GetDeviceEntriesInAColorTable gets the value of DeviceEntriesInAColorTable for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertyDeviceEntriesInAColorTable() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceEntriesInAColorTable")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceSpecificPens sets the value of DeviceSpecificPens for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertyDeviceSpecificPens(value uint32) (err error) {
	return instance.SetProperty("DeviceSpecificPens", value)
}

// GetDeviceSpecificPens gets the value of DeviceSpecificPens for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertyDeviceSpecificPens() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceSpecificPens")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHorizontalResolution sets the value of HorizontalResolution for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertyHorizontalResolution(value uint32) (err error) {
	return instance.SetProperty("HorizontalResolution", value)
}

// GetHorizontalResolution gets the value of HorizontalResolution for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertyHorizontalResolution() (value uint32, err error) {
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

// SetName sets the value of Name for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRefreshRate sets the value of RefreshRate for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertyRefreshRate(value int32) (err error) {
	return instance.SetProperty("RefreshRate", value)
}

// GetRefreshRate gets the value of RefreshRate for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertyRefreshRate() (value int32, err error) {
	retValue, err := instance.GetProperty("RefreshRate")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReservedSystemPaletteEntries sets the value of ReservedSystemPaletteEntries for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertyReservedSystemPaletteEntries(value uint32) (err error) {
	return instance.SetProperty("ReservedSystemPaletteEntries", value)
}

// GetReservedSystemPaletteEntries gets the value of ReservedSystemPaletteEntries for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertyReservedSystemPaletteEntries() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReservedSystemPaletteEntries")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemPaletteEntries sets the value of SystemPaletteEntries for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertySystemPaletteEntries(value uint32) (err error) {
	return instance.SetProperty("SystemPaletteEntries", value)
}

// GetSystemPaletteEntries gets the value of SystemPaletteEntries for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertySystemPaletteEntries() (value uint32, err error) {
	retValue, err := instance.GetProperty("SystemPaletteEntries")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVerticalResolution sets the value of VerticalResolution for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertyVerticalResolution(value uint32) (err error) {
	return instance.SetProperty("VerticalResolution", value)
}

// GetVerticalResolution gets the value of VerticalResolution for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertyVerticalResolution() (value uint32, err error) {
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

// SetVideoMode sets the value of VideoMode for the instance
func (instance *Win32_DisplayControllerConfiguration) SetPropertyVideoMode(value string) (err error) {
	return instance.SetProperty("VideoMode", value)
}

// GetVideoMode gets the value of VideoMode for the instance
func (instance *Win32_DisplayControllerConfiguration) GetPropertyVideoMode() (value string, err error) {
	retValue, err := instance.GetProperty("VideoMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
