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

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks struct {
	*Win32_PerfFormattedData

	//
	Internalbenefit uint64

	//
	Memorybrokerclerksize uint64

	//
	Periodicevictionspages uint64

	//
	PressureevictionspagesPersec uint64

	//
	Simulationbenefit uint64

	//
	Simulationsize uint64
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerksEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerksEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetInternalbenefit sets the value of Internalbenefit for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) SetPropertyInternalbenefit(value uint64) (err error) {
	return instance.SetProperty("Internalbenefit", (value))
}

// GetInternalbenefit gets the value of Internalbenefit for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) GetPropertyInternalbenefit() (value uint64, err error) {
	retValue, err := instance.GetProperty("Internalbenefit")
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

// SetMemorybrokerclerksize sets the value of Memorybrokerclerksize for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) SetPropertyMemorybrokerclerksize(value uint64) (err error) {
	return instance.SetProperty("Memorybrokerclerksize", (value))
}

// GetMemorybrokerclerksize gets the value of Memorybrokerclerksize for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) GetPropertyMemorybrokerclerksize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Memorybrokerclerksize")
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

// SetPeriodicevictionspages sets the value of Periodicevictionspages for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) SetPropertyPeriodicevictionspages(value uint64) (err error) {
	return instance.SetProperty("Periodicevictionspages", (value))
}

// GetPeriodicevictionspages gets the value of Periodicevictionspages for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) GetPropertyPeriodicevictionspages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Periodicevictionspages")
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

// SetPressureevictionspagesPersec sets the value of PressureevictionspagesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) SetPropertyPressureevictionspagesPersec(value uint64) (err error) {
	return instance.SetProperty("PressureevictionspagesPersec", (value))
}

// GetPressureevictionspagesPersec gets the value of PressureevictionspagesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) GetPropertyPressureevictionspagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PressureevictionspagesPersec")
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

// SetSimulationbenefit sets the value of Simulationbenefit for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) SetPropertySimulationbenefit(value uint64) (err error) {
	return instance.SetProperty("Simulationbenefit", (value))
}

// GetSimulationbenefit gets the value of Simulationbenefit for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) GetPropertySimulationbenefit() (value uint64, err error) {
	retValue, err := instance.GetProperty("Simulationbenefit")
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

// SetSimulationsize sets the value of Simulationsize for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) SetPropertySimulationsize(value uint64) (err error) {
	return instance.SetProperty("Simulationsize", (value))
}

// GetSimulationsize gets the value of Simulationsize for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryBrokerClerks) GetPropertySimulationsize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Simulationsize")
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
