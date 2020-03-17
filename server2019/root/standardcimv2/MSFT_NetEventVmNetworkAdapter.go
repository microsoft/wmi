// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetEventVmNetworkAdapter struct
type MSFT_NetEventVmNetworkAdapter struct {
	MSFT_NetEventPacketCaptureTarget

	//
	MacAddress string

	//
	PortName string

	//
	SwitchName string

	//
	VMId string

	//
	VMName string
}

// SetMacAddress sets the value of MacAddress for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) SetPropertyMacAddress(value string) (err error) {
	return instance.SetProperty("MacAddress", value)
}

// GetMacAddress gets the value of MacAddress for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) GetPropertyMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MacAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortName sets the value of PortName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) SetPropertyPortName(value string) (err error) {
	return instance.SetProperty("PortName", value)
}

// GetPortName gets the value of PortName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) GetPropertyPortName() (value string, err error) {
	retValue, err := instance.GetProperty("PortName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSwitchName sets the value of SwitchName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) SetPropertySwitchName(value string) (err error) {
	return instance.SetProperty("SwitchName", value)
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) GetPropertySwitchName() (value string, err error) {
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

// SetVMId sets the value of VMId for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) SetPropertyVMId(value string) (err error) {
	return instance.SetProperty("VMId", value)
}

// GetVMId gets the value of VMId for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) GetPropertyVMId() (value string, err error) {
	retValue, err := instance.GetProperty("VMId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVMName sets the value of VMName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) SetPropertyVMName(value string) (err error) {
	return instance.SetProperty("VMName", value)
}

// GetVMName gets the value of VMName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) GetPropertyVMName() (value string, err error) {
	retValue, err := instance.GetProperty("VMName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
