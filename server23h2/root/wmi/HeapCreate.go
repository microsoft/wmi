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

// HeapCreate struct
type HeapCreate struct {
	*HeapTrace

	//
	AllocatedSpace interface{}

	//
	CommittedSpace interface{}

	//
	HeapFlags uint32

	//
	HeapHandle uint32

	//
	ReservedSpace interface{}
}

func NewHeapCreateEx1(instance *cim.WmiInstance) (newInstance *HeapCreate, err error) {
	tmp, err := NewHeapTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapCreate{
		HeapTrace: tmp,
	}
	return
}

func NewHeapCreateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapCreate, err error) {
	tmp, err := NewHeapTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapCreate{
		HeapTrace: tmp,
	}
	return
}

// SetAllocatedSpace sets the value of AllocatedSpace for the instance
func (instance *HeapCreate) SetPropertyAllocatedSpace(value interface{}) (err error) {
	return instance.SetProperty("AllocatedSpace", (value))
}

// GetAllocatedSpace gets the value of AllocatedSpace for the instance
func (instance *HeapCreate) GetPropertyAllocatedSpace() (value interface{}, err error) {
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
func (instance *HeapCreate) SetPropertyCommittedSpace(value interface{}) (err error) {
	return instance.SetProperty("CommittedSpace", (value))
}

// GetCommittedSpace gets the value of CommittedSpace for the instance
func (instance *HeapCreate) GetPropertyCommittedSpace() (value interface{}, err error) {
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

// SetHeapFlags sets the value of HeapFlags for the instance
func (instance *HeapCreate) SetPropertyHeapFlags(value uint32) (err error) {
	return instance.SetProperty("HeapFlags", (value))
}

// GetHeapFlags gets the value of HeapFlags for the instance
func (instance *HeapCreate) GetPropertyHeapFlags() (value uint32, err error) {
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
func (instance *HeapCreate) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapCreate) GetPropertyHeapHandle() (value uint32, err error) {
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

// SetReservedSpace sets the value of ReservedSpace for the instance
func (instance *HeapCreate) SetPropertyReservedSpace(value interface{}) (err error) {
	return instance.SetProperty("ReservedSpace", (value))
}

// GetReservedSpace gets the value of ReservedSpace for the instance
func (instance *HeapCreate) GetPropertyReservedSpace() (value interface{}, err error) {
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
