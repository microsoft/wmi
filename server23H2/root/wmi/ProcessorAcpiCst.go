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

// ProcessorAcpiCst struct
type ProcessorAcpiCst struct {
	*cim.WmiInstance

	//
	Count uint32

	//
	State []ProcessorAcpiCstState
}

func NewProcessorAcpiCstEx1(instance *cim.WmiInstance) (newInstance *ProcessorAcpiCst, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ProcessorAcpiCst{
		WmiInstance: tmp,
	}
	return
}

func NewProcessorAcpiCstEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessorAcpiCst, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessorAcpiCst{
		WmiInstance: tmp,
	}
	return
}

// SetCount sets the value of Count for the instance
func (instance *ProcessorAcpiCst) SetPropertyCount(value uint32) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *ProcessorAcpiCst) GetPropertyCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("Count")
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

// SetState sets the value of State for the instance
func (instance *ProcessorAcpiCst) SetPropertyState(value []ProcessorAcpiCstState) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *ProcessorAcpiCst) GetPropertyState() (value []ProcessorAcpiCstState, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ProcessorAcpiCstState)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ProcessorAcpiCstState is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ProcessorAcpiCstState(valuetmp))
	}

	return
}
