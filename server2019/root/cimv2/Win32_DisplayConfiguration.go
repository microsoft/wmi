// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DisplayConfiguration struct
type Win32_DisplayConfiguration struct {
	CIM_Setting

	//
	BitsPerPel uint32

	//
	DeviceName string

	//
	DisplayFlags uint32

	//
	DisplayFrequency uint32

	//
	DitherType uint32

	//
	DriverVersion string

	//
	ICMIntent uint32

	//
	ICMMethod uint32

	//
	LogPixels uint32

	//
	PelsHeight uint32

	//
	PelsWidth uint32

	//
	SpecificationVersion uint32
}

// SetBitsPerPel sets the value of BitsPerPel for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyBitsPerPel(value uint32) (err error) {
	return instance.SetProperty("BitsPerPel", value)
}

// GetBitsPerPel gets the value of BitsPerPel for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyBitsPerPel() (value uint32, err error) {
	retValue, err := instance.GetProperty("BitsPerPel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceName sets the value of DeviceName for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyDeviceName(value string) (err error) {
	return instance.SetProperty("DeviceName", value)
}

// GetDeviceName gets the value of DeviceName for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyDeviceName() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayFlags sets the value of DisplayFlags for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyDisplayFlags(value uint32) (err error) {
	return instance.SetProperty("DisplayFlags", value)
}

// GetDisplayFlags gets the value of DisplayFlags for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyDisplayFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisplayFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayFrequency sets the value of DisplayFrequency for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyDisplayFrequency(value uint32) (err error) {
	return instance.SetProperty("DisplayFrequency", value)
}

// GetDisplayFrequency gets the value of DisplayFrequency for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyDisplayFrequency() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisplayFrequency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDitherType sets the value of DitherType for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyDitherType(value uint32) (err error) {
	return instance.SetProperty("DitherType", value)
}

// GetDitherType gets the value of DitherType for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyDitherType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DitherType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverVersion sets the value of DriverVersion for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyDriverVersion(value string) (err error) {
	return instance.SetProperty("DriverVersion", value)
}

// GetDriverVersion gets the value of DriverVersion for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyDriverVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DriverVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetICMIntent sets the value of ICMIntent for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyICMIntent(value uint32) (err error) {
	return instance.SetProperty("ICMIntent", value)
}

// GetICMIntent gets the value of ICMIntent for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyICMIntent() (value uint32, err error) {
	retValue, err := instance.GetProperty("ICMIntent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetICMMethod sets the value of ICMMethod for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyICMMethod(value uint32) (err error) {
	return instance.SetProperty("ICMMethod", value)
}

// GetICMMethod gets the value of ICMMethod for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyICMMethod() (value uint32, err error) {
	retValue, err := instance.GetProperty("ICMMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogPixels sets the value of LogPixels for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyLogPixels(value uint32) (err error) {
	return instance.SetProperty("LogPixels", value)
}

// GetLogPixels gets the value of LogPixels for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyLogPixels() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogPixels")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPelsHeight sets the value of PelsHeight for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyPelsHeight(value uint32) (err error) {
	return instance.SetProperty("PelsHeight", value)
}

// GetPelsHeight gets the value of PelsHeight for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyPelsHeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("PelsHeight")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPelsWidth sets the value of PelsWidth for the instance
func (instance *Win32_DisplayConfiguration) SetPropertyPelsWidth(value uint32) (err error) {
	return instance.SetProperty("PelsWidth", value)
}

// GetPelsWidth gets the value of PelsWidth for the instance
func (instance *Win32_DisplayConfiguration) GetPropertyPelsWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("PelsWidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecificationVersion sets the value of SpecificationVersion for the instance
func (instance *Win32_DisplayConfiguration) SetPropertySpecificationVersion(value uint32) (err error) {
	return instance.SetProperty("SpecificationVersion", value)
}

// GetSpecificationVersion gets the value of SpecificationVersion for the instance
func (instance *Win32_DisplayConfiguration) GetPropertySpecificationVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("SpecificationVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
