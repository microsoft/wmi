// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetEventVmSwitchProvider struct
type MSFT_NetEventVmSwitchProvider struct {
	MSFT_NetEventProviderBase

	//
	PortIds []uint32

	//
	SwitchName string
}

// SetPortIds sets the value of PortIds for the instance
func (instance *MSFT_NetEventVmSwitchProvider) SetPropertyPortIds(value []uint32) (err error) {
	return instance.SetProperty("PortIds", value)
}

// GetPortIds gets the value of PortIds for the instance
func (instance *MSFT_NetEventVmSwitchProvider) GetPropertyPortIds() (value []uint32, err error) {
	retValue, err := instance.GetProperty("PortIds")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSwitchName sets the value of SwitchName for the instance
func (instance *MSFT_NetEventVmSwitchProvider) SetPropertySwitchName(value string) (err error) {
	return instance.SetProperty("SwitchName", value)
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *MSFT_NetEventVmSwitchProvider) GetPropertySwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
