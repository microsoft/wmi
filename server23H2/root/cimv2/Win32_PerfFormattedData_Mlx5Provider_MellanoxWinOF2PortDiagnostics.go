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

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics struct {
	*Win32_PerfFormattedData

	//
	RXErrorLane0phy uint64

	//
	RXErrorLane0phyPerSec uint32

	//
	RXErrorLane1phy uint64

	//
	RXErrorLane1phyPerSec uint32

	//
	RXErrorLane2phy uint64

	//
	RXErrorLane2phyPerSec uint32

	//
	RXErrorLane3phy uint64

	//
	RXErrorLane3phyPerSec uint32

	//
	RXKbitsphy uint64

	//
	RXKbitsphyPerSec uint32

	//
	RXPCSCorrectedBitsphy uint64

	//
	RXPCSCorrectedBitsphyPerSec uint32

	//
	RXPCSSymbolErrorphy uint64

	//
	RXPCSSymbolErrorphyPerSec uint32
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnosticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetRXErrorLane0phy sets the value of RXErrorLane0phy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXErrorLane0phy(value uint64) (err error) {
	return instance.SetProperty("RXErrorLane0phy", (value))
}

// GetRXErrorLane0phy gets the value of RXErrorLane0phy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXErrorLane0phy() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXErrorLane0phy")
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

// SetRXErrorLane0phyPerSec sets the value of RXErrorLane0phyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXErrorLane0phyPerSec(value uint32) (err error) {
	return instance.SetProperty("RXErrorLane0phyPerSec", (value))
}

// GetRXErrorLane0phyPerSec gets the value of RXErrorLane0phyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXErrorLane0phyPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RXErrorLane0phyPerSec")
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

// SetRXErrorLane1phy sets the value of RXErrorLane1phy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXErrorLane1phy(value uint64) (err error) {
	return instance.SetProperty("RXErrorLane1phy", (value))
}

// GetRXErrorLane1phy gets the value of RXErrorLane1phy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXErrorLane1phy() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXErrorLane1phy")
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

// SetRXErrorLane1phyPerSec sets the value of RXErrorLane1phyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXErrorLane1phyPerSec(value uint32) (err error) {
	return instance.SetProperty("RXErrorLane1phyPerSec", (value))
}

// GetRXErrorLane1phyPerSec gets the value of RXErrorLane1phyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXErrorLane1phyPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RXErrorLane1phyPerSec")
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

// SetRXErrorLane2phy sets the value of RXErrorLane2phy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXErrorLane2phy(value uint64) (err error) {
	return instance.SetProperty("RXErrorLane2phy", (value))
}

// GetRXErrorLane2phy gets the value of RXErrorLane2phy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXErrorLane2phy() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXErrorLane2phy")
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

// SetRXErrorLane2phyPerSec sets the value of RXErrorLane2phyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXErrorLane2phyPerSec(value uint32) (err error) {
	return instance.SetProperty("RXErrorLane2phyPerSec", (value))
}

// GetRXErrorLane2phyPerSec gets the value of RXErrorLane2phyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXErrorLane2phyPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RXErrorLane2phyPerSec")
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

// SetRXErrorLane3phy sets the value of RXErrorLane3phy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXErrorLane3phy(value uint64) (err error) {
	return instance.SetProperty("RXErrorLane3phy", (value))
}

// GetRXErrorLane3phy gets the value of RXErrorLane3phy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXErrorLane3phy() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXErrorLane3phy")
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

// SetRXErrorLane3phyPerSec sets the value of RXErrorLane3phyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXErrorLane3phyPerSec(value uint32) (err error) {
	return instance.SetProperty("RXErrorLane3phyPerSec", (value))
}

// GetRXErrorLane3phyPerSec gets the value of RXErrorLane3phyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXErrorLane3phyPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RXErrorLane3phyPerSec")
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

// SetRXKbitsphy sets the value of RXKbitsphy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXKbitsphy(value uint64) (err error) {
	return instance.SetProperty("RXKbitsphy", (value))
}

// GetRXKbitsphy gets the value of RXKbitsphy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXKbitsphy() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXKbitsphy")
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

// SetRXKbitsphyPerSec sets the value of RXKbitsphyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXKbitsphyPerSec(value uint32) (err error) {
	return instance.SetProperty("RXKbitsphyPerSec", (value))
}

// GetRXKbitsphyPerSec gets the value of RXKbitsphyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXKbitsphyPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RXKbitsphyPerSec")
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

// SetRXPCSCorrectedBitsphy sets the value of RXPCSCorrectedBitsphy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXPCSCorrectedBitsphy(value uint64) (err error) {
	return instance.SetProperty("RXPCSCorrectedBitsphy", (value))
}

// GetRXPCSCorrectedBitsphy gets the value of RXPCSCorrectedBitsphy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXPCSCorrectedBitsphy() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXPCSCorrectedBitsphy")
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

// SetRXPCSCorrectedBitsphyPerSec sets the value of RXPCSCorrectedBitsphyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXPCSCorrectedBitsphyPerSec(value uint32) (err error) {
	return instance.SetProperty("RXPCSCorrectedBitsphyPerSec", (value))
}

// GetRXPCSCorrectedBitsphyPerSec gets the value of RXPCSCorrectedBitsphyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXPCSCorrectedBitsphyPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RXPCSCorrectedBitsphyPerSec")
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

// SetRXPCSSymbolErrorphy sets the value of RXPCSSymbolErrorphy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXPCSSymbolErrorphy(value uint64) (err error) {
	return instance.SetProperty("RXPCSSymbolErrorphy", (value))
}

// GetRXPCSSymbolErrorphy gets the value of RXPCSSymbolErrorphy for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXPCSSymbolErrorphy() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXPCSSymbolErrorphy")
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

// SetRXPCSSymbolErrorphyPerSec sets the value of RXPCSSymbolErrorphyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) SetPropertyRXPCSSymbolErrorphyPerSec(value uint32) (err error) {
	return instance.SetProperty("RXPCSSymbolErrorphyPerSec", (value))
}

// GetRXPCSSymbolErrorphyPerSec gets the value of RXPCSSymbolErrorphyPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2PortDiagnostics) GetPropertyRXPCSSymbolErrorphyPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RXPCSSymbolErrorphyPerSec")
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
