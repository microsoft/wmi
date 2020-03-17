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

// MSFT_NetAdapterRscCapabilities struct
type MSFT_NetAdapterRscCapabilities struct {
	cim.WmiInstance

	//
	IPv4Supported bool

	//
	IPv6Supported bool
}

// SetIPv4Supported sets the value of IPv4Supported for the instance
func (instance *MSFT_NetAdapterRscCapabilities) SetPropertyIPv4Supported(value bool) (err error) {
	return instance.SetProperty("IPv4Supported", value)
}

// GetIPv4Supported gets the value of IPv4Supported for the instance
func (instance *MSFT_NetAdapterRscCapabilities) GetPropertyIPv4Supported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4Supported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6Supported sets the value of IPv6Supported for the instance
func (instance *MSFT_NetAdapterRscCapabilities) SetPropertyIPv6Supported(value bool) (err error) {
	return instance.SetProperty("IPv6Supported", value)
}

// GetIPv6Supported gets the value of IPv6Supported for the instance
func (instance *MSFT_NetAdapterRscCapabilities) GetPropertyIPv6Supported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6Supported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
