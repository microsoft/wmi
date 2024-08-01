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

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters struct {
	*Win32_PerfFormattedData

	//
	EncapsulatedNonRssIPv4Only uint64

	//
	EncapsulatedNonRssIPv4PerTcp uint64

	//
	EncapsulatedNonRssIPv4PerUdp uint64

	//
	EncapsulatedNonRssIPv6Only uint64

	//
	EncapsulatedNonRssIPv6PerTcp uint64

	//
	EncapsulatedNonRssIPv6PerUdp uint64

	//
	EncapsulatedNonRssMisc uint64

	//
	EncapsulatedRssIPv4Only uint64

	//
	EncapsulatedRssIPv4PerTcp uint64

	//
	EncapsulatedRssIPv4PerUdp uint64

	//
	EncapsulatedRssIPv6Only uint64

	//
	EncapsulatedRssIPv6PerTcp uint64

	//
	EncapsulatedRssIPv6PerUdp uint64

	//
	EncapsulatedRssMisc uint64

	//
	NonRssIPv4Only uint64

	//
	NonRssIPv4PerTcp uint64

	//
	NonRssIPv4PerUdp uint64

	//
	NonRssIPv6Only uint64

	//
	NonRssIPv6PerTcp uint64

	//
	NonRssIPv6PerUdp uint64

	//
	NonRssMisc uint64

	//
	RssIPv4Only uint64

	//
	RssIPv4PerTcp uint64

	//
	RssIPv4PerUdp uint64

	//
	RssIPv6Only uint64

	//
	RssIPv6PerTcp uint64

	//
	RssIPv6PerUdp uint64

	//
	RssMisc uint64
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCountersEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCountersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetEncapsulatedNonRssIPv4Only sets the value of EncapsulatedNonRssIPv4Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedNonRssIPv4Only(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedNonRssIPv4Only", (value))
}

// GetEncapsulatedNonRssIPv4Only gets the value of EncapsulatedNonRssIPv4Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedNonRssIPv4Only() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedNonRssIPv4Only")
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

// SetEncapsulatedNonRssIPv4PerTcp sets the value of EncapsulatedNonRssIPv4PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedNonRssIPv4PerTcp(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedNonRssIPv4PerTcp", (value))
}

// GetEncapsulatedNonRssIPv4PerTcp gets the value of EncapsulatedNonRssIPv4PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedNonRssIPv4PerTcp() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedNonRssIPv4PerTcp")
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

// SetEncapsulatedNonRssIPv4PerUdp sets the value of EncapsulatedNonRssIPv4PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedNonRssIPv4PerUdp(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedNonRssIPv4PerUdp", (value))
}

// GetEncapsulatedNonRssIPv4PerUdp gets the value of EncapsulatedNonRssIPv4PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedNonRssIPv4PerUdp() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedNonRssIPv4PerUdp")
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

// SetEncapsulatedNonRssIPv6Only sets the value of EncapsulatedNonRssIPv6Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedNonRssIPv6Only(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedNonRssIPv6Only", (value))
}

// GetEncapsulatedNonRssIPv6Only gets the value of EncapsulatedNonRssIPv6Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedNonRssIPv6Only() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedNonRssIPv6Only")
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

// SetEncapsulatedNonRssIPv6PerTcp sets the value of EncapsulatedNonRssIPv6PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedNonRssIPv6PerTcp(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedNonRssIPv6PerTcp", (value))
}

// GetEncapsulatedNonRssIPv6PerTcp gets the value of EncapsulatedNonRssIPv6PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedNonRssIPv6PerTcp() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedNonRssIPv6PerTcp")
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

// SetEncapsulatedNonRssIPv6PerUdp sets the value of EncapsulatedNonRssIPv6PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedNonRssIPv6PerUdp(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedNonRssIPv6PerUdp", (value))
}

// GetEncapsulatedNonRssIPv6PerUdp gets the value of EncapsulatedNonRssIPv6PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedNonRssIPv6PerUdp() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedNonRssIPv6PerUdp")
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

