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

// MSFT_NetBranchCachePrimaryCache struct
type MSFT_NetBranchCachePrimaryCache struct {
	*MSFT_NetBranchCacheCache

	//
	CurrentActiveCacheSize uint64
}

func NewMSFT_NetBranchCachePrimaryCacheEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCachePrimaryCache, err error) {
	tmp, err := NewMSFT_NetBranchCacheCacheEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCachePrimaryCache{
		MSFT_NetBranchCacheCache: tmp,
	}
	return
}

func NewMSFT_NetBranchCachePrimaryCacheEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCachePrimaryCache, err error) {
	tmp, err := NewMSFT_NetBranchCacheCacheEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCachePrimaryCache{
		MSFT_NetBranchCacheCache: tmp,
	}
	return
}

// SetCurrentActiveCacheSize sets the value of CurrentActiveCacheSize for the instance
func (instance *MSFT_NetBranchCachePrimaryCache) SetPropertyCurrentActiveCacheSize(value uint64) (err error) {
	return instance.SetProperty("CurrentActiveCacheSize", value)
}

// GetCurrentActiveCacheSize gets the value of CurrentActiveCacheSize for the instance
func (instance *MSFT_NetBranchCachePrimaryCache) GetPropertyCurrentActiveCacheSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentActiveCacheSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
