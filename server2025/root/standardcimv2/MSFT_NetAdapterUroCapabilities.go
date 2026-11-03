// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
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

// MSFT_NetAdapterUroCapabilities struct
type MSFT_NetAdapterUroCapabilities struct {
	*cim.WmiInstance

	//
	Supported bool
}

func NewMSFT_NetAdapterUroCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterUroCapabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterUroCapabilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapterUroCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterUroCapabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterUroCapabilities{
		WmiInstance: tmp,
	}
	return
}

// SetSupported sets the value of Supported for the instance
func (instance *MSFT_NetAdapterUroCapabilities) SetPropertySupported(value bool) (err error) {
	return instance.SetProperty("Supported", (value))
}

// GetSupported gets the value of Supported for the instance
func (instance *MSFT_NetAdapterUroCapabilities) GetPropertySupported() (value bool, err error) {
	retValue, err := instance.GetProperty("Supported")
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