// SetEncapsulatedNonRssMisc sets the value of EncapsulatedNonRssMisc for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedNonRssMisc(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedNonRssMisc", (value))
}

// GetEncapsulatedNonRssMisc gets the value of EncapsulatedNonRssMisc for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedNonRssMisc() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedNonRssMisc")
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

// SetEncapsulatedRssIPv4Only sets the value of EncapsulatedRssIPv4Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedRssIPv4Only(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedRssIPv4Only", (value))
}

// GetEncapsulatedRssIPv4Only gets the value of EncapsulatedRssIPv4Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedRssIPv4Only() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedRssIPv4Only")
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

// SetEncapsulatedRssIPv4PerTcp sets the value of EncapsulatedRssIPv4PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedRssIPv4PerTcp(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedRssIPv4PerTcp", (value))
}

// GetEncapsulatedRssIPv4PerTcp gets the value of EncapsulatedRssIPv4PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedRssIPv4PerTcp() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedRssIPv4PerTcp")
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

// SetEncapsulatedRssIPv4PerUdp sets the value of EncapsulatedRssIPv4PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedRssIPv4PerUdp(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedRssIPv4PerUdp", (value))
}

// GetEncapsulatedRssIPv4PerUdp gets the value of EncapsulatedRssIPv4PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedRssIPv4PerUdp() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedRssIPv4PerUdp")
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

// SetEncapsulatedRssIPv6Only sets the value of EncapsulatedRssIPv6Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedRssIPv6Only(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedRssIPv6Only", (value))
}

// GetEncapsulatedRssIPv6Only gets the value of EncapsulatedRssIPv6Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedRssIPv6Only() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedRssIPv6Only")
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

// SetEncapsulatedRssIPv6PerTcp sets the value of EncapsulatedRssIPv6PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedRssIPv6PerTcp(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedRssIPv6PerTcp", (value))
}

// GetEncapsulatedRssIPv6PerTcp gets the value of EncapsulatedRssIPv6PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedRssIPv6PerTcp() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedRssIPv6PerTcp")
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

// SetEncapsulatedRssIPv6PerUdp sets the value of EncapsulatedRssIPv6PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedRssIPv6PerUdp(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedRssIPv6PerUdp", (value))
}

// GetEncapsulatedRssIPv6PerUdp gets the value of EncapsulatedRssIPv6PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedRssIPv6PerUdp() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedRssIPv6PerUdp")
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

// SetEncapsulatedRssMisc sets the value of EncapsulatedRssMisc for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyEncapsulatedRssMisc(value uint64) (err error) {
	return instance.SetProperty("EncapsulatedRssMisc", (value))
}

// GetEncapsulatedRssMisc gets the value of EncapsulatedRssMisc for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyEncapsulatedRssMisc() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncapsulatedRssMisc")
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

// SetNonRssIPv4Only sets the value of NonRssIPv4Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyNonRssIPv4Only(value uint64) (err error) {
	return instance.SetProperty("NonRssIPv4Only", (value))
}

// GetNonRssIPv4Only gets the value of NonRssIPv4Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyNonRssIPv4Only() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonRssIPv4Only")
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

// SetNonRssIPv4PerTcp sets the value of NonRssIPv4PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyNonRssIPv4PerTcp(value uint64) (err error) {
	return instance.SetProperty("NonRssIPv4PerTcp", (value))
}

// GetNonRssIPv4PerTcp gets the value of NonRssIPv4PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyNonRssIPv4PerTcp() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonRssIPv4PerTcp")
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

// SetNonRssIPv4PerUdp sets the value of NonRssIPv4PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyNonRssIPv4PerUdp(value uint64) (err error) {
	return instance.SetProperty("NonRssIPv4PerUdp", (value))
}

// GetNonRssIPv4PerUdp gets the value of NonRssIPv4PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyNonRssIPv4PerUdp() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonRssIPv4PerUdp")
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

// SetNonRssIPv6Only sets the value of NonRssIPv6Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyNonRssIPv6Only(value uint64) (err error) {
	return instance.SetProperty("NonRssIPv6Only", (value))
}

