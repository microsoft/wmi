// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic struct
type Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic struct {
	*Win32_PerfRawData

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
	PacketsReceivedDiscardedNoReceiveWQEs uint64

	//
	PacketsReceivedErrors uint64

	//
	PacketsReceivedFragmentsError uint64

	//
	PacketsReceivedFrameLengthError uint64

	//
	PacketsReceivedFrameTooLongError uint64

	//
	PacketsReceivedFrameUndersizeError uint64

	//
	PacketsReceivedJabbersError uint64

	//
	PacketsReceivedPerSec uint32

	//
	PacketsReceivedSymbolError uint64

	//
	PacketsReceivedUnsupportedOpcodeError uint64

	//
	PacketsSent uint64

	//
	PacketsSentPerSec uint32

	//
	PacketsTotal uint64

	//
	PacketsTotalPerSec uint32

	//
	RSCAborts uint64

	//
	RSCAveragePacketSize uint64

	//
	RSCCoalescedEvents uint64

	//
	RSCCoalescedOctets uint64

	//
	RSCCoalescedPackets uint64
}

func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTrafficEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTrafficEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyBytesReceived(value uint64) (err error) {
	return instance.SetProperty("BytesReceived", (value))
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetBytesSent sets the value of BytesSent for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyBytesSent(value uint64) (err error) {
	return instance.SetProperty("BytesSent", (value))
}

// GetBytesSent gets the value of BytesSent for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyBytesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesSent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetBytesTotal sets the value of BytesTotal for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyBytesTotal(value uint64) (err error) {
	return instance.SetProperty("BytesTotal", (value))
}

// GetBytesTotal gets the value of BytesTotal for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyBytesTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesTotal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetControlPackets sets the value of ControlPackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyControlPackets(value uint64) (err error) {
	return instance.SetProperty("ControlPackets", (value))
}

// GetControlPackets gets the value of ControlPackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyControlPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ControlPackets")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetKBytesReceivedPerSec sets the value of KBytesReceivedPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyKBytesReceivedPerSec(value uint32) (err error) {
	return instance.SetProperty("KBytesReceivedPerSec", (value))
}

// GetKBytesReceivedPerSec gets the value of KBytesReceivedPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyKBytesReceivedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("KBytesReceivedPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetKBytesSentPerSec sets the value of KBytesSentPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyKBytesSentPerSec(value uint32) (err error) {
	return instance.SetProperty("KBytesSentPerSec", (value))
}

// GetKBytesSentPerSec gets the value of KBytesSentPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyKBytesSentPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("KBytesSentPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetKBytesTotalPerSec sets the value of KBytesTotalPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyKBytesTotalPerSec(value uint32) (err error) {
	return instance.SetProperty("KBytesTotalPerSec", (value))
}

// GetKBytesTotalPerSec gets the value of KBytesTotalPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyKBytesTotalPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("KBytesTotalPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPacketsOutboundDiscarded sets the value of PacketsOutboundDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsOutboundDiscarded(value uint64) (err error) {
	return instance.SetProperty("PacketsOutboundDiscarded", (value))
}

// GetPacketsOutboundDiscarded gets the value of PacketsOutboundDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsOutboundDiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsOutboundDiscarded")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsOutboundErrors sets the value of PacketsOutboundErrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsOutboundErrors(value uint64) (err error) {
	return instance.SetProperty("PacketsOutboundErrors", (value))
}

// GetPacketsOutboundErrors gets the value of PacketsOutboundErrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsOutboundErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsOutboundErrors")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceived sets the value of PacketsReceived for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceived(value uint64) (err error) {
	return instance.SetProperty("PacketsReceived", (value))
}

// GetPacketsReceived gets the value of PacketsReceived for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceived")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceivedBadCRCError sets the value of PacketsReceivedBadCRCError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedBadCRCError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedBadCRCError", (value))
}

// GetPacketsReceivedBadCRCError gets the value of PacketsReceivedBadCRCError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedBadCRCError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedBadCRCError")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceivedDiscardedNoReceiveWQEs sets the value of PacketsReceivedDiscardedNoReceiveWQEs for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedDiscardedNoReceiveWQEs(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedDiscardedNoReceiveWQEs", (value))
}

// GetPacketsReceivedDiscardedNoReceiveWQEs gets the value of PacketsReceivedDiscardedNoReceiveWQEs for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedDiscardedNoReceiveWQEs() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedDiscardedNoReceiveWQEs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceivedErrors sets the value of PacketsReceivedErrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedErrors(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedErrors", (value))
}

// GetPacketsReceivedErrors gets the value of PacketsReceivedErrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedErrors")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceivedFragmentsError sets the value of PacketsReceivedFragmentsError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedFragmentsError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedFragmentsError", (value))
}

