// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetBranchCacheCache struct
type MSFT_NetBranchCacheCache struct {
	*CIM_LogicalElement

	//
	CacheFileDirectoryPath string

	//
	CurrentSizeOnDiskAsNumberOfBytes uint64

	//
	MaxCacheSizeAsNumberOfBytes uint64

	//
	MaxCacheSizeAsPercentageOfDiskVolume uint32
}

func NewMSFT_NetBranchCacheCacheEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheCache, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheCache{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewMSFT_NetBranchCacheCacheEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCacheCache, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheCache{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetCacheFileDirectoryPath sets the value of CacheFileDirectoryPath for the instance
func (instance *MSFT_NetBranchCacheCache) SetPropertyCacheFileDirectoryPath(value string) (err error) {
	return instance.SetProperty("CacheFileDirectoryPath", value)
}

// GetCacheFileDirectoryPath gets the value of CacheFileDirectoryPath for the instance
func (instance *MSFT_NetBranchCacheCache) GetPropertyCacheFileDirectoryPath() (value string, err error) {
	retValue, err := instance.GetProperty("CacheFileDirectoryPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentSizeOnDiskAsNumberOfBytes sets the value of CurrentSizeOnDiskAsNumberOfBytes for the instance
func (instance *MSFT_NetBranchCacheCache) SetPropertyCurrentSizeOnDiskAsNumberOfBytes(value uint64) (err error) {
	return instance.SetProperty("CurrentSizeOnDiskAsNumberOfBytes", value)
}

// GetCurrentSizeOnDiskAsNumberOfBytes gets the value of CurrentSizeOnDiskAsNumberOfBytes for the instance
func (instance *MSFT_NetBranchCacheCache) GetPropertyCurrentSizeOnDiskAsNumberOfBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentSizeOnDiskAsNumberOfBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxCacheSizeAsNumberOfBytes sets the value of MaxCacheSizeAsNumberOfBytes for the instance
func (instance *MSFT_NetBranchCacheCache) SetPropertyMaxCacheSizeAsNumberOfBytes(value uint64) (err error) {
	return instance.SetProperty("MaxCacheSizeAsNumberOfBytes", value)
}

// GetMaxCacheSizeAsNumberOfBytes gets the value of MaxCacheSizeAsNumberOfBytes for the instance
func (instance *MSFT_NetBranchCacheCache) GetPropertyMaxCacheSizeAsNumberOfBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxCacheSizeAsNumberOfBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxCacheSizeAsPercentageOfDiskVolume sets the value of MaxCacheSizeAsPercentageOfDiskVolume for the instance
func (instance *MSFT_NetBranchCacheCache) SetPropertyMaxCacheSizeAsPercentageOfDiskVolume(value uint32) (err error) {
	return instance.SetProperty("MaxCacheSizeAsPercentageOfDiskVolume", value)
}

// GetMaxCacheSizeAsPercentageOfDiskVolume gets the value of MaxCacheSizeAsPercentageOfDiskVolume for the instance
func (instance *MSFT_NetBranchCacheCache) GetPropertyMaxCacheSizeAsPercentageOfDiskVolume() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCacheSizeAsPercentageOfDiskVolume")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
