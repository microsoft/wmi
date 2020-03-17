// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetAdapterPowerManagement_Offload_NS struct
type MSFT_NetAdapterPowerManagement_Offload_NS struct {
	MSFT_NetAdapterPowerManagement_Offload

	//
	MacAddress string

	//
	RemoteIPv6Address string

	//
	SolicitedNodeIPv6Address string

	//
	TargetIPv6Addresses []string
}

// SetMacAddress sets the value of MacAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) SetPropertyMacAddress(value string) (err error) {
	return instance.SetProperty("MacAddress", value)
}

// GetMacAddress gets the value of MacAddress for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) GetPropertyMacAddress() (value string, err error) {
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

// SetRemoteIPv6Address sets the value of RemoteIPv6Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) SetPropertyRemoteIPv6Address(value string) (err error) {
	return instance.SetProperty("RemoteIPv6Address", value)
}

// GetRemoteIPv6Address gets the value of RemoteIPv6Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) GetPropertyRemoteIPv6Address() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteIPv6Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSolicitedNodeIPv6Address sets the value of SolicitedNodeIPv6Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) SetPropertySolicitedNodeIPv6Address(value string) (err error) {
	return instance.SetProperty("SolicitedNodeIPv6Address", value)
}

// GetSolicitedNodeIPv6Address gets the value of SolicitedNodeIPv6Address for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) GetPropertySolicitedNodeIPv6Address() (value string, err error) {
	retValue, err := instance.GetProperty("SolicitedNodeIPv6Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetIPv6Addresses sets the value of TargetIPv6Addresses for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) SetPropertyTargetIPv6Addresses(value []string) (err error) {
	return instance.SetProperty("TargetIPv6Addresses", value)
}

// GetTargetIPv6Addresses gets the value of TargetIPv6Addresses for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_NS) GetPropertyTargetIPv6Addresses() (value []string, err error) {
	retValue, err := instance.GetProperty("TargetIPv6Addresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
