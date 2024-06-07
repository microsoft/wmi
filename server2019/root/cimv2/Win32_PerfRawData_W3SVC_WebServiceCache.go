// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_PerfRawData_W3SVC_WebServiceCache struct
type Win32_PerfRawData_W3SVC_WebServiceCache struct { 
	*Win32_PerfRawData

	// 
	ActiveFlushedEntries uint32

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
	FileCacheHitsPercent uint32

	// 
	FileCacheHitsPercent_Base uint32

	// 
	FileCacheMisses uint32

	// 
	KernelCurrentURIsCached uint32

	// 
	KernelTotalFlushedURIs uint32

	// 
	KernelTotalURIsCached uint32

	// 
	KernelURICacheFlushes uint32

	// 
	KernelURICacheHits uint32

	// 
	KernelURICacheHitsPercent uint32

	// 
	KernelURICacheHitsPercent_Base uint32

	// 
	KernelUriCacheHitsPersec uint32

	// 
	KernelURICacheMisses uint32

	// 
	MaximumFileCacheMemoryUsage uint64

	// 
	MetadataCacheFlushes uint32

	// 
	MetadataCacheHits uint32

	// 
	MetadataCacheHitsPercent uint32

	// 
	MetadataCacheHitsPercent_Base uint32

	// 
	MetadataCacheMisses uint32

	// 
	OutputCacheCurrentFlushedItems uint32

	// 
	OutputCacheCurrentHitsPercent uint32

	// 
	OutputCacheCurrentHitsPercent_Base uint32

	// 
	OutputCacheCurrentItems uint32

	// 
	OutputCacheCurrentMemoryUsage uint64

	// 
	OutputCacheTotalFlushedItems uint32

	// 
	OutputCacheTotalFlushes uint32

	// 
	OutputCacheTotalHits uint32

	// 
	OutputCacheTotalMisses uint32

	// 
	TotalFilesCached uint32

	// 
	TotalFlushedFiles uint32

	// 
	TotalFlushedMetadata uint32

	// 
	TotalFlushedURIs uint32

	// 
	TotalMetadataCached uint32

	// 
	TotalURIsCached uint32

	// 
	URICacheFlushes uint32

	// 
	URICacheHits uint32

	// 
	URICacheHitsPercent uint32

	// 
	URICacheHitsPercent_Base uint32

	// 
	URICacheMisses uint32
}

	func NewWin32_PerfRawData_W3SVC_WebServiceCacheEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_W3SVC_WebServiceCache, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_W3SVC_WebServiceCache {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_W3SVC_WebServiceCacheEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_W3SVC_WebServiceCache, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_W3SVC_WebServiceCache {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetActiveFlushedEntries sets the value of ActiveFlushedEntries for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyActiveFlushedEntries(value uint32) (err error) { 
    return instance.SetProperty("ActiveFlushedEntries", (value))
}

// GetActiveFlushedEntries gets the value of ActiveFlushedEntries for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyActiveFlushedEntries()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ActiveFlushedEntries")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetCurrentFileCacheMemoryUsage sets the value of CurrentFileCacheMemoryUsage for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyCurrentFileCacheMemoryUsage(value uint64) (err error) { 
    return instance.SetProperty("CurrentFileCacheMemoryUsage", (value))
}

// GetCurrentFileCacheMemoryUsage gets the value of CurrentFileCacheMemoryUsage for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyCurrentFileCacheMemoryUsage()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CurrentFileCacheMemoryUsage")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetCurrentFilesCached sets the value of CurrentFilesCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyCurrentFilesCached(value uint32) (err error) { 
    return instance.SetProperty("CurrentFilesCached", (value))
}

// GetCurrentFilesCached gets the value of CurrentFilesCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyCurrentFilesCached()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentFilesCached")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetCurrentMetadataCached sets the value of CurrentMetadataCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyCurrentMetadataCached(value uint32) (err error) { 
    return instance.SetProperty("CurrentMetadataCached", (value))
}

