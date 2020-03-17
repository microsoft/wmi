// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_BootSourceSettingData struct
type Msvm_BootSourceSettingData struct {
	CIM_SettingData

	//
	BootSourceDescription string

	//
	BootSourceType uint32

	//
	FirmwareDevicePath string

	//
	OptionalData []uint8

	//
	OtherLocation string
}

// SetBootSourceDescription sets the value of BootSourceDescription for the instance
func (instance *Msvm_BootSourceSettingData) SetPropertyBootSourceDescription(value string) (err error) {
	return instance.SetProperty("BootSourceDescription", value)
}

// GetBootSourceDescription gets the value of BootSourceDescription for the instance
func (instance *Msvm_BootSourceSettingData) GetPropertyBootSourceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("BootSourceDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBootSourceType sets the value of BootSourceType for the instance
func (instance *Msvm_BootSourceSettingData) SetPropertyBootSourceType(value uint32) (err error) {
	return instance.SetProperty("BootSourceType", value)
}

// GetBootSourceType gets the value of BootSourceType for the instance
func (instance *Msvm_BootSourceSettingData) GetPropertyBootSourceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("BootSourceType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFirmwareDevicePath sets the value of FirmwareDevicePath for the instance
func (instance *Msvm_BootSourceSettingData) SetPropertyFirmwareDevicePath(value string) (err error) {
	return instance.SetProperty("FirmwareDevicePath", value)
}

// GetFirmwareDevicePath gets the value of FirmwareDevicePath for the instance
func (instance *Msvm_BootSourceSettingData) GetPropertyFirmwareDevicePath() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareDevicePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptionalData sets the value of OptionalData for the instance
func (instance *Msvm_BootSourceSettingData) SetPropertyOptionalData(value []uint8) (err error) {
	return instance.SetProperty("OptionalData", value)
}

// GetOptionalData gets the value of OptionalData for the instance
func (instance *Msvm_BootSourceSettingData) GetPropertyOptionalData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("OptionalData")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherLocation sets the value of OtherLocation for the instance
func (instance *Msvm_BootSourceSettingData) SetPropertyOtherLocation(value string) (err error) {
	return instance.SetProperty("OtherLocation", value)
}

// GetOtherLocation gets the value of OtherLocation for the instance
func (instance *Msvm_BootSourceSettingData) GetPropertyOtherLocation() (value string, err error) {
	retValue, err := instance.GetProperty("OtherLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
