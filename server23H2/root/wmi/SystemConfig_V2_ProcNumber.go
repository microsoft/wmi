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

// SystemConfig_V2_ProcNumber struct
type SystemConfig_V2_ProcNumber struct {
	*SystemConfig_V2

	//
	ProcessorCount uint32

	//
	ProcessorNumber []uint32
}

func NewSystemConfig_V2_ProcNumberEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_ProcNumber, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_ProcNumber{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_ProcNumberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_ProcNumber, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_ProcNumber{
		SystemConfig_V2: tmp,
	}
	return
}

// SetProcessorCount sets the value of ProcessorCount for the instance
func (instance *SystemConfig_V2_ProcNumber) SetPropertyProcessorCount(value uint32) (err error) {
	return instance.SetProperty("ProcessorCount", (value))
}

// GetProcessorCount gets the value of ProcessorCount for the instance
func (instance *SystemConfig_V2_ProcNumber) GetPropertyProcessorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessorCount")
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

// SetProcessorNumber sets the value of ProcessorNumber for the instance
func (instance *SystemConfig_V2_ProcNumber) SetPropertyProcessorNumber(value []uint32) (err error) {
	return instance.SetProperty("ProcessorNumber", (value))
}

// GetProcessorNumber gets the value of ProcessorNumber for the instance
func (instance *SystemConfig_V2_ProcNumber) GetPropertyProcessorNumber() (value []uint32, err error) {
	retValue, err := instance.GetProperty("ProcessorNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}
