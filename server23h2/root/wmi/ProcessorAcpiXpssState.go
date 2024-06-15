// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// ProcessorAcpiXpssState struct
type ProcessorAcpiXpssState struct {
	*cim.WmiInstance

	//
	BmLatency uint32

	//
	Control uint64

	//
	ControlMask uint64

	//
	Frequency uint32

	//
	Latency uint32

	//
	Power uint32

	//
	Status uint64

	//
	StatusMask uint64
}

func NewProcessorAcpiXpssStateEx1(instance *cim.WmiInstance) (newInstance *ProcessorAcpiXpssState, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ProcessorAcpiXpssState{
		WmiInstance: tmp,
	}
	return
}

func NewProcessorAcpiXpssStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessorAcpiXpssState, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessorAcpiXpssState{
		WmiInstance: tmp,
	}
	return
}

// SetBmLatency sets the value of BmLatency for the instance
func (instance *ProcessorAcpiXpssState) SetPropertyBmLatency(value uint32) (err error) {
	return instance.SetProperty("BmLatency", (value))
}

// GetBmLatency gets the value of BmLatency for the instance
func (instance *ProcessorAcpiXpssState) GetPropertyBmLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("BmLatency")
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

// SetControl sets the value of Control for the instance
func (instance *ProcessorAcpiXpssState) SetPropertyControl(value uint64) (err error) {
	return instance.SetProperty("Control", (value))
}

// GetControl gets the value of Control for the instance
func (instance *ProcessorAcpiXpssState) GetPropertyControl() (value uint64, err error) {
	retValue, err := instance.GetProperty("Control")
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

// SetControlMask sets the value of ControlMask for the instance
func (instance *ProcessorAcpiXpssState) SetPropertyControlMask(value uint64) (err error) {
	return instance.SetProperty("ControlMask", (value))
}

// GetControlMask gets the value of ControlMask for the instance
func (instance *ProcessorAcpiXpssState) GetPropertyControlMask() (value uint64, err error) {
	retValue, err := instance.GetProperty("ControlMask")
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

// SetFrequency sets the value of Frequency for the instance
func (instance *ProcessorAcpiXpssState) SetPropertyFrequency(value uint32) (err error) {
	return instance.SetProperty("Frequency", (value))
}

// GetFrequency gets the value of Frequency for the instance
func (instance *ProcessorAcpiXpssState) GetPropertyFrequency() (value uint32, err error) {
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

// SetLatency sets the value of Latency for the instance
func (instance *ProcessorAcpiXpssState) SetPropertyLatency(value uint32) (err error) {
	return instance.SetProperty("Latency", (value))
}

// GetLatency gets the value of Latency for the instance
func (instance *ProcessorAcpiXpssState) GetPropertyLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("Latency")
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

// SetPower sets the value of Power for the instance
func (instance *ProcessorAcpiXpssState) SetPropertyPower(value uint32) (err error) {
	return instance.SetProperty("Power", (value))
}

// GetPower gets the value of Power for the instance
func (instance *ProcessorAcpiXpssState) GetPropertyPower() (value uint32, err error) {
	retValue, err := instance.GetProperty("Power")
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

// SetStatus sets the value of Status for the instance
func (instance *ProcessorAcpiXpssState) SetPropertyStatus(value uint64) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *ProcessorAcpiXpssState) GetPropertyStatus() (value uint64, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetStatusMask sets the value of StatusMask for the instance
func (instance *ProcessorAcpiXpssState) SetPropertyStatusMask(value uint64) (err error) {
	return instance.SetProperty("StatusMask", (value))
}

// GetStatusMask gets the value of StatusMask for the instance
func (instance *ProcessorAcpiXpssState) GetPropertyStatusMask() (value uint64, err error) {
	retValue, err := instance.GetProperty("StatusMask")
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
