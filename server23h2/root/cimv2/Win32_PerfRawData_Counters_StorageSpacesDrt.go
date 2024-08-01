// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_Counters_StorageSpacesDrt struct
type Win32_PerfRawData_Counters_StorageSpacesDrt struct {
	*Win32_PerfRawData

	//
	CleanBytes uint64

	//
	CleanCount uint64

	//
	DirtyBytes uint64

	//
	DirtyCount uint64

	//
	FlushedBytes uint64

	//
	FlushedCount uint64

	//
	FlushingBytes uint64

	//
	FlushingCount uint64

	//
	Limit uint32

	//
	SynchronizingBytes uint64

	//
	SynchronizingCount uint64
}

func NewWin32_PerfRawData_Counters_StorageSpacesDrtEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_StorageSpacesDrt, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_StorageSpacesDrt{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_StorageSpacesDrtEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_StorageSpacesDrt, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_StorageSpacesDrt{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetCleanBytes sets the value of CleanBytes for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertyCleanBytes(value uint64) (err error) {
	return instance.SetProperty("CleanBytes", (value))
}

// GetCleanBytes gets the value of CleanBytes for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertyCleanBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("CleanBytes")
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

// SetCleanCount sets the value of CleanCount for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertyCleanCount(value uint64) (err error) {
	return instance.SetProperty("CleanCount", (value))
}

// GetCleanCount gets the value of CleanCount for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertyCleanCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("CleanCount")
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

// SetDirtyBytes sets the value of DirtyBytes for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertyDirtyBytes(value uint64) (err error) {
	return instance.SetProperty("DirtyBytes", (value))
}

// GetDirtyBytes gets the value of DirtyBytes for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertyDirtyBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("DirtyBytes")
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

// SetDirtyCount sets the value of DirtyCount for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertyDirtyCount(value uint64) (err error) {
	return instance.SetProperty("DirtyCount", (value))
}

// GetDirtyCount gets the value of DirtyCount for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertyDirtyCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("DirtyCount")
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

// SetFlushedBytes sets the value of FlushedBytes for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertyFlushedBytes(value uint64) (err error) {
	return instance.SetProperty("FlushedBytes", (value))
}

// GetFlushedBytes gets the value of FlushedBytes for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertyFlushedBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("FlushedBytes")
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

// SetFlushedCount sets the value of FlushedCount for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertyFlushedCount(value uint64) (err error) {
	return instance.SetProperty("FlushedCount", (value))
}

// GetFlushedCount gets the value of FlushedCount for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertyFlushedCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("FlushedCount")
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

// SetFlushingBytes sets the value of FlushingBytes for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertyFlushingBytes(value uint64) (err error) {
	return instance.SetProperty("FlushingBytes", (value))
}

// GetFlushingBytes gets the value of FlushingBytes for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertyFlushingBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("FlushingBytes")
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

// SetFlushingCount sets the value of FlushingCount for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertyFlushingCount(value uint64) (err error) {
	return instance.SetProperty("FlushingCount", (value))
}

// GetFlushingCount gets the value of FlushingCount for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertyFlushingCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("FlushingCount")
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

// SetLimit sets the value of Limit for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertyLimit(value uint32) (err error) {
	return instance.SetProperty("Limit", (value))
}

// GetLimit gets the value of Limit for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertyLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("Limit")
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

// SetSynchronizingBytes sets the value of SynchronizingBytes for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertySynchronizingBytes(value uint64) (err error) {
	return instance.SetProperty("SynchronizingBytes", (value))
}

// GetSynchronizingBytes gets the value of SynchronizingBytes for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertySynchronizingBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("SynchronizingBytes")
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

// SetSynchronizingCount sets the value of SynchronizingCount for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) SetPropertySynchronizingCount(value uint64) (err error) {
	return instance.SetProperty("SynchronizingCount", (value))
}

// GetSynchronizingCount gets the value of SynchronizingCount for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesDrt) GetPropertySynchronizingCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("SynchronizingCount")
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