// GetCurrentMetadataCached gets the value of CurrentMetadataCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyCurrentMetadataCached()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentMetadataCached")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetCurrentURIsCached sets the value of CurrentURIsCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyCurrentURIsCached(value uint32) (err error) { 
    return instance.SetProperty("CurrentURIsCached", (value))
}

// GetCurrentURIsCached gets the value of CurrentURIsCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyCurrentURIsCached()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CurrentURIsCached")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFileCacheFlushes sets the value of FileCacheFlushes for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyFileCacheFlushes(value uint32) (err error) { 
    return instance.SetProperty("FileCacheFlushes", (value))
}

// GetFileCacheFlushes gets the value of FileCacheFlushes for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyFileCacheFlushes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FileCacheFlushes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFileCacheHits sets the value of FileCacheHits for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyFileCacheHits(value uint32) (err error) { 
    return instance.SetProperty("FileCacheHits", (value))
}

// GetFileCacheHits gets the value of FileCacheHits for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyFileCacheHits()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FileCacheHits")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFileCacheHitsPercent sets the value of FileCacheHitsPercent for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyFileCacheHitsPercent(value uint32) (err error) { 
    return instance.SetProperty("FileCacheHitsPercent", (value))
}

// GetFileCacheHitsPercent gets the value of FileCacheHitsPercent for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyFileCacheHitsPercent()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FileCacheHitsPercent")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFileCacheHitsPercent_Base sets the value of FileCacheHitsPercent_Base for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyFileCacheHitsPercent_Base(value uint32) (err error) { 
    return instance.SetProperty("FileCacheHitsPercent_Base", (value))
}

// GetFileCacheHitsPercent_Base gets the value of FileCacheHitsPercent_Base for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyFileCacheHitsPercent_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FileCacheHitsPercent_Base")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFileCacheMisses sets the value of FileCacheMisses for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyFileCacheMisses(value uint32) (err error) { 
    return instance.SetProperty("FileCacheMisses", (value))
}

// GetFileCacheMisses gets the value of FileCacheMisses for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyFileCacheMisses()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FileCacheMisses")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetKernelCurrentURIsCached sets the value of KernelCurrentURIsCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyKernelCurrentURIsCached(value uint32) (err error) { 
    return instance.SetProperty("KernelCurrentURIsCached", (value))
}

// GetKernelCurrentURIsCached gets the value of KernelCurrentURIsCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyKernelCurrentURIsCached()(value uint32, err error) { 
    retValue, err := instance.GetProperty("KernelCurrentURIsCached")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetKernelTotalFlushedURIs sets the value of KernelTotalFlushedURIs for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyKernelTotalFlushedURIs(value uint32) (err error) { 
    return instance.SetProperty("KernelTotalFlushedURIs", (value))
}

// GetKernelTotalFlushedURIs gets the value of KernelTotalFlushedURIs for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyKernelTotalFlushedURIs()(value uint32, err error) { 
    retValue, err := instance.GetProperty("KernelTotalFlushedURIs")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetKernelTotalURIsCached sets the value of KernelTotalURIsCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyKernelTotalURIsCached(value uint32) (err error) { 
    return instance.SetProperty("KernelTotalURIsCached", (value))
}

// GetKernelTotalURIsCached gets the value of KernelTotalURIsCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyKernelTotalURIsCached()(value uint32, err error) { 
    retValue, err := instance.GetProperty("KernelTotalURIsCached")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetKernelURICacheFlushes sets the value of KernelURICacheFlushes for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyKernelURICacheFlushes(value uint32) (err error) { 
    return instance.SetProperty("KernelURICacheFlushes", (value))
}

// GetKernelURICacheFlushes gets the value of KernelURICacheFlushes for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyKernelURICacheFlushes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("KernelURICacheFlushes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetKernelURICacheHits sets the value of KernelURICacheHits for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyKernelURICacheHits(value uint32) (err error) { 
    return instance.SetProperty("KernelURICacheHits", (value))
}

