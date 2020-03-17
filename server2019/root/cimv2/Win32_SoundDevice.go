// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_SoundDevice struct
type Win32_SoundDevice struct {
	CIM_LogicalDevice

	//
	DMABufferSize uint16

	//
	Manufacturer string

	//
	MPU401Address uint32

	//
	ProductName string
}

// SetDMABufferSize sets the value of DMABufferSize for the instance
func (instance *Win32_SoundDevice) SetPropertyDMABufferSize(value uint16) (err error) {
	return instance.SetProperty("DMABufferSize", value)
}

// GetDMABufferSize gets the value of DMABufferSize for the instance
func (instance *Win32_SoundDevice) GetPropertyDMABufferSize() (value uint16, err error) {
	retValue, err := instance.GetProperty("DMABufferSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *Win32_SoundDevice) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *Win32_SoundDevice) GetPropertyManufacturer() (value string, err error) {
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

// SetMPU401Address sets the value of MPU401Address for the instance
func (instance *Win32_SoundDevice) SetPropertyMPU401Address(value uint32) (err error) {
	return instance.SetProperty("MPU401Address", value)
}

// GetMPU401Address gets the value of MPU401Address for the instance
func (instance *Win32_SoundDevice) GetPropertyMPU401Address() (value uint32, err error) {
	retValue, err := instance.GetProperty("MPU401Address")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProductName sets the value of ProductName for the instance
func (instance *Win32_SoundDevice) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", value)
}

// GetProductName gets the value of ProductName for the instance
func (instance *Win32_SoundDevice) GetPropertyProductName() (value string, err error) {
	retValue, err := instance.GetProperty("ProductName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
