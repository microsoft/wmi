// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor struct
type Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor struct {
	*Win32_PerfFormattedData

	//
	BackingTimerProgrammedPersec uint64

	//
	EarlyTimerPersec uint64

	//
	InterProcessorInterruptsPersec uint64

	//
	InterProcessorInterruptsSentPersec uint64

	//
	TimerInterruptsPersec uint64

	//
	TotalInterruptsPersec uint64
}

func NewWin32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessorEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetBackingTimerProgrammedPersec sets the value of BackingTimerProgrammedPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) SetPropertyBackingTimerProgrammedPersec(value uint64) (err error) {
	return instance.SetProperty("BackingTimerProgrammedPersec", (value))
}

// GetBackingTimerProgrammedPersec gets the value of BackingTimerProgrammedPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) GetPropertyBackingTimerProgrammedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BackingTimerProgrammedPersec")
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

// SetEarlyTimerPersec sets the value of EarlyTimerPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) SetPropertyEarlyTimerPersec(value uint64) (err error) {
	return instance.SetProperty("EarlyTimerPersec", (value))
}

// GetEarlyTimerPersec gets the value of EarlyTimerPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) GetPropertyEarlyTimerPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EarlyTimerPersec")
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

// SetInterProcessorInterruptsPersec sets the value of InterProcessorInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) SetPropertyInterProcessorInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("InterProcessorInterruptsPersec", (value))
}

// GetInterProcessorInterruptsPersec gets the value of InterProcessorInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) GetPropertyInterProcessorInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterProcessorInterruptsPersec")
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

// SetInterProcessorInterruptsSentPersec sets the value of InterProcessorInterruptsSentPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) SetPropertyInterProcessorInterruptsSentPersec(value uint64) (err error) {
	return instance.SetProperty("InterProcessorInterruptsSentPersec", (value))
}

// GetInterProcessorInterruptsSentPersec gets the value of InterProcessorInterruptsSentPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) GetPropertyInterProcessorInterruptsSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterProcessorInterruptsSentPersec")
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

// SetTimerInterruptsPersec sets the value of TimerInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) SetPropertyTimerInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("TimerInterruptsPersec", (value))
}

// GetTimerInterruptsPersec gets the value of TimerInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) GetPropertyTimerInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TimerInterruptsPersec")
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

// SetTotalInterruptsPersec sets the value of TotalInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) SetPropertyTotalInterruptsPersec(value uint64) (err error) {
	return instance.SetProperty("TotalInterruptsPersec", (value))
}

// GetTotalInterruptsPersec gets the value of TotalInterruptsPersec for the instance
func (instance *Win32_PerfFormattedData_HclPerfProvider_HyperVManagementVTLProcessor) GetPropertyTotalInterruptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalInterruptsPersec")
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
