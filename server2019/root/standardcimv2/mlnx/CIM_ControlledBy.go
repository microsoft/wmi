// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_ControlledBy struct
type CIM_ControlledBy struct {
	CIM_DeviceConnection

	//
	AccessMode ControlledBy_AccessMode

	//
	AccessPriority uint16

	//
	AccessState ControlledBy_AccessState

	//
	DeviceNumber string

	//
	NumberOfHardResets uint32

	//
	NumberOfSoftResets uint32

	//
	TimeOfDeviceReset string
}

// SetAccessMode sets the value of AccessMode for the instance
func (instance *CIM_ControlledBy) SetPropertyAccessMode(value ControlledBy_AccessMode) (err error) {
	return instance.SetProperty("AccessMode", value)
}

// GetAccessMode gets the value of AccessMode for the instance
func (instance *CIM_ControlledBy) GetPropertyAccessMode() (value ControlledBy_AccessMode, err error) {
	retValue, err := instance.GetProperty("AccessMode")
	if err != nil {
		return
	}
	value, ok := retValue.(ControlledBy_AccessMode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccessPriority sets the value of AccessPriority for the instance
func (instance *CIM_ControlledBy) SetPropertyAccessPriority(value uint16) (err error) {
	return instance.SetProperty("AccessPriority", value)
}

// GetAccessPriority gets the value of AccessPriority for the instance
func (instance *CIM_ControlledBy) GetPropertyAccessPriority() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessPriority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccessState sets the value of AccessState for the instance
func (instance *CIM_ControlledBy) SetPropertyAccessState(value ControlledBy_AccessState) (err error) {
	return instance.SetProperty("AccessState", value)
}

// GetAccessState gets the value of AccessState for the instance
func (instance *CIM_ControlledBy) GetPropertyAccessState() (value ControlledBy_AccessState, err error) {
	retValue, err := instance.GetProperty("AccessState")
	if err != nil {
		return
	}
	value, ok := retValue.(ControlledBy_AccessState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceNumber sets the value of DeviceNumber for the instance
func (instance *CIM_ControlledBy) SetPropertyDeviceNumber(value string) (err error) {
	return instance.SetProperty("DeviceNumber", value)
}

// GetDeviceNumber gets the value of DeviceNumber for the instance
func (instance *CIM_ControlledBy) GetPropertyDeviceNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfHardResets sets the value of NumberOfHardResets for the instance
func (instance *CIM_ControlledBy) SetPropertyNumberOfHardResets(value uint32) (err error) {
	return instance.SetProperty("NumberOfHardResets", value)
}

// GetNumberOfHardResets gets the value of NumberOfHardResets for the instance
func (instance *CIM_ControlledBy) GetPropertyNumberOfHardResets() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfHardResets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfSoftResets sets the value of NumberOfSoftResets for the instance
func (instance *CIM_ControlledBy) SetPropertyNumberOfSoftResets(value uint32) (err error) {
	return instance.SetProperty("NumberOfSoftResets", value)
}

// GetNumberOfSoftResets gets the value of NumberOfSoftResets for the instance
func (instance *CIM_ControlledBy) GetPropertyNumberOfSoftResets() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfSoftResets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeOfDeviceReset sets the value of TimeOfDeviceReset for the instance
func (instance *CIM_ControlledBy) SetPropertyTimeOfDeviceReset(value string) (err error) {
	return instance.SetProperty("TimeOfDeviceReset", value)
}

// GetTimeOfDeviceReset gets the value of TimeOfDeviceReset for the instance
func (instance *CIM_ControlledBy) GetPropertyTimeOfDeviceReset() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfDeviceReset")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
