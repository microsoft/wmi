// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP struct
type Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP struct {
	*Win32_PerfFormattedData

	//
	ActiveFlushedEntries uint32

	//
	ActiveRequests uint32

	//
	ActiveThreadsCount uint32

	//
	CurrentFileCacheMemoryUsage uint64

	//
	CurrentFilesCached uint32

	//
	CurrentMetadataCached uint32

	//
	CurrentURIsCached uint32

	//
	FileCacheFlushes uint32

	//
	FileCacheHits uint32

	//
	FileCacheHitsPersec uint32

	//
	FileCacheMisses uint32

	//
	FileCacheMissesPersec uint32

	//
	MaximumFileCacheMemoryUsage uint64

	//
	MaximumThreadsCount uint32

	//
	MetadataCacheFlushes uint32

	//
	MetadataCacheHits uint32

	//
	MetadataCacheHitsPersec uint32

	//
	MetadataCacheMisses uint32

	//
	MetadataCacheMissesPersec uint32

	//
	OutputCacheCurrentFlushedItems uint32

	//
	OutputCacheCurrentItems uint32

	//
	OutputCacheCurrentMemoryUsage uint64

	//
	OutputCacheHitsPersec uint32

	//
	OutputCacheMissesPersec uint32

	//
	OutputCacheTotalFlushedItems uint32

	//
	OutputCacheTotalFlushes uint32

	//
	OutputCacheTotalHits uint32

	//
	OutputCacheTotalMisses uint32

	//
	Percent401HTTPResponseSent uint32

	//
	Percent403HTTPResponseSent uint32

	//
	Percent404HTTPResponseSent uint32

	//
	Percent500HTTPResponseSent uint32

	//
	RequestsPerSec uint32

	//
	TotalFilesCached uint32

	//
	TotalFlushedFiles uint32

	//
	TotalFlushedMetadata uint32

	//
	TotalFlushedURIs uint32

	//
	TotalHTTPRequestsServed uint32

	//
	TotalMetadataCached uint32

	//
	TotalThreads uint32

	//
	TotalURIsCached uint32

	//
	URICacheFlushes uint32

	//
	URICacheHits uint32

	//
	UriCacheHitsPersec uint32

	//
	URICacheMisses uint32

	//
	UriCacheMissesPersec uint32

	//
	WebSocketActiveRequests uint32

	//
	WebSocketConnectionAttemptsPerSec uint32

	//
	WebSocketConnectionsAcceptedPerSec uint32

	//
	WebSocketConnectionsRejectedPerSec uint32
}

func NewWin32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WPEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetActiveFlushedEntries sets the value of ActiveFlushedEntries for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyActiveFlushedEntries(value uint32) (err error) {
	return instance.SetProperty("ActiveFlushedEntries", (value))
}

// GetActiveFlushedEntries gets the value of ActiveFlushedEntries for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyActiveFlushedEntries() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveFlushedEntries")
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

// SetActiveRequests sets the value of ActiveRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyActiveRequests(value uint32) (err error) {
	return instance.SetProperty("ActiveRequests", (value))
}

// GetActiveRequests gets the value of ActiveRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyActiveRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveRequests")
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

// SetActiveThreadsCount sets the value of ActiveThreadsCount for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyActiveThreadsCount(value uint32) (err error) {
	return instance.SetProperty("ActiveThreadsCount", (value))
}

// GetActiveThreadsCount gets the value of ActiveThreadsCount for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyActiveThreadsCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveThreadsCount")
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

// SetCurrentFileCacheMemoryUsage sets the value of CurrentFileCacheMemoryUsage for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyCurrentFileCacheMemoryUsage(value uint64) (err error) {
	return instance.SetProperty("CurrentFileCacheMemoryUsage", (value))
}

// GetCurrentFileCacheMemoryUsage gets the value of CurrentFileCacheMemoryUsage for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyCurrentFileCacheMemoryUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentFileCacheMemoryUsage")
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

// SetCurrentFilesCached sets the value of CurrentFilesCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyCurrentFilesCached(value uint32) (err error) {
	return instance.SetProperty("CurrentFilesCached", (value))
}

// GetCurrentFilesCached gets the value of CurrentFilesCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyCurrentFilesCached() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentFilesCached")
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

