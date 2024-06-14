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

// SampledProfile struct
type SampledProfile struct {
	*PerfInfo_V2

	//
	Count uint16

	//
	InstructionPointer uint32

	//
	Reserved uint16

	//
	ThreadId uint32
}

func NewSampledProfileEx1(instance *cim.WmiInstance) (newInstance *SampledProfile, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SampledProfile{
		PerfInfo_V2: tmp,
	}
	return
}

func NewSampledProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SampledProfile, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SampledProfile{
		PerfInfo_V2: tmp,
	}
	return
}

// SetCount sets the value of Count for the instance
func (instance *SampledProfile) SetPropertyCount(value uint16) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *SampledProfile) GetPropertyCount() (value uint16, err error) {
	retValue, err := instance.GetProperty("Count")
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

// SetInstructionPointer sets the value of InstructionPointer for the instance
func (instance *SampledProfile) SetPropertyInstructionPointer(value uint32) (err error) {
	return instance.SetProperty("InstructionPointer", (value))
}

// GetInstructionPointer gets the value of InstructionPointer for the instance
func (instance *SampledProfile) GetPropertyInstructionPointer() (value uint32, err error) {
	retValue, err := instance.GetProperty("InstructionPointer")
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

// SetReserved sets the value of Reserved for the instance
func (instance *SampledProfile) SetPropertyReserved(value uint16) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *SampledProfile) GetPropertyReserved() (value uint16, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetThreadId sets the value of ThreadId for the instance
func (instance *SampledProfile) SetPropertyThreadId(value uint32) (err error) {
	return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *SampledProfile) GetPropertyThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadId")
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
