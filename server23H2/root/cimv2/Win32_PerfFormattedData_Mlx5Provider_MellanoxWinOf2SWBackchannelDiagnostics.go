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

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics struct {
	*Win32_PerfFormattedData

	//
	CurrentlyActiveCapabilitiesBitmask uint32

	//
	IllegalOrUnsupportedReadConfigBlockOIDs uint64

	//
	IllegalOrUnsupportedWriteConfigBlockOIDs uint64

	//
	ReadConfigBlockOIDsFailedToApply uint64

	//
	ReadConfigBlockOIDsPerSec uint32

	//
	SupportedCapabilitiesBitmask uint32

	//
	WriteConfigBlockOIDsFailedToApply uint64

	//
	WriteConfigBlockOIDsPerSec uint32
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnosticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetCurrentlyActiveCapabilitiesBitmask sets the value of CurrentlyActiveCapabilitiesBitmask for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) SetPropertyCurrentlyActiveCapabilitiesBitmask(value uint32) (err error) {
	return instance.SetProperty("CurrentlyActiveCapabilitiesBitmask", (value))
}

// GetCurrentlyActiveCapabilitiesBitmask gets the value of CurrentlyActiveCapabilitiesBitmask for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) GetPropertyCurrentlyActiveCapabilitiesBitmask() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentlyActiveCapabilitiesBitmask")
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

// SetIllegalOrUnsupportedReadConfigBlockOIDs sets the value of IllegalOrUnsupportedReadConfigBlockOIDs for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) SetPropertyIllegalOrUnsupportedReadConfigBlockOIDs(value uint64) (err error) {
	return instance.SetProperty("IllegalOrUnsupportedReadConfigBlockOIDs", (value))
}

// GetIllegalOrUnsupportedReadConfigBlockOIDs gets the value of IllegalOrUnsupportedReadConfigBlockOIDs for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) GetPropertyIllegalOrUnsupportedReadConfigBlockOIDs() (value uint64, err error) {
	retValue, err := instance.GetProperty("IllegalOrUnsupportedReadConfigBlockOIDs")
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

// SetIllegalOrUnsupportedWriteConfigBlockOIDs sets the value of IllegalOrUnsupportedWriteConfigBlockOIDs for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) SetPropertyIllegalOrUnsupportedWriteConfigBlockOIDs(value uint64) (err error) {
	return instance.SetProperty("IllegalOrUnsupportedWriteConfigBlockOIDs", (value))
}

// GetIllegalOrUnsupportedWriteConfigBlockOIDs gets the value of IllegalOrUnsupportedWriteConfigBlockOIDs for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) GetPropertyIllegalOrUnsupportedWriteConfigBlockOIDs() (value uint64, err error) {
	retValue, err := instance.GetProperty("IllegalOrUnsupportedWriteConfigBlockOIDs")
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

// SetReadConfigBlockOIDsFailedToApply sets the value of ReadConfigBlockOIDsFailedToApply for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) SetPropertyReadConfigBlockOIDsFailedToApply(value uint64) (err error) {
	return instance.SetProperty("ReadConfigBlockOIDsFailedToApply", (value))
}

// GetReadConfigBlockOIDsFailedToApply gets the value of ReadConfigBlockOIDsFailedToApply for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) GetPropertyReadConfigBlockOIDsFailedToApply() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadConfigBlockOIDsFailedToApply")
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

// SetReadConfigBlockOIDsPerSec sets the value of ReadConfigBlockOIDsPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) SetPropertyReadConfigBlockOIDsPerSec(value uint32) (err error) {
	return instance.SetProperty("ReadConfigBlockOIDsPerSec", (value))
}

// GetReadConfigBlockOIDsPerSec gets the value of ReadConfigBlockOIDsPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) GetPropertyReadConfigBlockOIDsPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadConfigBlockOIDsPerSec")
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

// SetSupportedCapabilitiesBitmask sets the value of SupportedCapabilitiesBitmask for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) SetPropertySupportedCapabilitiesBitmask(value uint32) (err error) {
	return instance.SetProperty("SupportedCapabilitiesBitmask", (value))
}

// GetSupportedCapabilitiesBitmask gets the value of SupportedCapabilitiesBitmask for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) GetPropertySupportedCapabilitiesBitmask() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedCapabilitiesBitmask")
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

// SetWriteConfigBlockOIDsFailedToApply sets the value of WriteConfigBlockOIDsFailedToApply for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) SetPropertyWriteConfigBlockOIDsFailedToApply(value uint64) (err error) {
	return instance.SetProperty("WriteConfigBlockOIDsFailedToApply", (value))
}

// GetWriteConfigBlockOIDsFailedToApply gets the value of WriteConfigBlockOIDsFailedToApply for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) GetPropertyWriteConfigBlockOIDsFailedToApply() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteConfigBlockOIDsFailedToApply")
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

// SetWriteConfigBlockOIDsPerSec sets the value of WriteConfigBlockOIDsPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) SetPropertyWriteConfigBlockOIDsPerSec(value uint32) (err error) {
	return instance.SetProperty("WriteConfigBlockOIDsPerSec", (value))
}

// GetWriteConfigBlockOIDsPerSec gets the value of WriteConfigBlockOIDsPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOf2SWBackchannelDiagnostics) GetPropertyWriteConfigBlockOIDsPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteConfigBlockOIDsPerSec")
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
