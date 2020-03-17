// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetNatTransitionMonitoring struct
type MSFT_NetNatTransitionMonitoring struct {
	MSFT_NetSettingData

	//
	InboundAddress string

	//
	NatOutboundAddress string

	//
	OutboundAddress string

	//
	TransportProtocol uint32
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
