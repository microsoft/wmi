// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap struct
type Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap struct {
	*Win32_PerfRawData

	//
	VirtualDiskMapBytesAverage uint64

	//
	VirtualDiskMapBytesAverage_Base uint32

	//
	VirtualDiskMapBytesPersec uint64

	//
	VirtualDiskMapExtentLatencyms uint32

	//
	VirtualDiskMapExtentLatencyms_Base uint32

	//
	VirtualDiskMapExtentsAverage uint64

	//
	VirtualDiskMapExtentsAverage_Base uint32

	//
	VirtualDiskMapExtentsPersec uint64

	//
	VirtualDiskMapLatencyms uint32

	//
	VirtualDiskMapLatencyms_Base uint32

	//
	VirtualDiskMapsPersec uint64
}

func NewWin32_PerfRawData_Counters_StorageSpacesVirtualDiskMapEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_StorageSpacesVirtualDiskMapEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetVirtualDiskMapBytesAverage sets the value of VirtualDiskMapBytesAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapBytesAverage(value uint64) (err error) {
	return instance.SetProperty("VirtualDiskMapBytesAverage", (value))
}

// GetVirtualDiskMapBytesAverage gets the value of VirtualDiskMapBytesAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapBytesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapBytesAverage")
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

// SetVirtualDiskMapBytesAverage_Base sets the value of VirtualDiskMapBytesAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapBytesAverage_Base(value uint32) (err error) {
	return instance.SetProperty("VirtualDiskMapBytesAverage_Base", (value))
}

// GetVirtualDiskMapBytesAverage_Base gets the value of VirtualDiskMapBytesAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapBytesAverage_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapBytesAverage_Base")
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

// SetVirtualDiskMapBytesPersec sets the value of VirtualDiskMapBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapBytesPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualDiskMapBytesPersec", (value))
}

// GetVirtualDiskMapBytesPersec gets the value of VirtualDiskMapBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapBytesPersec")
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

// SetVirtualDiskMapExtentLatencyms sets the value of VirtualDiskMapExtentLatencyms for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapExtentLatencyms(value uint32) (err error) {
	return instance.SetProperty("VirtualDiskMapExtentLatencyms", (value))
}

// GetVirtualDiskMapExtentLatencyms gets the value of VirtualDiskMapExtentLatencyms for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapExtentLatencyms() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapExtentLatencyms")
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

// SetVirtualDiskMapExtentLatencyms_Base sets the value of VirtualDiskMapExtentLatencyms_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapExtentLatencyms_Base(value uint32) (err error) {
	return instance.SetProperty("VirtualDiskMapExtentLatencyms_Base", (value))
}

// GetVirtualDiskMapExtentLatencyms_Base gets the value of VirtualDiskMapExtentLatencyms_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapExtentLatencyms_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapExtentLatencyms_Base")
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

// SetVirtualDiskMapExtentsAverage sets the value of VirtualDiskMapExtentsAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapExtentsAverage(value uint64) (err error) {
	return instance.SetProperty("VirtualDiskMapExtentsAverage", (value))
}

// GetVirtualDiskMapExtentsAverage gets the value of VirtualDiskMapExtentsAverage for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapExtentsAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapExtentsAverage")
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

// SetVirtualDiskMapExtentsAverage_Base sets the value of VirtualDiskMapExtentsAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapExtentsAverage_Base(value uint32) (err error) {
	return instance.SetProperty("VirtualDiskMapExtentsAverage_Base", (value))
}

// GetVirtualDiskMapExtentsAverage_Base gets the value of VirtualDiskMapExtentsAverage_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapExtentsAverage_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapExtentsAverage_Base")
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

// SetVirtualDiskMapExtentsPersec sets the value of VirtualDiskMapExtentsPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapExtentsPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualDiskMapExtentsPersec", (value))
}

// GetVirtualDiskMapExtentsPersec gets the value of VirtualDiskMapExtentsPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapExtentsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapExtentsPersec")
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

// SetVirtualDiskMapLatencyms sets the value of VirtualDiskMapLatencyms for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapLatencyms(value uint32) (err error) {
	return instance.SetProperty("VirtualDiskMapLatencyms", (value))
}

// GetVirtualDiskMapLatencyms gets the value of VirtualDiskMapLatencyms for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapLatencyms() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapLatencyms")
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

// SetVirtualDiskMapLatencyms_Base sets the value of VirtualDiskMapLatencyms_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapLatencyms_Base(value uint32) (err error) {
	return instance.SetProperty("VirtualDiskMapLatencyms_Base", (value))
}

// GetVirtualDiskMapLatencyms_Base gets the value of VirtualDiskMapLatencyms_Base for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapLatencyms_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapLatencyms_Base")
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

// SetVirtualDiskMapsPersec sets the value of VirtualDiskMapsPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) SetPropertyVirtualDiskMapsPersec(value uint64) (err error) {
	return instance.SetProperty("VirtualDiskMapsPersec", (value))
}

// GetVirtualDiskMapsPersec gets the value of VirtualDiskMapsPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageSpacesVirtualDiskMap) GetPropertyVirtualDiskMapsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualDiskMapsPersec")
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
