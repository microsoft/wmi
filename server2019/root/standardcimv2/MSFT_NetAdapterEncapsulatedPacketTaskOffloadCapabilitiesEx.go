// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx struct
type MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx struct {
	*MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities

	//
	IsVxlanUDPPortConfigurable bool

	//
	VxlanUDPPortNumber uint16
}

func NewMSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesExEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx, err error) {
	tmp, err := NewMSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx{
		MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities: tmp,
	}
	return
}

func NewMSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesExEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx, err error) {
	tmp, err := NewMSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx{
		MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities: tmp,
	}
	return
}

// SetIsVxlanUDPPortConfigurable sets the value of IsVxlanUDPPortConfigurable for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx) SetPropertyIsVxlanUDPPortConfigurable(value bool) (err error) {
	return instance.SetProperty("IsVxlanUDPPortConfigurable", value)
}

// GetIsVxlanUDPPortConfigurable gets the value of IsVxlanUDPPortConfigurable for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx) GetPropertyIsVxlanUDPPortConfigurable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsVxlanUDPPortConfigurable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVxlanUDPPortNumber sets the value of VxlanUDPPortNumber for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx) SetPropertyVxlanUDPPortNumber(value uint16) (err error) {
	return instance.SetProperty("VxlanUDPPortNumber", value)
}

// GetVxlanUDPPortNumber gets the value of VxlanUDPPortNumber for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx) GetPropertyVxlanUDPPortNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("VxlanUDPPortNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
