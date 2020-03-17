// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_SoftwareElementVersionCheck struct
type CIM_SoftwareElementVersionCheck struct {
	CIM_Check

	//
	LowerSoftwareElementVersion string

	//
	SoftwareElementName string

	//
	SoftwareElementStateDesired uint16

	//
	TargetOperatingSystemDesired uint16

	//
	UpperSoftwareElementVersion string
}

// SetLowerSoftwareElementVersion sets the value of LowerSoftwareElementVersion for the instance
func (instance *CIM_SoftwareElementVersionCheck) SetPropertyLowerSoftwareElementVersion(value string) (err error) {
	return instance.SetProperty("LowerSoftwareElementVersion", value)
}

// GetLowerSoftwareElementVersion gets the value of LowerSoftwareElementVersion for the instance
func (instance *CIM_SoftwareElementVersionCheck) GetPropertyLowerSoftwareElementVersion() (value string, err error) {
	retValue, err := instance.GetProperty("LowerSoftwareElementVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSoftwareElementName sets the value of SoftwareElementName for the instance
func (instance *CIM_SoftwareElementVersionCheck) SetPropertySoftwareElementName(value string) (err error) {
	return instance.SetProperty("SoftwareElementName", value)
}

// GetSoftwareElementName gets the value of SoftwareElementName for the instance
func (instance *CIM_SoftwareElementVersionCheck) GetPropertySoftwareElementName() (value string, err error) {
	retValue, err := instance.GetProperty("SoftwareElementName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSoftwareElementStateDesired sets the value of SoftwareElementStateDesired for the instance
func (instance *CIM_SoftwareElementVersionCheck) SetPropertySoftwareElementStateDesired(value uint16) (err error) {
	return instance.SetProperty("SoftwareElementStateDesired", value)
}

// GetSoftwareElementStateDesired gets the value of SoftwareElementStateDesired for the instance
func (instance *CIM_SoftwareElementVersionCheck) GetPropertySoftwareElementStateDesired() (value uint16, err error) {
	retValue, err := instance.GetProperty("SoftwareElementStateDesired")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetOperatingSystemDesired sets the value of TargetOperatingSystemDesired for the instance
func (instance *CIM_SoftwareElementVersionCheck) SetPropertyTargetOperatingSystemDesired(value uint16) (err error) {
	return instance.SetProperty("TargetOperatingSystemDesired", value)
}

// GetTargetOperatingSystemDesired gets the value of TargetOperatingSystemDesired for the instance
func (instance *CIM_SoftwareElementVersionCheck) GetPropertyTargetOperatingSystemDesired() (value uint16, err error) {
	retValue, err := instance.GetProperty("TargetOperatingSystemDesired")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpperSoftwareElementVersion sets the value of UpperSoftwareElementVersion for the instance
func (instance *CIM_SoftwareElementVersionCheck) SetPropertyUpperSoftwareElementVersion(value string) (err error) {
	return instance.SetProperty("UpperSoftwareElementVersion", value)
}

// GetUpperSoftwareElementVersion gets the value of UpperSoftwareElementVersion for the instance
func (instance *CIM_SoftwareElementVersionCheck) GetPropertyUpperSoftwareElementVersion() (value string, err error) {
	retValue, err := instance.GetProperty("UpperSoftwareElementVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
