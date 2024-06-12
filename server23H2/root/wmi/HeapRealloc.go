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

// HeapRealloc struct
type HeapRealloc struct {
	*HeapTrace_V2

	//
	HeapHandle uint32

	//
	NewAllocAddress uint32

	//
	NewAllocSize interface{}

	//
	OldAllocAddress uint32

	//
	OldAllocSize interface{}

	//
	SourceId uint32
}

func NewHeapReallocEx1(instance *cim.WmiInstance) (newInstance *HeapRealloc, err error) {
	tmp, err := NewHeapTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapRealloc{
		HeapTrace_V2: tmp,
	}
	return
}

func NewHeapReallocEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapRealloc, err error) {
	tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapRealloc{
		HeapTrace_V2: tmp,
	}
	return
}

// SetHeapHandle sets the value of HeapHandle for the instance
func (instance *HeapRealloc) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapRealloc) GetPropertyHeapHandle() (value uint32, err error) {
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

// SetNewAllocAddress sets the value of NewAllocAddress for the instance
func (instance *HeapRealloc) SetPropertyNewAllocAddress(value uint32) (err error) {
	return instance.SetProperty("NewAllocAddress", (value))
}

// GetNewAllocAddress gets the value of NewAllocAddress for the instance
func (instance *HeapRealloc) GetPropertyNewAllocAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("NewAllocAddress")
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

// SetNewAllocSize sets the value of NewAllocSize for the instance
func (instance *HeapRealloc) SetPropertyNewAllocSize(value interface{}) (err error) {
	return instance.SetProperty("NewAllocSize", (value))
}

// GetNewAllocSize gets the value of NewAllocSize for the instance
func (instance *HeapRealloc) GetPropertyNewAllocSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("NewAllocSize")
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

// SetOldAllocAddress sets the value of OldAllocAddress for the instance
func (instance *HeapRealloc) SetPropertyOldAllocAddress(value uint32) (err error) {
	return instance.SetProperty("OldAllocAddress", (value))
}

// GetOldAllocAddress gets the value of OldAllocAddress for the instance
func (instance *HeapRealloc) GetPropertyOldAllocAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("OldAllocAddress")
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

// SetOldAllocSize sets the value of OldAllocSize for the instance
func (instance *HeapRealloc) SetPropertyOldAllocSize(value interface{}) (err error) {
	return instance.SetProperty("OldAllocSize", (value))
}

// GetOldAllocSize gets the value of OldAllocSize for the instance
func (instance *HeapRealloc) GetPropertyOldAllocSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("OldAllocSize")
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

// SetSourceId sets the value of SourceId for the instance
func (instance *HeapRealloc) SetPropertySourceId(value uint32) (err error) {
	return instance.SetProperty("SourceId", (value))
}

// GetSourceId gets the value of SourceId for the instance
func (instance *HeapRealloc) GetPropertySourceId() (value uint32, err error) {
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
