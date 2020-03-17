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

// MSFT_NetAdapterUdpSegmentationOffloadCapabilities struct
type MSFT_NetAdapterUdpSegmentationOffloadCapabilities struct {
	cim.WmiInstance

	//
	IPv4Encapsulation MSFT_NetAdapterUsoEncapsulationTypes

	//
	IPv4MaxOffloadSizeSupported uint32

	//
	IPv4MinSegmentCountSupported uint32

	//
	IPv4SubMssFinalSegmentSupported bool

	//
	IPv6Encapsulation MSFT_NetAdapterUsoEncapsulationTypes

	//
	IPv6IpExtensionHeadersSupported bool

	//
	IPv6MaxOffLoadSizeSupported uint32

	//
	IPv6MinSegmentCountSupported uint32

	//
	IPv6SubMssFinalSegmentSupported bool
}

// SetIPv4Encapsulation sets the value of IPv4Encapsulation for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv4Encapsulation(value MSFT_NetAdapterUsoEncapsulationTypes) (err error) {
	return instance.SetProperty("IPv4Encapsulation", value)
}

// GetIPv4Encapsulation gets the value of IPv4Encapsulation for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv4Encapsulation() (value MSFT_NetAdapterUsoEncapsulationTypes, err error) {
	retValue, err := instance.GetProperty("IPv4Encapsulation")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapterUsoEncapsulationTypes)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4MaxOffloadSizeSupported sets the value of IPv4MaxOffloadSizeSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv4MaxOffloadSizeSupported(value uint32) (err error) {
	return instance.SetProperty("IPv4MaxOffloadSizeSupported", value)
}

// GetIPv4MaxOffloadSizeSupported gets the value of IPv4MaxOffloadSizeSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv4MaxOffloadSizeSupported() (value uint32, err error) {
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
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv4MinSegmentCountSupported(value uint32) (err error) {
	return instance.SetProperty("IPv4MinSegmentCountSupported", value)
}

// GetIPv4MinSegmentCountSupported gets the value of IPv4MinSegmentCountSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv4MinSegmentCountSupported() (value uint32, err error) {
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

// SetIPv4SubMssFinalSegmentSupported sets the value of IPv4SubMssFinalSegmentSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv4SubMssFinalSegmentSupported(value bool) (err error) {
	return instance.SetProperty("IPv4SubMssFinalSegmentSupported", value)
}

// GetIPv4SubMssFinalSegmentSupported gets the value of IPv4SubMssFinalSegmentSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv4SubMssFinalSegmentSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4SubMssFinalSegmentSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6Encapsulation sets the value of IPv6Encapsulation for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv6Encapsulation(value MSFT_NetAdapterUsoEncapsulationTypes) (err error) {
	return instance.SetProperty("IPv6Encapsulation", value)
}

// GetIPv6Encapsulation gets the value of IPv6Encapsulation for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv6Encapsulation() (value MSFT_NetAdapterUsoEncapsulationTypes, err error) {
	retValue, err := instance.GetProperty("IPv6Encapsulation")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapterUsoEncapsulationTypes)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6IpExtensionHeadersSupported sets the value of IPv6IpExtensionHeadersSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv6IpExtensionHeadersSupported(value bool) (err error) {
	return instance.SetProperty("IPv6IpExtensionHeadersSupported", value)
}

// GetIPv6IpExtensionHeadersSupported gets the value of IPv6IpExtensionHeadersSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv6IpExtensionHeadersSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6IpExtensionHeadersSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6MaxOffLoadSizeSupported sets the value of IPv6MaxOffLoadSizeSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv6MaxOffLoadSizeSupported(value uint32) (err error) {
	return instance.SetProperty("IPv6MaxOffLoadSizeSupported", value)
}

// GetIPv6MaxOffLoadSizeSupported gets the value of IPv6MaxOffLoadSizeSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv6MaxOffLoadSizeSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6MaxOffLoadSizeSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6MinSegmentCountSupported sets the value of IPv6MinSegmentCountSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv6MinSegmentCountSupported(value uint32) (err error) {
	return instance.SetProperty("IPv6MinSegmentCountSupported", value)
}

// GetIPv6MinSegmentCountSupported gets the value of IPv6MinSegmentCountSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv6MinSegmentCountSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6MinSegmentCountSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6SubMssFinalSegmentSupported sets the value of IPv6SubMssFinalSegmentSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv6SubMssFinalSegmentSupported(value bool) (err error) {
	return instance.SetProperty("IPv6SubMssFinalSegmentSupported", value)
}

// GetIPv6SubMssFinalSegmentSupported gets the value of IPv6SubMssFinalSegmentSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv6SubMssFinalSegmentSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6SubMssFinalSegmentSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
