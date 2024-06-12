// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS struct {
	*Win32_PerfFormattedData

	//
	BytesReceived uint64

	//
	BytesSent uint64

	//
	BytesTotal uint64

	//
	KBytesReceivedPerSec uint32

	//
	KBytesSentPerSec uint32

	//
	KBytesTotalPerSec uint32

	//
	PacketsReceived uint64

	//
	packetsreceiveddiscarded uint64

	//
	PacketsReceivedPerSec uint32

	//
	PacketsSent uint64

	//
	PacketsSentPerSec uint32

	//
	PacketsTotal uint64

	//
	PacketsTotalPerSec uint32

	//
	RcvPauseDuration uint64

	//
	RcvPauseFrames uint64

	//
	SentPauseDuration uint64

	//
	SentPauseFrames uint64
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoSEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyBytesReceived(value uint64) (err error) {
	return instance.SetProperty("BytesReceived", (value))
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyBytesReceived() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyBytesSent(value uint64) (err error) {
	return instance.SetProperty("BytesSent", (value))
}

// GetBytesSent gets the value of BytesSent for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyBytesSent() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyBytesTotal(value uint64) (err error) {
	return instance.SetProperty("BytesTotal", (value))
}

// GetBytesTotal gets the value of BytesTotal for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyBytesTotal() (value uint64, err error) {
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

// SetKBytesReceivedPerSec sets the value of KBytesReceivedPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyKBytesReceivedPerSec(value uint32) (err error) {
	return instance.SetProperty("KBytesReceivedPerSec", (value))
}

// GetKBytesReceivedPerSec gets the value of KBytesReceivedPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyKBytesReceivedPerSec() (value uint32, err error) {
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyKBytesSentPerSec(value uint32) (err error) {
	return instance.SetProperty("KBytesSentPerSec", (value))
}

// GetKBytesSentPerSec gets the value of KBytesSentPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyKBytesSentPerSec() (value uint32, err error) {
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyKBytesTotalPerSec(value uint32) (err error) {
	return instance.SetProperty("KBytesTotalPerSec", (value))
}

// GetKBytesTotalPerSec gets the value of KBytesTotalPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyKBytesTotalPerSec() (value uint32, err error) {
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

// SetPacketsReceived sets the value of PacketsReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyPacketsReceived(value uint64) (err error) {
	return instance.SetProperty("PacketsReceived", (value))
}

// GetPacketsReceived gets the value of PacketsReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyPacketsReceived() (value uint64, err error) {
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

// Setpacketsreceiveddiscarded sets the value of packetsreceiveddiscarded for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertypacketsreceiveddiscarded(value uint64) (err error) {
	return instance.SetProperty("packetsreceiveddiscarded", (value))
}

// Getpacketsreceiveddiscarded gets the value of packetsreceiveddiscarded for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertypacketsreceiveddiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("packetsreceiveddiscarded")
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyPacketsReceivedPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsReceivedPerSec", (value))
}

// GetPacketsReceivedPerSec gets the value of PacketsReceivedPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyPacketsReceivedPerSec() (value uint32, err error) {
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

// SetPacketsSent sets the value of PacketsSent for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyPacketsSent(value uint64) (err error) {
	return instance.SetProperty("PacketsSent", (value))
}

// GetPacketsSent gets the value of PacketsSent for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyPacketsSent() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyPacketsSentPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsSentPerSec", (value))
}

// GetPacketsSentPerSec gets the value of PacketsSentPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyPacketsSentPerSec() (value uint32, err error) {
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyPacketsTotal(value uint64) (err error) {
	return instance.SetProperty("PacketsTotal", (value))
}

// GetPacketsTotal gets the value of PacketsTotal for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyPacketsTotal() (value uint64, err error) {
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyPacketsTotalPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsTotalPerSec", (value))
}

// GetPacketsTotalPerSec gets the value of PacketsTotalPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyPacketsTotalPerSec() (value uint32, err error) {
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

// SetRcvPauseDuration sets the value of RcvPauseDuration for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyRcvPauseDuration(value uint64) (err error) {
	return instance.SetProperty("RcvPauseDuration", (value))
}

// GetRcvPauseDuration gets the value of RcvPauseDuration for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyRcvPauseDuration() (value uint64, err error) {
	retValue, err := instance.GetProperty("RcvPauseDuration")
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

// SetRcvPauseFrames sets the value of RcvPauseFrames for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertyRcvPauseFrames(value uint64) (err error) {
	return instance.SetProperty("RcvPauseFrames", (value))
}

// GetRcvPauseFrames gets the value of RcvPauseFrames for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertyRcvPauseFrames() (value uint64, err error) {
	retValue, err := instance.GetProperty("RcvPauseFrames")
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

// SetSentPauseDuration sets the value of SentPauseDuration for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertySentPauseDuration(value uint64) (err error) {
	return instance.SetProperty("SentPauseDuration", (value))
}

// GetSentPauseDuration gets the value of SentPauseDuration for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertySentPauseDuration() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentPauseDuration")
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

// SetSentPauseFrames sets the value of SentPauseFrames for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) SetPropertySentPauseFrames(value uint64) (err error) {
	return instance.SetProperty("SentPauseFrames", (value))
}

// GetSentPauseFrames gets the value of SentPauseFrames for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortQoS) GetPropertySentPauseFrames() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentPauseFrames")
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
