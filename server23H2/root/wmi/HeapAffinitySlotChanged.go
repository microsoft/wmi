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

// HeapAffinitySlotChanged struct
type HeapAffinitySlotChanged struct {
	*HeapTrace_V2

	//
	Heap uint32

	//
	SlotIndex uint32

	//
	SubSegment uint32
}

func NewHeapAffinitySlotChangedEx1(instance *cim.WmiInstance) (newInstance *HeapAffinitySlotChanged, err error) {
	tmp, err := NewHeapTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapAffinitySlotChanged{
		HeapTrace_V2: tmp,
	}
	return
}

func NewHeapAffinitySlotChangedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapAffinitySlotChanged, err error) {
	tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapAffinitySlotChanged{
		HeapTrace_V2: tmp,
	}
	return
}

// SetHeap sets the value of Heap for the instance
func (instance *HeapAffinitySlotChanged) SetPropertyHeap(value uint32) (err error) {
	return instance.SetProperty("Heap", (value))
}

// GetHeap gets the value of Heap for the instance
func (instance *HeapAffinitySlotChanged) GetPropertyHeap() (value uint32, err error) {
	retValue, err := instance.GetProperty("Heap")
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

// SetSlotIndex sets the value of SlotIndex for the instance
func (instance *HeapAffinitySlotChanged) SetPropertySlotIndex(value uint32) (err error) {
	return instance.SetProperty("SlotIndex", (value))
}

// GetSlotIndex gets the value of SlotIndex for the instance
func (instance *HeapAffinitySlotChanged) GetPropertySlotIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("SlotIndex")
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
func (instance *HeapAffinitySlotChanged) SetPropertySubSegment(value uint32) (err error) {
	return instance.SetProperty("SubSegment", (value))
}

// GetSubSegment gets the value of SubSegment for the instance
func (instance *HeapAffinitySlotChanged) GetPropertySubSegment() (value uint32, err error) {
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
