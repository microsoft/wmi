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

// Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo struct
type Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo struct {
	*Win32_PerfFormattedData

	//
	VirtualDiskFlushLatencyms uint32

	//
	VirtualDiskReadBytesPersec uint64

	//
	VirtualDiskReadLatencyms uint32

	//
	VirtualDiskWriteBytesPersec uint64

	//
	VirtualDiskWriteLatencyms uint32
}

func NewWin32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIoEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetVirtualDiskFlushLatencyms sets the value of VirtualDiskFlushLatencyms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskFlushLatencyms(value uint32) (err error) {
	return instance.SetProperty("VirtualDiskFlushLatencyms", (value))
}

// GetVirtualDiskFlushLatencyms gets the value of VirtualDiskFlushLatencyms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskFlushLatencyms() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualDiskFlushLatencyms")
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

// SetVirtualDiskReadBytesPersec sets the value of VirtualDiskReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualDiskReadBytesPersec", (value))
}

// GetVirtualDiskReadBytesPersec gets the value of VirtualDiskReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualDiskReadBytesPersec")
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

// SetVirtualDiskReadLatencyms sets the value of VirtualDiskReadLatencyms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskReadLatencyms(value uint32) (err error) {
	return instance.SetProperty("VirtualDiskReadLatencyms", (value))
}

// GetVirtualDiskReadLatencyms gets the value of VirtualDiskReadLatencyms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskReadLatencyms() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualDiskReadLatencyms")
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

// SetVirtualDiskWriteBytesPersec sets the value of VirtualDiskWriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualDiskWriteBytesPersec", (value))
}

// GetVirtualDiskWriteBytesPersec gets the value of VirtualDiskWriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualDiskWriteBytesPersec")
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

// SetVirtualDiskWriteLatencyms sets the value of VirtualDiskWriteLatencyms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo) SetPropertyVirtualDiskWriteLatencyms(value uint32) (err error) {
	return instance.SetProperty("VirtualDiskWriteLatencyms", (value))
}

// GetVirtualDiskWriteLatencyms gets the value of VirtualDiskWriteLatencyms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesVirtualDiskIo) GetPropertyVirtualDiskWriteLatencyms() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualDiskWriteLatencyms")
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
