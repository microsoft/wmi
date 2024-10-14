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

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1 struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1 struct {
	*Win32_PerfFormattedData

	//
	ICMCCQHit uint64

	//
	ICMCCQMiss uint64

	//
	ICMCQPReceiveHit uint64

	//
	ICMCQPReceiveMiss uint64

	//
	ICMCQPSendHit uint64

	//
	ICMCQPSendMiss uint64

	//
	ICMCSRQHit uint64

	//
	ICMCSRQMiss uint64
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1Ex1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetICMCCQHit sets the value of ICMCCQHit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) SetPropertyICMCCQHit(value uint64) (err error) {
	return instance.SetProperty("ICMCCQHit", (value))
}

// GetICMCCQHit gets the value of ICMCCQHit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) GetPropertyICMCCQHit() (value uint64, err error) {
	retValue, err := instance.GetProperty("ICMCCQHit")
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

// SetICMCCQMiss sets the value of ICMCCQMiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) SetPropertyICMCCQMiss(value uint64) (err error) {
	return instance.SetProperty("ICMCCQMiss", (value))
}

// GetICMCCQMiss gets the value of ICMCCQMiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) GetPropertyICMCCQMiss() (value uint64, err error) {
	retValue, err := instance.GetProperty("ICMCCQMiss")
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

// SetICMCQPReceiveHit sets the value of ICMCQPReceiveHit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) SetPropertyICMCQPReceiveHit(value uint64) (err error) {
	return instance.SetProperty("ICMCQPReceiveHit", (value))
}

// GetICMCQPReceiveHit gets the value of ICMCQPReceiveHit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) GetPropertyICMCQPReceiveHit() (value uint64, err error) {
	retValue, err := instance.GetProperty("ICMCQPReceiveHit")
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

// SetICMCQPReceiveMiss sets the value of ICMCQPReceiveMiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) SetPropertyICMCQPReceiveMiss(value uint64) (err error) {
	return instance.SetProperty("ICMCQPReceiveMiss", (value))
}

// GetICMCQPReceiveMiss gets the value of ICMCQPReceiveMiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) GetPropertyICMCQPReceiveMiss() (value uint64, err error) {
	retValue, err := instance.GetProperty("ICMCQPReceiveMiss")
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

// SetICMCQPSendHit sets the value of ICMCQPSendHit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) SetPropertyICMCQPSendHit(value uint64) (err error) {
	return instance.SetProperty("ICMCQPSendHit", (value))
}

// GetICMCQPSendHit gets the value of ICMCQPSendHit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) GetPropertyICMCQPSendHit() (value uint64, err error) {
	retValue, err := instance.GetProperty("ICMCQPSendHit")
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

// SetICMCQPSendMiss sets the value of ICMCQPSendMiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) SetPropertyICMCQPSendMiss(value uint64) (err error) {
	return instance.SetProperty("ICMCQPSendMiss", (value))
}

// GetICMCQPSendMiss gets the value of ICMCQPSendMiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) GetPropertyICMCQPSendMiss() (value uint64, err error) {
	retValue, err := instance.GetProperty("ICMCQPSendMiss")
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

// SetICMCSRQHit sets the value of ICMCSRQHit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) SetPropertyICMCSRQHit(value uint64) (err error) {
	return instance.SetProperty("ICMCSRQHit", (value))
}

// GetICMCSRQHit gets the value of ICMCSRQHit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) GetPropertyICMCSRQHit() (value uint64, err error) {
	retValue, err := instance.GetProperty("ICMCSRQHit")
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

// SetICMCSRQMiss sets the value of ICMCSRQMiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) SetPropertyICMCSRQMiss(value uint64) (err error) {
	return instance.SetProperty("ICMCSRQMiss", (value))
}

// GetICMCSRQMiss gets the value of ICMCSRQMiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2ICMCDiagCountersExt1) GetPropertyICMCSRQMiss() (value uint64, err error) {
	retValue, err := instance.GetProperty("ICMCSRQMiss")
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
