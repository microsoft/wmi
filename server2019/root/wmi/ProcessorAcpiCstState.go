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

// ProcessorAcpiCstState struct
type ProcessorAcpiCstState struct {
	*cim.WmiInstance

	//
	Latency uint16

	//
	PowerConsumption uint32

	//
	Register AcpiGenAddr

	//
	StateType uint8
}

func NewProcessorAcpiCstStateEx1(instance *cim.WmiInstance) (newInstance *ProcessorAcpiCstState, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ProcessorAcpiCstState{
		WmiInstance: tmp,
	}
	return
}

func NewProcessorAcpiCstStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessorAcpiCstState, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessorAcpiCstState{
		WmiInstance: tmp,
	}
	return
}

// SetLatency sets the value of Latency for the instance
func (instance *ProcessorAcpiCstState) SetPropertyLatency(value uint16) (err error) {
	return instance.SetProperty("Latency", (value))
}

// GetLatency gets the value of Latency for the instance
func (instance *ProcessorAcpiCstState) GetPropertyLatency() (value uint16, err error) {
	retValue, err := instance.GetProperty("Latency")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetPowerConsumption sets the value of PowerConsumption for the instance
func (instance *ProcessorAcpiCstState) SetPropertyPowerConsumption(value uint32) (err error) {
	return instance.SetProperty("PowerConsumption", (value))
}

// GetPowerConsumption gets the value of PowerConsumption for the instance
func (instance *ProcessorAcpiCstState) GetPropertyPowerConsumption() (value uint32, err error) {
	retValue, err := instance.GetProperty("PowerConsumption")
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

// SetRegister sets the value of Register for the instance
func (instance *ProcessorAcpiCstState) SetPropertyRegister(value AcpiGenAddr) (err error) {
	return instance.SetProperty("Register", (value))
}

// GetRegister gets the value of Register for the instance
func (instance *ProcessorAcpiCstState) GetPropertyRegister() (value AcpiGenAddr, err error) {
	retValue, err := instance.GetProperty("Register")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AcpiGenAddr)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AcpiGenAddr is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AcpiGenAddr(valuetmp)

	return
}

// SetStateType sets the value of StateType for the instance
func (instance *ProcessorAcpiCstState) SetPropertyStateType(value uint8) (err error) {
	return instance.SetProperty("StateType", (value))
}

// GetStateType gets the value of StateType for the instance
func (instance *ProcessorAcpiCstState) GetPropertyStateType() (value uint8, err error) {
	retValue, err := instance.GetProperty("StateType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
