// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetBranchCachePrimaryCache struct
type MSFT_NetBranchCachePrimaryCache struct {
	MSFT_NetBranchCacheCache

	//
	CurrentActiveCacheSize uint64
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
