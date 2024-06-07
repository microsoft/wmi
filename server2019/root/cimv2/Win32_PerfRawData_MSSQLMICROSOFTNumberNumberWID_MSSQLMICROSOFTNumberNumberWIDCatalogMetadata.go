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

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata struct { 
	*Win32_PerfRawData

	// 
	CacheEntriesCount uint64

	// 
	CacheEntriesPinnedCount uint64

	// 
	CacheHitRatio uint64

	// 
	CacheHitRatio_Base uint64
}

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadataEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetCacheEntriesCount sets the value of CacheEntriesCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata) SetPropertyCacheEntriesCount(value uint64) (err error) { 
    return instance.SetProperty("CacheEntriesCount", (value))
}

// GetCacheEntriesCount gets the value of CacheEntriesCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata) GetPropertyCacheEntriesCount()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CacheEntriesCount")
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

// SetCacheEntriesPinnedCount sets the value of CacheEntriesPinnedCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata) SetPropertyCacheEntriesPinnedCount(value uint64) (err error) { 
    return instance.SetProperty("CacheEntriesPinnedCount", (value))
}

// GetCacheEntriesPinnedCount gets the value of CacheEntriesPinnedCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata) GetPropertyCacheEntriesPinnedCount()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CacheEntriesPinnedCount")
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

// SetCacheHitRatio sets the value of CacheHitRatio for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata) SetPropertyCacheHitRatio(value uint64) (err error) { 
    return instance.SetProperty("CacheHitRatio", (value))
}

// GetCacheHitRatio gets the value of CacheHitRatio for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata) GetPropertyCacheHitRatio()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CacheHitRatio")
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

// SetCacheHitRatio_Base sets the value of CacheHitRatio_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata) SetPropertyCacheHitRatio_Base(value uint64) (err error) { 
    return instance.SetProperty("CacheHitRatio_Base", (value))
}

// GetCacheHitRatio_Base gets the value of CacheHitRatio_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCatalogMetadata) GetPropertyCacheHitRatio_Base()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CacheHitRatio_Base")
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

