// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic struct
type Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic struct {
	*Win32_PerfRawData

	//
	AllowedEthTypeAntiSpoofingBytesDiscarded uint64

	//
	AllowedEthTypeAntiSpoofingPacketsDiscarded uint64

	//
	BytesReceivedBroadcastPerSec uint32

	//
	BytesReceivedMulticastPerSec uint32

	//
	BytesReceivedPerSec uint32

	//
	BytesReceivedUnicastPerSec uint32

	//
	BytesSentBroadcastPerSec uint32

	//
	BytesSentMulticastPerSec uint32

	//
	BytesSentPerSec uint32

	//
	BytesSentUnicastPerSec uint32

	//
	MacAntiSpoofingBytesDiscarded uint64

	//
	MacAntiSpoofingPacketsDiscarded uint64

	//
	PacketsOutboundDiscarded uint64

	//
	PacketsOutboundErrors uint64

	//
	PacketsReceivedBroadcastPerSec uint32

	//
	PacketsReceivedDiscarded uint64

	//
	PacketsReceivedErrors uint64

	//
	PacketsReceivedMulticastPerSec uint32

	//
	PacketsReceivedUnicastPerSec uint32

	//
	PacketsSentBroadcastPerSec uint32

	//
	PacketsSentMulticastPerSec uint32

	//
	PacketsSentUnicastPerSec uint32

	//
	VlanAntiSpoofingBytesDiscarded uint64

	//
	VlanAntiSpoofingPacketsDiscarded uint64
}

func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTrafficEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTrafficEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAllowedEthTypeAntiSpoofingBytesDiscarded sets the value of AllowedEthTypeAntiSpoofingBytesDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyAllowedEthTypeAntiSpoofingBytesDiscarded(value uint64) (err error) {
	return instance.SetProperty("AllowedEthTypeAntiSpoofingBytesDiscarded", (value))
}

// GetAllowedEthTypeAntiSpoofingBytesDiscarded gets the value of AllowedEthTypeAntiSpoofingBytesDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyAllowedEthTypeAntiSpoofingBytesDiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllowedEthTypeAntiSpoofingBytesDiscarded")
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

// SetAllowedEthTypeAntiSpoofingPacketsDiscarded sets the value of AllowedEthTypeAntiSpoofingPacketsDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyAllowedEthTypeAntiSpoofingPacketsDiscarded(value uint64) (err error) {
	return instance.SetProperty("AllowedEthTypeAntiSpoofingPacketsDiscarded", (value))
}

// GetAllowedEthTypeAntiSpoofingPacketsDiscarded gets the value of AllowedEthTypeAntiSpoofingPacketsDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyAllowedEthTypeAntiSpoofingPacketsDiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllowedEthTypeAntiSpoofingPacketsDiscarded")
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

// SetBytesReceivedBroadcastPerSec sets the value of BytesReceivedBroadcastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyBytesReceivedBroadcastPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesReceivedBroadcastPerSec", (value))
}

// GetBytesReceivedBroadcastPerSec gets the value of BytesReceivedBroadcastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyBytesReceivedBroadcastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesReceivedBroadcastPerSec")
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

// SetBytesReceivedMulticastPerSec sets the value of BytesReceivedMulticastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyBytesReceivedMulticastPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesReceivedMulticastPerSec", (value))
}

// GetBytesReceivedMulticastPerSec gets the value of BytesReceivedMulticastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyBytesReceivedMulticastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesReceivedMulticastPerSec")
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

// SetBytesReceivedPerSec sets the value of BytesReceivedPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyBytesReceivedPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesReceivedPerSec", (value))
}

// GetBytesReceivedPerSec gets the value of BytesReceivedPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyBytesReceivedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesReceivedPerSec")
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

// SetBytesReceivedUnicastPerSec sets the value of BytesReceivedUnicastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyBytesReceivedUnicastPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesReceivedUnicastPerSec", (value))
}

// GetBytesReceivedUnicastPerSec gets the value of BytesReceivedUnicastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyBytesReceivedUnicastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesReceivedUnicastPerSec")
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

// SetBytesSentBroadcastPerSec sets the value of BytesSentBroadcastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyBytesSentBroadcastPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesSentBroadcastPerSec", (value))
}

// GetBytesSentBroadcastPerSec gets the value of BytesSentBroadcastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyBytesSentBroadcastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesSentBroadcastPerSec")
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

// SetBytesSentMulticastPerSec sets the value of BytesSentMulticastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyBytesSentMulticastPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesSentMulticastPerSec", (value))
}

// GetBytesSentMulticastPerSec gets the value of BytesSentMulticastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyBytesSentMulticastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesSentMulticastPerSec")
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

// SetBytesSentPerSec sets the value of BytesSentPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyBytesSentPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesSentPerSec", (value))
}

// GetBytesSentPerSec gets the value of BytesSentPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyBytesSentPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesSentPerSec")
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

// SetBytesSentUnicastPerSec sets the value of BytesSentUnicastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyBytesSentUnicastPerSec(value uint32) (err error) {
	return instance.SetProperty("BytesSentUnicastPerSec", (value))
}

// GetBytesSentUnicastPerSec gets the value of BytesSentUnicastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyBytesSentUnicastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesSentUnicastPerSec")
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