// GetNonRssIPv6Only gets the value of NonRssIPv6Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyNonRssIPv6Only() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonRssIPv6Only")
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

// SetNonRssIPv6PerTcp sets the value of NonRssIPv6PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyNonRssIPv6PerTcp(value uint64) (err error) {
	return instance.SetProperty("NonRssIPv6PerTcp", (value))
}

// GetNonRssIPv6PerTcp gets the value of NonRssIPv6PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyNonRssIPv6PerTcp() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonRssIPv6PerTcp")
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

// SetNonRssIPv6PerUdp sets the value of NonRssIPv6PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyNonRssIPv6PerUdp(value uint64) (err error) {
	return instance.SetProperty("NonRssIPv6PerUdp", (value))
}

// GetNonRssIPv6PerUdp gets the value of NonRssIPv6PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyNonRssIPv6PerUdp() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonRssIPv6PerUdp")
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

// SetNonRssMisc sets the value of NonRssMisc for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyNonRssMisc(value uint64) (err error) {
	return instance.SetProperty("NonRssMisc", (value))
}

// GetNonRssMisc gets the value of NonRssMisc for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyNonRssMisc() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonRssMisc")
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

// SetRssIPv4Only sets the value of RssIPv4Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyRssIPv4Only(value uint64) (err error) {
	return instance.SetProperty("RssIPv4Only", (value))
}

// GetRssIPv4Only gets the value of RssIPv4Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyRssIPv4Only() (value uint64, err error) {
	retValue, err := instance.GetProperty("RssIPv4Only")
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

// SetRssIPv4PerTcp sets the value of RssIPv4PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyRssIPv4PerTcp(value uint64) (err error) {
	return instance.SetProperty("RssIPv4PerTcp", (value))
}

// GetRssIPv4PerTcp gets the value of RssIPv4PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyRssIPv4PerTcp() (value uint64, err error) {
	retValue, err := instance.GetProperty("RssIPv4PerTcp")
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

// SetRssIPv4PerUdp sets the value of RssIPv4PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyRssIPv4PerUdp(value uint64) (err error) {
	return instance.SetProperty("RssIPv4PerUdp", (value))
}

// GetRssIPv4PerUdp gets the value of RssIPv4PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyRssIPv4PerUdp() (value uint64, err error) {
	retValue, err := instance.GetProperty("RssIPv4PerUdp")
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

// SetRssIPv6Only sets the value of RssIPv6Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyRssIPv6Only(value uint64) (err error) {
	return instance.SetProperty("RssIPv6Only", (value))
}

// GetRssIPv6Only gets the value of RssIPv6Only for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyRssIPv6Only() (value uint64, err error) {
	retValue, err := instance.GetProperty("RssIPv6Only")
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

// SetRssIPv6PerTcp sets the value of RssIPv6PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyRssIPv6PerTcp(value uint64) (err error) {
	return instance.SetProperty("RssIPv6PerTcp", (value))
}

// GetRssIPv6PerTcp gets the value of RssIPv6PerTcp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyRssIPv6PerTcp() (value uint64, err error) {
	retValue, err := instance.GetProperty("RssIPv6PerTcp")
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

// SetRssIPv6PerUdp sets the value of RssIPv6PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyRssIPv6PerUdp(value uint64) (err error) {
	return instance.SetProperty("RssIPv6PerUdp", (value))
}

// GetRssIPv6PerUdp gets the value of RssIPv6PerUdp for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyRssIPv6PerUdp() (value uint64, err error) {
	retValue, err := instance.GetProperty("RssIPv6PerUdp")
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

// SetRssMisc sets the value of RssMisc for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) SetPropertyRssMisc(value uint64) (err error) {
	return instance.SetProperty("RssMisc", (value))
}

// GetRssMisc gets the value of RssMisc for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2RssCounters) GetPropertyRssMisc() (value uint64, err error) {
	retValue, err := instance.GetProperty("RssMisc")
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
