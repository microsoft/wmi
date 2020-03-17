// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterLargeSendOffloadV1Capabilities struct
type MSFT_NetAdapterLargeSendOffloadV1Capabilities struct {
	cim.WmiInstance

	//
	IPv4Encapsulation MSFT_NetAdapterLsoEncapsulationTypes

	//
	IPv4IpOptionsSupported bool

	//
	IPv4MaxOffloadSizeSupported uint32

	//
	IPv4MinSegmentCountSupported uint32

	//
	IPv4TcpOptionsSupported bool
}

// SetIPv4Encapsulation sets the value of IPv4Encapsulation for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) SetPropertyIPv4Encapsulation(value MSFT_NetAdapterLsoEncapsulationTypes) (err error) {
	return instance.SetProperty("IPv4Encapsulation", value)
}

// GetIPv4Encapsulation gets the value of IPv4Encapsulation for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) GetPropertyIPv4Encapsulation() (value MSFT_NetAdapterLsoEncapsulationTypes, err error) {
	retValue, err := instance.GetProperty("IPv4Encapsulation")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapterLsoEncapsulationTypes)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4IpOptionsSupported sets the value of IPv4IpOptionsSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) SetPropertyIPv4IpOptionsSupported(value bool) (err error) {
	return instance.SetProperty("IPv4IpOptionsSupported", value)
}

// GetIPv4IpOptionsSupported gets the value of IPv4IpOptionsSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) GetPropertyIPv4IpOptionsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4IpOptionsSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4MaxOffloadSizeSupported sets the value of IPv4MaxOffloadSizeSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) SetPropertyIPv4MaxOffloadSizeSupported(value uint32) (err error) {
	return instance.SetProperty("IPv4MaxOffloadSizeSupported", value)
}

// GetIPv4MaxOffloadSizeSupported gets the value of IPv4MaxOffloadSizeSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) GetPropertyIPv4MaxOffloadSizeSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4MaxOffloadSizeSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4MinSegmentCountSupported sets the value of IPv4MinSegmentCountSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) SetPropertyIPv4MinSegmentCountSupported(value uint32) (err error) {
	return instance.SetProperty("IPv4MinSegmentCountSupported", value)
}

// GetIPv4MinSegmentCountSupported gets the value of IPv4MinSegmentCountSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) GetPropertyIPv4MinSegmentCountSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4MinSegmentCountSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4TcpOptionsSupported sets the value of IPv4TcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) SetPropertyIPv4TcpOptionsSupported(value bool) (err error) {
	return instance.SetProperty("IPv4TcpOptionsSupported", value)
}

// GetIPv4TcpOptionsSupported gets the value of IPv4TcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) GetPropertyIPv4TcpOptionsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4TcpOptionsSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
