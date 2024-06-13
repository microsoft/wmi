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

// ProcessorBiosCStates struct
type ProcessorBiosCStates struct {
	*MSProcessorClass

	//
	Active bool

	//
	Cst ProcessorAcpiCst

	//
	CStateVersionInUse uint32

	//
	FadtC2Latency uint16

	//
	FadtC3Latency uint16

	//
	InstanceName string

	//
	NtNumber uint32

	//
	Reserved1 uint32

	//
	Reserved2 uint32

	//
	Reserved3 uint64
}

func NewProcessorBiosCStatesEx1(instance *cim.WmiInstance) (newInstance *ProcessorBiosCStates, err error) {
	tmp, err := NewMSProcessorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProcessorBiosCStates{
		MSProcessorClass: tmp,
	}
	return
}

func NewProcessorBiosCStatesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessorBiosCStates, err error) {
	tmp, err := NewMSProcessorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessorBiosCStates{
		MSProcessorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *ProcessorBiosCStates) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ProcessorBiosCStates) GetPropertyActive() (value bool, err error) {
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

// SetCst sets the value of Cst for the instance
func (instance *ProcessorBiosCStates) SetPropertyCst(value ProcessorAcpiCst) (err error) {
	return instance.SetProperty("Cst", (value))
}

// GetCst gets the value of Cst for the instance
func (instance *ProcessorBiosCStates) GetPropertyCst() (value ProcessorAcpiCst, err error) {
	retValue, err := instance.GetProperty("Cst")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ProcessorAcpiCst)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ProcessorAcpiCst is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ProcessorAcpiCst(valuetmp)

	return
}

// SetCStateVersionInUse sets the value of CStateVersionInUse for the instance
func (instance *ProcessorBiosCStates) SetPropertyCStateVersionInUse(value uint32) (err error) {
	return instance.SetProperty("CStateVersionInUse", (value))
}

// GetCStateVersionInUse gets the value of CStateVersionInUse for the instance
func (instance *ProcessorBiosCStates) GetPropertyCStateVersionInUse() (value uint32, err error) {
	retValue, err := instance.GetProperty("CStateVersionInUse")
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

// SetFadtC2Latency sets the value of FadtC2Latency for the instance
func (instance *ProcessorBiosCStates) SetPropertyFadtC2Latency(value uint16) (err error) {
	return instance.SetProperty("FadtC2Latency", (value))
}

// GetFadtC2Latency gets the value of FadtC2Latency for the instance
func (instance *ProcessorBiosCStates) GetPropertyFadtC2Latency() (value uint16, err error) {
	retValue, err := instance.GetProperty("FadtC2Latency")
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

// SetFadtC3Latency sets the value of FadtC3Latency for the instance
func (instance *ProcessorBiosCStates) SetPropertyFadtC3Latency(value uint16) (err error) {
	return instance.SetProperty("FadtC3Latency", (value))
}

// GetFadtC3Latency gets the value of FadtC3Latency for the instance
func (instance *ProcessorBiosCStates) GetPropertyFadtC3Latency() (value uint16, err error) {
	retValue, err := instance.GetProperty("FadtC3Latency")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ProcessorBiosCStates) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ProcessorBiosCStates) GetPropertyInstanceName() (value string, err error) {
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

// SetNtNumber sets the value of NtNumber for the instance
func (instance *ProcessorBiosCStates) SetPropertyNtNumber(value uint32) (err error) {
	return instance.SetProperty("NtNumber", (value))
}

// GetNtNumber gets the value of NtNumber for the instance
func (instance *ProcessorBiosCStates) GetPropertyNtNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("NtNumber")
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

// SetReserved1 sets the value of Reserved1 for the instance
func (instance *ProcessorBiosCStates) SetPropertyReserved1(value uint32) (err error) {
	return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *ProcessorBiosCStates) GetPropertyReserved1() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved1")
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

// SetReserved2 sets the value of Reserved2 for the instance
func (instance *ProcessorBiosCStates) SetPropertyReserved2(value uint32) (err error) {
	return instance.SetProperty("Reserved2", (value))
}

// GetReserved2 gets the value of Reserved2 for the instance
func (instance *ProcessorBiosCStates) GetPropertyReserved2() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved2")
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

// SetReserved3 sets the value of Reserved3 for the instance
func (instance *ProcessorBiosCStates) SetPropertyReserved3(value uint64) (err error) {
	return instance.SetProperty("Reserved3", (value))
}

// GetReserved3 gets the value of Reserved3 for the instance
func (instance *ProcessorBiosCStates) GetPropertyReserved3() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reserved3")
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
