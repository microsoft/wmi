// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ProcessorPerformance struct
type ProcessorPerformance struct {
	*Win32_PerfRawData

	//
	Active bool

	//
	frequency uint32

	//
	InstanceName string

	//
	percentage uint32

	//
	power uint32
}

func NewProcessorPerformanceEx1(instance *cim.WmiInstance) (newInstance *ProcessorPerformance, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProcessorPerformance{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewProcessorPerformanceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessorPerformance, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessorPerformance{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *ProcessorPerformance) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ProcessorPerformance) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// Setfrequency sets the value of frequency for the instance
func (instance *ProcessorPerformance) SetPropertyfrequency(value uint32) (err error) {
	return instance.SetProperty("frequency", (value))
}

// Getfrequency gets the value of frequency for the instance
func (instance *ProcessorPerformance) GetPropertyfrequency() (value uint32, err error) {
	retValue, err := instance.GetProperty("frequency")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ProcessorPerformance) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ProcessorPerformance) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// Setpercentage sets the value of percentage for the instance
func (instance *ProcessorPerformance) SetPropertypercentage(value uint32) (err error) {
	return instance.SetProperty("percentage", (value))
}

// Getpercentage gets the value of percentage for the instance
func (instance *ProcessorPerformance) GetPropertypercentage() (value uint32, err error) {
	retValue, err := instance.GetProperty("percentage")
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

// Setpower sets the value of power for the instance
func (instance *ProcessorPerformance) SetPropertypower(value uint32) (err error) {
	return instance.SetProperty("power", (value))
}

// Getpower gets the value of power for the instance
func (instance *ProcessorPerformance) GetPropertypower() (value uint32, err error) {
	retValue, err := instance.GetProperty("power")
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
