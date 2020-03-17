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

// MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities struct
type MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities struct {
	cim.WmiInstance

	//
	LsoV2Supported uint32

	//
	ReceiveChecksumOffloadSupported uint32

	//
	RssSupported uint32

	//
	TransmitChecksumOffloadSupported uint32

	//
	VmqSupported uint32
}

// SetLsoV2Supported sets the value of LsoV2Supported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) SetPropertyLsoV2Supported(value uint32) (err error) {
	return instance.SetProperty("LsoV2Supported", value)
}

// GetLsoV2Supported gets the value of LsoV2Supported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) GetPropertyLsoV2Supported() (value uint32, err error) {
	retValue, err := instance.GetProperty("LsoV2Supported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveChecksumOffloadSupported sets the value of ReceiveChecksumOffloadSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) SetPropertyReceiveChecksumOffloadSupported(value uint32) (err error) {
	return instance.SetProperty("ReceiveChecksumOffloadSupported", value)
}

// GetReceiveChecksumOffloadSupported gets the value of ReceiveChecksumOffloadSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) GetPropertyReceiveChecksumOffloadSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReceiveChecksumOffloadSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRssSupported sets the value of RssSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) SetPropertyRssSupported(value uint32) (err error) {
	return instance.SetProperty("RssSupported", value)
}

// GetRssSupported gets the value of RssSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) GetPropertyRssSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("RssSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTransmitChecksumOffloadSupported sets the value of TransmitChecksumOffloadSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) SetPropertyTransmitChecksumOffloadSupported(value uint32) (err error) {
	return instance.SetProperty("TransmitChecksumOffloadSupported", value)
}

// GetTransmitChecksumOffloadSupported gets the value of TransmitChecksumOffloadSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) GetPropertyTransmitChecksumOffloadSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("TransmitChecksumOffloadSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmqSupported sets the value of VmqSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) SetPropertyVmqSupported(value uint32) (err error) {
	return instance.SetProperty("VmqSupported", value)
}

// GetVmqSupported gets the value of VmqSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) GetPropertyVmqSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("VmqSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
