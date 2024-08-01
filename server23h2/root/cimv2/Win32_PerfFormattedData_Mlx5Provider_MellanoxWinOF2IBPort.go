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

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort struct {
	*Win32_PerfFormattedData

	//
	LinkDowned uint32

	//
	LinkErrorRecoveryCounter uint32

	//
	LocalLinkIntegrityErrors uint32

	//
	PortMultiCastReceivePackets uint32

	//
	PortMultiCastXmitPackets uint32

	//
	PortReceiveConstraintErrors uint32

	//
	PortReceiveData uint32

	//
	PortReceiveErrors uint32

	//
	PortReceivePackets uint32

	//
	PortReceiveRemotePhysicalErrors uint32

	//
	PortReceiveSwitchRelayErrors uint32

	//
	PortUnicastReceivePackets uint32

	//
	PortUnicastXmitPackets uint32

	//
	PortXmitConstraintErrors uint32

	//
	PortXmitData uint32

	//
	PortXmitDiscards uint32

	//
	PortXmitPackets uint32

	//
	PortXmitWait uint32

	//
	SymbolError uint32

	//
	VL15Dropped uint32
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPortEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetLinkDowned sets the value of LinkDowned for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyLinkDowned(value uint32) (err error) {
	return instance.SetProperty("LinkDowned", (value))
}

// GetLinkDowned gets the value of LinkDowned for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyLinkDowned() (value uint32, err error) {
	retValue, err := instance.GetProperty("LinkDowned")
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

// SetLinkErrorRecoveryCounter sets the value of LinkErrorRecoveryCounter for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyLinkErrorRecoveryCounter(value uint32) (err error) {
	return instance.SetProperty("LinkErrorRecoveryCounter", (value))
}

// GetLinkErrorRecoveryCounter gets the value of LinkErrorRecoveryCounter for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyLinkErrorRecoveryCounter() (value uint32, err error) {
	retValue, err := instance.GetProperty("LinkErrorRecoveryCounter")
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

// SetLocalLinkIntegrityErrors sets the value of LocalLinkIntegrityErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyLocalLinkIntegrityErrors(value uint32) (err error) {
	return instance.SetProperty("LocalLinkIntegrityErrors", (value))
}

// GetLocalLinkIntegrityErrors gets the value of LocalLinkIntegrityErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyLocalLinkIntegrityErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("LocalLinkIntegrityErrors")
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

// SetPortMultiCastReceivePackets sets the value of PortMultiCastReceivePackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortMultiCastReceivePackets(value uint32) (err error) {
	return instance.SetProperty("PortMultiCastReceivePackets", (value))
}

// GetPortMultiCastReceivePackets gets the value of PortMultiCastReceivePackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortMultiCastReceivePackets() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortMultiCastReceivePackets")
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

// SetPortMultiCastXmitPackets sets the value of PortMultiCastXmitPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortMultiCastXmitPackets(value uint32) (err error) {
	return instance.SetProperty("PortMultiCastXmitPackets", (value))
}

// GetPortMultiCastXmitPackets gets the value of PortMultiCastXmitPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortMultiCastXmitPackets() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortMultiCastXmitPackets")
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

// SetPortReceiveConstraintErrors sets the value of PortReceiveConstraintErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortReceiveConstraintErrors(value uint32) (err error) {
	return instance.SetProperty("PortReceiveConstraintErrors", (value))
}

// GetPortReceiveConstraintErrors gets the value of PortReceiveConstraintErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortReceiveConstraintErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortReceiveConstraintErrors")
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

// SetPortReceiveData sets the value of PortReceiveData for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortReceiveData(value uint32) (err error) {
	return instance.SetProperty("PortReceiveData", (value))
}

// GetPortReceiveData gets the value of PortReceiveData for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortReceiveData() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortReceiveData")
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

// SetPortReceiveErrors sets the value of PortReceiveErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortReceiveErrors(value uint32) (err error) {
	return instance.SetProperty("PortReceiveErrors", (value))
}

// GetPortReceiveErrors gets the value of PortReceiveErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortReceiveErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortReceiveErrors")
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

