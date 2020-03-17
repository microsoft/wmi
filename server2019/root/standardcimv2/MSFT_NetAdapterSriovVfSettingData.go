// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetAdapterSriovVfSettingData struct
type MSFT_NetAdapterSriovVfSettingData struct {
	MSFT_NetAdapterSettingData

	//
	CurrentMacAddress string

	//
	FunctionID uint16

	//
	PermanentMacAddress string

	//
	SwitchID uint32

	//
	VmFriendlyName string

	//
	VmID string

	//
	VmNicID string

	//
	VPortID []uint32
}

// SetCurrentMacAddress sets the value of CurrentMacAddress for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyCurrentMacAddress(value string) (err error) {
	return instance.SetProperty("CurrentMacAddress", value)
}

// GetCurrentMacAddress gets the value of CurrentMacAddress for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyCurrentMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentMacAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFunctionID sets the value of FunctionID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyFunctionID(value uint16) (err error) {
	return instance.SetProperty("FunctionID", value)
}

// GetFunctionID gets the value of FunctionID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyFunctionID() (value uint16, err error) {
	retValue, err := instance.GetProperty("FunctionID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPermanentMacAddress sets the value of PermanentMacAddress for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyPermanentMacAddress(value string) (err error) {
	return instance.SetProperty("PermanentMacAddress", value)
}

// GetPermanentMacAddress gets the value of PermanentMacAddress for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyPermanentMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("PermanentMacAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSwitchID sets the value of SwitchID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertySwitchID(value uint32) (err error) {
	return instance.SetProperty("SwitchID", value)
}

// GetSwitchID gets the value of SwitchID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertySwitchID() (value uint32, err error) {
	retValue, err := instance.GetProperty("SwitchID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmFriendlyName sets the value of VmFriendlyName for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyVmFriendlyName(value string) (err error) {
	return instance.SetProperty("VmFriendlyName", value)
}

// GetVmFriendlyName gets the value of VmFriendlyName for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyVmFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("VmFriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmID sets the value of VmID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyVmID(value string) (err error) {
	return instance.SetProperty("VmID", value)
}

// GetVmID gets the value of VmID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyVmID() (value string, err error) {
	retValue, err := instance.GetProperty("VmID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmNicID sets the value of VmNicID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyVmNicID(value string) (err error) {
	return instance.SetProperty("VmNicID", value)
}

// GetVmNicID gets the value of VmNicID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyVmNicID() (value string, err error) {
	retValue, err := instance.GetProperty("VmNicID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVPortID sets the value of VPortID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) SetPropertyVPortID(value []uint32) (err error) {
	return instance.SetProperty("VPortID", value)
}

// GetVPortID gets the value of VPortID for the instance
func (instance *MSFT_NetAdapterSriovVfSettingData) GetPropertyVPortID() (value []uint32, err error) {
	retValue, err := instance.GetProperty("VPortID")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
