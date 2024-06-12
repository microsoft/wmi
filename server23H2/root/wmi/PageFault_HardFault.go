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

// PageFault_HardFault struct
type PageFault_HardFault struct {
	*PageFault_V2

	//
	ByteCount uint32

	//
	FileObject uint32

	//
	InitialTime interface{}

	//
	ReadOffset uint64

	//
	TThreadId uint32

	//
	VirtualAddress uint32
}

func NewPageFault_HardFaultEx1(instance *cim.WmiInstance) (newInstance *PageFault_HardFault, err error) {
	tmp, err := NewPageFault_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &PageFault_HardFault{
		PageFault_V2: tmp,
	}
	return
}

func NewPageFault_HardFaultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PageFault_HardFault, err error) {
	tmp, err := NewPageFault_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PageFault_HardFault{
		PageFault_V2: tmp,
	}
	return
}

// SetByteCount sets the value of ByteCount for the instance
func (instance *PageFault_HardFault) SetPropertyByteCount(value uint32) (err error) {
	return instance.SetProperty("ByteCount", (value))
}

// GetByteCount gets the value of ByteCount for the instance
func (instance *PageFault_HardFault) GetPropertyByteCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ByteCount")
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

// SetFileObject sets the value of FileObject for the instance
func (instance *PageFault_HardFault) SetPropertyFileObject(value uint32) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *PageFault_HardFault) GetPropertyFileObject() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileObject")
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

// SetInitialTime sets the value of InitialTime for the instance
func (instance *PageFault_HardFault) SetPropertyInitialTime(value interface{}) (err error) {
	return instance.SetProperty("InitialTime", (value))
}

// GetInitialTime gets the value of InitialTime for the instance
func (instance *PageFault_HardFault) GetPropertyInitialTime() (value interface{}, err error) {
	retValue, err := instance.GetProperty("InitialTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetReadOffset sets the value of ReadOffset for the instance
func (instance *PageFault_HardFault) SetPropertyReadOffset(value uint64) (err error) {
	return instance.SetProperty("ReadOffset", (value))
}

// GetReadOffset gets the value of ReadOffset for the instance
func (instance *PageFault_HardFault) GetPropertyReadOffset() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadOffset")
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

// SetTThreadId sets the value of TThreadId for the instance
func (instance *PageFault_HardFault) SetPropertyTThreadId(value uint32) (err error) {
	return instance.SetProperty("TThreadId", (value))
}

// GetTThreadId gets the value of TThreadId for the instance
func (instance *PageFault_HardFault) GetPropertyTThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("TThreadId")
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

// SetVirtualAddress sets the value of VirtualAddress for the instance
func (instance *PageFault_HardFault) SetPropertyVirtualAddress(value uint32) (err error) {
	return instance.SetProperty("VirtualAddress", (value))
}

// GetVirtualAddress gets the value of VirtualAddress for the instance
func (instance *PageFault_HardFault) GetPropertyVirtualAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualAddress")
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
