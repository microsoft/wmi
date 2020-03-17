// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_UnitaryComputerSystem struct
type CIM_UnitaryComputerSystem struct {
	CIM_ComputerSystem

	//
	InitialLoadInfo []string

	//
	LastLoadInfo string

	//
	PowerManagementCapabilities []uint16

	//
	PowerManagementSupported bool

	//
	PowerState uint16

	//
	ResetCapability uint16
}

// SetInitialLoadInfo sets the value of InitialLoadInfo for the instance
func (instance *CIM_UnitaryComputerSystem) SetPropertyInitialLoadInfo(value []string) (err error) {
	return instance.SetProperty("InitialLoadInfo", value)
}

// GetInitialLoadInfo gets the value of InitialLoadInfo for the instance
func (instance *CIM_UnitaryComputerSystem) GetPropertyInitialLoadInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("InitialLoadInfo")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastLoadInfo sets the value of LastLoadInfo for the instance
func (instance *CIM_UnitaryComputerSystem) SetPropertyLastLoadInfo(value string) (err error) {
	return instance.SetProperty("LastLoadInfo", value)
}

// GetLastLoadInfo gets the value of LastLoadInfo for the instance
func (instance *CIM_UnitaryComputerSystem) GetPropertyLastLoadInfo() (value string, err error) {
	retValue, err := instance.GetProperty("LastLoadInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerManagementCapabilities sets the value of PowerManagementCapabilities for the instance
func (instance *CIM_UnitaryComputerSystem) SetPropertyPowerManagementCapabilities(value []uint16) (err error) {
	return instance.SetProperty("PowerManagementCapabilities", value)
}

// GetPowerManagementCapabilities gets the value of PowerManagementCapabilities for the instance
func (instance *CIM_UnitaryComputerSystem) GetPropertyPowerManagementCapabilities() (value []uint16, err error) {
	retValue, err := instance.GetProperty("PowerManagementCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerManagementSupported sets the value of PowerManagementSupported for the instance
func (instance *CIM_UnitaryComputerSystem) SetPropertyPowerManagementSupported(value bool) (err error) {
	return instance.SetProperty("PowerManagementSupported", value)
}

// GetPowerManagementSupported gets the value of PowerManagementSupported for the instance
func (instance *CIM_UnitaryComputerSystem) GetPropertyPowerManagementSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("PowerManagementSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerState sets the value of PowerState for the instance
func (instance *CIM_UnitaryComputerSystem) SetPropertyPowerState(value uint16) (err error) {
	return instance.SetProperty("PowerState", value)
}

// GetPowerState gets the value of PowerState for the instance
func (instance *CIM_UnitaryComputerSystem) GetPropertyPowerState() (value uint16, err error) {
	retValue, err := instance.GetProperty("PowerState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResetCapability sets the value of ResetCapability for the instance
func (instance *CIM_UnitaryComputerSystem) SetPropertyResetCapability(value uint16) (err error) {
	return instance.SetProperty("ResetCapability", value)
}

// GetResetCapability gets the value of ResetCapability for the instance
func (instance *CIM_UnitaryComputerSystem) GetPropertyResetCapability() (value uint16, err error) {
	retValue, err := instance.GetProperty("ResetCapability")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="PowerState" type="uint16 "></param>
// <param name="Time" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_UnitaryComputerSystem) SetPowerState( /* IN */ PowerState uint16,
	/* IN */ Time string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPowerState", PowerState, Time)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
