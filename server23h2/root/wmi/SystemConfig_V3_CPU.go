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

// SystemConfig_V3_CPU struct
type SystemConfig_V3_CPU struct {
	*SystemConfig_V3

	//
	AllocationGranularity uint32

	//
	ComputerName []byte

	//
	DomainName []byte

	//
	HighestUserAddress uint32

	//
	HyperThreadingFlag uint32

	//
	MemorySpeed uint32

	//
	MemSize uint32

	//
	MHz uint32

	//
	NumberOfProcessors uint32

	//
	NxEnabled uint8

	//
	PaeEnabled uint8

	//
	PageSize uint32

	//
	ProcessorArchitecture uint16

	//
	ProcessorLevel uint16

	//
	ProcessorRevision uint16
}

func NewSystemConfig_V3_CPUEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V3_CPU, err error) {
	tmp, err := NewSystemConfig_V3Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V3_CPU{
		SystemConfig_V3: tmp,
	}
	return
}

func NewSystemConfig_V3_CPUEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V3_CPU, err error) {
	tmp, err := NewSystemConfig_V3Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V3_CPU{
		SystemConfig_V3: tmp,
	}
	return
}

// SetAllocationGranularity sets the value of AllocationGranularity for the instance
func (instance *SystemConfig_V3_CPU) SetPropertyAllocationGranularity(value uint32) (err error) {
	return instance.SetProperty("AllocationGranularity", (value))
}

// GetAllocationGranularity gets the value of AllocationGranularity for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyAllocationGranularity() (value uint32, err error) {
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
func (instance *SystemConfig_V3_CPU) SetPropertyComputerName(value []byte) (err error) {
	return instance.SetProperty("ComputerName", (value))
}

// GetComputerName gets the value of ComputerName for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyComputerName() (value []byte, err error) {
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
func (instance *SystemConfig_V3_CPU) SetPropertyDomainName(value []byte) (err error) {
	return instance.SetProperty("DomainName", (value))
}

// GetDomainName gets the value of DomainName for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyDomainName() (value []byte, err error) {
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

// SetHighestUserAddress sets the value of HighestUserAddress for the instance
func (instance *SystemConfig_V3_CPU) SetPropertyHighestUserAddress(value uint32) (err error) {
	return instance.SetProperty("HighestUserAddress", (value))
}

// GetHighestUserAddress gets the value of HighestUserAddress for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyHighestUserAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("HighestUserAddress")
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

// SetHyperThreadingFlag sets the value of HyperThreadingFlag for the instance
func (instance *SystemConfig_V3_CPU) SetPropertyHyperThreadingFlag(value uint32) (err error) {
	return instance.SetProperty("HyperThreadingFlag", (value))
}

// GetHyperThreadingFlag gets the value of HyperThreadingFlag for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyHyperThreadingFlag() (value uint32, err error) {
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

// SetMemorySpeed sets the value of MemorySpeed for the instance
func (instance *SystemConfig_V3_CPU) SetPropertyMemorySpeed(value uint32) (err error) {
	return instance.SetProperty("MemorySpeed", (value))
}

// GetMemorySpeed gets the value of MemorySpeed for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyMemorySpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("MemorySpeed")
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
func (instance *SystemConfig_V3_CPU) SetPropertyMemSize(value uint32) (err error) {
	return instance.SetProperty("MemSize", (value))
}

// GetMemSize gets the value of MemSize for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyMemSize() (value uint32, err error) {
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
func (instance *SystemConfig_V3_CPU) SetPropertyMHz(value uint32) (err error) {
	return instance.SetProperty("MHz", (value))
}

// GetMHz gets the value of MHz for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyMHz() (value uint32, err error) {
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
func (instance *SystemConfig_V3_CPU) SetPropertyNumberOfProcessors(value uint32) (err error) {
	return instance.SetProperty("NumberOfProcessors", (value))
}

// GetNumberOfProcessors gets the value of NumberOfProcessors for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyNumberOfProcessors() (value uint32, err error) {
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

// SetNxEnabled sets the value of NxEnabled for the instance
func (instance *SystemConfig_V3_CPU) SetPropertyNxEnabled(value uint8) (err error) {
	return instance.SetProperty("NxEnabled", (value))
}

// GetNxEnabled gets the value of NxEnabled for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyNxEnabled() (value uint8, err error) {
	retValue, err := instance.GetProperty("NxEnabled")
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

// SetPaeEnabled sets the value of PaeEnabled for the instance
func (instance *SystemConfig_V3_CPU) SetPropertyPaeEnabled(value uint8) (err error) {
	return instance.SetProperty("PaeEnabled", (value))
}

// GetPaeEnabled gets the value of PaeEnabled for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyPaeEnabled() (value uint8, err error) {
	retValue, err := instance.GetProperty("PaeEnabled")
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

// SetPageSize sets the value of PageSize for the instance
func (instance *SystemConfig_V3_CPU) SetPropertyPageSize(value uint32) (err error) {
	return instance.SetProperty("PageSize", (value))
}

// GetPageSize gets the value of PageSize for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyPageSize() (value uint32, err error) {
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

// SetProcessorArchitecture sets the value of ProcessorArchitecture for the instance
func (instance *SystemConfig_V3_CPU) SetPropertyProcessorArchitecture(value uint16) (err error) {
	return instance.SetProperty("ProcessorArchitecture", (value))
}

// GetProcessorArchitecture gets the value of ProcessorArchitecture for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyProcessorArchitecture() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorArchitecture")
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

// SetProcessorLevel sets the value of ProcessorLevel for the instance
func (instance *SystemConfig_V3_CPU) SetPropertyProcessorLevel(value uint16) (err error) {
	return instance.SetProperty("ProcessorLevel", (value))
}

// GetProcessorLevel gets the value of ProcessorLevel for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyProcessorLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorLevel")
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

// SetProcessorRevision sets the value of ProcessorRevision for the instance
func (instance *SystemConfig_V3_CPU) SetPropertyProcessorRevision(value uint16) (err error) {
	return instance.SetProperty("ProcessorRevision", (value))
}

// GetProcessorRevision gets the value of ProcessorRevision for the instance
func (instance *SystemConfig_V3_CPU) GetPropertyProcessorRevision() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorRevision")
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
