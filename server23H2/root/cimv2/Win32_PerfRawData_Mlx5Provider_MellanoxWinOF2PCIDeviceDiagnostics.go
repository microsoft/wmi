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

// Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics struct
type Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics struct {
	*Win32_PerfRawData

	//
	AvailablePCIBW uint64

	//
	AvailablePCIBWPerSec uint32

	//
	PCIbackpressurecycles uint64

	//
	PCIbackpressurecyclesPerSec uint32

	//
	PCIlinkspeed uint64

	//
	PCIlinkwidth uint64

	//
	PCIreadbackpressurecycles uint64

	//
	PCIreadbackpressurecyclesPerSec uint32

	//
	PCIreadstucknoreceivebuffer uint64

	//
	PCIwritebackpressurecycles uint64

	//
	PCIwritebackpressurecyclesPerSec uint32

	//
	RXPACKETDROPSPCIeBUFFERS uint64

	//
	RXPACKETMARKEDPCIeBUFFERS uint64

	//
	RXPCIerrors uint64

	//
	TXPCIerrors uint64

	//
	TXPCIfatalerrors uint64

	//
	TXPCInonfatalerrors uint64

	//
	UsedPCIBW uint64

	//
	UsedPCIBWPerSec uint32
}

func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnosticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAvailablePCIBW sets the value of AvailablePCIBW for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyAvailablePCIBW(value uint64) (err error) {
	return instance.SetProperty("AvailablePCIBW", (value))
}

// GetAvailablePCIBW gets the value of AvailablePCIBW for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyAvailablePCIBW() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailablePCIBW")
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

// SetAvailablePCIBWPerSec sets the value of AvailablePCIBWPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyAvailablePCIBWPerSec(value uint32) (err error) {
	return instance.SetProperty("AvailablePCIBWPerSec", (value))
}

// GetAvailablePCIBWPerSec gets the value of AvailablePCIBWPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyAvailablePCIBWPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvailablePCIBWPerSec")
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

// SetPCIbackpressurecycles sets the value of PCIbackpressurecycles for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyPCIbackpressurecycles(value uint64) (err error) {
	return instance.SetProperty("PCIbackpressurecycles", (value))
}

// GetPCIbackpressurecycles gets the value of PCIbackpressurecycles for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyPCIbackpressurecycles() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCIbackpressurecycles")
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

// SetPCIbackpressurecyclesPerSec sets the value of PCIbackpressurecyclesPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyPCIbackpressurecyclesPerSec(value uint32) (err error) {
	return instance.SetProperty("PCIbackpressurecyclesPerSec", (value))
}

// GetPCIbackpressurecyclesPerSec gets the value of PCIbackpressurecyclesPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyPCIbackpressurecyclesPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PCIbackpressurecyclesPerSec")
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

// SetPCIlinkspeed sets the value of PCIlinkspeed for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyPCIlinkspeed(value uint64) (err error) {
	return instance.SetProperty("PCIlinkspeed", (value))
}

// GetPCIlinkspeed gets the value of PCIlinkspeed for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyPCIlinkspeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCIlinkspeed")
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

// SetPCIlinkwidth sets the value of PCIlinkwidth for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyPCIlinkwidth(value uint64) (err error) {
	return instance.SetProperty("PCIlinkwidth", (value))
}

// GetPCIlinkwidth gets the value of PCIlinkwidth for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyPCIlinkwidth() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCIlinkwidth")
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

// SetPCIreadbackpressurecycles sets the value of PCIreadbackpressurecycles for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyPCIreadbackpressurecycles(value uint64) (err error) {
	return instance.SetProperty("PCIreadbackpressurecycles", (value))
}

// GetPCIreadbackpressurecycles gets the value of PCIreadbackpressurecycles for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyPCIreadbackpressurecycles() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCIreadbackpressurecycles")
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

// SetPCIreadbackpressurecyclesPerSec sets the value of PCIreadbackpressurecyclesPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyPCIreadbackpressurecyclesPerSec(value uint32) (err error) {
	return instance.SetProperty("PCIreadbackpressurecyclesPerSec", (value))
}

// GetPCIreadbackpressurecyclesPerSec gets the value of PCIreadbackpressurecyclesPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyPCIreadbackpressurecyclesPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PCIreadbackpressurecyclesPerSec")
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

// SetPCIreadstucknoreceivebuffer sets the value of PCIreadstucknoreceivebuffer for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyPCIreadstucknoreceivebuffer(value uint64) (err error) {
	return instance.SetProperty("PCIreadstucknoreceivebuffer", (value))
}

