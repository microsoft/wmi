// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HeapFree struct
type HeapFree struct {
	*HeapTrace_V2

	//
	FreeAddress uint32

	//
	HeapHandle uint32

	//
	SourceId uint32
}

func NewHeapFreeEx1(instance *cim.WmiInstance) (newInstance *HeapFree, err error) {
	tmp, err := NewHeapTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapFree{
		HeapTrace_V2: tmp,
	}
	return
}

func NewHeapFreeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapFree, err error) {
	tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapFree{
		HeapTrace_V2: tmp,
	}
	return
}

// SetFreeAddress sets the value of FreeAddress for the instance
func (instance *HeapFree) SetPropertyFreeAddress(value uint32) (err error) {
	return instance.SetProperty("FreeAddress", (value))
}

// GetFreeAddress gets the value of FreeAddress for the instance
func (instance *HeapFree) GetPropertyFreeAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("FreeAddress")
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
func (instance *HeapFree) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapFree) GetPropertyHeapHandle() (value uint32, err error) {
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
func (instance *HeapFree) SetPropertySourceId(value uint32) (err error) {
	return instance.SetProperty("SourceId", (value))
}

// GetSourceId gets the value of SourceId for the instance
func (instance *HeapFree) GetPropertySourceId() (value uint32, err error) {
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
