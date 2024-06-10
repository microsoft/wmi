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

// HeapSubsegmentInitialize struct
type HeapSubsegmentInitialize struct {
	*HeapTrace_V2

	//
	AffinityIndex uint32

	//
	BlockCount interface{}

	//
	BlockSize interface{}

	//
	HeapHandle uint32

	//
	SubSegment uint32
}

func NewHeapSubsegmentInitializeEx1(instance *cim.WmiInstance) (newInstance *HeapSubsegmentInitialize, err error) {
	tmp, err := NewHeapTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapSubsegmentInitialize{
		HeapTrace_V2: tmp,
	}
	return
}

func NewHeapSubsegmentInitializeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapSubsegmentInitialize, err error) {
	tmp, err := NewHeapTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapSubsegmentInitialize{
		HeapTrace_V2: tmp,
	}
	return
}

// SetAffinityIndex sets the value of AffinityIndex for the instance
func (instance *HeapSubsegmentInitialize) SetPropertyAffinityIndex(value uint32) (err error) {
	return instance.SetProperty("AffinityIndex", (value))
}

// GetAffinityIndex gets the value of AffinityIndex for the instance
func (instance *HeapSubsegmentInitialize) GetPropertyAffinityIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("AffinityIndex")
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

// SetBlockCount sets the value of BlockCount for the instance
func (instance *HeapSubsegmentInitialize) SetPropertyBlockCount(value interface{}) (err error) {
	return instance.SetProperty("BlockCount", (value))
}

// GetBlockCount gets the value of BlockCount for the instance
func (instance *HeapSubsegmentInitialize) GetPropertyBlockCount() (value interface{}, err error) {
	retValue, err := instance.GetProperty("BlockCount")
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

// SetBlockSize sets the value of BlockSize for the instance
func (instance *HeapSubsegmentInitialize) SetPropertyBlockSize(value interface{}) (err error) {
	return instance.SetProperty("BlockSize", (value))
}

// GetBlockSize gets the value of BlockSize for the instance
func (instance *HeapSubsegmentInitialize) GetPropertyBlockSize() (value interface{}, err error) {
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
func (instance *HeapSubsegmentInitialize) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapSubsegmentInitialize) GetPropertyHeapHandle() (value uint32, err error) {
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
func (instance *HeapSubsegmentInitialize) SetPropertySubSegment(value uint32) (err error) {
	return instance.SetProperty("SubSegment", (value))
}

// GetSubSegment gets the value of SubSegment for the instance
func (instance *HeapSubsegmentInitialize) GetPropertySubSegment() (value uint32, err error) {
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