// SetCurrentMetadataCached sets the value of CurrentMetadataCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyCurrentMetadataCached(value uint32) (err error) {
	return instance.SetProperty("CurrentMetadataCached", (value))
}

// GetCurrentMetadataCached gets the value of CurrentMetadataCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyCurrentMetadataCached() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentMetadataCached")
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

// SetCurrentURIsCached sets the value of CurrentURIsCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyCurrentURIsCached(value uint32) (err error) {
	return instance.SetProperty("CurrentURIsCached", (value))
}

// GetCurrentURIsCached gets the value of CurrentURIsCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyCurrentURIsCached() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentURIsCached")
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

// SetFileCacheFlushes sets the value of FileCacheFlushes for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyFileCacheFlushes(value uint32) (err error) {
	return instance.SetProperty("FileCacheFlushes", (value))
}

// GetFileCacheFlushes gets the value of FileCacheFlushes for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyFileCacheFlushes() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileCacheFlushes")
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

// SetFileCacheHits sets the value of FileCacheHits for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyFileCacheHits(value uint32) (err error) {
	return instance.SetProperty("FileCacheHits", (value))
}

// GetFileCacheHits gets the value of FileCacheHits for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyFileCacheHits() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileCacheHits")
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

// SetFileCacheHitsPersec sets the value of FileCacheHitsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyFileCacheHitsPersec(value uint32) (err error) {
	return instance.SetProperty("FileCacheHitsPersec", (value))
}

// GetFileCacheHitsPersec gets the value of FileCacheHitsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyFileCacheHitsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileCacheHitsPersec")
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

// SetFileCacheMisses sets the value of FileCacheMisses for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyFileCacheMisses(value uint32) (err error) {
	return instance.SetProperty("FileCacheMisses", (value))
}

// GetFileCacheMisses gets the value of FileCacheMisses for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyFileCacheMisses() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileCacheMisses")
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

// SetFileCacheMissesPersec sets the value of FileCacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyFileCacheMissesPersec(value uint32) (err error) {
	return instance.SetProperty("FileCacheMissesPersec", (value))
}

// GetFileCacheMissesPersec gets the value of FileCacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyFileCacheMissesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileCacheMissesPersec")
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

// SetMaximumFileCacheMemoryUsage sets the value of MaximumFileCacheMemoryUsage for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyMaximumFileCacheMemoryUsage(value uint64) (err error) {
	return instance.SetProperty("MaximumFileCacheMemoryUsage", (value))
}

// GetMaximumFileCacheMemoryUsage gets the value of MaximumFileCacheMemoryUsage for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyMaximumFileCacheMemoryUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaximumFileCacheMemoryUsage")
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

// SetMaximumThreadsCount sets the value of MaximumThreadsCount for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyMaximumThreadsCount(value uint32) (err error) {
	return instance.SetProperty("MaximumThreadsCount", (value))
}

// GetMaximumThreadsCount gets the value of MaximumThreadsCount for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyMaximumThreadsCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumThreadsCount")
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

// SetMetadataCacheFlushes sets the value of MetadataCacheFlushes for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyMetadataCacheFlushes(value uint32) (err error) {
	return instance.SetProperty("MetadataCacheFlushes", (value))
}

// GetMetadataCacheFlushes gets the value of MetadataCacheFlushes for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyMetadataCacheFlushes() (value uint32, err error) {
	retValue, err := instance.GetProperty("MetadataCacheFlushes")
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

// SetMetadataCacheHits sets the value of MetadataCacheHits for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyMetadataCacheHits(value uint32) (err error) {
	return instance.SetProperty("MetadataCacheHits", (value))
}

// GetMetadataCacheHits gets the value of MetadataCacheHits for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyMetadataCacheHits() (value uint32, err error) {
	retValue, err := instance.GetProperty("MetadataCacheHits")
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

// SetMetadataCacheHitsPersec sets the value of MetadataCacheHitsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyMetadataCacheHitsPersec(value uint32) (err error) {
	return instance.SetProperty("MetadataCacheHitsPersec", (value))
}

// GetMetadataCacheHitsPersec gets the value of MetadataCacheHitsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyMetadataCacheHitsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("MetadataCacheHitsPersec")
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

