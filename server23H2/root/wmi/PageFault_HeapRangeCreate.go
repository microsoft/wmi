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

// PageFault_HeapRangeCreate struct
type PageFault_HeapRangeCreate struct {
	*PageFault_V2

	//
	FirstRangeSize interface{}

	//
	HeapHandle uint32

	//
	HRCreateFlags uint32
}

func NewPageFault_HeapRangeCreateEx1(instance *cim.WmiInstance) (newInstance *PageFault_HeapRangeCreate, err error) {
	tmp, err := NewPageFault_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &PageFault_HeapRangeCreate{
		PageFault_V2: tmp,
	}
	return
}

func NewPageFault_HeapRangeCreateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PageFault_HeapRangeCreate, err error) {
	tmp, err := NewPageFault_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PageFault_HeapRangeCreate{
		PageFault_V2: tmp,
	}
	return
}

// SetFirstRangeSize sets the value of FirstRangeSize for the instance
func (instance *PageFault_HeapRangeCreate) SetPropertyFirstRangeSize(value interface{}) (err error) {
	return instance.SetProperty("FirstRangeSize", (value))
}

// GetFirstRangeSize gets the value of FirstRangeSize for the instance
func (instance *PageFault_HeapRangeCreate) GetPropertyFirstRangeSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("FirstRangeSize")
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
func (instance *PageFault_HeapRangeCreate) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *PageFault_HeapRangeCreate) GetPropertyHeapHandle() (value uint32, err error) {
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

// SetHRCreateFlags sets the value of HRCreateFlags for the instance
func (instance *PageFault_HeapRangeCreate) SetPropertyHRCreateFlags(value uint32) (err error) {
	return instance.SetProperty("HRCreateFlags", (value))
}

// GetHRCreateFlags gets the value of HRCreateFlags for the instance
func (instance *PageFault_HeapRangeCreate) GetPropertyHRCreateFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("HRCreateFlags")
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
