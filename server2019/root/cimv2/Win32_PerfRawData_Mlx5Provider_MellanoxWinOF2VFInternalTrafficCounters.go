// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters struct
type Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters struct {
	*Win32_PerfRawData

	//
	Rxoctets uint64

	//
	Rxpackets uint64

	//
	Txoctets uint64

	//
	Txpackets uint64
}

func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCountersEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCountersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetRxoctets sets the value of Rxoctets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters) SetPropertyRxoctets(value uint64) (err error) {
	return instance.SetProperty("Rxoctets", (value))
}

// GetRxoctets gets the value of Rxoctets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters) GetPropertyRxoctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rxoctets")
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

// SetRxpackets sets the value of Rxpackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters) SetPropertyRxpackets(value uint64) (err error) {
	return instance.SetProperty("Rxpackets", (value))
}

// GetRxpackets gets the value of Rxpackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters) GetPropertyRxpackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("Rxpackets")
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

// SetTxoctets sets the value of Txoctets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters) SetPropertyTxoctets(value uint64) (err error) {
	return instance.SetProperty("Txoctets", (value))
}

// GetTxoctets gets the value of Txoctets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters) GetPropertyTxoctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("Txoctets")
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

// SetTxpackets sets the value of Txpackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters) SetPropertyTxpackets(value uint64) (err error) {
	return instance.SetProperty("Txpackets", (value))
}

// GetTxpackets gets the value of Txpackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2VFInternalTrafficCounters) GetPropertyTxpackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("Txpackets")
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