// GetPCIreadstucknoreceivebuffer gets the value of PCIreadstucknoreceivebuffer for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyPCIreadstucknoreceivebuffer() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCIreadstucknoreceivebuffer")
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

// SetPCIwritebackpressurecycles sets the value of PCIwritebackpressurecycles for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyPCIwritebackpressurecycles(value uint64) (err error) {
	return instance.SetProperty("PCIwritebackpressurecycles", (value))
}

// GetPCIwritebackpressurecycles gets the value of PCIwritebackpressurecycles for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyPCIwritebackpressurecycles() (value uint64, err error) {
	retValue, err := instance.GetProperty("PCIwritebackpressurecycles")
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

// SetPCIwritebackpressurecyclesPerSec sets the value of PCIwritebackpressurecyclesPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyPCIwritebackpressurecyclesPerSec(value uint32) (err error) {
	return instance.SetProperty("PCIwritebackpressurecyclesPerSec", (value))
}

// GetPCIwritebackpressurecyclesPerSec gets the value of PCIwritebackpressurecyclesPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyPCIwritebackpressurecyclesPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PCIwritebackpressurecyclesPerSec")
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

// SetRXPACKETDROPSPCIeBUFFERS sets the value of RXPACKETDROPSPCIeBUFFERS for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyRXPACKETDROPSPCIeBUFFERS(value uint64) (err error) {
	return instance.SetProperty("RXPACKETDROPSPCIeBUFFERS", (value))
}

// GetRXPACKETDROPSPCIeBUFFERS gets the value of RXPACKETDROPSPCIeBUFFERS for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyRXPACKETDROPSPCIeBUFFERS() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXPACKETDROPSPCIeBUFFERS")
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

// SetRXPACKETMARKEDPCIeBUFFERS sets the value of RXPACKETMARKEDPCIeBUFFERS for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyRXPACKETMARKEDPCIeBUFFERS(value uint64) (err error) {
	return instance.SetProperty("RXPACKETMARKEDPCIeBUFFERS", (value))
}

// GetRXPACKETMARKEDPCIeBUFFERS gets the value of RXPACKETMARKEDPCIeBUFFERS for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyRXPACKETMARKEDPCIeBUFFERS() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXPACKETMARKEDPCIeBUFFERS")
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

// SetRXPCIerrors sets the value of RXPCIerrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyRXPCIerrors(value uint64) (err error) {
	return instance.SetProperty("RXPCIerrors", (value))
}

// GetRXPCIerrors gets the value of RXPCIerrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyRXPCIerrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXPCIerrors")
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

// SetTXPCIerrors sets the value of TXPCIerrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyTXPCIerrors(value uint64) (err error) {
	return instance.SetProperty("TXPCIerrors", (value))
}

// GetTXPCIerrors gets the value of TXPCIerrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyTXPCIerrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("TXPCIerrors")
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

// SetTXPCIfatalerrors sets the value of TXPCIfatalerrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyTXPCIfatalerrors(value uint64) (err error) {
	return instance.SetProperty("TXPCIfatalerrors", (value))
}

// GetTXPCIfatalerrors gets the value of TXPCIfatalerrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyTXPCIfatalerrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("TXPCIfatalerrors")
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

// SetTXPCInonfatalerrors sets the value of TXPCInonfatalerrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyTXPCInonfatalerrors(value uint64) (err error) {
	return instance.SetProperty("TXPCInonfatalerrors", (value))
}

// GetTXPCInonfatalerrors gets the value of TXPCInonfatalerrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyTXPCInonfatalerrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("TXPCInonfatalerrors")
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

// SetUsedPCIBW sets the value of UsedPCIBW for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyUsedPCIBW(value uint64) (err error) {
	return instance.SetProperty("UsedPCIBW", (value))
}

// GetUsedPCIBW gets the value of UsedPCIBW for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyUsedPCIBW() (value uint64, err error) {
	retValue, err := instance.GetProperty("UsedPCIBW")
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

// SetUsedPCIBWPerSec sets the value of UsedPCIBWPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) SetPropertyUsedPCIBWPerSec(value uint32) (err error) {
	return instance.SetProperty("UsedPCIBWPerSec", (value))
}

// GetUsedPCIBWPerSec gets the value of UsedPCIBWPerSec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2PCIDeviceDiagnostics) GetPropertyUsedPCIBWPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("UsedPCIBWPerSec")
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