// SetMacAntiSpoofingBytesDiscarded sets the value of MacAntiSpoofingBytesDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyMacAntiSpoofingBytesDiscarded(value uint64) (err error) {
	return instance.SetProperty("MacAntiSpoofingBytesDiscarded", (value))
}

// GetMacAntiSpoofingBytesDiscarded gets the value of MacAntiSpoofingBytesDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyMacAntiSpoofingBytesDiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("MacAntiSpoofingBytesDiscarded")
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

// SetMacAntiSpoofingPacketsDiscarded sets the value of MacAntiSpoofingPacketsDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyMacAntiSpoofingPacketsDiscarded(value uint64) (err error) {
	return instance.SetProperty("MacAntiSpoofingPacketsDiscarded", (value))
}

// GetMacAntiSpoofingPacketsDiscarded gets the value of MacAntiSpoofingPacketsDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyMacAntiSpoofingPacketsDiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("MacAntiSpoofingPacketsDiscarded")
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

// SetPacketsOutboundDiscarded sets the value of PacketsOutboundDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyPacketsOutboundDiscarded(value uint64) (err error) {
	return instance.SetProperty("PacketsOutboundDiscarded", (value))
}

// GetPacketsOutboundDiscarded gets the value of PacketsOutboundDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyPacketsOutboundDiscarded() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyPacketsOutboundErrors(value uint64) (err error) {
	return instance.SetProperty("PacketsOutboundErrors", (value))
}

// GetPacketsOutboundErrors gets the value of PacketsOutboundErrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyPacketsOutboundErrors() (value uint64, err error) {
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

// SetPacketsReceivedBroadcastPerSec sets the value of PacketsReceivedBroadcastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyPacketsReceivedBroadcastPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsReceivedBroadcastPerSec", (value))
}

// GetPacketsReceivedBroadcastPerSec gets the value of PacketsReceivedBroadcastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyPacketsReceivedBroadcastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedBroadcastPerSec")
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

// SetPacketsReceivedDiscarded sets the value of PacketsReceivedDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyPacketsReceivedDiscarded(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedDiscarded", (value))
}

// GetPacketsReceivedDiscarded gets the value of PacketsReceivedDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyPacketsReceivedDiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedDiscarded")
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
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyPacketsReceivedErrors(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedErrors", (value))
}

// GetPacketsReceivedErrors gets the value of PacketsReceivedErrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyPacketsReceivedErrors() (value uint64, err error) {
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

// SetPacketsReceivedMulticastPerSec sets the value of PacketsReceivedMulticastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyPacketsReceivedMulticastPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsReceivedMulticastPerSec", (value))
}

// GetPacketsReceivedMulticastPerSec gets the value of PacketsReceivedMulticastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyPacketsReceivedMulticastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedMulticastPerSec")
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

// SetPacketsReceivedUnicastPerSec sets the value of PacketsReceivedUnicastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyPacketsReceivedUnicastPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsReceivedUnicastPerSec", (value))
}

// GetPacketsReceivedUnicastPerSec gets the value of PacketsReceivedUnicastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyPacketsReceivedUnicastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedUnicastPerSec")
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

// SetPacketsSentBroadcastPerSec sets the value of PacketsSentBroadcastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyPacketsSentBroadcastPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsSentBroadcastPerSec", (value))
}

// GetPacketsSentBroadcastPerSec gets the value of PacketsSentBroadcastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyPacketsSentBroadcastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsSentBroadcastPerSec")
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

// SetPacketsSentMulticastPerSec sets the value of PacketsSentMulticastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyPacketsSentMulticastPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsSentMulticastPerSec", (value))
}

// GetPacketsSentMulticastPerSec gets the value of PacketsSentMulticastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyPacketsSentMulticastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsSentMulticastPerSec")
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

// SetPacketsSentUnicastPerSec sets the value of PacketsSentUnicastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyPacketsSentUnicastPerSec(value uint32) (err error) {
	return instance.SetProperty("PacketsSentUnicastPerSec", (value))
}

// GetPacketsSentUnicastPerSec gets the value of PacketsSentUnicastPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyPacketsSentUnicastPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsSentUnicastPerSec")
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

// SetVlanAntiSpoofingBytesDiscarded sets the value of VlanAntiSpoofingBytesDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyVlanAntiSpoofingBytesDiscarded(value uint64) (err error) {
	return instance.SetProperty("VlanAntiSpoofingBytesDiscarded", (value))
}

// GetVlanAntiSpoofingBytesDiscarded gets the value of VlanAntiSpoofingBytesDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyVlanAntiSpoofingBytesDiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("VlanAntiSpoofingBytesDiscarded")
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

// SetVlanAntiSpoofingPacketsDiscarded sets the value of VlanAntiSpoofingPacketsDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) SetPropertyVlanAntiSpoofingPacketsDiscarded(value uint64) (err error) {
	return instance.SetProperty("VlanAntiSpoofingPacketsDiscarded", (value))
}

// GetVlanAntiSpoofingPacketsDiscarded gets the value of VlanAntiSpoofingPacketsDiscarded for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFPortTraffic) GetPropertyVlanAntiSpoofingPacketsDiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("VlanAntiSpoofingPacketsDiscarded")
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