// SetMetadataCacheMisses sets the value of MetadataCacheMisses for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyMetadataCacheMisses(value uint32) (err error) {
	return instance.SetProperty("MetadataCacheMisses", (value))
}

// GetMetadataCacheMisses gets the value of MetadataCacheMisses for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyMetadataCacheMisses() (value uint32, err error) {
	retValue, err := instance.GetProperty("MetadataCacheMisses")
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

// SetMetadataCacheMissesPersec sets the value of MetadataCacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyMetadataCacheMissesPersec(value uint32) (err error) {
	return instance.SetProperty("MetadataCacheMissesPersec", (value))
}

// GetMetadataCacheMissesPersec gets the value of MetadataCacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyMetadataCacheMissesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("MetadataCacheMissesPersec")
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

// SetOutputCacheCurrentFlushedItems sets the value of OutputCacheCurrentFlushedItems for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyOutputCacheCurrentFlushedItems(value uint32) (err error) {
	return instance.SetProperty("OutputCacheCurrentFlushedItems", (value))
}

// GetOutputCacheCurrentFlushedItems gets the value of OutputCacheCurrentFlushedItems for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyOutputCacheCurrentFlushedItems() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutputCacheCurrentFlushedItems")
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

// SetOutputCacheCurrentItems sets the value of OutputCacheCurrentItems for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyOutputCacheCurrentItems(value uint32) (err error) {
	return instance.SetProperty("OutputCacheCurrentItems", (value))
}

// GetOutputCacheCurrentItems gets the value of OutputCacheCurrentItems for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyOutputCacheCurrentItems() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutputCacheCurrentItems")
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

// SetOutputCacheCurrentMemoryUsage sets the value of OutputCacheCurrentMemoryUsage for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyOutputCacheCurrentMemoryUsage(value uint64) (err error) {
	return instance.SetProperty("OutputCacheCurrentMemoryUsage", (value))
}

// GetOutputCacheCurrentMemoryUsage gets the value of OutputCacheCurrentMemoryUsage for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyOutputCacheCurrentMemoryUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutputCacheCurrentMemoryUsage")
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

// SetOutputCacheHitsPersec sets the value of OutputCacheHitsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyOutputCacheHitsPersec(value uint32) (err error) {
	return instance.SetProperty("OutputCacheHitsPersec", (value))
}

// GetOutputCacheHitsPersec gets the value of OutputCacheHitsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyOutputCacheHitsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutputCacheHitsPersec")
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

// SetOutputCacheMissesPersec sets the value of OutputCacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyOutputCacheMissesPersec(value uint32) (err error) {
	return instance.SetProperty("OutputCacheMissesPersec", (value))
}

// GetOutputCacheMissesPersec gets the value of OutputCacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyOutputCacheMissesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutputCacheMissesPersec")
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

// SetOutputCacheTotalFlushedItems sets the value of OutputCacheTotalFlushedItems for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyOutputCacheTotalFlushedItems(value uint32) (err error) {
	return instance.SetProperty("OutputCacheTotalFlushedItems", (value))
}

// GetOutputCacheTotalFlushedItems gets the value of OutputCacheTotalFlushedItems for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyOutputCacheTotalFlushedItems() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutputCacheTotalFlushedItems")
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

// SetOutputCacheTotalFlushes sets the value of OutputCacheTotalFlushes for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyOutputCacheTotalFlushes(value uint32) (err error) {
	return instance.SetProperty("OutputCacheTotalFlushes", (value))
}

// GetOutputCacheTotalFlushes gets the value of OutputCacheTotalFlushes for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyOutputCacheTotalFlushes() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutputCacheTotalFlushes")
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

// SetOutputCacheTotalHits sets the value of OutputCacheTotalHits for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyOutputCacheTotalHits(value uint32) (err error) {
	return instance.SetProperty("OutputCacheTotalHits", (value))
}

// GetOutputCacheTotalHits gets the value of OutputCacheTotalHits for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyOutputCacheTotalHits() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutputCacheTotalHits")
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

// SetOutputCacheTotalMisses sets the value of OutputCacheTotalMisses for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyOutputCacheTotalMisses(value uint32) (err error) {
	return instance.SetProperty("OutputCacheTotalMisses", (value))
}

// GetOutputCacheTotalMisses gets the value of OutputCacheTotalMisses for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyOutputCacheTotalMisses() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutputCacheTotalMisses")
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

