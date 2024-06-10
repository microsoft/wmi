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

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics struct {
	*Win32_PerfFormattedData

	//
	AsyncEQOverrun uint64

	//
	CompletionEQOverrun uint64

	//
	CQOverrun uint64

	//
	Currentqueuesunderprocessorhandle uint64

	//
	InvalidCommands uint64

	//
	PacketsReceiveddroppedduetoSteering uint64

	//
	PacketsReceiveddroppedduetoVPortDown uint64

	//
	PacketsReceivedWQEtoosmall uint64

	//
	PacketsTransmitteddroppedduetoVPortDown uint64

	//
	QuotaExceededCommand uint64

	//
	SendQueuePriorityUpdateFlow uint64

	//
	Totalqueuesunderprocessorhandle uint64
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnosticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAsyncEQOverrun sets the value of AsyncEQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyAsyncEQOverrun(value uint64) (err error) {
	return instance.SetProperty("AsyncEQOverrun", (value))
}

// GetAsyncEQOverrun gets the value of AsyncEQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyAsyncEQOverrun() (value uint64, err error) {
	retValue, err := instance.GetProperty("AsyncEQOverrun")
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

// SetCompletionEQOverrun sets the value of CompletionEQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyCompletionEQOverrun(value uint64) (err error) {
	return instance.SetProperty("CompletionEQOverrun", (value))
}

// GetCompletionEQOverrun gets the value of CompletionEQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyCompletionEQOverrun() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompletionEQOverrun")
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

// SetCQOverrun sets the value of CQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyCQOverrun(value uint64) (err error) {
	return instance.SetProperty("CQOverrun", (value))
}

// GetCQOverrun gets the value of CQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyCQOverrun() (value uint64, err error) {
	retValue, err := instance.GetProperty("CQOverrun")
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

// SetCurrentqueuesunderprocessorhandle sets the value of Currentqueuesunderprocessorhandle for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyCurrentqueuesunderprocessorhandle(value uint64) (err error) {
	return instance.SetProperty("Currentqueuesunderprocessorhandle", (value))
}

// GetCurrentqueuesunderprocessorhandle gets the value of Currentqueuesunderprocessorhandle for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyCurrentqueuesunderprocessorhandle() (value uint64, err error) {
	retValue, err := instance.GetProperty("Currentqueuesunderprocessorhandle")
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

// SetInvalidCommands sets the value of InvalidCommands for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyInvalidCommands(value uint64) (err error) {
	return instance.SetProperty("InvalidCommands", (value))
}

// GetInvalidCommands gets the value of InvalidCommands for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyInvalidCommands() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvalidCommands")
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

// SetPacketsReceiveddroppedduetoSteering sets the value of PacketsReceiveddroppedduetoSteering for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyPacketsReceiveddroppedduetoSteering(value uint64) (err error) {
	return instance.SetProperty("PacketsReceiveddroppedduetoSteering", (value))
}

// GetPacketsReceiveddroppedduetoSteering gets the value of PacketsReceiveddroppedduetoSteering for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyPacketsReceiveddroppedduetoSteering() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceiveddroppedduetoSteering")
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

// SetPacketsReceiveddroppedduetoVPortDown sets the value of PacketsReceiveddroppedduetoVPortDown for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyPacketsReceiveddroppedduetoVPortDown(value uint64) (err error) {
	return instance.SetProperty("PacketsReceiveddroppedduetoVPortDown", (value))
}

// GetPacketsReceiveddroppedduetoVPortDown gets the value of PacketsReceiveddroppedduetoVPortDown for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyPacketsReceiveddroppedduetoVPortDown() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceiveddroppedduetoVPortDown")
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

// SetPacketsReceivedWQEtoosmall sets the value of PacketsReceivedWQEtoosmall for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyPacketsReceivedWQEtoosmall(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedWQEtoosmall", (value))
}

// GetPacketsReceivedWQEtoosmall gets the value of PacketsReceivedWQEtoosmall for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyPacketsReceivedWQEtoosmall() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedWQEtoosmall")
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

// SetPacketsTransmitteddroppedduetoVPortDown sets the value of PacketsTransmitteddroppedduetoVPortDown for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyPacketsTransmitteddroppedduetoVPortDown(value uint64) (err error) {
	return instance.SetProperty("PacketsTransmitteddroppedduetoVPortDown", (value))
}

// GetPacketsTransmitteddroppedduetoVPortDown gets the value of PacketsTransmitteddroppedduetoVPortDown for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyPacketsTransmitteddroppedduetoVPortDown() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsTransmitteddroppedduetoVPortDown")
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

// SetQuotaExceededCommand sets the value of QuotaExceededCommand for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyQuotaExceededCommand(value uint64) (err error) {
	return instance.SetProperty("QuotaExceededCommand", (value))
}

// GetQuotaExceededCommand gets the value of QuotaExceededCommand for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyQuotaExceededCommand() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuotaExceededCommand")
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

// SetSendQueuePriorityUpdateFlow sets the value of SendQueuePriorityUpdateFlow for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertySendQueuePriorityUpdateFlow(value uint64) (err error) {
	return instance.SetProperty("SendQueuePriorityUpdateFlow", (value))
}

// GetSendQueuePriorityUpdateFlow gets the value of SendQueuePriorityUpdateFlow for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertySendQueuePriorityUpdateFlow() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendQueuePriorityUpdateFlow")
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

// SetTotalqueuesunderprocessorhandle sets the value of Totalqueuesunderprocessorhandle for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) SetPropertyTotalqueuesunderprocessorhandle(value uint64) (err error) {
	return instance.SetProperty("Totalqueuesunderprocessorhandle", (value))
}

// GetTotalqueuesunderprocessorhandle gets the value of Totalqueuesunderprocessorhandle for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2VFDiagnostics) GetPropertyTotalqueuesunderprocessorhandle() (value uint64, err error) {
	retValue, err := instance.GetProperty("Totalqueuesunderprocessorhandle")
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
