// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ProcessorBiosTStates struct
type ProcessorBiosTStates struct {
	*MSProcessorClass

	//
	Active bool

	//
	FadtDutyOffset uint8

	//
	FadtDutyWidth uint8

	//
	InstanceName string

	//
	NtNumber uint32

	//
	Ptc AcpiControlStatus

	//
	Reserved1 uint32

	//
	Reserved2 uint32

	//
	Reserved3 uint64

	//
	Tpc uint32

	//
	Tss ProcessorAcpiTss

	//
	TStateVersionInUse uint32
}

func NewProcessorBiosTStatesEx1(instance *cim.WmiInstance) (newInstance *ProcessorBiosTStates, err error) {
	tmp, err := NewMSProcessorClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProcessorBiosTStates{
		MSProcessorClass: tmp,
	}
	return
}

func NewProcessorBiosTStatesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessorBiosTStates, err error) {
	tmp, err := NewMSProcessorClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessorBiosTStates{
		MSProcessorClass: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *ProcessorBiosTStates) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ProcessorBiosTStates) GetPropertyActive() (value bool, err error) {
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

// SetFadtDutyOffset sets the value of FadtDutyOffset for the instance
func (instance *ProcessorBiosTStates) SetPropertyFadtDutyOffset(value uint8) (err error) {
	return instance.SetProperty("FadtDutyOffset", (value))
}

// GetFadtDutyOffset gets the value of FadtDutyOffset for the instance
func (instance *ProcessorBiosTStates) GetPropertyFadtDutyOffset() (value uint8, err error) {
	retValue, err := instance.GetProperty("FadtDutyOffset")
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

// SetFadtDutyWidth sets the value of FadtDutyWidth for the instance
func (instance *ProcessorBiosTStates) SetPropertyFadtDutyWidth(value uint8) (err error) {
	return instance.SetProperty("FadtDutyWidth", (value))
}

// GetFadtDutyWidth gets the value of FadtDutyWidth for the instance
func (instance *ProcessorBiosTStates) GetPropertyFadtDutyWidth() (value uint8, err error) {
	retValue, err := instance.GetProperty("FadtDutyWidth")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ProcessorBiosTStates) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ProcessorBiosTStates) GetPropertyInstanceName() (value string, err error) {
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
func (instance *ProcessorBiosTStates) SetPropertyNtNumber(value uint32) (err error) {
	return instance.SetProperty("NtNumber", (value))
}

// GetNtNumber gets the value of NtNumber for the instance
func (instance *ProcessorBiosTStates) GetPropertyNtNumber() (value uint32, err error) {
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

// SetPtc sets the value of Ptc for the instance
func (instance *ProcessorBiosTStates) SetPropertyPtc(value AcpiControlStatus) (err error) {
	return instance.SetProperty("Ptc", (value))
}

// GetPtc gets the value of Ptc for the instance
func (instance *ProcessorBiosTStates) GetPropertyPtc() (value AcpiControlStatus, err error) {
	retValue, err := instance.GetProperty("Ptc")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AcpiControlStatus)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AcpiControlStatus is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AcpiControlStatus(valuetmp)

	return
}

// SetReserved1 sets the value of Reserved1 for the instance
func (instance *ProcessorBiosTStates) SetPropertyReserved1(value uint32) (err error) {
	return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *ProcessorBiosTStates) GetPropertyReserved1() (value uint32, err error) {
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
func (instance *ProcessorBiosTStates) SetPropertyReserved2(value uint32) (err error) {
	return instance.SetProperty("Reserved2", (value))
}

// GetReserved2 gets the value of Reserved2 for the instance
func (instance *ProcessorBiosTStates) GetPropertyReserved2() (value uint32, err error) {
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
func (instance *ProcessorBiosTStates) SetPropertyReserved3(value uint64) (err error) {
	return instance.SetProperty("Reserved3", (value))
}

// GetReserved3 gets the value of Reserved3 for the instance
func (instance *ProcessorBiosTStates) GetPropertyReserved3() (value uint64, err error) {
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

// SetTpc sets the value of Tpc for the instance
func (instance *ProcessorBiosTStates) SetPropertyTpc(value uint32) (err error) {
	return instance.SetProperty("Tpc", (value))
}

// GetTpc gets the value of Tpc for the instance
func (instance *ProcessorBiosTStates) GetPropertyTpc() (value uint32, err error) {
	retValue, err := instance.GetProperty("Tpc")
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

// SetTss sets the value of Tss for the instance
func (instance *ProcessorBiosTStates) SetPropertyTss(value ProcessorAcpiTss) (err error) {
	return instance.SetProperty("Tss", (value))
}

// GetTss gets the value of Tss for the instance
func (instance *ProcessorBiosTStates) GetPropertyTss() (value ProcessorAcpiTss, err error) {
	retValue, err := instance.GetProperty("Tss")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ProcessorAcpiTss)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ProcessorAcpiTss is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ProcessorAcpiTss(valuetmp)

	return
}

// SetTStateVersionInUse sets the value of TStateVersionInUse for the instance
func (instance *ProcessorBiosTStates) SetPropertyTStateVersionInUse(value uint32) (err error) {
	return instance.SetProperty("TStateVersionInUse", (value))
}

// GetTStateVersionInUse gets the value of TStateVersionInUse for the instance
func (instance *ProcessorBiosTStates) GetPropertyTStateVersionInUse() (value uint32, err error) {
	retValue, err := instance.GetProperty("TStateVersionInUse")
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
