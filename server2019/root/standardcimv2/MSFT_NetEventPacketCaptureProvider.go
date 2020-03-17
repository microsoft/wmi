// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetEventPacketCaptureProvider struct
type MSFT_NetEventPacketCaptureProvider struct {
	MSFT_NetEventProviderBase

	//
	CaptureType uint8

	//
	EtherType []uint16

	//
	IPAddresses []string

	//
	IPProtocols []uint8

	//
	LinkLayerAddress []string

	//
	MultiLayer bool

	//
	TruncationLength uint16

	//
	VmCaptureDirection uint8
}

// SetCaptureType sets the value of CaptureType for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyCaptureType(value uint8) (err error) {
	return instance.SetProperty("CaptureType", value)
}

// GetCaptureType gets the value of CaptureType for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyCaptureType() (value uint8, err error) {
	retValue, err := instance.GetProperty("CaptureType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEtherType sets the value of EtherType for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyEtherType(value []uint16) (err error) {
	return instance.SetProperty("EtherType", value)
}

// GetEtherType gets the value of EtherType for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyEtherType() (value []uint16, err error) {
	retValue, err := instance.GetProperty("EtherType")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPAddresses sets the value of IPAddresses for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyIPAddresses(value []string) (err error) {
	return instance.SetProperty("IPAddresses", value)
}

// GetIPAddresses gets the value of IPAddresses for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyIPAddresses() (value []string, err error) {
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

// SetIPProtocols sets the value of IPProtocols for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyIPProtocols(value []uint8) (err error) {
	return instance.SetProperty("IPProtocols", value)
}

// GetIPProtocols gets the value of IPProtocols for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyIPProtocols() (value []uint8, err error) {
	retValue, err := instance.GetProperty("IPProtocols")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkLayerAddress sets the value of LinkLayerAddress for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyLinkLayerAddress(value []string) (err error) {
	return instance.SetProperty("LinkLayerAddress", value)
}

// GetLinkLayerAddress gets the value of LinkLayerAddress for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyLinkLayerAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("LinkLayerAddress")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMultiLayer sets the value of MultiLayer for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyMultiLayer(value bool) (err error) {
	return instance.SetProperty("MultiLayer", value)
}

// GetMultiLayer gets the value of MultiLayer for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyMultiLayer() (value bool, err error) {
	retValue, err := instance.GetProperty("MultiLayer")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTruncationLength sets the value of TruncationLength for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyTruncationLength(value uint16) (err error) {
	return instance.SetProperty("TruncationLength", value)
}

// GetTruncationLength gets the value of TruncationLength for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyTruncationLength() (value uint16, err error) {
	retValue, err := instance.GetProperty("TruncationLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmCaptureDirection sets the value of VmCaptureDirection for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyVmCaptureDirection(value uint8) (err error) {
	return instance.SetProperty("VmCaptureDirection", value)
}

// GetVmCaptureDirection gets the value of VmCaptureDirection for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyVmCaptureDirection() (value uint8, err error) {
	retValue, err := instance.GetProperty("VmCaptureDirection")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
