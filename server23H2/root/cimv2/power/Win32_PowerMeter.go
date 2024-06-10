// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.power
//
// ////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PowerMeter struct
type Win32_PowerMeter struct {
	*CIM_NumericSensor

	//
	AveragingInterval uint32

	//
	BudgetEnabled bool

	//
	BudgetWriteable bool

	//
	ConfiguredBudget uint32

	//
	MaximumAveragingInterval uint32

	//
	MaxOperatingBudget uint32

	//
	MeterType uint32

	//
	MinimumAveragingInterval uint32

	//
	MinOperatingBudget uint32

	//
	SamplingPeriod uint32

	//
	SupportCapabilities uint32
}

func NewWin32_PowerMeterEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerMeter, err error) {
	tmp, err := NewCIM_NumericSensorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerMeter{
		CIM_NumericSensor: tmp,
	}
	return
}

func NewWin32_PowerMeterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerMeter, err error) {
	tmp, err := NewCIM_NumericSensorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerMeter{
		CIM_NumericSensor: tmp,
	}
	return
}

// SetAveragingInterval sets the value of AveragingInterval for the instance
func (instance *Win32_PowerMeter) SetPropertyAveragingInterval(value uint32) (err error) {
	return instance.SetProperty("AveragingInterval", (value))
}

// GetAveragingInterval gets the value of AveragingInterval for the instance
func (instance *Win32_PowerMeter) GetPropertyAveragingInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("AveragingInterval")
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

// SetBudgetEnabled sets the value of BudgetEnabled for the instance
func (instance *Win32_PowerMeter) SetPropertyBudgetEnabled(value bool) (err error) {
	return instance.SetProperty("BudgetEnabled", (value))
}

// GetBudgetEnabled gets the value of BudgetEnabled for the instance
func (instance *Win32_PowerMeter) GetPropertyBudgetEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("BudgetEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetBudgetWriteable sets the value of BudgetWriteable for the instance
func (instance *Win32_PowerMeter) SetPropertyBudgetWriteable(value bool) (err error) {
	return instance.SetProperty("BudgetWriteable", (value))
}

// GetBudgetWriteable gets the value of BudgetWriteable for the instance
func (instance *Win32_PowerMeter) GetPropertyBudgetWriteable() (value bool, err error) {
	retValue, err := instance.GetProperty("BudgetWriteable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetConfiguredBudget sets the value of ConfiguredBudget for the instance
func (instance *Win32_PowerMeter) SetPropertyConfiguredBudget(value uint32) (err error) {
	return instance.SetProperty("ConfiguredBudget", (value))
}

// GetConfiguredBudget gets the value of ConfiguredBudget for the instance
func (instance *Win32_PowerMeter) GetPropertyConfiguredBudget() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConfiguredBudget")
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

// SetMaximumAveragingInterval sets the value of MaximumAveragingInterval for the instance
func (instance *Win32_PowerMeter) SetPropertyMaximumAveragingInterval(value uint32) (err error) {
	return instance.SetProperty("MaximumAveragingInterval", (value))
}

// GetMaximumAveragingInterval gets the value of MaximumAveragingInterval for the instance
func (instance *Win32_PowerMeter) GetPropertyMaximumAveragingInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumAveragingInterval")
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

// SetMaxOperatingBudget sets the value of MaxOperatingBudget for the instance
func (instance *Win32_PowerMeter) SetPropertyMaxOperatingBudget(value uint32) (err error) {
	return instance.SetProperty("MaxOperatingBudget", (value))
}

// GetMaxOperatingBudget gets the value of MaxOperatingBudget for the instance
func (instance *Win32_PowerMeter) GetPropertyMaxOperatingBudget() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxOperatingBudget")
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

// SetMeterType sets the value of MeterType for the instance
func (instance *Win32_PowerMeter) SetPropertyMeterType(value uint32) (err error) {
	return instance.SetProperty("MeterType", (value))
}

// GetMeterType gets the value of MeterType for the instance
func (instance *Win32_PowerMeter) GetPropertyMeterType() (value uint32, err error) {
	retValue, err := instance.GetProperty("MeterType")
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

// SetMinimumAveragingInterval sets the value of MinimumAveragingInterval for the instance
func (instance *Win32_PowerMeter) SetPropertyMinimumAveragingInterval(value uint32) (err error) {
	return instance.SetProperty("MinimumAveragingInterval", (value))
}

// GetMinimumAveragingInterval gets the value of MinimumAveragingInterval for the instance
func (instance *Win32_PowerMeter) GetPropertyMinimumAveragingInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumAveragingInterval")
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

// SetMinOperatingBudget sets the value of MinOperatingBudget for the instance
func (instance *Win32_PowerMeter) SetPropertyMinOperatingBudget(value uint32) (err error) {
	return instance.SetProperty("MinOperatingBudget", (value))
}

// GetMinOperatingBudget gets the value of MinOperatingBudget for the instance
func (instance *Win32_PowerMeter) GetPropertyMinOperatingBudget() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinOperatingBudget")
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

// SetSamplingPeriod sets the value of SamplingPeriod for the instance
func (instance *Win32_PowerMeter) SetPropertySamplingPeriod(value uint32) (err error) {
	return instance.SetProperty("SamplingPeriod", (value))
}

// GetSamplingPeriod gets the value of SamplingPeriod for the instance
func (instance *Win32_PowerMeter) GetPropertySamplingPeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("SamplingPeriod")
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

// SetSupportCapabilities sets the value of SupportCapabilities for the instance
func (instance *Win32_PowerMeter) SetPropertySupportCapabilities(value uint32) (err error) {
	return instance.SetProperty("SupportCapabilities", (value))
}

// GetSupportCapabilities gets the value of SupportCapabilities for the instance
func (instance *Win32_PowerMeter) GetPropertySupportCapabilities() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportCapabilities")
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