// GetKernelURICacheHits gets the value of KernelURICacheHits for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyKernelURICacheHits()(value uint32, err error) { 
    retValue, err := instance.GetProperty("KernelURICacheHits")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetKernelURICacheHitsPercent sets the value of KernelURICacheHitsPercent for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyKernelURICacheHitsPercent(value uint32) (err error) { 
    return instance.SetProperty("KernelURICacheHitsPercent", (value))
}

// GetKernelURICacheHitsPercent gets the value of KernelURICacheHitsPercent for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyKernelURICacheHitsPercent()(value uint32, err error) { 
    retValue, err := instance.GetProperty("KernelURICacheHitsPercent")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetKernelURICacheHitsPercent_Base sets the value of KernelURICacheHitsPercent_Base for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyKernelURICacheHitsPercent_Base(value uint32) (err error) { 
    return instance.SetProperty("KernelURICacheHitsPercent_Base", (value))
}

// GetKernelURICacheHitsPercent_Base gets the value of KernelURICacheHitsPercent_Base for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyKernelURICacheHitsPercent_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("KernelURICacheHitsPercent_Base")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetKernelUriCacheHitsPersec sets the value of KernelUriCacheHitsPersec for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyKernelUriCacheHitsPersec(value uint32) (err error) { 
    return instance.SetProperty("KernelUriCacheHitsPersec", (value))
}

// GetKernelUriCacheHitsPersec gets the value of KernelUriCacheHitsPersec for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyKernelUriCacheHitsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("KernelUriCacheHitsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetKernelURICacheMisses sets the value of KernelURICacheMisses for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyKernelURICacheMisses(value uint32) (err error) { 
    return instance.SetProperty("KernelURICacheMisses", (value))
}

// GetKernelURICacheMisses gets the value of KernelURICacheMisses for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyKernelURICacheMisses()(value uint32, err error) { 
    retValue, err := instance.GetProperty("KernelURICacheMisses")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMaximumFileCacheMemoryUsage sets the value of MaximumFileCacheMemoryUsage for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyMaximumFileCacheMemoryUsage(value uint64) (err error) { 
    return instance.SetProperty("MaximumFileCacheMemoryUsage", (value))
}

// GetMaximumFileCacheMemoryUsage gets the value of MaximumFileCacheMemoryUsage for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyMaximumFileCacheMemoryUsage()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MaximumFileCacheMemoryUsage")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetMetadataCacheFlushes sets the value of MetadataCacheFlushes for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyMetadataCacheFlushes(value uint32) (err error) { 
    return instance.SetProperty("MetadataCacheFlushes", (value))
}

// GetMetadataCacheFlushes gets the value of MetadataCacheFlushes for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyMetadataCacheFlushes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MetadataCacheFlushes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMetadataCacheHits sets the value of MetadataCacheHits for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyMetadataCacheHits(value uint32) (err error) { 
    return instance.SetProperty("MetadataCacheHits", (value))
}

// GetMetadataCacheHits gets the value of MetadataCacheHits for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyMetadataCacheHits()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MetadataCacheHits")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMetadataCacheHitsPercent sets the value of MetadataCacheHitsPercent for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyMetadataCacheHitsPercent(value uint32) (err error) { 
    return instance.SetProperty("MetadataCacheHitsPercent", (value))
}

// GetMetadataCacheHitsPercent gets the value of MetadataCacheHitsPercent for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyMetadataCacheHitsPercent()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MetadataCacheHitsPercent")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMetadataCacheHitsPercent_Base sets the value of MetadataCacheHitsPercent_Base for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyMetadataCacheHitsPercent_Base(value uint32) (err error) { 
    return instance.SetProperty("MetadataCacheHitsPercent_Base", (value))
}

// GetMetadataCacheHitsPercent_Base gets the value of MetadataCacheHitsPercent_Base for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyMetadataCacheHitsPercent_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MetadataCacheHitsPercent_Base")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMetadataCacheMisses sets the value of MetadataCacheMisses for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyMetadataCacheMisses(value uint32) (err error) { 
    return instance.SetProperty("MetadataCacheMisses", (value))
}

