// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterRscCapabilities struct
type MSFT_NetAdapterRscCapabilities struct {
	*cim.WmiInstance

	//
	IPv4Supported bool

	//
	IPv6Supported bool
}

func NewMSFT_NetAdapterRscCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterRscCapabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterRscCapabilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapterRscCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterRscCapabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterRscCapabilities{
		WmiInstance: tmp,
	}
	return
}

// SetIPv4Supported sets the value of IPv4Supported for the instance
func (instance *MSFT_NetAdapterRscCapabilities) SetPropertyIPv4Supported(value bool) (err error) {
	return instance.SetProperty("IPv4Supported", (value))
}

// GetIPv4Supported gets the value of IPv4Supported for the instance
func (instance *MSFT_NetAdapterRscCapabilities) GetPropertyIPv4Supported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4Supported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetIPv6Supported sets the value of IPv6Supported for the instance
func (instance *MSFT_NetAdapterRscCapabilities) SetPropertyIPv6Supported(value bool) (err error) {
	return instance.SetProperty("IPv6Supported", (value))
}

// GetIPv6Supported gets the value of IPv6Supported for the instance
func (instance *MSFT_NetAdapterRscCapabilities) GetPropertyIPv6Supported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6Supported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
