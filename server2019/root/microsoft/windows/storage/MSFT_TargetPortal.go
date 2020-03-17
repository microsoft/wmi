// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_TargetPortal struct
type MSFT_TargetPortal struct {
	MSFT_StorageObject

	//
	IPv4Address string

	//
	IPv6Address string

	//
	PortNumber uint32

	//
	SubnetMask string
}

// SetIPv4Address sets the value of IPv4Address for the instance
func (instance *MSFT_TargetPortal) SetPropertyIPv4Address(value string) (err error) {
	return instance.SetProperty("IPv4Address", value)
}

// GetIPv4Address gets the value of IPv4Address for the instance
func (instance *MSFT_TargetPortal) GetPropertyIPv4Address() (value string, err error) {
	retValue, err := instance.GetProperty("IPv4Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6Address sets the value of IPv6Address for the instance
func (instance *MSFT_TargetPortal) SetPropertyIPv6Address(value string) (err error) {
	return instance.SetProperty("IPv6Address", value)
}

// GetIPv6Address gets the value of IPv6Address for the instance
func (instance *MSFT_TargetPortal) GetPropertyIPv6Address() (value string, err error) {
	retValue, err := instance.GetProperty("IPv6Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortNumber sets the value of PortNumber for the instance
func (instance *MSFT_TargetPortal) SetPropertyPortNumber(value uint32) (err error) {
	return instance.SetProperty("PortNumber", value)
}

// GetPortNumber gets the value of PortNumber for the instance
func (instance *MSFT_TargetPortal) GetPropertyPortNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubnetMask sets the value of SubnetMask for the instance
func (instance *MSFT_TargetPortal) SetPropertySubnetMask(value string) (err error) {
	return instance.SetProperty("SubnetMask", value)
}

// GetSubnetMask gets the value of SubnetMask for the instance
func (instance *MSFT_TargetPortal) GetPropertySubnetMask() (value string, err error) {
	retValue, err := instance.GetProperty("SubnetMask")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
