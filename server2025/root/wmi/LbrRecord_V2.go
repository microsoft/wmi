// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// LbrRecord_V2 struct
type LbrRecord_V2 struct {
	*LBR_V2

	//
	Entries []LbrRecordEntry

	//
	Options uint32

	//
	ProcessId uint32

	//
	ThreadId uint32

	//
	TimeStamp uint64
}

func NewLbrRecord_V2Ex1(instance *cim.WmiInstance) (newInstance *LbrRecord_V2, err error) {
	tmp, err := NewLBR_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &LbrRecord_V2{
		LBR_V2: tmp,
	}
	return
}

func NewLbrRecord_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LbrRecord_V2, err error) {
	tmp, err := NewLBR_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LbrRecord_V2{
		LBR_V2: tmp,
	}
	return
}

// SetEntries sets the value of Entries for the instance
func (instance *LbrRecord_V2) SetPropertyEntries(value []LbrRecordEntry) (err error) {
	return instance.SetProperty("Entries", (value))
}

// GetEntries gets the value of Entries for the instance
func (instance *LbrRecord_V2) GetPropertyEntries() (value []LbrRecordEntry, err error) {
	retValue, err := instance.GetProperty("Entries")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(LbrRecordEntry)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " LbrRecordEntry is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, LbrRecordEntry(valuetmp))
	}

	return
}

// SetOptions sets the value of Options for the instance
func (instance *LbrRecord_V2) SetPropertyOptions(value uint32) (err error) {
	return instance.SetProperty("Options", (value))
}

// GetOptions gets the value of Options for the instance
func (instance *LbrRecord_V2) GetPropertyOptions() (value uint32, err error) {
	retValue, err := instance.GetProperty("Options")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *LbrRecord_V2) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *LbrRecord_V2) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetThreadId sets the value of ThreadId for the instance
func (instance *LbrRecord_V2) SetPropertyThreadId(value uint32) (err error) {
	return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *LbrRecord_V2) GetPropertyThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadId")
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

// SetTimeStamp sets the value of TimeStamp for the instance
func (instance *LbrRecord_V2) SetPropertyTimeStamp(value uint64) (err error) {
	return instance.SetProperty("TimeStamp", (value))
}

// GetTimeStamp gets the value of TimeStamp for the instance
func (instance *LbrRecord_V2) GetPropertyTimeStamp() (value uint64, err error) {
	retValue, err := instance.GetProperty("TimeStamp")
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
