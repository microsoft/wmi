// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter struct
type Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter struct {
	Win32_PerfFormattedData

	//
	BytesDropped uint64

	//
	BytesReceivedPersec uint64

	//
	BytesSentPersec uint64

	//
	FramesDropped uint64

	//
	FramesReceivedPersec uint64

	//
	FramesSentPersec uint64
}

// SetBytesDropped sets the value of BytesDropped for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) SetPropertyBytesDropped(value uint64) (err error) {
	return instance.SetProperty("BytesDropped", value)
}

// GetBytesDropped gets the value of BytesDropped for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) GetPropertyBytesDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesReceivedPersec sets the value of BytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) SetPropertyBytesReceivedPersec(value uint64) (err error) {
	return instance.SetProperty("BytesReceivedPersec", value)
}

// GetBytesReceivedPersec gets the value of BytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) GetPropertyBytesReceivedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceivedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesSentPersec sets the value of BytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) SetPropertyBytesSentPersec(value uint64) (err error) {
	return instance.SetProperty("BytesSentPersec", value)
}

// GetBytesSentPersec gets the value of BytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) GetPropertyBytesSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesSentPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFramesDropped sets the value of FramesDropped for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) SetPropertyFramesDropped(value uint64) (err error) {
	return instance.SetProperty("FramesDropped", value)
}

// GetFramesDropped gets the value of FramesDropped for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) GetPropertyFramesDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("FramesDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFramesReceivedPersec sets the value of FramesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) SetPropertyFramesReceivedPersec(value uint64) (err error) {
	return instance.SetProperty("FramesReceivedPersec", value)
}

// GetFramesReceivedPersec gets the value of FramesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) GetPropertyFramesReceivedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FramesReceivedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFramesSentPersec sets the value of FramesSentPersec for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) SetPropertyFramesSentPersec(value uint64) (err error) {
	return instance.SetProperty("FramesSentPersec", value)
}

// GetFramesSentPersec gets the value of FramesSentPersec for the instance
func (instance *Win32_PerfFormattedData_EthernetPerfProvider_HyperVLegacyNetworkAdapter) GetPropertyFramesSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FramesSentPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
