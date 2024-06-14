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

// PageFault_HeapRangeTypeGroup struct
type PageFault_HeapRangeTypeGroup struct {
	*PageFault_V2

	//
	HeapHandle uint32

	//
	HRAddress uint32

	//
	HRSize interface{}
}

func NewPageFault_HeapRangeTypeGroupEx1(instance *cim.WmiInstance) (newInstance *PageFault_HeapRangeTypeGroup, err error) {
	tmp, err := NewPageFault_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &PageFault_HeapRangeTypeGroup{
		PageFault_V2: tmp,
	}
	return
}

func NewPageFault_HeapRangeTypeGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PageFault_HeapRangeTypeGroup, err error) {
	tmp, err := NewPageFault_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PageFault_HeapRangeTypeGroup{
		PageFault_V2: tmp,
	}
	return
}

// SetHeapHandle sets the value of HeapHandle for the instance
func (instance *PageFault_HeapRangeTypeGroup) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *PageFault_HeapRangeTypeGroup) GetPropertyHeapHandle() (value uint32, err error) {
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

// SetHRAddress sets the value of HRAddress for the instance
func (instance *PageFault_HeapRangeTypeGroup) SetPropertyHRAddress(value uint32) (err error) {
	return instance.SetProperty("HRAddress", (value))
}

// GetHRAddress gets the value of HRAddress for the instance
func (instance *PageFault_HeapRangeTypeGroup) GetPropertyHRAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("HRAddress")
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

// SetHRSize sets the value of HRSize for the instance
func (instance *PageFault_HeapRangeTypeGroup) SetPropertyHRSize(value interface{}) (err error) {
	return instance.SetProperty("HRSize", (value))
}

// GetHRSize gets the value of HRSize for the instance
func (instance *PageFault_HeapRangeTypeGroup) GetPropertyHRSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("HRSize")
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