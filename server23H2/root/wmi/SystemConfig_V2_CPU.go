// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SystemConfig_V2_CPU struct
type SystemConfig_V2_CPU struct {
	*SystemConfig_V2

	//
	AllocationGranularity uint32

	//
	ComputerName []byte

	//
	DomainName []byte

	//
	HyperThreadingFlag uint32

	//
	MemSize uint32

	//
	MHz uint32

	//
	NumberOfProcessors uint32

	//
	PageSize uint32
}

func NewSystemConfig_V2_CPUEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_CPU, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_CPU{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_CPUEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_CPU, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_CPU{
		SystemConfig_V2: tmp,
	}
	return
}

// SetAllocationGranularity sets the value of AllocationGranularity for the instance
func (instance *SystemConfig_V2_CPU) SetPropertyAllocationGranularity(value uint32) (err error) {
	return instance.SetProperty("AllocationGranularity", (value))
}

// GetAllocationGranularity gets the value of AllocationGranularity for the instance
func (instance *SystemConfig_V2_CPU) GetPropertyAllocationGranularity() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationGranularity")
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

// SetComputerName sets the value of ComputerName for the instance
func (instance *SystemConfig_V2_CPU) SetPropertyComputerName(value []byte) (err error) {
	return instance.SetProperty("ComputerName", (value))
}

// GetComputerName gets the value of ComputerName for the instance
func (instance *SystemConfig_V2_CPU) GetPropertyComputerName() (value []byte, err error) {
	retValue, err := instance.GetProperty("ComputerName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}

// SetDomainName sets the value of DomainName for the instance
func (instance *SystemConfig_V2_CPU) SetPropertyDomainName(value []byte) (err error) {
	return instance.SetProperty("DomainName", (value))
}

// GetDomainName gets the value of DomainName for the instance
func (instance *SystemConfig_V2_CPU) GetPropertyDomainName() (value []byte, err error) {
	retValue, err := instance.GetProperty("DomainName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}

// SetHyperThreadingFlag sets the value of HyperThreadingFlag for the instance
func (instance *SystemConfig_V2_CPU) SetPropertyHyperThreadingFlag(value uint32) (err error) {
	return instance.SetProperty("HyperThreadingFlag", (value))
}

// GetHyperThreadingFlag gets the value of HyperThreadingFlag for the instance
func (instance *SystemConfig_V2_CPU) GetPropertyHyperThreadingFlag() (value uint32, err error) {
	retValue, err := instance.GetProperty("HyperThreadingFlag")
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

// SetMemSize sets the value of MemSize for the instance
func (instance *SystemConfig_V2_CPU) SetPropertyMemSize(value uint32) (err error) {
	return instance.SetProperty("MemSize", (value))
}

// GetMemSize gets the value of MemSize for the instance
func (instance *SystemConfig_V2_CPU) GetPropertyMemSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MemSize")
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

// SetMHz sets the value of MHz for the instance
func (instance *SystemConfig_V2_CPU) SetPropertyMHz(value uint32) (err error) {
	return instance.SetProperty("MHz", (value))
}

// GetMHz gets the value of MHz for the instance
func (instance *SystemConfig_V2_CPU) GetPropertyMHz() (value uint32, err error) {
	retValue, err := instance.GetProperty("MHz")
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

// SetNumberOfProcessors sets the value of NumberOfProcessors for the instance
func (instance *SystemConfig_V2_CPU) SetPropertyNumberOfProcessors(value uint32) (err error) {
	return instance.SetProperty("NumberOfProcessors", (value))
}

// GetNumberOfProcessors gets the value of NumberOfProcessors for the instance
func (instance *SystemConfig_V2_CPU) GetPropertyNumberOfProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfProcessors")
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

// SetPageSize sets the value of PageSize for the instance
func (instance *SystemConfig_V2_CPU) SetPropertyPageSize(value uint32) (err error) {
	return instance.SetProperty("PageSize", (value))
}

// GetPageSize gets the value of PageSize for the instance
func (instance *SystemConfig_V2_CPU) GetPropertyPageSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PageSize")
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
