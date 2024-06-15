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

// ProcessorAcpiCsdDependency struct
type ProcessorAcpiCsdDependency struct {
	*cim.WmiInstance

	//
	CoordType uint32

	//
	Domain uint32

	//
	Index uint32

	//
	NumberOfEntries uint8

	//
	NumProcessors uint32

	//
	Reserved1 uint32

	//
	Reserved2 uint64

	//
	Revision uint8
}

func NewProcessorAcpiCsdDependencyEx1(instance *cim.WmiInstance) (newInstance *ProcessorAcpiCsdDependency, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ProcessorAcpiCsdDependency{
		WmiInstance: tmp,
	}
	return
}

func NewProcessorAcpiCsdDependencyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProcessorAcpiCsdDependency, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProcessorAcpiCsdDependency{
		WmiInstance: tmp,
	}
	return
}

// SetCoordType sets the value of CoordType for the instance
func (instance *ProcessorAcpiCsdDependency) SetPropertyCoordType(value uint32) (err error) {
	return instance.SetProperty("CoordType", (value))
}

// GetCoordType gets the value of CoordType for the instance
func (instance *ProcessorAcpiCsdDependency) GetPropertyCoordType() (value uint32, err error) {
	retValue, err := instance.GetProperty("CoordType")
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

// SetDomain sets the value of Domain for the instance
func (instance *ProcessorAcpiCsdDependency) SetPropertyDomain(value uint32) (err error) {
	return instance.SetProperty("Domain", (value))
}

// GetDomain gets the value of Domain for the instance
func (instance *ProcessorAcpiCsdDependency) GetPropertyDomain() (value uint32, err error) {
	retValue, err := instance.GetProperty("Domain")
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

// SetIndex sets the value of Index for the instance
func (instance *ProcessorAcpiCsdDependency) SetPropertyIndex(value uint32) (err error) {
	return instance.SetProperty("Index", (value))
}

// GetIndex gets the value of Index for the instance
func (instance *ProcessorAcpiCsdDependency) GetPropertyIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("Index")
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

// SetNumberOfEntries sets the value of NumberOfEntries for the instance
func (instance *ProcessorAcpiCsdDependency) SetPropertyNumberOfEntries(value uint8) (err error) {
	return instance.SetProperty("NumberOfEntries", (value))
}

// GetNumberOfEntries gets the value of NumberOfEntries for the instance
func (instance *ProcessorAcpiCsdDependency) GetPropertyNumberOfEntries() (value uint8, err error) {
	retValue, err := instance.GetProperty("NumberOfEntries")
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

// SetNumProcessors sets the value of NumProcessors for the instance
func (instance *ProcessorAcpiCsdDependency) SetPropertyNumProcessors(value uint32) (err error) {
	return instance.SetProperty("NumProcessors", (value))
}

// GetNumProcessors gets the value of NumProcessors for the instance
func (instance *ProcessorAcpiCsdDependency) GetPropertyNumProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumProcessors")
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
func (instance *ProcessorAcpiCsdDependency) SetPropertyReserved1(value uint32) (err error) {
	return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *ProcessorAcpiCsdDependency) GetPropertyReserved1() (value uint32, err error) {
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
func (instance *ProcessorAcpiCsdDependency) SetPropertyReserved2(value uint64) (err error) {
	return instance.SetProperty("Reserved2", (value))
}

// GetReserved2 gets the value of Reserved2 for the instance
func (instance *ProcessorAcpiCsdDependency) GetPropertyReserved2() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reserved2")
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

// SetRevision sets the value of Revision for the instance
func (instance *ProcessorAcpiCsdDependency) SetPropertyRevision(value uint8) (err error) {
	return instance.SetProperty("Revision", (value))
}

// GetRevision gets the value of Revision for the instance
func (instance *ProcessorAcpiCsdDependency) GetPropertyRevision() (value uint8, err error) {
	retValue, err := instance.GetProperty("Revision")
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