// GetMetadataCacheMisses gets the value of MetadataCacheMisses for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyMetadataCacheMisses()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MetadataCacheMisses")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOutputCacheCurrentFlushedItems sets the value of OutputCacheCurrentFlushedItems for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyOutputCacheCurrentFlushedItems(value uint32) (err error) { 
    return instance.SetProperty("OutputCacheCurrentFlushedItems", (value))
}

// GetOutputCacheCurrentFlushedItems gets the value of OutputCacheCurrentFlushedItems for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyOutputCacheCurrentFlushedItems()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OutputCacheCurrentFlushedItems")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOutputCacheCurrentHitsPercent sets the value of OutputCacheCurrentHitsPercent for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyOutputCacheCurrentHitsPercent(value uint32) (err error) { 
    return instance.SetProperty("OutputCacheCurrentHitsPercent", (value))
}

// GetOutputCacheCurrentHitsPercent gets the value of OutputCacheCurrentHitsPercent for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyOutputCacheCurrentHitsPercent()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OutputCacheCurrentHitsPercent")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOutputCacheCurrentHitsPercent_Base sets the value of OutputCacheCurrentHitsPercent_Base for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyOutputCacheCurrentHitsPercent_Base(value uint32) (err error) { 
    return instance.SetProperty("OutputCacheCurrentHitsPercent_Base", (value))
}

// GetOutputCacheCurrentHitsPercent_Base gets the value of OutputCacheCurrentHitsPercent_Base for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyOutputCacheCurrentHitsPercent_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OutputCacheCurrentHitsPercent_Base")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOutputCacheCurrentItems sets the value of OutputCacheCurrentItems for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyOutputCacheCurrentItems(value uint32) (err error) { 
    return instance.SetProperty("OutputCacheCurrentItems", (value))
}

// GetOutputCacheCurrentItems gets the value of OutputCacheCurrentItems for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyOutputCacheCurrentItems()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OutputCacheCurrentItems")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOutputCacheCurrentMemoryUsage sets the value of OutputCacheCurrentMemoryUsage for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyOutputCacheCurrentMemoryUsage(value uint64) (err error) { 
    return instance.SetProperty("OutputCacheCurrentMemoryUsage", (value))
}

// GetOutputCacheCurrentMemoryUsage gets the value of OutputCacheCurrentMemoryUsage for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyOutputCacheCurrentMemoryUsage()(value uint64, err error) { 
    retValue, err := instance.GetProperty("OutputCacheCurrentMemoryUsage")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetOutputCacheTotalFlushedItems sets the value of OutputCacheTotalFlushedItems for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyOutputCacheTotalFlushedItems(value uint32) (err error) { 
    return instance.SetProperty("OutputCacheTotalFlushedItems", (value))
}

// GetOutputCacheTotalFlushedItems gets the value of OutputCacheTotalFlushedItems for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyOutputCacheTotalFlushedItems()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OutputCacheTotalFlushedItems")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOutputCacheTotalFlushes sets the value of OutputCacheTotalFlushes for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyOutputCacheTotalFlushes(value uint32) (err error) { 
    return instance.SetProperty("OutputCacheTotalFlushes", (value))
}

// GetOutputCacheTotalFlushes gets the value of OutputCacheTotalFlushes for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyOutputCacheTotalFlushes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OutputCacheTotalFlushes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOutputCacheTotalHits sets the value of OutputCacheTotalHits for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyOutputCacheTotalHits(value uint32) (err error) { 
    return instance.SetProperty("OutputCacheTotalHits", (value))
}

// GetOutputCacheTotalHits gets the value of OutputCacheTotalHits for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyOutputCacheTotalHits()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OutputCacheTotalHits")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOutputCacheTotalMisses sets the value of OutputCacheTotalMisses for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyOutputCacheTotalMisses(value uint32) (err error) { 
    return instance.SetProperty("OutputCacheTotalMisses", (value))
}

