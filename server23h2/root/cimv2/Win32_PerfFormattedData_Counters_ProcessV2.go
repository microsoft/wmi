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

// Win32_PerfFormattedData_Counters_ProcessV2 struct
type Win32_PerfFormattedData_Counters_ProcessV2 struct {
	*Win32_PerfFormattedData

	//
	CreatingProcessID uint32

	//
	ElapsedTime uint64

	//
	HandleCount uint32

	//
	IODataBytesPersec uint64

	//
	IODataOperationsPersec uint64

	//
	IOOtherBytesPersec uint64

	//
	IOOtherOperationsPersec uint64

	//
	IOReadBytesPersec uint64

	//
	IOReadOperationsPersec uint64

	//
	IOWriteBytesPersec uint64

	//
	IOWriteOperationsPersec uint64

	//
	PageFaultsPersec uint32

	//
	PageFileBytes uint64

	//
	PageFileBytesPeak uint64

	//
	PercentPrivilegedTime uint64

	//
	PercentProcessorTime uint64

	//
	PercentUserTime uint64

	//
	PoolNonpagedBytes uint64

	//
	PoolPagedBytes uint64

	//
	PriorityBase uint32

	//
	PrivateBytes uint64

	//
	ProcessID uint32

	//
	ThreadCount uint32

	//
	VirtualBytes uint64

	//
	VirtualBytesPeak uint64

	//
	WorkingSet uint64

	//
	WorkingSetPeak uint64

	//
	WorkingSetPrivate uint64
}

func NewWin32_PerfFormattedData_Counters_ProcessV2Ex1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_ProcessV2, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_ProcessV2{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_ProcessV2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_ProcessV2, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_ProcessV2{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetCreatingProcessID sets the value of CreatingProcessID for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyCreatingProcessID(value uint32) (err error) {
	return instance.SetProperty("CreatingProcessID", (value))
}

// GetCreatingProcessID gets the value of CreatingProcessID for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyCreatingProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("CreatingProcessID")
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

// SetElapsedTime sets the value of ElapsedTime for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyElapsedTime(value uint64) (err error) {
	return instance.SetProperty("ElapsedTime", (value))
}

// GetElapsedTime gets the value of ElapsedTime for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyElapsedTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("ElapsedTime")
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

// SetHandleCount sets the value of HandleCount for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyHandleCount(value uint32) (err error) {
	return instance.SetProperty("HandleCount", (value))
}

// GetHandleCount gets the value of HandleCount for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyHandleCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("HandleCount")
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

// SetIODataBytesPersec sets the value of IODataBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyIODataBytesPersec(value uint64) (err error) {
	return instance.SetProperty("IODataBytesPersec", (value))
}

// GetIODataBytesPersec gets the value of IODataBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyIODataBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IODataBytesPersec")
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

// SetIODataOperationsPersec sets the value of IODataOperationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyIODataOperationsPersec(value uint64) (err error) {
	return instance.SetProperty("IODataOperationsPersec", (value))
}

// GetIODataOperationsPersec gets the value of IODataOperationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyIODataOperationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IODataOperationsPersec")
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

// SetIOOtherBytesPersec sets the value of IOOtherBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyIOOtherBytesPersec(value uint64) (err error) {
	return instance.SetProperty("IOOtherBytesPersec", (value))
}

// GetIOOtherBytesPersec gets the value of IOOtherBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyIOOtherBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOOtherBytesPersec")
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

// SetIOOtherOperationsPersec sets the value of IOOtherOperationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyIOOtherOperationsPersec(value uint64) (err error) {
	return instance.SetProperty("IOOtherOperationsPersec", (value))
}

// GetIOOtherOperationsPersec gets the value of IOOtherOperationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyIOOtherOperationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOOtherOperationsPersec")
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

// SetIOReadBytesPersec sets the value of IOReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyIOReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("IOReadBytesPersec", (value))
}

// GetIOReadBytesPersec gets the value of IOReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyIOReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadBytesPersec")
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

// SetIOReadOperationsPersec sets the value of IOReadOperationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyIOReadOperationsPersec(value uint64) (err error) {
	return instance.SetProperty("IOReadOperationsPersec", (value))
}

// GetIOReadOperationsPersec gets the value of IOReadOperationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyIOReadOperationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadOperationsPersec")
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

// SetIOWriteBytesPersec sets the value of IOWriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyIOWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("IOWriteBytesPersec", (value))
}

// GetIOWriteBytesPersec gets the value of IOWriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyIOWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWriteBytesPersec")
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

// SetIOWriteOperationsPersec sets the value of IOWriteOperationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyIOWriteOperationsPersec(value uint64) (err error) {
	return instance.SetProperty("IOWriteOperationsPersec", (value))
}

// GetIOWriteOperationsPersec gets the value of IOWriteOperationsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyIOWriteOperationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWriteOperationsPersec")
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

// SetPageFaultsPersec sets the value of PageFaultsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyPageFaultsPersec(value uint32) (err error) {
	return instance.SetProperty("PageFaultsPersec", (value))
}

// GetPageFaultsPersec gets the value of PageFaultsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyPageFaultsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PageFaultsPersec")
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

