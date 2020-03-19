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

// MSFT_NetNatTransitionMonitoring struct
type MSFT_NetNatTransitionMonitoring struct {
	*MSFT_NetSettingData

	//
	InboundAddress string

	//
	NatOutboundAddress string

	//
	OutboundAddress string

	//
	TransportProtocol uint32
}

func NewMSFT_NetNatTransitionMonitoringEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetNatTransitionMonitoring, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatTransitionMonitoring{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetNatTransitionMonitoringEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetNatTransitionMonitoring, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNatTransitionMonitoring{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetInboundAddress sets the value of InboundAddress for the instance
func (instance *MSFT_NetNatTransitionMonitoring) SetPropertyInboundAddress(value string) (err error) {
	return instance.SetProperty("InboundAddress", value)
}

// GetInboundAddress gets the value of InboundAddress for the instance
func (instance *MSFT_NetNatTransitionMonitoring) GetPropertyInboundAddress() (value string, err error) {
	retValue, err := instance.GetProperty("InboundAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNatOutboundAddress sets the value of NatOutboundAddress for the instance
func (instance *MSFT_NetNatTransitionMonitoring) SetPropertyNatOutboundAddress(value string) (err error) {
	return instance.SetProperty("NatOutboundAddress", value)
}

// GetNatOutboundAddress gets the value of NatOutboundAddress for the instance
func (instance *MSFT_NetNatTransitionMonitoring) GetPropertyNatOutboundAddress() (value string, err error) {
	retValue, err := instance.GetProperty("NatOutboundAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutboundAddress sets the value of OutboundAddress for the instance
func (instance *MSFT_NetNatTransitionMonitoring) SetPropertyOutboundAddress(value string) (err error) {
	return instance.SetProperty("OutboundAddress", value)
}

// GetOutboundAddress gets the value of OutboundAddress for the instance
func (instance *MSFT_NetNatTransitionMonitoring) GetPropertyOutboundAddress() (value string, err error) {
	retValue, err := instance.GetProperty("OutboundAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTransportProtocol sets the value of TransportProtocol for the instance
func (instance *MSFT_NetNatTransitionMonitoring) SetPropertyTransportProtocol(value uint32) (err error) {
	return instance.SetProperty("TransportProtocol", value)
}

// GetTransportProtocol gets the value of TransportProtocol for the instance
func (instance *MSFT_NetNatTransitionMonitoring) GetPropertyTransportProtocol() (value uint32, err error) {
	retValue, err := instance.GetProperty("TransportProtocol")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
