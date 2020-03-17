// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_PciExpress struct
type Msvm_PciExpress struct {
	CIM_LogicalDevice

	//
	DeviceInstancePath string

	//
	FunctionNumber uint16

	//
	LocationPath string
}

// SetDeviceInstancePath sets the value of DeviceInstancePath for the instance
func (instance *Msvm_PciExpress) SetPropertyDeviceInstancePath(value string) (err error) {
	return instance.SetProperty("DeviceInstancePath", value)
}

// GetDeviceInstancePath gets the value of DeviceInstancePath for the instance
func (instance *Msvm_PciExpress) GetPropertyDeviceInstancePath() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceInstancePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFunctionNumber sets the value of FunctionNumber for the instance
func (instance *Msvm_PciExpress) SetPropertyFunctionNumber(value uint16) (err error) {
	return instance.SetProperty("FunctionNumber", value)
}

// GetFunctionNumber gets the value of FunctionNumber for the instance
func (instance *Msvm_PciExpress) GetPropertyFunctionNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("FunctionNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocationPath sets the value of LocationPath for the instance
func (instance *Msvm_PciExpress) SetPropertyLocationPath(value string) (err error) {
	return instance.SetProperty("LocationPath", value)
}

// GetLocationPath gets the value of LocationPath for the instance
func (instance *Msvm_PciExpress) GetPropertyLocationPath() (value string, err error) {
	retValue, err := instance.GetProperty("LocationPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
