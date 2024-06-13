// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2 struct
type Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2 struct {
	*Win32_PerfFormattedData

	//
	CacheReadPopulateL1Bytes uint64

	//
	CacheReadPopulateL1BytesPersec uint64

	//
	CacheReadPopulateL2Bytes uint64

	//
	CacheReadPopulateL2BytesPersec uint64

	//
	CacheWritePopulateL0Bytes uint64

	//
	CacheWritePopulateL0BytesPersec uint64

	//
	CacheWritePopulateL1Bytes uint64

	//
	CacheWritePopulateL1BytesPersec uint64

	//
	CacheWritePopulateL2Bytes uint64

	//
	CacheWritePopulateL2BytesPersec uint64

	//
	HeatMapFreeMemory uint64

	//
	HeatMapWindow uint64

	//
	RateDiskVRCReads uint64

	//
	VRCHitReadBytes uint64

	//
	VRCHitReadBytesPersec uint64

	//
	VRCHitReads uint64

	//
	VRCHitReadsPersec uint64

	//
	VRCPopulateBytes uint64

	//
	VRCPopulateBytesPersec uint64

	//
	VRCPopulates uint64

	//
	VRCPopulatesPersec uint64
}

func NewWin32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2Ex1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetCacheReadPopulateL1Bytes sets the value of CacheReadPopulateL1Bytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyCacheReadPopulateL1Bytes(value uint64) (err error) {
	return instance.SetProperty("CacheReadPopulateL1Bytes", (value))
}

// GetCacheReadPopulateL1Bytes gets the value of CacheReadPopulateL1Bytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyCacheReadPopulateL1Bytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheReadPopulateL1Bytes")
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

// SetCacheReadPopulateL1BytesPersec sets the value of CacheReadPopulateL1BytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyCacheReadPopulateL1BytesPersec(value uint64) (err error) {
	return instance.SetProperty("CacheReadPopulateL1BytesPersec", (value))
}

// GetCacheReadPopulateL1BytesPersec gets the value of CacheReadPopulateL1BytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyCacheReadPopulateL1BytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheReadPopulateL1BytesPersec")
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

// SetCacheReadPopulateL2Bytes sets the value of CacheReadPopulateL2Bytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyCacheReadPopulateL2Bytes(value uint64) (err error) {
	return instance.SetProperty("CacheReadPopulateL2Bytes", (value))
}

// GetCacheReadPopulateL2Bytes gets the value of CacheReadPopulateL2Bytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyCacheReadPopulateL2Bytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheReadPopulateL2Bytes")
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

// SetCacheReadPopulateL2BytesPersec sets the value of CacheReadPopulateL2BytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyCacheReadPopulateL2BytesPersec(value uint64) (err error) {
	return instance.SetProperty("CacheReadPopulateL2BytesPersec", (value))
}

// GetCacheReadPopulateL2BytesPersec gets the value of CacheReadPopulateL2BytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyCacheReadPopulateL2BytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheReadPopulateL2BytesPersec")
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

// SetCacheWritePopulateL0Bytes sets the value of CacheWritePopulateL0Bytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyCacheWritePopulateL0Bytes(value uint64) (err error) {
	return instance.SetProperty("CacheWritePopulateL0Bytes", (value))
}

// GetCacheWritePopulateL0Bytes gets the value of CacheWritePopulateL0Bytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyCacheWritePopulateL0Bytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheWritePopulateL0Bytes")
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

// SetCacheWritePopulateL0BytesPersec sets the value of CacheWritePopulateL0BytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyCacheWritePopulateL0BytesPersec(value uint64) (err error) {
	return instance.SetProperty("CacheWritePopulateL0BytesPersec", (value))
}

// GetCacheWritePopulateL0BytesPersec gets the value of CacheWritePopulateL0BytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyCacheWritePopulateL0BytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheWritePopulateL0BytesPersec")
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

// SetCacheWritePopulateL1Bytes sets the value of CacheWritePopulateL1Bytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyCacheWritePopulateL1Bytes(value uint64) (err error) {
	return instance.SetProperty("CacheWritePopulateL1Bytes", (value))
}

// GetCacheWritePopulateL1Bytes gets the value of CacheWritePopulateL1Bytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyCacheWritePopulateL1Bytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheWritePopulateL1Bytes")
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

// SetCacheWritePopulateL1BytesPersec sets the value of CacheWritePopulateL1BytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyCacheWritePopulateL1BytesPersec(value uint64) (err error) {
	return instance.SetProperty("CacheWritePopulateL1BytesPersec", (value))
}

// GetCacheWritePopulateL1BytesPersec gets the value of CacheWritePopulateL1BytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyCacheWritePopulateL1BytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheWritePopulateL1BytesPersec")
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

