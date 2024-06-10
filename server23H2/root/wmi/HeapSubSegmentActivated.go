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

// HeapSubSegmentActivated struct
type HeapSubSegmentActivated struct {
	*HeapTrace_V2

	//
	HeapHandle uint32

	//
	SubSegment uint32
}

func NewHeapSubSegmentActivatedEx1(instance *cim.WmiInstance) (newInstance *HeapSubSegmentActivated, err error) {
	tmp, err := NewHeapTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapSubSegmentActivated{
		HeapTrace_V2: tmp,
	}
	return
}

func NewHeapSubSegmentActivatedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapSubSegmentActivated, err error) {
	tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapSubSegmentActivated{
		HeapTrace_V2: tmp,
	}
	return
}

// SetHeapHandle sets the value of HeapHandle for the instance
func (instance *HeapSubSegmentActivated) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapSubSegmentActivated) GetPropertyHeapHandle() (value uint32, err error) {
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
func (instance *HeapSubSegmentActivated) SetPropertySubSegment(value uint32) (err error) {
	return instance.SetProperty("SubSegment", (value))
}

// GetSubSegment gets the value of SubSegment for the instance
func (instance *HeapSubSegmentActivated) GetPropertySubSegment() (value uint32, err error) {
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
