// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_ComputerSystem struct
type CIM_ComputerSystem struct {
	CIM_System

	//
	Dedicated []ComputerSystem_Dedicated

	//
	OtherDedicatedDescriptions []string

	//
	PowerManagementCapabilities []ComputerSystem_PowerManagementCapabilities

	//
	ResetCapability ComputerSystem_ResetCapability
}

// SetDedicated sets the value of Dedicated for the instance
func (instance *CIM_ComputerSystem) SetPropertyDedicated(value []ComputerSystem_Dedicated) (err error) {
	return instance.SetProperty("Dedicated", value)
}

// GetDedicated gets the value of Dedicated for the instance
func (instance *CIM_ComputerSystem) GetPropertyDedicated() (value []ComputerSystem_Dedicated, err error) {
	retValue, err := instance.GetProperty("Dedicated")
	if err != nil {
		return
	}
	value, ok := retValue.([]ComputerSystem_Dedicated)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherDedicatedDescriptions sets the value of OtherDedicatedDescriptions for the instance
func (instance *CIM_ComputerSystem) SetPropertyOtherDedicatedDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherDedicatedDescriptions", value)
}

// GetOtherDedicatedDescriptions gets the value of OtherDedicatedDescriptions for the instance
func (instance *CIM_ComputerSystem) GetPropertyOtherDedicatedDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherDedicatedDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerManagementCapabilities sets the value of PowerManagementCapabilities for the instance
func (instance *CIM_ComputerSystem) SetPropertyPowerManagementCapabilities(value []ComputerSystem_PowerManagementCapabilities) (err error) {
	return instance.SetProperty("PowerManagementCapabilities", value)
}

// GetPowerManagementCapabilities gets the value of PowerManagementCapabilities for the instance
func (instance *CIM_ComputerSystem) GetPropertyPowerManagementCapabilities() (value []ComputerSystem_PowerManagementCapabilities, err error) {
	retValue, err := instance.GetProperty("PowerManagementCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.([]ComputerSystem_PowerManagementCapabilities)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResetCapability sets the value of ResetCapability for the instance
func (instance *CIM_ComputerSystem) SetPropertyResetCapability(value ComputerSystem_ResetCapability) (err error) {
	return instance.SetProperty("ResetCapability", value)
}

// GetResetCapability gets the value of ResetCapability for the instance
func (instance *CIM_ComputerSystem) GetPropertyResetCapability() (value ComputerSystem_ResetCapability, err error) {
	retValue, err := instance.GetProperty("ResetCapability")
	if err != nil {
		return
	}
	value, ok := retValue.(ComputerSystem_ResetCapability)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="PowerState" type="ComputerSystem_PowerState "></param>
// <param name="Time" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ComputerSystem) SetPowerState( /* IN */ PowerState ComputerSystem_PowerState,
	/* IN */ Time string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPowerState", PowerState, Time)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