// GetOutputCacheTotalMisses gets the value of OutputCacheTotalMisses for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyOutputCacheTotalMisses()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OutputCacheTotalMisses")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetTotalFilesCached sets the value of TotalFilesCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyTotalFilesCached(value uint32) (err error) { 
    return instance.SetProperty("TotalFilesCached", (value))
}

// GetTotalFilesCached gets the value of TotalFilesCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyTotalFilesCached()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalFilesCached")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetTotalFlushedFiles sets the value of TotalFlushedFiles for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyTotalFlushedFiles(value uint32) (err error) { 
    return instance.SetProperty("TotalFlushedFiles", (value))
}

// GetTotalFlushedFiles gets the value of TotalFlushedFiles for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyTotalFlushedFiles()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalFlushedFiles")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetTotalFlushedMetadata sets the value of TotalFlushedMetadata for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyTotalFlushedMetadata(value uint32) (err error) { 
    return instance.SetProperty("TotalFlushedMetadata", (value))
}

// GetTotalFlushedMetadata gets the value of TotalFlushedMetadata for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyTotalFlushedMetadata()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalFlushedMetadata")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetTotalFlushedURIs sets the value of TotalFlushedURIs for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyTotalFlushedURIs(value uint32) (err error) { 
    return instance.SetProperty("TotalFlushedURIs", (value))
}

// GetTotalFlushedURIs gets the value of TotalFlushedURIs for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyTotalFlushedURIs()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalFlushedURIs")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetTotalMetadataCached sets the value of TotalMetadataCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyTotalMetadataCached(value uint32) (err error) { 
    return instance.SetProperty("TotalMetadataCached", (value))
}

// GetTotalMetadataCached gets the value of TotalMetadataCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyTotalMetadataCached()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalMetadataCached")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetTotalURIsCached sets the value of TotalURIsCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyTotalURIsCached(value uint32) (err error) { 
    return instance.SetProperty("TotalURIsCached", (value))
}

// GetTotalURIsCached gets the value of TotalURIsCached for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyTotalURIsCached()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TotalURIsCached")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetURICacheFlushes sets the value of URICacheFlushes for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyURICacheFlushes(value uint32) (err error) { 
    return instance.SetProperty("URICacheFlushes", (value))
}

// GetURICacheFlushes gets the value of URICacheFlushes for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyURICacheFlushes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("URICacheFlushes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetURICacheHits sets the value of URICacheHits for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyURICacheHits(value uint32) (err error) { 
    return instance.SetProperty("URICacheHits", (value))
}

// GetURICacheHits gets the value of URICacheHits for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyURICacheHits()(value uint32, err error) { 
    retValue, err := instance.GetProperty("URICacheHits")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetURICacheHitsPercent sets the value of URICacheHitsPercent for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyURICacheHitsPercent(value uint32) (err error) { 
    return instance.SetProperty("URICacheHitsPercent", (value))
}

// GetURICacheHitsPercent gets the value of URICacheHitsPercent for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyURICacheHitsPercent()(value uint32, err error) { 
    retValue, err := instance.GetProperty("URICacheHitsPercent")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetURICacheHitsPercent_Base sets the value of URICacheHitsPercent_Base for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyURICacheHitsPercent_Base(value uint32) (err error) { 
    return instance.SetProperty("URICacheHitsPercent_Base", (value))
}

// GetURICacheHitsPercent_Base gets the value of URICacheHitsPercent_Base for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyURICacheHitsPercent_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("URICacheHitsPercent_Base")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetURICacheMisses sets the value of URICacheMisses for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) SetPropertyURICacheMisses(value uint32) (err error) { 
    return instance.SetProperty("URICacheMisses", (value))
}

// GetURICacheMisses gets the value of URICacheMisses for the instance
func (instance *Win32_PerfRawData_W3SVC_WebServiceCache) GetPropertyURICacheMisses()(value uint32, err error) { 
    retValue, err := instance.GetProperty("URICacheMisses")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

