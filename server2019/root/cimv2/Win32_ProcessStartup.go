// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ProcessStartup struct
type Win32_ProcessStartup struct {
	Win32_MethodParameterClass

	//
	CreateFlags uint32

	//
	EnvironmentVariables []string

	//
	ErrorMode uint16

	//
	FillAttribute uint32

	//
	PriorityClass uint32

	//
	ShowWindow uint16

	//
	Title string

	//
	WinstationDesktop string

	//
	X uint32

	//
	XCountChars uint32

	//
	XSize uint32

	//
	Y uint32

	//
	YCountChars uint32

	//
	YSize uint32
}

// SetCreateFlags sets the value of CreateFlags for the instance
func (instance *Win32_ProcessStartup) SetPropertyCreateFlags(value uint32) (err error) {
	return instance.SetProperty("CreateFlags", value)
}

// GetCreateFlags gets the value of CreateFlags for the instance
func (instance *Win32_ProcessStartup) GetPropertyCreateFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("CreateFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnvironmentVariables sets the value of EnvironmentVariables for the instance
func (instance *Win32_ProcessStartup) SetPropertyEnvironmentVariables(value []string) (err error) {
	return instance.SetProperty("EnvironmentVariables", value)
}

// GetEnvironmentVariables gets the value of EnvironmentVariables for the instance
func (instance *Win32_ProcessStartup) GetPropertyEnvironmentVariables() (value []string, err error) {
	retValue, err := instance.GetProperty("EnvironmentVariables")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorMode sets the value of ErrorMode for the instance
func (instance *Win32_ProcessStartup) SetPropertyErrorMode(value uint16) (err error) {
	return instance.SetProperty("ErrorMode", value)
}

// GetErrorMode gets the value of ErrorMode for the instance
func (instance *Win32_ProcessStartup) GetPropertyErrorMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFillAttribute sets the value of FillAttribute for the instance
func (instance *Win32_ProcessStartup) SetPropertyFillAttribute(value uint32) (err error) {
	return instance.SetProperty("FillAttribute", value)
}

// GetFillAttribute gets the value of FillAttribute for the instance
func (instance *Win32_ProcessStartup) GetPropertyFillAttribute() (value uint32, err error) {
	retValue, err := instance.GetProperty("FillAttribute")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriorityClass sets the value of PriorityClass for the instance
func (instance *Win32_ProcessStartup) SetPropertyPriorityClass(value uint32) (err error) {
	return instance.SetProperty("PriorityClass", value)
}

// GetPriorityClass gets the value of PriorityClass for the instance
func (instance *Win32_ProcessStartup) GetPropertyPriorityClass() (value uint32, err error) {
	retValue, err := instance.GetProperty("PriorityClass")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShowWindow sets the value of ShowWindow for the instance
func (instance *Win32_ProcessStartup) SetPropertyShowWindow(value uint16) (err error) {
	return instance.SetProperty("ShowWindow", value)
}

// GetShowWindow gets the value of ShowWindow for the instance
func (instance *Win32_ProcessStartup) GetPropertyShowWindow() (value uint16, err error) {
	retValue, err := instance.GetProperty("ShowWindow")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTitle sets the value of Title for the instance
func (instance *Win32_ProcessStartup) SetPropertyTitle(value string) (err error) {
	return instance.SetProperty("Title", value)
}

// GetTitle gets the value of Title for the instance
func (instance *Win32_ProcessStartup) GetPropertyTitle() (value string, err error) {
	retValue, err := instance.GetProperty("Title")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWinstationDesktop sets the value of WinstationDesktop for the instance
func (instance *Win32_ProcessStartup) SetPropertyWinstationDesktop(value string) (err error) {
	return instance.SetProperty("WinstationDesktop", value)
}

// GetWinstationDesktop gets the value of WinstationDesktop for the instance
func (instance *Win32_ProcessStartup) GetPropertyWinstationDesktop() (value string, err error) {
	retValue, err := instance.GetProperty("WinstationDesktop")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetX sets the value of X for the instance
func (instance *Win32_ProcessStartup) SetPropertyX(value uint32) (err error) {
	return instance.SetProperty("X", value)
}

// GetX gets the value of X for the instance
func (instance *Win32_ProcessStartup) GetPropertyX() (value uint32, err error) {
	retValue, err := instance.GetProperty("X")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetXCountChars sets the value of XCountChars for the instance
func (instance *Win32_ProcessStartup) SetPropertyXCountChars(value uint32) (err error) {
	return instance.SetProperty("XCountChars", value)
}

// GetXCountChars gets the value of XCountChars for the instance
func (instance *Win32_ProcessStartup) GetPropertyXCountChars() (value uint32, err error) {
	retValue, err := instance.GetProperty("XCountChars")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetXSize sets the value of XSize for the instance
func (instance *Win32_ProcessStartup) SetPropertyXSize(value uint32) (err error) {
	return instance.SetProperty("XSize", value)
}

// GetXSize gets the value of XSize for the instance
func (instance *Win32_ProcessStartup) GetPropertyXSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("XSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetY sets the value of Y for the instance
func (instance *Win32_ProcessStartup) SetPropertyY(value uint32) (err error) {
	return instance.SetProperty("Y", value)
}

// GetY gets the value of Y for the instance
func (instance *Win32_ProcessStartup) GetPropertyY() (value uint32, err error) {
	retValue, err := instance.GetProperty("Y")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetYCountChars sets the value of YCountChars for the instance
func (instance *Win32_ProcessStartup) SetPropertyYCountChars(value uint32) (err error) {
	return instance.SetProperty("YCountChars", value)
}

// GetYCountChars gets the value of YCountChars for the instance
func (instance *Win32_ProcessStartup) GetPropertyYCountChars() (value uint32, err error) {
	retValue, err := instance.GetProperty("YCountChars")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetYSize sets the value of YSize for the instance
func (instance *Win32_ProcessStartup) SetPropertyYSize(value uint32) (err error) {
	return instance.SetProperty("YSize", value)
}

// GetYSize gets the value of YSize for the instance
func (instance *Win32_ProcessStartup) GetPropertyYSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("YSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
