// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PerformanceState struct
type PerformanceState struct {
	*cim.WmiInstance

	//
	Flags uint32

	//
	Frequency uint32

	//
	PercentFrequency uint32
}

func NewPerformanceStateEx1(instance *cim.WmiInstance) (newInstance *PerformanceState, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PerformanceState{
		WmiInstance: tmp,
	}
	return
}

func NewPerformanceStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PerformanceState, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PerformanceState{
		WmiInstance: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *PerformanceState) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *PerformanceState) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetFrequency sets the value of Frequency for the instance
func (instance *PerformanceState) SetPropertyFrequency(value uint32) (err error) {
	return instance.SetProperty("Frequency", (value))
}

// GetFrequency gets the value of Frequency for the instance
func (instance *PerformanceState) GetPropertyFrequency() (value uint32, err error) {
	retValue, err := instance.GetProperty("Frequency")
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

// SetPercentFrequency sets the value of PercentFrequency for the instance
func (instance *PerformanceState) SetPropertyPercentFrequency(value uint32) (err error) {
	return instance.SetProperty("PercentFrequency", (value))
}

// GetPercentFrequency gets the value of PercentFrequency for the instance
func (instance *PerformanceState) GetPropertyPercentFrequency() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentFrequency")
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
