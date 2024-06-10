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

// Header_Extension_TypeGroup struct
type Header_Extension_TypeGroup struct {
	*EventTraceEvent

	//
	GroupMask1 uint32

	//
	GroupMask2 uint32

	//
	GroupMask3 uint32

	//
	GroupMask4 uint32

	//
	GroupMask5 uint32

	//
	GroupMask6 uint32

	//
	GroupMask7 uint32

	//
	GroupMask8 uint32

	//
	KernelEventVersion uint32
}

func NewHeader_Extension_TypeGroupEx1(instance *cim.WmiInstance) (newInstance *Header_Extension_TypeGroup, err error) {
	tmp, err := NewEventTraceEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Header_Extension_TypeGroup{
		EventTraceEvent: tmp,
	}
	return
}

func NewHeader_Extension_TypeGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Header_Extension_TypeGroup, err error) {
	tmp, err := NewEventTraceEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Header_Extension_TypeGroup{
		EventTraceEvent: tmp,
	}
	return
}

// SetGroupMask1 sets the value of GroupMask1 for the instance
func (instance *Header_Extension_TypeGroup) SetPropertyGroupMask1(value uint32) (err error) {
	return instance.SetProperty("GroupMask1", (value))
}

// GetGroupMask1 gets the value of GroupMask1 for the instance
func (instance *Header_Extension_TypeGroup) GetPropertyGroupMask1() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupMask1")
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

// SetGroupMask2 sets the value of GroupMask2 for the instance
func (instance *Header_Extension_TypeGroup) SetPropertyGroupMask2(value uint32) (err error) {
	return instance.SetProperty("GroupMask2", (value))
}

// GetGroupMask2 gets the value of GroupMask2 for the instance
func (instance *Header_Extension_TypeGroup) GetPropertyGroupMask2() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupMask2")
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

// SetGroupMask3 sets the value of GroupMask3 for the instance
func (instance *Header_Extension_TypeGroup) SetPropertyGroupMask3(value uint32) (err error) {
	return instance.SetProperty("GroupMask3", (value))
}

// GetGroupMask3 gets the value of GroupMask3 for the instance
func (instance *Header_Extension_TypeGroup) GetPropertyGroupMask3() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupMask3")
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

// SetGroupMask4 sets the value of GroupMask4 for the instance
func (instance *Header_Extension_TypeGroup) SetPropertyGroupMask4(value uint32) (err error) {
	return instance.SetProperty("GroupMask4", (value))
}

// GetGroupMask4 gets the value of GroupMask4 for the instance
func (instance *Header_Extension_TypeGroup) GetPropertyGroupMask4() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupMask4")
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

// SetGroupMask5 sets the value of GroupMask5 for the instance
func (instance *Header_Extension_TypeGroup) SetPropertyGroupMask5(value uint32) (err error) {
	return instance.SetProperty("GroupMask5", (value))
}

// GetGroupMask5 gets the value of GroupMask5 for the instance
func (instance *Header_Extension_TypeGroup) GetPropertyGroupMask5() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupMask5")
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

// SetGroupMask6 sets the value of GroupMask6 for the instance
func (instance *Header_Extension_TypeGroup) SetPropertyGroupMask6(value uint32) (err error) {
	return instance.SetProperty("GroupMask6", (value))
}

// GetGroupMask6 gets the value of GroupMask6 for the instance
func (instance *Header_Extension_TypeGroup) GetPropertyGroupMask6() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupMask6")
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

// SetGroupMask7 sets the value of GroupMask7 for the instance
func (instance *Header_Extension_TypeGroup) SetPropertyGroupMask7(value uint32) (err error) {
	return instance.SetProperty("GroupMask7", (value))
}

// GetGroupMask7 gets the value of GroupMask7 for the instance
func (instance *Header_Extension_TypeGroup) GetPropertyGroupMask7() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupMask7")
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

// SetGroupMask8 sets the value of GroupMask8 for the instance
func (instance *Header_Extension_TypeGroup) SetPropertyGroupMask8(value uint32) (err error) {
	return instance.SetProperty("GroupMask8", (value))
}

// GetGroupMask8 gets the value of GroupMask8 for the instance
func (instance *Header_Extension_TypeGroup) GetPropertyGroupMask8() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupMask8")
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

// SetKernelEventVersion sets the value of KernelEventVersion for the instance
func (instance *Header_Extension_TypeGroup) SetPropertyKernelEventVersion(value uint32) (err error) {
	return instance.SetProperty("KernelEventVersion", (value))
}

// GetKernelEventVersion gets the value of KernelEventVersion for the instance
func (instance *Header_Extension_TypeGroup) GetPropertyKernelEventVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("KernelEventVersion")
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