// SetPercent401HTTPResponseSent sets the value of Percent401HTTPResponseSent for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyPercent401HTTPResponseSent(value uint32) (err error) {
	return instance.SetProperty("Percent401HTTPResponseSent", (value))
}

// GetPercent401HTTPResponseSent gets the value of Percent401HTTPResponseSent for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyPercent401HTTPResponseSent() (value uint32, err error) {
	retValue, err := instance.GetProperty("Percent401HTTPResponseSent")
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

// SetPercent403HTTPResponseSent sets the value of Percent403HTTPResponseSent for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyPercent403HTTPResponseSent(value uint32) (err error) {
	return instance.SetProperty("Percent403HTTPResponseSent", (value))
}

// GetPercent403HTTPResponseSent gets the value of Percent403HTTPResponseSent for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyPercent403HTTPResponseSent() (value uint32, err error) {
	retValue, err := instance.GetProperty("Percent403HTTPResponseSent")
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

// SetPercent404HTTPResponseSent sets the value of Percent404HTTPResponseSent for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyPercent404HTTPResponseSent(value uint32) (err error) {
	return instance.SetProperty("Percent404HTTPResponseSent", (value))
}

// GetPercent404HTTPResponseSent gets the value of Percent404HTTPResponseSent for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyPercent404HTTPResponseSent() (value uint32, err error) {
	retValue, err := instance.GetProperty("Percent404HTTPResponseSent")
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

// SetPercent500HTTPResponseSent sets the value of Percent500HTTPResponseSent for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyPercent500HTTPResponseSent(value uint32) (err error) {
	return instance.SetProperty("Percent500HTTPResponseSent", (value))
}

// GetPercent500HTTPResponseSent gets the value of Percent500HTTPResponseSent for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyPercent500HTTPResponseSent() (value uint32, err error) {
	retValue, err := instance.GetProperty("Percent500HTTPResponseSent")
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

// SetRequestsPerSec sets the value of RequestsPerSec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyRequestsPerSec(value uint32) (err error) {
	return instance.SetProperty("RequestsPerSec", (value))
}

// GetRequestsPerSec gets the value of RequestsPerSec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyRequestsPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RequestsPerSec")
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

// SetTotalFilesCached sets the value of TotalFilesCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyTotalFilesCached(value uint32) (err error) {
	return instance.SetProperty("TotalFilesCached", (value))
}

// GetTotalFilesCached gets the value of TotalFilesCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyTotalFilesCached() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalFilesCached")
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

// SetTotalFlushedFiles sets the value of TotalFlushedFiles for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyTotalFlushedFiles(value uint32) (err error) {
	return instance.SetProperty("TotalFlushedFiles", (value))
}

// GetTotalFlushedFiles gets the value of TotalFlushedFiles for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyTotalFlushedFiles() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalFlushedFiles")
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

// SetTotalFlushedMetadata sets the value of TotalFlushedMetadata for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyTotalFlushedMetadata(value uint32) (err error) {
	return instance.SetProperty("TotalFlushedMetadata", (value))
}

// GetTotalFlushedMetadata gets the value of TotalFlushedMetadata for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyTotalFlushedMetadata() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalFlushedMetadata")
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

// SetTotalFlushedURIs sets the value of TotalFlushedURIs for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyTotalFlushedURIs(value uint32) (err error) {
	return instance.SetProperty("TotalFlushedURIs", (value))
}

// GetTotalFlushedURIs gets the value of TotalFlushedURIs for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyTotalFlushedURIs() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalFlushedURIs")
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

// SetTotalHTTPRequestsServed sets the value of TotalHTTPRequestsServed for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyTotalHTTPRequestsServed(value uint32) (err error) {
	return instance.SetProperty("TotalHTTPRequestsServed", (value))
}

// GetTotalHTTPRequestsServed gets the value of TotalHTTPRequestsServed for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyTotalHTTPRequestsServed() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalHTTPRequestsServed")
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

// SetTotalMetadataCached sets the value of TotalMetadataCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyTotalMetadataCached(value uint32) (err error) {
	return instance.SetProperty("TotalMetadataCached", (value))
}

// GetTotalMetadataCached gets the value of TotalMetadataCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyTotalMetadataCached() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalMetadataCached")
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

