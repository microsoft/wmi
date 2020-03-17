// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetEventWFPCaptureProvider struct
type MSFT_NetEventWFPCaptureProvider struct {
	MSFT_NetEventProviderBase

	//
	CaptureLayerSet uint64

	//
	DiscardedEvents bool

	//
	IPAddresses []string

	//
	TCPPorts []uint16

	//
	UDPPorts []uint16
}

// SetCaptureLayerSet sets the value of CaptureLayerSet for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) SetPropertyCaptureLayerSet(value uint64) (err error) {
	return instance.SetProperty("CaptureLayerSet", value)
}

// GetCaptureLayerSet gets the value of CaptureLayerSet for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) GetPropertyCaptureLayerSet() (value uint64, err error) {
	retValue, err := instance.GetProperty("CaptureLayerSet")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiscardedEvents sets the value of DiscardedEvents for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) SetPropertyDiscardedEvents(value bool) (err error) {
	return instance.SetProperty("DiscardedEvents", value)
}

// GetDiscardedEvents gets the value of DiscardedEvents for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) GetPropertyDiscardedEvents() (value bool, err error) {
	retValue, err := instance.GetProperty("DiscardedEvents")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPAddresses sets the value of IPAddresses for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) SetPropertyIPAddresses(value []string) (err error) {
	return instance.SetProperty("IPAddresses", value)
}

// GetIPAddresses gets the value of IPAddresses for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) GetPropertyIPAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTCPPorts sets the value of TCPPorts for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) SetPropertyTCPPorts(value []uint16) (err error) {
	return instance.SetProperty("TCPPorts", value)
}

// GetTCPPorts gets the value of TCPPorts for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) GetPropertyTCPPorts() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TCPPorts")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUDPPorts sets the value of UDPPorts for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) SetPropertyUDPPorts(value []uint16) (err error) {
	return instance.SetProperty("UDPPorts", value)
}

// GetUDPPorts gets the value of UDPPorts for the instance
func (instance *MSFT_NetEventWFPCaptureProvider) GetPropertyUDPPorts() (value []uint16, err error) {
	retValue, err := instance.GetProperty("UDPPorts")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
