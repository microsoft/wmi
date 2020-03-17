// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters struct
type Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters struct {
	Win32_PerfFormattedData

	//
	BytesReceived uint64

	//
	BytesSent uint64

	//
	BytesTotal uint64

	//
	ControlPackets uint64

	//
	KBytesReceivedPerSec uint32

	//
	KBytesSentPerSec uint32

	//
	KBytesTotalPerSec uint32

	//
	PacketsOutboundDiscarded uint64

	//
	PacketsOutboundErrors uint64

	//
	PacketsReceived uint64

	//
	PacketsReceivedBadCRCError uint64

	//
	PacketsReceivedDiscarded uint64

	//
	PacketsReceivedErrors uint64

	//
	PacketsReceivedFrameLengthError uint64

	//
	PacketsReceivedPerSec uint32

	//
	PacketsReceivedSymbolError uint64

	//
	PacketsSent uint64

	//
	PacketsSentPerSec uint32

	//
	PacketsTotal uint64

	//
	PacketsTotalPerSec uint32
}

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyBytesReceived(value uint64) (err error) {
	return instance.SetProperty("BytesReceived", value)
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesSent sets the value of BytesSent for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyBytesSent(value uint64) (err error) {
	return instance.SetProperty("BytesSent", value)
}

// GetBytesSent gets the value of BytesSent for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyBytesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesSent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesTotal sets the value of BytesTotal for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyBytesTotal(value uint64) (err error) {
	return instance.SetProperty("BytesTotal", value)
}

// GetBytesTotal gets the value of BytesTotal for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyBytesTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetControlPackets sets the value of ControlPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyControlPackets(value uint64) (err error) {
	return instance.SetProperty("ControlPackets", value)
}

// GetControlPackets gets the value of ControlPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyControlPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ControlPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKBytesReceivedPerSec sets the value of KBytesReceivedPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyKBytesReceivedPerSec(value uint32) (err error) {
	return instance.SetProperty("KBytesReceivedPerSec", value)
}

// GetKBytesReceivedPerSec gets the value of KBytesReceivedPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyKBytesReceivedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("KBytesReceivedPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKBytesSentPerSec sets the value of KBytesSentPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyKBytesSentPerSec(value uint32) (err error) {
	return instance.SetProperty("KBytesSentPerSec", value)
}

// GetKBytesSentPerSec gets the value of KBytesSentPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyKBytesSentPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("KBytesSentPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKBytesTotalPerSec sets the value of KBytesTotalPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyKBytesTotalPerSec(value uint32) (err error) {
	return instance.SetProperty("KBytesTotalPerSec", value)
}

// GetKBytesTotalPerSec gets the value of KBytesTotalPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyKBytesTotalPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("KBytesTotalPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsOutboundDiscarded sets the value of PacketsOutboundDiscarded for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsOutboundDiscarded(value uint64) (err error) {
	return instance.SetProperty("PacketsOutboundDiscarded", value)
}

// GetPacketsOutboundDiscarded gets the value of PacketsOutboundDiscarded for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsOutboundDiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsOutboundDiscarded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsOutboundErrors sets the value of PacketsOutboundErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsOutboundErrors(value uint64) (err error) {
	return instance.SetProperty("PacketsOutboundErrors", value)
}

// GetPacketsOutboundErrors gets the value of PacketsOutboundErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsOutboundErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsOutboundErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsReceived sets the value of PacketsReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsReceived(value uint64) (err error) {
	return instance.SetProperty("PacketsReceived", value)
}

// GetPacketsReceived gets the value of PacketsReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsReceivedBadCRCError sets the value of PacketsReceivedBadCRCError for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsReceivedBadCRCError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedBadCRCError", value)
}

// GetPacketsReceivedBadCRCError gets the value of PacketsReceivedBadCRCError for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsReceivedBadCRCError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedBadCRCError")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsReceivedDiscarded sets the value of PacketsReceivedDiscarded for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsReceivedDiscarded(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedDiscarded", value)
}

// GetPacketsReceivedDiscarded gets the value of PacketsReceivedDiscarded for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsReceivedDiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedDiscarded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsReceivedErrors sets the value of PacketsReceivedErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsReceivedErrors(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedErrors", value)
}

// GetPacketsReceivedErrors gets the value of PacketsReceivedErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsReceivedErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsReceivedFrameLengthError sets the value of PacketsReceivedFrameLengthError for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsReceivedFrameLengthError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedFrameLengthError", value)
}

// GetPacketsReceivedFrameLengthError gets the value of PacketsReceivedFrameLengthError for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsReceivedFrameLengthError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedFrameLengthError")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsReceivedPerSec sets the value of PacketsReceivedPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsReceivedPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsReceivedPerSec", value)
}

// GetPacketsReceivedPerSec gets the value of PacketsReceivedPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsReceivedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsReceivedSymbolError sets the value of PacketsReceivedSymbolError for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsReceivedSymbolError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedSymbolError", value)
}

// GetPacketsReceivedSymbolError gets the value of PacketsReceivedSymbolError for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsReceivedSymbolError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedSymbolError")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsSent sets the value of PacketsSent for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsSent(value uint64) (err error) {
	return instance.SetProperty("PacketsSent", value)
}

// GetPacketsSent gets the value of PacketsSent for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsSent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsSentPerSec sets the value of PacketsSentPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsSentPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsSentPerSec", value)
}

// GetPacketsSentPerSec gets the value of PacketsSentPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsSentPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsSentPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsTotal sets the value of PacketsTotal for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsTotal(value uint64) (err error) {
	return instance.SetProperty("PacketsTotal", value)
}

// GetPacketsTotal gets the value of PacketsTotal for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsTotalPerSec sets the value of PacketsTotalPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) SetPropertyPacketsTotalPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsTotalPerSec", value)
}

// GetPacketsTotalPerSec gets the value of PacketsTotalPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx4EthProvider_MellanoxAdapterTrafficCounters) GetPropertyPacketsTotalPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsTotalPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
