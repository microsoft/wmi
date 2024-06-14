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

// HeapSnapShot struct
type HeapSnapShot struct {
	*HeapTrace

	//
	AllocatedSpace interface{}

	//
	CommittedSpace interface{}

	//
	FreeListLength uint32

	//
	FreeSpace interface{}

	//
	HeapFlags uint32

	//
	HeapHandle uint32

	//
	LargeUCRSpace interface{}

	//
	ProcessId uint32

	//
	ReservedSpace interface{}

	//
	UCRLength uint32
}

func NewHeapSnapShotEx1(instance *cim.WmiInstance) (newInstance *HeapSnapShot, err error) {
	tmp, err := NewHeapTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapSnapShot{
		HeapTrace: tmp,
	}
	return
}

func NewHeapSnapShotEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapSnapShot, err error) {
	tmp, err := NewHeapTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapSnapShot{
		HeapTrace: tmp,
	}
	return
}

// SetAllocatedSpace sets the value of AllocatedSpace for the instance
func (instance *HeapSnapShot) SetPropertyAllocatedSpace(value interface{}) (err error) {
	return instance.SetProperty("AllocatedSpace", (value))
}

// GetAllocatedSpace gets the value of AllocatedSpace for the instance
func (instance *HeapSnapShot) GetPropertyAllocatedSpace() (value interface{}, err error) {
	retValue, err := instance.GetProperty("AllocatedSpace")
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

// SetCommittedSpace sets the value of CommittedSpace for the instance
func (instance *HeapSnapShot) SetPropertyCommittedSpace(value interface{}) (err error) {
	return instance.SetProperty("CommittedSpace", (value))
}

// GetCommittedSpace gets the value of CommittedSpace for the instance
func (instance *HeapSnapShot) GetPropertyCommittedSpace() (value interface{}, err error) {
	retValue, err := instance.GetProperty("CommittedSpace")
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

// SetFreeListLength sets the value of FreeListLength for the instance
func (instance *HeapSnapShot) SetPropertyFreeListLength(value uint32) (err error) {
	return instance.SetProperty("FreeListLength", (value))
}

// GetFreeListLength gets the value of FreeListLength for the instance
func (instance *HeapSnapShot) GetPropertyFreeListLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("FreeListLength")
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

// SetFreeSpace sets the value of FreeSpace for the instance
func (instance *HeapSnapShot) SetPropertyFreeSpace(value interface{}) (err error) {
	return instance.SetProperty("FreeSpace", (value))
}

// GetFreeSpace gets the value of FreeSpace for the instance
func (instance *HeapSnapShot) GetPropertyFreeSpace() (value interface{}, err error) {
	retValue, err := instance.GetProperty("FreeSpace")
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

// SetHeapFlags sets the value of HeapFlags for the instance
func (instance *HeapSnapShot) SetPropertyHeapFlags(value uint32) (err error) {
	return instance.SetProperty("HeapFlags", (value))
}

// GetHeapFlags gets the value of HeapFlags for the instance
func (instance *HeapSnapShot) GetPropertyHeapFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("HeapFlags")
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

// SetHeapHandle sets the value of HeapHandle for the instance
func (instance *HeapSnapShot) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapSnapShot) GetPropertyHeapHandle() (value uint32, err error) {
	retValue, err := instance.GetProperty("HeapHandle")
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

// SetLargeUCRSpace sets the value of LargeUCRSpace for the instance
func (instance *HeapSnapShot) SetPropertyLargeUCRSpace(value interface{}) (err error) {
	return instance.SetProperty("LargeUCRSpace", (value))
}

// GetLargeUCRSpace gets the value of LargeUCRSpace for the instance
func (instance *HeapSnapShot) GetPropertyLargeUCRSpace() (value interface{}, err error) {
	retValue, err := instance.GetProperty("LargeUCRSpace")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *HeapSnapShot) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *HeapSnapShot) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetReservedSpace sets the value of ReservedSpace for the instance
func (instance *HeapSnapShot) SetPropertyReservedSpace(value interface{}) (err error) {
	return instance.SetProperty("ReservedSpace", (value))
}

// GetReservedSpace gets the value of ReservedSpace for the instance
func (instance *HeapSnapShot) GetPropertyReservedSpace() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ReservedSpace")
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

// SetUCRLength sets the value of UCRLength for the instance
func (instance *HeapSnapShot) SetPropertyUCRLength(value uint32) (err error) {
	return instance.SetProperty("UCRLength", (value))
}

// GetUCRLength gets the value of UCRLength for the instance
func (instance *HeapSnapShot) GetPropertyUCRLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("UCRLength")
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
