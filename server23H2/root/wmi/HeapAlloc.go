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

// HeapAlloc struct
type HeapAlloc struct {
	*HeapTrace_V2

	//
	AllocAddress uint32

	//
	AllocSize interface{}

	//
	HeapHandle uint32

	//
	SourceId uint32
}

func NewHeapAllocEx1(instance *cim.WmiInstance) (newInstance *HeapAlloc, err error) {
	tmp, err := NewHeapTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapAlloc{
		HeapTrace_V2: tmp,
	}
	return
}

func NewHeapAllocEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapAlloc, err error) {
	tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapAlloc{
		HeapTrace_V2: tmp,
	}
	return
}

// SetAllocAddress sets the value of AllocAddress for the instance
func (instance *HeapAlloc) SetPropertyAllocAddress(value uint32) (err error) {
	return instance.SetProperty("AllocAddress", (value))
}

// GetAllocAddress gets the value of AllocAddress for the instance
func (instance *HeapAlloc) GetPropertyAllocAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocAddress")
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

// SetAllocSize sets the value of AllocSize for the instance
func (instance *HeapAlloc) SetPropertyAllocSize(value interface{}) (err error) {
	return instance.SetProperty("AllocSize", (value))
}

// GetAllocSize gets the value of AllocSize for the instance
func (instance *HeapAlloc) GetPropertyAllocSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("AllocSize")
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

// SetHeapHandle sets the value of HeapHandle for the instance
func (instance *HeapAlloc) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapAlloc) GetPropertyHeapHandle() (value uint32, err error) {
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

// SetSourceId sets the value of SourceId for the instance
func (instance *HeapAlloc) SetPropertySourceId(value uint32) (err error) {
	return instance.SetProperty("SourceId", (value))
}

// GetSourceId gets the value of SourceId for the instance
func (instance *HeapAlloc) GetPropertySourceId() (value uint32, err error) {
	retValue, err := instance.GetProperty("SourceId")
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