// SetPortReceivePackets sets the value of PortReceivePackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortReceivePackets(value uint32) (err error) {
	return instance.SetProperty("PortReceivePackets", (value))
}

// GetPortReceivePackets gets the value of PortReceivePackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortReceivePackets() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortReceivePackets")
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

// SetPortReceiveRemotePhysicalErrors sets the value of PortReceiveRemotePhysicalErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortReceiveRemotePhysicalErrors(value uint32) (err error) {
	return instance.SetProperty("PortReceiveRemotePhysicalErrors", (value))
}

// GetPortReceiveRemotePhysicalErrors gets the value of PortReceiveRemotePhysicalErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortReceiveRemotePhysicalErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortReceiveRemotePhysicalErrors")
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

// SetPortReceiveSwitchRelayErrors sets the value of PortReceiveSwitchRelayErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortReceiveSwitchRelayErrors(value uint32) (err error) {
	return instance.SetProperty("PortReceiveSwitchRelayErrors", (value))
}

// GetPortReceiveSwitchRelayErrors gets the value of PortReceiveSwitchRelayErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortReceiveSwitchRelayErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortReceiveSwitchRelayErrors")
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

// SetPortUnicastReceivePackets sets the value of PortUnicastReceivePackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortUnicastReceivePackets(value uint32) (err error) {
	return instance.SetProperty("PortUnicastReceivePackets", (value))
}

// GetPortUnicastReceivePackets gets the value of PortUnicastReceivePackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortUnicastReceivePackets() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortUnicastReceivePackets")
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

// SetPortUnicastXmitPackets sets the value of PortUnicastXmitPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortUnicastXmitPackets(value uint32) (err error) {
	return instance.SetProperty("PortUnicastXmitPackets", (value))
}

// GetPortUnicastXmitPackets gets the value of PortUnicastXmitPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortUnicastXmitPackets() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortUnicastXmitPackets")
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

// SetPortXmitConstraintErrors sets the value of PortXmitConstraintErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortXmitConstraintErrors(value uint32) (err error) {
	return instance.SetProperty("PortXmitConstraintErrors", (value))
}

// GetPortXmitConstraintErrors gets the value of PortXmitConstraintErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortXmitConstraintErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortXmitConstraintErrors")
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

// SetPortXmitData sets the value of PortXmitData for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortXmitData(value uint32) (err error) {
	return instance.SetProperty("PortXmitData", (value))
}

// GetPortXmitData gets the value of PortXmitData for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortXmitData() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortXmitData")
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

// SetPortXmitDiscards sets the value of PortXmitDiscards for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortXmitDiscards(value uint32) (err error) {
	return instance.SetProperty("PortXmitDiscards", (value))
}

// GetPortXmitDiscards gets the value of PortXmitDiscards for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortXmitDiscards() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortXmitDiscards")
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

// SetPortXmitPackets sets the value of PortXmitPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortXmitPackets(value uint32) (err error) {
	return instance.SetProperty("PortXmitPackets", (value))
}

// GetPortXmitPackets gets the value of PortXmitPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortXmitPackets() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortXmitPackets")
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

// SetPortXmitWait sets the value of PortXmitWait for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyPortXmitWait(value uint32) (err error) {
	return instance.SetProperty("PortXmitWait", (value))
}

// GetPortXmitWait gets the value of PortXmitWait for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyPortXmitWait() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortXmitWait")
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

// SetSymbolError sets the value of SymbolError for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertySymbolError(value uint32) (err error) {
	return instance.SetProperty("SymbolError", (value))
}

// GetSymbolError gets the value of SymbolError for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertySymbolError() (value uint32, err error) {
	retValue, err := instance.GetProperty("SymbolError")
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

// SetVL15Dropped sets the value of VL15Dropped for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) SetPropertyVL15Dropped(value uint32) (err error) {
	return instance.SetProperty("VL15Dropped", (value))
}

// GetVL15Dropped gets the value of VL15Dropped for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2IBPort) GetPropertyVL15Dropped() (value uint32, err error) {
	retValue, err := instance.GetProperty("VL15Dropped")
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
