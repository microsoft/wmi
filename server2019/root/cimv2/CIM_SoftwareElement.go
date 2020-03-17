// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_SoftwareElement struct
type CIM_SoftwareElement struct {
	CIM_LogicalElement

	//
	BuildNumber string

	//
	CodeSet string

	//
	IdentificationCode string

	//
	LanguageEdition string

	//
	Manufacturer string

	//
	OtherTargetOS string

	//
	SerialNumber string

	//
	SoftwareElementID string

	//
	SoftwareElementState uint16

	//
	TargetOperatingSystem uint16

	//
	Version string
}

// SetBuildNumber sets the value of BuildNumber for the instance
func (instance *CIM_SoftwareElement) SetPropertyBuildNumber(value string) (err error) {
	return instance.SetProperty("BuildNumber", value)
}

// GetBuildNumber gets the value of BuildNumber for the instance
func (instance *CIM_SoftwareElement) GetPropertyBuildNumber() (value string, err error) {
	retValue, err := instance.GetProperty("BuildNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCodeSet sets the value of CodeSet for the instance
func (instance *CIM_SoftwareElement) SetPropertyCodeSet(value string) (err error) {
	return instance.SetProperty("CodeSet", value)
}

// GetCodeSet gets the value of CodeSet for the instance
func (instance *CIM_SoftwareElement) GetPropertyCodeSet() (value string, err error) {
	retValue, err := instance.GetProperty("CodeSet")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdentificationCode sets the value of IdentificationCode for the instance
func (instance *CIM_SoftwareElement) SetPropertyIdentificationCode(value string) (err error) {
	return instance.SetProperty("IdentificationCode", value)
}

// GetIdentificationCode gets the value of IdentificationCode for the instance
func (instance *CIM_SoftwareElement) GetPropertyIdentificationCode() (value string, err error) {
	retValue, err := instance.GetProperty("IdentificationCode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLanguageEdition sets the value of LanguageEdition for the instance
func (instance *CIM_SoftwareElement) SetPropertyLanguageEdition(value string) (err error) {
	return instance.SetProperty("LanguageEdition", value)
}

// GetLanguageEdition gets the value of LanguageEdition for the instance
func (instance *CIM_SoftwareElement) GetPropertyLanguageEdition() (value string, err error) {
	retValue, err := instance.GetProperty("LanguageEdition")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *CIM_SoftwareElement) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *CIM_SoftwareElement) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherTargetOS sets the value of OtherTargetOS for the instance
func (instance *CIM_SoftwareElement) SetPropertyOtherTargetOS(value string) (err error) {
	return instance.SetProperty("OtherTargetOS", value)
}

// GetOtherTargetOS gets the value of OtherTargetOS for the instance
func (instance *CIM_SoftwareElement) GetPropertyOtherTargetOS() (value string, err error) {
	retValue, err := instance.GetProperty("OtherTargetOS")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *CIM_SoftwareElement) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", value)
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *CIM_SoftwareElement) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSoftwareElementID sets the value of SoftwareElementID for the instance
func (instance *CIM_SoftwareElement) SetPropertySoftwareElementID(value string) (err error) {
	return instance.SetProperty("SoftwareElementID", value)
}

// GetSoftwareElementID gets the value of SoftwareElementID for the instance
func (instance *CIM_SoftwareElement) GetPropertySoftwareElementID() (value string, err error) {
	retValue, err := instance.GetProperty("SoftwareElementID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSoftwareElementState sets the value of SoftwareElementState for the instance
func (instance *CIM_SoftwareElement) SetPropertySoftwareElementState(value uint16) (err error) {
	return instance.SetProperty("SoftwareElementState", value)
}

// GetSoftwareElementState gets the value of SoftwareElementState for the instance
func (instance *CIM_SoftwareElement) GetPropertySoftwareElementState() (value uint16, err error) {
	retValue, err := instance.GetProperty("SoftwareElementState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetOperatingSystem sets the value of TargetOperatingSystem for the instance
func (instance *CIM_SoftwareElement) SetPropertyTargetOperatingSystem(value uint16) (err error) {
	return instance.SetProperty("TargetOperatingSystem", value)
}

// GetTargetOperatingSystem gets the value of TargetOperatingSystem for the instance
func (instance *CIM_SoftwareElement) GetPropertyTargetOperatingSystem() (value uint16, err error) {
	retValue, err := instance.GetProperty("TargetOperatingSystem")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *CIM_SoftwareElement) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *CIM_SoftwareElement) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
