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

// Heap_SubsegmentGroup struct
type Heap_SubsegmentGroup struct {
	*HeapTrace_V2

	//
	BlockSize interface{}

	//
	HeapHandle uint32

	//
	SubSegment uint32

	//
	SubSegmentSize interface{}
}

func NewHeap_SubsegmentGroupEx1(instance *cim.WmiInstance) (newInstance *Heap_SubsegmentGroup, err error) {
	tmp, err := NewHeapTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Heap_SubsegmentGroup{
		HeapTrace_V2: tmp,
	}
	return
}

func NewHeap_SubsegmentGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Heap_SubsegmentGroup, err error) {
	tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Heap_SubsegmentGroup{
		HeapTrace_V2: tmp,
	}
	return
}

// SetBlockSize sets the value of BlockSize for the instance
func (instance *Heap_SubsegmentGroup) SetPropertyBlockSize(value interface{}) (err error) {
	return instance.SetProperty("BlockSize", (value))
}

// GetBlockSize gets the value of BlockSize for the instance
func (instance *Heap_SubsegmentGroup) GetPropertyBlockSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("BlockSize")
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
func (instance *Heap_SubsegmentGroup) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *Heap_SubsegmentGroup) GetPropertyHeapHandle() (value uint32, err error) {
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

// SetSubSegment sets the value of SubSegment for the instance
func (instance *Heap_SubsegmentGroup) SetPropertySubSegment(value uint32) (err error) {
	return instance.SetProperty("SubSegment", (value))
}

// GetSubSegment gets the value of SubSegment for the instance
func (instance *Heap_SubsegmentGroup) GetPropertySubSegment() (value uint32, err error) {
	retValue, err := instance.GetProperty("SubSegment")
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

// SetSubSegmentSize sets the value of SubSegmentSize for the instance
func (instance *Heap_SubsegmentGroup) SetPropertySubSegmentSize(value interface{}) (err error) {
	return instance.SetProperty("SubSegmentSize", (value))
}

// GetSubSegmentSize gets the value of SubSegmentSize for the instance
func (instance *Heap_SubsegmentGroup) GetPropertySubSegmentSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("SubSegmentSize")
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