// GetPacketsReceivedFragmentsError gets the value of PacketsReceivedFragmentsError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedFragmentsError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedFragmentsError")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceivedFrameLengthError sets the value of PacketsReceivedFrameLengthError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedFrameLengthError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedFrameLengthError", (value))
}

// GetPacketsReceivedFrameLengthError gets the value of PacketsReceivedFrameLengthError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedFrameLengthError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedFrameLengthError")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceivedFrameTooLongError sets the value of PacketsReceivedFrameTooLongError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedFrameTooLongError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedFrameTooLongError", (value))
}

// GetPacketsReceivedFrameTooLongError gets the value of PacketsReceivedFrameTooLongError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedFrameTooLongError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedFrameTooLongError")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceivedFrameUndersizeError sets the value of PacketsReceivedFrameUndersizeError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedFrameUndersizeError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedFrameUndersizeError", (value))
}

// GetPacketsReceivedFrameUndersizeError gets the value of PacketsReceivedFrameUndersizeError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedFrameUndersizeError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedFrameUndersizeError")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceivedJabbersError sets the value of PacketsReceivedJabbersError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedJabbersError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedJabbersError", (value))
}

// GetPacketsReceivedJabbersError gets the value of PacketsReceivedJabbersError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedJabbersError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedJabbersError")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceivedPerSec sets the value of PacketsReceivedPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsReceivedPerSec", (value))
}

// GetPacketsReceivedPerSec gets the value of PacketsReceivedPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPacketsReceivedSymbolError sets the value of PacketsReceivedSymbolError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedSymbolError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedSymbolError", (value))
}

// GetPacketsReceivedSymbolError gets the value of PacketsReceivedSymbolError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedSymbolError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedSymbolError")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsReceivedUnsupportedOpcodeError sets the value of PacketsReceivedUnsupportedOpcodeError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsReceivedUnsupportedOpcodeError(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedUnsupportedOpcodeError", (value))
}

// GetPacketsReceivedUnsupportedOpcodeError gets the value of PacketsReceivedUnsupportedOpcodeError for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsReceivedUnsupportedOpcodeError() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedUnsupportedOpcodeError")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsSent sets the value of PacketsSent for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsSent(value uint64) (err error) {
	return instance.SetProperty("PacketsSent", (value))
}

// GetPacketsSent gets the value of PacketsSent for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsSent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsSentPerSec sets the value of PacketsSentPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsSentPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsSentPerSec", (value))
}

// GetPacketsSentPerSec gets the value of PacketsSentPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsSentPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsSentPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPacketsTotal sets the value of PacketsTotal for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsTotal(value uint64) (err error) {
	return instance.SetProperty("PacketsTotal", (value))
}

// GetPacketsTotal gets the value of PacketsTotal for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsTotal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsTotalPerSec sets the value of PacketsTotalPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyPacketsTotalPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsTotalPerSec", (value))
}

// GetPacketsTotalPerSec gets the value of PacketsTotalPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyPacketsTotalPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsTotalPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetRSCAborts sets the value of RSCAborts for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyRSCAborts(value uint64) (err error) {
	return instance.SetProperty("RSCAborts", (value))
}

// GetRSCAborts gets the value of RSCAborts for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyRSCAborts() (value uint64, err error) {
	retValue, err := instance.GetProperty("RSCAborts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRSCAveragePacketSize sets the value of RSCAveragePacketSize for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyRSCAveragePacketSize(value uint64) (err error) {
	return instance.SetProperty("RSCAveragePacketSize", (value))
}

// GetRSCAveragePacketSize gets the value of RSCAveragePacketSize for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyRSCAveragePacketSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("RSCAveragePacketSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRSCCoalescedEvents sets the value of RSCCoalescedEvents for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyRSCCoalescedEvents(value uint64) (err error) {
	return instance.SetProperty("RSCCoalescedEvents", (value))
}

// GetRSCCoalescedEvents gets the value of RSCCoalescedEvents for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyRSCCoalescedEvents() (value uint64, err error) {
	retValue, err := instance.GetProperty("RSCCoalescedEvents")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRSCCoalescedOctets sets the value of RSCCoalescedOctets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyRSCCoalescedOctets(value uint64) (err error) {
	return instance.SetProperty("RSCCoalescedOctets", (value))
}

// GetRSCCoalescedOctets gets the value of RSCCoalescedOctets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyRSCCoalescedOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("RSCCoalescedOctets")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRSCCoalescedPackets sets the value of RSCCoalescedPackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) SetPropertyRSCCoalescedPackets(value uint64) (err error) {
	return instance.SetProperty("RSCCoalescedPackets", (value))
}

// GetRSCCoalescedPackets gets the value of RSCCoalescedPackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PortTraffic) GetPropertyRSCCoalescedPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("RSCCoalescedPackets")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
