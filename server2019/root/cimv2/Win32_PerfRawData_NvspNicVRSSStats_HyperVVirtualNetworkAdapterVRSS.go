// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS struct
type Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS struct {
	Win32_PerfRawData

	//
	PostmoveReceivePacketsPerSecond uint64

	//
	PostmoveReceivePacketsTotal uint64

	//
	PostmoveSendPacketCompletionsPerSecond uint64

	//
	PostmoveSendPacketCompletionsTotal uint64

	//
	PostmoveSendPacketsPerSecond uint64

	//
	PostmoveSendPacketsTotal uint64

	//
	ReceivePacketPerSecond uint64

	//
	ReceivePacketTotal uint64

	//
	ReceiveProcessor uint32

	//
	ReceiveProcessorGroup uint32

	//
	SendPacketCompletionsPerSecond uint64

	//
	SendPacketCompletionsTotal uint64

	//
	SendPacketPerSecond uint64

	//
	SendPacketTotal uint64

	//
	SendProcessor uint32

	//
	SendProcessorGroup uint32
}

// SetPostmoveReceivePacketsPerSecond sets the value of PostmoveReceivePacketsPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertyPostmoveReceivePacketsPerSecond(value uint64) (err error) {
	return instance.SetProperty("PostmoveReceivePacketsPerSecond", value)
}

// GetPostmoveReceivePacketsPerSecond gets the value of PostmoveReceivePacketsPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertyPostmoveReceivePacketsPerSecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("PostmoveReceivePacketsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPostmoveReceivePacketsTotal sets the value of PostmoveReceivePacketsTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertyPostmoveReceivePacketsTotal(value uint64) (err error) {
	return instance.SetProperty("PostmoveReceivePacketsTotal", value)
}

// GetPostmoveReceivePacketsTotal gets the value of PostmoveReceivePacketsTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertyPostmoveReceivePacketsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("PostmoveReceivePacketsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPostmoveSendPacketCompletionsPerSecond sets the value of PostmoveSendPacketCompletionsPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertyPostmoveSendPacketCompletionsPerSecond(value uint64) (err error) {
	return instance.SetProperty("PostmoveSendPacketCompletionsPerSecond", value)
}

// GetPostmoveSendPacketCompletionsPerSecond gets the value of PostmoveSendPacketCompletionsPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertyPostmoveSendPacketCompletionsPerSecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("PostmoveSendPacketCompletionsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPostmoveSendPacketCompletionsTotal sets the value of PostmoveSendPacketCompletionsTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertyPostmoveSendPacketCompletionsTotal(value uint64) (err error) {
	return instance.SetProperty("PostmoveSendPacketCompletionsTotal", value)
}

// GetPostmoveSendPacketCompletionsTotal gets the value of PostmoveSendPacketCompletionsTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertyPostmoveSendPacketCompletionsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("PostmoveSendPacketCompletionsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPostmoveSendPacketsPerSecond sets the value of PostmoveSendPacketsPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertyPostmoveSendPacketsPerSecond(value uint64) (err error) {
	return instance.SetProperty("PostmoveSendPacketsPerSecond", value)
}

// GetPostmoveSendPacketsPerSecond gets the value of PostmoveSendPacketsPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertyPostmoveSendPacketsPerSecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("PostmoveSendPacketsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPostmoveSendPacketsTotal sets the value of PostmoveSendPacketsTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertyPostmoveSendPacketsTotal(value uint64) (err error) {
	return instance.SetProperty("PostmoveSendPacketsTotal", value)
}

// GetPostmoveSendPacketsTotal gets the value of PostmoveSendPacketsTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertyPostmoveSendPacketsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("PostmoveSendPacketsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivePacketPerSecond sets the value of ReceivePacketPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertyReceivePacketPerSecond(value uint64) (err error) {
	return instance.SetProperty("ReceivePacketPerSecond", value)
}

// GetReceivePacketPerSecond gets the value of ReceivePacketPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertyReceivePacketPerSecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivePacketPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivePacketTotal sets the value of ReceivePacketTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertyReceivePacketTotal(value uint64) (err error) {
	return instance.SetProperty("ReceivePacketTotal", value)
}

// GetReceivePacketTotal gets the value of ReceivePacketTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertyReceivePacketTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivePacketTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveProcessor sets the value of ReceiveProcessor for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertyReceiveProcessor(value uint32) (err error) {
	return instance.SetProperty("ReceiveProcessor", value)
}

// GetReceiveProcessor gets the value of ReceiveProcessor for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertyReceiveProcessor() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReceiveProcessor")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveProcessorGroup sets the value of ReceiveProcessorGroup for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertyReceiveProcessorGroup(value uint32) (err error) {
	return instance.SetProperty("ReceiveProcessorGroup", value)
}

// GetReceiveProcessorGroup gets the value of ReceiveProcessorGroup for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertyReceiveProcessorGroup() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReceiveProcessorGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSendPacketCompletionsPerSecond sets the value of SendPacketCompletionsPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertySendPacketCompletionsPerSecond(value uint64) (err error) {
	return instance.SetProperty("SendPacketCompletionsPerSecond", value)
}

// GetSendPacketCompletionsPerSecond gets the value of SendPacketCompletionsPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertySendPacketCompletionsPerSecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendPacketCompletionsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSendPacketCompletionsTotal sets the value of SendPacketCompletionsTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertySendPacketCompletionsTotal(value uint64) (err error) {
	return instance.SetProperty("SendPacketCompletionsTotal", value)
}

// GetSendPacketCompletionsTotal gets the value of SendPacketCompletionsTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertySendPacketCompletionsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendPacketCompletionsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSendPacketPerSecond sets the value of SendPacketPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertySendPacketPerSecond(value uint64) (err error) {
	return instance.SetProperty("SendPacketPerSecond", value)
}

// GetSendPacketPerSecond gets the value of SendPacketPerSecond for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertySendPacketPerSecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendPacketPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSendPacketTotal sets the value of SendPacketTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertySendPacketTotal(value uint64) (err error) {
	return instance.SetProperty("SendPacketTotal", value)
}

// GetSendPacketTotal gets the value of SendPacketTotal for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertySendPacketTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendPacketTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSendProcessor sets the value of SendProcessor for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertySendProcessor(value uint32) (err error) {
	return instance.SetProperty("SendProcessor", value)
}

// GetSendProcessor gets the value of SendProcessor for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertySendProcessor() (value uint32, err error) {
	retValue, err := instance.GetProperty("SendProcessor")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSendProcessorGroup sets the value of SendProcessorGroup for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) SetPropertySendProcessorGroup(value uint32) (err error) {
	return instance.SetProperty("SendProcessorGroup", value)
}

// GetSendProcessorGroup gets the value of SendProcessorGroup for the instance
func (instance *Win32_PerfRawData_NvspNicVRSSStats_HyperVVirtualNetworkAdapterVRSS) GetPropertySendProcessorGroup() (value uint32, err error) {
	retValue, err := instance.GetProperty("SendProcessorGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