// SetTotalThreads sets the value of TotalThreads for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyTotalThreads(value uint32) (err error) {
	return instance.SetProperty("TotalThreads", (value))
}

// GetTotalThreads gets the value of TotalThreads for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyTotalThreads() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalThreads")
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

// SetTotalURIsCached sets the value of TotalURIsCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyTotalURIsCached(value uint32) (err error) {
	return instance.SetProperty("TotalURIsCached", (value))
}

// GetTotalURIsCached gets the value of TotalURIsCached for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyTotalURIsCached() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalURIsCached")
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

// SetURICacheFlushes sets the value of URICacheFlushes for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyURICacheFlushes(value uint32) (err error) {
	return instance.SetProperty("URICacheFlushes", (value))
}

// GetURICacheFlushes gets the value of URICacheFlushes for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyURICacheFlushes() (value uint32, err error) {
	retValue, err := instance.GetProperty("URICacheFlushes")
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

// SetURICacheHits sets the value of URICacheHits for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyURICacheHits(value uint32) (err error) {
	return instance.SetProperty("URICacheHits", (value))
}

// GetURICacheHits gets the value of URICacheHits for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyURICacheHits() (value uint32, err error) {
	retValue, err := instance.GetProperty("URICacheHits")
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

// SetUriCacheHitsPersec sets the value of UriCacheHitsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyUriCacheHitsPersec(value uint32) (err error) {
	return instance.SetProperty("UriCacheHitsPersec", (value))
}

// GetUriCacheHitsPersec gets the value of UriCacheHitsPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyUriCacheHitsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("UriCacheHitsPersec")
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

// SetURICacheMisses sets the value of URICacheMisses for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyURICacheMisses(value uint32) (err error) {
	return instance.SetProperty("URICacheMisses", (value))
}

// GetURICacheMisses gets the value of URICacheMisses for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyURICacheMisses() (value uint32, err error) {
	retValue, err := instance.GetProperty("URICacheMisses")
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

// SetUriCacheMissesPersec sets the value of UriCacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyUriCacheMissesPersec(value uint32) (err error) {
	return instance.SetProperty("UriCacheMissesPersec", (value))
}

// GetUriCacheMissesPersec gets the value of UriCacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyUriCacheMissesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("UriCacheMissesPersec")
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

// SetWebSocketActiveRequests sets the value of WebSocketActiveRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyWebSocketActiveRequests(value uint32) (err error) {
	return instance.SetProperty("WebSocketActiveRequests", (value))
}

// GetWebSocketActiveRequests gets the value of WebSocketActiveRequests for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyWebSocketActiveRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("WebSocketActiveRequests")
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

// SetWebSocketConnectionAttemptsPerSec sets the value of WebSocketConnectionAttemptsPerSec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyWebSocketConnectionAttemptsPerSec(value uint32) (err error) {
	return instance.SetProperty("WebSocketConnectionAttemptsPerSec", (value))
}

// GetWebSocketConnectionAttemptsPerSec gets the value of WebSocketConnectionAttemptsPerSec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyWebSocketConnectionAttemptsPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WebSocketConnectionAttemptsPerSec")
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

// SetWebSocketConnectionsAcceptedPerSec sets the value of WebSocketConnectionsAcceptedPerSec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyWebSocketConnectionsAcceptedPerSec(value uint32) (err error) {
	return instance.SetProperty("WebSocketConnectionsAcceptedPerSec", (value))
}

// GetWebSocketConnectionsAcceptedPerSec gets the value of WebSocketConnectionsAcceptedPerSec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyWebSocketConnectionsAcceptedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WebSocketConnectionsAcceptedPerSec")
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

// SetWebSocketConnectionsRejectedPerSec sets the value of WebSocketConnectionsRejectedPerSec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) SetPropertyWebSocketConnectionsRejectedPerSec(value uint32) (err error) {
	return instance.SetProperty("WebSocketConnectionsRejectedPerSec", (value))
}

// GetWebSocketConnectionsRejectedPerSec gets the value of WebSocketConnectionsRejectedPerSec for the instance
func (instance *Win32_PerfFormattedData_W3SVCW3WPCounterProvider_W3SVCW3WP) GetPropertyWebSocketConnectionsRejectedPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WebSocketConnectionsRejectedPerSec")
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
