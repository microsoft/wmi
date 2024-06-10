// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.SDDC.Management
//
// ////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SDDC_Processor struct
type SDDC_Processor struct {
	*cim.WmiInstance

	//
	Manufacturer string

	//
	MaxClockSpeedInMHz uint32

	//
	Model string

	//
	NumberOfCores uint32

	//
	NumberOfLogicalProcessors uint32
}

func NewSDDC_ProcessorEx1(instance *cim.WmiInstance) (newInstance *SDDC_Processor, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_Processor{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_ProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_Processor, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_Processor{
		WmiInstance: tmp,
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *SDDC_Processor) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *SDDC_Processor) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetMaxClockSpeedInMHz sets the value of MaxClockSpeedInMHz for the instance
func (instance *SDDC_Processor) SetPropertyMaxClockSpeedInMHz(value uint32) (err error) {
	return instance.SetProperty("MaxClockSpeedInMHz", (value))
}

// GetMaxClockSpeedInMHz gets the value of MaxClockSpeedInMHz for the instance
func (instance *SDDC_Processor) GetPropertyMaxClockSpeedInMHz() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxClockSpeedInMHz")
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

// SetModel sets the value of Model for the instance
func (instance *SDDC_Processor) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", (value))
}

// GetModel gets the value of Model for the instance
func (instance *SDDC_Processor) GetPropertyModel() (value string, err error) {
	retValue, err := instance.GetProperty("Model")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetNumberOfCores sets the value of NumberOfCores for the instance
func (instance *SDDC_Processor) SetPropertyNumberOfCores(value uint32) (err error) {
	return instance.SetProperty("NumberOfCores", (value))
}

// GetNumberOfCores gets the value of NumberOfCores for the instance
func (instance *SDDC_Processor) GetPropertyNumberOfCores() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfCores")
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

// SetNumberOfLogicalProcessors sets the value of NumberOfLogicalProcessors for the instance
func (instance *SDDC_Processor) SetPropertyNumberOfLogicalProcessors(value uint32) (err error) {
	return instance.SetProperty("NumberOfLogicalProcessors", (value))
}

// GetNumberOfLogicalProcessors gets the value of NumberOfLogicalProcessors for the instance
func (instance *SDDC_Processor) GetPropertyNumberOfLogicalProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfLogicalProcessors")
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
