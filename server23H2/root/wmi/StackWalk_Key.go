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

// StackWalk_Key struct
type StackWalk_Key struct {
	*StackWalk

	//
	EventTimeStamp uint64

	//
	StackKey uint32

	//
	StackProcess uint32

	//
	StackThread uint32
}

func NewStackWalk_KeyEx1(instance *cim.WmiInstance) (newInstance *StackWalk_Key, err error) {
	tmp, err := NewStackWalkEx1(instance)

	if err != nil {
		return
	}
	newInstance = &StackWalk_Key{
		StackWalk: tmp,
	}
	return
}

func NewStackWalk_KeyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *StackWalk_Key, err error) {
	tmp, err := NewStackWalkEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &StackWalk_Key{
		StackWalk: tmp,
	}
	return
}

// SetEventTimeStamp sets the value of EventTimeStamp for the instance
func (instance *StackWalk_Key) SetPropertyEventTimeStamp(value uint64) (err error) {
	return instance.SetProperty("EventTimeStamp", (value))
}

// GetEventTimeStamp gets the value of EventTimeStamp for the instance
func (instance *StackWalk_Key) GetPropertyEventTimeStamp() (value uint64, err error) {
	retValue, err := instance.GetProperty("EventTimeStamp")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStackKey sets the value of StackKey for the instance
func (instance *StackWalk_Key) SetPropertyStackKey(value uint32) (err error) {
	return instance.SetProperty("StackKey", (value))
}

// GetStackKey gets the value of StackKey for the instance
func (instance *StackWalk_Key) GetPropertyStackKey() (value uint32, err error) {
	retValue, err := instance.GetProperty("StackKey")
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

// SetStackProcess sets the value of StackProcess for the instance
func (instance *StackWalk_Key) SetPropertyStackProcess(value uint32) (err error) {
	return instance.SetProperty("StackProcess", (value))
}

// GetStackProcess gets the value of StackProcess for the instance
func (instance *StackWalk_Key) GetPropertyStackProcess() (value uint32, err error) {
	retValue, err := instance.GetProperty("StackProcess")
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

// SetStackThread sets the value of StackThread for the instance
func (instance *StackWalk_Key) SetPropertyStackThread(value uint32) (err error) {
	return instance.SetProperty("StackThread", (value))
}

// GetStackThread gets the value of StackThread for the instance
func (instance *StackWalk_Key) GetPropertyStackThread() (value uint32, err error) {
	retValue, err := instance.GetProperty("StackThread")
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
