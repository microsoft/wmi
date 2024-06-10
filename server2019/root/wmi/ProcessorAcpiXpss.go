// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ProcessorAcpiXpss struct
type ProcessorAcpiXpss struct {
	*cim.WmiInstance

	//
	Active bool

	//
	Count uint32

	//
	InstanceName string

	//
	State []ProcessorAcpiXpssState
}

func NewProcessorAcpiXpssEx1(instance *cim.WmiInstance) (newInstance *ProcessorAcpiXpss, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ProcessorAcpiXpss{
		WmiInstance: tmp,
	}
	return
}

func NewProcessorAcpiXpssEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessorAcpiXpss, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessorAcpiXpss{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *ProcessorAcpiXpss) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ProcessorAcpiXpss) GetPropertyActive() (value bool, err error) {
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

// SetCount sets the value of Count for the instance
func (instance *ProcessorAcpiXpss) SetPropertyCount(value uint32) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *ProcessorAcpiXpss) GetPropertyCount() (value uint32, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ProcessorAcpiXpss) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ProcessorAcpiXpss) GetPropertyInstanceName() (value string, err error) {
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

// SetState sets the value of State for the instance
func (instance *ProcessorAcpiXpss) SetPropertyState(value []ProcessorAcpiXpssState) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *ProcessorAcpiXpss) GetPropertyState() (value []ProcessorAcpiXpssState, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ProcessorAcpiXpssState)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ProcessorAcpiXpssState is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ProcessorAcpiXpssState(valuetmp))
	}

	return
}