// SetCacheWritePopulateL2Bytes sets the value of CacheWritePopulateL2Bytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyCacheWritePopulateL2Bytes(value uint64) (err error) {
	return instance.SetProperty("CacheWritePopulateL2Bytes", (value))
}

// GetCacheWritePopulateL2Bytes gets the value of CacheWritePopulateL2Bytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyCacheWritePopulateL2Bytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheWritePopulateL2Bytes")
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

// SetCacheWritePopulateL2BytesPersec sets the value of CacheWritePopulateL2BytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyCacheWritePopulateL2BytesPersec(value uint64) (err error) {
	return instance.SetProperty("CacheWritePopulateL2BytesPersec", (value))
}

// GetCacheWritePopulateL2BytesPersec gets the value of CacheWritePopulateL2BytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyCacheWritePopulateL2BytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheWritePopulateL2BytesPersec")
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

// SetHeatMapFreeMemory sets the value of HeatMapFreeMemory for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyHeatMapFreeMemory(value uint64) (err error) {
	return instance.SetProperty("HeatMapFreeMemory", (value))
}

// GetHeatMapFreeMemory gets the value of HeatMapFreeMemory for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyHeatMapFreeMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("HeatMapFreeMemory")
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

// SetHeatMapWindow sets the value of HeatMapWindow for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyHeatMapWindow(value uint64) (err error) {
	return instance.SetProperty("HeatMapWindow", (value))
}

// GetHeatMapWindow gets the value of HeatMapWindow for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyHeatMapWindow() (value uint64, err error) {
	retValue, err := instance.GetProperty("HeatMapWindow")
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

// SetRateDiskVRCReads sets the value of RateDiskVRCReads for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyRateDiskVRCReads(value uint64) (err error) {
	return instance.SetProperty("RateDiskVRCReads", (value))
}

// GetRateDiskVRCReads gets the value of RateDiskVRCReads for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyRateDiskVRCReads() (value uint64, err error) {
	retValue, err := instance.GetProperty("RateDiskVRCReads")
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

// SetVRCHitReadBytes sets the value of VRCHitReadBytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyVRCHitReadBytes(value uint64) (err error) {
	return instance.SetProperty("VRCHitReadBytes", (value))
}

// GetVRCHitReadBytes gets the value of VRCHitReadBytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyVRCHitReadBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRCHitReadBytes")
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

// SetVRCHitReadBytesPersec sets the value of VRCHitReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyVRCHitReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("VRCHitReadBytesPersec", (value))
}

// GetVRCHitReadBytesPersec gets the value of VRCHitReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyVRCHitReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRCHitReadBytesPersec")
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

// SetVRCHitReads sets the value of VRCHitReads for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyVRCHitReads(value uint64) (err error) {
	return instance.SetProperty("VRCHitReads", (value))
}

// GetVRCHitReads gets the value of VRCHitReads for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyVRCHitReads() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRCHitReads")
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

// SetVRCHitReadsPersec sets the value of VRCHitReadsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyVRCHitReadsPersec(value uint64) (err error) {
	return instance.SetProperty("VRCHitReadsPersec", (value))
}

// GetVRCHitReadsPersec gets the value of VRCHitReadsPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyVRCHitReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRCHitReadsPersec")
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

// SetVRCPopulateBytes sets the value of VRCPopulateBytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyVRCPopulateBytes(value uint64) (err error) {
	return instance.SetProperty("VRCPopulateBytes", (value))
}

// GetVRCPopulateBytes gets the value of VRCPopulateBytes for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyVRCPopulateBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRCPopulateBytes")
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

// SetVRCPopulateBytesPersec sets the value of VRCPopulateBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyVRCPopulateBytesPersec(value uint64) (err error) {
	return instance.SetProperty("VRCPopulateBytesPersec", (value))
}

// GetVRCPopulateBytesPersec gets the value of VRCPopulateBytesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyVRCPopulateBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRCPopulateBytesPersec")
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

// SetVRCPopulates sets the value of VRCPopulates for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyVRCPopulates(value uint64) (err error) {
	return instance.SetProperty("VRCPopulates", (value))
}

// GetVRCPopulates gets the value of VRCPopulates for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyVRCPopulates() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRCPopulates")
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

// SetVRCPopulatesPersec sets the value of VRCPopulatesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) SetPropertyVRCPopulatesPersec(value uint64) (err error) {
	return instance.SetProperty("VRCPopulatesPersec", (value))
}

// GetVRCPopulatesPersec gets the value of VRCPopulatesPersec for the instance
func (instance *Win32_PerfFormattedData_ClusBfltPerfProvider_ClusterStorageHybridDisks2) GetPropertyVRCPopulatesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRCPopulatesPersec")
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