// SetPageFileBytes sets the value of PageFileBytes for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyPageFileBytes(value uint64) (err error) {
	return instance.SetProperty("PageFileBytes", (value))
}

// GetPageFileBytes gets the value of PageFileBytes for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyPageFileBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageFileBytes")
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

// SetPageFileBytesPeak sets the value of PageFileBytesPeak for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyPageFileBytesPeak(value uint64) (err error) {
	return instance.SetProperty("PageFileBytesPeak", (value))
}

// GetPageFileBytesPeak gets the value of PageFileBytesPeak for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyPageFileBytesPeak() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageFileBytesPeak")
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

// SetPercentPrivilegedTime sets the value of PercentPrivilegedTime for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyPercentPrivilegedTime(value uint64) (err error) {
	return instance.SetProperty("PercentPrivilegedTime", (value))
}

// GetPercentPrivilegedTime gets the value of PercentPrivilegedTime for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyPercentPrivilegedTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentPrivilegedTime")
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

// SetPercentProcessorTime sets the value of PercentProcessorTime for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyPercentProcessorTime(value uint64) (err error) {
	return instance.SetProperty("PercentProcessorTime", (value))
}

// GetPercentProcessorTime gets the value of PercentProcessorTime for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyPercentProcessorTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentProcessorTime")
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

// SetPercentUserTime sets the value of PercentUserTime for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyPercentUserTime(value uint64) (err error) {
	return instance.SetProperty("PercentUserTime", (value))
}

// GetPercentUserTime gets the value of PercentUserTime for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyPercentUserTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentUserTime")
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

// SetPoolNonpagedBytes sets the value of PoolNonpagedBytes for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyPoolNonpagedBytes(value uint64) (err error) {
	return instance.SetProperty("PoolNonpagedBytes", (value))
}

// GetPoolNonpagedBytes gets the value of PoolNonpagedBytes for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyPoolNonpagedBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("PoolNonpagedBytes")
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

// SetPoolPagedBytes sets the value of PoolPagedBytes for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyPoolPagedBytes(value uint64) (err error) {
	return instance.SetProperty("PoolPagedBytes", (value))
}

// GetPoolPagedBytes gets the value of PoolPagedBytes for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyPoolPagedBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("PoolPagedBytes")
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

// SetPriorityBase sets the value of PriorityBase for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyPriorityBase(value uint32) (err error) {
	return instance.SetProperty("PriorityBase", (value))
}

// GetPriorityBase gets the value of PriorityBase for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyPriorityBase() (value uint32, err error) {
	retValue, err := instance.GetProperty("PriorityBase")
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

// SetPrivateBytes sets the value of PrivateBytes for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyPrivateBytes(value uint64) (err error) {
	return instance.SetProperty("PrivateBytes", (value))
}

// GetPrivateBytes gets the value of PrivateBytes for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyPrivateBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("PrivateBytes")
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

// SetProcessID sets the value of ProcessID for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyProcessID(value uint32) (err error) {
	return instance.SetProperty("ProcessID", (value))
}

// GetProcessID gets the value of ProcessID for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessID")
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

// SetThreadCount sets the value of ThreadCount for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyThreadCount(value uint32) (err error) {
	return instance.SetProperty("ThreadCount", (value))
}

// GetThreadCount gets the value of ThreadCount for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyThreadCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadCount")
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

// SetVirtualBytes sets the value of VirtualBytes for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyVirtualBytes(value uint64) (err error) {
	return instance.SetProperty("VirtualBytes", (value))
}

// GetVirtualBytes gets the value of VirtualBytes for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyVirtualBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualBytes")
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

// SetVirtualBytesPeak sets the value of VirtualBytesPeak for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyVirtualBytesPeak(value uint64) (err error) {
	return instance.SetProperty("VirtualBytesPeak", (value))
}

// GetVirtualBytesPeak gets the value of VirtualBytesPeak for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyVirtualBytesPeak() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualBytesPeak")
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

// SetWorkingSet sets the value of WorkingSet for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyWorkingSet(value uint64) (err error) {
	return instance.SetProperty("WorkingSet", (value))
}

// GetWorkingSet gets the value of WorkingSet for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyWorkingSet() (value uint64, err error) {
	retValue, err := instance.GetProperty("WorkingSet")
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

// SetWorkingSetPeak sets the value of WorkingSetPeak for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyWorkingSetPeak(value uint64) (err error) {
	return instance.SetProperty("WorkingSetPeak", (value))
}

// GetWorkingSetPeak gets the value of WorkingSetPeak for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyWorkingSetPeak() (value uint64, err error) {
	retValue, err := instance.GetProperty("WorkingSetPeak")
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

// SetWorkingSetPrivate sets the value of WorkingSetPrivate for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) SetPropertyWorkingSetPrivate(value uint64) (err error) {
	return instance.SetProperty("WorkingSetPrivate", (value))
}

// GetWorkingSetPrivate gets the value of WorkingSetPrivate for the instance
func (instance *Win32_PerfFormattedData_Counters_ProcessV2) GetPropertyWorkingSetPrivate() (value uint64, err error) {
	retValue, err := instance.GetProperty("WorkingSetPrivate")
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
