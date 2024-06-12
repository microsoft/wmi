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

// HeapSubsegmentReuseThresholdActivated struct
type HeapSubsegmentReuseThresholdActivated struct {
	*HeapTrace_V2

	//
	BucketIndex uint32

	//
	Heap uint32

	//
	SubSegment uint32
}

func NewHeapSubsegmentReuseThresholdActivatedEx1(instance *cim.WmiInstance) (newInstance *HeapSubsegmentReuseThresholdActivated, err error) {
	tmp, err := NewHeapTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapSubsegmentReuseThresholdActivated{
		HeapTrace_V2: tmp,
	}
	return
}

func NewHeapSubsegmentReuseThresholdActivatedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapSubsegmentReuseThresholdActivated, err error) {
	tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapSubsegmentReuseThresholdActivated{
		HeapTrace_V2: tmp,
	}
	return
}

// SetBucketIndex sets the value of BucketIndex for the instance
func (instance *HeapSubsegmentReuseThresholdActivated) SetPropertyBucketIndex(value uint32) (err error) {
	return instance.SetProperty("BucketIndex", (value))
}

// GetBucketIndex gets the value of BucketIndex for the instance
func (instance *HeapSubsegmentReuseThresholdActivated) GetPropertyBucketIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("BucketIndex")
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

// SetHeap sets the value of Heap for the instance
func (instance *HeapSubsegmentReuseThresholdActivated) SetPropertyHeap(value uint32) (err error) {
	return instance.SetProperty("Heap", (value))
}

// GetHeap gets the value of Heap for the instance
func (instance *HeapSubsegmentReuseThresholdActivated) GetPropertyHeap() (value uint32, err error) {
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

// SetSubSegment sets the value of SubSegment for the instance
func (instance *HeapSubsegmentReuseThresholdActivated) SetPropertySubSegment(value uint32) (err error) {
	return instance.SetProperty("SubSegment", (value))
}

// GetSubSegment gets the value of SubSegment for the instance
func (instance *HeapSubsegmentReuseThresholdActivated) GetPropertySubSegment() (value uint32, err error) {
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
